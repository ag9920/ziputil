package ziputil

import (
	"archive/zip"
	"io/ioutil"
	"strings"
)

type File struct {
	RelativePath string
	Content      []byte
}

var (
	ignoreFileSubStr = []string{
		"MACOSX",
		".DS_Store",
	}
)

// GetFilesFromZip load zip file from path, decompress it, retrieve information.
func GetFilesFromZip(path string) ([]File, error) {
	zf, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	purePath := strings.TrimRight(path, ".zip")

	infos := make([]File, 0)
	for _, file := range zf.File {
		if file.FileInfo().IsDir() {
			continue
		}
		subPath := strings.TrimPrefix(file.Name, purePath+"/")
		if subPath == "" || needIgnore(subPath) {
			continue
		}
		content, err := readAll(file)
		if err != nil {
			return nil, err
		}
		infos = append(infos, File{
			RelativePath: subPath,
			Content:      content,
		})
	}
	return infos, nil
}

// readAll is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func readAll(file *zip.File) ([]byte, error) {
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fc.Close()

	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func needIgnore(s string) bool {
	for _, i := range ignoreFileSubStr {
		if strings.Contains(s, i) {
			return true
		}
	}
	return false
}
