package utils

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestParseBody(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name       string
		body       string
		want       TestStruct
		shouldFail bool
	}{
		{
			name: "Valid JSON",
			body: `{"name": "John", "age": 30}`,
			want: TestStruct{Name: "John", Age: 30},
		},
		{
			name:       "Invalid JSON",
			body:       `{"name": "John", "age": "thirty"}`,
			want:       TestStruct{},
			shouldFail: true,
		},
		{
			name:       "Empty Body",
			body:       ``,
			want:       TestStruct{},
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tt.body))
			var got TestStruct
			ParseBody(req, &got)

			if !tt.shouldFail && (got != tt.want) {
				t.Errorf("ParseBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
