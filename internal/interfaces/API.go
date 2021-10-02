// Package interfaces defines common interfaces between the service's components
package interfaces

import (
	"github.com/o-aloqaily/go-project-starter/pkg/router"
)

// API is the interface which any router must implement
type API interface {
	Handlers() router.Router
}
