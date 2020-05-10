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

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		if count >= 100 {
			break
		}
		fmt.Println(v)
		count ++
	}
}

