package main

import (
	"basics/basics"
	"errors"
	"flag"
	"log"
	"strings"
)

func main() {

	runType := flag.String("runtype", "Variables", "Which basics package you want run")
	flag.Parse()
	if strings.EqualFold(*runType, "Variables") {
		basics.Variables()
	} else if strings.EqualFold(*runType, "Primitives") {
		basics.Primitives()
	} else if strings.EqualFold(*runType, "Arrays") {
		basics.Arrays()
	} else if strings.EqualFold(*runType, "Slices") {
		basics.Slices()
	} else if strings.EqualFold(*runType, "Maps") {
		basics.Maps()
	} else if strings.EqualFold(*runType, "Struct") {
		basics.Struct()
	} else {
		err := errors.New("runtype not found")
		log.Fatal(err)
	}
}
