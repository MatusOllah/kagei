package main

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func initImages(m map[int]string) ([]*ebiten.Image, error) {
	var imgs []*ebiten.Image
	for _, path := range m {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return nil, err
		}

		if opts.ResizeImages {
			img = resizeImage(img)
		}

		imgs = append(imgs, img)
	}

	return imgs, nil
}

func resizeImage(img *ebiten.Image) *ebiten.Image {
	newImg := ebiten.NewImage(opts.Width, opts.Height)
	if opts.Verbose {
		slog.Info("resizing image", "old", img.Bounds(), "new", newImg.Bounds())
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(img.Bounds().Dx())/2, -float64(img.Bounds().Dy())/2)
	op.GeoM.Translate(float64(opts.Width)/2, float64(opts.Height)/2)
	newImg.DrawImage(img, op)

	return newImg
}
