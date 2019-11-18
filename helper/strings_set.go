// Copyright 2019 pangush
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package helper

import "encoding/json"

const OPTIONS_EX = "EX"// seconds – 设置键key的过期时间，单位时秒
const OPTIONS_PX = "PX"// milliseconds – 设置键key的过期时间，单位时毫秒
const OPTIONS_NX = "NX"// 只有键key不存在的时候才会设置key的值
const OPTIONS_XX = "XX"// 只有键key存在的时候才会设置key的值

// Set redis strings set value
func (strings *Strings) SetString(value string, args ...interface{}) error {
	conn := strings.keys.engine.conn
	key := strings.keys.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetBytes(value []byte, args ...interface{}) error {
	conn := strings.keys.engine.conn
	key := strings.keys.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetUint64(value uint64, args ...interface{}) error {
	conn := strings.keys.engine.conn
	key := strings.keys.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetStruct(bean interface{}, args ...interface{}) error {
	conn := strings.keys.engine.conn
	key := strings.keys.engine.key

	jsonBytes, err := json.Marshal(bean)

	err = strings.set(conn, key, jsonBytes, args...)
	if err != nil {
		return err
	}

	return nil
}

