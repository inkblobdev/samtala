# Samtala

Samtala is a decentralized, distributed chat application that creates peer-to-peer chatrooms with robust encryption for both in-transit and at-rest data.

## Notice
Samtala is a toy project for now.

## Objectives

- [ ] create a base network enabled zig executable(samtala core) that can transmit tcp messages.
- [ ] create a rendezvous server in go to facilitate discovery of the different Samtala clients.
- [ ] enable samtala core to do tcp hole punching to connect to the peers.
- [ ] implement distributed persistance on samtala core.

Innitialy I'll focus on get this zig core functionall afterwards a gui can be implemented.

## Features

- Decentralized, peer-to-peer chatrooms
- End-to-end encryption using TLS and PGP
- Centralized server for discovery and P2P connection establishment
- Direct P2P transfer for all messages and files
- Separate core executable and GUI components

## Architecture

Samtala consists of two main components:

1. **Core Executable**: Runs in the background and handles:
   - Encryption
   - Data replication
   - Storage
   - Consensus logic

2. **GUI**: User interface that communicates with the core via named pipes

Each chatroom functions as its own P2P cluster, ensuring privacy and scalability.

## How It Works

1. Users connect to a central server for initial discovery and P2P connection establishment.
2. Once connected, all communication (messages, files, etc.) occurs directly between peers.
3. The core executable manages all backend operations, while the GUI provides a user-friendly interface.
