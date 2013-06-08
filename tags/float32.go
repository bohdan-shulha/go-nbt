package tags

import (
	"io"
	"encoding/binary"
)

const TAG_FLOAT32 int8 = 0x05

func ReadFloat32(r io.Reader) (float32, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return 0, "", err
	}

	body, err := ReadFloat32Body(r)
	if err != nil {
		return 0, "", err
	}

	return body, name, nil
}

func ReadFloat32Body(r io.Reader) (float32, error) {
	var body float32
	err := binary.Read(r, binary.BigEndian, &body)
	if err != nil {
		return 0, err
	}

	return body, nil
}

func WriteFloat32(w io.Writer, name string, body float32) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteFloat32Body(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteFloat32Body(w io.Writer, body float32) error {
	err := binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}
	return nil
}
