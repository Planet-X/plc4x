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

// BACnetConfirmedServiceRequestConfirmedCOVNotification is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotification
type BACnetConfirmedServiceRequestConfirmedCOVNotification interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() BACnetContextTagUnsignedInteger
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
}

// BACnetConfirmedServiceRequestConfirmedCOVNotificationExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedCOVNotification.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedCOVNotificationExactly interface {
	BACnetConfirmedServiceRequestConfirmedCOVNotification
	isBACnetConfirmedServiceRequestConfirmedCOVNotification() bool
}

// _BACnetConfirmedServiceRequestConfirmedCOVNotification is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedCOVNotification struct {
	*_BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	LifetimeInSeconds           BACnetContextTagUnsignedInteger
	ListOfValues                BACnetPropertyValues
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLifetimeInSeconds() BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotification factory function for _BACnetConfirmedServiceRequestConfirmedCOVNotification
func NewBACnetConfirmedServiceRequestConfirmedCOVNotification(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, lifetimeInSeconds BACnetContextTagUnsignedInteger, listOfValues BACnetPropertyValues, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedCOVNotification {
	_result := &_BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:    subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:     initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:      monitoredObjectIdentifier,
		LifetimeInSeconds:              lifetimeInSeconds,
		ListOfValues:                   listOfValues,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedCOVNotification(structType interface{}) BACnetConfirmedServiceRequestConfirmedCOVNotification {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotification"
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits(ctx)

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedCOVNotification, error) {
	return BACnetConfirmedServiceRequestConfirmedCOVNotificationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedCOVNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriberProcessIdentifier")
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field of BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	subscriberProcessIdentifier := _subscriberProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriberProcessIdentifier")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initiatingDeviceIdentifier")
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field of BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	initiatingDeviceIdentifier := _initiatingDeviceIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initiatingDeviceIdentifier")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field of BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	monitoredObjectIdentifier := _monitoredObjectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredObjectIdentifier")
	}

	// Simple Field (lifetimeInSeconds)
	if pullErr := readBuffer.PullContext("lifetimeInSeconds"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lifetimeInSeconds")
	}
	_lifetimeInSeconds, _lifetimeInSecondsErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lifetimeInSecondsErr != nil {
		return nil, errors.Wrap(_lifetimeInSecondsErr, "Error parsing 'lifetimeInSeconds' field of BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	lifetimeInSeconds := _lifetimeInSeconds.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lifetimeInSeconds"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lifetimeInSeconds")
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParseWithBuffer(ctx, readBuffer, uint8(uint8(4)), BACnetObjectType(monitoredObjectIdentifier.GetObjectType()))
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field of BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	listOfValues := _listOfValues.(BACnetPropertyValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedCOVNotification{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SubscriberProcessIdentifier: subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:  initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:   monitoredObjectIdentifier,
		LifetimeInSeconds:           lifetimeInSeconds,
		ListOfValues:                listOfValues,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedCOVNotification")
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
		}
		_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetSubscriberProcessIdentifier())
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for initiatingDeviceIdentifier")
		}
		_initiatingDeviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetInitiatingDeviceIdentifier())
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for initiatingDeviceIdentifier")
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
		}
		_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoredObjectIdentifier())
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Simple Field (lifetimeInSeconds)
		if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lifetimeInSeconds")
		}
		_lifetimeInSecondsErr := writeBuffer.WriteSerializable(ctx, m.GetLifetimeInSeconds())
		if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lifetimeInSeconds")
		}
		if _lifetimeInSecondsErr != nil {
			return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
		}

		// Simple Field (listOfValues)
		if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfValues")
		}
		_listOfValuesErr := writeBuffer.WriteSerializable(ctx, m.GetListOfValues())
		if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfValues")
		}
		if _listOfValuesErr != nil {
			return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedCOVNotification")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) isBACnetConfirmedServiceRequestConfirmedCOVNotification() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
