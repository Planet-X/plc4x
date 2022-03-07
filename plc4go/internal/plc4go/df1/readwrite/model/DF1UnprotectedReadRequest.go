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
type DF1UnprotectedReadRequest struct {
	*DF1Command
	Address uint16
	Size    uint8
}

// The corresponding interface
type IDF1UnprotectedReadRequest interface {
	IDF1Command
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetSize returns Size (property field)
	GetSize() uint8
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
func (m *DF1UnprotectedReadRequest) GetCommandCode() uint8 {
	return 0x01
}

func (m *DF1UnprotectedReadRequest) InitializeParent(parent *DF1Command, status uint8, transactionCounter uint16) {
	m.DF1Command.Status = status
	m.DF1Command.TransactionCounter = transactionCounter
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *DF1UnprotectedReadRequest) GetAddress() uint16 {
	return m.Address
}

func (m *DF1UnprotectedReadRequest) GetSize() uint8 {
	return m.Size
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDF1UnprotectedReadRequest factory function for DF1UnprotectedReadRequest
func NewDF1UnprotectedReadRequest(address uint16, size uint8, status uint8, transactionCounter uint16) *DF1Command {
	child := &DF1UnprotectedReadRequest{
		Address:    address,
		Size:       size,
		DF1Command: NewDF1Command(status, transactionCounter),
	}
	child.Child = child
	return child.DF1Command
}

func CastDF1UnprotectedReadRequest(structType interface{}) *DF1UnprotectedReadRequest {
	if casted, ok := structType.(DF1UnprotectedReadRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(DF1Command); ok {
		return CastDF1UnprotectedReadRequest(casted.Child)
	}
	if casted, ok := structType.(*DF1Command); ok {
		return CastDF1UnprotectedReadRequest(casted.Child)
	}
	return nil
}

func (m *DF1UnprotectedReadRequest) GetTypeName() string {
	return "DF1UnprotectedReadRequest"
}

func (m *DF1UnprotectedReadRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1UnprotectedReadRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *DF1UnprotectedReadRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1UnprotectedReadRequestParse(readBuffer utils.ReadBuffer) (*DF1Command, error) {
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (size)
	_size, _sizeErr := readBuffer.ReadUint8("size", 8)
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field")
	}
	size := _size

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1UnprotectedReadRequest{
		Address:    address,
		Size:       size,
		DF1Command: &DF1Command{},
	}
	_child.DF1Command.Child = _child
	return _child.DF1Command, nil
}

func (m *DF1UnprotectedReadRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (size)
		size := uint8(m.Size)
		_sizeErr := writeBuffer.WriteUint8("size", 8, (size))
		if _sizeErr != nil {
			return errors.Wrap(_sizeErr, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1UnprotectedReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
