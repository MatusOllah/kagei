# kagei

[![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/kagei)](https://goreportcard.com/report/github.com/MatusOllah/kagei)
 
**kagei** (Kage Interpreter) is a CLI tool for testing [Kage shaders](https://ebitengine.org/en/documents/shader.html).

## Building & Installing

1. Install Go

2. Install a C compiler (only on macOS, Linux and FreeBSD)

3. Run: `go install github.com/MatusOllah/kagei@latest`

## Usage

To test a Kage shader run:

```shell
kagei path/to/shader.kage
```

### Useful flags

- `-v` - prints verbose output
- `--version` - print version
- `-w` - screen width
- `-g` - screen height
- `--fps-counter` - FPS counter
- `--vsync` - VSync
- `-I` - source image (for example `-I0:image.png`)
- `-U` - uniform variables (for example `-Unum:69`)
- `--resize-images` - resize source images
