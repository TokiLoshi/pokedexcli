package ascii

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

var renderer *figlet4go.AsciiRender

var colorMap = map[string]figlet4go.Color{
	"red": figlet4go.ColorRed, 
	"blue": figlet4go.ColorBlue,
	"yellow": figlet4go.ColorYellow,
	"green": figlet4go.ColorGreen, 
	"magenta": figlet4go.ColorMagenta,
	"cyan": figlet4go.ColorCyan,
}

func init() {
	renderer = figlet4go.NewAsciiRender()
}

func RenderText(text string) (string, error) {
	renderedStr, err := renderer.Render(text)
	if err != nil {
		return "", fmt.Errorf("error rendering ascii: %w", err)
	}
	return renderedStr, nil
}

func RenderTextOptions(text, col1, col2 string) (string, error) {
	options := figlet4go.NewRenderOptions()
	
	colour1, ok1 := colorMap[col1]
	if !ok1 {
		return "", fmt.Errorf("invalid color: %s", colour1)
	}
	
	colour2, ok2 := colorMap[col2]
	if !ok2 {
		return "", fmt.Errorf("invalid color: %s", colour2)
	}

	
	options.FontColor = []figlet4go.Color{
		colour1, 
		colour2,
	}

	renderStr, err := renderer.RenderOpts(text, options)
	if err != nil {
		return "", fmt.Errorf("error with options: %w", err)
	}
	return renderStr, nil
}
	// options := figlet4go.NewRenderOptions()
  //   options.FontColor = []figlet4go.Color{
  //       figlet4go.ColorRed,
  //       figlet4go.ColorMagenta,
	// 	}
	
	// renderStr, _ := ascii.RenderOpts("Pokedex", options)