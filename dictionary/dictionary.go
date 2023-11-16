package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "dictionary.txt"

type Dictionary struct {
}

func New() *Dictionary {
	return &Dictionary{}
}

func (d *Dictionary) Add(word, definition string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s|%s\n", word, definition)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dictionary) Get(word string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 2 && parts[0] == word {
			return parts[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("word not found")
}

func (d *Dictionary) Remove(word string) error {
	lines, err := readLines(filename)
	if err != nil {
		return err
	}

	var newLines []string
	removed := false
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 || parts[0] != word {
			newLines = append(newLines, line)
		} else {
			removed = true
		}
	}

	if !removed {
		return fmt.Errorf("word not found")
	}

	if err := writeLines(filename, newLines); err != nil {
		return err
	}
	return nil
}

func (d *Dictionary) List() (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}
