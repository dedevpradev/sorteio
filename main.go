package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var fp = flag.String("f", "items.txt", "file path")

func main() {
	flag.Parse()

	file, err := os.Open(*fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var names []string
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(names[rand.Intn(len(names))])
}
