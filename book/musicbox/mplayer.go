package main

import (
	"LearnGo/book/musicbox/library"
	"LearnGo/book/musicbox/mp"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handlelibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, " : ", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.Music{strconv.Itoa(id), tokens[2], tokens[3],
				tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.Remove(tokens[2])
		}else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}

	e , _ := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player:")
	fmt.Println("    lib list -- View the existing music lib")
	fmt.Println("    lib add <name><artist><source><type> -- Add a music to the music lib")
	fmt.Println("    lib remove <name> -- Remove the specified music from the lib")
	fmt.Println("    play <name> -- Play the specified music")

	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please enter command:")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e"{
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handlelibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		}else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}