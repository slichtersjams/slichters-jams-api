package app

import "context"

type IGetContext interface {
	getContext() context.Context
}
