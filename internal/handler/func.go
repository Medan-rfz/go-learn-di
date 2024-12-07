package handler

import "net/http"

// Func some func
// @Summary some func
// @Tags funcs
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /func [get]
func (h *Handler) Func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := h.repo.Func(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("some data"))
	}
}
