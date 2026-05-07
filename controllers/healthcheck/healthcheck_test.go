package healthcheck

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/logan1o1/go_chess_server/mocks"
	"github.com/logan1o1/go_chess_server/models"
)

func TestCheckServerHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		mockResp   models.HealthCheckResponse
		mockErr    error
		wantStatus int
	}{
		{
			name:       "success",
			mockResp:   models.HealthCheckResponse{Status: "ok"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "service error",
			mockErr:    errors.New("db error"),
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.MockHealthCheckService{
				Response: tt.mockResp,
				Err:      tt.mockErr,
			}
			controller := NewHealthCkeckController(mockService)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
			controller.CheckServerHealthCheck(ctx)
			if w.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, w.Code)
			}
		})
	}
}
