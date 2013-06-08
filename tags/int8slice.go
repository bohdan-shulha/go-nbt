package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT8_SLICE int8 = 0x07

func ReadInt8Slice(r io.Reader) ([]int8, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return nil, "", err
	}

	body, err := ReadInt8SliceBody(r)
	if err != nil {
		return nil, "", err
	}

	return body, name, nil
}

func ReadInt8SliceBody(r io.Reader) ([]int8, error) {
	var length int32
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	body := make([]int8, length)
	err = binary.Read(r, binary.BigEndian, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func WriteInt8Slice(w io.Writer, name string, body []int8) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteInt8SliceBody(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteInt8SliceBody(w io.Writer, body []int8) error {
	err := binary.Write(w, binary.BigEndian, int32(len(body)))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, body)
	if err != nil {
		return err
	}
	
	return nil
}
