package tags

import (
	"io"
	"encoding/binary"
)

const TAG_STRING int8 = 0x08

func ReadString(r io.Reader) (string, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return "", "", err
	}

	body, err := ReadStringBody(r)
	if err != nil {
		return "", "", err
	}

	return body, name, nil
}

func ReadStringBody(r io.Reader) (string, error) {
	var length uint16
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}

	stringBytes := make([]byte, length)
	_, err  = r.Read(stringBytes)
	if err != nil {
		return "", err
	}

	return string(stringBytes), nil
}

func WriteString(w io.Writer, name string, body string) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteStringBody(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteStringBody(w io.Writer, body string) error {
	stringBytes := []byte(body)
	err := binary.Write(w, binary.BigEndian, int16(len(stringBytes)))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, stringBytes)
	if err != nil {
		return err
	}

	return nil
}