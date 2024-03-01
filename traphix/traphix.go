// traphix/traphix.go
package traphix

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

// ANSI color code templates
const (
	ansiColorTemplate = "\033[48;2;%d;%d;%dm  \033[0m"
	reset             = "\033[0m"
)

// Pixel represents a single pixel with RGB color.
type Pixel struct {
	R, G, B int
}

// Window represents a colored window as a 2D array of pixels.
type Window struct {
	Width, Height int
	Pixels        [][]Pixel
}

// NewWindow creates a new Window with the specified width and height.
func NewWindow(width, height int) *Window {
	window := &Window{
		Width:  width,
		Height: height,
		Pixels: make([][]Pixel, height),
	}

	for i := range window.Pixels {
		window.Pixels[i] = make([]Pixel, width)
	}

	return window
}

// SetAllPixelsColor sets the color of all pixels in the window.
func (w *Window) SetAllPixelsColor(r, g, b int) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.SetPixel(x, y, r, g, b)
		}
	}
}

// SetPixel sets the color of a specific pixel at the specified position.
func (w *Window) SetPixel(x, y, r, g, b int) {
	w.Pixels[y][x] = Pixel{R: r, G: g, B: b}
}

// ClearScreen clears the terminal screen.
func (w *Window) ClearScreen() {
	cmd := exec.Command("clear") // Use "clear" for Unix-like systems, "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// PrintWindow prints the window with ANSI color codes.
func (w *Window) PrintWindow() {
	for _, row := range w.Pixels {
		for _, pixel := range row {
			colorCode := fmt.Sprintf(ansiColorTemplate, pixel.R, pixel.G, pixel.B)
			fmt.Print(colorCode) // Print only the color code, no spacing
		}
		fmt.Println()
	}
}

// Background sets the background color of the window.
func (w *Window) Background(r, g, b int) {
	w.SetAllPixelsColor(r, g, b)
}

// Window creates a new Window with the specified size and sets the background color.
func Screen(width, height, r, g, b int) *Window {
	window := NewWindow(width, height)
	window.Background(r, g, b)
	return window
}

// DrawLine draws a line from (x0, y0) to (x1, y1) using Bresenham's algorithm.
func (w *Window) DrawLine(x0, y0, x1, y1, r, g, b int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	var sx, sy int

	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	for {
		w.SetPixel(x0, y0, r, g, b)

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

// DrawCircle draws a circle with center (cx, cy) and radius r using trigonometry.
func (w *Window) DrawCircle(cx, cy, r, rr, rg, rb int) {
	for theta := 0.0; theta < 360.0; theta += 0.1 {
		x := int(math.Round(float64(cx) + float64(r)*math.Cos(theta*math.Pi/180.0)))
		y := int(math.Round(float64(cy) + float64(r)*math.Sin(theta*math.Pi/180.0)))

		// Draw only if the pixel is inside the window
		if x >= 0 && x < w.Width && y >= 0 && y < w.Height {
			w.SetPixel(x, y, rr, rg, rb)
		}
	}
}

// Helper function to get the absolute value.
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
