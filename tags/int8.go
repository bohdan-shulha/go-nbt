package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT8 int8 = 0x01

func ReadInt8(r io.Reader) (int8, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadInt8Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadInt8Body(r io.Reader) (int8, error) {
	var body int8
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteInt8(w io.Writer, name string, body int8) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteInt8Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteInt8Body(w io.Writer, body int8) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}

	return nil
}
