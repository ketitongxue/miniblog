// Copyright 2025 JuZX <wo_sakura@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ketitongxue/miniblog.

package errno

import (
	"net/http"
	"testing"

	"github.com/onexstack/onexstack/pkg/errorsx"
)

func TestErrorCodes(t *testing.T) {
	tests := []struct {
		name     string
		err      *errorsx.ErrorX
		wantCode int
		wantMsg  string
	}{
		{
			name:     "OK",
			err:      OK,
			wantCode: http.StatusOK,
			wantMsg:  "",
		},
		{
			name:     "ErrPageNotFound",
			err:      ErrPageNotFound,
			wantCode: http.StatusNotFound,
			wantMsg:  "Page not found.",
		},
		{
			name:     "ErrSignToken",
			err:      ErrSignToken,
			wantCode: http.StatusUnauthorized,
			wantMsg:  "Error occurred while signing the JSON web token.",
		},
		{
			name:     "ErrTokenInvalid",
			err:      ErrTokenInvalid,
			wantCode: http.StatusUnauthorized,
			wantMsg:  "Token was invalid.",
		},
		{
			name:     "ErrDBRead",
			err:      ErrDBRead,
			wantCode: http.StatusInternalServerError,
			wantMsg:  "Database read failure.",
		},
		{
			name:     "ErrDBWrite",
			err:      ErrDBWrite,
			wantCode: http.StatusInternalServerError,
			wantMsg:  "Database write failure.",
		},
		{
			name:     "ErrAddRole",
			err:      ErrAddRole,
			wantCode: http.StatusInternalServerError,
			wantMsg:  "Error occurred while adding the role.",
		},
		{
			name:     "ErrRemoveRole",
			err:      ErrRemoveRole,
			wantCode: http.StatusInternalServerError,
			wantMsg:  "Error occurred while removing the role.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Code != tt.wantCode {
				t.Errorf("%s: got code %d, want %d", tt.name, tt.err.Code, tt.wantCode)
			}
			if tt.err.Message != tt.wantMsg {
				t.Errorf("%s: got message %q, want %q", tt.name, tt.err.Message, tt.wantMsg)
			}
		})
	}
}
func TestErrorCodesNotNil(t *testing.T) {
	errorCodes := []*errorsx.ErrorX{
		OK,
		ErrInternal,
		ErrNotFound,
		ErrBind,
		ErrInvalidArgument,
		ErrUnauthenticated,
		ErrPermissionDenied,
		ErrOperationFailed,
		ErrPageNotFound,
		ErrSignToken,
		ErrTokenInvalid,
		ErrDBRead,
		ErrDBWrite,
		ErrAddRole,
		ErrRemoveRole,
	}

	for i, err := range errorCodes {
		if err == nil {
			t.Errorf("Error code at index %d is nil", i)
		}
	}
}

func TestErrorReason(t *testing.T) {
	tests := []struct {
		name       string
		err        *errorsx.ErrorX
		wantReason string
	}{
		{
			name:       "ErrPageNotFound",
			err:        ErrPageNotFound,
			wantReason: "NotFound.PageNotFound",
		},
		{
			name:       "ErrSignToken",
			err:        ErrSignToken,
			wantReason: "Unauthenticated.SignToken",
		},
		{
			name:       "ErrTokenInvalid",
			err:        ErrTokenInvalid,
			wantReason: "Unauthenticated.TokenInvalid",
		},
		{
			name:       "ErrDBRead",
			err:        ErrDBRead,
			wantReason: "InternalError.DBRead",
		},
		{
			name:       "ErrDBWrite",
			err:        ErrDBWrite,
			wantReason: "InternalError.DBWrite",
		},
		{
			name:       "ErrAddRole",
			err:        ErrAddRole,
			wantReason: "InternalError.AddRole",
		},
		{
			name:       "ErrRemoveRole",
			err:        ErrRemoveRole,
			wantReason: "InternalError.RemoveRole",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Reason != tt.wantReason {
				t.Errorf("%s: got reason %q, want %q", tt.name, tt.err.Reason, tt.wantReason)
			}
		})
	}
}

func TestHTTPStatusCodes(t *testing.T) {
	tests := []struct {
		name           string
		err            *errorsx.ErrorX
		expectedStatus int
	}{
		{"OK should be 200", OK, http.StatusOK},
		{"ErrPageNotFound should be 404", ErrPageNotFound, http.StatusNotFound},
		{"ErrSignToken should be 401", ErrSignToken, http.StatusUnauthorized},
		{"ErrTokenInvalid should be 401", ErrTokenInvalid, http.StatusUnauthorized},
		{"ErrDBRead should be 500", ErrDBRead, http.StatusInternalServerError},
		{"ErrDBWrite should be 500", ErrDBWrite, http.StatusInternalServerError},
		{"ErrAddRole should be 500", ErrAddRole, http.StatusInternalServerError},
		{"ErrRemoveRole should be 500", ErrRemoveRole, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.err.Code)
			}
		})
	}
}
