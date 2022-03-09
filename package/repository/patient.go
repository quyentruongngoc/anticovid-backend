package repository

import (
	"time"
)

type Patient struct {
	BaseModel
	UserUUID             string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	Gender               string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Ethnicity            string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	HIType               int
	HIExpire             time.Time
	BirthDay             time.Time
	HINumber             string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	GoogleMap            string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Occupations          string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Nationality          string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	OtherContact         string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Relative             string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	RelativeAddr         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	RelativePhone        string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	RelativeGoogleMap    string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	RelativeOtherContact string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

// func (s *Storage) UpdatePatientInfo(in patient.PatientInfo) (patient.PatientInfo, error) {
// 	data, err := s.DescribeAccount(in.UserUUID)
// 	if err != nil {
// 		return patient.PatientInfo{}, err
// 	}

// 	if data.Account.User != in.UserUUID {
// 		return patient.PatientInfo{}, fmt.Errorf("Error: User not found")
// 	}

// 	t, err := time.Parse(time.RFC3339, in.Patient.HIExpire)
// 	if err != nil {
// 		t = time.Now()
// 	}
// 	record := Patient{
// 		UserUUID:             in.UserUUID,
// 		Gender:               in.Patient.Gender,
// 		Ethnicity:            in.Patient.Ethnicity,
// 		HIType:               in.Patient.HIType,
// 		HIExpire:             t,
// 		HINumber:             in.Patient.HINumber,
// 		GoogleMap:            in.Patient.GoogleMap,
// 		Occupations:          in.Patient.Occupations,
// 		Nationality:          in.Patient.Nationality,
// 		OtherContact:         in.Patient.OtherContact,
// 		Relative:             in.Relative.Name,
// 		RelativeAddr:         in.Relative.Addr,
// 		RelativePhone:        in.Relative.Phone,
// 		RelativeGoogleMap:    in.Relative.GoogleMap,
// 		RelativeOtherContact: in.Relative.OtherContact,
// 	}

// 	t, err = time.Parse(time.RFC3339, in.Patient.BirthDay)
// 	if err != nil {
// 		t = time.Now()
// 	}

// 	record.BirthDay = t

// 	// update patient to database
// 	// if err := db.Model(&Patient{}).Debug().Where(&Patient{
// 	// 	UserUUID: in.UserUUID,
// 	// }).Update(&record).Error; err != nil {
// 	// 	// in case not found, create it
// 	// 	// if gorm.IsRecordNotFoundError(err) {

// 	// 	// }
// 	// 	if cerr := db.Model(&Patient{}).Debug().Create(&record).Error; cerr != nil {
// 	// 		return patient.PatientInfo{}, err
// 	// 	}
// 	// }

// 	if db.Model(&record).Debug().Where(&Patient{
// 		UserUUID: in.UserUUID,
// 	}).Updates(&record).RowsAffected == 0 {
// 		db.Debug().Create(&record)
// 	}

// 	return in, nil
// }

// func (s *Storage) GetPatientInfo(user string) (patient.PatientInfo, error) {
// 	var record Patient
// 	var ret patient.PatientInfo

// 	result := s.db.Debug().Where(&Patient{
// 		UserUUID: user,
// 	}).First(&record)
// 	if result.Error != nil {
// 		return patient.PatientInfo{}, fmt.Errorf("Failed to search user in database: %v : %v", user, result.Error)
// 	}

// 	ret = patient.PatientInfo{
// 		UserUUID: record.UserUUID,
// 		Patient: patient.Patient{
// 			Gender:       record.Gender,
// 			Ethnicity:    record.Ethnicity,
// 			HIType:       record.HIType,
// 			HIExpire:     record.HIExpire.Format(time.RFC3339),
// 			HINumber:     record.HINumber,
// 			GoogleMap:    record.GoogleMap,
// 			Occupations:  record.Occupations,
// 			Nationality:  record.Nationality,
// 			OtherContact: record.OtherContact,
// 			BirthDay:     record.BirthDay.Format(time.RFC3339),
// 		},
// 		Relative: patient.Relative{
// 			Name:         record.Relative,
// 			Addr:         record.RelativeAddr,
// 			Phone:        record.RelativePhone,
// 			GoogleMap:    record.RelativeGoogleMap,
// 			OtherContact: record.RelativeOtherContact,
// 		},
// 	}

// 	return ret, nil
// }
