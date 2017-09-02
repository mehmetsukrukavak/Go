package main

import "fmt"
type mapsType map[string]string
func main() {
	////var colors map[string] string
	//
	//colors := make(map[int]string)
	//colors[10]="#ffffff"
	//
	//delete(colors,10)

	colors := mapsType{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	colors.printMaps()
}


func printMap(c mapsType){
	for color,hex :=range c{
		fmt.Println("Hex code for", color,"is",hex)
	}
}

func (c mapsType) printMaps(){
	for color,hex :=range c{
		fmt.Println("Hex code for", color,"is",hex)
	}
}
