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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventParameterChangeOfBitstringListOfBitstringValues is the corresponding interface of BACnetEventParameterChangeOfBitstringListOfBitstringValues
type BACnetEventParameterChangeOfBitstringListOfBitstringValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfBitstringValues returns ListOfBitstringValues (property field)
	GetListOfBitstringValues() []BACnetApplicationTagBitString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfBitstringListOfBitstringValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfBitstringListOfBitstringValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfBitstringListOfBitstringValuesExactly interface {
	BACnetEventParameterChangeOfBitstringListOfBitstringValues
	isBACnetEventParameterChangeOfBitstringListOfBitstringValues() bool
}

// _BACnetEventParameterChangeOfBitstringListOfBitstringValues is the data-structure of this message
type _BACnetEventParameterChangeOfBitstringListOfBitstringValues struct {
	OpeningTag            BACnetOpeningTag
	ListOfBitstringValues []BACnetApplicationTagBitString
	ClosingTag            BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetListOfBitstringValues() []BACnetApplicationTagBitString {
	return m.ListOfBitstringValues
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfBitstringListOfBitstringValues factory function for _BACnetEventParameterChangeOfBitstringListOfBitstringValues
func NewBACnetEventParameterChangeOfBitstringListOfBitstringValues(openingTag BACnetOpeningTag, listOfBitstringValues []BACnetApplicationTagBitString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	return &_BACnetEventParameterChangeOfBitstringListOfBitstringValues{OpeningTag: openingTag, ListOfBitstringValues: listOfBitstringValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfBitstringListOfBitstringValues(structType interface{}) BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfBitstringListOfBitstringValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfBitstringListOfBitstringValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfBitstringListOfBitstringValues"
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfBitstringValues) > 0 {
		for _, element := range m.ListOfBitstringValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfBitstringListOfBitstringValuesParse(theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
	return BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfBitstringValues)
	if pullErr := readBuffer.PullContext("listOfBitstringValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfBitstringValues")
	}
	// Terminated array
	var listOfBitstringValues []BACnetApplicationTagBitString
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfBitstringValues' field of BACnetEventParameterChangeOfBitstringListOfBitstringValues")
			}
			listOfBitstringValues = append(listOfBitstringValues, _item.(BACnetApplicationTagBitString))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfBitstringValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfBitstringValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfBitstringListOfBitstringValues{
		TagNumber:             tagNumber,
		OpeningTag:            openingTag,
		ListOfBitstringValues: listOfBitstringValues,
		ClosingTag:            closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfBitstringValues)
	if pushErr := writeBuffer.PushContext("listOfBitstringValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfBitstringValues")
	}
	for _curItem, _element := range m.GetListOfBitstringValues() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetListOfBitstringValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfBitstringValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfBitstringValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfBitstringValues")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) isBACnetEventParameterChangeOfBitstringListOfBitstringValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
