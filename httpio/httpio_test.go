package httpio_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ekediala/expensix/httpio"
)

func TestTextMiddleware(t *testing.T) {
	t.Parallel()
	s := "hello"

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	httpio.Text(s).ServeHTTP(w, r)

	if got := w.Body.String(); !strings.Contains(got, s) {
		t.Errorf("expected %q to contain %q", got, s)
	}
}

func TestCodeMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	httpio.Code(http.StatusBadGateway, httpio.OK).ServeHTTP(w, r)

	if got := w.Code; got != http.StatusBadGateway {
		t.Errorf("expected code to be %d, got %d", http.StatusBadGateway, got)
	}
}

func TestHTMLMiddleware(t *testing.T) {
	
}
