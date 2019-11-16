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

// Set redis strings set value
func (strings *Strings) SetString(value string, args ...interface{}) error {
	conn := strings.engine.conn
	key := strings.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetBytes(value []byte, args ...interface{}) error {
	conn := strings.engine.conn
	key := strings.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetUint64(value uint64, args ...interface{}) error {
	conn := strings.engine.conn
	key := strings.engine.key

	err := strings.set(conn, key, value, args...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) SetStruct(bean interface{}, args ...interface{}) error {
	conn := strings.engine.conn
	key := strings.engine.key

	jsonBytes, err := json.Marshal(bean)

	err = strings.set(conn, key, jsonBytes, args...)
	if err != nil {
		return err
	}

	return nil
}

