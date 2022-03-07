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
type NotTransmittedSyncLoss struct {
	*Confirmation
}

// The corresponding interface
type INotTransmittedSyncLoss interface {
	IConfirmation
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
func (m *NotTransmittedSyncLoss) GetConfirmationType() byte {
	return 0x25
}

func (m *NotTransmittedSyncLoss) InitializeParent(parent *Confirmation, alpha *Alpha) {
	m.Confirmation.Alpha = alpha
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewNotTransmittedSyncLoss factory function for NotTransmittedSyncLoss
func NewNotTransmittedSyncLoss(alpha *Alpha) *Confirmation {
	child := &NotTransmittedSyncLoss{
		Confirmation: NewConfirmation(alpha),
	}
	child.Child = child
	return child.Confirmation
}

func CastNotTransmittedSyncLoss(structType interface{}) *NotTransmittedSyncLoss {
	if casted, ok := structType.(NotTransmittedSyncLoss); ok {
		return &casted
	}
	if casted, ok := structType.(*NotTransmittedSyncLoss); ok {
		return casted
	}
	if casted, ok := structType.(Confirmation); ok {
		return CastNotTransmittedSyncLoss(casted.Child)
	}
	if casted, ok := structType.(*Confirmation); ok {
		return CastNotTransmittedSyncLoss(casted.Child)
	}
	return nil
}

func (m *NotTransmittedSyncLoss) GetTypeName() string {
	return "NotTransmittedSyncLoss"
}

func (m *NotTransmittedSyncLoss) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NotTransmittedSyncLoss) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *NotTransmittedSyncLoss) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NotTransmittedSyncLossParse(readBuffer utils.ReadBuffer) (*Confirmation, error) {
	if pullErr := readBuffer.PullContext("NotTransmittedSyncLoss"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotTransmittedSyncLoss"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NotTransmittedSyncLoss{
		Confirmation: &Confirmation{},
	}
	_child.Confirmation.Child = _child
	return _child.Confirmation, nil
}

func (m *NotTransmittedSyncLoss) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotTransmittedSyncLoss"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("NotTransmittedSyncLoss"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NotTransmittedSyncLoss) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
