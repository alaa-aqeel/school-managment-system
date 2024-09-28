package helpers

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func UUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return id.String()
}
