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

public class ServerOnNetwork extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "12191";
  }

  // Properties.
  protected final long recordId;
  protected final PascalString serverName;
  protected final PascalString discoveryUrl;
  protected final int noOfServerCapabilities;
  protected final List<PascalString> serverCapabilities;

  public ServerOnNetwork(
      long recordId,
      PascalString serverName,
      PascalString discoveryUrl,
      int noOfServerCapabilities,
      List<PascalString> serverCapabilities) {
    super();
    this.recordId = recordId;
    this.serverName = serverName;
    this.discoveryUrl = discoveryUrl;
    this.noOfServerCapabilities = noOfServerCapabilities;
    this.serverCapabilities = serverCapabilities;
  }

  public long getRecordId() {
    return recordId;
  }

  public PascalString getServerName() {
    return serverName;
  }

  public PascalString getDiscoveryUrl() {
    return discoveryUrl;
  }

  public int getNoOfServerCapabilities() {
    return noOfServerCapabilities;
  }

  public List<PascalString> getServerCapabilities() {
    return serverCapabilities;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ServerOnNetwork");

    // Simple Field (recordId)
    writeSimpleField("recordId", recordId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (serverName)
    writeSimpleField("serverName", serverName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (discoveryUrl)
    writeSimpleField("discoveryUrl", discoveryUrl, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfServerCapabilities)
    writeSimpleField(
        "noOfServerCapabilities", noOfServerCapabilities, writeSignedInt(writeBuffer, 32));

    // Array Field (serverCapabilities)
    writeComplexTypeArrayField("serverCapabilities", serverCapabilities, writeBuffer);

    writeBuffer.popContext("ServerOnNetwork");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ServerOnNetwork _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (recordId)
    lengthInBits += 32;

    // Simple field (serverName)
    lengthInBits += serverName.getLengthInBits();

    // Simple field (discoveryUrl)
    lengthInBits += discoveryUrl.getLengthInBits();

    // Simple field (noOfServerCapabilities)
    lengthInBits += 32;

    // Array field
    if (serverCapabilities != null) {
      int i = 0;
      for (PascalString element : serverCapabilities) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverCapabilities.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ServerOnNetwork");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long recordId = readSimpleField("recordId", readUnsignedLong(readBuffer, 32));

    PascalString serverName =
        readSimpleField(
            "serverName",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString discoveryUrl =
        readSimpleField(
            "discoveryUrl",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfServerCapabilities =
        readSimpleField("noOfServerCapabilities", readSignedInt(readBuffer, 32));

    List<PascalString> serverCapabilities =
        readCountArrayField(
            "serverCapabilities",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerCapabilities);

    readBuffer.closeContext("ServerOnNetwork");
    // Create the instance
    return new ServerOnNetworkBuilderImpl(
        recordId, serverName, discoveryUrl, noOfServerCapabilities, serverCapabilities);
  }

  public static class ServerOnNetworkBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long recordId;
    private final PascalString serverName;
    private final PascalString discoveryUrl;
    private final int noOfServerCapabilities;
    private final List<PascalString> serverCapabilities;

    public ServerOnNetworkBuilderImpl(
        long recordId,
        PascalString serverName,
        PascalString discoveryUrl,
        int noOfServerCapabilities,
        List<PascalString> serverCapabilities) {
      this.recordId = recordId;
      this.serverName = serverName;
      this.discoveryUrl = discoveryUrl;
      this.noOfServerCapabilities = noOfServerCapabilities;
      this.serverCapabilities = serverCapabilities;
    }

    public ServerOnNetwork build() {
      ServerOnNetwork serverOnNetwork =
          new ServerOnNetwork(
              recordId, serverName, discoveryUrl, noOfServerCapabilities, serverCapabilities);
      return serverOnNetwork;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ServerOnNetwork)) {
      return false;
    }
    ServerOnNetwork that = (ServerOnNetwork) o;
    return (getRecordId() == that.getRecordId())
        && (getServerName() == that.getServerName())
        && (getDiscoveryUrl() == that.getDiscoveryUrl())
        && (getNoOfServerCapabilities() == that.getNoOfServerCapabilities())
        && (getServerCapabilities() == that.getServerCapabilities())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRecordId(),
        getServerName(),
        getDiscoveryUrl(),
        getNoOfServerCapabilities(),
        getServerCapabilities());
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
