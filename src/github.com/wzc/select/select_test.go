package selectrace

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestSynchroniseRacer(t *testing.T) {
	serverA := makeDelayServer(3 * time.Second)
	serverB := makeDelayServer(4 * time.Second)

	//defer代表在serverA执行完成后执行
	defer serverA.Close()
	defer serverB.Close()

	_, err := SynchroniseRacer(serverA.URL, serverB.URL, 3*time.Second)

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}
