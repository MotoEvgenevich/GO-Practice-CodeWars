package main

import "fmt"

func main() {
	length, width, height := 10.2, 7.2, 8.2
	fmt.Println(GetVolumeOfCuboid(length, width, height))
}

func GetVolumeOfCuboid(length, width, height float64) float64 {
	cubeVolume := length * width * height
	return cubeVolume
}
