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

# System Design Diagram

### Sequential diagram for the Avalanche Consensus Protocol implementation

 The diagram shows how transactions are processed, how nodes query each other for preferences, and how consensus is reached

```mermaid 
---
title: how Avalanch transactions are processed
---
sequenceDiagram
    participant NodeA as Node A
    participant NodeB as Node B
    participant NodeC as Node C
    participant NodeD as Node D
    participant Network as Network

    NodeA->>Network: Create RandomTransaction()
    Network-->>NodeA: Return Transaction

    NodeA->>NodeA: HandleMessage(Transaction)
    NodeA->>NodeB: QueryNode(Transaction)
    NodeB-->>NodeA: Respond with Preference

    NodeA->>NodeC: QueryNode(Transaction)
    NodeC-->>NodeA: Respond with Preference

    NodeA->>NodeD: QueryNode(Transaction)
    NodeD-->>NodeA: Respond with Preference

    alt Majority Preference is True
        NodeA->>NodeA: IncrementVote()
    else No Majority Preference
        NodeA->>NodeA: DecrementVote()
    end

    alt Confidence >= Threshold
        NodeA->>NodeA: Accept Transaction
    else Confidence <= -Threshold
        NodeA->>NodeA: Reject Transaction
    end

    Note over NodeA: Continue querying nodes until consensus is reached

```

Explanation
Transaction Creation:

Node A creates a random transaction by calling RandomTransaction().
The network returns the transaction to Node A.
Handling the Transaction:

Node A processes the transaction by calling HandleMessage().
Querying Peers:

Node A queries Node B for its preference regarding the transaction.
Node B responds with its preference.
Node A queries Node C for its preference.
Node C responds with its preference.
Node A queries Node D for its preference.
Node D responds with its preference.
Voting Based on Preferences:

If the majority of nodes agree with the preference, Node A increments the vote count.
If there is no majority agreement, Node A decrements the vote count.
Consensus Decision:

If the confidence level reaches the threshold, Node A accepts the transaction.
If the confidence level drops below the negative threshold, Node A rejects the transaction.
Continued Querying:

Node A continues to query other nodes until consensus is reached, either by accepting or rejecting the transaction.

### Class diagram for the Avalanche Consensus Protocol implementation

```mermaid
classDiagram
    class Transaction {
        +uint64 Nonce
        +int32 Data
        +RandomTransaction() *Transaction
        +Serialize() []byte
        +Hash() string
        +ConflictWith(other *Transaction) bool
    }

    class Node {
        +int64 ID
        +map~string, TxState~ Mempool
        +NewNode(id int64) *Node
        +HandleMessage(origin int64, msg Message)
        +QueryNode(peer *Node, txHash string) bool
    }

    class TxState {
        +Transaction Tx
        +int Confidence
        +bool Accepted
        +bool Rejected
        +int ConfidenceThreshold
        +int Alpha
        +int Beta
        +int SnowflakeCounter
        +int SnowballCounter
        +bool Preference
        +NewTxState(tx *Transaction) *TxState
        +IncrementVote()
        +DecrementVote()
    }

    class Network {
        +map~int64, Node~ nodes
        +NewNetwork(n int64) *Network
        +Run()
        +simulateNodeActivity(node *Node)
        +queryPeers(node *Node, txHash string) bool
        +getRandomPeers(nodeID int64, count int) []*Node
    }

    class Message {
        <<interface>>
    }

    class MessageTransaction {
        +Transaction Tx
    }

    Node "1" *-- "many" TxState : manages
    Network "1" *-- "many" Node : contains
    MessageTransaction --|> Message : implements
    Node "1" *-- "many" Message : processes
    TxState "1" *-- "1" Transaction : includes


```
