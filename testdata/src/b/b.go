package b

import ctx "context"

func f(){
	_ = ctx.Background() // OK
	_ = ctx.TODO() // NG
}
