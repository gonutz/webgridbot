package main

import (
	"image"

	"github.com/gonutz/auto"
)

const (
	screenLeft  = 820.0
	screenTop   = 219.0
	screenSize  = 765.0
	tileCount   = 30.0
	tileSize    = screenSize / tileCount
	targetRed   = 10
	targetGreen = 132
	targetBlue  = 255
)

var (
	lastTileX = -1
	lastTileY = -1
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	running := true
	auto.SetOnKeyboardEvent(func(e *auto.KeyboardEvent) {
		if e.Key == auto.KeyEscape {
			running = false
		}
	})

	for running {
		if err := clickBlue(); err != nil {
			return err
		}
	}
	return nil
}

func clickBlue() error {
	img, err := auto.CaptureScreen(screenLeft, screenTop, screenSize, screenSize)
	if err != nil {
		return err
	}

	rgba := img.(*image.RGBA)
	for tileY := 0; tileY < tileCount; tileY++ {
		for tileX := 0; tileX < tileCount; tileX++ {
			x := round(screenLeft + tileSize/2 + float64(tileX)*tileSize)
			y := round(screenTop + tileSize/2 + float64(tileY)*tileSize)

			i := rgba.PixOffset(x, y)
			r := rgba.Pix[i]
			g := rgba.Pix[i+1]
			b := rgba.Pix[i+2]
			if r == targetRed && g == targetGreen && b == targetBlue {
				// Only click on this tile when it differs from the last clicked
				// tile. This way we do not keep hitting the same tile over and
				// over again.
				if !(tileX == lastTileX && tileY == lastTileY) {
					auto.ClickLeftMouseAt(x, y)
					lastTileX, lastTileY = tileX, tileY
				}
				return nil
			}
		}
	}

	return nil
}

func round(x float64) int {
	return int(x + 0.5)
}
