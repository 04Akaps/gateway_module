package common

import (
	"github.com/04Akaps/gateway_module/log"
	"github.com/bytedance/sonic"
)

type jsonHandler struct {
	marshal   func(val interface{}) ([]byte, error)
	unmarshal func(buf []byte, val interface{}) error
}

var JsonHandler jsonHandler

func init() {
	JsonHandler = jsonHandler{
		marshal:   sonic.Marshal,
		unmarshal: sonic.Unmarshal,
	}
}

func (j jsonHandler) Marshal(v interface{}) ([]byte, error) {
	bytes, err := j.marshal(v)

	if err != nil {
		log.Log.Error("Failed to marshal")
		return nil, err
	}

	return bytes, nil
}

func (j jsonHandler) Unmarshal(v []byte, buffer interface{}) error {
	err := j.unmarshal(v, buffer)

	if err != nil {
		log.Log.Error("Failed to marshal")
		return err
	}

	return nil
}

func (j jsonHandler) Handle(v interface{}, buffer interface{}) error {
	bytes, err := j.Marshal(v)

	if err != nil {
		return err
	}

	err = j.Unmarshal(bytes, buffer)

	if err != nil {
		return err
	}

	return nil
}
