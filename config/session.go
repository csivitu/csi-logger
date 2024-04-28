package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New(session.Config{
    Expiration: 48 * time.Hour,  
})
