package tiles

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var masterTexture *ebiten.Image

type Water struct {
	x float64
	y float64
}

func init() {

	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/water.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	masterTexture, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	masterTexture.DrawImage(origEbitenImage, op)
}

func NewWater(xPos float64, yPos float64) *Water {
	return &Water{x: xPos, y: yPos}
}

func (s *Water) Draw(screen *ebiten.Image, xOffset float64, yOffset float64) {
	w, h := masterTexture.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(s.x, s.y)
	op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(masterTexture, op)
}

func (s *Water) GetPos() (float64, float64) {
	return s.x, s.y
}
