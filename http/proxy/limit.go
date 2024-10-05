package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var counter = map[string]*Limit{}

type Limit struct {
	count int
	ttl   time.Time
}

type LimitProxy struct {
	key   string
	limit int
	ttl   time.Duration
}

func newLimitProxy(key string, limit int, ttl time.Duration) LimitProxy {
	return LimitProxy{
		key:   key,
		limit: limit,
		ttl:   ttl,
	}
}

func (p LimitProxy) Accept(key string) bool {
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()

	if r, ok := counter[path]; ok && r.count >= p.limit && r.ttl.After(time.Now()) {
		c.Response().SetStatusCode(429)
		fmt.Printf("request limit reached for %s \n", path)
		return nil
	} else if !ok {
		counter[path] = &Limit{
			count: 0,
			ttl:   time.Now().Add(p.ttl),
		}
	}
	c.SendString("Go - 103 Http Package")
	counter[path].count++

	return nil
}
