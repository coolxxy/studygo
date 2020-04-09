package golang

import (
	"testing"
)

var once = Instance()

func TestInstance(t *testing.T) {
	tests := []struct {
		name string
		want *singleton
	}{
		{
			"first",
			once,
		},
		{
			"second",
			once,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Instance(); got != once {
				t.Errorf("Instance() = %v, want %v", got, tt.want)
			}
		})
	}
}
