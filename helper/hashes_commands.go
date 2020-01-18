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

import "github.com/pangush/redigo/redis"

// Set redis strings set value
func (hashes *Hashes) HSet(field string, value string) error {
	conn := hashes.keys.engine.conn
	key := hashes.keys.engine.key

	if key == "" {
		return ErrKeyEmpty
	}

	_, err := conn.Do("HSET",  key, field, value)
	if err != nil {
		return err
	}

	return nil
}

func (hashes *Hashes) HGet(field string) (string, error) {
	conn := hashes.keys.engine.conn
	key := hashes.keys.engine.key

	if key == "" {
		return "", ErrKeyEmpty
	}

	return redis.String(conn.Do("HGet",  key, field))
}

func (hashes *Hashes) HExists(field string) (bool, error) {
	conn := hashes.keys.engine.conn
	key := hashes.keys.engine.key

	if key == "" {
		return false, ErrKeyEmpty
	}

	bo, err := redis.Bool(conn.Do("HEXISTS",  key, field))
	if err != nil {
		return false, err
	}

	return bo, nil
}

func (hashes *Hashes) HDel(field string, args ...interface{}) error {
	conn := hashes.keys.engine.conn
	key := hashes.keys.engine.key

	if key == "" {
		return ErrKeyEmpty
	}

	newArgs := make([]interface{}, 0, len(args) + 2)
	newArgs = append(newArgs, key)
	newArgs = append(newArgs, field)
	newArgs = append(newArgs, args...)

	_, err := conn.Do("HDEL", newArgs...)
	if err != nil {
		return err
	}

	return nil
}

