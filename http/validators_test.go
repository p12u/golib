package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	z "github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `json:"name"`
}

func (u User) Validator() Parser {
	return z.Struct(z.Schema{
		"name": z.String().Len(8),
	})
}

type Query struct {
	StartDate time.Time `json:"startDate"`
}

func (q Query) Validator() Parser {
	return z.Struct(z.Schema{
		"startDate": z.Time().
			Required(z.Message("startDate is required")).
			PreTransform(func(data any, ctx z.ParseCtx) (out any, err error) {
				str, ok := data.(string)
				if !ok {
					return nil, errors.New("not a string")
				}

				return time.Parse(time.DateOnly, str)
			}),
	})
}

func TestValidated(t *testing.T) {
	userJson := `{"name":"Jon Snow","email":"jon@labstack.com"}`
	responseJson := `{"name":"Jon Snow"}`
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/?startDate=2025-05-01",
		strings.NewReader(userJson),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := EBQ(func(c echo.Context, u *User, q *Query) error {
		assert.Equal(t, 2025, q.StartDate.Year())
		return c.JSON(http.StatusOK, u)
	})

	// Assertions
	err := handler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseJson, strings.Trim(rec.Body.String(), "\n"))
	}

}
