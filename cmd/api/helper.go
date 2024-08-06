package main

import (
	"baseGo/internal/dto/payload"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	InvalidParameter          = "invalid %s parameter"
	InvalidRequestIdParameter = "invalid request id parameter"
	InvalidUUID               = "invalid uuid"

	JsonUnknownField             = "json: unknown field"
	IncorrectJsonTypeForField    = "body contains incorrect JSON type for field %q"
	IncorrectJsonTypeAtCharacter = "body contains incorrect JSON type (at character %d)"
	BadlyFormedJsonAtCharacter   = "body contains badly-formed JSON at (character %d)"
	BadlyFormedJson              = "body contains badly-formed JSON"
	BodyMustNotBeEmpty           = "body must not be empty"
	BodyContainsUnknownKey       = "body contains unknown key %s"
	RequestBodyTooLarge          = "http: request body too large"
	BodyTooLargeMessage          = "body must not be larger than %d bytes"
	BodyContainSingleJsonValue   = "body must contain a single JSON value"
	FailureToCommitDBTransaction = "failed to commit transaction"
)

func (app *application) writeJson(writer http.ResponseWriter, status int, data payload.BaseResponse,
	headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		writer.Header()[key] = value
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if _, err := writer.Write(js); err != nil {
		app.logger.PrintError(err, nil)
		return err
	}
	app.logger.PrintInfo("response: ", map[string]interface{}{
		"status": status,
		"body":   data,
		"time":   time.Now(),
	})
	return nil
}

func (app *application) readJson(writer http.ResponseWriter, request *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	request.Body = http.MaxBytesReader(writer, request.Body, int64(maxBytes))

	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf(BadlyFormedJsonAtCharacter, syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New(BadlyFormedJson)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf(IncorrectJsonTypeForField,
					unmarshalTypeError.Field)
			}
			return fmt.Errorf(IncorrectJsonTypeAtCharacter,
				unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New(BodyMustNotBeEmpty)

		case strings.HasPrefix(err.Error(), JsonUnknownField):
			fieldName := strings.TrimPrefix(err.Error(), JsonUnknownField)
			return fmt.Errorf(BodyContainsUnknownKey, fieldName)

		case err.Error() == RequestBodyTooLarge:
			return fmt.Errorf(BodyTooLargeMessage, maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New(BodyContainSingleJsonValue)
	}
	return nil
}

func (app *application) readStrings(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

func (app *application) readIDParam(r *http.Request, key string) (uuid.UUID, error) {
	params := r.PathValue(key)

	id, err := uuid.Parse(params)
	if err != nil {
		return uuid.Nil, fmt.Errorf(InvalidParameter, key)
	}

	return id, nil
}

func (app *application) commitDBTransaction(ctx context.Context, tx pgx.Tx) error {
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf(FailureToCommitDBTransaction)
	}
	return nil
}
