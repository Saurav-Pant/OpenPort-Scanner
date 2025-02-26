# Go Port Scanner

A fast and efficient port scanner written in Go that uses concurrent goroutines to scan multiple ports simultaneously.

## Features

- **Concurrent port scanning** using goroutines and worker pools
- **Configurable port range scanning** (default: 1-1024)
- **Adjustable timeout settings**
- **User-friendly output** showing open ports and scan duration
- **Efficient worker pool implementation** to manage system resources

## Requirements

- Go 1.11 or higher

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/go-port-scanner.git
cd go-port-scanner
```

## Usage

1. Build and run the application:

   ```bash
   go run main.go
   ```

2. Enter the target IP address or domain name when prompted:

   ```bash
   Enter target IP or domain: example.com
   ```

3. The scanner will begin checking ports and display results in real-time:

   ```bash
   Scanning ports...
   [+] Port 80 is open
   [+] Port 443 is open
   Open ports: [80 443]

   Scan completed in 1.234s
   ```

## Configuration

You can modify the following constants in `main.go` to adjust the scanner's behavior:

- `WORKER_POOL_SIZE`: Number of concurrent workers (default: 100)
- `startPort`: Beginning of port range to scan (default: 1)
- `endPort`: End of port range to scan (default: 1024)
- `timeout`: Connection timeout duration (default: 500ms)

## Performance

This port scanner implements concurrent scanning using Go's goroutines and channels, making it significantly faster than sequential scanning. The **worker pool pattern** is used to limit the number of concurrent connections and prevent system resource exhaustion.

