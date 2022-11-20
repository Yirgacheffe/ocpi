package versions

import (
	"encoding/json"
	"log"
	"net/http"
)

type VersionHandler struct {
	Repo VersionRepository
}

func NewHandler(repo *PGVersionRepository) *VersionHandler {
	return &VersionHandler{
		Repo: repo,
	}
}

func (v *VersionHandler) ListVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	results, err := v.Repo.FindAll()
	if err != nil {
		log.Printf("not found %s", err)
		return
	}

	j, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}
