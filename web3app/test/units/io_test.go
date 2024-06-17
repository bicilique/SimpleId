package testing

import (
	"SimpleId/internal/utils"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func getProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectRoot := cwd
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, ".git")); os.IsNotExist(err) {
			projectRoot = filepath.Dir(projectRoot)
			if projectRoot == "/" || projectRoot == "." {
				return "", fmt.Errorf("project root not found")
			}
		} else {
			break
		}
	}

	return projectRoot, nil
}
func TestIo(t *testing.T) {
	utils.WriteStringToFile("simmer", "gas.txt")

	// from Executable Directory
	ex, _ := os.Executable()
	fmt.Println("Executable DIR:", filepath.Dir(ex))

	// Current working directory
	dir, _ := os.Getwd()
	fmt.Println("CWD:", dir)

	// Relative on runtime DIR:
	_, b, _, _ := runtime.Caller(0)
	d1 := path.Join(path.Dir(b))
	root := filepath.Join(filepath.Dir(b), "../..")
	fmt.Println("Relative", d1)
	fmt.Println(root)
}

func TestRoot(t *testing.T) {
	root, err := getProjectRoot()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Project Root:", root)
	}
	utils.WriteStringToFile("deployResults", root+"/"+"nama")

}
