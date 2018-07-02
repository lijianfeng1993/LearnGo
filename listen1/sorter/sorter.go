package main

import (
	"fmt"
	"flag"
	"LearnGo/listen1/sorter/algorithms/bubblesort"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sort.")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "bubblesort", "Sort Algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Fail to open input file " , infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected.")
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	//return values, nil
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Fail to create outfile ", outfile)
		return err
	}

	defer file.Close()

	for _ , value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			bubblesort.Bubblesort(values)
		default:
			fmt.Println("Sorting Algorithm ", *algorithm ," is not support.")
		}
		t2 := time.Now()
		fmt.Println("Sort process cost ", t2.Sub(t1), " to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}