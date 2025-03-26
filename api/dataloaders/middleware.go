package dataloaders

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

func MWareDataLoader(repoSvc pgx.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		loaders := newLoaders(ctx, repoSvc)
		augmentedCtx := context.WithValue(ctx, key, loaders)
		c.Request = c.Request.WithContext(augmentedCtx)
		c.Next()
	}
}
