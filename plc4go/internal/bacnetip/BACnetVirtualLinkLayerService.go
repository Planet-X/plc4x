/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package bacnetip

import (
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type _MultiplexClient struct {
	*Client
	multiplexer *UDPMultiplexer
}

func _New_MultiplexClient(multiplexer *UDPMultiplexer) (*_MultiplexClient, error) {
	m := &_MultiplexClient{
		multiplexer: multiplexer,
	}
	var err error
	m.Client, err = NewClient(nil, m)
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}
	return m, nil
}

func (m *_MultiplexClient) Confirmation(pdu _PDU) error {
	return m.multiplexer.Confirmation(m, pdu)
}

type _MultiplexServer struct {
	*Server
	multiplexer *UDPMultiplexer
}

func _New_MultiplexServer(multiplexer *UDPMultiplexer) (*_MultiplexServer, error) {
	m := &_MultiplexServer{
		multiplexer: multiplexer,
	}
	var err error
	m.Server, err = NewServer(nil, m)
	if err != nil {
		return nil, errors.Wrap(err, "error creating server")
	}
	return m, nil
}

func (m *_MultiplexServer) Indication(pdu _PDU) error {
	return m.multiplexer.Indication(m, pdu)
}

type UDPMultiplexer struct {
	address            *Address
	addrTuple          *AddressTuple[string, uint16]
	addrBroadcastTuple *AddressTuple[string, uint16]
	direct             *_MultiplexClient
	directPort         *UDPDirector
	broadcast          *_MultiplexClient
	broadcastPort      *UDPDirector
	annexH             *_MultiplexServer
	annexJ             *_MultiplexServer
}

func NewUDPMultiplexer(address interface{}, noBroadcast bool) (*UDPMultiplexer, error) {
	log.Debug().Msgf("NewUDPMultiplexer %v noBroadcast=%t", address, noBroadcast)
	u := &UDPMultiplexer{}

	// check for some options
	specialBroadcast := false
	if address == nil {
		address, _ := NewAddress()
		u.address = address
		u.addrTuple = &AddressTuple[string, uint16]{"", 47808}
		u.addrBroadcastTuple = &AddressTuple[string, uint16]{"255.255.255.255", 47808}
	} else {
		// allow the address to be cast
		if caddress, ok := address.(*Address); ok {
			u.address = caddress
		} else if caddress, ok := address.(Address); ok {
			u.address = &caddress
		} else {
			newAddress, err := NewAddress(address)
			if err != nil {
				return nil, errors.Wrap(err, "error parsing address")
			}
			u.address = newAddress
		}

		// promote the normal and broadcast tuples
		u.addrTuple = u.address.AddrTuple
		u.addrBroadcastTuple = u.address.AddrBroadcastTuple

		// check for no broadcasting (loopback interface)
		if u.addrBroadcastTuple == nil {
			noBroadcast = true
		} else if u.addrTuple == u.addrBroadcastTuple {
			// old school broadcast address
			u.addrBroadcastTuple = &AddressTuple[string, uint16]{"255.255.255.255", u.addrTuple.Right}
		} else {
			specialBroadcast = true
		}
	}

	log.Debug().Msgf("address: %v", u.address)
	log.Debug().Msgf("addrTuple: %v", u.addrTuple)
	log.Debug().Msgf("addrBroadcastTuple: %v", u.addrBroadcastTuple)
	log.Debug().Msgf("route_aware: %v", settings.RouteAware)

	// create and bind direct address
	var err error
	u.direct, err = _New_MultiplexClient(u)
	if err != nil {
		return nil, errors.Wrap(err, "error creating multiplex client")
	}
	u.directPort, err = NewUDPDirector(*u.addrTuple, nil, nil, nil, nil)
	if err := bind(u.direct, u.directPort); err != nil {
		return nil, errors.Wrap(err, "error binding ports")
	}

	// create and bind the broadcast address for non-Windows
	if specialBroadcast && !noBroadcast {
		u.broadcast, err = _New_MultiplexClient(u)
		if err != nil {
			return nil, errors.Wrap(err, "error creating broadcast multiplex client")
		}
		reuse := true
		u.broadcastPort, err = NewUDPDirector(*u.addrBroadcastTuple, nil, &reuse, nil, nil)
		if err := bind(u.direct, u.directPort); err != nil {
			return nil, errors.Wrap(err, "error binding ports")
		}
	}

	// create and bind the Annex H and J servers
	u.annexH, err = _New_MultiplexServer(u)
	if err != nil {
		return nil, errors.Wrap(err, "error creating annexH")
	}
	u.annexJ, err = _New_MultiplexServer(u)
	if err != nil {
		return nil, errors.Wrap(err, "error creating annexJ")
	}
	return u, nil
}

