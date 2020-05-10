# go-sort-pipeline
a distributed external sorting demo in go

## 1. external sorting pipeline on one machine
The external sorting pipeline goes like this. We divide the original file into 4 parts, then read and sort them separately. We have a barrier(this is done by go in goroutine) to wait till all 4 nodes have done sorting, then begin to merge. Finally use a sink node to write the result in a file.

![pipeline](https://github.com/yngyuan/go-sort-pipeline/blob/master/pipeline.png?raw=true)

## 2. distributed external sorting pipline
There could be times that we cannot even store data on one machine. So we develop the pipline to a distributed one.

InMemSort --> Writer Sink ========> Reader Source --> Merge

                     


