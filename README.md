<img src="https://static1.pixcee.dev/external/strawstacks/logo.png" width="128px"></img>

# Strawhouse
Bare minimal file store engine, featured lightweight indexing, signed checksum file integrity, and pre-signed token with access control.

## Architecture
- Store file using host filesystem-backed without overhead.
- Supports signed checksums and flags (file-level validation for compromised direct filesystem modifications) using [xattr](https://en.wikipedia.org/wiki/Extended_file_attributes) stored directly in [inode](https://en.wikipedia.org/wiki/Inode) for lightning fast access and no database overhead.
- Reengineered stateless pre-signed token for clients to upload and retrieve file with expiration time, path restriction, and custom attribute support. Used only 27-bytes token with bit-by-bit optimization (see [performance benchmark](https://github.com/strawstacks/strawhouse/wiki/Benchmark)).
- Get file directly from HTTP GET (with just path and token), upload using pre-defined API endpoints.
- Store file metadata in high performance key-value embed database: [pogreb](https://github.com/akrylysov/pogreb).
- Designed and optimized for store and serve millions of static files from the ground up.
- Structured to support file validation, file event hooks in the future.

## Module

Get into each modules for more documentation:

| [Backend](#backend)                       | [Driver](#driver)                        | [Command](#command)                         |
|-------------------------------------------|------------------------------------------|---------------------------------------------|
| File server with HTTP and gRPC interface. | Go library for interacting with backend. | Command line tool to test and manage files. |

### Backend

File server with HTTP and gRPC interface.

1. **Configuration**
   
   Create a configuration file `config.yaml` with the following content:
   ```yaml
   webListen: ["tcp", ":3000"]
   protoListen: ["tcp", ":3001"]
   dataRoot: ./local/data/
   pogrebPath: ./local/pogreb/
   key: 6AnxPZy.... # base64-encoded 48 bytes key: `openssl rand -base64 48`
   ```
   Note:
   - Listen address referred by [net.Listen](https://golang.org/pkg/net/#Listen). Example: `["tcp", ":3000"]`, `["unix", "/tmp/strawhouse.sock"]`.
   - `dataRoot` is the root directory of static files.
   - `pogrebPath` is the directory for metadata database.

2. Choice 1: **Try with `go run`**
   ```bash
   git clone https://github.com/strawstacks/strawhouse.git
   cd strawhouse/
   go run ./backend --config ./local/config.yaml
   ```

3. Choice 2: **Compile binary from source**
   ```bash
   go build -o ./local/backend ./backend
   mv ./local/backend /usr/local/bin/strawhousebackd
   strawhousebackd --config /etc/strawhouse/backend/config.yaml
   ```
   
4. Choice 3: **Download pre-built binary**
   ```bash
   sudo wget -O /usr/local/bin/strawhousebackd https://github.com/strawstacks/strawhouse/releases/download/v0.1.0/strawhousebackd_linux_arm64
   sudo chmod +x /usr/local/bin/strawhousebackd
   strawhousebackd --config /etc/strawhouse/backend/config.yaml
   ```
   
5. **Using service manager**
   
   Create a service file `/etc/systemd/system/strawhousebackd.service`:
   ```ini
   [Unit]
   Description=Strawhouse Backend Service
   After=network.target

   [Service]
   Type=simple
   ExecStart=/usr/local/bin/strawhousebackd --config /etc/strawhouse/backend/config.yaml
   Restart=on-failure

   [Install]
   WantedBy=multi-user.target
   ```
   Then run:
   ```bash
   systemctl enable strawhousebackd
   systemctl start strawhousebackd
   ```

Why not Docker?
  - Docker is not recommended for high-performance file server due to the overhead of filesystem mapping and network stack.

### Driver

Go library for interacting with backend.

```bash
go get -u github.com/strawstacks/strawhouse/driver
```

```go
func main() {
   st := strawhouse.New("6AnxPZy....", "localhost:3001") // key, gRPC address
   defer st.Close()
   
   mode := strawhouse.SignatureModeFile // or strawhouse.SignatureModeDirectory
   action := strawhouse.SignatureActionUpload // or strawhouse.SignatureActionGet
   depth := 2 // Only effect in get action: for path /a/b/c, depth of 2 means allow access all files under /a/b, for upload action, it's ignored and allow user to upload to /a/b/c only.
   expired := time.Now().Add(time.Duration(20) * time.Second) // 20 seconds
   path := "/a/b/c" // Relative path to dataRoot that grant user access
   token := st.Signature.Generate(1, mode, action, uint32(depth), expired, path, nil)
   fmt.Println(token)
}
```
### Command

Command line tool to test and manage files.

```bash
go install github.com/strawstacks/strawhouse/command
```

First time configuration:
```bash
strawc config --name key # Backend's key
strawc config --name server # Backend's gRPC address
```

Sign token:
```bash
strawc --action upload --depth 1 --expired 60 --mode dir --path photos/
```

## Credits

Developed by **[BSthun](https://github.com/BSthun)**, originated to use internally at [Connected Tech](https://www.connectedtech.co.th) and [Pixcee](https://www.pixcee.app/).

The name of **Straw House** inspired by pronunciation of "Store House" (in the meaning of file store), plus [Facebook's Haystack](https://engineering.fb.com/2009/04/30/core-infra/needle-in-a-haystack-efficient-storage-of-billions-of-photos/) and a song [Jade - Straw House](https://open.spotify.com/track/50uwQoov3D7ASWwfmRVHQI?si=9081a42990ba4233).