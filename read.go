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
