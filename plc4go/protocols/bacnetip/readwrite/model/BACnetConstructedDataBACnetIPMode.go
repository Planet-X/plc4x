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

// BACnetConstructedDataBACnetIPMode is the corresponding interface of BACnetConstructedDataBACnetIPMode
type BACnetConstructedDataBACnetIPMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBacnetIpMode returns BacnetIpMode (property field)
	GetBacnetIpMode() BACnetIPModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetIPModeTagged
}

// BACnetConstructedDataBACnetIPModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPMode.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPModeExactly interface {
	BACnetConstructedDataBACnetIPMode
	isBACnetConstructedDataBACnetIPMode() bool
}

// _BACnetConstructedDataBACnetIPMode is the data-structure of this message
type _BACnetConstructedDataBACnetIPMode struct {
	*_BACnetConstructedData
	BacnetIpMode BACnetIPModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPMode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetBacnetIpMode() BACnetIPModeTagged {
	return m.BacnetIpMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMode) GetActualValue() BACnetIPModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetIPModeTagged(m.GetBacnetIpMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPMode factory function for _BACnetConstructedDataBACnetIPMode
func NewBACnetConstructedDataBACnetIPMode(bacnetIpMode BACnetIPModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPMode {
	_result := &_BACnetConstructedDataBACnetIPMode{
		BacnetIpMode:           bacnetIpMode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPMode(structType interface{}) BACnetConstructedDataBACnetIPMode {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPMode) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPMode"
}

func (m *_BACnetConstructedDataBACnetIPMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bacnetIpMode)
	lengthInBits += m.BacnetIpMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPModeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPMode, error) {
	return BACnetConstructedDataBACnetIPModeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bacnetIpMode)
	if pullErr := readBuffer.PullContext("bacnetIpMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bacnetIpMode")
	}
	_bacnetIpMode, _bacnetIpModeErr := BACnetIPModeTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _bacnetIpModeErr != nil {
		return nil, errors.Wrap(_bacnetIpModeErr, "Error parsing 'bacnetIpMode' field of BACnetConstructedDataBACnetIPMode")
	}
	bacnetIpMode := _bacnetIpMode.(BACnetIPModeTagged)
	if closeErr := readBuffer.CloseContext("bacnetIpMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bacnetIpMode")
	}

	// Virtual field
	_actualValue := bacnetIpMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPMode{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BacnetIpMode: bacnetIpMode,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPMode")
		}

		// Simple Field (bacnetIpMode)
		if pushErr := writeBuffer.PushContext("bacnetIpMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bacnetIpMode")
		}
		_bacnetIpModeErr := writeBuffer.WriteSerializable(ctx, m.GetBacnetIpMode())
		if popErr := writeBuffer.PopContext("bacnetIpMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bacnetIpMode")
		}
		if _bacnetIpModeErr != nil {
			return errors.Wrap(_bacnetIpModeErr, "Error serializing 'bacnetIpMode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPMode) isBACnetConstructedDataBACnetIPMode() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
