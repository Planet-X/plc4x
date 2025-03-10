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

// BACnetLogRecordLogDatumRealValue is the corresponding interface of BACnetLogRecordLogDatumRealValue
type BACnetLogRecordLogDatumRealValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetContextTagReal
}

// BACnetLogRecordLogDatumRealValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumRealValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumRealValueExactly interface {
	BACnetLogRecordLogDatumRealValue
	isBACnetLogRecordLogDatumRealValue() bool
}

// _BACnetLogRecordLogDatumRealValue is the data-structure of this message
type _BACnetLogRecordLogDatumRealValue struct {
	*_BACnetLogRecordLogDatum
	RealValue BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumRealValue) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumRealValue) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumRealValue) GetRealValue() BACnetContextTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumRealValue factory function for _BACnetLogRecordLogDatumRealValue
func NewBACnetLogRecordLogDatumRealValue(realValue BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumRealValue {
	_result := &_BACnetLogRecordLogDatumRealValue{
		RealValue:                realValue,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumRealValue(structType interface{}) BACnetLogRecordLogDatumRealValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumRealValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumRealValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumRealValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumRealValue"
}

func (m *_BACnetLogRecordLogDatumRealValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumRealValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordLogDatumRealValueParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumRealValue, error) {
	return BACnetLogRecordLogDatumRealValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumRealValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumRealValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumRealValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumRealValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
	_realValue, _realValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field of BACnetLogRecordLogDatumRealValue")
	}
	realValue := _realValue.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumRealValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumRealValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumRealValue{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		RealValue: realValue,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumRealValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumRealValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumRealValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumRealValue")
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realValue")
		}
		_realValueErr := writeBuffer.WriteSerializable(ctx, m.GetRealValue())
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realValue")
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumRealValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumRealValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumRealValue) isBACnetLogRecordLogDatumRealValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumRealValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
