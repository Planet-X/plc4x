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
type BACnetConfirmedServiceRequestReadPropertyMultiple struct {
	*BACnetConfirmedServiceRequest
	Data []*BACnetReadAccessSpecification

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestReadPropertyMultiple interface {
	IBACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []*BACnetReadAccessSpecification
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
func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetServiceChoice() uint8 {
	return 0x0E
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetData() []*BACnetReadAccessSpecification {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadPropertyMultiple factory function for BACnetConfirmedServiceRequestReadPropertyMultiple
func NewBACnetConfirmedServiceRequestReadPropertyMultiple(data []*BACnetReadAccessSpecification, len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestReadPropertyMultiple{
		Data:                          data,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestReadPropertyMultiple(structType interface{}) *BACnetConfirmedServiceRequestReadPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReadPropertyMultiple(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReadPropertyMultiple(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadPropertyMultiple"
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadPropertyMultipleParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]*BACnetReadAccessSpecification, 0)
	{
		_dataLength := len
		_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
		for readBuffer.GetPos() < _dataEndPos {
			_item, _err := BACnetReadAccessSpecificationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestReadPropertyMultiple{
		Data:                          data,
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestReadPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
