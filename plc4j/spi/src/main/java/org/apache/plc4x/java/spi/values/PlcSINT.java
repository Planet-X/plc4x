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
package org.apache.plc4x.java.spi.values;

import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.math.BigDecimal;
import java.math.BigInteger;

public class PlcSINT extends PlcIECValue<Byte> {

    private static final String VALUE_OUT_OF_RANGE = "Value of type %s is out of range %d - %d for a %s Value";
    static Byte minValue = Byte.MIN_VALUE;
    static Byte maxValue = Byte.MAX_VALUE;

    public static PlcSINT of(Object value) {
        if (value instanceof Boolean) {
            return new PlcSINT((Boolean) value);
        } else if (value instanceof Byte) {
            return new PlcSINT((Byte) value);
        } else if (value instanceof Short) {
            return new PlcSINT((Short) value);
        } else if (value instanceof Integer) {
            return new PlcSINT((Integer) value);
        } else if (value instanceof Long) {
            return new PlcSINT((Long) value);
        } else if (value instanceof Float) {
            return new PlcSINT((Float) value);
        } else if (value instanceof Double) {
            return new PlcSINT((Double) value);
        } else if (value instanceof BigInteger) {
            return new PlcSINT((BigInteger) value);
        } else if (value instanceof BigDecimal) {
            return new PlcSINT((BigDecimal) value);
        } else {
            return new PlcSINT((String) value);
        }
    }

    public PlcSINT(Boolean value) {
        this.value = value ? Byte.valueOf((byte) 1) : Byte.valueOf((byte) 0);
        this.isNullable = false;
    }

    public PlcSINT(Byte value) {
        this.value = value;
        this.isNullable = false;
    }

    public PlcSINT(Short value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = false;
    }

    public PlcSINT(Integer value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = false;
    }

    public PlcSINT(Long value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = false;
    }

    public PlcSINT(Float value) {
        if ((value < minValue) || (value > maxValue) || (value % 1 != 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = false;
    }

    public PlcSINT(Double value) {
        if ((value < minValue) || (value > maxValue) || (value % 1 != 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = false;
    }

    public PlcSINT(BigInteger value) {
        if ((value.compareTo(BigInteger.valueOf(minValue)) < 0) || (value.compareTo(BigInteger.valueOf(maxValue)) > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = true;
    }

    public PlcSINT(BigDecimal value) {
        if ((value.compareTo(BigDecimal.valueOf(minValue)) < 0) || (value.compareTo(BigDecimal.valueOf(maxValue)) > 0) || (value.scale() > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.byteValue();
        this.isNullable = true;
    }

    public PlcSINT(String value) {
        try {
            this.value = Byte.valueOf(value.trim());
            this.isNullable = false;
        } catch (Exception e) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()), e);
        }
    }

    public PlcSINT(byte value) {
        this.value = value;
        this.isNullable = false;
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.SINT;
    }

    @Override
    public boolean isBoolean() {
        return true;
    }

    @Override
    public boolean getBoolean() {
        return (value != null) && !value.equals((byte) 0);
    }

    @Override
    public boolean isByte() {
        return true;
    }

    @Override
    public byte getByte() {
        return value;
    }

    @Override
    public boolean isShort() {
        return true;
    }

    @Override
    public short getShort() {
        return value.shortValue();
    }

    @Override
    public boolean isInteger() {
        return true;
    }

    @Override
    public int getInteger() {
        return value.intValue();
    }

    @Override
    public boolean isLong() {
        return true;
    }

    @Override
    public long getLong() {
        return value.longValue();
    }

    @Override
    public boolean isBigInteger() {
        return true;
    }

    @Override
    public BigInteger getBigInteger() {
        return BigInteger.valueOf(getLong());
    }

    @Override
    public boolean isFloat() {
        return true;
    }

    @Override
    public float getFloat() {
        return value.floatValue();
    }

    @Override
    public boolean isDouble() {
        return true;
    }

    @Override
    public double getDouble() {
        return value.doubleValue();
    }

    @Override
    public boolean isBigDecimal() {
        return true;
    }

    @Override
    public BigDecimal getBigDecimal() {
        return BigDecimal.valueOf(getFloat());
    }

    @Override
    public boolean isString() {
        return true;
    }

    @Override
    public String getString() {
        return toString();
    }

    @Override
    public String toString() {
        return Byte.toString(value);
    }

    public byte[] getBytes() {
        byte[] bytes = new byte[1];
        bytes[0] = (byte) (value & 0xff);
        return bytes;
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        writeBuffer.writeByte(getClass().getSimpleName(), value);
    }

}
