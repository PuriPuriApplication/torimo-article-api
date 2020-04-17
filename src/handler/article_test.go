package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"torimo-article-api/src/handler"
	"torimo-article-api/tests/interactor"
	"torimo-article-api/tests/presenter"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	articleJSON = `"{"title":"test title","body":"test body","status":"public","userId":1,"shopId":1,"categoryIds":[1]}"`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// h := InitArticleHandler()
	h := handler.NewArticleHandler(
		interactor.NewArticleInteractor(),
		presenter.NewArticlePresenter(),
	)

	// Assertions
	if assert.NoError(t, h.CreateArticle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:email")
// 	c.SetParamNames("email")
// 	c.SetParamValues("jon@labstack.com")
// 	h := &handler{mockDB}

// 	// Assertions
// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
