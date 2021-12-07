package divvy

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"sigs.k8s.io/yaml"
)

func ErrorStringsToError(errmsgs []string) error {
	if len(errmsgs) != 0 {
		o, err := yaml.Marshal(errmsgs)
		if err != nil {
			panic(err)
		}
		return errors.Wrap(errors.ErrInvalidRequest, string(o))
	}
	return nil
}
