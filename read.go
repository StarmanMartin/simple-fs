package fs

import (
	"bufio"
	"os"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string, lineCounts int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCounts--
		if lineCounts == 0 {
			return lines, scanner.Err()
		}
	}

	return lines, scanner.Err()
}

//Exists checks if a folders or a file exists.
//
//Parameter
//
//`path` *string* Absolute path to the folder or file
//
//return
//
//`bool` `true` if the folder and exists
//`error` only if file exists but not able to read
func Exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
