package server

import (
	"fmt"
	"net/http"

	"github.com/1gkx/openmetrics/internal/router"
)

func Start(port string) error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		router.InitRouter(),
		)
}
