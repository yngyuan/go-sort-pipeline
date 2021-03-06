package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init()  {
	startTime = time.Now()
}

func ArraySource( a ...int) <-chan int {
	out := make(chan int)
	// use go routine
	go func() {
		for _,v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("read done:", time.Now().Sub(startTime))
		// Sort
		sort.Ints(a)
		fmt.Println("in-memory sort done:", time.Now().Sub(startTime))
		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge( in1, in2 <-chan int) <- chan int {
	out := make(chan int)
	go func() {
		// two channels could have different size of data
		// so we must check if channel is ok
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1<=v2) {
				out <- v1
				v1, ok1 = <- in1
			} else {
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
		fmt.Println("merge done:", time.Now().Sub(startTime))
	}()
	return out
}

//ReaderSource use chunkSize to control the size to read
func ReaderSource(reader io.Reader,
	chunkSize int) <- chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead +=n
			if n >0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <- chan int) {
	for v:= range in {
		buffer := make([] byte, 8)
		binary.BigEndian.PutUint64(
			buffer, uint64(v))
		writer.Write(buffer)
	}
}

//RandomSource Generate random number
func RandomSource(count int) <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//MergeN Merge n inputs
func MergeN(inputs ...<- chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) /2
	return Merge(
	MergeN(inputs[:m]...),
	MergeN(inputs[m:]...))
}

