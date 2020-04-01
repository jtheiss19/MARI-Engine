package player

type PlrView struct {
	xPos   float64
	yPos   float64
	speed  float64
	width  float64
	height float64
}

func NewPlrView() *PlrView {
	return &PlrView{xPos: 0, yPos: 0, speed: 1, width: 1280, height: 720}
}

func (s *PlrView) MovePlrViewBy(newX float64, newY float64) {
	s.xPos += newX * s.speed
	s.yPos += newY * s.speed
}

func (s *PlrView) CanView(xPosTest float64, yPosTest float64) bool {
	if s.xPos <= xPosTest+64 && xPosTest <= s.xPos+s.width {
		if s.yPos <= yPosTest+64 && yPosTest <= s.yPos+s.height {
			return true
		}
	}
	return false
}

func (s *PlrView) GetPos() (float64, float64) {
	return s.xPos, s.yPos
}
