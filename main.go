package main

import (
	"api/api/src/config"
	"api/api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func initNewRelic() {
	_, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func main() {
	config.Carregar()
	fmt.Printf("Escutando na porta %d", config.Porta)
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
