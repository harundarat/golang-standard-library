package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Harun", "Al", "Rasyid"})
	_ = writer.Write([]string{"Vincent", "van", "Googh"})
	_ = writer.Write([]string{"Radit", "van", "Adit"})

	writer.Flush()
}
