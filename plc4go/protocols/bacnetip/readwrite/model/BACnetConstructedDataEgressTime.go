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

// BACnetConstructedDataEgressTime is the corresponding interface of BACnetConstructedDataEgressTime
type BACnetConstructedDataEgressTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEgressTime returns EgressTime (property field)
	GetEgressTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataEgressTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEgressTime.
// This is useful for switch cases.
type BACnetConstructedDataEgressTimeExactly interface {
	BACnetConstructedDataEgressTime
	isBACnetConstructedDataEgressTime() bool
}

// _BACnetConstructedDataEgressTime is the data-structure of this message
type _BACnetConstructedDataEgressTime struct {
	*_BACnetConstructedData
	EgressTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEgressTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EGRESS_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEgressTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEgressTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetEgressTime() BACnetApplicationTagUnsignedInteger {
	return m.EgressTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEgressTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetEgressTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEgressTime factory function for _BACnetConstructedDataEgressTime
func NewBACnetConstructedDataEgressTime(egressTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEgressTime {
	_result := &_BACnetConstructedDataEgressTime{
		EgressTime:             egressTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEgressTime(structType interface{}) BACnetConstructedDataEgressTime {
	if casted, ok := structType.(BACnetConstructedDataEgressTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEgressTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEgressTime) GetTypeName() string {
	return "BACnetConstructedDataEgressTime"
}

func (m *_BACnetConstructedDataEgressTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (egressTime)
	lengthInBits += m.EgressTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEgressTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataEgressTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEgressTime, error) {
	return BACnetConstructedDataEgressTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataEgressTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEgressTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEgressTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEgressTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (egressTime)
	if pullErr := readBuffer.PullContext("egressTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for egressTime")
	}
	_egressTime, _egressTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _egressTimeErr != nil {
		return nil, errors.Wrap(_egressTimeErr, "Error parsing 'egressTime' field of BACnetConstructedDataEgressTime")
	}
	egressTime := _egressTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("egressTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for egressTime")
	}

	// Virtual field
	_actualValue := egressTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEgressTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEgressTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEgressTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EgressTime: egressTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEgressTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEgressTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEgressTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEgressTime")
		}

		// Simple Field (egressTime)
		if pushErr := writeBuffer.PushContext("egressTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for egressTime")
		}
		_egressTimeErr := writeBuffer.WriteSerializable(ctx, m.GetEgressTime())
		if popErr := writeBuffer.PopContext("egressTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for egressTime")
		}
		if _egressTimeErr != nil {
			return errors.Wrap(_egressTimeErr, "Error serializing 'egressTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEgressTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEgressTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEgressTime) isBACnetConstructedDataEgressTime() bool {
	return true
}

func (m *_BACnetConstructedDataEgressTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
