package main

import (
	"os"

	"github.com/andresidrim/cesupa-hospital/env"
)

func main() {
	os.Remove(env.DB_URL)
}
