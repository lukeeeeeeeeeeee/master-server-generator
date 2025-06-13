# msgen — Master Server Generator
#### Author: Bocaletto Luca

**msgen** is a cross-platform CLI tool written in Go that scaffolds a standalone, persistent “server” binary. The generated server:

- Runs as a background service/daemon on Windows & Linux  
- Auto-starts at boot (Windows Service / systemd unit)  
- Automatically restarts on crash or stop (service recovery policy)  
- Periodically collects machine info (IP, MAC, hostname, OS, CPU, memory, disk, users)  
- Sends templated email reports via SMTP (STARTTLS/SMTPS)  
- Is fully self-contained: configuration is embedded or loaded from `config.yaml`  

---

## 🚀 Features

- Single executable (`server.exe` on Windows, `server` on Linux)  
- Embedded or external YAML configuration  
- Cron-style scheduling (`@every 30m`, `0 8 * * *`, etc.)  
- System metrics via [gopsutil]  
- Email delivery via [gomail] with TLS & retry  
- Windows Service integration (register/unregister)  
- systemd unit file support for Linux  
- Structured console logging (zerolog)  
- Dockerfile for containerized deployment  
- GitHub Actions CI workflow  

---

## 📋 Prerequisites

- Go 1.21+ toolchain  
- Git  
- On Windows: Administrative privileges to register a Service  
- On Linux: `systemd` environment and `sudo` for unit deployment  

---

## 🔧 Installation

1. Clone the repository and install `msgen`:
    ```bash
    git clone https://github.com/bocaletto-luca/msgen.git
    cd msgen
    go install ./cmd/msgen
    ```
2. Verify `msgen` is in your `$GOPATH/bin` or `$GOBIN`:
    ```bash
    msgen --help
    ```

---

## ⚙️ Configuration

Create or edit a YAML config file (`config.yaml`). You can either keep it alongside your server binary or let it fall back to the embedded default.

```yaml
# config.yaml
schedule: "@every 30m"

smtp:
  host: "smtp.example.com"
  port: 587
  username: "alerts@example.com"
  password: "supersecret"
  from: "alerts@example.com"
  to:
    - "admin1@example.com"
    - "admin2@example.com"

modules:
  - ip
  - mac
  - os
  - cpu
  - mem
  - disk
  - users
```

**Environment variable overrides**  
You may replace any field with `${ENV_VAR}` syntax and export the corresponding environment variables before generating or running the server.

---

## ▶️ Generating Your Server

Run `msgen` pointing at your config:

```bash
msgen --config path/to/config.yaml --out ./dist
```

This will create:

- `dist/config.yaml` (copied)  
- `dist/server.go` (templated Go source)  

---

## 🛠️ Building Executables

### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o dist/server.exe dist/server.go
```

### Linux

```bash
GOOS=linux   GOARCH=amd64 go build -o dist/server   dist/server.go
chmod +x dist/server
```

---

## ⚙️ Deployment

### Windows Service

1. Open PowerShell as Administrator.
2. Register the service:
   ```powershell
   sc.exe create ms-server `
     binPath= "C:\path\to\server.exe --config C:\path\to\config.yaml" `
     start= auto
   sc.exe failure ms-server reset= 0 actions= restart/5000
   ```
3. Start the service:
   ```powershell
   sc.exe start ms-server
   ```
4. To stop & remove:
   ```powershell
   sc.exe stop ms-server
   sc.exe delete ms-server
   ```

### Linux systemd Unit

1. Copy binary and config to `/opt/ms-server/`:
   ```bash
   sudo mkdir -p /opt/ms-server
   sudo cp dist/server /opt/ms-server/
   sudo cp dist/config.yaml /opt/ms-server/
   ```
2. Create `/etc/systemd/system/ms-server.service`:
   ```ini
   [Unit]
   Description=msgen Server Daemon
   After=network.target

   [Service]
   ExecStart=/opt/ms-server/server --config /opt/ms-server/config.yaml
   Restart=always
   RestartSec=5
   User=root
   LimitNOFILE=4096

   [Install]
   WantedBy=multi-user.target
   ```
3. Enable & start:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable ms-server
   sudo systemctl start ms-server
   ```
4. Check status & logs:
   ```bash
   sudo systemctl status ms-server
   journalctl -u ms-server -f
   ```

---

## 🐳 Docker

Build and run in a container:

```bash
docker build -t bocaletto-luca/ms-server .
docker run -d \
  -v $(pwd)/config.yaml:/app/config.yaml:ro \
  -e SMTP_PASSWORD \
  --name ms-server \
  bocaletto-luca/ms-server
```

---

## 🔍 Health Checks

The server binary exposes a basic HTTP health endpoint at `http://localhost:8000/healthz`. You can integrate this with load balancers or monitoring systems.

---

## 🧪 Testing & CI

- Unit tests and linters:
  ```bash
  go test ./...
  go vet ./...
  ```
- CI workflow on GitHub Actions builds & tests for multiple Go versions.

---

## 📝 License

This project is licensed under the **GPL License**. See [LICENSE](LICENSE) for details.

---

**Author:** Bocaletto Luca ([@bocaletto-luca](https://github.com/bocaletto-luca))  
