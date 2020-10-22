/*
 * rabbitmq-quacker
 *
 *
 * Contact: zhangwb@shinetechchina.com
 */

package main

import (
	"fmt"
	"rabbitmq-quacker/app"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Server started\n")

	rabbitmqConfig := app.QuackerConfig{
		Host:     getEnv("QUACKER_HOST", "127.0.0.1"), // "rabbitmq.osvie.com",
		Port:     getEnv("QUACKER_PORT", "1883"),
		Username: getEnv("QUACKER_USERNAME", ""),
		Password: getEnv("QUACKER_PASSWORD", ""),
		Exchange: getEnvOrFail("QUACKER_EXCHANGE"),
		Topic:    getEnvOrFail("QUACKER_TOPIC"),
		Interval: getEnv("QUACKER_INTERVAL", "1000"),
		DataFile: getEnv("QUACKER_DATAFILE", "/data.json"),
	}

	quacker := app.NewQuacker(rabbitmqConfig)
	runtime.SetFinalizer(&quacker, func(obj *app.Quacker) {
		obj.Close()
	})

	quacker.Start()
	defer quacker.Close()
}

func getEnv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found == false {
		return defaultValue
	}
	return value
}

func getEnvOrFail(key string) string {
	value, found := os.LookupEnv(key)
	if found == false {
		panic("Need env key" + key)
	}
	return value
}
