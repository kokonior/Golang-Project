package code

import (
	"fmt"
	"testing"
	"time"
)

func SayHello() {
	fmt.Println("hello world")
}

func TestSayHello(t *testing.T) {
	go SayHello()
	fmt.Println("OKE")

	// if the block program is come to an end, but the thread still has
	// a gorountine, it will get killed automatically
	time.Sleep(1 * time.Second)
}

func TestEqualMap(t *testing.T) {
	// map1 := map[string]string{}
	map2 := make(map[string]string)

	fmt.Println(map2)
}
