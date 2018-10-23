package inputter

import (
	"io"

	ejson "encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/ktr0731/evans/adapter/protobuf"
	"github.com/ktr0731/evans/entity"
	"github.com/ktr0731/evans/usecase/port"
	"github.com/pkg/errors"
)

type json struct {
	decoder *ejson.Decoder
}

func NewJSON(in io.Reader) port.Inputter {
	return &json{
		decoder: ejson.NewDecoder(in),
	}
}

func (i *json) Input(reqType entity.Message) (proto.Message, error) {
	req := protobuf.NewDynamicMessage(reqType)
	err := i.decoder.Decode(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read input from json")
	}
	return req, nil
}
