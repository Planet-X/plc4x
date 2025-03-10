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

// BACnetHostNPortEnclosed is the corresponding interface of BACnetHostNPortEnclosed
type BACnetHostNPortEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetBacnetHostNPort returns BacnetHostNPort (property field)
	GetBacnetHostNPort() BACnetHostNPort
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetHostNPortEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetHostNPortEnclosed.
// This is useful for switch cases.
type BACnetHostNPortEnclosedExactly interface {
	BACnetHostNPortEnclosed
	isBACnetHostNPortEnclosed() bool
}

// _BACnetHostNPortEnclosed is the data-structure of this message
type _BACnetHostNPortEnclosed struct {
	OpeningTag      BACnetOpeningTag
	BacnetHostNPort BACnetHostNPort
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostNPortEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetHostNPortEnclosed) GetBacnetHostNPort() BACnetHostNPort {
	return m.BacnetHostNPort
}

func (m *_BACnetHostNPortEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostNPortEnclosed factory function for _BACnetHostNPortEnclosed
func NewBACnetHostNPortEnclosed(openingTag BACnetOpeningTag, bacnetHostNPort BACnetHostNPort, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetHostNPortEnclosed {
	return &_BACnetHostNPortEnclosed{OpeningTag: openingTag, BacnetHostNPort: bacnetHostNPort, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostNPortEnclosed(structType interface{}) BACnetHostNPortEnclosed {
	if casted, ok := structType.(BACnetHostNPortEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostNPortEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostNPortEnclosed) GetTypeName() string {
	return "BACnetHostNPortEnclosed"
}

func (m *_BACnetHostNPortEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (bacnetHostNPort)
	lengthInBits += m.BacnetHostNPort.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostNPortEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostNPortEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetHostNPortEnclosed, error) {
	return BACnetHostNPortEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetHostNPortEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetHostNPortEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPortEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPortEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetHostNPortEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (bacnetHostNPort)
	if pullErr := readBuffer.PullContext("bacnetHostNPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bacnetHostNPort")
	}
	_bacnetHostNPort, _bacnetHostNPortErr := BACnetHostNPortParseWithBuffer(ctx, readBuffer)
	if _bacnetHostNPortErr != nil {
		return nil, errors.Wrap(_bacnetHostNPortErr, "Error parsing 'bacnetHostNPort' field of BACnetHostNPortEnclosed")
	}
	bacnetHostNPort := _bacnetHostNPort.(BACnetHostNPort)
	if closeErr := readBuffer.CloseContext("bacnetHostNPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bacnetHostNPort")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetHostNPortEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostNPortEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPortEnclosed")
	}

	// Create the instance
	return &_BACnetHostNPortEnclosed{
		TagNumber:       tagNumber,
		OpeningTag:      openingTag,
		BacnetHostNPort: bacnetHostNPort,
		ClosingTag:      closingTag,
	}, nil
}

func (m *_BACnetHostNPortEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostNPortEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostNPortEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPortEnclosed")
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

	// Simple Field (bacnetHostNPort)
	if pushErr := writeBuffer.PushContext("bacnetHostNPort"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bacnetHostNPort")
	}
	_bacnetHostNPortErr := writeBuffer.WriteSerializable(ctx, m.GetBacnetHostNPort())
	if popErr := writeBuffer.PopContext("bacnetHostNPort"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bacnetHostNPort")
	}
	if _bacnetHostNPortErr != nil {
		return errors.Wrap(_bacnetHostNPortErr, "Error serializing 'bacnetHostNPort' field")
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

	if popErr := writeBuffer.PopContext("BACnetHostNPortEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPortEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetHostNPortEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetHostNPortEnclosed) isBACnetHostNPortEnclosed() bool {
	return true
}

func (m *_BACnetHostNPortEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
