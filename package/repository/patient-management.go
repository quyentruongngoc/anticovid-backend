package repository

import (
	"time"
)

type PatientMgmt struct {
	PatientUUID        string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	TreatmentPlace     string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	TreatmentMap       string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	DoctorUUID         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	AdministrativeUUID string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	HealthcareUUID     string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	SupplyUUID         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	ReceivedTime       time.Time
	Transfer           bool
	TransferTime       time.Time
	TransferNote       string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	StopTreatment      bool
	StopTreatmentTime  time.Time
	StopTreatmentNote  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	EndStreatment      bool
	EndStreatmentTime  time.Time
	EndStreatmentNote  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Dead               bool
	DeadTime           time.Time
	DeadNote           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

// func (s *Storage) UpdatePatientMgmt(in patient.PatientMgmt) (patient.PatientMgmt, error) {
// 	data, err := s.DescribeAccount(in.UserUUID)
// 	if err != nil {
// 		return patient.PatientMgmt{}, err
// 	}

// 	if data.Account.User != in.UserUUID {
// 		return patient.PatientMgmt{}, fmt.Errorf("Error: User not found")
// 	}

// 	t, err := time.Parse(time.RFC3339, in.ReceivedTime)
// 	if err != nil {
// 		return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.ReceivedTime, err)
// 	}
// 	record := PatientMgmt{
// 		PatientUUID:        in.UserUUID,
// 		ReceivedTime:       t,
// 		TreatmentPlace:     in.TreatmentPlace,
// 		TreatmentMap:       in.TreatmentMap,
// 		DoctorUUID:         in.Doctor.Account.User,
// 		AdministrativeUUID: in.Administrative.Account.User,
// 		HealthcareUUID:     in.Healthcare.Account.User,
// 		SupplyUUID:         in.Supply.Account.User,
// 		Transfer:           in.Transfer.Status,
// 		TransferTime:       t,
// 		TransferNote:       in.Transfer.Note,
// 		StopTreatment:      in.StopTreatment.Status,
// 		StopTreatmentTime:  t,
// 		StopTreatmentNote:  in.StopTreatment.Note,
// 		EndStreatment:      in.EndTreatment.Status,
// 		EndStreatmentTime:  t,
// 		EndStreatmentNote:  in.EndTreatment.Note,
// 		Dead:               in.Dead.Status,
// 		DeadTime:           t,
// 		DeadNote:           in.Dead.Note,
// 	}

// 	t, err = time.Parse(time.RFC3339, in.Transfer.Time)
// 	if err == nil {
// 		record.TransferTime = t
// 		// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.Transfer.Time, err)
// 	}

// 	t, err = time.Parse(time.RFC3339, in.StopTreatment.Time)
// 	if err == nil {
// 		record.StopTreatmentTime = t
// 		// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.StopTreatment.Time, err)
// 	}

// 	t, err = time.Parse(time.RFC3339, in.EndTreatment.Time)
// 	if err == nil {
// 		record.EndStreatmentTime = t
// 		// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.EndTreatment.Time, err)
// 	}

// 	t, err = time.Parse(time.RFC3339, in.Dead.Time)
// 	if err == nil {
// 		record.DeadTime = t
// 		// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.Dead.Time, err)
// 	}

// 	if db.Model(&record).Debug().Where(&PatientMgmt{
// 		PatientUUID: in.UserUUID,
// 	}).Updates(&record).RowsAffected == 0 {
// 		db.Debug().Create(&record)
// 	}

// 	// Just use pointer but it cause alot of nil check, so I fixed here
// 	db.Model(&record).Debug().Where(&PatientMgmt{
// 		PatientUUID: in.UserUUID,
// 	}).Updates(map[string]interface{}{
// 		"transfer":       in.Transfer.Status,
// 		"stop_treatment": in.StopTreatment.Status,
// 		"end_streatment": in.EndTreatment.Status,
// 		"dead":           in.Dead.Status,
// 	})

// 	temp := Account{}
// 	if in.Dead.Status || in.EndTreatment.Status || in.Transfer.Status || in.StopTreatment.Status {
// 		temp.Discharge = true
// 	} else {
// 		temp.Discharge = false
// 	}

// 	db.Model(&temp).Debug().Where(&Account{
// 		User: in.UserUUID,
// 	}).Updates(&temp)

// 	return in, nil
// }

// func (s *Storage) descPatientMgmtByStaff(user string, search string, role uint) ([]PatientMgmt, error) {
// 	var records []PatientMgmt
// 	where := &PatientMgmt{}

// 	switch role {
// 	case constant.DoctorRole:
// 		where.DoctorUUID = user
// 	case constant.HealthcareRole:
// 		where.HealthcareUUID = user
// 	case constant.SupplyRole:
// 		where.SupplyUUID = user
// 	case constant.AdministrativeRole:
// 		where.AdministrativeUUID = user
// 	default:
// 		return []PatientMgmt{}, fmt.Errorf("Invalid role")
// 	}

// 	result := s.db.Debug().Where(&where).Find(&records)
// 	if result.Error != nil {
// 		return []PatientMgmt{}, fmt.Errorf("Failed to search user in database: %v : %v", user, result.Error)
// 	}

// 	return records, nil

// }

// func (s *Storage) GetPatientMgmt(user string) (patient.PatientMgmt, error) {
// 	var record PatientMgmt
// 	var ret patient.PatientMgmt

// 	result := s.db.Debug().Where(&PatientMgmt{
// 		PatientUUID: user,
// 	}).First(&record)
// 	if result.Error != nil {
// 		return patient.PatientMgmt{}, fmt.Errorf("Failed to search user in database: %v : %v", user, result.Error)
// 	}

// 	ret = patient.PatientMgmt{
// 		UserUUID:       user,
// 		ReceivedTime:   record.ReceivedTime.Format(time.RFC3339),
// 		TreatmentPlace: record.TreatmentPlace,
// 		TreatmentMap:   record.TreatmentMap,
// 		// DoctorUUID:         record.DoctorUUID,
// 		// AdministrativeUUID: record.AdministrativeUUID,
// 		// HealthcareUUID:     record.HealthcareUUID,
// 		// SupplyUUID:         record.SupplyUUID,
// 		Transfer: patient.Status{
// 			Time:   record.TransferTime.Format(time.RFC3339),
// 			Status: record.Transfer,
// 			Note:   record.TransferNote,
// 		},
// 		StopTreatment: patient.Status{
// 			Time:   record.StopTreatmentTime.Format(time.RFC3339),
// 			Status: record.StopTreatment,
// 			Note:   record.StopTreatmentNote,
// 		},
// 		EndTreatment: patient.Status{
// 			Time:   record.EndStreatmentTime.Format(time.RFC3339),
// 			Status: record.EndStreatment,
// 			Note:   record.EndStreatmentNote,
// 		},
// 		Dead: patient.Status{
// 			Time:   record.DeadTime.Format(time.RFC3339),
// 			Status: record.Dead,
// 			Note:   record.DeadNote,
// 		},
// 	}

// 	ret.Administrative, _ = s.DescribeAccount(record.AdministrativeUUID)
// 	// ret.Administrative.Passwd = ""
// 	// ret.Administrative.CreateBy = ""
// 	// ret.Administrative.Token = ""
// 	ret.Doctor, _ = s.DescribeAccount(record.DoctorUUID)
// 	// ret.Doctor.Passwd = ""
// 	// ret.Doctor.CreateBy = ""
// 	// ret.Doctor.Token = ""
// 	ret.Supply, _ = s.DescribeAccount(record.SupplyUUID)
// 	// ret.Supply.Passwd = ""
// 	// ret.Supply.CreateBy = ""
// 	// ret.Supply.Token = ""
// 	ret.Healthcare, _ = s.DescribeAccount(record.HealthcareUUID)
// 	// ret.Healthcare.Passwd = ""
// 	// ret.Healthcare.CreateBy = ""
// 	// ret.Healthcare.Token = ""

// 	return ret, nil
// }
