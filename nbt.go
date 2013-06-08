package nbt

import (
	"nbt/tags"
	"io"
	"encoding/binary"
	"errors"
)

func ReadFrom(r io.Reader) (tags.Compound, string, error) {
	var tagType int8
	err := binary.Read(r, binary.BigEndian, &tagType)
	if err != nil {
		return nil, "", err
	}

	if tagType != tags.TAG_COMPOUND {
		return nil, "", errors.New("nbt: no compound tag found")
	}
	return tags.ReadCompound(r)
}

func WriteTo(w io.Writer, tag tags.Compound, name string) error {
	err := binary.Write(w, binary.BigEndian, tags.TAG_COMPOUND)
	if err != nil {
		return err
	}
	return tags.WriteCompound(w, name, tag)
}

func Find(c tags.Compound, searchName string) tags.Tag {
	for name, tag := range c {
		if name == searchName {
			return tag
		} else if inner, ok := tag.(tags.Compound); ok {
			if found := Find(inner, searchName); found != nil {
				return found
			}
		}
	}
	return nil
}
