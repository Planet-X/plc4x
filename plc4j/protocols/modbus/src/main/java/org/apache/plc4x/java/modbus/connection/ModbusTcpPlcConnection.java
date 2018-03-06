/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.modbus.connection;

import io.netty.channel.*;
import org.apache.plc4x.java.base.connection.TcpSocketChannelFactory;
import org.apache.plc4x.java.modbus.netty.ModbusProtocol;
import org.apache.plc4x.java.modbus.netty.ModbusTcpProtocol;
import org.apache.plc4x.java.modbus.netty.Plc4XModbusProtocol;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.util.concurrent.CompletableFuture;

public class ModbusTcpPlcConnection extends BaseModbusPlcConnection {

    private static final int MODBUS_TCP_PORT = 502;

    private static final Logger logger = LoggerFactory.getLogger(ModbusTcpPlcConnection.class);

    public ModbusTcpPlcConnection(InetAddress address, String params) {
        super(new TcpSocketChannelFactory(address, MODBUS_TCP_PORT), params);

        logger.info("Configured ModbusTcpPlcConnection with: host-name {}", address.getHostAddress());
    }

    @Override
    protected ChannelHandler getChannelHandler(CompletableFuture<Void> sessionSetupCompleteFuture) {
        return new ChannelInitializer() {
            @Override
            protected void initChannel(Channel channel) {
                // Build the protocol stack for communicating with the s7 protocol.
                ChannelPipeline pipeline = channel.pipeline();
                pipeline.addLast(new ModbusTcpProtocol());
                pipeline.addLast(new ModbusProtocol());
                pipeline.addLast(new Plc4XModbusProtocol());
            }
        };
    }

}
