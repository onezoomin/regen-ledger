package server

import "fmt"

func ormError(err error) error {
	if err != nil {
		return fmt.Errorf("ORM: %w", err)
	}
	return nil
}

func eventError(err error) error {
	if err != nil {
		return fmt.Errorf("EVENT: %w", err)
	}
	return nil
}
