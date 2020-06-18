package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Exists is check file or dir si exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// AllFiles is get all files of the folder
func AllFiles(dirPth string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			tempFiles, err := AllFiles(dirPth + PthSep + fi.Name())
			if err != nil {
				return nil, err
			}
			for _, f := range tempFiles {
				files = append(files, f)
			}
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

// GetCurrentDirectory is get current folder
func GetCurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

// GetParentDirectory is get parrent folder
func GetParentDirectory(dirctory string) string {
	dir := strings.Replace(dirctory, "\\", "/", -1)
	return dirctory[0:strings.LastIndex(dir, "/")]
}

// GetCurrentExeAbsPath get abs path
func GetCurrentExeAbsPath() (string, error) {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "\\", "/", -1), nil
}

// BuildSubDir is Create directories based on strings according to certain rules
func BuildSubDir(fileName string) string {
	windowsNotDir := []string{"CON", "PRN", "AUX", "CLOCK$", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "LPT1"}
	dirList := strings.Split(fileName, "_")
	var sunDir string = ""
	maxNum := 3
	if len(dirList)-1 < 3 {
		maxNum = len(dirList) - 1
	}
	for i := 0; i < maxNum; i++ {
		find := false
		for _, v := range windowsNotDir {
			if v == strings.ToUpper(dirList[i]) {
				find = true
				break
			}
		}
		if !find {
			sunDir += dirList[i]
			sunDir += "/"
		}
	}
	return sunDir
}

// GetFilesAndDirs is get all files and folders from parrent dir
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, dirs, nil
}

// ChmodDirRecursion is change dir Jurisdiction
func ChmodDirRecursion(dir string, m os.FileMode) error {
	files, dirs, err := GetFilesAndDirs(dir)
	if err != nil {
		return err
	}
	for _, v := range files {
		os.Chmod(v, m)
	}
	for _, v := range dirs {
		os.Chmod(v, m)
	}
	return nil
}

// FileOrDirSize is get file or dir size
func FileOrDirSize(fileDir string) int64 {
	st, err := os.Stat(fileDir)
	if err != nil {
		return 0
	}
	if st.IsDir() {
		files, _, err := GetFilesAndDirs(fileDir)
		if err != nil {
			return 0
		}
		var totalSize int64 = 0
		for _, f := range files {
			s := FileOrDirSize(f)
			totalSize += s
		}
		return totalSize
	}
	return st.Size()
}

// FileIsLock check file is lock or unlock
func FileIsLock(filePath string) bool {
	if !Exists(filePath) {
		return false
	}
	fi, _ := os.Stat(filePath)
	if fi.IsDir() {
		return false
	}
	tf, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeDevice)
	if tf != nil {
		defer tf.Close()
	}
	if err != nil {
		return true
	}
	return false
}
