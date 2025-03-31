package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {

	tests := []struct {
		name    string
		a       int
		b       int
		checkFn func(t *testing.T, a, b, total int)
	}{
		{
			name: "First Case",
			a:    10,
			b:    20,
			checkFn: func(t *testing.T, a, b, total int) {
				assert.Equal(t, a+b, total)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Addition(tt.a, tt.b)
			tt.checkFn(t, tt.a, tt.b, got)
		})
	}
}
