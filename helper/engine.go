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

// Engine is the major struct of redisgo, it means a redis manager
// Commonly, an application only need one engine
type Engine struct {
	key 	string
	conn 	redis.Conn
	//keys	*Keys
}

func NewEngine(conn redis.Conn, key string) *Engine {
	return &Engine{
		key:  	key,
		conn: 	conn,
		//keys:	&Keys{
		//	key:key,
		//},
	}
}
