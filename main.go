package main

import (
	"fmt"
	"image"
	img_color "image/color"
	"image/png"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessevdk/go-flags"
)

var Version string = "1.2.0"

type Game struct {
	shader   *ebiten.Shader
	images   []*ebiten.Image
	uniforms map[string]interface{}
}

func NewGame() (*Game, error) {
	if opts.Verbose {
		slog.Info("compiling shader")
	}
	b, err := os.ReadFile(opts.PosArgs.ShaderPath)
	if err != nil {
		return nil, err
	}

	s, err := ebiten.NewShader(b)
	if err != nil {
		fmt.Println(opts.PosArgs.ShaderPath + ":" + err.Error())
		os.Exit(1)
	}
	if opts.Verbose {
		slog.Info("done compiling shader")
	}

	if opts.Verbose {
		slog.Info("loading source images")
	}
	images, err := initImages(opts.Image)
	if err != nil {
		return nil, err
	}
	if opts.Verbose {
		slog.Info("done loading source images")
	}

	uniforms := make(map[string]interface{})
	for k, v := range opts.UniformBool {
		uniforms[k] = v
	}
	for k, v := range opts.UniformInt {
		uniforms[k] = v
	}
	for k, v := range opts.UniformFloat {
		uniforms[k] = v
	}

	return &Game{
		shader:   s,
		images:   images,
		uniforms: uniforms,
	}, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(img_color.Black)

	op := &ebiten.DrawRectShaderOptions{}
	copy(op.Images[:], g.images)
	op.Uniforms = g.uniforms
	screen.DrawRectShader(screen.Bounds().Dx(), screen.Bounds().Dy(), g.shader, op)

	if opts.ExportImage != "" {
		exportImage(screen, opts.ExportImage)
		os.Exit(0)
	}

	if opts.FPSCounter {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return opts.Width, opts.Height
}

func exportImage(img image.Image, path string) {
	if opts.Verbose {
		slog.Info("exporting image", "path", path)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}

func main() {
	if _, err := flags.NewParser(&opts, flags.HelpFlag|flags.IgnoreUnknown|flags.PassDoubleDash).Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if opts.Version {
		cyan := color.New(color.Bold, color.FgCyan).SprintFunc()
		white := color.New(color.Bold).SprintFunc()
		fmt.Println(cyan("kagei"), white("version"), "v"+Version)
		fmt.Println(cyan("Go"), white("version"), runtime.Version())
		os.Exit(0)
	}

	if opts.Verbose {
		slog.Info(fmt.Sprintf("kagei version %s", Version))
		slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	}

	if opts.Verbose {
		slog.Info("opts", "opts", fmt.Sprintf("%+v", opts))
	}

	ebiten.SetVsyncEnabled(false)
	if opts.VSync {
		ebiten.SetVsyncEnabled(true)
	}

	if opts.Verbose {
		slog.Info("initializing game")
	}
	beforeInit := time.Now()
	g, err := NewGame()
	if err != nil {
		panic(err)
	}
	if opts.Verbose {
		slog.Info("init game done", "time", time.Since(beforeInit))
	}

	ebiten.SetWindowSize(opts.Width, opts.Height)
	ebiten.SetWindowTitle("kagei")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
