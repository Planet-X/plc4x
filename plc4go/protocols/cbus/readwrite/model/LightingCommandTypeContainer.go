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

// LightingCommandTypeContainer is an enum
type LightingCommandTypeContainer uint8

type ILightingCommandTypeContainer interface {
	utils.Serializable
	NumBytes() uint8
	CommandType() LightingCommandType
}

const (
	LightingCommandTypeContainer_LightingCommandOff                       LightingCommandTypeContainer = 0x01
	LightingCommandTypeContainer_LightingCommandOn                        LightingCommandTypeContainer = 0x79
	LightingCommandTypeContainer_LightingCommandRampToLevel_Instantaneous LightingCommandTypeContainer = 0x02
	LightingCommandTypeContainer_LightingCommandRampToLevel_4Second       LightingCommandTypeContainer = 0x0A
	LightingCommandTypeContainer_LightingCommandRampToLevel_8Second       LightingCommandTypeContainer = 0x12
	LightingCommandTypeContainer_LightingCommandRampToLevel_12Second      LightingCommandTypeContainer = 0x1A
	LightingCommandTypeContainer_LightingCommandRampToLevel_20Second      LightingCommandTypeContainer = 0x22
	LightingCommandTypeContainer_LightingCommandRampToLevel_30Second      LightingCommandTypeContainer = 0x2A
	LightingCommandTypeContainer_LightingCommandRampToLevel_40Second      LightingCommandTypeContainer = 0x32
	LightingCommandTypeContainer_LightingCommandRampToLevel_60Second      LightingCommandTypeContainer = 0x3A
	LightingCommandTypeContainer_LightingCommandRampToLevel_90Second      LightingCommandTypeContainer = 0x42
	LightingCommandTypeContainer_LightingCommandRampToLevel_120Second     LightingCommandTypeContainer = 0x4A
	LightingCommandTypeContainer_LightingCommandRampToLevel_180Second     LightingCommandTypeContainer = 0x52
	LightingCommandTypeContainer_LightingCommandRampToLevel_300Second     LightingCommandTypeContainer = 0x5A
	LightingCommandTypeContainer_LightingCommandRampToLevel_420Second     LightingCommandTypeContainer = 0x62
	LightingCommandTypeContainer_LightingCommandRampToLevel_600Second     LightingCommandTypeContainer = 0x6A
	LightingCommandTypeContainer_LightingCommandRampToLevel_900Second     LightingCommandTypeContainer = 0x72
	LightingCommandTypeContainer_LightingCommandRampToLevel_1020Second    LightingCommandTypeContainer = 0x7A
	LightingCommandTypeContainer_LightingCommandTerminateRamp             LightingCommandTypeContainer = 0x09
	LightingCommandTypeContainer_LightingCommandLabel_0Bytes              LightingCommandTypeContainer = 0xA0
	LightingCommandTypeContainer_LightingCommandLabel_1Bytes              LightingCommandTypeContainer = 0xA1
	LightingCommandTypeContainer_LightingCommandLabel_2Bytes              LightingCommandTypeContainer = 0xA2
	LightingCommandTypeContainer_LightingCommandLabel_3Bytes              LightingCommandTypeContainer = 0xA3
	LightingCommandTypeContainer_LightingCommandLabel_4Bytes              LightingCommandTypeContainer = 0xA4
	LightingCommandTypeContainer_LightingCommandLabel_5Bytes              LightingCommandTypeContainer = 0xA5
	LightingCommandTypeContainer_LightingCommandLabel_6Bytes              LightingCommandTypeContainer = 0xA6
	LightingCommandTypeContainer_LightingCommandLabel_7Bytes              LightingCommandTypeContainer = 0xA7
	LightingCommandTypeContainer_LightingCommandLabel_8Bytes              LightingCommandTypeContainer = 0xA8
	LightingCommandTypeContainer_LightingCommandLabel_9Bytes              LightingCommandTypeContainer = 0xA9
	LightingCommandTypeContainer_LightingCommandLabel_10Bytes             LightingCommandTypeContainer = 0xAA
	LightingCommandTypeContainer_LightingCommandLabel_11Bytes             LightingCommandTypeContainer = 0xAB
	LightingCommandTypeContainer_LightingCommandLabel_12Bytes             LightingCommandTypeContainer = 0xAC
	LightingCommandTypeContainer_LightingCommandLabel_13Bytes             LightingCommandTypeContainer = 0xAD
	LightingCommandTypeContainer_LightingCommandLabel_14Bytes             LightingCommandTypeContainer = 0xAE
	LightingCommandTypeContainer_LightingCommandLabel_15Bytes             LightingCommandTypeContainer = 0xAF
	LightingCommandTypeContainer_LightingCommandLabel_16Bytes             LightingCommandTypeContainer = 0xB0
	LightingCommandTypeContainer_LightingCommandLabel_17Bytes             LightingCommandTypeContainer = 0xB1
	LightingCommandTypeContainer_LightingCommandLabel_18Bytes             LightingCommandTypeContainer = 0xB2
	LightingCommandTypeContainer_LightingCommandLabel_19Bytes             LightingCommandTypeContainer = 0xB3
	LightingCommandTypeContainer_LightingCommandLabel_20Bytes             LightingCommandTypeContainer = 0xB4
	LightingCommandTypeContainer_LightingCommandLabel_21Bytes             LightingCommandTypeContainer = 0xB5
	LightingCommandTypeContainer_LightingCommandLabel_22Bytes             LightingCommandTypeContainer = 0xB6
	LightingCommandTypeContainer_LightingCommandLabel_23Bytes             LightingCommandTypeContainer = 0xB7
	LightingCommandTypeContainer_LightingCommandLabel_24Bytes             LightingCommandTypeContainer = 0xB8
	LightingCommandTypeContainer_LightingCommandLabel_25Bytes             LightingCommandTypeContainer = 0xB9
	LightingCommandTypeContainer_LightingCommandLabel_26Bytes             LightingCommandTypeContainer = 0xBA
	LightingCommandTypeContainer_LightingCommandLabel_27Bytes             LightingCommandTypeContainer = 0xBB
	LightingCommandTypeContainer_LightingCommandLabel_28Bytes             LightingCommandTypeContainer = 0xBC
	LightingCommandTypeContainer_LightingCommandLabel_29Bytes             LightingCommandTypeContainer = 0xBD
	LightingCommandTypeContainer_LightingCommandLabel_30Bytes             LightingCommandTypeContainer = 0xBE
	LightingCommandTypeContainer_LightingCommandLabel_32Bytes             LightingCommandTypeContainer = 0xBF
)

