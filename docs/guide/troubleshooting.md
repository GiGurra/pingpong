# Troubleshooting

## Connection Fails

### "Connection refused"

The port isn't open or nothing is listening.

**Checklist:**
- [ ] Is `pingpong listen` running on the host?
- [ ] Is the firewall allowing the port?
- [ ] Is port forwarding configured on the router?

### "Connection timed out"

The request isn't reaching the host at all.

**Checklist:**
- [ ] Is the public IP correct? (`curl ifconfig.me` on host)
- [ ] Is the port correct?
- [ ] Is port forwarding configured for the right internal IP?
- [ ] Are you behind CGNAT? (See below)

### "No route to host"

Network-level connectivity issue.

**Checklist:**
- [ ] Is the host connected to the internet?
- [ ] Is the public IP valid?
- [ ] Try a different port (some ISPs block common ports)

## UDP-Specific Issues

### UDP "works" but game doesn't connect

UDP is connectionless - pingpong can't fully verify the path.

**Try:**
1. Test with TCP first to verify basic connectivity
2. Check that you're testing all required ports (some games use multiple)
3. Verify game-specific firewall rules

### No response from UDP ping

UDP packets might be dropped silently.

**Checklist:**
- [ ] Verify the host shows "Listening on port X via UDP"
- [ ] Check firewall allows UDP on that port
- [ ] Some routers have UDP timeout issues - try TCP to verify setup

## Finding Your Public IP

**On the host:**
```bash
curl ifconfig.me
```

**Alternative methods:**
```bash
curl api.ipify.org
curl icanhazip.com
curl checkip.amazonaws.com
```

**Don't use your local IP** (192.168.x.x, 10.x.x.x, 172.16-31.x.x)

## CGNAT Detection

Carrier-Grade NAT means you share a public IP with other customers.

**How to check:**

1. Find your router's WAN IP (in router admin page)
2. Find your public IP: `curl ifconfig.me`
3. If they're different, you're probably behind CGNAT

**Signs of CGNAT:**
- WAN IP starts with 100.64.x.x
- WAN IP is in private range but public IP is different
- Your ISP is a mobile carrier or budget provider

**Solutions:**
- Request a public IP from your ISP (may cost extra)
- Use a VPN with port forwarding support
- Use a cloud relay/tunnel service

## Firewall Verification

### Test locally first

On the host machine:
```bash
# Terminal 1
pingpong listen --conn-type tcp --port 25565

# Terminal 2 (same machine)
pingpong ping --conn-type tcp --addr 127.0.0.1 --port 25565
```

If this fails, it's a local firewall issue.

### Temporarily disable firewall

**For testing only:**

Windows:
```powershell
netsh advfirewall set allprofiles state off
# Don't forget to re-enable!
netsh advfirewall set allprofiles state on
```

macOS:
```bash
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --setglobalstate off
# Re-enable after testing
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --setglobalstate on
```

Linux:
```bash
sudo ufw disable
# Re-enable after testing
sudo ufw enable
```

## Port Already in Use

```
failed to listen on port 25565: bind: address already in use
```

Something else is using the port.

**Find what's using it:**

Windows:
```powershell
netstat -ano | findstr :25565
```

macOS/Linux:
```bash
lsof -i :25565
```

**Solutions:**
- Close the other application
- Use a different port for testing

## Still Not Working?

1. **Verify each step independently:**
   - Local test works? (ping 127.0.0.1)
   - LAN test works? (ping from another device on same network)
   - External test works? (ping from friend)

2. **Simplify:**
   - Try a high port number (49152-65535)
   - Try TCP instead of UDP
   - Temporarily disable all firewalls

3. **Check router logs:**
   - Look for blocked connections
   - Verify port forwarding rule is active

4. **Try a mobile hotspot:**
   - Bypasses your router entirely
   - If it works, the problem is your router config
