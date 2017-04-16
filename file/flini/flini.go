package flini

import (
	"bufio"
	"io"
	"os"

	"strings"

	"github.com/pkg/errors"
)

type File struct {
	Sections map[string][]Section
}

type Section struct {
	Name   string
	Values map[string][]interface{}
}

func ParseFile(path string) (File, error) {
	stream, err := os.Open(path)
	if err != nil {
		return File{}, errors.Wrap(err, "couldn't open the INI")
	}
	return Parse(stream)
}

func Parse(r io.Reader) (File, error) {
	reader := bufio.NewReader(r)

	ret := File{Sections: map[string][]Section{}}

	curSection := &Section{Values: map[string][]interface{}{}}

	lineNum := -1
	for {
		lineNum++
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return File{}, errors.Wrapf(err, "Error on line %v", lineNum)
		}
		commentIdx := strings.Index(line, ";")

		// Whole section can be commented as [;Name], skip it then
		if commentIdx > 0 && line[commentIdx-1] == '[' {
			if curSection.Name != "" {
				ret.Sections[curSection.Name] = append(ret.Sections[curSection.Name], *curSection)
			}
			curSection = nil
			continue
		}
		if commentIdx >= 0 {
			line = line[:commentIdx]
		}
		line = strings.Trim(line, "\r \t\n")
		if len(line) == 0 {
			continue
		}
		// Section has begun
		if line[0] == '[' {
			if curSection != nil && curSection.Name != "" {
				ret.Sections[curSection.Name] = append(ret.Sections[curSection.Name], *curSection)
			}
			curSection = &Section{Values: map[string][]interface{}{}, Name: line[1 : len(line)-1]}
			continue
		}

		if curSection == nil {
			continue // skipping this commented category
		}

		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}

		fieldName := fields[0]
		fieldName = strings.Trim(fieldName, "\r \t\n")
		fieldValue := fields[1]
		fieldValue = strings.Trim(fieldValue, "\r \t\n")
		curSection.Values[fieldName] = append(curSection.Values[fieldName], fieldValue)
	}
	return ret, nil
}
