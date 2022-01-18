package uuid

import uuid "github.com/satori/go.uuid"

func CreateUUID() string {
	u4 := uuid.NewV4()
	return u4.String()
}
