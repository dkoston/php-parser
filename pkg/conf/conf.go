package conf

import (
	"github.com/dkoston/php-parser/pkg/errors"
	"github.com/dkoston/php-parser/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
