package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RequestTimeoutMiddleware(timeout *time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), *timeout)
		defer cancel()

		opDone := make(chan struct{})

		go func() {
			c.Next()
			opDone <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			c.AbortWithStatus(http.StatusRequestTimeout)
		case <-opDone:

		}
	}
}
