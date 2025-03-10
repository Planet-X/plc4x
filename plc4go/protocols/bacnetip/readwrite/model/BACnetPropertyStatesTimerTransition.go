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

// BACnetPropertyStatesTimerTransition is the corresponding interface of BACnetPropertyStatesTimerTransition
type BACnetPropertyStatesTimerTransition interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetTimerTransition returns TimerTransition (property field)
	GetTimerTransition() BACnetTimerTransitionTagged
}

// BACnetPropertyStatesTimerTransitionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesTimerTransition.
// This is useful for switch cases.
type BACnetPropertyStatesTimerTransitionExactly interface {
	BACnetPropertyStatesTimerTransition
	isBACnetPropertyStatesTimerTransition() bool
}

// _BACnetPropertyStatesTimerTransition is the data-structure of this message
type _BACnetPropertyStatesTimerTransition struct {
	*_BACnetPropertyStates
	TimerTransition BACnetTimerTransitionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesTimerTransition) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesTimerTransition) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesTimerTransition) GetTimerTransition() BACnetTimerTransitionTagged {
	return m.TimerTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesTimerTransition factory function for _BACnetPropertyStatesTimerTransition
func NewBACnetPropertyStatesTimerTransition(timerTransition BACnetTimerTransitionTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesTimerTransition {
	_result := &_BACnetPropertyStatesTimerTransition{
		TimerTransition:       timerTransition,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesTimerTransition(structType interface{}) BACnetPropertyStatesTimerTransition {
	if casted, ok := structType.(BACnetPropertyStatesTimerTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesTimerTransition) GetTypeName() string {
	return "BACnetPropertyStatesTimerTransition"
}

func (m *_BACnetPropertyStatesTimerTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timerTransition)
	lengthInBits += m.TimerTransition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesTimerTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesTimerTransitionParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesTimerTransition, error) {
	return BACnetPropertyStatesTimerTransitionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesTimerTransitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesTimerTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesTimerTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerTransition)
	if pullErr := readBuffer.PullContext("timerTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timerTransition")
	}
	_timerTransition, _timerTransitionErr := BACnetTimerTransitionTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _timerTransitionErr != nil {
		return nil, errors.Wrap(_timerTransitionErr, "Error parsing 'timerTransition' field of BACnetPropertyStatesTimerTransition")
	}
	timerTransition := _timerTransition.(BACnetTimerTransitionTagged)
	if closeErr := readBuffer.CloseContext("timerTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timerTransition")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesTimerTransition")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesTimerTransition{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		TimerTransition:       timerTransition,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesTimerTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesTimerTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesTimerTransition")
		}

		// Simple Field (timerTransition)
		if pushErr := writeBuffer.PushContext("timerTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timerTransition")
		}
		_timerTransitionErr := writeBuffer.WriteSerializable(ctx, m.GetTimerTransition())
		if popErr := writeBuffer.PopContext("timerTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timerTransition")
		}
		if _timerTransitionErr != nil {
			return errors.Wrap(_timerTransitionErr, "Error serializing 'timerTransition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesTimerTransition")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesTimerTransition) isBACnetPropertyStatesTimerTransition() bool {
	return true
}

func (m *_BACnetPropertyStatesTimerTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
