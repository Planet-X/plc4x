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
package org.apache.plc4x.java.abeth.readwrite;

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

public abstract class DF1RequestMessage implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getCommandCode();

  // Properties.
  protected final short destinationAddress;
  protected final short sourceAddress;
  protected final short status;
  protected final int transactionCounter;

  public DF1RequestMessage(
      short destinationAddress, short sourceAddress, short status, int transactionCounter) {
    super();
    this.destinationAddress = destinationAddress;
    this.sourceAddress = sourceAddress;
    this.status = status;
    this.transactionCounter = transactionCounter;
  }

  public short getDestinationAddress() {
    return destinationAddress;
  }

  public short getSourceAddress() {
    return sourceAddress;
  }

  public short getStatus() {
    return status;
  }

  public int getTransactionCounter() {
    return transactionCounter;
  }

  protected abstract void serializeDF1RequestMessageChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DF1RequestMessage");

    // Simple Field (destinationAddress)
    writeSimpleField("destinationAddress", destinationAddress, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (sourceAddress)
    writeSimpleField("sourceAddress", sourceAddress, writeUnsignedShort(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x0000, writeUnsignedInt(writeBuffer, 16));

    // Discriminator Field (commandCode) (Used as input to a switch field)
    writeDiscriminatorField("commandCode", getCommandCode(), writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (transactionCounter)
    writeSimpleField("transactionCounter", transactionCounter, writeUnsignedInt(writeBuffer, 16));

    // Switch field (Serialize the sub-type)
    serializeDF1RequestMessageChild(writeBuffer);

    writeBuffer.popContext("DF1RequestMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DF1RequestMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (destinationAddress)
    lengthInBits += 8;

    // Simple field (sourceAddress)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Discriminator Field (commandCode)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Simple field (transactionCounter)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static DF1RequestMessage staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DF1RequestMessage staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DF1RequestMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short destinationAddress =
        readSimpleField("destinationAddress", readUnsignedShort(readBuffer, 8));

    short sourceAddress = readSimpleField("sourceAddress", readUnsignedShort(readBuffer, 8));

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    short commandCode = readDiscriminatorField("commandCode", readUnsignedShort(readBuffer, 8));

    short status = readSimpleField("status", readUnsignedShort(readBuffer, 8));

    int transactionCounter = readSimpleField("transactionCounter", readUnsignedInt(readBuffer, 16));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    DF1RequestMessageBuilder builder = null;
    if (EvaluationHelper.equals(commandCode, (short) 0x0F)) {
      builder = DF1CommandRequestMessage.staticParseDF1RequestMessageBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandCode="
              + commandCode
              + "]");
    }

    readBuffer.closeContext("DF1RequestMessage");
    // Create the instance
    DF1RequestMessage _dF1RequestMessage =
        builder.build(destinationAddress, sourceAddress, status, transactionCounter);
    return _dF1RequestMessage;
  }

  public interface DF1RequestMessageBuilder {
    DF1RequestMessage build(
        short destinationAddress, short sourceAddress, short status, int transactionCounter);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DF1RequestMessage)) {
      return false;
    }
    DF1RequestMessage that = (DF1RequestMessage) o;
    return (getDestinationAddress() == that.getDestinationAddress())
        && (getSourceAddress() == that.getSourceAddress())
        && (getStatus() == that.getStatus())
        && (getTransactionCounter() == that.getTransactionCounter())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getDestinationAddress(), getSourceAddress(), getStatus(), getTransactionCounter());
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
