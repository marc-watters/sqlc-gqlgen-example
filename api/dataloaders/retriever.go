package dataloaders

import "context"

type (
	Retriever interface {
		Retrieve(context.Context) *Loaders
	}
	retriever struct {
		key contextKey
	}
)

func NewRetriever() Retriever {
	return &retriever{key: key}
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}
