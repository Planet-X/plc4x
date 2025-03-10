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

public class LogicAssignment implements Message {

  // Properties.
  protected final boolean greaterOfOrLogic;
  protected final boolean reStrikeDelay;
  protected final boolean assignedToGav16;
  protected final boolean assignedToGav15;
  protected final boolean assignedToGav14;
  protected final boolean assignedToGav13;

  // Reserved Fields
  private Boolean reservedField0;
  private Boolean reservedField1;

  public LogicAssignment(
      boolean greaterOfOrLogic,
      boolean reStrikeDelay,
      boolean assignedToGav16,
      boolean assignedToGav15,
      boolean assignedToGav14,
      boolean assignedToGav13) {
    super();
    this.greaterOfOrLogic = greaterOfOrLogic;
    this.reStrikeDelay = reStrikeDelay;
    this.assignedToGav16 = assignedToGav16;
    this.assignedToGav15 = assignedToGav15;
    this.assignedToGav14 = assignedToGav14;
    this.assignedToGav13 = assignedToGav13;
  }

  public boolean getGreaterOfOrLogic() {
    return greaterOfOrLogic;
  }

  public boolean getReStrikeDelay() {
    return reStrikeDelay;
  }

  public boolean getAssignedToGav16() {
    return assignedToGav16;
  }

  public boolean getAssignedToGav15() {
    return assignedToGav15;
  }

  public boolean getAssignedToGav14() {
    return assignedToGav14;
  }

  public boolean getAssignedToGav13() {
    return assignedToGav13;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("LogicAssignment");

    // Simple Field (greaterOfOrLogic)
    writeSimpleField("greaterOfOrLogic", greaterOfOrLogic, writeBoolean(writeBuffer));

    // Simple Field (reStrikeDelay)
    writeSimpleField("reStrikeDelay", reStrikeDelay, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (boolean) false,
        writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (boolean) false,
        writeBoolean(writeBuffer));

    // Simple Field (assignedToGav16)
    writeSimpleField("assignedToGav16", assignedToGav16, writeBoolean(writeBuffer));

    // Simple Field (assignedToGav15)
    writeSimpleField("assignedToGav15", assignedToGav15, writeBoolean(writeBuffer));

    // Simple Field (assignedToGav14)
    writeSimpleField("assignedToGav14", assignedToGav14, writeBoolean(writeBuffer));

    // Simple Field (assignedToGav13)
    writeSimpleField("assignedToGav13", assignedToGav13, writeBoolean(writeBuffer));

    writeBuffer.popContext("LogicAssignment");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LogicAssignment _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (greaterOfOrLogic)
    lengthInBits += 1;

    // Simple field (reStrikeDelay)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (assignedToGav16)
    lengthInBits += 1;

    // Simple field (assignedToGav15)
    lengthInBits += 1;

    // Simple field (assignedToGav14)
    lengthInBits += 1;

    // Simple field (assignedToGav13)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static LogicAssignment staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static LogicAssignment staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LogicAssignment");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean greaterOfOrLogic = readSimpleField("greaterOfOrLogic", readBoolean(readBuffer));

    boolean reStrikeDelay = readSimpleField("reStrikeDelay", readBoolean(readBuffer));

    Boolean reservedField0 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    Boolean reservedField1 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean assignedToGav16 = readSimpleField("assignedToGav16", readBoolean(readBuffer));

    boolean assignedToGav15 = readSimpleField("assignedToGav15", readBoolean(readBuffer));

    boolean assignedToGav14 = readSimpleField("assignedToGav14", readBoolean(readBuffer));

    boolean assignedToGav13 = readSimpleField("assignedToGav13", readBoolean(readBuffer));

    readBuffer.closeContext("LogicAssignment");
    // Create the instance
    LogicAssignment _logicAssignment;
    _logicAssignment =
        new LogicAssignment(
            greaterOfOrLogic,
            reStrikeDelay,
            assignedToGav16,
            assignedToGav15,
            assignedToGav14,
            assignedToGav13);
    _logicAssignment.reservedField0 = reservedField0;
    _logicAssignment.reservedField1 = reservedField1;
    return _logicAssignment;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LogicAssignment)) {
      return false;
    }
    LogicAssignment that = (LogicAssignment) o;
    return (getGreaterOfOrLogic() == that.getGreaterOfOrLogic())
        && (getReStrikeDelay() == that.getReStrikeDelay())
        && (getAssignedToGav16() == that.getAssignedToGav16())
        && (getAssignedToGav15() == that.getAssignedToGav15())
        && (getAssignedToGav14() == that.getAssignedToGav14())
        && (getAssignedToGav13() == that.getAssignedToGav13())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getGreaterOfOrLogic(),
        getReStrikeDelay(),
        getAssignedToGav16(),
        getAssignedToGav15(),
        getAssignedToGav14(),
        getAssignedToGav13());
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
