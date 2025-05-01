package fetcher

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

type PackageInfo struct {
	Name    string
	Version string
	URL     string
}

// FetchPackageInfo simulates fetching metadata of a package.
func FetchPackageInfo(packageName string) (*PackageInfo, error) {
	// In a real scenario, you'd parse a repo's .db file (core.db, extra.db)
	// Here, we're simplifying to just returning dummy data for example purposes.
	return &PackageInfo{
		Name:    packageName,
		Version: "5.1.008-2",
		URL:     fmt.Sprintf("https://mirror.rackspace.com/archlinux/core/os/x86_64/%s-5.1.008-2-x86_64.pkg.tar.zst", packageName),
	}, nil
}

// DownloadPackage downloads a .pkg.tar.zst package from the given URL.
func DownloadPackage(info *PackageInfo) (string, error) {
	resp, err := http.Get(info.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the local file
	outFile, err := os.Create(fmt.Sprintf("%s.pkg.tar.zst", info.Name))
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Copy data to file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return "", err
	}

	return outFile.Name(), nil
}

