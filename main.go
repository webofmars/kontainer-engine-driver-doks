package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/rancher/kontainer-engine/types"
	"github.com/ribeiro-rodrigo/kontainer-engine-driver-doks/doks"
	"github.com/sirupsen/logrus"
)

var wg = &sync.WaitGroup{}

func main() {
	if os.Args[1] == "" {
		panic(errors.New("no port provided"))
	}

	port, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(fmt.Errorf("argument not parsable as int: %v", err))
	}

	addr := make(chan string)
	driver := doks.NewDriver()
	go types.NewServer(&driver, addr).ServeOrDie(fmt.Sprintf("127.0.0.1:%v", port))

	logrus.Debugf("DOKS driver up and running on at %v", <-addr)

	wg.Add(1)
	wg.Wait()
}
