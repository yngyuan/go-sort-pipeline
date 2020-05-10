# go-sort-pipeline
a distributed external sorting demo in go

## external sorting pipeline on one machine
The external sorting pipeline goes like this. We divide the original file into 4 parts, then read and sort them separately. We have a barrier to wait till all 4 nodes have done sorting, then begin mergine. Finally use a sink node to write the result in a file.

![pipeline](https://github.com/yngyuan/go-sort-pipeline/blob/master/pipeline.png?raw=true)

