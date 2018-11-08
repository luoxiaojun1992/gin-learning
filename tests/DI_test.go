package tests

import (
	"testing"
	"github.com/luoxiaojun1992/gin-learning/DI"
	"log"
)

func Test_Resolve(t *testing.T) {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	userService := DI.Resolve("UserService")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{name string})
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Benchmark_Resolve(b *testing.B)  {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	i := 0
	for ;i <= b.N;i++ {
		if DI.Resolve("UserService").(struct{name string}).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}
