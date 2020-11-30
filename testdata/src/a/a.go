package a

import (
	"context"
	c "context"
)

func f() {
	_ = context.Background() // OK
	_ = context.TODO() // want "don't use context.Background. Use context.TODO"

	_ = c.TODO() // want "don't use context.Background. Use context.TODO"
	_ = c.Background() // OK
}
