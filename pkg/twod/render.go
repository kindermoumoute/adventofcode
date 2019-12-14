package twod

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
)

var (
	RenderingEnabled = false
	started          = false
	picUpdate        = make(chan pixel.Picture)
)

type RenderingMap map[interface{}]color.RGBA

func (m Map) Render(renderingMap RenderingMap) {
	if !RenderingEnabled {
		return
	}
	if !started {
		go mainLoop()
		started = true
	}
	picUpdate <- m.loadPicture(renderingMap)
}

func mainLoop() {
	maxX, maxY := pixelgl.PrimaryMonitor().Size()
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, maxX, maxY),
		//Monitor:   pixelgl.PrimaryMonitor(),
		//Undecorated: true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	//win.SetSmooth(true)

	win.Clear(colornames.White)

	fps := time.Tick(time.Second / 144)
	camPos := pixel.ZV
	camSpeed := 500.0
	camZoom := 10.0
	camZoomSpeed := 1.2
	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		frames++
		select {
		case pic := <-picUpdate:
			win.Clear(colornames.White)
			sprite := pixel.NewSprite(pic, pic.Bounds())

			mat := pixel.IM
			mat = mat.Scaled(pixel.ZV, camZoom)
			mat = mat.Moved(win.Bounds().Center().Sub(camPos))
			sprite.Draw(win, mat)
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
		win.Update()
		<-fps
	}
}

func (m Map) rectangle() image.Rectangle {
	botL := m.FindBottomLeft()
	topR := m.FindTopRight()

	upLeft := image.Point{botL.X(), botL.Y()}
	lowRight := image.Point{topR.X() + 1, topR.Y() + 1}
	return image.Rectangle{upLeft, lowRight}
}

func (m Map) loadPicture(renderingMap RenderingMap) pixel.Picture {
	img := image.NewRGBA(m.rectangle())
	for k, v := range m {
		color, exist := renderingMap[v]
		if exist {
			img.Set(k.X(), k.Y(), color)
		}
	}
	return pixel.PictureDataFromImage(img)
}
