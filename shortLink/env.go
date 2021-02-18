package main

type Env struct {
	S Storage
}

func getEnv() *Env {
	r := NewRedisCli("127.0.0.1:6379", "", 1)
	return &Env{S: r}
}
