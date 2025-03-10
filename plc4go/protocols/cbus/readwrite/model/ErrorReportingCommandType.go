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

// ErrorReportingCommandType is an enum
type ErrorReportingCommandType uint8

type IErrorReportingCommandType interface {
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	ErrorReportingCommandType_DEPRECATED        ErrorReportingCommandType = 0x00
	ErrorReportingCommandType_ERROR_REPORT      ErrorReportingCommandType = 0x01
	ErrorReportingCommandType_ACKNOWLEDGE       ErrorReportingCommandType = 0x02
	ErrorReportingCommandType_CLEAR_MOST_SEVERE ErrorReportingCommandType = 0x03
)

var ErrorReportingCommandTypeValues []ErrorReportingCommandType

func init() {
	_ = errors.New
	ErrorReportingCommandTypeValues = []ErrorReportingCommandType{
		ErrorReportingCommandType_DEPRECATED,
		ErrorReportingCommandType_ERROR_REPORT,
		ErrorReportingCommandType_ACKNOWLEDGE,
		ErrorReportingCommandType_CLEAR_MOST_SEVERE,
	}
}

func (e ErrorReportingCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 8
		}
	case 0x01:
		{ /* '0x01' */
			return 8
		}
	case 0x02:
		{ /* '0x02' */
			return 8
		}
	case 0x03:
		{ /* '0x03' */
			return 8
		}
	default:
		{
			return 0
		}
	}
}

func ErrorReportingCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (ErrorReportingCommandType, error) {
	for _, sizeValue := range ErrorReportingCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func ErrorReportingCommandTypeByValue(value uint8) (enum ErrorReportingCommandType, ok bool) {
	switch value {
	case 0x00:
		return ErrorReportingCommandType_DEPRECATED, true
	case 0x01:
		return ErrorReportingCommandType_ERROR_REPORT, true
	case 0x02:
		return ErrorReportingCommandType_ACKNOWLEDGE, true
	case 0x03:
		return ErrorReportingCommandType_CLEAR_MOST_SEVERE, true
	}
	return 0, false
}

func ErrorReportingCommandTypeByName(value string) (enum ErrorReportingCommandType, ok bool) {
	switch value {
	case "DEPRECATED":
		return ErrorReportingCommandType_DEPRECATED, true
	case "ERROR_REPORT":
		return ErrorReportingCommandType_ERROR_REPORT, true
	case "ACKNOWLEDGE":
		return ErrorReportingCommandType_ACKNOWLEDGE, true
	case "CLEAR_MOST_SEVERE":
		return ErrorReportingCommandType_CLEAR_MOST_SEVERE, true
	}
	return 0, false
}

func ErrorReportingCommandTypeKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingCommandType(structType interface{}) ErrorReportingCommandType {
	castFunc := func(typ interface{}) ErrorReportingCommandType {
		if sErrorReportingCommandType, ok := typ.(ErrorReportingCommandType); ok {
			return sErrorReportingCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingCommandTypeParse(ctx context.Context, theBytes []byte) (ErrorReportingCommandType, error) {
	return ErrorReportingCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingCommandType, error) {
	val, err := readBuffer.ReadUint8("ErrorReportingCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingCommandType")
	}
	if enum, ok := ErrorReportingCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ErrorReportingCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingCommandType) PLC4XEnumName() string {
	switch e {
	case ErrorReportingCommandType_DEPRECATED:
		return "DEPRECATED"
	case ErrorReportingCommandType_ERROR_REPORT:
		return "ERROR_REPORT"
	case ErrorReportingCommandType_ACKNOWLEDGE:
		return "ACKNOWLEDGE"
	case ErrorReportingCommandType_CLEAR_MOST_SEVERE:
		return "CLEAR_MOST_SEVERE"
	}
	return ""
}

func (e ErrorReportingCommandType) String() string {
	return e.PLC4XEnumName()
}
