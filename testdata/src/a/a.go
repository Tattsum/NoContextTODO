package a

import (
	"context"
	c "context"
)

func f() {
	_ = context.Background() // OK
	_ = context.TODO() // want "NG"

	_ = c.TODO() // want "NG"
	_ = c.Background() // OK
}
