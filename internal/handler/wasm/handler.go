//go:build js && wasm

package wasm

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/mapper"
	"github.com/kingmidas74/gonesis-engine/internal/mapper/model"
)

func (h *Handler) serializeWorld(w contracts.World) string {
	res := mapper.NewWorld(w)

	r, e := json.Marshal(res)
	if e != nil {
		return h.serializeResponse(1, e.Error())
	}

	return h.serializeResponse(0, string(r))
}

func (h *Handler) serializeResponse(code int, message string) string {
	r, err := json.Marshal(model.Response{
		Code:    code,
		Message: message,
	})
	if err != nil {
		return err.Error()
	}
	return string(r)
}

func (h *Handler) deserializeConfiguration(configJson string) (*configuration.Configuration, error) {
	config := configuration.NewConfiguration()
	err := config.FromJson(configJson)
	return config, err
}
