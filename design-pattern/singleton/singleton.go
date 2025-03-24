package singleton

import "sync"

type Singleton struct {
	Id   int
	Name string
}

var instance *Singleton
var once sync.Once

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Id:   1,
			Name: "Rajesh Kumar",
		}
	})

	return instance
}
