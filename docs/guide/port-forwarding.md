# Port Forwarding Setup

Port forwarding tells your router to send incoming traffic on a specific port to your computer.

## Prerequisites

Before configuring port forwarding:

1. **Know your local IP** - Your computer's address on your home network
2. **Know the port** - The port your game uses
3. **Know the protocol** - TCP, UDP, or both

### Find Your Local IP

**Windows:**
```
ipconfig
```
Look for "IPv4 Address" (usually `192.168.x.x`)

**macOS/Linux:**
```bash
ip addr | grep inet
# or
ifconfig | grep inet
```

## Router Configuration

### General Steps

1. Open your router's admin page (usually `192.168.1.1` or `192.168.0.1`)
2. Log in (check your router for default credentials)
3. Find "Port Forwarding" or "NAT" settings
4. Add a new rule:
   - **Service Port / External Port**: The game port (e.g., 25565)
   - **Internal IP**: Your computer's local IP
   - **Internal Port**: Same as service port
   - **Protocol**: TCP, UDP, or Both

### Common Router Interfaces

**Most routers:**
- Advanced Settings → NAT → Port Forwarding

**Asus:**
- WAN → Virtual Server / Port Forwarding

**Netgear:**
- Advanced → Advanced Setup → Port Forwarding

**TP-Link:**
- Advanced → NAT Forwarding → Port Forwarding

## Firewall Configuration

Your computer's firewall also needs to allow the traffic.

### Windows Firewall

```powershell
# Allow TCP
netsh advfirewall firewall add rule name="Game Server TCP" dir=in action=allow protocol=tcp localport=25565

# Allow UDP
netsh advfirewall firewall add rule name="Game Server UDP" dir=in action=allow protocol=udp localport=25565
```

### macOS

System Settings → Network → Firewall → Options → Add application

Or temporarily disable:
```bash
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --setglobalstate off
```

### Linux (iptables)

```bash
# Allow TCP
sudo iptables -A INPUT -p tcp --dport 25565 -j ACCEPT

# Allow UDP
sudo iptables -A INPUT -p udp --dport 25565 -j ACCEPT
```

### Linux (ufw)

```bash
sudo ufw allow 25565/tcp
sudo ufw allow 25565/udp
```

## Testing Your Setup

After configuring:

1. **Start pingpong listener:**
   ```bash
   pingpong listen --conn-type tcp --port 25565
   ```

2. **Have a friend test:**
   ```bash
   pingpong ping --conn-type tcp --addr YOUR_PUBLIC_IP --port 25565
   ```

3. **If it works**, your port forwarding is correctly configured!

## Common Issues

### Double NAT

If you have two routers (ISP router + your router), you need to:
- Configure port forwarding on both, OR
- Put one router in bridge mode

### CGNAT

Some ISPs use Carrier-Grade NAT. Port forwarding won't work because you share a public IP with other customers.

Check if you have CGNAT:
1. Find your router's WAN IP
2. Compare to your public IP (`curl ifconfig.me`)
3. If they're different, you might be behind CGNAT

Solutions:
- Ask your ISP for a public IP
- Use a VPN with port forwarding (e.g., Mullvad)
- Use a cloud server as a relay

### Dynamic IP

Your public IP might change. Solutions:
- Use a Dynamic DNS service (No-IP, DuckDNS)
- Check your IP before each gaming session
