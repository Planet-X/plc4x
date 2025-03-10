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

// BACnetEventParameterChangeOfBitstring is the corresponding interface of BACnetEventParameterChangeOfBitstring
type BACnetEventParameterChangeOfBitstring interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetBitmask returns Bitmask (property field)
	GetBitmask() BACnetContextTagBitString
	// GetListOfBitstringValues returns ListOfBitstringValues (property field)
	GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfBitstringExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfBitstring.
// This is useful for switch cases.
type BACnetEventParameterChangeOfBitstringExactly interface {
	BACnetEventParameterChangeOfBitstring
	isBACnetEventParameterChangeOfBitstring() bool
}

// _BACnetEventParameterChangeOfBitstring is the data-structure of this message
type _BACnetEventParameterChangeOfBitstring struct {
	*_BACnetEventParameter
	OpeningTag            BACnetOpeningTag
	TimeDelay             BACnetContextTagUnsignedInteger
	Bitmask               BACnetContextTagBitString
	ListOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues
	ClosingTag            BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfBitstring) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfBitstring) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfBitstring) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfBitstring) GetBitmask() BACnetContextTagBitString {
	return m.Bitmask
}

func (m *_BACnetEventParameterChangeOfBitstring) GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	return m.ListOfBitstringValues
}

func (m *_BACnetEventParameterChangeOfBitstring) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfBitstring factory function for _BACnetEventParameterChangeOfBitstring
func NewBACnetEventParameterChangeOfBitstring(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, bitmask BACnetContextTagBitString, listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfBitstring {
	_result := &_BACnetEventParameterChangeOfBitstring{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		Bitmask:               bitmask,
		ListOfBitstringValues: listOfBitstringValues,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfBitstring(structType interface{}) BACnetEventParameterChangeOfBitstring {
	if casted, ok := structType.(BACnetEventParameterChangeOfBitstring); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfBitstring); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTypeName() string {
	return "BACnetEventParameterChangeOfBitstring"
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (bitmask)
	lengthInBits += m.Bitmask.GetLengthInBits(ctx)

	// Simple field (listOfBitstringValues)
	lengthInBits += m.ListOfBitstringValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfBitstringParse(theBytes []byte) (BACnetEventParameterChangeOfBitstring, error) {
	return BACnetEventParameterChangeOfBitstringParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterChangeOfBitstringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstring, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfBitstring"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfBitstring")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfBitstring")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterChangeOfBitstring")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (bitmask)
	if pullErr := readBuffer.PullContext("bitmask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitmask")
	}
	_bitmask, _bitmaskErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _bitmaskErr != nil {
		return nil, errors.Wrap(_bitmaskErr, "Error parsing 'bitmask' field of BACnetEventParameterChangeOfBitstring")
	}
	bitmask := _bitmask.(BACnetContextTagBitString)
	if closeErr := readBuffer.CloseContext("bitmask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitmask")
	}

	// Simple Field (listOfBitstringValues)
	if pullErr := readBuffer.PullContext("listOfBitstringValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfBitstringValues")
	}
	_listOfBitstringValues, _listOfBitstringValuesErr := BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _listOfBitstringValuesErr != nil {
		return nil, errors.Wrap(_listOfBitstringValuesErr, "Error parsing 'listOfBitstringValues' field of BACnetEventParameterChangeOfBitstring")
	}
	listOfBitstringValues := _listOfBitstringValues.(BACnetEventParameterChangeOfBitstringListOfBitstringValues)
	if closeErr := readBuffer.CloseContext("listOfBitstringValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfBitstringValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfBitstring")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfBitstring"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfBitstring")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfBitstring{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		Bitmask:               bitmask,
		ListOfBitstringValues: listOfBitstringValues,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfBitstring) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfBitstring) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfBitstring"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfBitstring")
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(ctx, m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (bitmask)
		if pushErr := writeBuffer.PushContext("bitmask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitmask")
		}
		_bitmaskErr := writeBuffer.WriteSerializable(ctx, m.GetBitmask())
		if popErr := writeBuffer.PopContext("bitmask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitmask")
		}
		if _bitmaskErr != nil {
			return errors.Wrap(_bitmaskErr, "Error serializing 'bitmask' field")
		}

		// Simple Field (listOfBitstringValues)
		if pushErr := writeBuffer.PushContext("listOfBitstringValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfBitstringValues")
		}
		_listOfBitstringValuesErr := writeBuffer.WriteSerializable(ctx, m.GetListOfBitstringValues())
		if popErr := writeBuffer.PopContext("listOfBitstringValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfBitstringValues")
		}
		if _listOfBitstringValuesErr != nil {
			return errors.Wrap(_listOfBitstringValuesErr, "Error serializing 'listOfBitstringValues' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfBitstring"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfBitstring")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfBitstring) isBACnetEventParameterChangeOfBitstring() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfBitstring) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
