package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	home := os.Getenv("HOME")
	wallpath := home + "/Pictures/wall"
	_, oldN := countFiles(wallpath)

	for {
		W, N := countFiles(wallpath)

		number := generateRandomNumber(N)
		for number == oldN {
			number = generateRandomNumber(N)
		}

		wallpaper := W[number]
		setWallpaper(wallpaper, wallpath)
		time.Sleep(300 * time.Second)
		oldN = number
	}
}

func countFiles(dirPath string) ([]string, int) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return []string{}, 0
	}
	filesarray := make([]string, 0)
	filenumbers := len(files)
	for _, file := range files {
		filesarray = append(filesarray, file.Name())
	}
	return (filesarray), filenumbers
}

func generateRandomNumber(max int) int {
	return rand.Intn(max)
}

func setWallpaper(wallpaper, wallpath string) {
	fmt.Println(wallpaper)
	script := fmt.Sprintf(`
		var Desktops = desktops();
		for (i=0; i<Desktops.length; i++) {
			d = Desktops[i];
			d.wallpaperPlugin = 'org.kde.image';
			d.currentConfigGroup = Array('Wallpaper', 'org.kde.image', 'General');
			d.writeConfig('Image', 'file://%s/%s');
		}`, wallpath, wallpaper)

	cmd := exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
