package kata

import "fmt"

func RGB(r, g, b int) (hex string) {
	for _, num := range []int{r, g, b} {
		if num < 0 {
			num = 0
		} else if num > 255 {
			num = 255
		}
		hex += fmt.Sprintf("%02X", num)
	}
	return
}
