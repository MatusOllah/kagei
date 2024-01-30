package main

var opts struct {
	PosArgs struct {
		ShaderPath string `description:"Path to Kage shader"`
	} `positional-args:"yes" required:"yes"`

	Version         bool               `long:"version" description:"Print version and exit"`
	Verbose         bool               `short:"v" long:"verbose" description:"Print verbose information"`
	Width           int                `short:"w" long:"width" description:"Window width" default:"640"`
	Height          int                `short:"g" long:"height" description:"Window height" default:"480"`
	FPSCounter      bool               `long:"fps-counter" description:"Show FPS counter"`
	TPSCounter      bool               `long:"tps-counter" description:"Show TPS counter"`
	ShowUniforms    bool               `long:"show-uniforms" description:"Show uniform variables"`
	VSync           bool               `long:"vsync" description:"Enable VSync"`
	Image           map[int]string     `short:"I" long:"image" description:"Source images"`
	UniformBool     map[string]bool    `long:"uniform-bool" description:"Uniform bool variable"`
	UniformInt      map[string]int     `long:"uniform-int" description:"Uniform int variable"`
	UniformFloat    map[string]float32 `long:"uniform-float" description:"Uniform float variable"`
	UniformVec      map[string][]float32
	UniformVecFunc  func(string) `long:"uniform-vec" description:"Uniform vecN variable"`
	UniformIvec     map[string][]int
	UniformIvecFunc func(string) `long:"uniform-ivec" description:"Uniform ivecN variable"`
	UniformJSON     string       `long:"uniform-json" description:"Uniforms using JSON"`
	UniformJSONFile string       `long:"uniform-json-file" description:"Uniforms using JSON file"`
	ResizeImages    bool         `long:"resize-images" description:"Resize source images to window size"`
	ExportImage     string       `long:"export-image" description:"Export rendered image and exit"`
}
