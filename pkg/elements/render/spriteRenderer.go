package render

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

//SpriteRenderer is the component that handles all
//rendering of sprites onto the screen
type SpriteRenderer struct {
	container *elements.Element
	Tex       *ebiten.Image
	Filename  string

	Type          string
	Width, Height float64
}

var textureMapMaster = make(map[string]*ebiten.Image)

//NewSpriteRenderer creates a SpriteRenderer which
//is the component that handles all rendering of
//sprites onto the screen
func NewSpriteRenderer(container *elements.Element, filename string) *SpriteRenderer {

	var tex *ebiten.Image
	var width, height int

	if textureMapMaster[filename] == nil {
		tex = textureFromPNG(filename)
		width, height = tex.Size()
		textureMapMaster[filename] = tex
	} else {
		tex = textureMapMaster[filename]
		width, height = tex.Size()
	}

	return &SpriteRenderer{
		container: container,
		Tex:       tex,
		Filename:  filename,
		Width:     float64(width),
		Height:    float64(height),
		Type:      "SpriteRenderer",
	}
}

//OnDraw Draws the stored texture file onto the screen
func (sr *SpriteRenderer) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(sr.Width)/2, -float64(sr.Height)/2)
	op.GeoM.Rotate(sr.container.Rotation)
	op.GeoM.Translate(float64(sr.Width)/2, float64(sr.Height)/2)
	op.GeoM.Translate(sr.container.XPos, sr.container.YPos)

	op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(sr.Tex, op)
	return nil
}

//OnUpdate is used to qualify SpriteRenderer as a component
func (sr *SpriteRenderer) OnUpdate(world []*elements.Element) error {
	return nil
}

func textureFromPNG(filename string) *ebiten.Image {
	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/"+filename, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	masterTexture, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	masterTexture.DrawImage(origEbitenImage, op)
	return masterTexture
}

func (sr *SpriteRenderer) OnCheck(elemC *elements.Element) error {
	return nil
}

func (sr *SpriteRenderer) OnUpdateServer(world []*elements.Element) error {
	return nil
}
