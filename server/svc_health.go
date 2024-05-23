package server

import (
	"net/http"
)

// Health represents current service health status.
type Health struct {
	PostgresPing string `json:"postgres_ping"`
}

func HealthCheck(dp databaseChecker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := Health{
			PostgresPing: "ok",
		}
		if _, err := dp.Ping(); err != nil {
			h.PostgresPing = err.Error()
		}
		encodeJSONResp(w, h, http.StatusOK)
	}
}

type databaseChecker interface {
	Ping() (ok bool, err error)
}
