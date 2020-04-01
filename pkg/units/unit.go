package units

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type IsUnit interface {
	Draw(screen *ebiten.Image, xOffset float64, yOffset float64)
	Update()
	GetPos() (float64, float64)
	GetType() string
}

type Unit struct {
	X     float64
	Y     float64
	V     float64
	Type  string
	Owner string
}

var devTexture *ebiten.Image

func NewUnit(kind string, owner string) *Unit {
	return &Unit{X: 0, Y: 0, V: .1, Type: kind, Owner: owner}
}

func init() {
	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/dev.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	devTexture, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	devTexture.DrawImage(origEbitenImage, op)
}

func (s *Unit) Draw(screen *ebiten.Image, xOffset float64, yOffset float64) {
	w, h := devTexture.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(1 * math.Pi * 270 / 360)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(s.X, s.Y)
	op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(devTexture, op)
}

func (s *Unit) Update() {
	s.X += s.V
	s.Y += s.V
}

func (s *Unit) GetPos() (float64, float64) {
	return s.X, s.Y
}

func (s *Unit) GetType() string {
	return s.Type
}
