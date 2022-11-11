# Cache lib
## Usage

```go
package main

import (
	"fmt"
	"github.com/iooojik-dev/go-cache/cache"
	"time"
)

type Any struct {
	Value int
}

func main() {
	c := cache.NewCacheStorage[Any](1 * time.Second)
	c.Set("1", Any{Value: 1}, time.Now().Add(5*time.Second))
	fmt.Println(c.Get("1"))
}
```