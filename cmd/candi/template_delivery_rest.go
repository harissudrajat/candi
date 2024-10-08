package main

import (
	"strings"

	"github.com/golangid/candi/candihelper"
)

const (
	deliveryRestTemplate = `// {{.Header}}

package resthandler

import (
	"encoding/json"
	"io"
	"net/http"{{if and .MongoDeps (not .SQLDeps)}}{{else}}
	"strconv"{{end}}

	"{{$.PackagePrefix}}/internal/modules/{{cleanPathModule .ModuleName}}/domain"
	"{{.PackagePrefix}}/pkg/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	restserver "github.com/golangid/candi/codebase/app/rest_server"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(uc usecase.Usecase, deps dependency.Dependency) *RestHandler {
	return &RestHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root interfaces.RESTRouter) {
	v1{{upper (camel .ModuleName)}} := root.Group(candihelper.V1+"/{{kebab .ModuleName}}", h.mw.HTTPBearerAuth)

	v1{{upper (camel .ModuleName)}}.GET("/", h.getAll{{upper (camel .ModuleName)}}, h.mw.HTTPPermissionACL("getAll{{upper (camel .ModuleName)}}"))
	v1{{upper (camel .ModuleName)}}.GET("/:id", h.getDetail{{upper (camel .ModuleName)}}ByID, h.mw.HTTPPermissionACL("getDetail{{upper (camel .ModuleName)}}"))
	v1{{upper (camel .ModuleName)}}.POST("/", h.create{{upper (camel .ModuleName)}}, h.mw.HTTPPermissionACL("create{{upper (camel .ModuleName)}}"))
	v1{{upper (camel .ModuleName)}}.PUT("/:id", h.update{{upper (camel .ModuleName)}}, h.mw.HTTPPermissionACL("update{{upper (camel .ModuleName)}}"))
	v1{{upper (camel .ModuleName)}}.DELETE("/:id", h.delete{{upper (camel .ModuleName)}}, h.mw.HTTPPermissionACL("delete{{upper (camel .ModuleName)}}"))
}

// GetAll{{upper (camel .ModuleName)}} documentation
// @Summary			Get All {{upper (camel .ModuleName)}}
// @Description		API for get all {{.ModuleName}}
// @Tags			{{upper (camel .ModuleName)}}
// @Accept			json
// @Produce			json
// @Param			page	query	string	false	"Page with default value is 1"
// @Param			limit	query	string	false	"Limit with default value is 10"
// @Param			search	query	string	false	"Search"
// @Param			orderBy	query	string	false	"Order By"
// @Param			sort	query	string	false	"Sort (ASC DESC)"
// @Success			200	{object}	domain.Response{{upper (camel .ModuleName)}}List
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/{{kebab .ModuleName}} [get]
func (h *RestHandler) getAll{{upper (camel .ModuleName)}}(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "{{upper (camel .ModuleName)}}DeliveryREST:GetAll{{upper (camel .ModuleName)}}")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using HTTPBearerAuth in middleware for this handler

	var filter domain.Filter{{upper (camel .ModuleName)}}
	if err := candihelper.ParseFromQueryParam(req.URL.Query(), &filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(rw)
		return
	}

	if err := h.validator.ValidateDocument("{{cleanPathModule .ModuleName}}/get_all", filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(rw)
		return
	}

	result, err := h.uc.{{upper (camel .ModuleName)}}().GetAll{{upper (camel .ModuleName)}}(ctx, &filter)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	message := "Success, with your user id (" + tokenClaim.Subject + ") and role (" + tokenClaim.Role + ")"
	response := wrapper.NewHTTPResponse(http.StatusOK, message, result.Data)
	response.Meta = result.Meta
	response.JSON(rw)
}

// GetDetail{{upper (camel .ModuleName)}} documentation
// @Summary			Get Detail {{upper (camel .ModuleName)}}
// @Description		API for get detail {{.ModuleName}}
// @Tags			{{upper (camel .ModuleName)}}
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.Response{{upper (camel .ModuleName)}}
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/{{kebab .ModuleName}}/{id} [get]
func (h *RestHandler) getDetail{{upper (camel .ModuleName)}}ByID(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "{{upper (camel .ModuleName)}}DeliveryREST:GetDetail{{upper (camel .ModuleName)}}ByID")
	defer trace.Finish()

	{{if and .MongoDeps (not .SQLDeps)}}id := restserver.URLParam(req, "id"){{else}}id, _ := strconv.Atoi(restserver.URLParam(req, "id")){{end}}
	data, err := h.uc.{{upper (camel .ModuleName)}}().GetDetail{{upper (camel .ModuleName)}}(ctx, id)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(rw)
}

// Create{{upper (camel .ModuleName)}} documentation
// @Summary			Create {{upper (camel .ModuleName)}}
// @Description		API for create {{.ModuleName}}
// @Tags			{{upper (camel .ModuleName)}}
// @Accept			json
// @Produce			json
// @Param			data	body	domain.Request{{upper (camel .ModuleName)}}	true	"Body Data"
// @Success			200	{object}	domain.Response{{upper (camel .ModuleName)}}
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/{{kebab .ModuleName}} [post]
func (h *RestHandler) create{{upper (camel .ModuleName)}}(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "{{upper (camel .ModuleName)}}DeliveryREST:Create{{upper (camel .ModuleName)}}")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("{{cleanPathModule .ModuleName}}/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.Request{{upper (camel .ModuleName)}}
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.{{upper (camel .ModuleName)}}().Create{{upper (camel .ModuleName)}}(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success", res).JSON(rw)
}

// Update{{upper (camel .ModuleName)}} documentation
// @Summary			Update {{upper (camel .ModuleName)}}
// @Description		API for update {{.ModuleName}}
// @Tags			{{upper (camel .ModuleName)}}
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Param			data	body	domain.Request{{upper (camel .ModuleName)}}	true	"Body Data"
// @Success			200	{object}	domain.Response{{upper (camel .ModuleName)}}
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/{{kebab .ModuleName}}/{id} [put]
func (h *RestHandler) update{{upper (camel .ModuleName)}}(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "{{upper (camel .ModuleName)}}DeliveryREST:Update{{upper (camel .ModuleName)}}")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("{{cleanPathModule .ModuleName}}/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.Request{{upper (camel .ModuleName)}}
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	{{if and .MongoDeps (not .SQLDeps)}}payload.ID = restserver.URLParam(req, "id"){{else}}payload.ID, _ = strconv.Atoi(restserver.URLParam(req, "id")){{end}}
	err := h.uc.{{upper (camel .ModuleName)}}().Update{{upper (camel .ModuleName)}}(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}

// Delete{{upper (camel .ModuleName)}} documentation
// @Summary			Delete {{upper (camel .ModuleName)}}
// @Description		API for delete {{.ModuleName}}
// @Tags			{{upper (camel .ModuleName)}}
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.Response{{upper (camel .ModuleName)}}
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/{{kebab .ModuleName}}/{id} [delete]
func (h *RestHandler) delete{{upper (camel .ModuleName)}}(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "{{upper (camel .ModuleName)}}DeliveryREST:Delete{{upper (camel .ModuleName)}}")
	defer trace.Finish()
	
	{{if and .MongoDeps (not .SQLDeps)}}id := restserver.URLParam(req, "id"){{else}}id, _ := strconv.Atoi(restserver.URLParam(req, "id")){{end}}
	if err := h.uc.{{upper (camel .ModuleName)}}().Delete{{upper (camel .ModuleName)}}(ctx, id); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}
`

	deliveryRestTestTemplate = `// {{.Header}}

package resthandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"{{$.PackagePrefix}}/internal/modules/{{cleanPathModule .ModuleName}}/domain"
	mockusecase "{{$.PackagePrefix}}/pkg/mocks/modules/{{cleanPathModule .ModuleName}}/usecase"
	mocksharedusecase "{{$.PackagePrefix}}/pkg/mocks/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody                       string
	wantValidateError, wantUsecaseError error
	wantRespCode                        int
}

var (
	errFoo = errors.New("Something error")
)

func TestNewRestHandler(t *testing.T) {
	mockMiddleware := &mockinterfaces.Middleware{}
	mockMiddleware.On("HTTPPermissionACL", mock.Anything).Return(func(http.Handler) http.Handler { return nil })
	mockValidator := &mockinterfaces.Validator{}

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetMiddleware").Return(mockMiddleware)
	mockDeps.On("GetValidator").Return(mockValidator)

	handler := NewRestHandler(nil, mockDeps)
	assert.NotNil(t, handler)

	mockRoute := &mockinterfaces.RESTRouter{}
	mockRoute.On("Group", mock.Anything, mock.Anything).Return(mockRoute)
	mockRoute.On("GET", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("POST", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("PUT", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("DELETE", mock.Anything, mock.Anything, mock.Anything)
	handler.Mount(mockRoute)
}

func TestRestHandler_getAll{{upper (camel .ModuleName)}}(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: "?page=str", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{camel .ModuleName}}Usecase := &mockusecase.{{upper (camel .ModuleName)}}Usecase{}
			{{camel .ModuleName}}Usecase.On("GetAll{{upper (camel .ModuleName)}}", mock.Anything, mock.Anything).Return(
				domain.Response{{upper (camel .ModuleName)}}List{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("{{upper (camel .ModuleName)}}").Return({{camel .ModuleName}}Usecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodGet, "/"+tt.reqBody, strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.getAll{{upper (camel .ModuleName)}}(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_getDetail{{upper (camel .ModuleName)}}ByID(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{camel .ModuleName}}Usecase := &mockusecase.{{upper (camel .ModuleName)}}Usecase{}
			{{camel .ModuleName}}Usecase.On("GetDetail{{upper (camel .ModuleName)}}", mock.Anything, mock.Anything).Return(domain.Response{{upper (camel .ModuleName)}}{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("{{upper (camel .ModuleName)}}").Return({{camel .ModuleName}}Usecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.getDetail{{upper (camel .ModuleName)}}ByID(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_create{{upper (camel .ModuleName)}}(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: ` + "`" + `{"email": "test@test.com"}` + "`" + `, wantUsecaseError: nil, wantRespCode: http.StatusCreated,
		},
		{
			name: "Testcase #2: Negative", reqBody: ` + "`" + `{"email": test@test.com}` + "`" + `, wantUsecaseError: nil, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqBody: ` + "`" + `{"email": "test@test.com"}` + "`" + `, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{camel .ModuleName}}Usecase := &mockusecase.{{upper (camel .ModuleName)}}Usecase{}
			{{camel .ModuleName}}Usecase.On("Create{{upper (camel .ModuleName)}}", mock.Anything, mock.Anything).Return(domain.Response{{upper (camel .ModuleName)}}{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("{{upper (camel .ModuleName)}}").Return({{camel .ModuleName}}Usecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.create{{upper (camel .ModuleName)}}(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_update{{upper (camel .ModuleName)}}(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: ` + "`" + `{"email": "test@test.com"}` + "`" + `, wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: ` + "`" + `{"email": test@test.com}` + "`" + `, wantValidateError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqBody: ` + "`" + `{"email": test@test.com}` + "`" + `, wantUsecaseError: nil, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #4: Negative", reqBody: ` + "`" + `{"email": "test@test.com"}` + "`" + `, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{camel .ModuleName}}Usecase := &mockusecase.{{upper (camel .ModuleName)}}Usecase{}
			{{camel .ModuleName}}Usecase.On("Update{{upper (camel .ModuleName)}}", mock.Anything, mock.Anything, mock.Anything).Return(tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("{{upper (camel .ModuleName)}}").Return({{camel .ModuleName}}Usecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.update{{upper (camel .ModuleName)}}(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_delete{{upper (camel .ModuleName)}}(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{camel .ModuleName}}Usecase := &mockusecase.{{upper (camel .ModuleName)}}Usecase{}
			{{camel .ModuleName}}Usecase.On("Delete{{upper (camel .ModuleName)}}", mock.Anything, mock.Anything).Return(tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("{{upper (camel .ModuleName)}}").Return({{camel .ModuleName}}Usecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.delete{{upper (camel .ModuleName)}}(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}
`
)

func getRestFuncTemplate(moduleName, usecaseName string) string {
	moduleName, usecaseName = strings.Title(moduleName), strings.Title(usecaseName)
	return `func (h *RestHandler) ` + candihelper.ToCamelCase(usecaseName) + `(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "` + moduleName + `DeliveryREST:` + usecaseName + `")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	var payload domain.Request` + usecaseName + `
	// if err := h.validator.ValidateDocument("` + strings.ToLower(moduleName) + `/` + candihelper.ToDelimited(usecaseName, '-') + `", body); err != nil {
	// 	wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
	// 	return
	// }
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.` + moduleName + `().` + usecaseName + `(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", res).JSON(rw)
}
`
}
