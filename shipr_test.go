package shipr

import "log"

func init() {
	if err := Connect("db"); err != nil {
		log.Fatal(err)
	}
}
