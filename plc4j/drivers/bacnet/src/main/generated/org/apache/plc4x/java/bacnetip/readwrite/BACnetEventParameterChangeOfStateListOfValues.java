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

public class BACnetEventParameterChangeOfStateListOfValues implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final List<BACnetPropertyStates> listOfValues;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public BACnetEventParameterChangeOfStateListOfValues(
      BACnetOpeningTag openingTag,
      List<BACnetPropertyStates> listOfValues,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.listOfValues = listOfValues;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public List<BACnetPropertyStates> getListOfValues() {
    return listOfValues;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventParameterChangeOfStateListOfValues");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Array Field (listOfValues)
    writeComplexTypeArrayField("listOfValues", listOfValues, writeBuffer);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterChangeOfStateListOfValues");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetEventParameterChangeOfStateListOfValues _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Array field
    if (listOfValues != null) {
      for (Message element : listOfValues) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterChangeOfStateListOfValues staticParse(
      ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Short tagNumber;
    if (args[0] instanceof Short) {
      tagNumber = (Short) args[0];
    } else if (args[0] instanceof String) {
      tagNumber = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, tagNumber);
  }

  public static BACnetEventParameterChangeOfStateListOfValues staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterChangeOfStateListOfValues");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    List<BACnetPropertyStates> listOfValues =
        readTerminatedArrayField(
            "listOfValues",
            new DataReaderComplexDefault<>(
                () -> BACnetPropertyStates.staticParse(readBuffer), readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterChangeOfStateListOfValues");
    // Create the instance
    BACnetEventParameterChangeOfStateListOfValues _bACnetEventParameterChangeOfStateListOfValues;
    _bACnetEventParameterChangeOfStateListOfValues =
        new BACnetEventParameterChangeOfStateListOfValues(
            openingTag, listOfValues, closingTag, tagNumber);
    return _bACnetEventParameterChangeOfStateListOfValues;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterChangeOfStateListOfValues)) {
      return false;
    }
    BACnetEventParameterChangeOfStateListOfValues that =
        (BACnetEventParameterChangeOfStateListOfValues) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getListOfValues() == that.getListOfValues())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getListOfValues(), getClosingTag());
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
