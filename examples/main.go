package main

import (
	"fmt"
	"jwt"
	"time"
)

func main()  {
	jwtUtil := jwt.New(&jwt.Conf{
		Secret:   []byte("e2hxjuiwrdhgkj1"),
		ExpireAt: 1,
	})
	name := "zhangsan"
	token, expireAt, err := jwtUtil.GenerateToken(name)
	if err != nil {
		panic(err)
	}
	_ = expireAt
	time.Sleep(2 * time.Second)
	claims, err := jwtUtil.ParseToken(token)
	if err != nil {
		panic(err)
	}
	if claims.UserName != name{
		panic("Don't match")
	}
	fmt.Println("pass")
}
