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

// BACnetProcessIdSelection is the corresponding interface of BACnetProcessIdSelection
type BACnetProcessIdSelection interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetProcessIdSelectionExactly can be used when we want exactly this type and not a type which fulfills BACnetProcessIdSelection.
// This is useful for switch cases.
type BACnetProcessIdSelectionExactly interface {
	BACnetProcessIdSelection
	isBACnetProcessIdSelection() bool
}

// _BACnetProcessIdSelection is the data-structure of this message
type _BACnetProcessIdSelection struct {
	_BACnetProcessIdSelectionChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetProcessIdSelectionChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetProcessIdSelectionParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetProcessIdSelection, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetProcessIdSelectionChild interface {
	utils.Serializable
	InitializeParent(parent BACnetProcessIdSelection, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetProcessIdSelection

	GetTypeName() string
	BACnetProcessIdSelection
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelection) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetProcessIdSelection) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProcessIdSelection factory function for _BACnetProcessIdSelection
func NewBACnetProcessIdSelection(peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelection {
	return &_BACnetProcessIdSelection{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelection(structType interface{}) BACnetProcessIdSelection {
	if casted, ok := structType.(BACnetProcessIdSelection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelection) GetTypeName() string {
	return "BACnetProcessIdSelection"
}

func (m *_BACnetProcessIdSelection) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetProcessIdSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProcessIdSelectionParse(theBytes []byte) (BACnetProcessIdSelection, error) {
	return BACnetProcessIdSelectionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetProcessIdSelectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProcessIdSelection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetProcessIdSelectionChildSerializeRequirement interface {
		BACnetProcessIdSelection
		InitializeParent(BACnetProcessIdSelection, BACnetTagHeader)
		GetParent() BACnetProcessIdSelection
	}
	var _childTemp interface{}
	var _child BACnetProcessIdSelectionChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetProcessIdSelectionNull
		_childTemp, typeSwitchError = BACnetProcessIdSelectionNullParseWithBuffer(ctx, readBuffer)
	case 0 == 0: // BACnetProcessIdSelectionValue
		_childTemp, typeSwitchError = BACnetProcessIdSelectionValueParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetProcessIdSelection")
	}
	_child = _childTemp.(BACnetProcessIdSelectionChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelection")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetProcessIdSelection) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetProcessIdSelection, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetProcessIdSelection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelection")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetProcessIdSelection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelection")
	}
	return nil
}

func (m *_BACnetProcessIdSelection) isBACnetProcessIdSelection() bool {
	return true
}

func (m *_BACnetProcessIdSelection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
