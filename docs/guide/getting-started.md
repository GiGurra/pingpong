# Getting Started

## Installation

### Using Go

```bash
go install github.com/gigurra/pingpong@latest
```

### From Source

```bash
git clone https://github.com/gigurra/pingpong.git
cd pingpong
go build
```

### Verify Installation

```bash
pingpong --help
```

## Your First Test

### Step 1: Find Your Public IP

The host needs to share their public IP with the tester:

```bash
curl ifconfig.me
```

This returns something like `203.0.113.42`.

### Step 2: Host Starts Listening

On the host machine:

```bash
pingpong listen --conn-type tcp --port 12345
```

You'll see:
```
Listening on port 12345 via TCP
```

### Step 3: Tester Pings

On the tester machine:

```bash
pingpong ping --conn-type tcp --addr 203.0.113.42 --port 12345
```

### Step 4: Check Results

**Success:**
```
Successfully connected to 203.0.113.42:12345 via TCP
INFO Received response response=pong
```

**Failure:**
```
Ping failed: failed to connect to 203.0.113.42:12345 via TCP: dial tcp 203.0.113.42:12345: i/o timeout
```

## Commands

### listen

Start a server that responds to pings.

```bash
pingpong listen --conn-type <tcp|udp> --port <port>
```

| Flag | Description |
|------|-------------|
| `--conn-type` | Protocol: `tcp` or `udp` |
| `--port` | Port number to listen on |

### ping

Test if a server is reachable.

```bash
pingpong ping --conn-type <tcp|udp> --addr <ip> --port <port>
```

| Flag | Description |
|------|-------------|
| `--conn-type` | Protocol: `tcp` or `udp` |
| `--addr` | Host's public IP address |
| `--port` | Port number to test |

## TCP vs UDP

Choose the protocol that matches your game:

| Protocol | Characteristics | Common Games |
|----------|----------------|--------------|
| **TCP** | Reliable, connection-oriented | Minecraft, Terraria, Factorio |
| **UDP** | Fast, connectionless | Valheim, CS2, Rust |

!!! tip
    If you're not sure, try TCP first - it's easier to debug.
