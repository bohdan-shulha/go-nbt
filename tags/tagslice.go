package tags

import (
	"io"
	"encoding/binary"
	"reflect"
)

const TAG_TAG_SLICE int8 = 0x09

func ReadTagSlice(r io.Reader) ([]Tag, string, error) {
	name, err := ReadStringBody(r)
	if err != nil {
		return nil, "", err
	}

	body, _, err := ReadTagSliceBody(r)
	if err != nil {
		return nil, "", err
	}

	return body, name, nil
}

func ReadTagSliceBody(r io.Reader) ([]Tag, int8, error) {
	var bodyType int8
	err := binary.Read(r, binary.BigEndian, &bodyType)
	if err != nil {
		return nil, 0, err
	}

	var length int32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, 0, err
	}

	body := make([]Tag, length)
	for i := length - 1; i >= 0; i-- {
		switch bodyType {
			case TAG_INT8:
				body[i], err = ReadInt8Body(r)
			case TAG_INT16:
				body[i], err = ReadInt16Body(r)
			case TAG_INT32:
				body[i], err = ReadInt32Body(r)
			case TAG_INT64:
				body[i], err = ReadInt64Body(r)
			case TAG_FLOAT32:
				body[i], err = ReadFloat32Body(r)
			case TAG_FLOAT64:
				body[i], err = ReadFloat64Body(r)
			case TAG_INT8_SLICE:
				body[i], err = ReadInt8SliceBody(r)
			case TAG_STRING:
				body[i], err = ReadStringBody(r)
			case TAG_TAG_SLICE:
				body[i], _, err = ReadTagSliceBody(r)
			case TAG_COMPOUND:
				body[i], err = ReadCompoundBody(r)
			case TAG_INT32_SLICE:
				body[i], err = ReadInt32SliceBody(r)
		}
		if err != nil {
			return nil, 0, err
		}
	}
	return body, bodyType, nil
}

func WriteTagSlice(w io.Writer, name string, body []Tag) error {
	err := WriteStringBody(w, name)
	if err != nil {
		return err
	}

	err = WriteTagSliceBody(w, body)
	if err != nil {
		return err
	}

	return nil
}

func WriteTagSliceBody(w io.Writer, body []Tag) error {
	var bodyType int8
	if len(body) > 0 {
		switch reflect.TypeOf(body[0]).Kind() {
			case reflect.Int8:
				bodyType = TAG_INT8
			case reflect.Int16:
				bodyType = TAG_INT16
			case reflect.Int32:
				bodyType = TAG_INT32
			case reflect.Int64:
				bodyType = TAG_INT64
			case reflect.Float32:
				bodyType = TAG_FLOAT32
			case reflect.Float64:
				bodyType = TAG_FLOAT64
			case reflect.Slice:
				vof := reflect.ValueOf(body[0])
				if vof.Len() > 0 {
					switch vof.Index(0).Kind() {
						case reflect.Int8:
							bodyType = TAG_INT8_SLICE
						case reflect.Int32:
							bodyType = TAG_INT32_SLICE
						default:
							bodyType = TAG_TAG_SLICE
					}
				}
			case reflect.Map:
				bodyType = TAG_COMPOUND
		}
	}
	if bodyType == 0 {
		bodyType = TAG_INT8
	}
	err := binary.Write(w, binary.BigEndian, bodyType)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, int32(len(body)))
	if err != nil {
		return err
	}

	for _, body := range body {
		switch bodyType {
			case TAG_INT8:
				err = WriteInt8Body(w, body.(int8))
			case TAG_INT16:
				err = WriteInt16Body(w, body.(int16))
			case TAG_INT32:
				err = WriteInt32Body(w, body.(int32))
			case TAG_INT64:
				err = WriteInt64Body(w, body.(int64))
			case TAG_FLOAT32:
				err = WriteFloat32Body(w, body.(float32))
			case TAG_FLOAT64:
				err = WriteFloat64Body(w, body.(float64))
			case TAG_INT8_SLICE:
				err = WriteInt8SliceBody(w, body.([]int8))
			case TAG_STRING:
				err = WriteStringBody(w, body.(string))
			case TAG_TAG_SLICE:
				err = WriteTagSliceBody(w, body.([]Tag))
			case TAG_COMPOUND:
				err = WriteCompoundBody(w, body.(Compound))
			case TAG_INT32_SLICE:
				err = WriteInt32SliceBody(w, body.([]int32))
		}
		if err != nil {
			return err
		}
	}
	return nil
}