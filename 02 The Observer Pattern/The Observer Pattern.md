### Unpacking the WeatherData class

```go
type WeatherData struct {}


func (w WeatherData) GetTemperature() float64 {
    return 0.0
}

func (w WeatherData) GetHumidity() float64 {
    return 0.0
}
    
func (w WeatherData) GetPressure() float64 {
    return 0.0
}

// WeatherData gets called whenever the weather measurements have been updated. 
func (w WeatherData) MeasurementsChanged() {
    // Your code goes here
}
```

### Taking a first, misguided SWAG at the Weather Station

```go
func (w WeatherData) MeasurementsChanged() {
    temp := w.GetTemperature()
    humidity := w.GetHumidity()
    pressure := w.GetPressure()

    currentConditionDisplay.Update(temp, humidity, pressure)
    statisticsDisplay.Update(temp, humidity, pressure)
    forecastDisplay.Update(temp, humidity, pressure)
}
```

> By coding to concrete implementations, we have no way to add or remove other display elements without making changes to the program.

### The Observer Pattern defined

> **The Observer Pattern** defines a one-to-many dependency between objects so that when one object changes state, all of its dependents are notified and updated automatically.

### The power of Loose Coupling

> **Design Principle** Strive for loosely coupled designs between objects that interact.

### Implementing the Weather Station

```go
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
```

### Implementing the Subject interface in WeatherData

```go
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

func(w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
    w.temperature = temperature
    w.humidity = humidity
    w.pressure = pressure
    w.MeasurementsChanged()
}
```

### Now, let's build those display elements

```go
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
    fmt.Printf("Current conditions: %.1fF degrees and %.1f%% humidity\n", c.temperature, c.humidity)
}
```


### Power up the Weather Station

```go
func main() {
	weatherData := &WeatherData{}
	_ = NewCurrentConditionsDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
```