func (m *UDPMultiplexer) Close() error {
	log.Debug().Msg("Close")

	// pass along the close to the director(s)
	m.directPort.Close()
	if m.broadcastPort != nil {
		m.broadcastPort.Close()
	}
	return nil
}

func (m *UDPMultiplexer) Indication(server *_MultiplexServer, pdu _PDU) error {
	log.Debug().Msgf("Indication %v\n%v", server, pdu)

	pduDestination := pdu.GetPDUDestination()

	// broadcast message
	var dest *Address
	if pduDestination.AddrType == LOCAL_BROADCAST_ADDRESS {
		// interface might not support broadcasts
		if m.addrBroadcastTuple == nil {
			return nil
		}

		address, err := NewAddress(*m.addrBroadcastTuple)
		if err != nil {
			return errors.Wrap(err, "error getting address from tuple")
		}
		dest = address
		log.Debug().Msgf("requesting local broadcast: %v", dest)
	} else if pduDestination.AddrType == LOCAL_STATION_ADDRESS {
		dest = pduDestination
	} else {
		return errors.New("invalid destination address type")
	}

	return m.directPort.Indication(NewPDUFromPDU(pdu, WithPDUDestination(dest)))
}

func (m *UDPMultiplexer) Confirmation(client *_MultiplexClient, pdu _PDU) error {
	log.Debug().Msgf("Confirmation %v\n%v", client, pdu)
	log.Debug().Msgf("client address: %v", client.multiplexer.address)

	// if this came from ourselves, dump it
	pduSource := pdu.GetPDUSource()
	if pduSource.Equals(m.address) {
		log.Debug().Msg("from us")
		return nil
	}

	// TODO: upstream this is a tuple but we don't have that here so we can work with what we have
	src := pduSource
	var dest *Address

	// match the destination in case the stack needs it
	if client == m.direct {
		log.Debug().Msg("direct to us")
		dest = m.address
	} else if client == m.broadcast {
		log.Debug().Msg("broadcast to us")
		dest = NewLocalBroadcast(nil)
	} else {
		return errors.New("Confirmation missmatch")
	}
	log.Debug().Msgf("dest: %s", dest)

	// must have at least one octet
	if pdu.GetMessage() == nil {
		log.Debug().Msg("no data")
		return nil
	}

	// TODO: we only support 0x81 at the moment
	if m.annexJ != nil {
		return m.annexJ.Response(NewPDU(pdu.GetMessage(), WithPDUSource(src), WithPDUDestination(dest)))
	}

	return nil
}

type AnnexJCodec struct {
	*Client
	*Server
}

func NewAnnexJCodec(cid *int, sid *int) (*AnnexJCodec, error) {
	log.Debug().Msgf("NewAnnexJCodec cid=%d sid=%d", cid, sid)
	a := &AnnexJCodec{}
	client, err := NewClient(cid, a)
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}
	a.Client = client
	server, err := NewServer(sid, a)
	if err != nil {
		return nil, errors.Wrap(err, "error creating server")
	}
	a.Server = server
	return a, nil
}

func (b *AnnexJCodec) Indication(pdu _PDU) error {
	// Note: our BVLC are all annexJ at the moment
	return b.Request(pdu)
}

func (b *AnnexJCodec) Confirmation(pdu _PDU) error {
	// Note: our BVLC are all annexJ at the moment
	return b.Response(pdu)
}

type _BIPSAP interface {
	_ServiceAccessPoint
	_Client
}

type BIPSAP struct {
	*ServiceAccessPoint
	rootStruct _BIPSAP
}

func NewBIPSAP(sapID *int, rootStruct _BIPSAP) (*BIPSAP, error) {
	log.Debug().Msgf("NewBIPSAP sapID=%d", sapID)
	b := &BIPSAP{}
	serviceAccessPoint, err := NewServiceAccessPoint(sapID, rootStruct)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating service access point")
	}
	b.ServiceAccessPoint = serviceAccessPoint
	b.rootStruct = rootStruct
	return b, nil
}

func (b *BIPSAP) SapIndication(pdu _PDU) error {
	log.Debug().Msgf("SapIndication\n%ss", pdu)
	// this is a request initiated by the ASE, send this downstream
	return b.rootStruct.Request(pdu)
}

