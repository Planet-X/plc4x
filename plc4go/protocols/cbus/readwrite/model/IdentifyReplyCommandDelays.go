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

// IdentifyReplyCommandDelays is the corresponding interface of IdentifyReplyCommandDelays
type IdentifyReplyCommandDelays interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetTerminalLevels returns TerminalLevels (property field)
	GetTerminalLevels() []byte
	// GetReStrikeDelay returns ReStrikeDelay (property field)
	GetReStrikeDelay() byte
}

// IdentifyReplyCommandDelaysExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandDelays.
// This is useful for switch cases.
type IdentifyReplyCommandDelaysExactly interface {
	IdentifyReplyCommandDelays
	isIdentifyReplyCommandDelays() bool
}

// _IdentifyReplyCommandDelays is the data-structure of this message
type _IdentifyReplyCommandDelays struct {
	*_IdentifyReplyCommand
	TerminalLevels []byte
	ReStrikeDelay  byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandDelays) GetAttribute() Attribute {
	return Attribute_Delays
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandDelays) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandDelays) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandDelays) GetTerminalLevels() []byte {
	return m.TerminalLevels
}

func (m *_IdentifyReplyCommandDelays) GetReStrikeDelay() byte {
	return m.ReStrikeDelay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandDelays factory function for _IdentifyReplyCommandDelays
func NewIdentifyReplyCommandDelays(terminalLevels []byte, reStrikeDelay byte, numBytes uint8) *_IdentifyReplyCommandDelays {
	_result := &_IdentifyReplyCommandDelays{
		TerminalLevels:        terminalLevels,
		ReStrikeDelay:         reStrikeDelay,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandDelays(structType interface{}) IdentifyReplyCommandDelays {
	if casted, ok := structType.(IdentifyReplyCommandDelays); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDelays); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandDelays) GetTypeName() string {
	return "IdentifyReplyCommandDelays"
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.TerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.TerminalLevels))
	}

	// Simple field (reStrikeDelay)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandDelaysParse(theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandDelays, error) {
	return IdentifyReplyCommandDelaysParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandDelaysParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandDelays, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDelays"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandDelays")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (terminalLevels)
	numberOfBytesterminalLevels := int(uint16(numBytes) - uint16(uint16(1)))
	terminalLevels, _readArrayErr := readBuffer.ReadByteArray("terminalLevels", numberOfBytesterminalLevels)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'terminalLevels' field of IdentifyReplyCommandDelays")
	}

	// Simple Field (reStrikeDelay)
	_reStrikeDelay, _reStrikeDelayErr := readBuffer.ReadByte("reStrikeDelay")
	if _reStrikeDelayErr != nil {
		return nil, errors.Wrap(_reStrikeDelayErr, "Error parsing 'reStrikeDelay' field of IdentifyReplyCommandDelays")
	}
	reStrikeDelay := _reStrikeDelay

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDelays"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandDelays")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandDelays{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		TerminalLevels: terminalLevels,
		ReStrikeDelay:  reStrikeDelay,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandDelays) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandDelays) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDelays"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandDelays")
		}

		// Array Field (terminalLevels)
		// Byte Array field (terminalLevels)
		if err := writeBuffer.WriteByteArray("terminalLevels", m.GetTerminalLevels()); err != nil {
			return errors.Wrap(err, "Error serializing 'terminalLevels' field")
		}

		// Simple Field (reStrikeDelay)
		reStrikeDelay := byte(m.GetReStrikeDelay())
		_reStrikeDelayErr := writeBuffer.WriteByte("reStrikeDelay", (reStrikeDelay))
		if _reStrikeDelayErr != nil {
			return errors.Wrap(_reStrikeDelayErr, "Error serializing 'reStrikeDelay' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDelays"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandDelays")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandDelays) isIdentifyReplyCommandDelays() bool {
	return true
}

func (m *_IdentifyReplyCommandDelays) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
