package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	URL = "/api/heartbeat"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

type Handler struct {
}

// Register adds the routes for the metric handler to the passed router.
func (h *Handler) Register(router *gin.Engine) {
	router.Any(URL, func(c *gin.Context) {
		HandlerFunc(h.Heartbeat).ServeHTTP(c.Writer, c.Request)
	})
}

// Heartbeat handles the heartbeat metric.
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
