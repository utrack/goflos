package flini

import (
	"bytes"
	"os"

	"encoding/binary"

	"io/ioutil"

	"io"

	"github.com/pkg/errors"
)

func ParseBINI(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't open the file")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't slurp BINI file")
	}
	return parseBINI(bytes.NewReader(b))
}
func parseBINI(r *bytes.Reader) (*File, error) {
	header := make([]byte, 4)
	_, err := r.Read(header)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read header")
	}

	if bytes.Compare(header, biniHeader) != 0 {
		return nil, errors.Errorf("unknown header for BINI: %v", string(header))
	}

	var version int32
	err = binary.Read(r, binary.LittleEndian, &version)
	if err != nil {
		return nil, err
	}
	if version != 1 {
		return nil, errors.Errorf("unknown BINI version %v", version)
	}

	var stringTableOffset int32
	err = binary.Read(r, binary.LittleEndian, &stringTableOffset)
	if err != nil {
		return nil, err
	}
	var stringTableLen = r.Size() - int64(stringTableOffset)
	if stringTableLen < 0 {
		return nil, errors.Errorf("StringTableLength < 0 : %v", stringTableLen)
	}
	r.Seek(int64(stringTableOffset), io.SeekStart)

	var stringTable stringTable
	{
		stringTableBytes := make([]byte, stringTableLen)
		_, err = r.Read(stringTableBytes)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't read StringTable")
		}
		stringTable = newStringTable(stringTableBytes)
	}
	r.Seek(12, io.SeekStart)

	ret := &File{Sections: map[string][]Section{}}
	pos := int64(12)
	for pos+4 < int64(stringTableOffset) {
		var nameOffset int16
		err = binary.Read(r, binary.LittleEndian, &nameOffset)
		if err != nil {
			return nil, errors.Wrap(err, "error when reading nameOffset")
		}
		name, ok := stringTable[nameOffset]
		if !ok {
			return nil, errors.Errorf("section name not found by offset %v", nameOffset)
		}

		var entryCount int16
		err = binary.Read(r, binary.LittleEndian, &entryCount)
		if err != nil {
			return nil, errors.Wrap(err, "error when reading entryCount")
		}

		curSection := Section{
			Name:   name,
			Values: map[string][]interface{}{},
		}

		for i := 0; i < int(entryCount); i++ {
			entName, entVals, err := parseEntry(r, stringTable)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing section %v", curSection.Name)
			}
			curSection.Values[entName] = entVals
		}

		pos, _ = r.Seek(0, io.SeekCurrent)

		ret.Sections[name] = append(ret.Sections[name], curSection)
	}

	return ret, nil
}

var (
	biniHeader = []byte("BINI")
)

type stringTable map[int16]string

func parseEntry(r *bytes.Reader, st stringTable) (string, []interface{}, error) {
	var nameOffset int16
	err := binary.Read(r, binary.LittleEndian, &nameOffset)
	if err != nil {
		return "", nil, errors.Wrap(err, "error when reading nameOffset")
	}
	name, ok := st[nameOffset]
	if !ok {
		return "", nil, errors.Errorf("entry name not found by offset %v", nameOffset)
	}

	valueCount, _ := r.ReadByte()

	values := make([]interface{}, 0, int(valueCount))
	for i := 0; i < int(valueCount); i++ {
		valType, err := r.ReadByte()
		if err != nil {
			return "", nil, err
		}
		var val interface{}
		switch valType {
		case 1:
			var v int32
			err = binary.Read(r, binary.LittleEndian, &v)
			val = v
		case 2:
			var v float32
			err = binary.Read(r, binary.LittleEndian, &v)
			val = v
		case 3:
			var idx int32
			err = binary.Read(r, binary.LittleEndian, &idx)
			v, ok := st[int16(idx)]
			if !ok {
				err = errors.New("corrupt BINI: couldn't find string by index")
			}
			val = v
		default:
			err = errors.Errorf("unknown BINI value type %v", valType)
		}
		if err != nil {
			return "", nil, errors.Wrap(err, "couldn't retrieve value")
		}

		values = append(values, val)
	}
	return name, values, nil
}

func newStringTable(st []byte) stringTable {
	ret := stringTable{}
	r := bytes.NewReader(st)

	curBuf := bytes.NewBuffer([]byte{})

	var curPos int16
	var curOffset int64
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		if b == 0x0 {
			ret[int16(curOffset)] = curBuf.String()
			curOffset, _ = r.Seek(0, io.SeekCurrent)
			curBuf.Reset()
			curPos++
			continue
		}
		if err = curBuf.WriteByte(b); err != nil {
			panic(err)
		}

		curPos++
	}
	return ret
}
