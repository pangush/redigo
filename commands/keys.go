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

import "github.com/pangush/redigo/redis"

// Redis DEL 删除指定的一批keys，如果删除中的某些key不存在，则直接忽略。
func Del(conn redis.Conn, args... interface{}) error {
	_, err := conn.Do("DEL",  args...)
	if err != nil {
		return err
	}

	return nil
}
