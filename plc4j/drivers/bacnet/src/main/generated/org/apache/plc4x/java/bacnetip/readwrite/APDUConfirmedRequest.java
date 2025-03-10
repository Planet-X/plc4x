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

public class APDUConfirmedRequest extends APDU implements Message {

  // Accessors for discriminator values.
  public ApduType getApduType() {
    return ApduType.CONFIRMED_REQUEST_PDU;
  }

  // Properties.
  protected final boolean segmentedMessage;
  protected final boolean moreFollows;
  protected final boolean segmentedResponseAccepted;
  protected final MaxSegmentsAccepted maxSegmentsAccepted;
  protected final MaxApduLengthAccepted maxApduLengthAccepted;
  protected final short invokeId;
  protected final Short sequenceNumber;
  protected final Short proposedWindowSize;
  protected final BACnetConfirmedServiceRequest serviceRequest;
  protected final BACnetConfirmedServiceChoice segmentServiceChoice;
  protected final byte[] segment;

  // Arguments.
  protected final Integer apduLength;
  // Reserved Fields
  private Byte reservedField0;

  public APDUConfirmedRequest(
      boolean segmentedMessage,
      boolean moreFollows,
      boolean segmentedResponseAccepted,
      MaxSegmentsAccepted maxSegmentsAccepted,
      MaxApduLengthAccepted maxApduLengthAccepted,
      short invokeId,
      Short sequenceNumber,
      Short proposedWindowSize,
      BACnetConfirmedServiceRequest serviceRequest,
      BACnetConfirmedServiceChoice segmentServiceChoice,
      byte[] segment,
      Integer apduLength) {
    super(apduLength);
    this.segmentedMessage = segmentedMessage;
    this.moreFollows = moreFollows;
    this.segmentedResponseAccepted = segmentedResponseAccepted;
    this.maxSegmentsAccepted = maxSegmentsAccepted;
    this.maxApduLengthAccepted = maxApduLengthAccepted;
    this.invokeId = invokeId;
    this.sequenceNumber = sequenceNumber;
    this.proposedWindowSize = proposedWindowSize;
    this.serviceRequest = serviceRequest;
    this.segmentServiceChoice = segmentServiceChoice;
    this.segment = segment;
    this.apduLength = apduLength;
  }

  public boolean getSegmentedMessage() {
    return segmentedMessage;
  }

  public boolean getMoreFollows() {
    return moreFollows;
  }

  public boolean getSegmentedResponseAccepted() {
    return segmentedResponseAccepted;
  }

  public MaxSegmentsAccepted getMaxSegmentsAccepted() {
    return maxSegmentsAccepted;
  }

  public MaxApduLengthAccepted getMaxApduLengthAccepted() {
    return maxApduLengthAccepted;
  }

  public short getInvokeId() {
    return invokeId;
  }

  public Short getSequenceNumber() {
    return sequenceNumber;
  }

  public Short getProposedWindowSize() {
    return proposedWindowSize;
  }

  public BACnetConfirmedServiceRequest getServiceRequest() {
    return serviceRequest;
  }

  public BACnetConfirmedServiceChoice getSegmentServiceChoice() {
    return segmentServiceChoice;
  }

  public byte[] getSegment() {
    return segment;
  }

  public int getApduHeaderReduction() {
    return (int) ((3) + ((((getSegmentedMessage()) ? 2 : 0))));
  }

  public int getSegmentReduction() {
    return (int)
        (((((getSegmentServiceChoice()) != (null)))
            ? ((getApduHeaderReduction()) + (1))
            : getApduHeaderReduction()));
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("APDUConfirmedRequest");

    // Simple Field (segmentedMessage)
    writeSimpleField("segmentedMessage", segmentedMessage, writeBoolean(writeBuffer));

    // Simple Field (moreFollows)
    writeSimpleField("moreFollows", moreFollows, writeBoolean(writeBuffer));

    // Simple Field (segmentedResponseAccepted)
    writeSimpleField(
        "segmentedResponseAccepted", segmentedResponseAccepted, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0,
        writeUnsignedByte(writeBuffer, 2));

    // Simple Field (maxSegmentsAccepted)
    writeSimpleEnumField(
        "maxSegmentsAccepted",
        "MaxSegmentsAccepted",
        maxSegmentsAccepted,
        new DataWriterEnumDefault<>(
            MaxSegmentsAccepted::getValue,
            MaxSegmentsAccepted::name,
            writeUnsignedByte(writeBuffer, 3)));

    // Simple Field (maxApduLengthAccepted)
    writeSimpleEnumField(
        "maxApduLengthAccepted",
        "MaxApduLengthAccepted",
        maxApduLengthAccepted,
        new DataWriterEnumDefault<>(
            MaxApduLengthAccepted::getValue,
            MaxApduLengthAccepted::name,
            writeUnsignedByte(writeBuffer, 4)));

    // Simple Field (invokeId)
    writeSimpleField("invokeId", invokeId, writeUnsignedShort(writeBuffer, 8));

    // Optional Field (sequenceNumber) (Can be skipped, if the value is null)
    writeOptionalField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedShort(writeBuffer, 8),
        getSegmentedMessage());

    // Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
    writeOptionalField(
        "proposedWindowSize",
        proposedWindowSize,
        writeUnsignedShort(writeBuffer, 8),
        getSegmentedMessage());

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    int apduHeaderReduction = getApduHeaderReduction();
    writeBuffer.writeVirtual("apduHeaderReduction", apduHeaderReduction);

    // Optional Field (serviceRequest) (Can be skipped, if the value is null)
    writeOptionalField(
        "serviceRequest",
        serviceRequest,
        new DataWriterComplexDefault<>(writeBuffer),
        !(getSegmentedMessage()));

    // Optional Field (segmentServiceChoice) (Can be skipped, if the value is null)
    writeOptionalEnumField(
        "segmentServiceChoice",
        "BACnetConfirmedServiceChoice",
        segmentServiceChoice,
        new DataWriterEnumDefault<>(
            BACnetConfirmedServiceChoice::getValue,
            BACnetConfirmedServiceChoice::name,
            writeUnsignedShort(writeBuffer, 8)),
        (getSegmentedMessage()) && ((getSequenceNumber()) != (0)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    int segmentReduction = getSegmentReduction();
    writeBuffer.writeVirtual("segmentReduction", segmentReduction);

    // Array Field (segment)
    writeByteArrayField("segment", segment, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("APDUConfirmedRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUConfirmedRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (segmentedMessage)
    lengthInBits += 1;

    // Simple field (moreFollows)
    lengthInBits += 1;

    // Simple field (segmentedResponseAccepted)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (maxSegmentsAccepted)
    lengthInBits += 3;

    // Simple field (maxApduLengthAccepted)
    lengthInBits += 4;

    // Simple field (invokeId)
    lengthInBits += 8;

    // Optional Field (sequenceNumber)
    if (sequenceNumber != null) {
      lengthInBits += 8;
    }

    // Optional Field (proposedWindowSize)
    if (proposedWindowSize != null) {
      lengthInBits += 8;
    }

    // A virtual field doesn't have any in- or output.

    // Optional Field (serviceRequest)
    if (serviceRequest != null) {
      lengthInBits += serviceRequest.getLengthInBits();
    }

    // Optional Field (segmentServiceChoice)
    if (segmentServiceChoice != null) {
      lengthInBits += 8;
    }

    // A virtual field doesn't have any in- or output.

    // Array field
    if (segment != null) {
      lengthInBits += 8 * segment.length;
    }

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("APDUConfirmedRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean segmentedMessage = readSimpleField("segmentedMessage", readBoolean(readBuffer));

    boolean moreFollows = readSimpleField("moreFollows", readBoolean(readBuffer));

    boolean segmentedResponseAccepted =
        readSimpleField("segmentedResponseAccepted", readBoolean(readBuffer));

    Byte reservedField0 = readReservedField("reserved", readUnsignedByte(readBuffer, 2), (byte) 0);

    MaxSegmentsAccepted maxSegmentsAccepted =
        readEnumField(
            "maxSegmentsAccepted",
            "MaxSegmentsAccepted",
            new DataReaderEnumDefault<>(
                MaxSegmentsAccepted::enumForValue, readUnsignedByte(readBuffer, 3)));

    MaxApduLengthAccepted maxApduLengthAccepted =
        readEnumField(
            "maxApduLengthAccepted",
            "MaxApduLengthAccepted",
            new DataReaderEnumDefault<>(
                MaxApduLengthAccepted::enumForValue, readUnsignedByte(readBuffer, 4)));

    short invokeId = readSimpleField("invokeId", readUnsignedShort(readBuffer, 8));

    Short sequenceNumber =
        readOptionalField("sequenceNumber", readUnsignedShort(readBuffer, 8), segmentedMessage);

    Short proposedWindowSize =
        readOptionalField("proposedWindowSize", readUnsignedShort(readBuffer, 8), segmentedMessage);
    int apduHeaderReduction =
        readVirtualField("apduHeaderReduction", int.class, (3) + ((((segmentedMessage) ? 2 : 0))));

    BACnetConfirmedServiceRequest serviceRequest =
        readOptionalField(
            "serviceRequest",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConfirmedServiceRequest.staticParse(
                        readBuffer, (long) ((apduLength) - (apduHeaderReduction))),
                readBuffer),
            !(segmentedMessage));
    // Validation
    if (!((((!(segmentedMessage)) && ((serviceRequest) != (null)))) || (segmentedMessage))) {
      throw new ParseValidationException("service request should be set");
    }

    BACnetConfirmedServiceChoice segmentServiceChoice =
        readOptionalField(
            "segmentServiceChoice",
            new DataReaderEnumDefault<>(
                BACnetConfirmedServiceChoice::enumForValue, readUnsignedShort(readBuffer, 8)),
            (segmentedMessage) && ((sequenceNumber) != (0)));
    int segmentReduction =
        readVirtualField(
            "segmentReduction",
            int.class,
            ((((segmentServiceChoice) != (null)))
                ? ((apduHeaderReduction) + (1))
                : apduHeaderReduction));

    byte[] segment =
        readBuffer.readByteArray(
            "segment",
            Math.toIntExact(
                ((segmentedMessage)
                    ? (((((apduLength) > (0))) ? ((apduLength) - (segmentReduction)) : 0))
                    : 0)));

    readBuffer.closeContext("APDUConfirmedRequest");
    // Create the instance
    return new APDUConfirmedRequestBuilderImpl(
        segmentedMessage,
        moreFollows,
        segmentedResponseAccepted,
        maxSegmentsAccepted,
        maxApduLengthAccepted,
        invokeId,
        sequenceNumber,
        proposedWindowSize,
        serviceRequest,
        segmentServiceChoice,
        segment,
        apduLength,
        reservedField0);
  }

  public static class APDUConfirmedRequestBuilderImpl implements APDU.APDUBuilder {
    private final boolean segmentedMessage;
    private final boolean moreFollows;
    private final boolean segmentedResponseAccepted;
    private final MaxSegmentsAccepted maxSegmentsAccepted;
    private final MaxApduLengthAccepted maxApduLengthAccepted;
    private final short invokeId;
    private final Short sequenceNumber;
    private final Short proposedWindowSize;
    private final BACnetConfirmedServiceRequest serviceRequest;
    private final BACnetConfirmedServiceChoice segmentServiceChoice;
    private final byte[] segment;
    private final Integer apduLength;
    private final Byte reservedField0;

    public APDUConfirmedRequestBuilderImpl(
        boolean segmentedMessage,
        boolean moreFollows,
        boolean segmentedResponseAccepted,
        MaxSegmentsAccepted maxSegmentsAccepted,
        MaxApduLengthAccepted maxApduLengthAccepted,
        short invokeId,
        Short sequenceNumber,
        Short proposedWindowSize,
        BACnetConfirmedServiceRequest serviceRequest,
        BACnetConfirmedServiceChoice segmentServiceChoice,
        byte[] segment,
        Integer apduLength,
        Byte reservedField0) {
      this.segmentedMessage = segmentedMessage;
      this.moreFollows = moreFollows;
      this.segmentedResponseAccepted = segmentedResponseAccepted;
      this.maxSegmentsAccepted = maxSegmentsAccepted;
      this.maxApduLengthAccepted = maxApduLengthAccepted;
      this.invokeId = invokeId;
      this.sequenceNumber = sequenceNumber;
      this.proposedWindowSize = proposedWindowSize;
      this.serviceRequest = serviceRequest;
      this.segmentServiceChoice = segmentServiceChoice;
      this.segment = segment;
      this.apduLength = apduLength;
      this.reservedField0 = reservedField0;
    }

    public APDUConfirmedRequest build(Integer apduLength) {

      APDUConfirmedRequest aPDUConfirmedRequest =
          new APDUConfirmedRequest(
              segmentedMessage,
              moreFollows,
              segmentedResponseAccepted,
              maxSegmentsAccepted,
              maxApduLengthAccepted,
              invokeId,
              sequenceNumber,
              proposedWindowSize,
              serviceRequest,
              segmentServiceChoice,
              segment,
              apduLength);
      aPDUConfirmedRequest.reservedField0 = reservedField0;
      return aPDUConfirmedRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUConfirmedRequest)) {
      return false;
    }
    APDUConfirmedRequest that = (APDUConfirmedRequest) o;
    return (getSegmentedMessage() == that.getSegmentedMessage())
        && (getMoreFollows() == that.getMoreFollows())
        && (getSegmentedResponseAccepted() == that.getSegmentedResponseAccepted())
        && (getMaxSegmentsAccepted() == that.getMaxSegmentsAccepted())
        && (getMaxApduLengthAccepted() == that.getMaxApduLengthAccepted())
        && (getInvokeId() == that.getInvokeId())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getProposedWindowSize() == that.getProposedWindowSize())
        && (getServiceRequest() == that.getServiceRequest())
        && (getSegmentServiceChoice() == that.getSegmentServiceChoice())
        && (getSegment() == that.getSegment())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSegmentedMessage(),
        getMoreFollows(),
        getSegmentedResponseAccepted(),
        getMaxSegmentsAccepted(),
        getMaxApduLengthAccepted(),
        getInvokeId(),
        getSequenceNumber(),
        getProposedWindowSize(),
        getServiceRequest(),
        getSegmentServiceChoice(),
        getSegment());
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
