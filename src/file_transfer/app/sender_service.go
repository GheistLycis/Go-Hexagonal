package file_transfer

import domain "Go-Hexagonal/src/file_transfer/domain"

type SenderService struct {
	listener domain.ListenerPort
}

func NewSenderService(l domain.ListenerPort) domain.SenderServicePort {
	return &SenderService{listener: l}
}

func (s *SenderService) Write(b []byte) error {
	return nil
}
