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
	"encoding/binary"
	"fmt"

	"github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DefaultPlcWriteRequestResult struct {
	Request  model.PlcWriteRequest
	Response model.PlcWriteResponse
	Err      error
}

func (d *DefaultPlcWriteRequestResult) GetRequest() model.PlcWriteRequest {
	return d.Request
}

func (d *DefaultPlcWriteRequestResult) GetResponse() model.PlcWriteResponse {
	return d.Response
}

func (d *DefaultPlcWriteRequestResult) GetErr() error {
	return d.Err
}

func (d *DefaultPlcWriteRequestResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcWriteRequestResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcWriteRequestResult"); err != nil {
		return err
	}

	if d.Request != nil {
		if serializableField, ok := d.Request.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("request"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("request"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.Request)
			if err := writeBuffer.WriteString("request", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if d.Response != nil {
		if serializableField, ok := d.Response.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("response"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("response"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.Response)
			if err := writeBuffer.WriteString("response", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if d.Err != nil {
		if err := writeBuffer.WriteString("err", uint32(len(d.Err.Error())*8), "UTF-8", d.Err.Error()); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("PlcWriteRequestResult"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcWriteRequestResult) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
