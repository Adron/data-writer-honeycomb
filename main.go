package main

import (
	"fmt"
	"github.com/dmgk/faker"
	"github.com/honeycombio/libhoney-go"
	"os"
	"strconv"
	"time"
)

func main() {
	goRoutineCount := 10
	perRoutineCount := 100000

	libhoney.Init(libhoney.Config{
		WriteKey: os.Getenv("THRASHER_HONEYCOMBKEY"),
		Dataset:  "elastic-comparison",
	})

	defer libhoney.Close()

	fmt.Println("Starting @ " + time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00"))

	for i := 1; i < goRoutineCount; i++ {
		go loadHoneycombData(perRoutineCount, "load_"+strconv.Itoa(i))
	}

	fmt.Print("Working...\n\npress key to end.\n")
	fmt.Scanln()
	fmt.Println("Finished @ " + time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00"))
}

func loadHoneycombData(repeat int, ident string) {
	for i := 0; i < repeat; i++ {

		hostLogSource := faker.Internet().IpV4Address()
		duration, _ := time.ParseDuration(faker.Number().Between(1, 999) + "ms")

		libhoney.SendNow(map[string]interface{}{
			"time":                      time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00"),
			"@version":                  "1",
			"actconn":                   faker.Number().Between(100, 600),
			"backend_name":              faker.Number().Between(100, 125) + "_trapster_" + ident,
			"backend_queue":             faker.Number().Between(0, 24),
			"beconn":                    "70",
			"bytes_read":                faker.Number().Between(64, 1024),
			"captured_request_cookie":   "-",
			"captured_response_cookie":  "-",
			"client_ip":                 faker.Internet().IpV4Address(),
			"client_port":               faker.Number().Between(50124, 50127),
			"facility":                  faker.Number().Between(3, 20),
			"facility_label":            "local0",
			"feconn":                    faker.Number().Between(420, 471),
			"frontend_name":             "default_" + ident,
			"host":                      "127.0.0.1",
			"hostname":                  hostLogSource,
			"http_request":              faker.Internet().Url(),
			"http_status_code":          "500",
			"http_verb":                 "GET",
			"http_version":              "1.1",
			"logsource":                 hostLogSource,
			"pid":                       faker.Number().Between(1, 32768),
			"priority":                  faker.Number().Between(1, 200),
			"program":                   "haproxy",
			"request_header_host":       "registry-x.wherethingsgo.io",
			"request_header_user_agent": faker.Hacker().Phrases(),
			"retries":                   faker.Number().Between(1, 3),
			"server_name":               faker.App().Name(),
			"service":                   "haproxy",
			"severity":                  faker.Number().Between(0, 7),
			"severity_label":            faker.Hacker().Verb(),
			"source_host":               hostLogSource,
			"srv_queue":                 faker.Number().Between(0, 9),
			"srvconn":                   faker.Number().Between(0, 999),
			"tags":                      "[\"haproxy\",\"" + faker.Hacker().Adjective() + "\",\"" + faker.Hacker().Noun() + "\"]",
			"termination_state":         "PT--",
			"time_backend_connect":      -1,
			"time_backend_response":     -1,
			"time_duration":             2069,
			"time_queue":                2070,
			"time_request":              -1,
			"timestamp":                 faker.Time().Backward(duration),
			"type":                      "haproxy",
		})
	}
}