func (b *BIPSAP) SapConfirmation(pdu _PDU) error {
	log.Debug().Msgf("SapConfirmation\n%s", pdu)
	// this is a response from the ASE, send this downstream
	return b.rootStruct.Request(pdu)
}

type BIPSimple struct {
	*BIPSAP
	*Client
	*Server
}

func NewBIPSimple(sapID *int, cid *int, sid *int) (*BIPSimple, error) {
	log.Debug().Msgf("NewBIPSimple sapID=%d cid=%d sid=%d", sapID, cid, sid)
	b := &BIPSimple{}
	bipsap, err := NewBIPSAP(sapID, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating bisap")
	}
	b.BIPSAP = bipsap
	client, err := NewClient(cid, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}
	b.Client = client
	server, err := NewServer(sid, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating server")
	}
	b.Server = server
	return b, nil
}

func (b *BIPSimple) Indication(pdu _PDU) error {
	log.Debug().Msgf("Indication %s", pdu)

	// check for local stations
	switch pdu.GetPDUDestination().AddrType {
	case LOCAL_STATION_ADDRESS:
		// make an original unicast PDU
		xpdu := readWriteModel.NewBVLCOriginalUnicastNPDU(pdu.GetMessage().(readWriteModel.NPDU), 0)
		log.Debug().Msgf("xpdu:\n%s", xpdu)

		// send it downstream
		return b.Request(NewPDUFromPDUWithNewMessage(pdu, xpdu))
	case LOCAL_BROADCAST_ADDRESS:
		// make an original broadcast PDU
		xpdu := readWriteModel.NewBVLCOriginalBroadcastNPDU(pdu.GetMessage().(readWriteModel.NPDU), 0)

		log.Debug().Msgf("xpdu:\n%s", xpdu)

		// send it downstream
		return b.Request(NewPDUFromPDUWithNewMessage(pdu, xpdu))
	default:
		return errors.Errorf("invalid destination address: %s", pdu.GetPDUDestination())
	}
}

func (b *BIPSimple) Confirmation(pdu _PDU) error {
	log.Debug().Msgf("Confirmation %s", pdu)

	switch msg := pdu.GetMessage().(type) {
	// some kind of response to a request
	case readWriteModel.BVLCResultExactly:
		// send this to the service access point
		return b.SapResponse(pdu)
	case readWriteModel.BVLCReadBroadcastDistributionTableAckExactly:
		// send this to the service access point
		return b.SapResponse(pdu)
	case readWriteModel.BVLCReadForeignDeviceTableAckExactly:
		// send this to the service access point
		return b.SapResponse(pdu)
	case readWriteModel.BVLCOriginalUnicastNPDUExactly:
		// build a vanilla PDU
		xpdu := NewPDU(msg.GetNpdu(), WithPDUSource(pdu.GetPDUSource()), WithPDUDestination(pdu.GetPDUDestination()))
		log.Debug().Msgf("xpdu: %s", xpdu)

		// send it upstream
		return b.Response(xpdu)
	case readWriteModel.BVLCOriginalBroadcastNPDUExactly:
		// build a PDU with a local broadcast address
		xpdu := NewPDU(msg.GetNpdu(), WithPDUSource(pdu.GetPDUSource()), WithPDUDestination(NewLocalBroadcast(nil)))
		log.Debug().Msgf("xpdu: %s", xpdu)

		// send it upstream
		return b.Response(xpdu)
	case readWriteModel.BVLCForwardedNPDUExactly:
		// build a PDU with the source from the real source
		ip := msg.GetIp()
		port := msg.GetPort()
		source, err := NewAddress(append(ip, uint16ToPort(port)...))
		if err != nil {
			return errors.Wrap(err, "error building a ip")
		}
		xpdu := NewPDU(msg.GetNpdu(), WithPDUSource(source), WithPDUDestination(NewLocalBroadcast(nil)))
		log.Debug().Msgf("xpdu: %s", xpdu)

		// send it upstream
		return b.Response(xpdu)
	case readWriteModel.BVLCWriteBroadcastDistributionTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCReadBroadcastDistributionTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
		// build a response
	case readWriteModel.BVLCRegisterForeignDeviceExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCReadForeignDeviceTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCDeleteForeignDeviceTableEntryExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCDistributeBroadcastToNetworkExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	default:
		log.Warn().Msgf("invalid pdu type %T", msg)
		return nil
	}
}

