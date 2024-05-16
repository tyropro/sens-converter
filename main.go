package main

import (
	"flag"
	"fmt"
	"math"
)

var baseSens float64
var baseDpi float64
var newDpi float64
var gameFrom string
var gameTo string
var eDpi bool

var gameConvert bool
var gameFromConstant float64
var gameToConstant float64

var oldEdpi float64
var newEdpi float64
var cm360 float64
var newSens float64

func init() {
	flag.Float64Var(&baseSens, "s", 1, "The sensitivity to convert")
	flag.Float64Var(&baseDpi, "b", 800, "Original DPI to convert from")
	flag.Float64Var(&newDpi, "n", 800, "New DPI to convert to")
	flag.StringVar(&gameFrom, "f", "undefined", "Game to convert from")
	flag.StringVar(&gameTo, "t", "undefined", "Game to convert to")
	flag.BoolVar(&eDpi, "e", false, "Prints the eDPI instead of the sensitivity")
	flag.Parse()
}

// THESE VALUES ARE CONSTANT AND ARE NOT TO BE CHANGED
// use switch statements as go does not support struct indexing
func getConstants() {
	switch gameFrom {
		case "aimlabs":
			gameFromConstant = 18288.0
		case "apex":
			gameFromConstant = 41563.65
		case "ark":
			gameFromConstant = 138545.42
		case "cs":
			gameFromConstant = 41563.65
		case "destiny":
			gameFromConstant = 138545.42
		case "finals":
			gameFromConstant = 1159212.01
		case "fn":
			gameFromConstant = 164608.46
		case "minecraft":
			gameFromConstant = 91089.19
		case "ow":
			gameFromConstant = 138545.42
		case "roblox":
			gameFromConstant = 2298.09
		case "val":
			gameFromConstant = 13062.86
	}

	switch gameTo {
		case "aimlabs":
			gameToConstant = 18288.0
		case "apex":
			gameToConstant = 41563.65
		case "ark":
			gameToConstant = 138545.42
		case "cs":
			gameToConstant = 41563.65
		case "destiny":
			gameToConstant = 138545.42
		case "finals":
			gameToConstant = 1159212.01
		case "fn":
			gameToConstant = 164608.46
		case "minecraft":
			gameToConstant = 91089.19
		case "ow":
			gameToConstant = 138545.42
		case "roblox":
			gameToConstant = 2298.09
		case "val":
			gameToConstant = 13062.86
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	gameConvert = gameFrom != "undefined" && gameTo != "undefined"

	// get constant values if no game is specified
	if !gameConvert {
		gameFrom = "cs"
		gameTo = "cs"
	}

	getConstants()

	oldEdpi = baseSens * baseDpi

	if !eDpi {
		cm360 = gameFromConstant / oldEdpi; // get cm360 of sens
		newEdpi = gameToConstant / cm360; // get edpi of new game
		newSens = newEdpi / newDpi; // get sens of new game
	} else {
		newSens = oldEdpi; // print eDPI if no conversion
	}

	fmt.Println(roundFloat(newSens, 3))
}
