package handler

import (
	"bytes"
	"github.com/atadzan/todo-app"
	"github.com/atadzan/todo-app/pkg/service"
	mock_service "github.com/atadzan/todo-app/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createItem(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTodoItem, item todo.TodoItem)

	testTable := []struct {
		name                  string
		userId                int
		inputBody             string
		inputItem             todo.TodoItem
		mockBehaviour         mockBehaviour
		expectedStatusCode    int
		expectedRequestedBody string
	}{
		{
			name:      "OK",
			userId:    1,
			inputBody: `{"Title":"Item-1", "Description-1": "Item Description"}`,
			inputItem: todo.TodoItem{
				Title:       "Item-1",
				Description: "Item Description",
			},
			mockBehaviour: func(s *mock_service.MockTodoItem, item todo.TodoItem) {
				s.EXPECT().Create(1, item)
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			//	Test Server
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			//	Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
