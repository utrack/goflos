package flini

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"strings"

	"github.com/pkg/errors"
)

type File struct {
	Sections map[string][]Section
}

type Section struct {
	Name     string
	Settings map[string]Settings
}

type Settings []Values

func (s Settings) V() Values {
	if len(s) > 0 {
		return s[0]
	}
	return Values{}
}

type Values []interface{}

func (v Values) String(pos int) string {
	val := v[pos]
	switch val := val.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case float32:
		return fmt.Sprint(val)
	default:
		panic(fmt.Sprintf("Unknown var type %v", reflect.TypeOf(val).String()))
	}
}

func (v Values) Float(pos int) (float32, error) {
	if pos >= len(v) {
		return 0, errors.Errorf("No such item at pos %v", pos)
	}
	if pos, ok := v[pos].(float32); ok {
		return pos, nil
	}
	return 0, errors.Errorf("Item at %v is not float32, %v", pos, reflect.TypeOf(v[pos]).String())
}

func ParseFile(path string) (*File, error) {
	stream, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't open the INI")
	}
	return Parse(stream)
}

func Parse(r io.ReadSeeker) (*File, error) {
	hdr := make([]byte, 4)
	_, err := r.Read(hdr)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read file header")
	}
	r.Seek(0, io.SeekStart)
	if bytes.Compare(hdr, biniHeader) == 0 {
		return ParseBINI(r)
	}
	return ParseINI(r)
}

func ParseINI(r io.Reader) (*File, error) {
	reader := bufio.NewReader(r)

	ret := &File{Sections: map[string][]Section{}}

	curSection := &Section{Settings: map[string]Settings{}}

	lineNum := -1
	for {
		lineNum++
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrapf(err, "Error on line %v", lineNum)
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
			curSection = &Section{Settings: map[string]Settings{}, Name: line[1 : len(line)-1]}
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
		curSection.Settings[fieldName] = append(curSection.Settings[fieldName], valuesFromSetting(fieldValue))
	}
	return ret, nil
}

func valuesFromSetting(str string) []interface{} {
	vals := strings.Split(str, ",")
	ret := make([]interface{}, 0, len(vals))

	for _, val := range vals {
		val = strings.Trim(val, " \t\n\r")
		if val[0] == '"' {
			// We see ""s, this is a string
			val = strings.Trim(val, `"`)
			ret = append(ret, val)
			continue
		}
		if strings.Index(val, ".") != -1 {
			// float
			f, err := strconv.ParseFloat(val, 32)
			if err == nil {
				ret = append(ret, float32(f))
				continue
			}
		}

		if i, err := strconv.Atoi(val); err == nil {
			ret = append(ret, int(i))
			continue
		}
		ret = append(ret, val)
	}
	return ret
}
