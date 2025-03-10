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

// LightingDataLabel is the corresponding interface of LightingDataLabel
type LightingDataLabel interface {
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLabelOptions returns LabelOptions (property field)
	GetLabelOptions() LightingLabelOptions
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
}

// LightingDataLabelExactly can be used when we want exactly this type and not a type which fulfills LightingDataLabel.
// This is useful for switch cases.
type LightingDataLabelExactly interface {
	LightingDataLabel
	isLightingDataLabel() bool
}

// _LightingDataLabel is the data-structure of this message
type _LightingDataLabel struct {
	*_LightingData
	Group        byte
	LabelOptions LightingLabelOptions
	Language     *Language
	Data         []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataLabel) InitializeParent(parent LightingData, commandTypeContainer LightingCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_LightingDataLabel) GetParent() LightingData {
	return m._LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataLabel) GetGroup() byte {
	return m.Group
}

func (m *_LightingDataLabel) GetLabelOptions() LightingLabelOptions {
	return m.LabelOptions
}

func (m *_LightingDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_LightingDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataLabel factory function for _LightingDataLabel
func NewLightingDataLabel(group byte, labelOptions LightingLabelOptions, language *Language, data []byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataLabel {
	_result := &_LightingDataLabel{
		Group:         group,
		LabelOptions:  labelOptions,
		Language:      language,
		Data:          data,
		_LightingData: NewLightingData(commandTypeContainer),
	}
	_result._LightingData._LightingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataLabel(structType interface{}) LightingDataLabel {
	if casted, ok := structType.(LightingDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataLabel) GetTypeName() string {
	return "LightingDataLabel"
}

func (m *_LightingDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	// Simple field (labelOptions)
	lengthInBits += m.LabelOptions.GetLengthInBits(ctx)

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_LightingDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingDataLabelParse(theBytes []byte, commandTypeContainer LightingCommandTypeContainer) (LightingDataLabel, error) {
	return LightingDataLabelParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func LightingDataLabelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer LightingCommandTypeContainer) (LightingDataLabel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field of LightingDataLabel")
	}
	group := _group

	// Simple Field (labelOptions)
	if pullErr := readBuffer.PullContext("labelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for labelOptions")
	}
	_labelOptions, _labelOptionsErr := LightingLabelOptionsParseWithBuffer(ctx, readBuffer)
	if _labelOptionsErr != nil {
		return nil, errors.Wrap(_labelOptionsErr, "Error parsing 'labelOptions' field of LightingDataLabel")
	}
	labelOptions := _labelOptions.(LightingLabelOptions)
	if closeErr := readBuffer.CloseContext("labelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for labelOptions")
	}

	// Optional Field (language) (Can be skipped, if a given expression evaluates to false)
	var language *Language = nil
	if bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON)) {
		if pullErr := readBuffer.PullContext("language"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for language")
		}
		_val, _err := LanguageParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'language' field of LightingDataLabel")
		}
		language = &_val
		if closeErr := readBuffer.CloseContext("language"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for language")
		}
	}
	// Byte Array field (data)
	numberOfBytesdata := int((uint16(commandTypeContainer.NumBytes()) - uint16((utils.InlineIf((bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON))), func() interface{} { return uint16((uint16(3))) }, func() interface{} { return uint16((uint16(2))) }).(uint16)))))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of LightingDataLabel")
	}

	if closeErr := readBuffer.CloseContext("LightingDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataLabel")
	}

	// Create a partially initialized instance
	_child := &_LightingDataLabel{
		_LightingData: &_LightingData{},
		Group:         group,
		LabelOptions:  labelOptions,
		Language:      language,
		Data:          data,
	}
	_child._LightingData._LightingDataChildRequirements = _child
	return _child, nil
}

func (m *_LightingDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataLabel")
		}

		// Simple Field (group)
		group := byte(m.GetGroup())
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		// Simple Field (labelOptions)
		if pushErr := writeBuffer.PushContext("labelOptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for labelOptions")
		}
		_labelOptionsErr := writeBuffer.WriteSerializable(ctx, m.GetLabelOptions())
		if popErr := writeBuffer.PopContext("labelOptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for labelOptions")
		}
		if _labelOptionsErr != nil {
			return errors.Wrap(_labelOptionsErr, "Error serializing 'labelOptions' field")
		}

		// Optional Field (language) (Can be skipped, if the value is null)
		var language *Language = nil
		if m.GetLanguage() != nil {
			if pushErr := writeBuffer.PushContext("language"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for language")
			}
			language = m.GetLanguage()
			_languageErr := writeBuffer.WriteSerializable(ctx, language)
			if popErr := writeBuffer.PopContext("language"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for language")
			}
			if _languageErr != nil {
				return errors.Wrap(_languageErr, "Error serializing 'language' field")
			}
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataLabel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataLabel) isLightingDataLabel() bool {
	return true
}

func (m *_LightingDataLabel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
