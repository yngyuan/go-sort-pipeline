package main

import (
	"fmt"
	"go-sort-pipeline/pipeline"
	"os"
)

func main()  {
	const filename = "large.in"
	const n = 10000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	pipeline.WriterSink(file, p)

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(file)
	count := 0
	for v := range p {
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
