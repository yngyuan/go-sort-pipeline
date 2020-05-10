package main

import (
	"bufio"
	"fmt"
	"go-sort-pipeline/pipeline"
	"os"
)

func main()  {
	const filename = "sort.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	p := pipeline.RandomSource(n)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// now read the file we generated earlier
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {    // to use range, the channel must be closed
		if count >= 100 {
			break
		}
		fmt.Println(v)
		count ++
	}
}

func mergeDemo()  {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(
				3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(
				7, 9, 0, 3, 1, 8, 16, 11)))

	for v := range p {
		fmt.Println(v)
	}
}