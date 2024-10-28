package main

func main() {

}

func GetSize(w, h, d int) [2]int {
	area := w * h
	volume := w * h * d
	var arr [2]int = [2]int{area, volume}
	return arr
}
