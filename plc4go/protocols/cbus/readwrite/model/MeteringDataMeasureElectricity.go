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

// MeteringDataMeasureElectricity is the corresponding interface of MeteringDataMeasureElectricity
type MeteringDataMeasureElectricity interface {
	utils.LengthAware
	utils.Serializable
	MeteringData
}

// MeteringDataMeasureElectricityExactly can be used when we want exactly this type and not a type which fulfills MeteringDataMeasureElectricity.
// This is useful for switch cases.
type MeteringDataMeasureElectricityExactly interface {
	MeteringDataMeasureElectricity
	isMeteringDataMeasureElectricity() bool
}

// _MeteringDataMeasureElectricity is the data-structure of this message
type _MeteringDataMeasureElectricity struct {
	*_MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataMeasureElectricity) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataMeasureElectricity) GetParent() MeteringData {
	return m._MeteringData
}

// NewMeteringDataMeasureElectricity factory function for _MeteringDataMeasureElectricity
func NewMeteringDataMeasureElectricity(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureElectricity {
	_result := &_MeteringDataMeasureElectricity{
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureElectricity(structType interface{}) MeteringDataMeasureElectricity {
	if casted, ok := structType.(MeteringDataMeasureElectricity); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureElectricity); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureElectricity) GetTypeName() string {
	return "MeteringDataMeasureElectricity"
}

func (m *_MeteringDataMeasureElectricity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureElectricity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataMeasureElectricityParse(theBytes []byte) (MeteringDataMeasureElectricity, error) {
	return MeteringDataMeasureElectricityParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataMeasureElectricityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureElectricity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureElectricity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureElectricity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureElectricity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureElectricity")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataMeasureElectricity{
		_MeteringData: &_MeteringData{},
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataMeasureElectricity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureElectricity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureElectricity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureElectricity")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureElectricity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureElectricity")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureElectricity) isMeteringDataMeasureElectricity() bool {
	return true
}

func (m *_MeteringDataMeasureElectricity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
