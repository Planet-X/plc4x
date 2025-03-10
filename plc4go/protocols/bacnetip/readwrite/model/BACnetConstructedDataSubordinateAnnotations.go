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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSubordinateAnnotations is the corresponding interface of BACnetConstructedDataSubordinateAnnotations
type BACnetConstructedDataSubordinateAnnotations interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateAnnotations returns SubordinateAnnotations (property field)
	GetSubordinateAnnotations() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataSubordinateAnnotationsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSubordinateAnnotations.
// This is useful for switch cases.
type BACnetConstructedDataSubordinateAnnotationsExactly interface {
	BACnetConstructedDataSubordinateAnnotations
	isBACnetConstructedDataSubordinateAnnotations() bool
}

// _BACnetConstructedDataSubordinateAnnotations is the data-structure of this message
type _BACnetConstructedDataSubordinateAnnotations struct {
	*_BACnetConstructedData
	NumberOfDataElements   BACnetApplicationTagUnsignedInteger
	SubordinateAnnotations []BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateAnnotations) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_ANNOTATIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateAnnotations) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateAnnotations) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetSubordinateAnnotations() []BACnetApplicationTagCharacterString {
	return m.SubordinateAnnotations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateAnnotations) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateAnnotations factory function for _BACnetConstructedDataSubordinateAnnotations
func NewBACnetConstructedDataSubordinateAnnotations(numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateAnnotations []BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateAnnotations {
	_result := &_BACnetConstructedDataSubordinateAnnotations{
		NumberOfDataElements:   numberOfDataElements,
		SubordinateAnnotations: subordinateAnnotations,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubordinateAnnotations(structType interface{}) BACnetConstructedDataSubordinateAnnotations {
	if casted, ok := structType.(BACnetConstructedDataSubordinateAnnotations); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateAnnotations); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetTypeName() string {
	return "BACnetConstructedDataSubordinateAnnotations"
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.SubordinateAnnotations) > 0 {
		for _, element := range m.SubordinateAnnotations {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateAnnotations) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSubordinateAnnotationsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSubordinateAnnotations, error) {
	return BACnetConstructedDataSubordinateAnnotationsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSubordinateAnnotationsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSubordinateAnnotations, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateAnnotations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateAnnotations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataSubordinateAnnotations")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (subordinateAnnotations)
	if pullErr := readBuffer.PullContext("subordinateAnnotations", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subordinateAnnotations")
	}
	// Terminated array
	var subordinateAnnotations []BACnetApplicationTagCharacterString
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateAnnotations' field of BACnetConstructedDataSubordinateAnnotations")
			}
			subordinateAnnotations = append(subordinateAnnotations, _item.(BACnetApplicationTagCharacterString))
		}
	}
	if closeErr := readBuffer.CloseContext("subordinateAnnotations", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subordinateAnnotations")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateAnnotations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateAnnotations")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSubordinateAnnotations{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements:   numberOfDataElements,
		SubordinateAnnotations: subordinateAnnotations,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSubordinateAnnotations) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubordinateAnnotations) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateAnnotations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateAnnotations")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (subordinateAnnotations)
		if pushErr := writeBuffer.PushContext("subordinateAnnotations", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subordinateAnnotations")
		}
		for _curItem, _element := range m.GetSubordinateAnnotations() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetSubordinateAnnotations()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subordinateAnnotations' field")
			}
		}
		if popErr := writeBuffer.PopContext("subordinateAnnotations", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subordinateAnnotations")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateAnnotations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateAnnotations")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateAnnotations) isBACnetConstructedDataSubordinateAnnotations() bool {
	return true
}

func (m *_BACnetConstructedDataSubordinateAnnotations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
