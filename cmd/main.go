package main

import (
	"fmt"

	"github.com/grizante/jast/linux/sensors"
	"github.com/grizante/jast/storage"
)


func main() {
        storage := storage.NewStorage()
        si := sensors.SensorInfo{}
        si.SaveTemperatures(storage)
        fmt.Printf("%v", storage.Read())
}