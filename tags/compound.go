package tags

import (
	"io"
	"encoding/binary"
	"reflect"
)

const TAG_COMPOUND int8 = 0x0A
type Compound map[string]Tag

func ReadCompound(r io.Reader) (Compound, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return nil, "", err
	}

	body, err := ReadCompoundBody(r)
	if err != nil {
		return nil, "", err
	}

	return body, name, nil
}

func ReadCompoundBody(r io.Reader) (Compound, error) {
	body := make(Compound)
	var nextTagType int8
	var nextTag Tag
	var err error
	var name string
	for {
		err = binary.Read(r, binary.BigEndian, &nextTagType)
		if err != nil {
			return nil, err
		}

		switch nextTagType {
			case TAG_END:
				return body, nil
			case TAG_INT8:
				nextTag, name, err = ReadInt8(r)
			case TAG_INT16:
				nextTag, name, err = ReadInt16(r)
			case TAG_INT32:
				nextTag, name, err = ReadInt32(r)
			case TAG_INT64:
				nextTag, name, err = ReadInt64(r)
			case TAG_FLOAT32:
				nextTag, name, err = ReadFloat32(r)
			case TAG_FLOAT64:
				nextTag, name, err = ReadFloat64(r)
			case TAG_INT8_SLICE:
				nextTag, name, err = ReadInt8Slice(r)
			case TAG_STRING:
				nextTag, name, err = ReadString(r)
			case TAG_TAG_SLICE:
				nextTag, name, err = ReadTagSlice(r)
			case TAG_COMPOUND:
				nextTag, name, err = ReadCompound(r)
			case TAG_INT32_SLICE:
				nextTag, name, err = ReadInt32Slice(r)
		}
		if err != nil {
			return nil, err
		}
		body[name] = nextTag
	}
	return body, nil
}

func WriteCompound(w io.Writer, name string, body Compound) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteCompoundBody(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteCompoundBody(w io.Writer, body Compound) error {
	var err error
	for name, tag := range body {
		switch reflect.TypeOf(tag).Kind() {
			case reflect.Int8:
				err = binary.Write(w, binary.BigEndian, TAG_INT8)
				if err != nil {
					return err
				}

				err = WriteInt8(w, name, tag.(int8))
			case reflect.Int16:
				err = binary.Write(w, binary.BigEndian, TAG_INT16)
				if err != nil {
					return err
				}

				err = WriteInt16(w, name, tag.(int16))
			case reflect.Int32:
				err = binary.Write(w, binary.BigEndian, TAG_INT32)
				if err != nil {
					return err
				}

				err = WriteInt32(w, name, tag.(int32))
			case reflect.Int64:
				err = binary.Write(w, binary.BigEndian, TAG_INT64)
				if err != nil {
					return err
				}

				err = WriteInt64(w, name, tag.(int64))
			case reflect.Float32:
				err = binary.Write(w, binary.BigEndian, TAG_FLOAT32)
				if err != nil {
					return err
				}

				err = WriteFloat32(w, name, tag.(float32))
			case reflect.Float64:
				err = binary.Write(w, binary.BigEndian, TAG_FLOAT64)
				if err != nil {
					return err
				}

				err = WriteFloat64(w, name, tag.(float64))
			case reflect.Slice:
				vof := reflect.ValueOf(tag)
				if vof.Len() > 0 {
					switch vof.Index(0).Kind() {
						case reflect.Int8:
							err = binary.Write(w, binary.BigEndian, TAG_INT8_SLICE)
							if err != nil {
								return err
							}

							err = WriteInt8Slice(w, name, tag.([]int8))
						case reflect.Int32:
							err = binary.Write(w, binary.BigEndian, TAG_INT32_SLICE)
							if err != nil {
								return err
							}

							err = WriteInt32Slice(w, name, tag.([]int32))
						default:
							err = binary.Write(w, binary.BigEndian, TAG_TAG_SLICE)
							if err != nil {
								return err
							}

							err = WriteTagSlice(w, name, tag.([]Tag))
					}
				} else {
					err = binary.Write(w, binary.BigEndian, TAG_TAG_SLICE)
					if err != nil {
						return err
					}

					err = WriteTagSlice(w, name, tag.([]Tag))
				}
			case reflect.String:
				err = binary.Write(w, binary.BigEndian, TAG_STRING)
				if err != nil {
					return err
				}

				err = WriteString(w, name, tag.(string))
			case reflect.Map:
				err = binary.Write(w, binary.BigEndian, TAG_COMPOUND)
				if err != nil {
					return err
				}
				
				err = WriteCompound(w, name, tag.(Compound))
		}
		if err != nil {
			return err
		}
	}
	err = binary.Write(w, binary.BigEndian, TAG_END)
	if err != nil {
		return err
	}
	return nil
}
