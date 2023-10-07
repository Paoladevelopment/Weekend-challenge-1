package mario

func PrintMarioPyramid(height int) string {

	pyramid := ""
	for i := 0; i <= height; i++ {
		for j := 0; j < height; j++ {
			if j >= height-i {
				pyramid += "#"
			} else {
				pyramid += " "
			}
		}
		pyramid += "\n"
	}

	return pyramid
}