package main

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"
	"boa/pkg/fetcher"
	"boa/pkg/extractor"
	"boa/pkg/install"
)

var rootCmd = &cobra.Command{
	Use:   "boa",
	Short: "A lightweight package manager for Arch packages",
}

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install a package from Arch repositories",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		// Fetch package metadata
		packageInfo, err := fetcher.FetchPackageInfo(packageName)
		if err != nil {
			log.Fatal("Failed to fetch package info:", err)
		}

		// Download package
		packageFilePath, err := fetcher.DownloadPackage(packageInfo)
		if err != nil {
			log.Fatal("Failed to download package:", err)
		}

		// Extract and install the package
		err = extractor.ExtractPackage(packageFilePath, "/home/boa")
		if err != nil {
			log.Fatal("Failed to extract package:", err)
		}

		// Install the package
		err = install.InstallPackage(packageName, "/home/boa")
		if err != nil {
			log.Fatal("Failed to install package:", err)
		}

		fmt.Printf("Package '%s' installed successfully!\n", packageName)
	},
}

func main() {
	rootCmd.AddCommand(installCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

