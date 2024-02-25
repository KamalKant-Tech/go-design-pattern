package db

// singleton pattern for DB client
// Link Referral: https://refactoring.guru/design-patterns/singleton
import (
	"fmt"
	"sync"
)

type Db struct {
}

var dbInstance *Db

var lock = &sync.Mutex{}

func GetInstance() *Db {
	if dbInstance != nil {
		fmt.Println("Instance already created!!!")
	} else {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			dbInstance = &Db{}
			fmt.Println("Creating the instance now")
		} else {
			fmt.Println("Instance already created from goroutines!!!")
		}
	}
	return dbInstance
}
