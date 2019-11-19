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

package commands

import (
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewPool() *redis.Pool {
	pool := &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			//if _, err := c.Do("AUTH", ""); err != nil {
			//  c.Close()
			//  return nil, err
			//}
			if _, err := c.Do("SELECT", "0"); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	return pool
}

func TestDel(t *testing.T) {
	pool := NewPool()
	conn := pool.Get()
	defer conn.Close()

	key := "TestStrings"
	key1 := "TestStrings"

	err := Del(conn, key, key1)

	assert.NoError(t, err)
}

