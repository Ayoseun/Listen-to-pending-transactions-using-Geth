# Listen-to-pending-transactions-using-Geth
This Tutorial demonstrates how to view transactions in the blockchain mempool using golang, Geth

- ### Prerequisites
Before we delve into the code, ensure you have the following prerequisites in place:

 1. Golang installed on your system, you can install Golang from here

 2. A basic understanding of Ethereum and its transaction lifecycle

 3. An IDE (Visual studio code)

### Getting Started

Create a file called main.go and then run 
go mod init main

This will create a file called go.mod like below

```go
created go.mod
```

Now create a file called main.go in the same directory as the go.mod file we created earlier.

At this point, we can now install the go-ethereum package.

To do this run the command below

```go
go get github.com/ethereum/go-ethereum/
```

it will now appear in our go.mod file and a file named go.sum will also be automatically created carrying all the dependencies.



_Great , we are now set _ ✅

1. Importing Required Packages

The initial part of the code is concerned with importing the necessary packages that enable interaction with Ethereum networks and WebSockets. Notably, we import the following:

```go
import (
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient/gethclient"
    "github.com/ethereum/go-ethereum/rpc"
)

```

* let's talk about these a bit

 - "context" package is essential for handling the context of the application.

 - "fmt" package allows us to format and print output messages.

 - "log" package provides error logging functionalities.

 - "github.com/ethereum/go-ethereum/common" offers Ethereum common data structures.

 - "github.com/ethereum/go-ethereum/ethclient/gethclient" provides access to the Ethereum client.

 - "github.com/ethereum/go-ethereum/rpc" enables communication with Ethereum nodes via RPC.


2. Connecting to an Ethereum Node

To interact with the Ethereum network, we need a connection to a node. In this case, i am connecting using the WebSocket-based RPC client from Infura, a popular Ethereum node provider. The node URL is specified as:

```go
client, err := rpc.Dial("wss://polygon-mumbai.infura.io/ws/v3/b2ea..................")
if err != nil {
    log.Fatal(err)
}
```

Now it is good to note that you will need to use the Infura wss for this to work properly.

So create an Infura account and then switch to websocket RPCs like in the image below

Infura website image

Do not forget to replace the Infura URL with your preferred Ethereum node's WebSocket endpoint.


3. Subscribing to Pending Transactions

Once we have established a connection to the Ethereum node, we can subscribe to pending transactions using the gethclient.New method:

```go
ctx := context.Background()

pendingTxSub := make(chan common.Hash)
sub, err := gethclient.New(client).SubscribePendingTransactions(ctx, pendingTxSub)
if err != nil {
    log.Fatal(err)
}
defer sub.Unsubscribe()
```

We create a channel pendingTxSub to receive incoming pending transaction hashes. The SubscribePendingTransactions method helps us subscribe to these pending transactions in real-time. Don't forget to unsubscribe from the subscription once you are done with it, hence the defer sub.Unsubscribe().


4. Processing Incoming Pending Transactions

Furthermore, we set up a loop to process incoming pending transactions:

```go
for {
    select {
    case err := <-sub.Err():
        log.Fatal(err)
    case txs := <-pendingTxSub:
        fmt.Printf("Pending Transaction Hash: %s\n", txs.Hex())
    }
}
```

Here,
we use a select statement to listen for two types of channels: sub.Err() to handle errors and pendingTxSub to process incoming pending transaction hashes. When a new pending transaction is detected, the hash is printed to the console.

Complete code
```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient/gethclient"
    "github.com/ethereum/go-ethereum/rpc"
)

func main() {
    // Connect to an Ethereum node using an HTTP-based RPC client
    client, err := rpc.Dial("<RPC HERE>")
    if err != nil {
        log.Fatal(err)
    }

    // Create a context to handle the subscription
    ctx := context.Background()

    // Subscribe to pending transactions
     pendingTxSub := make(chan common.Hash)
    sub, err := gethclient.New(client).SubscribePendingTransactions(ctx,pendingTxSub)
    if err != nil {
        log.Fatal(err)
    }
    defer sub.Unsubscribe()

    // Start processing incoming pending transactions
    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case txs := <-pendingTxSub:
            fmt.Printf("Pending Transaction Hash: %s\n", txs.Hex())

        }
    }
}
```

Congratulations! 🚀

You have now successfully built a simple Golang application that connects to an Ethereum node, subscribes to pending transactions, and processes them in real-time. This forms the foundation for creating more sophisticated transaction monitoring systems, dApp analytics tools, and auditing solutions.

Feel free to expand on this code and integrate additional features, such as storing transaction data in a database or analyzing the gas fees associated with transactions.



