package delta

import (
	"os"

	"github.com/balacode/go-delta"
	"github.com/iimrudy/vUtils/files"
)

// delta of files & execute patches on file

// p1, p2 are the files path from which the delta is calculated
func GenerateDiffFormFilesByPath(p1, p2 string) (*delta.Delta, error) {
	if f1, err := os.Open(p1); err != nil {
		return nil, err
	} else if f2, err := os.Open(p2); err != nil {
		return nil, err
	} else {
		return GenerateDiffFormFiles(f1, f2)
	}
}

// f1, f2 are the files from which the delta is calculated
func GenerateDiffFormFiles(f1, f2 *os.File) (*delta.Delta, error) {
	if fb1, err := files.ReadFileAsBytes(f1); err != nil {
		return nil, err
	} else if fb2, err := files.ReadFileAsBytes(f2); err != nil {
		return nil, err
	} else {
		return cDelta(fb1, fb2), nil
	}
}

func cDelta(b1, b2 []byte) *delta.Delta {
	d := delta.Make(b1, b2)
	return &d
}

// apple diff to files
