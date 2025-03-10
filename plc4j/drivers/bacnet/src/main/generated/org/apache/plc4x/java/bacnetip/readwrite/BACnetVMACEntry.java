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

public class BACnetVMACEntry implements Message {

  // Properties.
  protected final BACnetContextTagOctetString virtualMacAddress;
  protected final BACnetContextTagOctetString nativeMacAddress;

  public BACnetVMACEntry(
      BACnetContextTagOctetString virtualMacAddress, BACnetContextTagOctetString nativeMacAddress) {
    super();
    this.virtualMacAddress = virtualMacAddress;
    this.nativeMacAddress = nativeMacAddress;
  }

  public BACnetContextTagOctetString getVirtualMacAddress() {
    return virtualMacAddress;
  }

  public BACnetContextTagOctetString getNativeMacAddress() {
    return nativeMacAddress;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetVMACEntry");

    // Optional Field (virtualMacAddress) (Can be skipped, if the value is null)
    writeOptionalField(
        "virtualMacAddress", virtualMacAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (nativeMacAddress) (Can be skipped, if the value is null)
    writeOptionalField(
        "nativeMacAddress", nativeMacAddress, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetVMACEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetVMACEntry _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (virtualMacAddress)
    if (virtualMacAddress != null) {
      lengthInBits += virtualMacAddress.getLengthInBits();
    }

    // Optional Field (nativeMacAddress)
    if (nativeMacAddress != null) {
      lengthInBits += nativeMacAddress.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetVMACEntry staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetVMACEntry staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetVMACEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagOctetString virtualMacAddress =
        readOptionalField(
            "virtualMacAddress",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagOctetString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.OCTET_STRING)),
                readBuffer));

    BACnetContextTagOctetString nativeMacAddress =
        readOptionalField(
            "nativeMacAddress",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagOctetString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.OCTET_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetVMACEntry");
    // Create the instance
    BACnetVMACEntry _bACnetVMACEntry;
    _bACnetVMACEntry = new BACnetVMACEntry(virtualMacAddress, nativeMacAddress);
    return _bACnetVMACEntry;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetVMACEntry)) {
      return false;
    }
    BACnetVMACEntry that = (BACnetVMACEntry) o;
    return (getVirtualMacAddress() == that.getVirtualMacAddress())
        && (getNativeMacAddress() == that.getNativeMacAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getVirtualMacAddress(), getNativeMacAddress());
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
