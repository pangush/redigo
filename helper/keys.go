package helper

import "github.com/pangush/redigo/redis"

type Keys struct {
	engine 	*Engine
	//key 	string
}

func NewKeys(conn redis.Conn, key string) *Keys {
	return &Keys{
		engine: NewEngine(conn, key),
	}
}

// Redis DEL 删除指定的一批keys，如果删除中的某些key不存在，则直接忽略。
func (k *Keys) Del() (bool, error)  {
	conn := k.engine.conn
	key := k.engine.key

	if key == "" {
		return false, ErrKeyEmpty
	}

	bo, err := redis.Bool(conn.Do("DEL",  key))

	return bo, err
}

func (k *Keys) Type() {

}

// 以秒为单位，返回给定 key 的剩余生存时间
func (k *Keys) TTl() {

}

// Redis Expire 命令用于设置 key 的过期时间。key 过期后将不再可用。
// EXPIREAT 的作用和 EXPIRE类似，都用于为 key 设置生存时间。
// 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳 Unix timestamp 。
// 1 如果设置了过期时间 0 如果没有设置过期时间，或者不能设置过期时间
func (k *Keys) Expire() {

}

// Redis PTLL 命令以毫秒为单位返回 key 的剩余过期时间
func (k *Keys) PTTL() {

}

// Redis EXISTS 命令用于检查给定 key 是否存在。
func (k *Keys) Exists() (bool, error) {
	conn := k.engine.conn
	key := k.engine.key

	if key == "" {
		return false, ErrKeyEmpty
	}

	bo, err := redis.Bool(conn.Do("EXISTS",  key))

	return bo, err
}

// Redis ExpireAt 命令用于以 UNIX 时间戳(unix timestamp)格式设置 key 的过期时间。key 过期后将不再可用。