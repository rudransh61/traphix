// main.go
package main

import (
	"traphix/traphix"
	"time"
	"math"
)

func main() {
	// Create a window with size 40x20 and set the background color to black
	window := traphix.Screen(40, 20, 0, 0, 0)

	// Set the initial position and velocity of the ball
	ballX, ballY := window.Width/2, window.Height/2
	velocityX, velocityY := 1, 1

	// Loop to create the bouncing ball animation
	for {
		// Clear the screen for each frame
		window.ClearScreen()
		window.Background(0,0,0)

		// Update the ball's position based on velocity
		ballX += velocityX
		ballY += velocityY

		// Bounce off the walls
		if ballX <= 3 || ballX >= window.Width-5 {
			velocityX = -velocityX
		}
		if ballY <= 3 || ballY >= window.Height-5 {
			velocityY = -velocityY
		}

		// Draw a colorful ball that smoothly changes colors
		rr := int(128 + 127*math.Sin(float64(time.Now().UnixNano()%1000000000)/1000000000.0))
		rg := int(128 + 127*math.Sin(float64(time.Now().UnixNano()%1000000000)/1000000000.0+2))
		rb := int(128 + 127*math.Sin(float64(time.Now().UnixNano()%1000000000)/1000000000.0+4))
		window.DrawCircle(ballX, ballY, 3, rr, rg, rb)
		window.DrawLine(0,0,0,19,0,0,255)
		window.DrawLine(39,0,39,19,0,0,255)
		window.DrawLine(0,0,39,0,0,0,255)
		window.DrawLine(0,19,39,19,0,0,255)
		// Print the window with the drawn ball
		window.PrintWindow()

		// Pause for a short duration to control the animation speed
		time.Sleep(50 * time.Millisecond)
	}
}
