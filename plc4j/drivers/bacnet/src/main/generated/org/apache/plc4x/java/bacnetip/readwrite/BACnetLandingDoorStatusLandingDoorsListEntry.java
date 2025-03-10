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

public class BACnetLandingDoorStatusLandingDoorsListEntry implements Message {

  // Properties.
  protected final BACnetContextTagUnsignedInteger floorNumber;
  protected final BACnetDoorStatusTagged doorStatus;

  public BACnetLandingDoorStatusLandingDoorsListEntry(
      BACnetContextTagUnsignedInteger floorNumber, BACnetDoorStatusTagged doorStatus) {
    super();
    this.floorNumber = floorNumber;
    this.doorStatus = doorStatus;
  }

  public BACnetContextTagUnsignedInteger getFloorNumber() {
    return floorNumber;
  }

  public BACnetDoorStatusTagged getDoorStatus() {
    return doorStatus;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetLandingDoorStatusLandingDoorsListEntry");

    // Simple Field (floorNumber)
    writeSimpleField("floorNumber", floorNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (doorStatus)
    writeSimpleField("doorStatus", doorStatus, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLandingDoorStatusLandingDoorsListEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetLandingDoorStatusLandingDoorsListEntry _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (floorNumber)
    lengthInBits += floorNumber.getLengthInBits();

    // Simple field (doorStatus)
    lengthInBits += doorStatus.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLandingDoorStatusLandingDoorsListEntry staticParse(
      ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetLandingDoorStatusLandingDoorsListEntry staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetLandingDoorStatusLandingDoorsListEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger floorNumber =
        readSimpleField(
            "floorNumber",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetDoorStatusTagged doorStatus =
        readSimpleField(
            "doorStatus",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetDoorStatusTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetLandingDoorStatusLandingDoorsListEntry");
    // Create the instance
    BACnetLandingDoorStatusLandingDoorsListEntry _bACnetLandingDoorStatusLandingDoorsListEntry;
    _bACnetLandingDoorStatusLandingDoorsListEntry =
        new BACnetLandingDoorStatusLandingDoorsListEntry(floorNumber, doorStatus);
    return _bACnetLandingDoorStatusLandingDoorsListEntry;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLandingDoorStatusLandingDoorsListEntry)) {
      return false;
    }
    BACnetLandingDoorStatusLandingDoorsListEntry that =
        (BACnetLandingDoorStatusLandingDoorsListEntry) o;
    return (getFloorNumber() == that.getFloorNumber())
        && (getDoorStatus() == that.getDoorStatus())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getFloorNumber(), getDoorStatus());
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
