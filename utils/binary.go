package utils

import (
	"bytes"
	"encoding/binary"
)

func Encode(tag int32, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, tag); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, int32(len(data))); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(b []byte) (int32, []byte, error) {
	buf := bytes.NewBuffer(b)
	var tag, length int32

	if err := binary.Read(buf, binary.BigEndian, &tag); err != nil {
		return 0, []byte{}, err
	}

	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return 0, []byte{}, err
	}

	data := make([]byte, length)
	if err := binary.Read(buf, binary.BigEndian, &data); err != nil {
		return 0, []byte{}, err
	}
	return tag, data, nil
}
