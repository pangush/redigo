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
