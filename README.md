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
- `--uniform-bool` - uniform bool variable (for example `--uniform-bool Foo:true`)
- `--uniform-int` - uniform int variable (for example `--uniform-int Foo:1`)
- `--uniform-float` - uniform float variable (for example `--uniform-float Foo:0.5`)
- `--resize-images` - resize source images
- `--export-image` - export rendered image and exit (aka headless mode)
