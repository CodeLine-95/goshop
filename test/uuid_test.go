package test

import (
	"fmt"
	"github.com/gofrs/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Println(u1)
}
