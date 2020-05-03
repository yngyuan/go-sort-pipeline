package main

import (
	"fmt"
	"go-sort-pipeline/pipeline"
)

func main()  {
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
