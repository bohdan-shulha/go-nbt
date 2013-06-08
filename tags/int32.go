package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT32 int8 = 0x03

func ReadInt32(r io.Reader) (int32, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadInt32Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadInt32Body(r io.Reader) (int32, error) {
	var body int32
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteInt32(w io.Writer, name string, body int32) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}
	
	err = WriteInt32Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteInt32Body(w io.Writer, body int32) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}

	return nil
}