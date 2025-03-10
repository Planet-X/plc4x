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

package model

import (
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// APDU is the corresponding interface of APDU
type APDU interface {
	utils.LengthAware
	utils.Serializable
	// GetApduType returns ApduType (discriminator field)
	GetApduType() ApduType
}

// APDUExactly can be used when we want exactly this type and not a type which fulfills APDU.
// This is useful for switch cases.
type APDUExactly interface {
	APDU
	isAPDU() bool
}

// _APDU is the data-structure of this message
type _APDU struct {
	_APDUChildRequirements

	// Arguments.
	ApduLength uint16
}

type _APDUChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetApduType() ApduType
}

type APDUParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child APDU, serializeChildFunction func() error) error
	GetTypeName() string
}

type APDUChild interface {
	utils.Serializable
	InitializeParent(parent APDU)
	GetParent() *APDU

	GetTypeName() string
	APDU
}

// NewAPDU factory function for _APDU
func NewAPDU(apduLength uint16) *_APDU {
	return &_APDU{ApduLength: apduLength}
}

// Deprecated: use the interface for direct cast
func CastAPDU(structType interface{}) APDU {
	if casted, ok := structType.(APDU); ok {
		return casted
	}
	if casted, ok := structType.(*APDU); ok {
		return *casted
	}
	return nil
}

func (m *_APDU) GetTypeName() string {
	return "APDU"
}

func (m *_APDU) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apduType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_APDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUParse(theBytes []byte, apduLength uint16) (APDU, error) {
	return APDUParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (apduType) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("apduType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apduType")
	}
	apduType_temp, _apduTypeErr := ApduTypeParseWithBuffer(ctx, readBuffer)
	var apduType ApduType = apduType_temp
	if closeErr := readBuffer.CloseContext("apduType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apduType")
	}
	if _apduTypeErr != nil {
		return nil, errors.Wrap(_apduTypeErr, "Error parsing 'apduType' field of APDU")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type APDUChildSerializeRequirement interface {
		APDU
		InitializeParent(APDU)
		GetParent() APDU
	}
	var _childTemp interface{}
	var _child APDUChildSerializeRequirement
	var typeSwitchError error
	switch {
	case apduType == ApduType_CONFIRMED_REQUEST_PDU: // APDUConfirmedRequest
		_childTemp, typeSwitchError = APDUConfirmedRequestParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_UNCONFIRMED_REQUEST_PDU: // APDUUnconfirmedRequest
		_childTemp, typeSwitchError = APDUUnconfirmedRequestParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_SIMPLE_ACK_PDU: // APDUSimpleAck
		_childTemp, typeSwitchError = APDUSimpleAckParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_COMPLEX_ACK_PDU: // APDUComplexAck
		_childTemp, typeSwitchError = APDUComplexAckParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_SEGMENT_ACK_PDU: // APDUSegmentAck
		_childTemp, typeSwitchError = APDUSegmentAckParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_ERROR_PDU: // APDUError
		_childTemp, typeSwitchError = APDUErrorParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_REJECT_PDU: // APDUReject
		_childTemp, typeSwitchError = APDURejectParseWithBuffer(ctx, readBuffer, apduLength)
	case apduType == ApduType_ABORT_PDU: // APDUAbort
		_childTemp, typeSwitchError = APDUAbortParseWithBuffer(ctx, readBuffer, apduLength)
	case 0 == 0: // APDUUnknown
		_childTemp, typeSwitchError = APDUUnknownParseWithBuffer(ctx, readBuffer, apduLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [apduType=%v]", apduType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of APDU")
	}
	_child = _childTemp.(APDUChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("APDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDU")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_APDU) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child APDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("APDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for APDU")
	}

	// Discriminator Field (apduType) (Used as input to a switch field)
	apduType := ApduType(child.GetApduType())
	if pushErr := writeBuffer.PushContext("apduType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for apduType")
	}
	_apduTypeErr := writeBuffer.WriteSerializable(ctx, apduType)
	if popErr := writeBuffer.PopContext("apduType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for apduType")
	}

	if _apduTypeErr != nil {
		return errors.Wrap(_apduTypeErr, "Error serializing 'apduType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("APDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for APDU")
	}
	return nil
}

////
// Arguments Getter

func (m *_APDU) GetApduLength() uint16 {
	return m.ApduLength
}

//
////

func (m *_APDU) isAPDU() bool {
	return true
}

func (m *_APDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
