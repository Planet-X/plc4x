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
package org.apache.plc4x.java.opcua.readwrite;

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

public class SignatureData extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "458";
  }

  // Properties.
  protected final PascalString algorithm;
  protected final PascalByteString signature;

  public SignatureData(PascalString algorithm, PascalByteString signature) {
    super();
    this.algorithm = algorithm;
    this.signature = signature;
  }

  public PascalString getAlgorithm() {
    return algorithm;
  }

  public PascalByteString getSignature() {
    return signature;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SignatureData");

    // Simple Field (algorithm)
    writeSimpleField("algorithm", algorithm, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (signature)
    writeSimpleField("signature", signature, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SignatureData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SignatureData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (algorithm)
    lengthInBits += algorithm.getLengthInBits();

    // Simple field (signature)
    lengthInBits += signature.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("SignatureData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString algorithm =
        readSimpleField(
            "algorithm",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalByteString signature =
        readSimpleField(
            "signature",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SignatureData");
    // Create the instance
    return new SignatureDataBuilderImpl(algorithm, signature);
  }

  public static class SignatureDataBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString algorithm;
    private final PascalByteString signature;

    public SignatureDataBuilderImpl(PascalString algorithm, PascalByteString signature) {
      this.algorithm = algorithm;
      this.signature = signature;
    }

    public SignatureData build() {
      SignatureData signatureData = new SignatureData(algorithm, signature);
      return signatureData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SignatureData)) {
      return false;
    }
    SignatureData that = (SignatureData) o;
    return (getAlgorithm() == that.getAlgorithm())
        && (getSignature() == that.getSignature())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAlgorithm(), getSignature());
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
