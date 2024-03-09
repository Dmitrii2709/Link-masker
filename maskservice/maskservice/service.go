package maskservice

import (
	"fmt"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

type Service struct {
	prod producer
	pres presenter
}

func NewServiceName(prod producer, pres presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

func (s *Service) Run() error {

	data, err := s.prod.produce()
	if err != nil {
		return fmt.Errorf("Service.producer.produce: %w", err)
	}

	data[0] = s.MaskingSpam(data[0])

	err = s.pres.present(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) MaskingSpam(a string) string {

	s1 := "http://"
	s2 := " "
	s3 := "*"
	var x []byte

	for i := 0; i < len(a); i++ {
		if len(a) > len(s1) && i < (len(a)-6) && a[i:i+len(s1)] == s1 {
			x = append(x, a[i:i+len(s1)]...)

			for j := i + len(s1); j < len(a); j++ {
				if a[j] != s2[0] {
					x = append(x, s3[0])
				} else {
					x = append(x, s2[0])
					i += len(x) - len(a[:i+1])
					break
				}
			}
			i += len(x) - len(a[:i+1])
		} else {
			x = append(x, a[i])
		}
	}
	return string(x)
}
