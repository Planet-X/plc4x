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

// BACnetFaultParameterFaultExtendedParametersEntryOctetString is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryOctetString
type BACnetFaultParameterFaultExtendedParametersEntryOctetString interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
}

// BACnetFaultParameterFaultExtendedParametersEntryOctetStringExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryOctetString.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryOctetStringExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryOctetString
	isBACnetFaultParameterFaultExtendedParametersEntryOctetString() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryOctetString is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryOctetString struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	OctetStringValue BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryOctetString factory function for _BACnetFaultParameterFaultExtendedParametersEntryOctetString
func NewBACnetFaultParameterFaultExtendedParametersEntryOctetString(octetStringValue BACnetApplicationTagOctetString, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryOctetString{
		OctetStringValue: octetStringValue,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryOctetString(structType interface{}) BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryOctetString"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryOctetStringParse(theBytes []byte) (BACnetFaultParameterFaultExtendedParametersEntryOctetString, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryOctetStringParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryOctetStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octetStringValue)
	if pullErr := readBuffer.PullContext("octetStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for octetStringValue")
	}
	_octetStringValue, _octetStringValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _octetStringValueErr != nil {
		return nil, errors.Wrap(_octetStringValueErr, "Error parsing 'octetStringValue' field of BACnetFaultParameterFaultExtendedParametersEntryOctetString")
	}
	octetStringValue := _octetStringValue.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("octetStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for octetStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryOctetString{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		OctetStringValue: octetStringValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
		}

		// Simple Field (octetStringValue)
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for octetStringValue")
		}
		_octetStringValueErr := writeBuffer.WriteSerializable(ctx, m.GetOctetStringValue())
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for octetStringValue")
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) isBACnetFaultParameterFaultExtendedParametersEntryOctetString() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
