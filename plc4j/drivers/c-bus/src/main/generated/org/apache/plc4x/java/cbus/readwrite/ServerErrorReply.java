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
package org.apache.plc4x.java.cbus.readwrite;

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

public class ServerErrorReply extends ReplyOrConfirmation implements Message {

  // Accessors for discriminator values.

  // Constant values.
  public static final Byte ERRORMARKER = 0x21;

  // Arguments.
  protected final CBusOptions cBusOptions;
  protected final RequestContext requestContext;

  public ServerErrorReply(byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext) {
    super(peekedByte, cBusOptions, requestContext);
    this.cBusOptions = cBusOptions;
    this.requestContext = requestContext;
  }

  public byte getErrorMarker() {
    return ERRORMARKER;
  }

  @Override
  protected void serializeReplyOrConfirmationChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ServerErrorReply");

    // Const Field (errorMarker)
    writeConstField("errorMarker", ERRORMARKER, writeByte(writeBuffer, 8));

    writeBuffer.popContext("ServerErrorReply");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ServerErrorReply _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (errorMarker)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ReplyOrConfirmationBuilder staticParseReplyOrConfirmationBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions, RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("ServerErrorReply");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte errorMarker =
        readConstField("errorMarker", readByte(readBuffer, 8), ServerErrorReply.ERRORMARKER);

    readBuffer.closeContext("ServerErrorReply");
    // Create the instance
    return new ServerErrorReplyBuilderImpl(cBusOptions, requestContext);
  }

  public static class ServerErrorReplyBuilderImpl
      implements ReplyOrConfirmation.ReplyOrConfirmationBuilder {
    private final CBusOptions cBusOptions;
    private final RequestContext requestContext;

    public ServerErrorReplyBuilderImpl(CBusOptions cBusOptions, RequestContext requestContext) {
      this.cBusOptions = cBusOptions;
      this.requestContext = requestContext;
    }

    public ServerErrorReply build(
        byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext) {
      ServerErrorReply serverErrorReply =
          new ServerErrorReply(peekedByte, cBusOptions, requestContext);
      return serverErrorReply;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ServerErrorReply)) {
      return false;
    }
    ServerErrorReply that = (ServerErrorReply) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
