/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadReadVarResponse is the corresponding interface of S7PayloadReadVarResponse
type S7PayloadReadVarResponse interface {
	utils.LengthAware
	utils.Serializable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7VarPayloadDataItem
}

// S7PayloadReadVarResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadReadVarResponse.
// This is useful for switch cases.
type S7PayloadReadVarResponseExactly interface {
	S7PayloadReadVarResponse
	isS7PayloadReadVarResponse() bool
}

// _S7PayloadReadVarResponse is the data-structure of this message
type _S7PayloadReadVarResponse struct {
	*_S7Payload
	Items []S7VarPayloadDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadReadVarResponse) GetParameterParameterType() uint8 {
	return 0x04
}

func (m *_S7PayloadReadVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadReadVarResponse) InitializeParent(parent S7Payload) {}

func (m *_S7PayloadReadVarResponse) GetParent() S7Payload {
	return m._S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadReadVarResponse) GetItems() []S7VarPayloadDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadReadVarResponse factory function for _S7PayloadReadVarResponse
func NewS7PayloadReadVarResponse(items []S7VarPayloadDataItem, parameter S7Parameter) *_S7PayloadReadVarResponse {
	_result := &_S7PayloadReadVarResponse{
		Items:      items,
		_S7Payload: NewS7Payload(parameter),
	}
	_result._S7Payload._S7PayloadChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadReadVarResponse(structType interface{}) S7PayloadReadVarResponse {
	if casted, ok := structType.(S7PayloadReadVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadReadVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadReadVarResponse) GetTypeName() string {
	return "S7PayloadReadVarResponse"
}

func (m *_S7PayloadReadVarResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadReadVarResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadReadVarResponseParse(theBytes []byte, messageType uint8, parameter S7Parameter) (S7PayloadReadVarResponse, error) {
	return S7PayloadReadVarResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), messageType, parameter)
}

func S7PayloadReadVarResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (S7PayloadReadVarResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadReadVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadReadVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]S7VarPayloadDataItem, CastS7ParameterReadVarResponse(parameter).GetNumItems())
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		_numItems := uint16(CastS7ParameterReadVarResponse(parameter).GetNumItems())
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := S7VarPayloadDataItemParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of S7PayloadReadVarResponse")
			}
			items[_curItem] = _item.(S7VarPayloadDataItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadReadVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadReadVarResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadReadVarResponse{
		_S7Payload: &_S7Payload{
			Parameter: parameter,
		},
		Items: items,
	}
	_child._S7Payload._S7PayloadChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadReadVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadReadVarResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadReadVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadReadVarResponse")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("S7PayloadReadVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadReadVarResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadReadVarResponse) isS7PayloadReadVarResponse() bool {
	return true
}

func (m *_S7PayloadReadVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
