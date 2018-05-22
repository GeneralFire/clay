// Code generated by protoc-gen-goclay
// source: sum.proto
// DO NOT EDIT!

/*
Package sumpb is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay Service interface.
*/
package sumpb

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/utrack/clay/transport"
	"github.com/utrack/clay/transport/httpruntime"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ chi.Router

// SummatorDesc is a descriptor/registrator for the SummatorServer.
type SummatorDesc struct {
	svc SummatorServer
}

// NewSummatorServiceDesc creates new registrator for the SummatorServer.
func NewSummatorServiceDesc(svc SummatorServer) *SummatorDesc {
	return &SummatorDesc{svc: svc}
}

// RegisterGRPC implements service registrator interface.
func (d *SummatorDesc) RegisterGRPC(s *grpc.Server) {
	RegisterSummatorServer(s, d.svc)
}

// SwaggerDef returns this file's Swagger definition.
func (d *SummatorDesc) SwaggerDef() []byte {
	return _swaggerDef_sum_proto
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *SummatorDesc) RegisterHTTP(mux transport.Router) {

	// Handlers for Sum

	mux.MethodFunc(pattern_goclay_Summator_Sum_0, "GET", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req SumRequest
		err := unmarshaler_goclay_Summator_Sum_0(r, &req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't parse request"))
			return
		}

		ret, err := d.svc.Sum(r.Context(), &req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "returned from handler"))
			return
		}

		_, outbound := httpruntime.MarshalerForRequest(r)
		w.Header().Set("Content-Type", outbound.ContentType())
		err = outbound.Marshal(w, ret)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
			return
		}
	})

}

var _swaggerDef_sum_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "sum.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/sum/{a}/{b}": {
      "get": {
        "operationId": "Sum",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/sumpbSumResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "a",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "b",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Summator"
        ]
      }
    }
  },
  "definitions": {
    "sumpbSumRequest": {
      "type": "object",
      "properties": {
        "a": {
          "type": "string",
          "format": "int64",
          "description": "A is the number we're adding to. Can't be zero for the sake of example."
        },
        "b": {
          "type": "string",
          "format": "int64",
          "description": "B is the number we're adding."
        }
      },
      "description": "SumRequest is a request for Summator service."
    },
    "sumpbSumResponse": {
      "type": "object",
      "properties": {
        "sum": {
          "type": "string",
          "format": "int64"
        },
        "error": {
          "type": "string"
        }
      }
    }
  }
}

`)

var (
	pattern_goclay_Summator_Sum_0     = "/v1/example/sum/{a}/{b}"
	unmarshaler_goclay_Summator_Sum_0 = func(r *http.Request, req *SumRequest) error {

		rctx := chi.RouteContext(r.Context())
		if rctx == nil {
			panic("Only chi router is supported for GETs atm")
		}
		for pos, k := range rctx.URLParams.Keys {
			runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos])
		}
		return nil

	}
)
