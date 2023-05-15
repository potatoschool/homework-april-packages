package game

type Settings struct {
	screenWidth  int
	screenHeight int
}

func (s *Settings) SetWidth(width int) {
	s.screenWidth = width
}

func (s *Settings) SetHeight(heigth int) {
	s.screenHeight = heigth
}
