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
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse struct {
	*S7PayloadUserDataItem
	Result     uint8
	Reserved01 uint8
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse interface {
	IS7PayloadUserDataItem
	// GetResult returns Result (property field)
	GetResult() uint8
	// GetReserved01 returns Reserved01 (property field)
	GetReserved01() uint8
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
func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetDataLength() uint16 {
	return 0x02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetResult() uint8 {
	return m.Result
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetReserved01() uint8 {
	return m.Reserved01
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse factory function for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(result uint8, reserved01 uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse{
		Result:                result,
		Reserved01:            reserved01,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Child = child
	return child.S7PayloadUserDataItem
}

func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(structType interface{}) *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (result)
	_result, _resultErr := readBuffer.ReadUint8("result", 8)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result

	// Simple Field (reserved01)
	_reserved01, _reserved01Err := readBuffer.ReadUint8("reserved01", 8)
	if _reserved01Err != nil {
		return nil, errors.Wrap(_reserved01Err, "Error parsing 'reserved01' field")
	}
	reserved01 := _reserved01

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse{
		Result:                result,
		Reserved01:            reserved01,
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child.S7PayloadUserDataItem, nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		result := uint8(m.Result)
		_resultErr := writeBuffer.WriteUint8("result", 8, (result))
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (reserved01)
		reserved01 := uint8(m.Reserved01)
		_reserved01Err := writeBuffer.WriteUint8("reserved01", 8, (reserved01))
		if _reserved01Err != nil {
			return errors.Wrap(_reserved01Err, "Error serializing 'reserved01' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
