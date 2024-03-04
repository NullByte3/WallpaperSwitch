package main

import (
	"github.com/gonutz/w32/v2"
	"github.com/reujab/wallpaper"
	"os"
	"path/filepath"
	"time"
)

var (
	debug                 = false
	selected              = ""
	files, directory, err = getWallpapers()
)

func hideTerminal() {
	console := w32.GetConsoleWindow()
	if console == 0 {
		return
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(console, w32.SW_HIDE)
	}
}
func main() {
	hideTerminal()
	for {
		if err != nil || len(files) == 0 || files == nil {
			files, directory, err = getWallpapers()
		}
		files, selected = selectElement(files)
		wallpaper.SetFromFile(directory + "/" + selected)
		time.Sleep(15 * time.Second)

	}
}

func getWallpapers() ([]string, string, error) {
	wallpapersDir := filepath.Join(os.Getenv("USERPROFILE"), "Documents", "wallpapers")
	files, err := os.ReadDir(wallpapersDir)
	if err != nil {
		return nil, "", err
	}

	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	return fileNames, wallpapersDir, nil
}

func selectElement(slice []string) ([]string, string) {
	if len(slice) == 0 {
		return slice, ""
	}
	value := slice[0]
	copy(slice[:], slice[1:])
	slice = slice[:len(slice)-1]
	return slice, value
}
