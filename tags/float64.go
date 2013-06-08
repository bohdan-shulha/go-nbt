package tags

import (
	"io"
	"encoding/binary"
)

const TAG_FLOAT64 int8 = 0x06

func ReadFloat64(r io.Reader) (float64, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadFloat64Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadFloat64Body(r io.Reader) (float64, error) {
	var body float64
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteFloat64(w io.Writer, name string, body float64) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteFloat64Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteFloat64Body(w io.Writer, body float64) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}
	
	return nil
}