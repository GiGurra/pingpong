# pingpong

[![Docs](https://img.shields.io/badge/docs-GitHub%20Pages-blue)](https://gigurra.github.io/pingpong/)

Simple CLI tool to test port forwarding for multiplayer gaming.

## The Problem

You want to host a game server, but you're not sure if your port forwarding is working. Your friend tries to connect and it fails - but is it the port forwarding, the firewall, or something else?

## The Solution

Test the port before starting the game:

```bash
# You (the host)
pingpong listen --conn-type tcp --port 25565

# Your friend
pingpong ping --conn-type tcp --addr YOUR_PUBLIC_IP --port 25565
```

If it works, you'll see "pong" - your port forwarding is good!

## Installation

```bash
go install github.com/gigurra/pingpong@latest
```

## Quick Start

### Host (Server)

```bash
# TCP (Minecraft, most games)
pingpong listen --conn-type tcp --port 25565

# UDP (Valheim, real-time games)
pingpong listen --conn-type udp --port 2456
```

### Tester (Client)

```bash
# TCP
pingpong ping --conn-type tcp --addr 203.0.113.42 --port 25565

# UDP
pingpong ping --conn-type udp --addr 203.0.113.42 --port 2456
```

## Common Game Ports

| Game | Port | Protocol |
|------|------|----------|
| Minecraft | 25565 | TCP |
| Valheim | 2456-2458 | UDP |
| Terraria | 7777 | TCP |
| CS2 | 27015 | UDP |

## Documentation

See the [full documentation](https://gigurra.github.io/pingpong/) for troubleshooting and setup guides.

## License

MIT
