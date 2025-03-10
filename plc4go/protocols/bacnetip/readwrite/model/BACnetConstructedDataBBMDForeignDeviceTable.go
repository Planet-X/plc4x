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

// BACnetConstructedDataBBMDForeignDeviceTable is the corresponding interface of BACnetConstructedDataBBMDForeignDeviceTable
type BACnetConstructedDataBBMDForeignDeviceTable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBbmdForeignDeviceTable returns BbmdForeignDeviceTable (property field)
	GetBbmdForeignDeviceTable() []BACnetBDTEntry
}

// BACnetConstructedDataBBMDForeignDeviceTableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBBMDForeignDeviceTable.
// This is useful for switch cases.
type BACnetConstructedDataBBMDForeignDeviceTableExactly interface {
	BACnetConstructedDataBBMDForeignDeviceTable
	isBACnetConstructedDataBBMDForeignDeviceTable() bool
}

// _BACnetConstructedDataBBMDForeignDeviceTable is the data-structure of this message
type _BACnetConstructedDataBBMDForeignDeviceTable struct {
	*_BACnetConstructedData
	BbmdForeignDeviceTable []BACnetBDTEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_FOREIGN_DEVICE_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetBbmdForeignDeviceTable() []BACnetBDTEntry {
	return m.BbmdForeignDeviceTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDForeignDeviceTable factory function for _BACnetConstructedDataBBMDForeignDeviceTable
func NewBACnetConstructedDataBBMDForeignDeviceTable(bbmdForeignDeviceTable []BACnetBDTEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBBMDForeignDeviceTable {
	_result := &_BACnetConstructedDataBBMDForeignDeviceTable{
		BbmdForeignDeviceTable: bbmdForeignDeviceTable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDForeignDeviceTable(structType interface{}) BACnetConstructedDataBBMDForeignDeviceTable {
	if casted, ok := structType.(BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDForeignDeviceTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetTypeName() string {
	return "BACnetConstructedDataBBMDForeignDeviceTable"
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.BbmdForeignDeviceTable) > 0 {
		for _, element := range m.BbmdForeignDeviceTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBBMDForeignDeviceTableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDForeignDeviceTable, error) {
	return BACnetConstructedDataBBMDForeignDeviceTableParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBBMDForeignDeviceTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDForeignDeviceTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDForeignDeviceTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDForeignDeviceTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (bbmdForeignDeviceTable)
	if pullErr := readBuffer.PullContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdForeignDeviceTable")
	}
	// Terminated array
	var bbmdForeignDeviceTable []BACnetBDTEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetBDTEntryParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'bbmdForeignDeviceTable' field of BACnetConstructedDataBBMDForeignDeviceTable")
			}
			bbmdForeignDeviceTable = append(bbmdForeignDeviceTable, _item.(BACnetBDTEntry))
		}
	}
	if closeErr := readBuffer.CloseContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdForeignDeviceTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDForeignDeviceTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDForeignDeviceTable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBBMDForeignDeviceTable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BbmdForeignDeviceTable: bbmdForeignDeviceTable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDForeignDeviceTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDForeignDeviceTable")
		}

		// Array Field (bbmdForeignDeviceTable)
		if pushErr := writeBuffer.PushContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bbmdForeignDeviceTable")
		}
		for _curItem, _element := range m.GetBbmdForeignDeviceTable() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetBbmdForeignDeviceTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'bbmdForeignDeviceTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("bbmdForeignDeviceTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bbmdForeignDeviceTable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDForeignDeviceTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDForeignDeviceTable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) isBACnetConstructedDataBBMDForeignDeviceTable() bool {
	return true
}

func (m *_BACnetConstructedDataBBMDForeignDeviceTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
