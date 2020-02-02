package main

import (
	"fmt"
)

type Subject interface {
    RegisterObserver(o Observer)
    RemoveObserver(o Observer)
    NotifyObservers()
}

type Observer interface {
    Update(temp, humidity, pressure float64)
}

type DisplayElement interface {
    Display()
}

type WeatherData struct {
    observers []Observer
    temperature float64
    humidity float64
    pressure float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
    w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
    for i, observer := range w.observers {
        if o == observer {
            w.observers = append(w.observers[:i], w.observers[i+1:]...)
            return
        }
    }
}

func (w *WeatherData) NotifyObservers() {
    for _, o := range w.observers {
        o.Update(w.temperature, w.humidity, w.pressure)
    }
}

func (w *WeatherData) MeasurementsChanged() {
    w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
    w.temperature = temperature
    w.humidity = humidity
    w.pressure = pressure
    w.MeasurementsChanged()
}

type CurrentConditionsDisplay struct {
    temperature float64
    humidity float64
    weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
    currentConditionsDisplay := CurrentConditionsDisplay{weatherData: weatherData}
    weatherData.RegisterObserver(&currentConditionsDisplay)
    return &currentConditionsDisplay
}

func (c *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
    c.temperature = temperature
    c.humidity = humidity
    c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
    fmt.Printf("Current conditions: %.2fF degrees and %.2f%% humidity\n", c.temperature, c.humidity)
}

func main() {
	weatherData := &WeatherData{}
	_ = NewCurrentConditionsDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
