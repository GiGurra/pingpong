# PingPong

A simple CLI tool to test port forwarding configuration when playing games together online.

## What does it do?

PingPong helps you verify that port forwarding is set up correctly before starting a multiplayer gaming session. One player runs the `listen` command to act as a server, while another player uses the `ping` command to test if they can reach the server through the forwarded port.

## Installation

```bash
go install github.com/gigurra/pingpong@latest
```

Or build from source:

```bash
git clone https://github.com/gigurra/pingpong.git
cd pingpong
go build
```

## Usage

### Server Mode (Host)

The player hosting the game runs the listener to accept incoming connections:

```bash
# Listen on TCP port 25565
pingpong listen --conn-type tcp --port 25565

# Listen on UDP port 27015
pingpong listen --conn-type udp --port 27015
```

Make sure to:
1. Configure port forwarding on your router for the specified port
2. Allow the port through your firewall
3. Share your public IP address with the other player

### Client Mode (Tester)

The other player tests if they can reach the host:

```bash
# Test TCP connection
pingpong ping --conn-type tcp --addr <host-ip> --port 25565

# Test UDP connection
pingpong ping --conn-type udp --addr <host-ip> --port 27015
```

If successful, you'll see a "Successfully connected" message and receive a "pong" response.

## Supported Protocols

- **TCP**: Reliable connection-oriented protocol (most common for game servers)
- **UDP**: Fast connectionless protocol (used by many real-time games)

## Common Use Cases

### Minecraft Server
```bash
# Host
pingpong listen --conn-type tcp --port 25565

# Tester
pingpong ping --conn-type tcp --addr 203.0.113.42 --port 25565
```

### Valheim Server
```bash
# Host
pingpong listen --conn-type udp --port 2456

# Tester
pingpong ping --conn-type udp --addr 203.0.113.42 --port 2456
```

## Troubleshooting

**Connection fails?**
- Verify port forwarding is configured on the router
- Check firewall settings on both host and tester machines
- Ensure you're using the host's public IP (not local network IP)
- Confirm the correct protocol (TCP vs UDP) is being used
- Some ISPs block certain ports or use CGNAT which prevents port forwarding

**How to find your public IP?**
```bash
curl ifconfig.me
```

## License

MIT

## Contributing

Pull requests are welcome! Feel free to open an issue if you encounter any problems.
