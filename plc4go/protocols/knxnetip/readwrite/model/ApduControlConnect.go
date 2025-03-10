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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduControlConnect is the corresponding interface of ApduControlConnect
type ApduControlConnect interface {
	utils.LengthAware
	utils.Serializable
	ApduControl
}

// ApduControlConnectExactly can be used when we want exactly this type and not a type which fulfills ApduControlConnect.
// This is useful for switch cases.
type ApduControlConnectExactly interface {
	ApduControlConnect
	isApduControlConnect() bool
}

// _ApduControlConnect is the data-structure of this message
type _ApduControlConnect struct {
	*_ApduControl
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlConnect) GetControlType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlConnect) InitializeParent(parent ApduControl) {}

func (m *_ApduControlConnect) GetParent() ApduControl {
	return m._ApduControl
}

// NewApduControlConnect factory function for _ApduControlConnect
func NewApduControlConnect() *_ApduControlConnect {
	_result := &_ApduControlConnect{
		_ApduControl: NewApduControl(),
	}
	_result._ApduControl._ApduControlChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlConnect(structType interface{}) ApduControlConnect {
	if casted, ok := structType.(ApduControlConnect); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlConnect); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlConnect) GetTypeName() string {
	return "ApduControlConnect"
}

func (m *_ApduControlConnect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduControlConnect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduControlConnectParse(theBytes []byte) (ApduControlConnect, error) {
	return ApduControlConnectParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ApduControlConnectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApduControlConnect, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlConnect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlConnect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlConnect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlConnect")
	}

	// Create a partially initialized instance
	_child := &_ApduControlConnect{
		_ApduControl: &_ApduControl{},
	}
	_child._ApduControl._ApduControlChildRequirements = _child
	return _child, nil
}

func (m *_ApduControlConnect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlConnect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlConnect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlConnect")
		}

		if popErr := writeBuffer.PopContext("ApduControlConnect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlConnect")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlConnect) isApduControlConnect() bool {
	return true
}

func (m *_ApduControlConnect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
