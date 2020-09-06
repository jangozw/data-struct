package main

import (
	"fmt"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	var expire time.Duration
	expire = 100 * time.Second
	nowTime := time.Now()
	expireTime := nowTime.Add(expire * time.Second)

	fmt.Println(expireTime.Unix())
}
