package ipfinder

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockHTTPClient struct {
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.GetFunc(url)
}

func TestIPFinder_GetIP(t *testing.T) {
	tests := []struct {
		name        string
		mockGetFunc func(url string) (*http.Response, error)
		want        string
		wantErr     bool
	}{
		{
			name: "Success",
			mockGetFunc: func(url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("192.168.1.1")),
				}, nil
			},
			want:    "192.168.1.1",
			wantErr: false,
		},
		{
			name: "Error case - bad status code",
			mockGetFunc: func(url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(strings.NewReader("")),
				}, nil
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Error case - network error",
			mockGetFunc: func(url string) (*http.Response, error) {
				return nil, errors.New("network error")
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				GetFunc: tt.mockGetFunc,
			}

			finder := &IPFinder{
				URL:    "http://example.com",
				Client: mockClient,
			}

			got, err := finder.GetIP()

			if (err != nil) != tt.wantErr {
				t.Errorf("IPFinder.GetIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IPFinder.GetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
