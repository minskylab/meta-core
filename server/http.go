// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package server

import (
	"context"
	"net/http"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	httpoption "github.com/RussellLuo/kun/pkg/httpoption2"
	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPRouter(svc Service, codecs httpcodec.Codecs, opts ...httpoption.Option) chi.Router {
	r := chi.NewRouter()
	options := httpoption.NewOptions(opts...)

	r.Method("GET", "/api", oas2.Handler(OASv2APIDoc, options.ResponseSchema()))

	var codec httpcodec.Codec
	var validator httpoption.Validator
	var kitOptions []kithttp.ServerOption

	codec = codecs.EncodeDecoder("DoneProcess")
	validator = options.RequestValidator("DoneProcess")
	r.Method(
		"POST", "/process/done",
		kithttp.NewServer(
			MakeEndpointOfDoneProcess(svc),
			decodeDoneProcessRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("ObtainProcess")
	validator = options.RequestValidator("ObtainProcess")
	r.Method(
		"GET", "/process/{processId}",
		kithttp.NewServer(
			MakeEndpointOfObtainProcess(svc),
			decodeObtainProcessRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PlaceStack")
	validator = options.RequestValidator("PlaceStack")
	r.Method(
		"POST", "/stack/deploy",
		kithttp.NewServer(
			MakeEndpointOfPlaceStack(svc),
			decodePlaceStackRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PurgeStack")
	validator = options.RequestValidator("PurgeStack")
	r.Method(
		"POST", "/stack/purge",
		kithttp.NewServer(
			MakeEndpointOfPurgeStack(svc),
			decodePurgeStackRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("StartProcess")
	validator = options.RequestValidator("StartProcess")
	r.Method(
		"POST", "/process/start",
		kithttp.NewServer(
			MakeEndpointOfStartProcess(svc),
			decodeStartProcessRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("UpdateStateProcess")
	validator = options.RequestValidator("UpdateStateProcess")
	r.Method(
		"POST", "/process/update",
		kithttp.NewServer(
			MakeEndpointOfUpdateStateProcess(svc),
			decodeUpdateStateProcessRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	return r
}

func NewHTTPRouterWithOAS(svc Service, codecs httpcodec.Codecs, schema oas2.Schema) chi.Router {
	return NewHTTPRouter(svc, codecs, httpoption.ResponseSchema(schema))
}

func decodeDoneProcessRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req DoneProcessRequest

		if err := codec.DecodeRequestBody(r, &_req.ProcessIdentity); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeObtainProcessRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req ObtainProcessRequest

		processId := []string{chi.URLParam(r, "processId")}
		if err := codec.DecodeRequestParam("processId", processId, &_req.ProcessId); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePlaceStackRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PlaceStackRequest

		if err := codec.DecodeRequestBody(r, &_req.Definition); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePurgeStackRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PurgeStackRequest

		if err := codec.DecodeRequestBody(r, &_req.S); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeStartProcessRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req StartProcessRequest

		if err := codec.DecodeRequestBody(r, &_req.Definition); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeUpdateStateProcessRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req UpdateStateProcessRequest

		if err := codec.DecodeRequestBody(r, &_req.Updater); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}