type BIPForeign struct {
	*BIPSAP
	*Client
	*Server
	*OneShotTask
	registrationStatus      int
	bbmdAddress             *Address
	bbmdTimeToLive          *int
	registrationTimeoutTask *OneShotFunctionTask
}

func NewBIPForeign(addr *Address, ttl *int, sapID *int, cid *int, sid *int) (*BIPForeign, error) {
	log.Debug().Msgf("NewBIPForeign addr=%s ttl=%d sapID=%d cid=%d sid=%d", addr, ttl, sapID, cid, sid)
	b := &BIPForeign{}
	bipsap, err := NewBIPSAP(sapID, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating bisap")
	}
	b.BIPSAP = bipsap
	client, err := NewClient(cid, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}
	b.Client = client
	server, err := NewServer(sid, b)
	if err != nil {
		return nil, errors.Wrap(err, "error creating server")
	}
	b.Server = server
	b.OneShotTask = NewOneShotTask(b, nil)

	// -2=unregistered, -1=not attempted or no ack, 0=OK, >0 error
	b.registrationStatus = -1

	// clear the BBMD address and time-to-live
	b.bbmdAddress = nil
	b.bbmdTimeToLive = nil

	// used in tracking active registration timeouts
	b.registrationTimeoutTask = OneShotFunction(b._registration_expired)

	// registration provided
	if addr != nil {
		// a little error checking
		if ttl == nil {
			return nil, errors.New("BBMD address and time-to-live must both be specified")
		}

		if err := b.register(*addr, *ttl); err != nil {
			return nil, errors.Wrap(err, "error registering")
		}
	}

	return b, nil
}

func (b *BIPForeign) Indication(pdu _PDU) error {
	log.Debug().Msgf("Indication %s", pdu)

	// check for local stations
	switch pdu.GetPDUDestination().AddrType {
	case LOCAL_STATION_ADDRESS:
		// make an original unicast PDU
		xpdu := readWriteModel.NewBVLCOriginalUnicastNPDU(pdu.GetMessage().(readWriteModel.NPDU), 0)
		log.Debug().Msgf("xpdu:\n%s", xpdu)

		// send it downstream
		return b.Request(NewPDUFromPDUWithNewMessage(pdu, xpdu))
	case LOCAL_BROADCAST_ADDRESS:
		// check the BBMD registration status, we may not be registered
		if b.registrationStatus != 0 {
			log.Debug().Msg("packet dropped, unregistered")
			return nil
		}

		// make an original broadcast PDU
		xpdu := readWriteModel.NewBVLCOriginalBroadcastNPDU(pdu.GetMessage().(readWriteModel.NPDU), 0)

		log.Debug().Msgf("xpdu:\n%s", xpdu)

		// send it downstream
		return b.Request(NewPDUFromPDUWithNewMessage(pdu, xpdu))
	default:
		return errors.Errorf("invalid destination address: %s", pdu.GetPDUDestination())
	}
}

func (b *BIPForeign) Confirmation(pdu _PDU) error {
	log.Debug().Msgf("Confirmation %s", pdu)

	switch msg := pdu.GetMessage().(type) {
	// check for a registration request result
	case readWriteModel.BVLCResultExactly:
		// if we are unbinding, do nothing
		if b.registrationStatus == -2 {
			return nil
		}

		// make sure we have a bind request in process

		// make sure the result is from the bbmd

		if !pdu.GetPDUSource().Equals(b.bbmdAddress) {
			log.Debug().Msg("packet dropped, not from the BBMD")
			return nil
		}
		// save the result code as the status
		b.registrationStatus = int(msg.GetCode())

		// If successful, track registration timeout
		if b.registrationStatus == 0 {
			b._start_track_registration()
		}

		return nil
	case readWriteModel.BVLCOriginalUnicastNPDUExactly:
		// build a vanilla PDU
		xpdu := NewPDU(msg.GetNpdu(), WithPDUSource(pdu.GetPDUSource()), WithPDUDestination(pdu.GetPDUDestination()))
		log.Debug().Msgf("xpdu: %s", xpdu)

		// send it upstream
		return b.Response(xpdu)
	case readWriteModel.BVLCForwardedNPDUExactly:
		// check the BBMD registration status, we may not be registered
		if b.registrationStatus != 0 {
			log.Debug().Msg("packet dropped, unregistered")
			return nil
		}

		// make sure the forwarded PDU from the bbmd
		if !pdu.GetPDUSource().Equals(b.bbmdAddress) {
			log.Debug().Msg("packet dropped, not from the BBMD")
			return nil
		}

		// build a PDU with the source from the real source
		ip := msg.GetIp()
		port := msg.GetPort()
		source, err := NewAddress(append(ip, uint16ToPort(port)...))
		if err != nil {
			return errors.Wrap(err, "error building a ip")
		}
		xpdu := NewPDU(msg.GetNpdu(), WithPDUSource(source), WithPDUDestination(NewLocalBroadcast(nil)))
		log.Debug().Msgf("xpdu: %s", xpdu)

		// send it upstream
		return b.Response(xpdu)
	case readWriteModel.BVLCReadBroadcastDistributionTableAckExactly:
		// send this to the service access point
		return b.SapResponse(pdu)
	case readWriteModel.BVLCReadForeignDeviceTableAckExactly:
		// send this to the service access point
		return b.SapResponse(pdu)
	case readWriteModel.BVLCWriteBroadcastDistributionTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCReadBroadcastDistributionTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCRegisterForeignDeviceExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCReadForeignDeviceTableExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCDeleteForeignDeviceTableEntryExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCDistributeBroadcastToNetworkExactly:
		// build a response
		result := readWriteModel.NewBVLCResult(readWriteModel.BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK)
		xpdu := NewPDU(result, WithPDUDestination(pdu.GetPDUSource()))

		// send it downstream
		return b.Request(xpdu)
	case readWriteModel.BVLCOriginalBroadcastNPDUExactly:
		log.Debug().Msg("packet dropped")
		return nil
	default:
		log.Warn().Msgf("invalid pdu type %T", msg)
		return nil
	}
}

