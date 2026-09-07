# Go Consensus Engine

A lightweight, modular Proof-of-Work (PoW) blockchain consensus engine built entirely from scratch in Go. 

This project demonstrates core protocol engineering principles, including cryptographic immutability, state validation, and peer-to-peer consensus simulation without relying on external frameworks.

**Core Features**
* **Cryptographic Hashing:** Secures blocks using `crypto/sha256` for deterministic hash generation.
* **Proof-of-Work Mining:** Implements a target-difficulty mining algorithm requiring computational effort (nonce iteration) to mint new blocks.
* **Chain Validation:** Includes integrity checks to instantly detect data tampering or invalid chain links.
* **Modular Architecture:** Separates consensus logic into a dedicated `core` package for clean, scalable code organization.

**Tech Stack**
* Language: Go (Golang)
* Dependencies: Standard Library Only (`crypto`, `encoding`, `strings`)

**Getting Started**

1. Clone the repository:
git clone https://github.com/shashini-karunarathna/go-consensus-engine.git

2. Navigate to the project directory:
cd go-consensus-engine

3. Run the blockchain simulation:
go run main.go

**Future Roadmap**
* Implement a structured P2P network using `libp2p`.
* Migrate state persistence to `LevelDB` or `BadgerDB`.
* Add elliptic curve cryptography (ECDSA) for transaction signing.
