package main

var opts struct {
	PosArgs struct {
		ShaderPath string `description:"Path to Kage shader"`
	} `positional-args:"yes" required:"yes"`
	Version      bool                   `long:"version" description:"Print version and exit"`
	Verbose      bool                   `short:"v" long:"verbose" description:"Print verbose information"`
	Width        int                    `short:"w" long:"width" description:"Window width" default:"640"`
	Height       int                    `short:"g" long:"height" description:"Window height" default:"480"`
	FPSCounter   bool                   `long:"fps-counter" description:"Show FPS counter"`
	VSync        bool                   `long:"vsync" description:"Enable VSync"`
	Image        map[int]string         `short:"I" long:"image" description:"Source images"`
	Uniform      map[string]interface{} `short:"U" long:"uniform" description:"Uniform variables"`
	ResizeImages bool                   `long:"resize-images" description:"Resize source images to window size"`
	ExportImage  string                 `long:"export-image" description:"Export rendered image and exit"`
}
