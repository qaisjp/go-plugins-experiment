package structs

import (
	"strconv"
)

// some random struct name
type Context struct {
	IP   string // some random fields
	Port int64
}

func (c *Context) GetHost() string {
	return c.IP + ":" + strconv.FormatInt(c.Port, 10)
}
