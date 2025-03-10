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

// BACnetConstructedDataLargeAnalogValueFaultHighLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueFaultHighLimit
type BACnetConstructedDataLargeAnalogValueFaultHighLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
}

// BACnetConstructedDataLargeAnalogValueFaultHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLargeAnalogValueFaultHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataLargeAnalogValueFaultHighLimitExactly interface {
	BACnetConstructedDataLargeAnalogValueFaultHighLimit
	isBACnetConstructedDataLargeAnalogValueFaultHighLimit() bool
}

// _BACnetConstructedDataLargeAnalogValueFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueFaultHighLimit struct {
	*_BACnetConstructedData
	FaultHighLimit BACnetApplicationTagDouble
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagDouble {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetFaultHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLargeAnalogValueFaultHighLimit factory function for _BACnetConstructedDataLargeAnalogValueFaultHighLimit
func NewBACnetConstructedDataLargeAnalogValueFaultHighLimit(faultHighLimit BACnetApplicationTagDouble, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	_result := &_BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		FaultHighLimit:         faultHighLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueFaultHighLimit(structType interface{}) BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueFaultHighLimit"
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLargeAnalogValueFaultHighLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueFaultHighLimit, error) {
	return BACnetConstructedDataLargeAnalogValueFaultHighLimitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLargeAnalogValueFaultHighLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultHighLimit)
	if pullErr := readBuffer.PullContext("faultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultHighLimit")
	}
	_faultHighLimit, _faultHighLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _faultHighLimitErr != nil {
		return nil, errors.Wrap(_faultHighLimitErr, "Error parsing 'faultHighLimit' field of BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}
	faultHighLimit := _faultHighLimit.(BACnetApplicationTagDouble)
	if closeErr := readBuffer.CloseContext("faultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultHighLimit")
	}

	// Virtual field
	_actualValue := faultHighLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultHighLimit: faultHighLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}

		// Simple Field (faultHighLimit)
		if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
		}
		_faultHighLimitErr := writeBuffer.WriteSerializable(ctx, m.GetFaultHighLimit())
		if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultHighLimit")
		}
		if _faultHighLimitErr != nil {
			return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) isBACnetConstructedDataLargeAnalogValueFaultHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
