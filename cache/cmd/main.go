package main

import (
	"fmt"

	"github.com/cache/pkg/cache"
)

func main() {
	c := cache.NewSimpleCache()
	c.Set("myKey", "myValue")
	fmt.Println(c.Get("myKey")) // Output: myValue
}
