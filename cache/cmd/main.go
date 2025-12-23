package main

import (
	"fmt"

	"github.com/pfremaux/golibs/cache/pkg/cache"
)

func main() {
	c := cache.NewSimpleCache[string]()
	c.Set("myKey", "myValue")
	fmt.Println(c.Get("myKey")) // Output: myValue
}
