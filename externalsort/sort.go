package main

import (
	"bufio"
	"fmt"
	"go-sort-pipeline/pipeline"
	"os"
	"strconv"
)

func main()  {
	// TODO handle case when filesize%chunkCount != 0
	// here we divide the file of 512 bytes to 4 parts
	p := createNetworkPipeline("sort.in", 512, 4)
	writeToFile(p, "sort.out")
	printFile("sort.out")
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v:= range p {
		fmt.Println(v)
	}
}

func writeToFile(p <-chan int, filename string)  {
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline(
	filename string,
	fileSize,
	chunkCount int) <-chan int  {
	chunkSize := fileSize/ chunkCount
	pipeline.Init()

	sortResults := []<-chan int{}
	for i:= 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		// TODO should close this
		if err !=nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(
	filename string,
	filesize,
	chunkCount int) <-chan int  {
	chunkSize := filesize/ chunkCount
	pipeline.Init()

	sortAddr := []string{}
	for i:= 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		// TODO should close this
		if err !=nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)
		addr := ":"+strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}