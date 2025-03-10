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

// BACnetLogDataLogDataEntryIntegerValue is the corresponding interface of BACnetLogDataLogDataEntryIntegerValue
type BACnetLogDataLogDataEntryIntegerValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetContextTagSignedInteger
}

// BACnetLogDataLogDataEntryIntegerValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryIntegerValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryIntegerValueExactly interface {
	BACnetLogDataLogDataEntryIntegerValue
	isBACnetLogDataLogDataEntryIntegerValue() bool
}

// _BACnetLogDataLogDataEntryIntegerValue is the data-structure of this message
type _BACnetLogDataLogDataEntryIntegerValue struct {
	*_BACnetLogDataLogDataEntry
	IntegerValue BACnetContextTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryIntegerValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetIntegerValue() BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryIntegerValue factory function for _BACnetLogDataLogDataEntryIntegerValue
func NewBACnetLogDataLogDataEntryIntegerValue(integerValue BACnetContextTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryIntegerValue {
	_result := &_BACnetLogDataLogDataEntryIntegerValue{
		IntegerValue:               integerValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryIntegerValue(structType interface{}) BACnetLogDataLogDataEntryIntegerValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryIntegerValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryIntegerValue"
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogDataEntryIntegerValueParse(theBytes []byte) (BACnetLogDataLogDataEntryIntegerValue, error) {
	return BACnetLogDataLogDataEntryIntegerValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogDataLogDataEntryIntegerValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryIntegerValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
	_integerValue, _integerValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetLogDataLogDataEntryIntegerValue")
	}
	integerValue := _integerValue.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryIntegerValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryIntegerValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		IntegerValue:               integerValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryIntegerValue")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := writeBuffer.WriteSerializable(ctx, m.GetIntegerValue())
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryIntegerValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) isBACnetLogDataLogDataEntryIntegerValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
