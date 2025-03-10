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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class APDUSimpleAck extends APDU implements Message {

  // Accessors for discriminator values.
  public ApduType getApduType() {
    return ApduType.SIMPLE_ACK_PDU;
  }

  // Properties.
  protected final short originalInvokeId;
  protected final BACnetConfirmedServiceChoice serviceChoice;

  // Arguments.
  protected final Integer apduLength;
  // Reserved Fields
  private Byte reservedField0;

  public APDUSimpleAck(
      short originalInvokeId, BACnetConfirmedServiceChoice serviceChoice, Integer apduLength) {
    super(apduLength);
    this.originalInvokeId = originalInvokeId;
    this.serviceChoice = serviceChoice;
    this.apduLength = apduLength;
  }

  public short getOriginalInvokeId() {
    return originalInvokeId;
  }

  public BACnetConfirmedServiceChoice getServiceChoice() {
    return serviceChoice;
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("APDUSimpleAck");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0,
        writeUnsignedByte(writeBuffer, 4));

    // Simple Field (originalInvokeId)
    writeSimpleField("originalInvokeId", originalInvokeId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (serviceChoice)
    writeSimpleEnumField(
        "serviceChoice",
        "BACnetConfirmedServiceChoice",
        serviceChoice,
        new DataWriterEnumDefault<>(
            BACnetConfirmedServiceChoice::getValue,
            BACnetConfirmedServiceChoice::name,
            writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("APDUSimpleAck");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUSimpleAck _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 4;

    // Simple field (originalInvokeId)
    lengthInBits += 8;

    // Simple field (serviceChoice)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("APDUSimpleAck");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readUnsignedByte(readBuffer, 4), (byte) 0);

    short originalInvokeId = readSimpleField("originalInvokeId", readUnsignedShort(readBuffer, 8));

    BACnetConfirmedServiceChoice serviceChoice =
        readEnumField(
            "serviceChoice",
            "BACnetConfirmedServiceChoice",
            new DataReaderEnumDefault<>(
                BACnetConfirmedServiceChoice::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("APDUSimpleAck");
    // Create the instance
    return new APDUSimpleAckBuilderImpl(
        originalInvokeId, serviceChoice, apduLength, reservedField0);
  }

  public static class APDUSimpleAckBuilderImpl implements APDU.APDUBuilder {
    private final short originalInvokeId;
    private final BACnetConfirmedServiceChoice serviceChoice;
    private final Integer apduLength;
    private final Byte reservedField0;

    public APDUSimpleAckBuilderImpl(
        short originalInvokeId,
        BACnetConfirmedServiceChoice serviceChoice,
        Integer apduLength,
        Byte reservedField0) {
      this.originalInvokeId = originalInvokeId;
      this.serviceChoice = serviceChoice;
      this.apduLength = apduLength;
      this.reservedField0 = reservedField0;
    }

    public APDUSimpleAck build(Integer apduLength) {

      APDUSimpleAck aPDUSimpleAck = new APDUSimpleAck(originalInvokeId, serviceChoice, apduLength);
      aPDUSimpleAck.reservedField0 = reservedField0;
      return aPDUSimpleAck;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUSimpleAck)) {
      return false;
    }
    APDUSimpleAck that = (APDUSimpleAck) o;
    return (getOriginalInvokeId() == that.getOriginalInvokeId())
        && (getServiceChoice() == that.getServiceChoice())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOriginalInvokeId(), getServiceChoice());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
