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

import (
	"encoding/json"
	"github.com/pangush/redigo/redis"
)

// GetBytes redis strings get []byte
func (strings *Strings) GetBytes() ([]byte, error) {
	conn := strings.engine.conn
	key := strings.engine.key

	if key == "" {
		return nil, ErrKeyEmpty
	}

	reply, err := strings.get(conn, key)
	if err != nil {
		return nil, err
	}

	bytes, err := redis.Bytes(reply, err)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetString redis strings get string
func (strings *Strings) GetString() (string, error) {
	conn := strings.engine.conn
	key := strings.engine.key

	value, err := redis.String(strings.get(conn, key))
	if err != nil {
		return "", err
	}

	return value, nil
}

func (strings *Strings) GetUint64() (uint64, error) {
	conn := strings.engine.conn
	key := strings.engine.key

	value, err := redis.Uint64(strings.get(conn, key))
	if err != nil {
		return uint64(0), err
	}

	return value, nil
}

func (strings *Strings) GetStruct(bean interface{}) error {
	conn := strings.engine.conn
	key :=  strings.engine.key

	value, err := strings.get(conn, key)
	if err != nil {
		return err
	}

	bytes, err := redis.Bytes(value, err)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &bean)
	if err != nil {
		return err
	}

	return nil
}