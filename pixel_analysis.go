package main

type Pixel struct {
	mean            float32
	variance        float32
	gradient_energy uint16
	bit_entropy     int16
}

type cfaPattern struct {
	r  Pixel
	g1 Pixel
	g2 Pixel
	b  Pixel
}

type Grid struct {
	row     int16
	col     int16
	pattern cfaPattern
}

func ProcessImageFrames() {}

func ExtractPixelStats() {}
