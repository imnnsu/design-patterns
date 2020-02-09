package singleton

type singleton struct{}

var uniqueInstance *singleton

func newSingleton() *singleton {
	return &singleton{}
}

func GetInstance() *singleton {
	if uniqueInstance == nil {
		uniqueInstance = newSingleton()
	}
	return uniqueInstance
}
