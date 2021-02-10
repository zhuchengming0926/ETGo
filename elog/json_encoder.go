/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: json_encoder.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/9 16:44
 */

package elog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RegisterETJSONEncoder() error {
	return zap.RegisterEncoder("et-json", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewETJSONEncoder(cfg), nil
	})
}

type jsonHexEncoder struct {
	zapcore.Encoder
}

func NewETJSONEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	jsonEncoder := zapcore.NewJSONEncoder(cfg)
	return &jsonHexEncoder{
		Encoder: jsonEncoder,
	}
}

