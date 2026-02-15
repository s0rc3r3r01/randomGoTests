package getstatus

import (
	"net/http"
	"net/http/httptest"
	"randomGoTests/httpgordle/internal/api"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/games/", nil)
	require.NoError(t, err)
	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()
	Handle(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"solution":"",
"wordlength":0,"status":""}`, recorder.Body.String())
}
