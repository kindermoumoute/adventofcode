package twod

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"sync"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	RenderingEnabled = false
	started          = false
	picUpdate        = make(chan pixel.Picture)
	RenderingMap     = make(map[interface{}]color.Color)
	fps              = (<-chan time.Time)(nil)

	lock = sync.Mutex{}
)

func SetFPS(rate int) {
	lock.Lock()
	if rate == 0 {
		fps = nil
	} else {
		fps = time.Tick(time.Second / time.Duration(rate))
	}
	lock.Unlock()
}

func (m Map) Render() {
	if !RenderingEnabled {
		return
	}
	if !started {
		go mainLoop()
		started = true
	}
	picUpdate <- m.loadPicture()
}

func mainLoop() {
	maxX, maxY := pixelgl.PrimaryMonitor().Size()
	cfg := pixelgl.WindowConfig{
		Title:  "Advent of code!",
		Bounds: pixel.R(0, 0, maxX, maxY),
		//Monitor:   pixelgl.PrimaryMonitor(),
		//Undecorated: true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)

	camPos := pixel.ZV
	camSpeed := 500.0
	camZoom := 10.0
	camZoomSpeed := 1.2
	frames := 0
	second := time.Tick(time.Second)
	lastTime := time.Now()
	lastPic := pixel.Picture(nil)

	for !win.Closed() {
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()
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
		if win.MouseScroll().Y != 0 {
			drawPic(win, lastPic, camZoom, camPos)
		}

		frames++
		select {
		case pic := <-picUpdate:
			drawPic(win, pic, camZoom, camPos)
			lastPic = pic
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
		win.Update()
		lock.Lock()
		if fps != nil {
			<-fps
		}
		lock.Unlock()
	}
}

func drawPic(win *pixelgl.Window, pic pixel.Picture, camZoom float64, camPos pixel.Vec) {
	if pic == nil {
		return
	}
	win.Clear(colornames.White)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, camZoom)
	mat = mat.Moved(win.Bounds().Center().Sub(camPos))
	sprite.Draw(win, mat)
}

func (m Map) rectangle() image.Rectangle {
	botL := m.FindBottomLeft()
	topR := m.FindTopRight()

	upLeft := image.Point{botL.X(), botL.Y()}
	lowRight := image.Point{topR.X() + 1, topR.Y() + 1}
	return image.Rectangle{upLeft, lowRight}
}

func (m Map) loadPicture() pixel.Picture {
	img := image.NewRGBA(m.rectangle())
	for k, v := range m {
		matchingColor, exist := RenderingMap[v]
		if exist {
			img.Set(k.X(), k.Y(), matchingColor)
		}
	}
	return pixel.PictureDataFromImage(img)
}
