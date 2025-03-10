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

// BACnetTimerStateChangeValueNoValue is the corresponding interface of BACnetTimerStateChangeValueNoValue
type BACnetTimerStateChangeValueNoValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetNoValue returns NoValue (property field)
	GetNoValue() BACnetContextTagNull
}

// BACnetTimerStateChangeValueNoValueExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueNoValue.
// This is useful for switch cases.
type BACnetTimerStateChangeValueNoValueExactly interface {
	BACnetTimerStateChangeValueNoValue
	isBACnetTimerStateChangeValueNoValue() bool
}

// _BACnetTimerStateChangeValueNoValue is the data-structure of this message
type _BACnetTimerStateChangeValueNoValue struct {
	*_BACnetTimerStateChangeValue
	NoValue BACnetContextTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueNoValue) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueNoValue) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueNoValue) GetNoValue() BACnetContextTagNull {
	return m.NoValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueNoValue factory function for _BACnetTimerStateChangeValueNoValue
func NewBACnetTimerStateChangeValueNoValue(noValue BACnetContextTagNull, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueNoValue {
	_result := &_BACnetTimerStateChangeValueNoValue{
		NoValue:                      noValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueNoValue(structType interface{}) BACnetTimerStateChangeValueNoValue {
	if casted, ok := structType.(BACnetTimerStateChangeValueNoValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueNoValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueNoValue) GetTypeName() string {
	return "BACnetTimerStateChangeValueNoValue"
}

func (m *_BACnetTimerStateChangeValueNoValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noValue)
	lengthInBits += m.NoValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueNoValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueNoValueParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueNoValue, error) {
	return BACnetTimerStateChangeValueNoValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueNoValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueNoValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueNoValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueNoValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noValue)
	if pullErr := readBuffer.PullContext("noValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for noValue")
	}
	_noValue, _noValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_NULL))
	if _noValueErr != nil {
		return nil, errors.Wrap(_noValueErr, "Error parsing 'noValue' field of BACnetTimerStateChangeValueNoValue")
	}
	noValue := _noValue.(BACnetContextTagNull)
	if closeErr := readBuffer.CloseContext("noValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for noValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueNoValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueNoValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueNoValue{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		NoValue: noValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueNoValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueNoValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueNoValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueNoValue")
		}

		// Simple Field (noValue)
		if pushErr := writeBuffer.PushContext("noValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for noValue")
		}
		_noValueErr := writeBuffer.WriteSerializable(ctx, m.GetNoValue())
		if popErr := writeBuffer.PopContext("noValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for noValue")
		}
		if _noValueErr != nil {
			return errors.Wrap(_noValueErr, "Error serializing 'noValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueNoValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueNoValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueNoValue) isBACnetTimerStateChangeValueNoValue() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueNoValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
