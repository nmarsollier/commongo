package mktools

import (
	"bytes"
	"io"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/security"
	"github.com/nmarsollier/commongo/strs"
	"github.com/nmarsollier/commongo/test/mockgen"
)

// Http Mocks
func ExpectHttpToken(mock *mockgen.MockHTTPClient, user *security.User) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(strs.ToJson(user))),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}

func ExpectHttpUnauthorized(mock *mockgen.MockHTTPClient) {
	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}
