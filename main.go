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
	perRoutineCount := 10

	libhoney.Init(libhoney.Config{
		WriteKey: os.Getenv("THRASHER_HONEYCOMBKEY"),
		Dataset:  "sample-data-comparison",
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
		libhoney.SendNow(makeMap())
		fmt.Printf("Finished %s", ident)
	}
}

func makeMap(ident string) map[string]interface{} {
	return map[string]interface{}{
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
	}
}