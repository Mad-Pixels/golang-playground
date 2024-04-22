package entrypoint

import (
	"encoding/json"
	"net/http"

	"github.com/Mad-Pixels/golang-playground/apps"
	"github.com/Mad-Pixels/golang-playground/apps/pkg/k8s"
)

type probeResponse struct {
	Host    string `json:"host"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

func handlerLivenessProbe(w http.ResponseWriter, _ *http.Request) {
	response := probeResponse{
		Status:  "OK",
		Message: "LivenessProbe",
		Host:    apps.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func handlerReadinessProbe(w http.ResponseWriter, _ *http.Request) {
	response := probeResponse{
		Status:  "OK",
		Message: "ReadnessProbe",
		Host:    apps.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func playground(w http.ResponseWriter, r *http.Request) {
	k8s.Pod()
	w.Write([]byte("welcome"))
}
