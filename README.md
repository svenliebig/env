# env

this is a minimal package for go that allows you to read environment variables from a `.env` file, that is located in the root of your working directory.

## Usage

```go
package main

import (
    "fmt"
    "github.com/svenliebig/env"
)

func main() {
    env.Load()
    fmt.Println(env.Get("MY_ENV_VAR"))
}
```
