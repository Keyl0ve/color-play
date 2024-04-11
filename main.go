package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Black("Black")
	color.Red("Red")
	color.Blue("Blue")
	color.Green("Green")
	color.Yellow("Yellow")
	color.White("White")
	color.Magenta("Magenta")
	color.Cyan("Cyan")
	color.HiBlack("HiBlack")
	color.HiRed("HiRed")
	color.HiBlue("HiBlue")
	color.HiWhite("HiWhite")
	color.HiYellow("HiYellow")
	color.HiCyan("HiCyan")
	color.HiGreen("HiGreen")
	color.HiMagenta("HiMagenta")

	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	// Create a custom print function for convenience
	red1 := color.New(color.FgRed).PrintfFunc()
	red1("Warning")
	red1("Error: %s", "Something went wrong")

	// Mix up multiple attributes
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("Don't forget this...")

}
