package service

import (
	"fmt"
	"testing"
)

func TestCreateUID(t *testing.T) {

	for i := 0; i < 100; i++ {
		if id, err := UIDWorker.NextId(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}

}
