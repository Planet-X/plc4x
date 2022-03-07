/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type COTPPacketConnectionResponse struct {
	*COTPPacket
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass

	// Arguments.
	CotpLen uint16
}

// The corresponding interface
type ICOTPPacketConnectionResponse interface {
	ICOTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketConnectionResponse) GetTpduCode() uint8 {
	return 0xD0
}

func (m *COTPPacketConnectionResponse) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.COTPPacket.Parameters = parameters
	m.COTPPacket.Payload = payload
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *COTPPacketConnectionResponse) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *COTPPacketConnectionResponse) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *COTPPacketConnectionResponse) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCOTPPacketConnectionResponse factory function for COTPPacketConnectionResponse
func NewCOTPPacketConnectionResponse(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, parameters []*COTPParameter, payload *S7Message, cotpLen uint16) *COTPPacket {
	child := &COTPPacketConnectionResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
		COTPPacket:           NewCOTPPacket(parameters, payload, cotpLen),
	}
	child.Child = child
	return child.COTPPacket
}

func CastCOTPPacketConnectionResponse(structType interface{}) *COTPPacketConnectionResponse {
	if casted, ok := structType.(COTPPacketConnectionResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPPacketConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(COTPPacket); ok {
		return CastCOTPPacketConnectionResponse(casted.Child)
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return CastCOTPPacketConnectionResponse(casted.Child)
	}
	return nil
}

func (m *COTPPacketConnectionResponse) GetTypeName() string {
	return "COTPPacketConnectionResponse"
}

func (m *COTPPacketConnectionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPPacketConnectionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPPacketConnectionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketConnectionResponseParse(readBuffer utils.ReadBuffer, cotpLen uint16) (*COTPPacket, error) {
	if pullErr := readBuffer.PullContext("COTPPacketConnectionResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
	_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}
	destinationReference := _destinationReference

	// Simple Field (sourceReference)
	_sourceReference, _sourceReferenceErr := readBuffer.ReadUint16("sourceReference", 16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field")
	}
	sourceReference := _sourceReference

	// Simple Field (protocolClass)
	if pullErr := readBuffer.PullContext("protocolClass"); pullErr != nil {
		return nil, pullErr
	}
	_protocolClass, _protocolClassErr := COTPProtocolClassParse(readBuffer)
	if _protocolClassErr != nil {
		return nil, errors.Wrap(_protocolClassErr, "Error parsing 'protocolClass' field")
	}
	protocolClass := _protocolClass
	if closeErr := readBuffer.CloseContext("protocolClass"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("COTPPacketConnectionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketConnectionResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
		COTPPacket:           &COTPPacket{},
	}
	_child.COTPPacket.Child = _child
	return _child.COTPPacket, nil
}

func (m *COTPPacketConnectionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketConnectionResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationReference)
		destinationReference := uint16(m.DestinationReference)
		_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (sourceReference)
		sourceReference := uint16(m.SourceReference)
		_sourceReferenceErr := writeBuffer.WriteUint16("sourceReference", 16, (sourceReference))
		if _sourceReferenceErr != nil {
			return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
		}

		// Simple Field (protocolClass)
		if pushErr := writeBuffer.PushContext("protocolClass"); pushErr != nil {
			return pushErr
		}
		_protocolClassErr := m.ProtocolClass.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("protocolClass"); popErr != nil {
			return popErr
		}
		if _protocolClassErr != nil {
			return errors.Wrap(_protocolClassErr, "Error serializing 'protocolClass' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketConnectionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPPacketConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
