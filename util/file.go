package util

import (
	"bufio"
	"os"
)

func ForEachLine(f string, reader func(string)) {
	file, err := os.Open(f)
	Check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reader(scanner.Text())
	}
	err = file.Close()
	Check(err)
}
