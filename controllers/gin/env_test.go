package gin_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/sinhashubham95/go-actuator/commons"
)

func TestHandleEnv(t *testing.T) {
	w := setupRouterAndGetResponse(t, models.Env, commons.EnvEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var env map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&env)
	assert.NoError(t, err)
	assert.NotEmpty(t, env)
}