var LightingCommandTypeContainerValues []LightingCommandTypeContainer

func init() {
	_ = errors.New
	LightingCommandTypeContainerValues = []LightingCommandTypeContainer{
		LightingCommandTypeContainer_LightingCommandOff,
		LightingCommandTypeContainer_LightingCommandOn,
		LightingCommandTypeContainer_LightingCommandRampToLevel_Instantaneous,
		LightingCommandTypeContainer_LightingCommandRampToLevel_4Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_8Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_12Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_20Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_30Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_40Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_60Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_90Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_120Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_180Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_300Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_420Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_600Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_900Second,
		LightingCommandTypeContainer_LightingCommandRampToLevel_1020Second,
		LightingCommandTypeContainer_LightingCommandTerminateRamp,
		LightingCommandTypeContainer_LightingCommandLabel_0Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_1Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_2Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_3Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_4Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_5Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_6Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_7Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_8Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_9Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_10Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_11Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_12Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_13Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_14Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_15Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_16Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_17Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_18Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_19Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_20Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_21Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_22Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_23Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_24Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_25Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_26Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_27Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_28Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_29Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_30Bytes,
		LightingCommandTypeContainer_LightingCommandLabel_32Bytes,
	}
}

func (e LightingCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 1
		}
	case 0x09:
		{ /* '0x09' */
			return 1
		}
	case 0x0A:
		{ /* '0x0A' */
			return 2
		}
	case 0x12:
		{ /* '0x12' */
			return 2
		}
	case 0x1A:
		{ /* '0x1A' */
			return 2
		}
	case 0x22:
		{ /* '0x22' */
			return 2
		}
	case 0x2A:
		{ /* '0x2A' */
			return 2
		}
	case 0x32:
		{ /* '0x32' */
			return 2
		}
	case 0x3A:
		{ /* '0x3A' */
			return 2
		}
	case 0x42:
		{ /* '0x42' */
			return 2
		}
	case 0x4A:
		{ /* '0x4A' */
			return 2
		}
	case 0x52:
		{ /* '0x52' */
			return 2
		}
	case 0x5A:
		{ /* '0x5A' */
			return 2
		}
	case 0x62:
		{ /* '0x62' */
			return 2
		}
	case 0x6A:
		{ /* '0x6A' */
			return 2
		}
	case 0x72:
		{ /* '0x72' */
			return 2
		}
	case 0x79:
		{ /* '0x79' */
			return 1
		}
	case 0x7A:
		{ /* '0x7A' */
			return 2
		}
	case 0xA0:
		{ /* '0xA0' */
			return 0
		}
	case 0xA1:
		{ /* '0xA1' */
			return 1
		}
	case 0xA2:
		{ /* '0xA2' */
			return 2
		}
	case 0xA3:
		{ /* '0xA3' */
			return 3
		}
	case 0xA4:
		{ /* '0xA4' */
			return 4
		}
	case 0xA5:
		{ /* '0xA5' */
			return 5
		}
	case 0xA6:
		{ /* '0xA6' */
			return 6
		}
	case 0xA7:
		{ /* '0xA7' */
			return 7
		}
	case 0xA8:
		{ /* '0xA8' */
			return 8
		}
	case 0xA9:
		{ /* '0xA9' */
			return 9
		}
	case 0xAA:
		{ /* '0xAA' */
			return 10
		}
	case 0xAB:
		{ /* '0xAB' */
			return 11
		}
	case 0xAC:
		{ /* '0xAC' */
			return 12
		}
	case 0xAD:
		{ /* '0xAD' */
			return 13
		}
	case 0xAE:
		{ /* '0xAE' */
			return 14
		}
	case 0xAF:
		{ /* '0xAF' */
			return 15
		}
	case 0xB0:
		{ /* '0xB0' */
			return 16
		}
	case 0xB1:
		{ /* '0xB1' */
			return 17
		}
	case 0xB2:
		{ /* '0xB2' */
			return 18
		}
	case 0xB3:
		{ /* '0xB3' */
			return 19
		}
	case 0xB4:
		{ /* '0xB4' */
			return 20
		}
	case 0xB5:
		{ /* '0xB5' */
			return 21
		}
	case 0xB6:
		{ /* '0xB6' */
			return 22
		}
	case 0xB7:
		{ /* '0xB7' */
			return 23
		}
	case 0xB8:
		{ /* '0xB8' */
			return 24
		}
	case 0xB9:
		{ /* '0xB9' */
			return 25
		}
	case 0xBA:
		{ /* '0xBA' */
			return 26
		}
	case 0xBB:
		{ /* '0xBB' */
			return 27
		}
	case 0xBC:
		{ /* '0xBC' */
			return 28
		}
	case 0xBD:
		{ /* '0xBD' */
			return 29
		}
	case 0xBE:
		{ /* '0xBE' */
			return 30
		}
	case 0xBF:
		{ /* '0xBF' */
			return 31
		}
	default:
		{
			return 0
		}
	}
}

func LightingCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (LightingCommandTypeContainer, error) {
	for _, sizeValue := range LightingCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e LightingCommandTypeContainer) CommandType() LightingCommandType {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return LightingCommandType_OFF
		}
	case 0x02:
		{ /* '0x02' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x09:
		{ /* '0x09' */
			return LightingCommandType_TERMINATE_RAMP
		}
	case 0x0A:
		{ /* '0x0A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x12:
		{ /* '0x12' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x1A:
		{ /* '0x1A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x22:
		{ /* '0x22' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x2A:
		{ /* '0x2A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x32:
		{ /* '0x32' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x3A:
		{ /* '0x3A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x42:
		{ /* '0x42' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x4A:
		{ /* '0x4A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x52:
		{ /* '0x52' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x5A:
		{ /* '0x5A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x62:
		{ /* '0x62' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x6A:
		{ /* '0x6A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x72:
		{ /* '0x72' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0x79:
		{ /* '0x79' */
			return LightingCommandType_ON
		}
	case 0x7A:
		{ /* '0x7A' */
			return LightingCommandType_RAMP_TO_LEVEL
		}
	case 0xA0:
		{ /* '0xA0' */
			return LightingCommandType_LABEL
		}
	case 0xA1:
		{ /* '0xA1' */
			return LightingCommandType_LABEL
		}
	case 0xA2:
		{ /* '0xA2' */
			return LightingCommandType_LABEL
		}
	case 0xA3:
		{ /* '0xA3' */
			return LightingCommandType_LABEL
		}
	case 0xA4:
		{ /* '0xA4' */
			return LightingCommandType_LABEL
		}
	case 0xA5:
		{ /* '0xA5' */
			return LightingCommandType_LABEL
		}
	case 0xA6:
		{ /* '0xA6' */
			return LightingCommandType_LABEL
		}
	case 0xA7:
		{ /* '0xA7' */
			return LightingCommandType_LABEL
		}
	case 0xA8:
		{ /* '0xA8' */
			return LightingCommandType_LABEL
		}
	case 0xA9:
		{ /* '0xA9' */
			return LightingCommandType_LABEL
		}
	case 0xAA:
		{ /* '0xAA' */
			return LightingCommandType_LABEL
		}
	case 0xAB:
		{ /* '0xAB' */
			return LightingCommandType_LABEL
		}
	case 0xAC:
		{ /* '0xAC' */
			return LightingCommandType_LABEL
		}
	case 0xAD:
		{ /* '0xAD' */
			return LightingCommandType_LABEL
		}
	case 0xAE:
		{ /* '0xAE' */
			return LightingCommandType_LABEL
		}
	case 0xAF:
		{ /* '0xAF' */
			return LightingCommandType_LABEL
		}
	case 0xB0:
		{ /* '0xB0' */
			return LightingCommandType_LABEL
		}
	case 0xB1:
		{ /* '0xB1' */
			return LightingCommandType_LABEL
		}
	case 0xB2:
		{ /* '0xB2' */
			return LightingCommandType_LABEL
		}
	case 0xB3:
		{ /* '0xB3' */
			return LightingCommandType_LABEL
		}
	case 0xB4:
		{ /* '0xB4' */
			return LightingCommandType_LABEL
		}
	case 0xB5:
		{ /* '0xB5' */
			return LightingCommandType_LABEL
		}
	case 0xB6:
		{ /* '0xB6' */
			return LightingCommandType_LABEL
		}
	case 0xB7:
		{ /* '0xB7' */
			return LightingCommandType_LABEL
		}
	case 0xB8:
		{ /* '0xB8' */
			return LightingCommandType_LABEL
		}
	case 0xB9:
		{ /* '0xB9' */
			return LightingCommandType_LABEL
		}
	case 0xBA:
		{ /* '0xBA' */
			return LightingCommandType_LABEL
		}
	case 0xBB:
		{ /* '0xBB' */
			return LightingCommandType_LABEL
		}
	case 0xBC:
		{ /* '0xBC' */
			return LightingCommandType_LABEL
		}
	case 0xBD:
		{ /* '0xBD' */
			return LightingCommandType_LABEL
		}
	case 0xBE:
		{ /* '0xBE' */
			return LightingCommandType_LABEL
		}
	case 0xBF:
		{ /* '0xBF' */
			return LightingCommandType_LABEL
		}
	default:
		{
			return 0
		}
	}
}

