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
type MFuncPropCommandReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// The corresponding interface
type IMFuncPropCommandReq interface {
	ICEMI
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
func (m *MFuncPropCommandReq) GetMessageCode() uint8 {
	return 0xF8
}

func (m *MFuncPropCommandReq) InitializeParent(parent *CEMI) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewMFuncPropCommandReq factory function for MFuncPropCommandReq
func NewMFuncPropCommandReq(size uint16) *CEMI {
	child := &MFuncPropCommandReq{
		CEMI: NewCEMI(size),
	}
	child.Child = child
	return child.CEMI
}

func CastMFuncPropCommandReq(structType interface{}) *MFuncPropCommandReq {
	if casted, ok := structType.(MFuncPropCommandReq); ok {
		return &casted
	}
	if casted, ok := structType.(*MFuncPropCommandReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastMFuncPropCommandReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastMFuncPropCommandReq(casted.Child)
	}
	return nil
}

func (m *MFuncPropCommandReq) GetTypeName() string {
	return "MFuncPropCommandReq"
}

func (m *MFuncPropCommandReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MFuncPropCommandReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *MFuncPropCommandReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MFuncPropCommandReqParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("MFuncPropCommandReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCommandReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MFuncPropCommandReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child.CEMI, nil
}

func (m *MFuncPropCommandReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCommandReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MFuncPropCommandReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MFuncPropCommandReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
