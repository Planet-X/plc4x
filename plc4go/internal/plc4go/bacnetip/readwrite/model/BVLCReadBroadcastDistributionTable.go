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
type BVLCReadBroadcastDistributionTable struct {
	*BVLC
}

// The corresponding interface
type IBVLCReadBroadcastDistributionTable interface {
	IBVLC
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
func (m *BVLCReadBroadcastDistributionTable) GetBvlcFunction() uint8 {
	return 0x02
}

func (m *BVLCReadBroadcastDistributionTable) InitializeParent(parent *BVLC) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBVLCReadBroadcastDistributionTable factory function for BVLCReadBroadcastDistributionTable
func NewBVLCReadBroadcastDistributionTable() *BVLC {
	child := &BVLCReadBroadcastDistributionTable{
		BVLC: NewBVLC(),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCReadBroadcastDistributionTable(structType interface{}) *BVLCReadBroadcastDistributionTable {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTable); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCReadBroadcastDistributionTable(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCReadBroadcastDistributionTable(casted.Child)
	}
	return nil
}

func (m *BVLCReadBroadcastDistributionTable) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTable"
}

func (m *BVLCReadBroadcastDistributionTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCReadBroadcastDistributionTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BVLCReadBroadcastDistributionTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCReadBroadcastDistributionTableParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTable"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTable"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCReadBroadcastDistributionTable{
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCReadBroadcastDistributionTable) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTable"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTable"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCReadBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
