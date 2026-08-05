package shared

import "net/http"

type HealthCheck struct {
	Status string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorWithStatus(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method tidak diizinkan")
		return
	}

	Success(w, HealthCheck{Status: "OK"})
}
