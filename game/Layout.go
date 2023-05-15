package game

func (view *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return view.settings.screenWidth, view.settings.screenHeight
}
