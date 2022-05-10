# ziputil


A Simple Tool on Dealing with Zip File, written in Go.


- GetFilesFromZip

`func GetFilesFromZip(path string) ([]File, error)`

load zip file from path, decompress it, retrieve information.
