/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dubbo

import (
	"mosn.io/mosn/pkg/protocol"
	"mosn.io/mosn/pkg/protocol/xprotocol"
	"mosn.io/mosn/pkg/types"
)

type Header struct {
	Magic   []byte
	Flag    byte
	Status  byte
	Id      uint64
	DataLen uint32

	Event           int // 1 mean ping
	TwoWay          int // 1 mean req & resp pair
	Direction       int // 1 mean req
	SerializationId int // 2 mean hessian
	protocol.CommonHeader
}

type Frame struct {
	Header
	rawData []byte // raw data
	payload []byte // raw payload

	data    types.IoBuffer // wrapper of data
	content types.IoBuffer // wrapper of payload
}

// ~ XFrame
func (r *Frame) GetRequestId() uint64 {
	return r.Header.Id
}

func (r *Frame) SetRequestId(id uint64) {
	r.Header.Id = id
}

func (r *Frame) IsHeartbeatFrame() bool {
	return r.Header.Event != 0
}

func (r *Frame) GetStreamType() xprotocol.StreamType {
	switch r.Direction {
	case EventRequest:
		return xprotocol.Request
	case EventResponse:
		return xprotocol.Response
	default:
		return xprotocol.Request
	}
}

func (r *Frame) GetHeader() types.HeaderMap {
	return r
}

func (r *Frame) GetData() types.IoBuffer {
	return r.content
}

func (r *Frame) SetData(data types.IoBuffer) {
	r.content = data
}

func (r *Frame) GetStatusCode() uint32 {
	return uint32(r.Header.Status)
}
