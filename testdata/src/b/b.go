package b

import ctx "context"

func f(){
	_ = ctx.Background() // OK
	_ = ctx.TODO() // want "don't use context.Background. Use context.TODO"
}
