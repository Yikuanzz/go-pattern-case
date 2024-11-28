package profile

import (
	"log"
	"time"
)

func Duration(invoance time.Time, name string) {
	elapsed := time.Since(invoance)
	log.Printf("%s took %s", name, elapsed)
}
