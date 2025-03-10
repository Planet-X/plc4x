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

public class BACnetClosingTag implements Message {

  // Properties.
  protected final BACnetTagHeader header;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetClosingTag(BACnetTagHeader header, Short tagNumberArgument) {
    super();
    this.header = header;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetClosingTag");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetClosingTag");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetClosingTag _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetClosingTag staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Short tagNumberArgument;
    if (args[0] instanceof Short) {
      tagNumberArgument = (Short) args[0];
    } else if (args[0] instanceof String) {
      tagNumberArgument = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, tagNumberArgument);
  }

  public static BACnetClosingTag staticParse(ReadBuffer readBuffer, Short tagNumberArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetClosingTag");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getActualTagNumber()) == (tagNumberArgument))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }
    // Validation
    if (!((header.getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS))) {
      throw new ParseValidationException("should be a context tag");
    }
    // Validation
    if (!((header.getLengthValueType()) == (7))) {
      throw new ParseValidationException("closing tag should have a value of 7");
    }

    readBuffer.closeContext("BACnetClosingTag");
    // Create the instance
    BACnetClosingTag _bACnetClosingTag;
    _bACnetClosingTag = new BACnetClosingTag(header, tagNumberArgument);
    return _bACnetClosingTag;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetClosingTag)) {
      return false;
    }
    BACnetClosingTag that = (BACnetClosingTag) o;
    return (getHeader() == that.getHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader());
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
