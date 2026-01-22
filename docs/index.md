# pingpong

Simple CLI tool to test port forwarding for multiplayer gaming.

## Why pingpong?

Setting up port forwarding for game servers is frustrating:

1. You configure port forwarding on your router
2. You start the game server
3. Your friend can't connect
4. Is it the port forwarding? The firewall? The game?

**pingpong answers one question: Is the port reachable?**

## How It Works

```
┌─────────────┐                      ┌─────────────┐
│   Host      │                      │   Tester    │
│             │                      │             │
│  pingpong   │◄────── Internet ─────│  pingpong   │
│  listen     │                      │  ping       │
│             │                      │             │
└─────────────┘                      └─────────────┘
     ▲                                     │
     │                                     │
     └──── Router (port forwarding) ◄──────┘
```

1. **Host** runs `pingpong listen` on the game port
2. **Tester** runs `pingpong ping` to the host's public IP
3. If they get "pong" back, port forwarding works!

## Quick Example

**Host (you):**
```bash
pingpong listen --conn-type tcp --port 25565
# Listening on port 25565 via TCP
```

**Tester (your friend):**
```bash
pingpong ping --conn-type tcp --addr 203.0.113.42 --port 25565
# Successfully connected to 203.0.113.42:25565 via TCP
# Received response: pong
```

## Supported Protocols

| Protocol | Use Case |
|----------|----------|
| **TCP** | Most game servers (Minecraft, Terraria) |
| **UDP** | Real-time games (Valheim, CS2) |

## Installation

```bash
go install github.com/gigurra/pingpong@latest
```

## Next Steps

- [Getting Started](guide/getting-started.md) - Installation and first test
- [Port Forwarding Setup](guide/port-forwarding.md) - How to configure your router
- [Game Examples](guide/games.md) - Common games and their ports
- [Troubleshooting](guide/troubleshooting.md) - When things don't work
