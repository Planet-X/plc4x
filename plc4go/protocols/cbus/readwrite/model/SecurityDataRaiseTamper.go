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

// SecurityDataRaiseTamper is the corresponding interface of SecurityDataRaiseTamper
type SecurityDataRaiseTamper interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataRaiseTamperExactly can be used when we want exactly this type and not a type which fulfills SecurityDataRaiseTamper.
// This is useful for switch cases.
type SecurityDataRaiseTamperExactly interface {
	SecurityDataRaiseTamper
	isSecurityDataRaiseTamper() bool
}

// _SecurityDataRaiseTamper is the data-structure of this message
type _SecurityDataRaiseTamper struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataRaiseTamper) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataRaiseTamper) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataRaiseTamper factory function for _SecurityDataRaiseTamper
func NewSecurityDataRaiseTamper(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataRaiseTamper {
	_result := &_SecurityDataRaiseTamper{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataRaiseTamper(structType interface{}) SecurityDataRaiseTamper {
	if casted, ok := structType.(SecurityDataRaiseTamper); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataRaiseTamper); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataRaiseTamper) GetTypeName() string {
	return "SecurityDataRaiseTamper"
}

func (m *_SecurityDataRaiseTamper) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataRaiseTamper) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataRaiseTamperParse(theBytes []byte) (SecurityDataRaiseTamper, error) {
	return SecurityDataRaiseTamperParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataRaiseTamperParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataRaiseTamper, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataRaiseTamper"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataRaiseTamper")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataRaiseTamper"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataRaiseTamper")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataRaiseTamper{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataRaiseTamper) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataRaiseTamper) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataRaiseTamper"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataRaiseTamper")
		}

		if popErr := writeBuffer.PopContext("SecurityDataRaiseTamper"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataRaiseTamper")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataRaiseTamper) isSecurityDataRaiseTamper() bool {
	return true
}

func (m *_SecurityDataRaiseTamper) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
