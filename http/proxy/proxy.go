package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

var Proxies = []Proxy{
	newLimitProxy("user", 3, time.Minute*3),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range Proxies {
		if v.Accept(c.Params("key")) {
			return v.Proxy(c)
		}
	}
	c.Response().SetStatusCode(400)
	return nil
}
