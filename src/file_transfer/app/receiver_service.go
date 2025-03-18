package file_transfer

import domain "Go-Hexagonal/src/file_transfer/domain"

type ReceiverService struct {
	listener domain.ListenerPort
}

func NewReceiverService(l domain.ListenerPort) domain.ReceiverServicePort {
	return &ReceiverService{listener: l}
}

func (s *ReceiverService) Read(b []byte) error {
	return nil
}
