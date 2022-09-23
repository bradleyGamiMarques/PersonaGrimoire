// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Arcana defines model for Arcana.
type Arcana struct {
	// ArcanaID A universally unique identifier for identifying one of the 22 Major Arcana.
	ArcanaID uuid.UUID `gorm:"type:char(36);primary_key" json:"ArcanaID"`

	// ArcanaName The name of the Major Arcana.
	ArcanaName string `json:"ArcanaName" validate:"min=3,max=10"`

	// ArcanaNumber A number from 0 - 21 representing one of the 22 Major Arcana from tarot cards.
	ArcanaNumber int `json:"ArcanaNumber" validate:"min=0,max=21"`

	// ArcanaNumeral A Roman Numeral representation of the Arcana Number. The exeception is the Fool Arcana which does not have an associated numeral. It is represented with "0".
	ArcanaNumeral string `json:"ArcanaNumeral" validate:"min=1,max=3"`
}

// ErrorBaseResponse defines model for ErrorBaseResponse.
type ErrorBaseResponse struct {
	Code    int    `json:"Code"`
	Data    string `json:"Data"`
	Error   bool   `json:"Error"`
	Message string `json:"Message"`
	Ping    bool   `json:"Ping"`
}

// P5Persona defines model for P5Persona.
type P5Persona struct {
	// Arcana A universally unique identifier for identifying one of the 22 Major Arcana. Each Persona has one arcana.
	Arcana uuid.UUID `gorm:"type:char(36);primary_key" json:"Arcana"`

	// CreatedAt Represents when the Persona was added to the database.
	CreatedAt time.Time `json:"-"`

	// DLC Represents if the Persona is only available via Downloadable Content.
	DLC bool `json:"DLC"`

	// Level The level of the Persona 5 when first encountered during fusion. The main character must be at least this level in order to fuse this Persona.
	Level int `json:"Level" validate:"min=1,max=99"`

	// Name The name of the Persona 5 Persona.
	Name string `json:"Name" validate:"min=1,max=24"`

	// PersonaID A universally unique identifier for identifying a Persona 5 Persona.
	PersonaID uuid.UUID `gorm:"type:char(36);primary_key" json:"PersonaID"`

	// Skill A universally unique identifier for identifying a skill that a Persona can learn.
	Skill *uuid.UUID `gorm:"type:char(36);primary_key" json:"Skill,omitempty"`

	// TreasureDemon Represents if the Persona is a treasure demon. Unique field to Persona 5 and Persona 5 Royal.
	TreasureDemon bool `json:"TreasureDemon"`

	// UpdatedAt Represents the last time when the Persona was updated in the database.
	UpdatedAt time.Time `json:"-"`
}

// P5PersonaSkill defines model for P5PersonaSkill.
type P5PersonaSkill struct {
	// CreatedAt Represents when the skill was added to the database.
	CreatedAt *time.Time `json:"-"`

	// SkillCost The cost to use the skill.
	SkillCost *string `json:"SkillCost,omitempty"`

	// SkillEffect The in-game description of what the skill does when used by the player.
	SkillEffect *string `json:"SkillEffect,omitempty"`

	// SkillID A universally unique identifier for identifying a skill that a Persona can learn.
	SkillID *uuid.UUID `gorm:"type:char(36);primary_key" json:"SkillID,omitempty"`

	// SkillName The in-game name for the skill.
	SkillName *string `json:"SkillName,omitempty"`

	// UpdatedAt Represents the last time when the skill was updated in the database.
	UpdatedAt *time.Time `json:"-"`
}

// P5PersonaStats defines model for P5PersonaStats.
type P5PersonaStats struct {
	// Agility An integer that represents the Persona's Agility stat.
	Agility int `json:"Agility" validate:"min=1,max=99"`

	// Endurance An integer that represents the Persona's Endurance stat.
	Endurance int `json:"Endurance" validate:"min=1,max=99"`

	// Luck An integer that represents the Persona's Luck stat.
	Luck int `json:"Luck" validate:"min=1,max=99"`

	// Magic An integer that represents the Persona's Magic stat.
	Magic int `json:"Magic" validate:"min=1,max=99"`

	// StatsID A universally unique identifier for identifying a Persona 5 Persona's stats.
	StatsID uuid.UUID `gorm:"type:char(36);primary_key" json:"StatsID"`

	// Strength An integer that represents the Persona's Strength stat.
	Strength int `json:"Strength" validate:"min=1,max=99"`
}

// BadRequest defines model for BadRequest.
type BadRequest = ErrorBaseResponse

// Forbidden defines model for Forbidden.
type Forbidden = ErrorBaseResponse

// MissingSubject defines model for MissingSubject.
type MissingSubject = ErrorBaseResponse

// NoContent defines model for NoContent.
type NoContent = map[string]interface{}

// NotFound defines model for NotFound.
type NotFound = ErrorBaseResponse

// ServerError defines model for ServerError.
type ServerError = ErrorBaseResponse

// Unauthorized defines model for Unauthorized.
type Unauthorized = ErrorBaseResponse

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Foo request
	Foo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) Foo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFooRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewFooRequest generates requests for Foo
func NewFooRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/foo")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// Foo request
	FooWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FooResponse, error)
}

type FooResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r FooResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FooResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FooWithResponse request returning *FooResponse
func (c *ClientWithResponses) FooWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FooResponse, error) {
	rsp, err := c.Foo(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFooResponse(rsp)
}

// ParseFooResponse parses an HTTP response from a FooWithResponse call
func ParseFooResponse(rsp *http.Response) (*FooResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FooResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /foo)
	Foo(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Foo converts echo context to params.
func (w *ServerInterfaceWrapper) Foo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Foo(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/foo", wrapper.Foo)

}

type BadRequestJSONResponse ErrorBaseResponse

type ForbiddenJSONResponse ErrorBaseResponse

type MissingSubjectJSONResponse ErrorBaseResponse

type NoContentJSONResponse map[string]interface{}

type NotFoundJSONResponse ErrorBaseResponse

type ServerErrorJSONResponse ErrorBaseResponse

type UnauthorizedJSONResponse ErrorBaseResponse

type FooRequestObject struct {
}

type FooResponseObject interface {
	VisitFooResponse(w http.ResponseWriter) error
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /foo)
	Foo(ctx context.Context, request FooRequestObject) (FooResponseObject, error)
}

type StrictHandlerFunc func(ctx echo.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// Foo operation middleware
func (sh *strictHandler) Foo(ctx echo.Context) error {
	var request FooRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Foo(ctx.Request().Context(), request.(FooRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Foo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(FooResponseObject); ok {
		return validResponse.VisitFooResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"ksOl6BsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
