package proxy

import (
	"fmt"
	"task-management/server/spec/authspec"
	"task-management/server/spec/customerspec"
	"task-management/server/spec/taskspec"
)

const (
	BasePath       = "/"
	localhost      = "http://localhost"
	ProxyPortHttp  = ":80"
	ProxyPortHttps = ":443"
)

var Routes = map[string]string{
	taskspec.BasePath:     fmt.Sprintf("%s%s", localhost, taskspec.Port),
	authspec.BasePath:     fmt.Sprintf("%s%s", localhost, authspec.Port),
	customerspec.BasePath: fmt.Sprintf("%s%s", localhost, customerspec.Port),
}
