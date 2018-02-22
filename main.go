package main

import (
	"fmt"
	"github.com/dmgk/faker"
	"github.com/honeycombio/libhoney-go"
	"os"
	"time"
)

func main() {
	fmt.Println("Key in use THRASHER_HONEYCOMBKEY", os.Getenv("THRASHER_HONEYCOMBKEY"))

	libhoney.Init(libhoney.Config{
		WriteKey: os.Getenv("THRASHER_HONEYCOMBKEY"),
		Dataset:  "logstash-sample-data-mocker-example",
	})

	defer libhoney.Close() // Flush any pending calls to Honeycomb

	go loadHoneycombData(1000000, "load 1")
	go loadHoneycombData(1000000, "load 2")
	go loadHoneycombData(1000000, "load 3")
	go loadHoneycombData(1000000, "load 4")
	go loadHoneycombData(1000000, "load 5")

	fmt.Print("Working... press key to end.\n")
	fmt.Scanln()
	fmt.Println("Finished.\n")
}

func loadHoneycombData(repeat int, ident string) {
	fmt.Print("started ping " + ident + "\n")
	for i := 0; i < repeat; i++ {
		libhoney.SendNow(map[string]interface{}{
			"time":                      time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00"),
			"@version":                  "1",
			"actconn":                   faker.Number().Between(100, 600),
			"backend_name":              "645_trapster_" + ident,
			"backend_queue":             faker.Number().Between(0, 24),
			"beconn":                    "70",
			"bytes_read":                144,
			"captured_request_cookie":   "-",
			"captured_response_cookie":  "-",
			"client_ip":                 faker.Internet().IpV4Address(),
			"client_port":               faker.Number().Between(50124, 50127),
			"facility":                  faker.Number().Between(3, 20),
			"facility_label":            "local0",
			"feconn":                    "471",
			"frontend_name":             "default_" + ident,
			"host":                      "127.0.0.1",
			"hostname":                  "ip-10-1-47-121",
			"http_request":              "/v2/library/ubuntu/manifests/14.04",
			"http_status_code":          "500",
			"http_verb":                 "GET",
			"http_version":              "1.1",
			"logsource":                 "ip-10-1-47-121",
			"pid":                       faker.Number().Between(1111, 8888),
			"priority":                  134,
			"program":                   "haproxy",
			"request_header_host":       "registry-1.docker.io",
			"request_header_user_agent": "docker/17.03.2-ce go/go1.7.5 git-commit/f5ec1e2 kernel/4.4.0-75-generic os/linux arch/amd64 UpstreamClient(Docker-Client/17.03.2",
			"retries":                   faker.Number().Between(1, 3),
			"server_name":               "\u003cNOSRV\u003e",
			"service":                   "haproxy",
			"severity":                  faker.Number().Between(0, 7),
			"severity_label":            "Informational",
			"source_host":               "ip-10-1-47-121",
			"srv_queue":                 "0",
			"srvconn":                   "0",
			"tags":                      "[\"haproxy\",\"" + faker.Hacker().Adjective() + "\",\"" + faker.Hacker().Noun() + "\"]",
			"termination_state":         "PT--",
			"time_backend_connect":      -1,
			"time_backend_response":     -1,
			"time_duration":             2069,
			"time_queue":                2070,
			"time_request":              -1,
			"timestamp":                 time.Now().UTC().Format("Jan _2 15:04:05"),
			"type":                      "haproxy",
		})
	}
	fmt.Print("done ping " + ident + "\n")
}
