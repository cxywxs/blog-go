package tools

import (
	"github.com/satori/go.uuid"
)

func CreateUuid() string {
	u1 := uuid.NewV4()
	return u1.String()
}
