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

// BACnetConstructedDataValueChangeTime is the corresponding interface of BACnetConstructedDataValueChangeTime
type BACnetConstructedDataValueChangeTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetValueChangeTime returns ValueChangeTime (property field)
	GetValueChangeTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataValueChangeTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataValueChangeTime.
// This is useful for switch cases.
type BACnetConstructedDataValueChangeTimeExactly interface {
	BACnetConstructedDataValueChangeTime
	isBACnetConstructedDataValueChangeTime() bool
}

// _BACnetConstructedDataValueChangeTime is the data-structure of this message
type _BACnetConstructedDataValueChangeTime struct {
	*_BACnetConstructedData
	ValueChangeTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValueChangeTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_CHANGE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueChangeTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataValueChangeTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetValueChangeTime() BACnetDateTime {
	return m.ValueChangeTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetValueChangeTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValueChangeTime factory function for _BACnetConstructedDataValueChangeTime
func NewBACnetConstructedDataValueChangeTime(valueChangeTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValueChangeTime {
	_result := &_BACnetConstructedDataValueChangeTime{
		ValueChangeTime:        valueChangeTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueChangeTime(structType interface{}) BACnetConstructedDataValueChangeTime {
	if casted, ok := structType.(BACnetConstructedDataValueChangeTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueChangeTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueChangeTime) GetTypeName() string {
	return "BACnetConstructedDataValueChangeTime"
}

func (m *_BACnetConstructedDataValueChangeTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (valueChangeTime)
	lengthInBits += m.ValueChangeTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValueChangeTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataValueChangeTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValueChangeTime, error) {
	return BACnetConstructedDataValueChangeTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataValueChangeTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValueChangeTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueChangeTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (valueChangeTime)
	if pullErr := readBuffer.PullContext("valueChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueChangeTime")
	}
	_valueChangeTime, _valueChangeTimeErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _valueChangeTimeErr != nil {
		return nil, errors.Wrap(_valueChangeTimeErr, "Error parsing 'valueChangeTime' field of BACnetConstructedDataValueChangeTime")
	}
	valueChangeTime := _valueChangeTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("valueChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueChangeTime")
	}

	// Virtual field
	_actualValue := valueChangeTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueChangeTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataValueChangeTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ValueChangeTime: valueChangeTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataValueChangeTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValueChangeTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueChangeTime")
		}

		// Simple Field (valueChangeTime)
		if pushErr := writeBuffer.PushContext("valueChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for valueChangeTime")
		}
		_valueChangeTimeErr := writeBuffer.WriteSerializable(ctx, m.GetValueChangeTime())
		if popErr := writeBuffer.PopContext("valueChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for valueChangeTime")
		}
		if _valueChangeTimeErr != nil {
			return errors.Wrap(_valueChangeTimeErr, "Error serializing 'valueChangeTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueChangeTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValueChangeTime) isBACnetConstructedDataValueChangeTime() bool {
	return true
}

func (m *_BACnetConstructedDataValueChangeTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
