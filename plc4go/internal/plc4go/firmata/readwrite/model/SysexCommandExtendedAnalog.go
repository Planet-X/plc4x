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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type SysexCommandExtendedAnalog struct {
	*SysexCommand
}

// The corresponding interface
type ISysexCommandExtendedAnalog interface {
	ISysexCommand
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
func (m *SysexCommandExtendedAnalog) GetCommandType() uint8 {
	return 0x6F
}

func (m *SysexCommandExtendedAnalog) GetResponse() bool {
	return false
}

func (m *SysexCommandExtendedAnalog) InitializeParent(parent *SysexCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewSysexCommandExtendedAnalog factory function for SysexCommandExtendedAnalog
func NewSysexCommandExtendedAnalog() *SysexCommand {
	child := &SysexCommandExtendedAnalog{
		SysexCommand: NewSysexCommand(),
	}
	child.Child = child
	return child.SysexCommand
}

func CastSysexCommandExtendedAnalog(structType interface{}) *SysexCommandExtendedAnalog {
	if casted, ok := structType.(SysexCommandExtendedAnalog); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandExtendedAnalog); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandExtendedAnalog(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandExtendedAnalog(casted.Child)
	}
	return nil
}

func (m *SysexCommandExtendedAnalog) GetTypeName() string {
	return "SysexCommandExtendedAnalog"
}

func (m *SysexCommandExtendedAnalog) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandExtendedAnalog) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandExtendedAnalog) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandExtendedAnalogParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandExtendedAnalog"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandExtendedAnalog"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandExtendedAnalog{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child.SysexCommand, nil
}

func (m *SysexCommandExtendedAnalog) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExtendedAnalog"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandExtendedAnalog"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandExtendedAnalog) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
