package configurer

import (
	"time"

	"github.com/pkg/errors"
)

// Duration custom duration for toml configs
type Duration struct {
	time.Duration
}

// UnmarshalText method satisfying toml unmarshal interface
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return errors.Wrap(err, "unmarshal duration type error")
}
