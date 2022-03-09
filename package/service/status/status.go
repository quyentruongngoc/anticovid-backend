package status

import (
	"anti-corona-backend/package/api-process/status"
)

type Repository interface {
	AddStatusRecord(status.StaffStatus) (status.StaffStatus, error)
	GetStatusRecord(status.StaffStatus) ([]status.StaffStatus, error)
	AddPatientStatus(status.PatientStatus) (status.PatientStatus, error)
	UpdatePatientStatus(status.PatientStatus) (status.PatientStatus, error)
	GetPatientStatus(status.PatientStatus) ([]status.PatientStatus, error)
	GetStaffRecordList(string) ([]status.StaffRecordList, error)
	GetPatientRecordList(string) ([]status.StaffRecordList, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) AddStatusRecord(in status.StaffStatus) (status.StaffStatus, error) {
	instance, err := s.repo.AddStatusRecord(in)
	if err != nil {
		return status.StaffStatus{}, err
	}

	return instance, nil
}

func (s *Service) GetStatusRecord(in status.StaffStatus) ([]status.StaffStatus, error) {
	instances, err := s.repo.GetStatusRecord(in)
	if err != nil {
		return []status.StaffStatus{}, err
	}

	return instances, nil
}

func (s *Service) AddPatientStatus(in status.PatientStatus) (status.PatientStatus, error) {
	instance, err := s.repo.AddPatientStatus(in)
	if err != nil {
		return status.PatientStatus{}, err
	}

	return instance, nil
}

func (s *Service) UpdatePatientStatus(in status.PatientStatus) (status.PatientStatus, error) {
	instance, err := s.repo.UpdatePatientStatus(in)
	if err != nil {
		return status.PatientStatus{}, err
	}

	return instance, nil
}

func (s *Service) GetPatientStatus(in status.PatientStatus) ([]status.PatientStatus, error) {
	instances, err := s.repo.GetPatientStatus(in)
	if err != nil {
		return []status.PatientStatus{}, err
	}

	return instances, nil
}

func (s *Service) GetStaffRecordList(patient string) ([]status.StaffRecordList, error) {
	instances, err := s.repo.GetStaffRecordList(patient)
	if err != nil {
		return []status.StaffRecordList{}, err
	}

	return instances, nil
}

func (s *Service) GetPatientRecordList(patient string) ([]status.StaffRecordList, error) {
	instances, err := s.repo.GetPatientRecordList(patient)
	if err != nil {
		return []status.StaffRecordList{}, err
	}

	return instances, nil
}
