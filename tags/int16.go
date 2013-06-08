package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT16 int8 = 0x02

func ReadInt16(r io.Reader) (int16, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadInt16Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadInt16Body(r io.Reader) (int16, error) {
	var body int16
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteInt16(w io.Writer, name string, body int16) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteInt16Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteInt16Body(w io.Writer, body int16) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}

	return nil
}