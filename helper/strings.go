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
	"github.com/pangush/redigo/redis"
)

type Strings struct {
	engine 			*Engine
}

// Strings return the wrapper of redis strings
func NewStrings(conn redis.Conn, key string) *Strings  {
	engine := NewEngine(conn, key)
	strings := &Strings{engine:engine}

	return strings
}

func (strings *Strings) set(conn redis.Conn, key string, value interface{}, args ...interface{}) error {
	if key == "" {
		return ErrKeyEmpty
	}

	newArgs := make([]interface{}, 0)
	newArgs = append(newArgs, key)
	newArgs = append(newArgs, value)
	newArgs = append(newArgs, args...)

	_, err := conn.Do("SET", newArgs...)
	if err != nil {
		return err
	}

	return nil
}

func (strings *Strings) get(conn redis.Conn, key string) (interface{}, error) {
	if key == "" {
		return "", ErrKeyEmpty
	}

	return conn.Do("GET",  key)
}

func (strings *Strings) Exists() (bool, error) {
	return strings.engine.Exists()
}

func (strings *Strings) Del() (bool, error) {
	return strings.engine.Del()
}