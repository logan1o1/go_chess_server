package healthcheck_service

import (
	"github.com/logan1o1/go_chess_server/mocks"
	"github.com/logan1o1/go_chess_server/models"
	"testing"
)

func TestCheckHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		mockDB  *mocks.MockDatabase
		want    models.HealthCheckResponse
		wantErr bool
	}{
		{
			name:    "successful health check",
			mockDB:  &mocks.MockDatabase{},
			want:    models.HealthCheckResponse{Status: "ok"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewHealthCheckService(tt.mockDB)
			got, err := service.CheckHealthCheck()
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got.Status != tt.want.Status {
				t.Errorf("expected status '%s', got '%s'", tt.want.Status, got.Status)
			}
		})
	}
}
