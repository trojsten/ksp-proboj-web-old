package compiler

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
)

func Unpack(r io.Reader, root string) error {
	decompress, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer decompress.Close()

	tr := tar.NewReader(decompress)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if header.Typeflag == tar.TypeDir {
			err := os.Mkdir(path.Join(root, header.Name), os.ModePerm)
			if err != nil {
				return err
			}
		} else if header.Typeflag == tar.TypeReg {
			file, err := os.Create(path.Join(root, header.Name))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
			file.Close()
		} else {
			return fmt.Errorf("unknown tar type: %v", header.Typeflag)
		}
	}

	return nil
}
