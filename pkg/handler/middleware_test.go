package handler

import (
	"github.com/atadzan/todo-app/pkg/service"
	mock_service "github.com/atadzan/todo-app/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		headerName          string
		headerValue         string
		token               string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "1",
		},
		{
			name:                "Invalid Header Name",
			headerName:          "",
			headerValue:         "Bearer token",
			token:               "token",
			mockBehaviour:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"empty auth header"}`,
		},
		{
			name:                "Invalid Header Value",
			headerName:          "Authorization",
			headerValue:         "Bearr token",
			token:               "token",
			mockBehaviour:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                "Empty token",
			headerName:          "Authorization",
			headerValue:         "Bearer",
			mockBehaviour:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"invalid auth header"}`,
		},
		{
			name:        "PArse error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
			},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"invalid token"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//	Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := Handler{services}

			//	Test Server
			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, "%d", id)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			//Make request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}

}

func TestGetUserId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}
	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name: "OK",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getUserId(test.ctx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, id, test.id)
		})
	}
}
