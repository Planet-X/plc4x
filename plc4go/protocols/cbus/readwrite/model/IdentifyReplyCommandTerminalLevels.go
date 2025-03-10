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

// IdentifyReplyCommandTerminalLevels is the corresponding interface of IdentifyReplyCommandTerminalLevels
type IdentifyReplyCommandTerminalLevels interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetTerminalLevels returns TerminalLevels (property field)
	GetTerminalLevels() []byte
}

// IdentifyReplyCommandTerminalLevelsExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandTerminalLevels.
// This is useful for switch cases.
type IdentifyReplyCommandTerminalLevelsExactly interface {
	IdentifyReplyCommandTerminalLevels
	isIdentifyReplyCommandTerminalLevels() bool
}

// _IdentifyReplyCommandTerminalLevels is the data-structure of this message
type _IdentifyReplyCommandTerminalLevels struct {
	*_IdentifyReplyCommand
	TerminalLevels []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) GetAttribute() Attribute {
	return Attribute_TerminalLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandTerminalLevels) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) GetTerminalLevels() []byte {
	return m.TerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandTerminalLevels factory function for _IdentifyReplyCommandTerminalLevels
func NewIdentifyReplyCommandTerminalLevels(terminalLevels []byte, numBytes uint8) *_IdentifyReplyCommandTerminalLevels {
	_result := &_IdentifyReplyCommandTerminalLevels{
		TerminalLevels:        terminalLevels,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandTerminalLevels(structType interface{}) IdentifyReplyCommandTerminalLevels {
	if casted, ok := structType.(IdentifyReplyCommandTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandTerminalLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandTerminalLevels"
}

func (m *_IdentifyReplyCommandTerminalLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.TerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.TerminalLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandTerminalLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandTerminalLevelsParse(theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandTerminalLevels, error) {
	return IdentifyReplyCommandTerminalLevelsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandTerminalLevelsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandTerminalLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandTerminalLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandTerminalLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (terminalLevels)
	numberOfBytesterminalLevels := int(numBytes)
	terminalLevels, _readArrayErr := readBuffer.ReadByteArray("terminalLevels", numberOfBytesterminalLevels)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'terminalLevels' field of IdentifyReplyCommandTerminalLevels")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandTerminalLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandTerminalLevels")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandTerminalLevels{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		TerminalLevels: terminalLevels,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandTerminalLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandTerminalLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandTerminalLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandTerminalLevels")
		}

		// Array Field (terminalLevels)
		// Byte Array field (terminalLevels)
		if err := writeBuffer.WriteByteArray("terminalLevels", m.GetTerminalLevels()); err != nil {
			return errors.Wrap(err, "Error serializing 'terminalLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandTerminalLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandTerminalLevels")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandTerminalLevels) isIdentifyReplyCommandTerminalLevels() bool {
	return true
}

func (m *_IdentifyReplyCommandTerminalLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
