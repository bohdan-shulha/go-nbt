package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT64 int8 = 0x04

func ReadInt64(r io.Reader) (int64, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadInt64Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadInt64Body(r io.Reader) (int64, error) {
	var body int64	
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteInt64(w io.Writer, name string, body int64) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteInt64Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteInt64Body(w io.Writer, body int64) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}
	
	return nil
}