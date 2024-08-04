package main

// SCP
// Single Responsibility Principle
// A class should have only one reason to change
// A class should have only one job
// In the example below, the Journal struct has two responsibilities:
// 1. To store journal entries
// 2. To save journal entries to a file
// The Journal struct should be split into two separate structs:
// 1. Journal struct to store journal entries
// 2. Persistence struct to save journal entries to a file

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String(lineSeperator string) string {
	return strings.Join(j.entries, lineSeperator)
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	text = fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, text)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	for idx := range j.entries {
		if idx == index {
			j.entries = append(j.entries[:idx], j.entries[idx+1:]...)
			return
		}
	}
}

type Stringer interface {
	String(lineSeperator string) string
}

type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(obj Stringer, filename string) {
	data := []byte(obj.String(p.lineSeperator))
	os.WriteFile(filename, data, 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	j.AddEntry("I slept all day.")
	fmt.Println(j.entries)
	j.RemoveEntry(2)
	fmt.Println(j.entries)

	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
