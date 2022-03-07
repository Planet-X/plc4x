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
type COTPParameterChecksum struct {
	*COTPParameter
	Crc uint8

	// Arguments.
	Rest uint8
}

// The corresponding interface
type ICOTPParameterChecksum interface {
	ICOTPParameter
	// GetCrc returns Crc (property field)
	GetCrc() uint8
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
func (m *COTPParameterChecksum) GetParameterType() uint8 {
	return 0xC3
}

func (m *COTPParameterChecksum) InitializeParent(parent *COTPParameter) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *COTPParameterChecksum) GetCrc() uint8 {
	return m.Crc
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCOTPParameterChecksum factory function for COTPParameterChecksum
func NewCOTPParameterChecksum(crc uint8, rest uint8) *COTPParameter {
	child := &COTPParameterChecksum{
		Crc:           crc,
		COTPParameter: NewCOTPParameter(rest),
	}
	child.Child = child
	return child.COTPParameter
}

func CastCOTPParameterChecksum(structType interface{}) *COTPParameterChecksum {
	if casted, ok := structType.(COTPParameterChecksum); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPParameterChecksum); ok {
		return casted
	}
	if casted, ok := structType.(COTPParameter); ok {
		return CastCOTPParameterChecksum(casted.Child)
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return CastCOTPParameterChecksum(casted.Child)
	}
	return nil
}

func (m *COTPParameterChecksum) GetTypeName() string {
	return "COTPParameterChecksum"
}

func (m *COTPParameterChecksum) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPParameterChecksum) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (crc)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPParameterChecksum) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterChecksumParse(readBuffer utils.ReadBuffer, rest uint8) (*COTPParameter, error) {
	if pullErr := readBuffer.PullContext("COTPParameterChecksum"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint8("crc", 8)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("COTPParameterChecksum"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPParameterChecksum{
		Crc:           crc,
		COTPParameter: &COTPParameter{},
	}
	_child.COTPParameter.Child = _child
	return _child.COTPParameter, nil
}

func (m *COTPParameterChecksum) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterChecksum"); pushErr != nil {
			return pushErr
		}

		// Simple Field (crc)
		crc := uint8(m.Crc)
		_crcErr := writeBuffer.WriteUint8("crc", 8, (crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterChecksum"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPParameterChecksum) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
