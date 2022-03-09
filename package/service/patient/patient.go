package patient

type Repository interface {
	// UpdatePatientInfo(patient.PatientInfo) (patient.PatientInfo, error)
	// GetPatientInfo(string) (patient.PatientInfo, error)
	// UpdatePatientMgmt(patient.PatientMgmt) (patient.PatientMgmt, error)
	// GetPatientMgmt(string) (patient.PatientMgmt, error)
	// UpdatePatientMedical(patient.MedicalData) (patient.MedicalData, error)
	// GetPatientMedical(string) (patient.MedicalData, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}

// func (s *Service) UpdatePatientInfo(in patient.PatientInfo) (patient.PatientInfo, error) {
// 	instance, err := s.repo.UpdatePatientInfo(in)
// 	if err != nil {
// 		return patient.PatientInfo{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) GetPatientInfo(user string) (patient.PatientInfo, error) {
// 	instance, err := s.repo.GetPatientInfo(user)
// 	if err != nil {
// 		return patient.PatientInfo{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) UpdatePatientMgmt(in patient.PatientMgmt) (patient.PatientMgmt, error) {
// 	instance, err := s.repo.UpdatePatientMgmt(in)
// 	if err != nil {
// 		return patient.PatientMgmt{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) GetPatientMgmt(user string) (patient.PatientMgmt, error) {
// 	instance, err := s.repo.GetPatientMgmt(user)
// 	if err != nil {
// 		return patient.PatientMgmt{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) UpdatePatientMedical(in patient.MedicalData) (patient.MedicalData, error) {
// 	instance, err := s.repo.UpdatePatientMedical(in)
// 	if err != nil {
// 		return patient.MedicalData{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) GetPatientMedical(user string) (patient.MedicalData, error) {
// 	instance, err := s.repo.GetPatientMedical(user)
// 	if err != nil {
// 		return patient.MedicalData{}, err
// 	}

// 	return instance, nil
// }
