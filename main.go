package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// set by goreleaser
var (
	version = "v0.0.0"
	date    = "-"
)

// logic

var wg sync.WaitGroup

func handleDir(dir string, patternlist []string) {
	defer wg.Done()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			fullPath := filepath.Join(dir, f.Name())
			if checkName(f.Name(), patternlist) {
				go exclude(fullPath)
			} else {
				go handleDir(fullPath, patternlist)
			}
		}
	}
}

func checkName(name string, patternList []string) bool {
	for _, pattern := range patternList {
		if match, _ := filepath.Match(pattern, name); match {
			return true
		}
	}

	return false
}

func exclude(dir string) {
	defer wg.Done()

	var err error

	switch runtime.GOOS {
	case "windows":
		err = excludeWindows(dir)
	case "linux":
		err = excludeLinux(dir)
	case "darwin":
		err = excludeDarwin(dir)
	default:
		log.Fatalf("Unknown GOOS %v", runtime.GOOS)
	}

	if err != nil {
		log.Printf("Error excluding %v: %v\n", dir, err)
	} else {
		log.Printf("Excluded %v\n", dir)
	}
}

func excludeWindows(dir string) error {
	ads, err := os.OpenFile(dir+":com.dropbox.ignored", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = ads.WriteString("1")
	if err != nil {
		_ = ads.Close()
		return err
	}

	err = ads.Close()
	if err != nil {
		return err
	}

	return nil
}

func excludeLinux(dir string) error {
	return exec.Command("attr", "-s", "com.dropbox.ignored", "-V", "1", dir).Run()
}

func excludeDarwin(dir string) error {
	return exec.Command("xattr", "-w", "com.dropbox.ignored", "1", dir).Run()
}

// command line handling

func defaultDropboxFolder() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homePath, "Dropbox")
}

func main() {
	var dropboxFolder string

	var rootCmd = &cobra.Command{
		Use:     "dropbox-ignorer <folder-name> [another-folder-name] [...]",
		Short:   "Quick and dirty way to make Dropbox ignore multiple folders by pattern (like node_modules)",
		Version: fmt.Sprintf("%v (built on %v)", version, date),
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Scanning %v for %v...\n", dropboxFolder, strings.Join(args, ", "))
			wg.Add(1)
			handleDir(dropboxFolder, args)
			wg.Wait()
		},
	}

	rootCmd.Flags().StringVarP(&dropboxFolder, "dropbox-folder", "d", defaultDropboxFolder(), "Override default Dropbox folder")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
