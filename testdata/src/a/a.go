package a

import (
	"context"
	c "context"
)

func f() {
	_ = context.Background() // OK
	_ = context.TODO() // want "don't use context.TODO. Use context.Background"

	_ = c.TODO() // want "don't use context.TODO. Use context.Background"
	_ = c.Background() // OK
}