// register starts the foreign device registration process with the given BBMD.
//
//        Registration will be renewed periodically according to the ttl value
//        until explicitly stopped by a call to `unregister`.
func (b *BIPForeign) register(addr Address, ttl int) error {
	// a little error checking
	if ttl <= 0 {
		return errors.New("time-to-live must be greater than zero")
	}

	// save the BBMD address and time-to-live
	b.bbmdAddress = &addr
	b.bbmdTimeToLive = &ttl

	// install this task to do registration renewal according to the TTL
	// and stop tracking any active registration timeouts
	var taskTime time.Time
	b.InstallTask(&taskTime, nil)
	b._stop_track_registration()
	return nil
}

// unregister stops the foreign device registration process.
//
// Immediately drops active foreign device registration and stops further
// registration renewals.
func (b *BIPForeign) unregister() {
	pdu := NewPDU(readWriteModel.NewBVLCRegisterForeignDevice(0), WithPDUDestination(b.bbmdAddress))

	// send it downstream
	if err := b.Request(pdu); err != nil {
		log.Debug().Err(err).Msg("error sending request")
		return
	}

	// change the status to unregistered
	b.registrationStatus = -2

	// clear the BBMD address and time-to-live
	b.bbmdAddress = nil
	b.bbmdTimeToLive = nil

	// unschedule registration renewal & timeout tracking if previously
	// scheduled
	b.SuspendTask()
	b._stop_track_registration()
}

// processTask is called when the registration request should be sent to the BBMD.
func (b *BIPForeign) processTask() error {
	pdu := NewPDU(readWriteModel.NewBVLCRegisterForeignDevice(uint16(*b.bbmdTimeToLive)), WithPDUDestination(b.bbmdAddress))

	// send it downstream
	if err := b.Request(pdu); err != nil {
		return errors.Wrap(err, "error sending request")
	}

	// schedule the next registration renewal
	var delta = time.Duration(*b.bbmdTimeToLive) * time.Second
	b.InstallTask(nil, &delta)
	return nil
}

// _start_track_registration From J.5.2.3 Foreign Device Table Operation (paraphrasing): if a
// foreign device does not renew its registration 30 seconds after its
// TTL expired then it will be removed from the BBMD's FDT.
//
// Thus, if we're registered and don't get a response to a subsequent
// renewal request 30 seconds after our TTL expired then we're
// definitely not registered anymore.
func (b *BIPForeign) _start_track_registration() {
	var delta = time.Duration(*b.bbmdTimeToLive)*time.Second + (30 * time.Second)
	b.registrationTimeoutTask.InstallTask(nil, &delta)
}

func (b *BIPForeign) _stop_track_registration() {
	b.registrationTimeoutTask.SuspendTask()
}

// _registration_expired is called when detecting that foreign device registration has definitely expired.
func (b *BIPForeign) _registration_expired() error {
	b.registrationStatus = -1 // Unregistered
	b._stop_track_registration()
	return nil
}
