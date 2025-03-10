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

// BACnetConstructedDataAccessEvent is the corresponding interface of BACnetConstructedDataAccessEvent
type BACnetConstructedDataAccessEvent interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() BACnetAccessEventTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessEventTagged
}

// BACnetConstructedDataAccessEventExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessEvent.
// This is useful for switch cases.
type BACnetConstructedDataAccessEventExactly interface {
	BACnetConstructedDataAccessEvent
	isBACnetConstructedDataAccessEvent() bool
}

// _BACnetConstructedDataAccessEvent is the data-structure of this message
type _BACnetConstructedDataAccessEvent struct {
	*_BACnetConstructedData
	AccessEvent BACnetAccessEventTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEvent) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEvent) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessEvent) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetAccessEvent() BACnetAccessEventTagged {
	return m.AccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetActualValue() BACnetAccessEventTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessEventTagged(m.GetAccessEvent())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEvent factory function for _BACnetConstructedDataAccessEvent
func NewBACnetConstructedDataAccessEvent(accessEvent BACnetAccessEventTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEvent {
	_result := &_BACnetConstructedDataAccessEvent{
		AccessEvent:            accessEvent,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEvent(structType interface{}) BACnetConstructedDataAccessEvent {
	if casted, ok := structType.(BACnetConstructedDataAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEvent) GetTypeName() string {
	return "BACnetConstructedDataAccessEvent"
}

func (m *_BACnetConstructedDataAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessEventParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEvent, error) {
	return BACnetConstructedDataAccessEventParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEvent)
	if pullErr := readBuffer.PullContext("accessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEvent")
	}
	_accessEvent, _accessEventErr := BACnetAccessEventTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _accessEventErr != nil {
		return nil, errors.Wrap(_accessEventErr, "Error parsing 'accessEvent' field of BACnetConstructedDataAccessEvent")
	}
	accessEvent := _accessEvent.(BACnetAccessEventTagged)
	if closeErr := readBuffer.CloseContext("accessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEvent")
	}

	// Virtual field
	_actualValue := accessEvent
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessEvent{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AccessEvent: accessEvent,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEvent")
		}

		// Simple Field (accessEvent)
		if pushErr := writeBuffer.PushContext("accessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEvent")
		}
		_accessEventErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEvent())
		if popErr := writeBuffer.PopContext("accessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEvent")
		}
		if _accessEventErr != nil {
			return errors.Wrap(_accessEventErr, "Error serializing 'accessEvent' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEvent) isBACnetConstructedDataAccessEvent() bool {
	return true
}

func (m *_BACnetConstructedDataAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
