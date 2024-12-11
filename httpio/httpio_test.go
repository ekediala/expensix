package httpio_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ekediala/expensix/httpio"
	"github.com/ekediala/expensix/templ/pages/health"
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
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	c := health.Health()
	buf := strings.Builder{}
	err := c.Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("expected nil, got %q", err)
	}

	httpio.Code(http.StatusTeapot, httpio.HTML(c, httpio.OK, httpio.Header{Value: httpio.AllowedMaxAge, Key: "Access-Control-Max-Age"})).ServeHTTP(w, r)

	if got := w.Code; w.Code != http.StatusTeapot {
		t.Errorf("expected w.Code to be %d, got %d", http.StatusTeapot, got)
	}

	if got := w.Header().Get("Access-Control-Max-Age"); got != httpio.AllowedMaxAge {
		t.Errorf("expected Access-Control-Max-Age to be %q, got %q", httpio.AllowedMaxAge, got)
	}

	if got := w.Body.String(); !strings.Contains(got, buf.String()) {
		t.Errorf("expected body to contain %s, got %s", buf.String(), got)
	}

}
