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

public abstract class NLM implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getMessageType();

  // Arguments.
  protected final Integer apduLength;

  public NLM(Integer apduLength) {
    super();
    this.apduLength = apduLength;
  }

  public boolean getIsVendorProprietaryMessage() {
    return (boolean) ((getMessageType()) >= (128));
  }

  protected abstract void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NLM");

    // Discriminator Field (messageType) (Used as input to a switch field)
    writeDiscriminatorField("messageType", getMessageType(), writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isVendorProprietaryMessage = getIsVendorProprietaryMessage();
    writeBuffer.writeVirtual("isVendorProprietaryMessage", isVendorProprietaryMessage);

    // Switch field (Serialize the sub-type)
    serializeNLMChild(writeBuffer);

    writeBuffer.popContext("NLM");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    NLM _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (messageType)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static NLM staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Integer apduLength;
    if (args[0] instanceof Integer) {
      apduLength = (Integer) args[0];
    } else if (args[0] instanceof String) {
      apduLength = Integer.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Integer or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, apduLength);
  }

  public static NLM staticParse(ReadBuffer readBuffer, Integer apduLength) throws ParseException {
    readBuffer.pullContext("NLM");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short messageType = readDiscriminatorField("messageType", readUnsignedShort(readBuffer, 8));
    boolean isVendorProprietaryMessage =
        readVirtualField("isVendorProprietaryMessage", boolean.class, (messageType) >= (128));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    NLMBuilder builder = null;
    if (EvaluationHelper.equals(messageType, (short) 0x00)) {
      builder = NLMWhoIsRouterToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x01)) {
      builder = NLMIAmRouterToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x02)) {
      builder = NLMICouldBeRouterToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x03)) {
      builder = NLMRejectRouterToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x04)) {
      builder = NLMRouterBusyToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x05)) {
      builder = NLMRouterAvailableToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x06)) {
      builder = NLMInitalizeRoutingTable.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x07)) {
      builder = NLMInitalizeRoutingTableAck.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x08)) {
      builder = NLMEstablishConnectionToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x09)) {
      builder = NLMDisconnectConnectionToNetwork.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0A)) {
      builder = NLMChallengeRequest.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0B)) {
      builder = NLMSecurityPayload.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0C)) {
      builder = NLMSecurityResponse.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0D)) {
      builder = NLMRequestKeyUpdate.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0E)) {
      builder = NLMUpdateKeyUpdate.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x0F)) {
      builder = NLMUpdateKeyDistributionKey.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x10)) {
      builder = NLMRequestMasterKey.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x11)) {
      builder = NLMSetMasterKey.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x12)) {
      builder = NLMWhatIsNetworkNumber.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (EvaluationHelper.equals(messageType, (short) 0x13)) {
      builder = NLMNetworkNumberIs.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (true && EvaluationHelper.equals(isVendorProprietaryMessage, (boolean) false)) {
      builder = NLMReserved.staticParseNLMBuilder(readBuffer, apduLength);
    } else if (true) {
      builder = NLMVendorProprietaryMessage.staticParseNLMBuilder(readBuffer, apduLength);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "messageType="
              + messageType
              + " "
              + "isVendorProprietaryMessage="
              + isVendorProprietaryMessage
              + "]");
    }

    readBuffer.closeContext("NLM");
    // Create the instance
    NLM _nLM = builder.build(apduLength);

    return _nLM;
  }

  public interface NLMBuilder {
    NLM build(Integer apduLength);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLM)) {
      return false;
    }
    NLM that = (NLM) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
