# NoContextTODO

`NoContextTODO` find wrong usage of context.TODO

```go

package main

import (
    "context"
    "fmt"
    "time"
)

func infiniteLoop(ctx context.Context) {
    innerCtx, cancel := context.WithCancel(ctx)
    defer cancel()
    for {
        fmt.Println("Waiting for time out")
        //time.Sleep(time.Second)
        // If I put sleep in here, fmt.Println doesn't output the "Exit now!"

        select {
        case <-innerCtx.Done():
            fmt.Println("Exit now!")
            fmt.Println("message:", ctx.Value("message").(string))
            return
        default:

        }
    }
}

func main() {
    ctx := context.TODO()
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    ctx = context.WithValue(ctx, "message", "hi")
    defer cancel()

    go infiniteLoop(ctx)

    select {
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}

```

```sh
$ go vet -vettool=`which NoContextTODO` main.go
./main.go:29:9: don't use context.Background. Use context.TODO
```
