/*
 * amqp-quacker
 *
 *
 * Contact: zhangwb@shinetechchina.com
 */

package main

import (
	"fmt"
	"log"
	"amqp-quacker/app"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Server started\n")

	amqpConfig := app.QuackerConfig{
		Host:     getEnv("QUACKER_HOST", "127.0.0.1"),
		Port:     getEnv("QUACKER_PORT", "5672"),
		Username: getEnv("QUACKER_USERNAME", ""),
		Password: getEnv("QUACKER_PASSWORD", ""),
		Exchange: getEnv("QUACKER_EXCHANGE",""),
		Topic:    getEnvOrFail("QUACKER_TOPIC"),
		Interval: getEnv("QUACKER_INTERVAL", "1000"),
		DataFile: getEnv("QUACKER_DATAFILE", "/data.json"),
	}

	quacker := app.NewQuacker(amqpConfig)
	runtime.SetFinalizer(&quacker, func(obj *app.Quacker) {
		obj.Close()
	})

	err := quacker.Start()
	if(err != nil){
		log.Println(err)
	}
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
