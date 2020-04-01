package units

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var masterTexture *ebiten.Image

type Destroyer struct {
	x float64
	y float64
	v float64
}

func init() {

	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/dest1.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	masterTexture, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	masterTexture.DrawImage(origEbitenImage, op)
}

func NewDestroyer() *Destroyer {
	return &Destroyer{x: 0, y: 0, v: .1}
}

func (s *Destroyer) Draw(screen *ebiten.Image, xOffset float64, yOffset float64) {
	w, h := masterTexture.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(1 * math.Pi * 270 / 360)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(masterTexture, op)
}

func (s *Destroyer) Update() {
	s.x += s.v
	s.y += s.v
}

func (s *Destroyer) GetPos() (float64, float64) {
	return s.x, s.y
}
