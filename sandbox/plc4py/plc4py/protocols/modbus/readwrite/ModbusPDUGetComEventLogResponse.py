#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from ctypes import c_bool
from ctypes import c_byte
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
from typing import List
import math


@dataclass
class ModbusPDUGetComEventLogResponse(PlcMessage, ModbusPDU):
    status: c_uint16
    event_count: c_uint16
    message_count: c_uint16
    events: List[c_byte]
    # Accessors for discriminator values.
    error_flag: c_bool = False
    function_flag: c_uint8 = 0x0C
    response: c_bool = True

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusPDUGetComEventLogResponse")

        # Implicit Field (byte_count) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        byte_count: c_uint8 = c_uint8(((COUNT(self.events())) + (6)))
        write_implicit_field(
            "byteCount", byte_count, write_unsigned_short(write_buffer, 8)
        )

        # Simple Field (status)
        write_simple_field("status", self.status, write_unsigned_int(write_buffer, 16))

        # Simple Field (eventCount)
        write_simple_field(
            "eventCount", self.event_count, write_unsigned_int(write_buffer, 16)
        )

        # Simple Field (messageCount)
        write_simple_field(
            "messageCount", self.message_count, write_unsigned_int(write_buffer, 16)
        )

        # Array Field (events)
        write_byte_array_field("events", self.events, writeByteArray(write_buffer, 8))

        write_buffer.pop_context("ModbusPDUGetComEventLogResponse")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUGetComEventLogResponse = self

        # Implicit Field (byteCount)
        length_in_bits += 8

        # Simple field (status)
        length_in_bits += 16

        # Simple field (eventCount)
        length_in_bits += 16

        # Simple field (messageCount)
        length_in_bits += 16

        # Array field
        if self.events is not None:
            length_in_bits += 8 * self.events.length

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: c_bool):
        read_buffer.pull_context("ModbusPDUGetComEventLogResponse")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        byte_count: c_uint8 = read_implicit_field(
            "byteCount", read_unsigned_short(read_buffer, 8)
        )

        status: c_uint16 = read_simple_field(
            "status", read_unsigned_int(read_buffer, 16)
        )

        event_count: c_uint16 = read_simple_field(
            "eventCount", read_unsigned_int(read_buffer, 16)
        )

        message_count: c_uint16 = read_simple_field(
            "messageCount", read_unsigned_int(read_buffer, 16)
        )

        events: List[c_byte] = read_buffer.read_byte_array(
            "events", int((byteCount) - (6))
        )

        read_buffer.close_context("ModbusPDUGetComEventLogResponse")
        # Create the instance
        return ModbusPDUGetComEventLogResponseBuilder(
            status, event_count, message_count, events
        )

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUGetComEventLogResponse):
            return False

        that: ModbusPDUGetComEventLogResponse = ModbusPDUGetComEventLogResponse(o)
        return (
            (self.status == that.status)
            and (self.event_count == that.event_count)
            and (self.message_count == that.message_count)
            and (self.events == that.events)
            and super().equals(that)
            and True
        )

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUGetComEventLogResponseBuilder(ModbusPDUBuilder):
    status: c_uint16
    eventCount: c_uint16
    messageCount: c_uint16
    events: List[c_byte]

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUGetComEventLogResponse:
        modbus_pdu_get_com_event_log_response: ModbusPDUGetComEventLogResponse = (
            ModbusPDUGetComEventLogResponse(
                self.status, self.event_count, self.message_count, self.events
            )
        )
        return modbus_pdu_get_com_event_log_response
