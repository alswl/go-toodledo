package services

import (
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/sirupsen/logrus"
)

const SettingBucket = "settings"

type SettingService interface {
	Find(name string) (string, error)
	Put(name, body string) error
	Delete(name string) error
}

type settingService struct {
	log logrus.FieldLogger
	db  dal.Backend
}

func NewSettingService(log logrus.FieldLogger, db dal.Backend) SettingService {
	return &settingService{log: log, db: db}
}

func (s *settingService) Find(name string) (string, error) {
	bytes, err := s.db.Get(SettingBucket, name)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *settingService) Put(name, body string) error {
	err := s.db.Put(SettingBucket, name, []byte(body))
	return err
}

func (s *settingService) Delete(name string) error {
	return s.db.Remove(SettingBucket, name)
}
