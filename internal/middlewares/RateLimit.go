package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/JGLTechnologies/gin-rate-limit"
	"time"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(429, gin.H{"message": "Too many requests. Try again in "+time.Until(info.ResetTime).String()})
}

func RateLimit(rate time.Duration, limit uint) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  rate,
		Limit: limit,
	})
	
	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc: keyFunc,
	})
}