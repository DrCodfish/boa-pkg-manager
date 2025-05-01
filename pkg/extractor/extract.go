package extractor

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"github.com/klauspost/compress/zstd"
)

// ExtractPackage extracts the downloaded .pkg.tar.zst to a specified directory.
func ExtractPackage(packagePath, installDir string) error {
	// Open the .pkg.tar.zst file
	file, err := os.Open(packagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decompress using zstd
	d, err := zstd.NewReader(file)
	if err != nil {
		return err
	}
	defer d.Close()

	// Untar the file
	tarReader := tar.NewReader(d)
	for {
		header, err := tarReader.Next()
		if err != nil {
			break
		}

		// Create the file path
		targetPath := fmt.Sprintf("%s/%s", installDir, header.Name)
		outFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		// Copy content to file
		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			return err
		}
	}

	return nil
}

