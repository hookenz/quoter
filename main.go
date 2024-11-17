package main

import (
	"fmt"
	"math/rand"
	"time"
)

// List of sample quotes
var quotes = []string{
	"“The only way to do great work is to love what you do.” - Steve Jobs",
	"“In the middle of difficulty lies opportunity.” - Albert Einstein",
	"“Your time is limited, don't waste it living someone else's life.” - Steve Jobs",
	"“You must be the change you wish to see in the world.” - Mahatma Gandhi",
	"“To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.” - Ralph Waldo Emerson",
	"“The best way to predict the future is to create it.” - Peter Drucker",
	"“It always seems impossible until it’s done.” - Nelson Mandela",
	"“Success is not final, failure is not fatal: It is the courage to continue that counts.” - Winston Churchill",
}

// ANSI color codes for foreground colors
var colors = []string{
	"\033[38;5;1m",  // Red
	"\033[38;5;2m",  // Green
	"\033[38;5;3m",  // Yellow
	"\033[38;5;4m",  // Blue
	"\033[38;5;5m",  // Magenta
	"\033[38;5;6m",  // Cyan
	"\033[38;5;7m",  // White
	"\033[38;5;8m",  // Bright Black
	"\033[38;5;9m",  // Bright Red
	"\033[38;5;10m", // Bright Green
}

// ANSI codes for background colors
var bgColors = []string{
	"\033[48;5;1m",  // Red Background
	"\033[48;5;2m",  // Green Background
	"\033[48;5;3m",  // Yellow Background
	"\033[48;5;4m",  // Blue Background
	"\033[48;5;5m",  // Magenta Background
	"\033[48;5;6m",  // Cyan Background
	"\033[48;5;7m",  // White Background
	"\033[48;5;8m",  // Bright Black Background
	"\033[48;5;9m",  // Bright Red Background
	"\033[48;5;10m", // Bright Green Background
}

// Text formatting codes
var bold = "\033[1m"
var underline = "\033[4m"
var reset = "\033[0m"

// Randomly select a quote and color, then print
func printRandomQuote() {
	// Select a random quote
	quote := quotes[rand.Intn(len(quotes))]

	// Select a random foreground color
	color := colors[rand.Intn(len(colors))]

	// 10% chance of changing the background color
	bgColor := ""
	if rand.Float64() <= 0.1 {
		bgColor = bgColors[rand.Intn(len(bgColors))]
	}

	// Randomly add bold or underline
	style := ""
	if rand.Intn(2) == 0 {
		style += bold
	}
	if rand.Intn(2) == 0 {
		style += underline
	}

	// Combine the styles and the quote
	styledQuote := color + bgColor + style + quote + reset

	// Print the styled quote
	fmt.Println(styledQuote)
}

func main() {
	limit := 2000

	// simulate some bursty behavior
	for {
		// fill log with 3k lines
		for i := 0; i < limit; i++ {
			printRandomQuote()
			time.Sleep(5 * time.Millisecond)
		}

		// wait 5 minutes
		time.Sleep(5 * time.Minute)

		// now one quote every 2s
		for i := 0; i < 200; i++ {
			printRandomQuote()
			time.Sleep(1 * time.Second)
		}

		limit = rand.Intn(5000)
	}
}
