package main

import (
	"basics/basics"
	goRoutine "basics/basics/goroutines"
	"errors"
	"flag"
	"log"
	"strings"
)

func main() {

	run := flag.String("run", "Variables", "Which basics package you want run")
	routine := flag.String("routine", "Basic", "Which goroutine example you want to run")
	flag.Parse()

	if strings.EqualFold(*run, "Variables") {
		basics.Variables()
	} else if strings.EqualFold(*run, "Primitives") {
		basics.Primitives()
	} else if strings.EqualFold(*run, "Arrays") {
		basics.Arrays()
	} else if strings.EqualFold(*run, "Slices") {
		basics.Slices()
	} else if strings.EqualFold(*run, "Maps") {
		basics.Maps()
	} else if strings.EqualFold(*run, "Struct") {
		basics.Struct()
	} else if strings.EqualFold(*run, "Controls") {
		basics.ControlStatements()
	} else if strings.EqualFold(*run, "Loop") {
		basics.Loop()
	} else if strings.EqualFold(*run, "Defer") {
		basics.Defer()
	} else if strings.EqualFold(*run, "Panic") {
		basics.PanicAndRecover()
	} else if strings.EqualFold(*run, "Pointer") {
		basics.Pointers()
	} else if strings.EqualFold(*run, "Function") {
		basics.Function()
	} else if strings.EqualFold(*run, "Interface") {
		basics.Interface()
	} else if strings.EqualFold(*run, "Gorountine") {
		runGoRoutineExamples(*routine)
	} else if strings.EqualFold(*run, "Go") {
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
	basics.Pointers()
	basics.Function()
	basics.Interface()
}

func runGoRoutineExamples(routine string) {
	if strings.EqualFold(routine, "Basic") {
		goRoutine.GoRoutine()
	} else if strings.EqualFold(routine, "NoSync") {
		goRoutine.NoSync()
	} else if strings.EqualFold(routine, "Mutex") {
		goRoutine.SyncWithMutex()
	} else {
		err := errors.New("routine example not found")
		log.Fatal(err)
	}
}
