/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetContextTagOctetString struct {
	*BACnetContextTag
	Payload *BACnetTagPayloadOctetString

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
}

// The corresponding interface
type IBACnetContextTagOctetString interface {
	IBACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadOctetString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagOctetString) GetDataType() BACnetDataType {
	return BACnetDataType_OCTET_STRING
}

func (m *BACnetContextTagOctetString) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagOctetString) GetPayload() *BACnetTagPayloadOctetString {
	return m.Payload
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetContextTagOctetString factory function for BACnetContextTagOctetString
func NewBACnetContextTagOctetString(payload *BACnetTagPayloadOctetString, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagOctetString{
		Payload:          payload,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagOctetString(structType interface{}) *BACnetContextTagOctetString {
	if casted, ok := structType.(BACnetContextTagOctetString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagOctetString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagOctetString(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagOctetString(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagOctetString) GetTypeName() string {
	return "BACnetContextTagOctetString"
}

func (m *BACnetContextTagOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetContextTagOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagOctetStringParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, header *BACnetTagHeader) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagOctetString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadOctetStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadOctetString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagOctetString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagOctetString{
		Payload:          CastBACnetTagPayloadOctetString(payload),
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagOctetString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagOctetString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
