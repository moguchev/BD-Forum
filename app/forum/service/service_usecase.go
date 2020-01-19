package service

import ()

func (s Service) Clear() error {
	err := s.Repository.Clear()

	if err != nil {
		return err
	}

	return nil
}
