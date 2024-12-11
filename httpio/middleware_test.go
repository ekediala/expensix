package httpio_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/ekediala/expensix/httpio"
)

func TestTraceMiddleware(t *testing.T) {
	t.Parallel()

	var ok bool

	var h httpio.Handler = func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		_, ok = httpio.GetTraceID(r.Context())
		return httpio.OK
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	httpio.TraceMiddleware(h).ServeHTTP(w, r)

	if !ok {
		t.Error("expected to get traceID, got false")
	}
}

func TestLoggingMiddleware(t *testing.T) {
	buf := strings.Builder{}
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	slog.SetDefault(logger)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	var handler httpio.Handler = func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		return httpio.Code(http.StatusOK, httpio.OK)
	}

	httpio.LoggingMiddleware(handler).ServeHTTP(w, r)

	if got := buf.String(); !strings.Contains(got, strconv.Itoa(http.StatusOK)) {
		t.Errorf("expected status code to be logged, got %s", got)
	}
}

func TestCORSMiddleware(t *testing.T) {
	t.Parallel()
	s := "hello"

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		httpio.Code(http.StatusCreated, httpio.Text(s)).ServeHTTP(w, r)
	}

	testHeaders := func(t *testing.T, w http.ResponseWriter) {
		t.Helper()
		if h := w.Header().Get("Access-Control-Allow-Origin"); h != "" {
			t.Errorf("expected Access-Control-Allow-Origin to be empty, got %q", h)
		}

		if h := w.Header().Get("Access-Control-Allow-Methods"); h != httpio.AllowedMethods {
			t.Errorf("expected Access-Control-Allow-Methods to be %q, got %q", httpio.AllowedMethods, h)
		}

		if h := w.Header().Get("Access-Control-Allow-Headers"); h != httpio.AllowedHeaders {
			t.Errorf("expected Access-Control-Allow-Headers to be %q, got %q", httpio.AllowedHeaders, h)
		}

		if h := w.Header().Get("Access-Control-Max-Age"); h != httpio.AllowedMaxAge {
			t.Errorf("expected Access-Control-Max-Age to be %q, got %q", httpio.AllowedMaxAge, h)
		}
	}

	t.Run("calls next handler", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		httpio.CORSMiddleware(handler).ServeHTTP(w, r)

		if got := w.Code; w.Code != http.StatusCreated {
			t.Errorf("expected status code %d got %d", http.StatusCreated, got)
		}

		if h := w.Header().Get("Content-Type"); h != "text/plain; charset=utf-8" {
			t.Errorf("expected Content-Type to be text/plain; charset=utf-8, got %q", h)
		}

		if got := w.Body.String(); !strings.Contains(got, s) {
			t.Errorf("expected %s to contain %q", got, s)
		}

		testHeaders(t, w)
	})

	t.Run("options returns ok", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodOptions, "/", nil)

		httpio.CORSMiddleware(handler).ServeHTTP(w, r)

		if got := w.Code; w.Code != http.StatusOK {
			t.Errorf("expected status code %d got %d", http.StatusOK, got)
		}
		testHeaders(t, w)
	})

}