func LightingCommandTypeContainerFirstEnumForFieldCommandType(value LightingCommandType) (LightingCommandTypeContainer, error) {
	for _, sizeValue := range LightingCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func LightingCommandTypeContainerByValue(value uint8) (enum LightingCommandTypeContainer, ok bool) {
	switch value {
	case 0x01:
		return LightingCommandTypeContainer_LightingCommandOff, true
	case 0x02:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_Instantaneous, true
	case 0x09:
		return LightingCommandTypeContainer_LightingCommandTerminateRamp, true
	case 0x0A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_4Second, true
	case 0x12:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_8Second, true
	case 0x1A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_12Second, true
	case 0x22:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_20Second, true
	case 0x2A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_30Second, true
	case 0x32:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_40Second, true
	case 0x3A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_60Second, true
	case 0x42:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_90Second, true
	case 0x4A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_120Second, true
	case 0x52:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_180Second, true
	case 0x5A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_300Second, true
	case 0x62:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_420Second, true
	case 0x6A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_600Second, true
	case 0x72:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_900Second, true
	case 0x79:
		return LightingCommandTypeContainer_LightingCommandOn, true
	case 0x7A:
		return LightingCommandTypeContainer_LightingCommandRampToLevel_1020Second, true
	case 0xA0:
		return LightingCommandTypeContainer_LightingCommandLabel_0Bytes, true
	case 0xA1:
		return LightingCommandTypeContainer_LightingCommandLabel_1Bytes, true
	case 0xA2:
		return LightingCommandTypeContainer_LightingCommandLabel_2Bytes, true
	case 0xA3:
		return LightingCommandTypeContainer_LightingCommandLabel_3Bytes, true
	case 0xA4:
		return LightingCommandTypeContainer_LightingCommandLabel_4Bytes, true
	case 0xA5:
		return LightingCommandTypeContainer_LightingCommandLabel_5Bytes, true
	case 0xA6:
		return LightingCommandTypeContainer_LightingCommandLabel_6Bytes, true
	case 0xA7:
		return LightingCommandTypeContainer_LightingCommandLabel_7Bytes, true
	case 0xA8:
		return LightingCommandTypeContainer_LightingCommandLabel_8Bytes, true
	case 0xA9:
		return LightingCommandTypeContainer_LightingCommandLabel_9Bytes, true
	case 0xAA:
		return LightingCommandTypeContainer_LightingCommandLabel_10Bytes, true
	case 0xAB:
		return LightingCommandTypeContainer_LightingCommandLabel_11Bytes, true
	case 0xAC:
		return LightingCommandTypeContainer_LightingCommandLabel_12Bytes, true
	case 0xAD:
		return LightingCommandTypeContainer_LightingCommandLabel_13Bytes, true
	case 0xAE:
		return LightingCommandTypeContainer_LightingCommandLabel_14Bytes, true
	case 0xAF:
		return LightingCommandTypeContainer_LightingCommandLabel_15Bytes, true
	case 0xB0:
		return LightingCommandTypeContainer_LightingCommandLabel_16Bytes, true
	case 0xB1:
		return LightingCommandTypeContainer_LightingCommandLabel_17Bytes, true
	case 0xB2:
		return LightingCommandTypeContainer_LightingCommandLabel_18Bytes, true
	case 0xB3:
		return LightingCommandTypeContainer_LightingCommandLabel_19Bytes, true
	case 0xB4:
		return LightingCommandTypeContainer_LightingCommandLabel_20Bytes, true
	case 0xB5:
		return LightingCommandTypeContainer_LightingCommandLabel_21Bytes, true
	case 0xB6:
		return LightingCommandTypeContainer_LightingCommandLabel_22Bytes, true
	case 0xB7:
		return LightingCommandTypeContainer_LightingCommandLabel_23Bytes, true
	case 0xB8:
		return LightingCommandTypeContainer_LightingCommandLabel_24Bytes, true
	case 0xB9:
		return LightingCommandTypeContainer_LightingCommandLabel_25Bytes, true
	case 0xBA:
		return LightingCommandTypeContainer_LightingCommandLabel_26Bytes, true
	case 0xBB:
		return LightingCommandTypeContainer_LightingCommandLabel_27Bytes, true
	case 0xBC:
		return LightingCommandTypeContainer_LightingCommandLabel_28Bytes, true
	case 0xBD:
		return LightingCommandTypeContainer_LightingCommandLabel_29Bytes, true
	case 0xBE:
		return LightingCommandTypeContainer_LightingCommandLabel_30Bytes, true
	case 0xBF:
		return LightingCommandTypeContainer_LightingCommandLabel_32Bytes, true
	}
	return 0, false
}

func LightingCommandTypeContainerByName(value string) (enum LightingCommandTypeContainer, ok bool) {
	switch value {
	case "LightingCommandOff":
		return LightingCommandTypeContainer_LightingCommandOff, true
	case "LightingCommandRampToLevel_Instantaneous":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_Instantaneous, true
	case "LightingCommandTerminateRamp":
		return LightingCommandTypeContainer_LightingCommandTerminateRamp, true
	case "LightingCommandRampToLevel_4Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_4Second, true
	case "LightingCommandRampToLevel_8Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_8Second, true
	case "LightingCommandRampToLevel_12Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_12Second, true
	case "LightingCommandRampToLevel_20Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_20Second, true
	case "LightingCommandRampToLevel_30Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_30Second, true
	case "LightingCommandRampToLevel_40Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_40Second, true
	case "LightingCommandRampToLevel_60Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_60Second, true
	case "LightingCommandRampToLevel_90Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_90Second, true
	case "LightingCommandRampToLevel_120Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_120Second, true
	case "LightingCommandRampToLevel_180Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_180Second, true
	case "LightingCommandRampToLevel_300Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_300Second, true
	case "LightingCommandRampToLevel_420Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_420Second, true
	case "LightingCommandRampToLevel_600Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_600Second, true
	case "LightingCommandRampToLevel_900Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_900Second, true
	case "LightingCommandOn":
		return LightingCommandTypeContainer_LightingCommandOn, true
	case "LightingCommandRampToLevel_1020Second":
		return LightingCommandTypeContainer_LightingCommandRampToLevel_1020Second, true
	case "LightingCommandLabel_0Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_0Bytes, true
	case "LightingCommandLabel_1Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_1Bytes, true
	case "LightingCommandLabel_2Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_2Bytes, true
	case "LightingCommandLabel_3Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_3Bytes, true
	case "LightingCommandLabel_4Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_4Bytes, true
	case "LightingCommandLabel_5Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_5Bytes, true
	case "LightingCommandLabel_6Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_6Bytes, true
	case "LightingCommandLabel_7Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_7Bytes, true
	case "LightingCommandLabel_8Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_8Bytes, true
	case "LightingCommandLabel_9Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_9Bytes, true
	case "LightingCommandLabel_10Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_10Bytes, true
	case "LightingCommandLabel_11Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_11Bytes, true
	case "LightingCommandLabel_12Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_12Bytes, true
	case "LightingCommandLabel_13Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_13Bytes, true
	case "LightingCommandLabel_14Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_14Bytes, true
	case "LightingCommandLabel_15Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_15Bytes, true
	case "LightingCommandLabel_16Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_16Bytes, true
	case "LightingCommandLabel_17Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_17Bytes, true
	case "LightingCommandLabel_18Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_18Bytes, true
	case "LightingCommandLabel_19Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_19Bytes, true
	case "LightingCommandLabel_20Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_20Bytes, true
	case "LightingCommandLabel_21Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_21Bytes, true
	case "LightingCommandLabel_22Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_22Bytes, true
	case "LightingCommandLabel_23Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_23Bytes, true
	case "LightingCommandLabel_24Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_24Bytes, true
	case "LightingCommandLabel_25Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_25Bytes, true
	case "LightingCommandLabel_26Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_26Bytes, true
	case "LightingCommandLabel_27Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_27Bytes, true
	case "LightingCommandLabel_28Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_28Bytes, true
	case "LightingCommandLabel_29Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_29Bytes, true
	case "LightingCommandLabel_30Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_30Bytes, true
	case "LightingCommandLabel_32Bytes":
		return LightingCommandTypeContainer_LightingCommandLabel_32Bytes, true
	}
	return 0, false
}

func LightingCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range LightingCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLightingCommandTypeContainer(structType interface{}) LightingCommandTypeContainer {
	castFunc := func(typ interface{}) LightingCommandTypeContainer {
		if sLightingCommandTypeContainer, ok := typ.(LightingCommandTypeContainer); ok {
			return sLightingCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m LightingCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m LightingCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingCommandTypeContainerParse(ctx context.Context, theBytes []byte) (LightingCommandTypeContainer, error) {
	return LightingCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("LightingCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LightingCommandTypeContainer")
	}
	if enum, ok := LightingCommandTypeContainerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return LightingCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e LightingCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LightingCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("LightingCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LightingCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case LightingCommandTypeContainer_LightingCommandOff:
		return "LightingCommandOff"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_Instantaneous:
		return "LightingCommandRampToLevel_Instantaneous"
	case LightingCommandTypeContainer_LightingCommandTerminateRamp:
		return "LightingCommandTerminateRamp"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_4Second:
		return "LightingCommandRampToLevel_4Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_8Second:
		return "LightingCommandRampToLevel_8Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_12Second:
		return "LightingCommandRampToLevel_12Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_20Second:
		return "LightingCommandRampToLevel_20Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_30Second:
		return "LightingCommandRampToLevel_30Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_40Second:
		return "LightingCommandRampToLevel_40Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_60Second:
		return "LightingCommandRampToLevel_60Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_90Second:
		return "LightingCommandRampToLevel_90Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_120Second:
		return "LightingCommandRampToLevel_120Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_180Second:
		return "LightingCommandRampToLevel_180Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_300Second:
		return "LightingCommandRampToLevel_300Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_420Second:
		return "LightingCommandRampToLevel_420Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_600Second:
		return "LightingCommandRampToLevel_600Second"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_900Second:
		return "LightingCommandRampToLevel_900Second"
	case LightingCommandTypeContainer_LightingCommandOn:
		return "LightingCommandOn"
	case LightingCommandTypeContainer_LightingCommandRampToLevel_1020Second:
		return "LightingCommandRampToLevel_1020Second"
	case LightingCommandTypeContainer_LightingCommandLabel_0Bytes:
		return "LightingCommandLabel_0Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_1Bytes:
		return "LightingCommandLabel_1Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_2Bytes:
		return "LightingCommandLabel_2Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_3Bytes:
		return "LightingCommandLabel_3Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_4Bytes:
		return "LightingCommandLabel_4Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_5Bytes:
		return "LightingCommandLabel_5Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_6Bytes:
		return "LightingCommandLabel_6Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_7Bytes:
		return "LightingCommandLabel_7Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_8Bytes:
		return "LightingCommandLabel_8Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_9Bytes:
		return "LightingCommandLabel_9Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_10Bytes:
		return "LightingCommandLabel_10Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_11Bytes:
		return "LightingCommandLabel_11Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_12Bytes:
		return "LightingCommandLabel_12Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_13Bytes:
		return "LightingCommandLabel_13Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_14Bytes:
		return "LightingCommandLabel_14Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_15Bytes:
		return "LightingCommandLabel_15Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_16Bytes:
		return "LightingCommandLabel_16Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_17Bytes:
		return "LightingCommandLabel_17Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_18Bytes:
		return "LightingCommandLabel_18Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_19Bytes:
		return "LightingCommandLabel_19Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_20Bytes:
		return "LightingCommandLabel_20Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_21Bytes:
		return "LightingCommandLabel_21Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_22Bytes:
		return "LightingCommandLabel_22Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_23Bytes:
		return "LightingCommandLabel_23Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_24Bytes:
		return "LightingCommandLabel_24Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_25Bytes:
		return "LightingCommandLabel_25Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_26Bytes:
		return "LightingCommandLabel_26Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_27Bytes:
		return "LightingCommandLabel_27Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_28Bytes:
		return "LightingCommandLabel_28Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_29Bytes:
		return "LightingCommandLabel_29Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_30Bytes:
		return "LightingCommandLabel_30Bytes"
	case LightingCommandTypeContainer_LightingCommandLabel_32Bytes:
		return "LightingCommandLabel_32Bytes"
	}
	return ""
}

func (e LightingCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
