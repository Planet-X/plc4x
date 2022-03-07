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
type S7PayloadWriteVarRequest struct {
	*S7Payload
	Items []*S7VarPayloadDataItem

	// Arguments.
	Parameter S7Parameter
}

// The corresponding interface
type IS7PayloadWriteVarRequest interface {
	IS7Payload
	// GetItems returns Items (property field)
	GetItems() []*S7VarPayloadDataItem
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
func (m *S7PayloadWriteVarRequest) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *S7PayloadWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

func (m *S7PayloadWriteVarRequest) InitializeParent(parent *S7Payload) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7PayloadWriteVarRequest) GetItems() []*S7VarPayloadDataItem {
	return m.Items
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7PayloadWriteVarRequest factory function for S7PayloadWriteVarRequest
func NewS7PayloadWriteVarRequest(items []*S7VarPayloadDataItem, parameter S7Parameter) *S7Payload {
	child := &S7PayloadWriteVarRequest{
		Items:     items,
		S7Payload: NewS7Payload(parameter),
	}
	child.Child = child
	return child.S7Payload
}

func CastS7PayloadWriteVarRequest(structType interface{}) *S7PayloadWriteVarRequest {
	if casted, ok := structType.(S7PayloadWriteVarRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(S7Payload); ok {
		return CastS7PayloadWriteVarRequest(casted.Child)
	}
	if casted, ok := structType.(*S7Payload); ok {
		return CastS7PayloadWriteVarRequest(casted.Child)
	}
	return nil
}

func (m *S7PayloadWriteVarRequest) GetTypeName() string {
	return "S7PayloadWriteVarRequest"
}

func (m *S7PayloadWriteVarRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadWriteVarRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadWriteVarRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadWriteVarRequestParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7Payload, error) {
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7VarPayloadDataItem, uint16(len(CastS7ParameterWriteVarRequest(parameter).GetItems())))
	{
		for curItem := uint16(0); curItem < uint16(uint16(len(CastS7ParameterWriteVarRequest(parameter).GetItems()))); curItem++ {
			_item, _err := S7VarPayloadDataItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastS7VarPayloadDataItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadWriteVarRequest{
		Items:     items,
		S7Payload: &S7Payload{},
	}
	_child.S7Payload.Child = _child
	return _child.S7Payload, nil
}

func (m *S7PayloadWriteVarRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarRequest"); pushErr != nil {
			return pushErr
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
