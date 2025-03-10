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

// BACnetConstructedDataDoorUnlockDelayTime is the corresponding interface of BACnetConstructedDataDoorUnlockDelayTime
type BACnetConstructedDataDoorUnlockDelayTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorUnlockDelayTime returns DoorUnlockDelayTime (property field)
	GetDoorUnlockDelayTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDoorUnlockDelayTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorUnlockDelayTime.
// This is useful for switch cases.
type BACnetConstructedDataDoorUnlockDelayTimeExactly interface {
	BACnetConstructedDataDoorUnlockDelayTime
	isBACnetConstructedDataDoorUnlockDelayTime() bool
}

// _BACnetConstructedDataDoorUnlockDelayTime is the data-structure of this message
type _BACnetConstructedDataDoorUnlockDelayTime struct {
	*_BACnetConstructedData
	DoorUnlockDelayTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_UNLOCK_DELAY_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetDoorUnlockDelayTime() BACnetApplicationTagUnsignedInteger {
	return m.DoorUnlockDelayTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDoorUnlockDelayTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorUnlockDelayTime factory function for _BACnetConstructedDataDoorUnlockDelayTime
func NewBACnetConstructedDataDoorUnlockDelayTime(doorUnlockDelayTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorUnlockDelayTime {
	_result := &_BACnetConstructedDataDoorUnlockDelayTime{
		DoorUnlockDelayTime:    doorUnlockDelayTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorUnlockDelayTime(structType interface{}) BACnetConstructedDataDoorUnlockDelayTime {
	if casted, ok := structType.(BACnetConstructedDataDoorUnlockDelayTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorUnlockDelayTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetTypeName() string {
	return "BACnetConstructedDataDoorUnlockDelayTime"
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (doorUnlockDelayTime)
	lengthInBits += m.DoorUnlockDelayTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDoorUnlockDelayTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorUnlockDelayTime, error) {
	return BACnetConstructedDataDoorUnlockDelayTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDoorUnlockDelayTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorUnlockDelayTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorUnlockDelayTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorUnlockDelayTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorUnlockDelayTime)
	if pullErr := readBuffer.PullContext("doorUnlockDelayTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorUnlockDelayTime")
	}
	_doorUnlockDelayTime, _doorUnlockDelayTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _doorUnlockDelayTimeErr != nil {
		return nil, errors.Wrap(_doorUnlockDelayTimeErr, "Error parsing 'doorUnlockDelayTime' field of BACnetConstructedDataDoorUnlockDelayTime")
	}
	doorUnlockDelayTime := _doorUnlockDelayTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("doorUnlockDelayTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorUnlockDelayTime")
	}

	// Virtual field
	_actualValue := doorUnlockDelayTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorUnlockDelayTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorUnlockDelayTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoorUnlockDelayTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoorUnlockDelayTime: doorUnlockDelayTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorUnlockDelayTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorUnlockDelayTime")
		}

		// Simple Field (doorUnlockDelayTime)
		if pushErr := writeBuffer.PushContext("doorUnlockDelayTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorUnlockDelayTime")
		}
		_doorUnlockDelayTimeErr := writeBuffer.WriteSerializable(ctx, m.GetDoorUnlockDelayTime())
		if popErr := writeBuffer.PopContext("doorUnlockDelayTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorUnlockDelayTime")
		}
		if _doorUnlockDelayTimeErr != nil {
			return errors.Wrap(_doorUnlockDelayTimeErr, "Error serializing 'doorUnlockDelayTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorUnlockDelayTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorUnlockDelayTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) isBACnetConstructedDataDoorUnlockDelayTime() bool {
	return true
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
