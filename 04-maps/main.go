package main

import "fmt"

func main() {
	colors := map[string]string{
		"green": "#123cad",
		"red":   "#123fff",
		"white": "#ffffff",
	}

	colors["red"] = "#ff0000"

	delete(colors, "green")

	printMapStrStr(colors)

	fmt.Println(colors)
}

func printMapStrStr(c map[string]string) {
	for key, value := range c {
		fmt.Printf("%v - %v\n", key, value)
	}
}
