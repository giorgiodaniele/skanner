
# skanner

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue)
![License](https://img.shields.io/badge/License-MIT-green)

Skanner is a simple TCP port scanner written in Go.

## Usage

Build with `go build`, then run `skanner -h` for usage information.

```bash
skanner -targets 127.0.0.1,192.168.1.1 -ports 22,80,443,1-1024 -timeout 500ms -workers 100
```