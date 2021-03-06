package response

import (
	"net/http"

	"api-billing/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func HTTPErrorHandler(err error, c echo.Context) {
	// handler *echo.HTTPError
	if he, ok := err.(*echo.HTTPError); ok {
		var msgErr string
		msgErr, ok := he.Message.(string)
		if !ok {
			msgErr = "¡Upps! algo inesperado ocurrió"

		}

		registryLogError(c, he.Code, err)
		err = c.JSON(he.Code, Response{Message: msgErr})

		return
	}

	// handler if NOT is a *model.Error
	e, ok := err.(*model.Error)
	if !ok {
		registryLogError(c, http.StatusInternalServerError, err)
		err = c.JSON(http.StatusInternalServerError, MessageResponse{
			Errors: []*Response{{Code: UnexpectedError, Message: "¡Upps! algo inesperado ocurrió"}},
		})

		return
	}

	// handler a *model.Error
	status, resp := getResponseError(e)
	registryLogError(c, status, err)

	err = c.JSON(status, resp)
}

func registryLogError(c echo.Context, status int, err error) {
	fields := logrus.Fields{
		"status":      status,
		"uri":         c.Path(),
		"query_param": c.QueryParams(),
		"remote_ip":   c.RealIP(),
		"method":      c.Request().Method,
	}

	if hasToken(c.Request()) {
		fields["user"] = GetUserID(c)
	}

	if e, ok := err.(*model.Error); ok {
		fields["where"] = e.Where()
		fields["who"] = e.Who()
	}

	if status >= 500 {
		logrus.WithFields(fields).Error(err)
		return
	}

	if status >= 400 {
		logrus.WithFields(fields).Warn(err)
		return
	}
}

func GetUserID(c echo.Context) uint {
	userID := c.Get("userID")
	if userID != nil {
		return userID.(uint)
	}
	return 0
}

func hasToken(r *http.Request) bool {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return false
	}
	return true
}

// getResponseError returns the status code and a Response
func getResponseError(err *model.Error) (outputStatus int, outputResponse MessageResponse) {
	if !err.HasCode() {
		err.SetCode(UnexpectedError)
	}

	if !err.HasAPIMessage() {
		err.SetErrorAsAPIMessage()
	}

	if err.HasData() {
		outputResponse.Data = err.Data()
	}

	// search if in the response map exists the err.Code
	if !err.IsFailureError() {
		if response, ok := responses[err.Code()]; ok {
			err.SetStatus(response.Status)
			err.SetAPIMessage(response.Message)
		}
	}

	// Failure is a special code for send the custom error message, and API message from the logic and not from responses map
	if err.IsFailureError() {
		if !err.HasStatus() {
			err.SetStatus(http.StatusBadRequest)
		}
	}

	if !err.HasStatus() {
		err.SetStatus(http.StatusInternalServerError)
	}

	outputStatus = err.Status()
	outputResponse.Errors = append(outputResponse.Errors, &Response{
		Code:    err.Code(),
		Message: err.APIMessage(),
	})

	return
}
