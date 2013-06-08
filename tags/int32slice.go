package tags

import (
	"io"
	"encoding/binary"
)

const TAG_INT32_SLICE int8 = 0x0B

func ReadInt32Slice(r io.Reader) ([]int32, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return nil, "", err
	}

	body, err := ReadInt32SliceBody(r)
	if err != nil {
		return nil, "", err
	}

	return body, name, nil
}

func ReadInt32SliceBody(r io.Reader) ([]int32, error) {	
	var length int32
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	body := make([]int32, length)
	err = binary.Read(r, binary.BigEndian, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func WriteInt32Slice(w io.Writer, name string, body []int32) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteInt32SliceBody(w, body)
	if err != nil {
		return err
	}
	return nil
}

func WriteInt32SliceBody(w io.Writer, body []int32) error {
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