package handler

import (
	"github.com/jesee-kuya/transport-system/middleware"
	"github.com/jesee-kuya/transport-system/service"
)

type Transport struct {
	Middleware            middleware.Middleware
	AuthService           service.Authentication
	SchoolService         service.SchoolService
	GuardianService       service.GuardianService
	PrivateParentService  service.PrivateParentService
	SchoolDriverService   service.SchoolDriverService
	PrivateDriverService  service.PrivateDriverService
}
