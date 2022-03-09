package repository

import (
	"anti-corona-backend/package/api-process/status"
	"anti-corona-backend/package/constant"
	"fmt"
	"time"
)

type StaffRecord struct {
	BaseModel
	UserUUID     string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Note         string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	RecordNumber int
	Role         int
	Type         int
	PatientUUID  string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	ParentID     int
}

func (s *Storage) AddStatusRecord(in status.StaffStatus) (status.StaffStatus, error) {
	record := StaffRecord{
		UserUUID:     in.UserUUID,
		Note:         in.Note,
		RecordNumber: in.RecordNumber,
		PatientUUID:  in.PatientUUID,
		Role:         in.Role,
		Type:         in.Type,
		ParentID:     in.ParentID,
	}

	result := s.db.Model(&record).Debug().Create(&record)
	if result.Error != nil {
		return status.StaffStatus{}, result.Error
	}

	return in, nil
}

func (s *Storage) getStatusRecord(where StaffRecord) ([]StaffRecord, error) {
	var records []StaffRecord

	result := s.db.Debug().Where(&where).Find(&records)

	if result.Error != nil {
		return []StaffRecord{}, fmt.Errorf("Failed to search staff record in database: %v : %v", where, result.Error)
	}

	return records, nil
}

func (s *Storage) GetStatusRecord(in status.StaffStatus) ([]status.StaffStatus, error) {
	var records []StaffRecord
	var ret []status.StaffStatus

	where := StaffRecord{
		Role:         in.Role,
		RecordNumber: in.RecordNumber,
		PatientUUID:  in.PatientUUID,
	}

	if in.RecordNumber == 255 {
		where.RecordNumber = 0
	}

	records, err := s.getStatusRecord(where)
	if err != nil {
		return []status.StaffStatus{}, err
	}

	// result := s.db.Debug().Where(&where).Find(&records)
	// if result.Error != nil {
	// 	return []status.StaffStatus{}, fmt.Errorf("Failed to search staff record in database: %v : %v", in, result.Error)
	// }

	for _, rec := range records {
		ins := status.StaffStatus{
			ParentID:     rec.ParentID,
			ID:           int(rec.ID),
			Note:         rec.Note,
			CreateAt:     rec.CreatedAt.Format(time.RFC3339),
			RecordNumber: rec.RecordNumber,
			UserUUID:     rec.UserUUID,
			PatientUUID:  rec.PatientUUID,
			Role:         rec.Role,
			Type:         rec.Type,
			Token:        "",
		}
		ret = append(ret, ins)
	}

	return ret, nil
}

func (s *Storage) GetStaffRecordList(patient string) ([]status.StaffRecordList, error) {
	// var records []StaffRecord
	var ret []status.StaffRecordList

	for i := constant.DoctorRole; i < constant.MaxRole; i++ {
		where := StaffRecord{
			Role:        i,
			PatientUUID: patient,
		}

		records, err := s.getStatusRecord(where)
		if err != nil {
			return []status.StaffRecordList{}, err
		}

		var tempList []status.RecordData

		for _, rec := range records {
			ins := status.RecordData{
				RecordNumber: rec.RecordNumber,
				RecordType:   rec.Type,
			}
			tempList = append(tempList, ins)
		}

		tempInst := status.StaffRecordList{
			Role:       i,
			RecordData: tempList,
		}

		ret = append(ret, tempInst)
	}

	return ret, nil

}
