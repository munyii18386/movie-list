package handlers

import (
	"errors"
	"encoding/json"
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

//RedisInstance represents a session.Store backed by redis.
type RedisInstance struct {
	Client          *redis.Client
	SessionDuration time.Duration
}

//NewRedisInstance constructs a new RedisStore
func NewRedisInstance(client *redis.Client, sessionDuration time.Duration) *RedisInstance {
	return &RedisInstance{
		Client:          client,
		SessionDuration: sessionDuration,
	}
}

//Save saves the provided `sessionState` and associated SessionID
func (rs *RedisInstance) Save(sid string, sessionState interface{}) error {
	key := "sid:" + sid
	state, err := json.Marshal(sessionState)
	if nil != err {
		return err
	}

	err = rs.Client.Set(key, state, rs.SessionDuration).Err()
	if err != nil {
		return err
	}
	return nil

}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisInstance) Get(sid string, sessionState interface{}) error {
	key := "sid:" + sid
	list := rs.Client.Pipeline()
	result := list.Get(key)
	if result.Err() != nil{
		return errors.New("no session state was found in the session store")
	}
	changeExpiration := list.Expire(key, rs.SessionDuration).Err()
	if changeExpiration != nil {
		return fmt.Errorf("error changing expiration of session <%s>:\n%s", sid, changeExpiration.Error())
	}
	_, err:= list.Exec()
	if err != nil {
		return fmt.Errorf("error getting sid <%s>:\n%v", string(sid), err.Error())
	}
	err = json.Unmarshal([]byte(result.Val()), sessionState)
	if err != nil {
		return fmt.Errorf("error unmarshaling sessionState: %s", err.Error())
	}
	return nil

}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisInstance) Delete(sid string) error {
	key := "sid:" + sid
	rs.Client.Del(key)
	return nil
}

// Find retrieves value by key
func (rs *RedisInstance) Find(sid string) {
	
	val := rs.Client.Do("GET", sid)

	if val == nil {
		fmt.Println("i exist")
	}

}