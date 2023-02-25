package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

type output struct {
	name        string
	id          string
	count       int
	point       float64
	temperature float64
	typ         string
}

func (o output) String4CSV() []string {
	return []string{
		o.name,
		o.id,
		fmt.Sprintf("%v", o.count),
		fmt.Sprintf("%v", o.point),
		fmt.Sprintf("%v", o.temperature),
		o.typ,
	}
}

func newOuntout() output {
	return output{
		name:        uuid.NewString(),
		id:          uuid.NewString(),
		count:       rand.Intn(100),
		point:       rand.Float64() * float64(rand.Intn(10000)),
		temperature: rand.Float64() * float64(rand.Intn(10000)),
		typ:         m[rand.Intn(9)],
	}
}

var m = map[int]string{
	0: "Go",
	1: "C",
	2: "Scala",
	3: "Python",
	4: "Java",
	5: "JavaScript",
	6: "TypeScript",
	7: "Ruby",
	8: "Rust",
	9: "Dart",
}

func main() {
	f, err := os.OpenFile("sample.csv", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	rand.Intn(10)
	w := csv.NewWriter(f)
	for i := 0; i < 1000; i++ {
		w.Write(newOuntout().String4CSV())
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
