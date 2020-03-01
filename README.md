# go-nats-cli-ws
Go nats.io client via websockets proxy 

```go
package main

import (
    mqs "github.com/9glt/go-nats-cli-ws"
    stan "github.com/nats-io/stan.go"
)

func main() {
    conn, err := mqs.New("wss://host.domain.tld/path?token=tok", "nats://user:pass@127.0.0.1:4222", "test-cluster", "client-1")
    if err != nil {
        panic(err)
    }
    
    sc := conn.UnderlyingConn()

    // use sc 
    
    // ...
}
```