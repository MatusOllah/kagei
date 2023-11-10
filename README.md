# kagei
 
**kagei** (Kage Interpreter) is a CLI tool for testing Kage shaders.

## Building & Installing

1. Install Go

2. Run: `go install github.com/MatusOllah/kagei@latest`

## Usage

To lookup a user run:

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
