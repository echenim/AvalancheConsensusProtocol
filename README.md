# Avalanche Consensus Protocol

The Avalanche Consensus Protocol is a novel consensus mechanism designed to provide high throughput, low latency, and robust security for blockchain and distributed ledger systems. This repository contains the implementation of the Avalanche consensus protocol in Go.

## Overview

Avalanche is a family of consensus protocols that achieve consensus through repeated randomized sampling and network gossiping. It is designed to be scalable, efficient, and secure, making it suitable for a wide range of applications, including cryptocurrencies, DeFi, supply chain management, gaming, IoT, and enterprise applications.

## Key Features

- **High Throughput**: Capable of handling thousands of transactions per second.
- **Low Latency**: Transactions can be confirmed within seconds.
- **Scalability**: Designed to scale with the number of nodes in the network.
- **Robustness**: Resistant to various types of attacks, including Sybil attacks and double-spending.
- **Energy Efficiency**: Does not require extensive computational resources.

## How It Works

1. **Transaction Initiation**: A node creates a transaction and broadcasts it to the network.
2. **Gossip Protocol**: The transaction is propagated through the network using a gossip protocol.
3. **Repeated Randomized Sampling**: Nodes repeatedly query a small, random subset of other nodes to determine their opinion on the validity of the transaction.
4. **Snowball and Snowflake**: Sub-protocols used to aggregate opinions and determine consensus.
5. **Consensus Decision**: Once a node reaches a high enough confidence level in the validity of the transaction, it considers it confirmed and propagates this decision to the network.

## Use Cases

- **Cryptocurrencies and Digital Assets**
- **DeFi (Decentralized Finance)**
- **Supply Chain Management**
- **Gaming**
- **Internet of Things (IoT)**
- **Enterprise Applications**







Below is a diagram illustrating the interaction between the components

```mermaid
classDiagram
    class Transaction {
        +uint64 Nonce
        +int32 Data
        +RandomTransaction() *Transaction
        +Serialize() []byte
        +Hash() string
    }

    class Network {
        +map~int64, Node~ nodes
        +NewNetwork(n int64) *Network
        +Run()
    }

    class Node {
        +sync.Mutex
        +int64 ID
        +map~string, TxState~ Mempool
        +NewNode(id int64) *Node
        +HandleMessage(origin int64, msg Message)
    }

    class Message {
        <<interface>>
    }

    class MessageTransaction {
        +Transaction Tx
    }

    class TxState {
        // Implementation of TxState, similar to Rust version
    }

    Network "1" *-- "many" Node : contains
    Node "1" *-- "many" TxState : manages
    MessageTransaction --|> Message : implements
    Node "1" o-- "many" Message : processes
    Transaction "1" *-- "1" TxState : represents

```

Creating a Random Transaction

```mermaid
sequenceDiagram
    participant User
    participant Transaction
    User->>Transaction: RandomTransaction()
    Transaction-->>User: *Transaction

```

Handling a Message
```mermaid
sequenceDiagram
    participant OriginNode
    participant TargetNode
    participant Message
    OriginNode->>TargetNode: Send Message
    TargetNode->>TargetNode: HandleMessage(origin, msg)

```