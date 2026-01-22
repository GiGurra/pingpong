# Game Examples

Common games and their default ports.

## Minecraft (Java Edition)

**Port:** 25565 (TCP)

```bash
# Host
pingpong listen --conn-type tcp --port 25565

# Tester
pingpong ping --conn-type tcp --addr <host-ip> --port 25565
```

## Minecraft (Bedrock Edition)

**Port:** 19132 (UDP)

```bash
# Host
pingpong listen --conn-type udp --port 19132

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 19132
```

## Valheim

**Ports:** 2456-2458 (UDP)

```bash
# Host (test main port)
pingpong listen --conn-type udp --port 2456

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 2456
```

!!! note
    Valheim uses three consecutive ports. Forward all three: 2456, 2457, 2458.

## Terraria

**Port:** 7777 (TCP)

```bash
# Host
pingpong listen --conn-type tcp --port 7777

# Tester
pingpong ping --conn-type tcp --addr <host-ip> --port 7777
```

## Factorio

**Port:** 34197 (UDP)

```bash
# Host
pingpong listen --conn-type udp --port 34197

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 34197
```

## Counter-Strike 2 / CS:GO

**Port:** 27015 (UDP)

```bash
# Host
pingpong listen --conn-type udp --port 27015

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 27015
```

## Rust

**Port:** 28015 (UDP)

```bash
# Host
pingpong listen --conn-type udp --port 28015

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 28015
```

## ARK: Survival Evolved

**Ports:** 7777 (UDP), 27015 (UDP)

```bash
# Host (test game port)
pingpong listen --conn-type udp --port 7777

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 7777
```

## 7 Days to Die

**Ports:** 26900 (TCP), 26900-26902 (UDP)

```bash
# Host (TCP)
pingpong listen --conn-type tcp --port 26900

# Host (UDP)
pingpong listen --conn-type udp --port 26900

# Tester
pingpong ping --conn-type tcp --addr <host-ip> --port 26900
pingpong ping --conn-type udp --addr <host-ip> --port 26900
```

## Palworld

**Port:** 8211 (UDP)

```bash
# Host
pingpong listen --conn-type udp --port 8211

# Tester
pingpong ping --conn-type udp --addr <host-ip> --port 8211
```

## Quick Reference

| Game | Port(s) | Protocol |
|------|---------|----------|
| Minecraft Java | 25565 | TCP |
| Minecraft Bedrock | 19132 | UDP |
| Valheim | 2456-2458 | UDP |
| Terraria | 7777 | TCP |
| Factorio | 34197 | UDP |
| CS2 | 27015 | UDP |
| Rust | 28015 | UDP |
| ARK | 7777, 27015 | UDP |
| 7 Days to Die | 26900 | TCP+UDP |
| Palworld | 8211 | UDP |
