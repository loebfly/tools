package fileT

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// SelfPath returns the absolute path of the current executable file
func (f Enter) SelfPath() string {
	abs, _ := filepath.Abs(os.Args[0])
	return abs
}

// SelfDir returns the directory of the current executable file
func (f Enter) SelfDir() string {
	return filepath.Dir(f.SelfPath())
}

// GetBasename returns the base name of the file
func (f Enter) GetBasename(filePath string) string {
	return path.Base(filePath)
}

// GetDir returns the directory of the file
func (f Enter) GetDir(filePath string) string {
	return path.Dir(filePath)
}

// GetExt returns the extension of the file
func (f Enter) GetExt(filePath string) string {
	return path.Ext(filePath)
}

// GetMTime returns the modification time of the file
func (f Enter) GetMTime(filePath string) (int64, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fi.ModTime().Unix(), nil
}

// GetSize returns the size of the file
func (f Enter) GetSize(filePath string) (size int64, err error) {
	fi, e := os.Stat(filePath)
	if e != nil {
		return 0, e
	}
	return fi.Size(), nil
}

// Rename renames a file
func (f Enter) Rename(filePath string, to string) error {
	return os.Rename(filePath, to)
}

// IsExist checks whether a file or directory exists
func (f Enter) IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// IsFile checks whether the path is a file,
func (f Enter) IsFile(filePath string) bool {
	fi, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// IsDir checks whether the path is a directory
func (f Enter) IsDir(filePath string) bool {
	fi, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

// WriteBytes writes bytes to a file
func (f Enter) WriteBytes(filePath string, data []byte) (int, error) {
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer func(fw *os.File) {
		err := fw.Close()
		if err != nil {
			panic(err)
		}
	}(fw)
	return fw.Write(data)
}

// WriteString writes string to a file
func (f Enter) WriteString(filePath string, data string) (int, error) {
	return f.WriteBytes(filePath, []byte(data))
}

// ReadBytes reads bytes from a file
func (f Enter) ReadBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadString reads string from a file
func (f Enter) ReadString(filePath string) (string, error) {
	b, err := f.ReadBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadLines reads lines from a file
func (f Enter) ReadLines(filePath string) ([]string, error) {
	b, err := f.ReadBytes(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), "\n"), nil
}

// Remove removes a file or directory
func (f Enter) Remove(path string) error {
	if f.IsDir(path) {
		return os.RemoveAll(path)
	} else if f.IsFile(path) {
		return os.Remove(path)
	} else {
		return errors.New("not a file or directory")
	}
}
