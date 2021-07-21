package gin_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleMetrics(t *testing.T) {
	w := setupRouterAndGetResponse(t, models.Metrics, commons.MetricsEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var metrics *models.MemStats
	err := json.NewDecoder(w.Body).Decode(&metrics)
	assert.NoError(t, err)
	assert.NotNil(t, metrics)
}
