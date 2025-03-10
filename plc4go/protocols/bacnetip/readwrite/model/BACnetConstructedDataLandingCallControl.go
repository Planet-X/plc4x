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

// BACnetConstructedDataLandingCallControl is the corresponding interface of BACnetConstructedDataLandingCallControl
type BACnetConstructedDataLandingCallControl interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLandingCallControl returns LandingCallControl (property field)
	GetLandingCallControl() BACnetLandingCallStatus
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLandingCallStatus
}

// BACnetConstructedDataLandingCallControlExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLandingCallControl.
// This is useful for switch cases.
type BACnetConstructedDataLandingCallControlExactly interface {
	BACnetConstructedDataLandingCallControl
	isBACnetConstructedDataLandingCallControl() bool
}

// _BACnetConstructedDataLandingCallControl is the data-structure of this message
type _BACnetConstructedDataLandingCallControl struct {
	*_BACnetConstructedData
	LandingCallControl BACnetLandingCallStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLandingCallControl) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LANDING_CALL_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLandingCallControl) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLandingCallControl) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetLandingCallControl() BACnetLandingCallStatus {
	return m.LandingCallControl
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetActualValue() BACnetLandingCallStatus {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLandingCallStatus(m.GetLandingCallControl())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLandingCallControl factory function for _BACnetConstructedDataLandingCallControl
func NewBACnetConstructedDataLandingCallControl(landingCallControl BACnetLandingCallStatus, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLandingCallControl {
	_result := &_BACnetConstructedDataLandingCallControl{
		LandingCallControl:     landingCallControl,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLandingCallControl(structType interface{}) BACnetConstructedDataLandingCallControl {
	if casted, ok := structType.(BACnetConstructedDataLandingCallControl); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLandingCallControl); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLandingCallControl) GetTypeName() string {
	return "BACnetConstructedDataLandingCallControl"
}

func (m *_BACnetConstructedDataLandingCallControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (landingCallControl)
	lengthInBits += m.LandingCallControl.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLandingCallControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLandingCallControlParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLandingCallControl, error) {
	return BACnetConstructedDataLandingCallControlParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLandingCallControlParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLandingCallControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLandingCallControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLandingCallControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (landingCallControl)
	if pullErr := readBuffer.PullContext("landingCallControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingCallControl")
	}
	_landingCallControl, _landingCallControlErr := BACnetLandingCallStatusParseWithBuffer(ctx, readBuffer)
	if _landingCallControlErr != nil {
		return nil, errors.Wrap(_landingCallControlErr, "Error parsing 'landingCallControl' field of BACnetConstructedDataLandingCallControl")
	}
	landingCallControl := _landingCallControl.(BACnetLandingCallStatus)
	if closeErr := readBuffer.CloseContext("landingCallControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingCallControl")
	}

	// Virtual field
	_actualValue := landingCallControl
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLandingCallControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLandingCallControl")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLandingCallControl{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LandingCallControl: landingCallControl,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLandingCallControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLandingCallControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLandingCallControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLandingCallControl")
		}

		// Simple Field (landingCallControl)
		if pushErr := writeBuffer.PushContext("landingCallControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for landingCallControl")
		}
		_landingCallControlErr := writeBuffer.WriteSerializable(ctx, m.GetLandingCallControl())
		if popErr := writeBuffer.PopContext("landingCallControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for landingCallControl")
		}
		if _landingCallControlErr != nil {
			return errors.Wrap(_landingCallControlErr, "Error serializing 'landingCallControl' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLandingCallControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLandingCallControl")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLandingCallControl) isBACnetConstructedDataLandingCallControl() bool {
	return true
}

func (m *_BACnetConstructedDataLandingCallControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
