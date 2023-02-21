package main

import "fmt"

const (
	LoopStart       = "нц "
	LoopEnd         = "кц"
	LoopReplacement = "for ______$1 {"
)

func LoopWhile(condition string) string {
	return fmt.Sprintf("нц пока %s", condition)
}

func LoopTimes(iterations int) string {
	return fmt.Sprintf("нц %d раз", iterations)
}
