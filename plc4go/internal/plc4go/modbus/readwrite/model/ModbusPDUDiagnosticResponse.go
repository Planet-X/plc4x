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
type ModbusPDUDiagnosticResponse struct {
	*ModbusPDU
	SubFunction uint16
	Data        uint16
}

// The corresponding interface
type IModbusPDUDiagnosticResponse interface {
	IModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
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
func (m *ModbusPDUDiagnosticResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUDiagnosticResponse) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *ModbusPDUDiagnosticResponse) GetResponse() bool {
	return bool(true)
}

func (m *ModbusPDUDiagnosticResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUDiagnosticResponse) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *ModbusPDUDiagnosticResponse) GetData() uint16 {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUDiagnosticResponse factory function for ModbusPDUDiagnosticResponse
func NewModbusPDUDiagnosticResponse(subFunction uint16, data uint16) *ModbusPDU {
	child := &ModbusPDUDiagnosticResponse{
		SubFunction: subFunction,
		Data:        data,
		ModbusPDU:   NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUDiagnosticResponse(structType interface{}) *ModbusPDUDiagnosticResponse {
	if casted, ok := structType.(ModbusPDUDiagnosticResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUDiagnosticResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUDiagnosticResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUDiagnosticResponse) GetTypeName() string {
	return "ModbusPDUDiagnosticResponse"
}

func (m *ModbusPDUDiagnosticResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUDiagnosticResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUDiagnosticResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUDiagnosticResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subFunction)
	_subFunction, _subFunctionErr := readBuffer.ReadUint16("subFunction", 16)
	if _subFunctionErr != nil {
		return nil, errors.Wrap(_subFunctionErr, "Error parsing 'subFunction' field")
	}
	subFunction := _subFunction

	// Simple Field (data)
	_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUDiagnosticResponse{
		SubFunction: subFunction,
		Data:        data,
		ModbusPDU:   &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUDiagnosticResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subFunction)
		subFunction := uint16(m.SubFunction)
		_subFunctionErr := writeBuffer.WriteUint16("subFunction", 16, (subFunction))
		if _subFunctionErr != nil {
			return errors.Wrap(_subFunctionErr, "Error serializing 'subFunction' field")
		}

		// Simple Field (data)
		data := uint16(m.Data)
		_dataErr := writeBuffer.WriteUint16("data", 16, (data))
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUDiagnosticResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
