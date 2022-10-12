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
	} else if strings.EqualFold(*runType, "Controls") {
		basics.ControlStatements()
	} else if strings.EqualFold(*runType, "Loop") {
		basics.Loop()
	} else if strings.EqualFold(*runType, "Defer") {
		basics.Defer()
	} else if strings.EqualFold(*runType, "Panic") {
		basics.PanicAndRecover()
	} else if strings.EqualFold(*runType, "All") {
		all()
	} else {
		err := errors.New("runtype not found")
		log.Fatal(err)
	}
}
func all() {
	basics.Variables()
	basics.Primitives()
	basics.Arrays()
	basics.Slices()
	basics.Maps()
	basics.Struct()
	basics.ControlStatements()
	basics.Loop()
	basics.Defer()
	basics.PanicAndRecover()
}
