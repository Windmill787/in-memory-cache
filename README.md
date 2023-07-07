# In memory cache package

You can store `interface{}` types in memory, access them by `string` key and delete them.

## Example

```go
package main

import "github.com/Windmill787/in-memory-cache"

func main() {
    // create cache
    cache := inmemorycache.NewCache()

    // set new value
    cache.Set("userId", 123)

    // get value by key
    cache.Get("userId")

    // delete value by key
    cache.Delete("userId")
}

```
