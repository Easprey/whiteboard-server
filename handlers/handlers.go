package handlers

import (
	"github.com/Easprey/whiteboard-server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// operations.<operationID>HandlerFunc expects an anonymous function with the appropriate
// signature to be passed.
// We call <operationID>(db *DataBase) so the handler can consume the database connection.
// <operationID>(db* DataBase) returns the anonymous function <operationID> expects.
// The returned anonymous function has access to the database connection.
// This paradigm is known as a closure.

func GetFingerPaths(db *DataBase) func(operations.FingerPathsGetParams) middleware.Responder {
	return func(params operations.FingerPathsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation operations.FingerPathsGet has not yet been implemented")
	}
}

func PostFingerPaths(db *DataBase) func(operations.FingerPathsPostParams) middleware.Responder {
	return func(params operations.FingerPathsPostParams) middleware.Responder {
		return middleware.NotImplemented("operation operations.FingerPathsPost has not yet been implemented")
	}
}
