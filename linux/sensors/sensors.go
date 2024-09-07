package sensors

import (
	"log"

	"github.com/grizante/jast/storage"
	"github.com/shirou/gopsutil/v4/sensors"
)

type SensorInfo struct {
	Name string
	Temperature float64
}

func (si *SensorInfo) writeTemperatures(data []sensors.TemperatureStat) *[]SensorInfo {
        var sensors []SensorInfo = make([]SensorInfo, len(data))
        for i := 0; i < len(sensors); i++ {
                sensors[i] = SensorInfo{Name: data[i].SensorKey, Temperature: data[i].Temperature}
        }
        return &sensors
}

func (si *SensorInfo) readTemperatures() []sensors.TemperatureStat {
        sensors, err := sensors.SensorsTemperatures()
        if err != nil {
                log.Fatal(err)
        }
        return sensors
}

func (si *SensorInfo) SaveTemperatures(storage *storage.Storage) {
        data := si.readTemperatures()
        sensorsInfo := si.writeTemperatures(data)
        for _, sensor := range *sensorsInfo {
                storage.Set(sensor.Name, sensor.Temperature)
        }
}
