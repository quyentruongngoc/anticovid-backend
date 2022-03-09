package repository

import (
	"fmt"
	"reflect"
	"strconv"
)

type PatientSelf struct {
	Bptc       string `json:"bptc,omitempty"`
	Dtd        string `json:"dtd,omitempty"`
	Bhs        string `json:"bhs,omitempty"`
	DtdKks     string `json:"dtd_kks,omitempty"`
	Bp         string `json:"bp,omitempty"`
	Lvbpk      string `json:"lvbpk,omitempty"`
	Utm        string `json:"utm,omitempty"`
	Utp        string `json:"utp,omitempty"`
	Utdc       string `json:"utdc,omitempty"`
	Gt         string `json:"gt,omitempty"`
	Stent      string `json:"stent,omitempty"`
	Tm         string `json:"tm,omitempty"`
	Blmmn      string `json:"blmmn,omitempty"`
	Tha        string `json:"tha,omitempty"`
	Hcd        string `json:"hcd,omitempty"`
	Bltk       string `json:"bltk,omitempty"`
	Gmd        string `json:"gmd,omitempty"`
	Btmt       string `json:"btmt,omitempty"`
	Sdcgn      string `json:"sdcgn,omitempty"`
	Htn        string `json:"htn,omitempty"`
	Sdchctucmd string `json:"sdchctucmd,omitempty"`
	Cbht       string `json:"cbht,omitempty"`
	Thdu       string `json:"thdu,omitempty"`
}

type PatientMedical struct {
	UserUUID              string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	Reason                string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	PathologicalProcess   string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Self                  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	SelfNote              string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Allergy               string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Medicine              string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Cigarette             string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Alcohol               string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Drug                  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	CharacteristicsOthers string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Family                string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Environment           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	AcuteIllness          string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Place                 string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	BodyNote              string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Height                int
	Weight                int
	Cyclic                string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Respiratory           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Digest                string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Kidney                string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Nerve                 string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Musculoskeletal       string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Endocrine             string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Ent                   string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Nutrition             string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	AgenciesOthers        string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	MainDisease           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	SideDisease           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Distinguish           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Prognosis             string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Treatments            string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Subclinical           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	SumDischarge          string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	SumTreatDir           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Vaccination           string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

func checkSelfBackground(in PatientSelf) bool {
	v := reflect.ValueOf(in)
	// values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		temp := v.Field(i).Interface()
		strTemp := fmt.Sprintf("%v", temp)
		isTrue, _ := strconv.Atoi(strTemp)
		if isTrue != 0 {
			return true
		}
	}

	return false
}

// func (s *Storage) UpdatePatientMedical(in patient.MedicalData) (patient.MedicalData, error) {

// 	data, err := s.DescribeAccount(in.UserUUID)
// 	if err != nil {
// 		return patient.MedicalData{}, err
// 	}

// 	if data.Account.User != in.UserUUID {
// 		return patient.MedicalData{}, fmt.Errorf("Error: User not found")
// 	}

// 	record := PatientMedical{
// 		UserUUID:              in.UserUUID,
// 		Reason:                in.Reason,
// 		PathologicalProcess:   in.Ask.PathologicalProcess,
// 		Self:                  in.Ask.MedicalHistory.Self,
// 		SelfNote:              in.Ask.MedicalHistory.SelfNote,
// 		Allergy:               in.Ask.MedicalHistory.Characteristics.Allergy,
// 		Medicine:              in.Ask.MedicalHistory.Characteristics.Medicine,
// 		Cigarette:             in.Ask.MedicalHistory.Characteristics.Cigarette,
// 		Alcohol:               in.Ask.MedicalHistory.Characteristics.Alcohol,
// 		Drug:                  in.Ask.MedicalHistory.Characteristics.Drug,
// 		CharacteristicsOthers: in.Ask.MedicalHistory.Characteristics.Others,
// 		Family:                in.Ask.MedicalHistory.Family,
// 		Environment:           in.Ask.MedicalHistory.Epidemiology.Environment,
// 		AcuteIllness:          in.Ask.MedicalHistory.Epidemiology.AcuteIllness,
// 		Place:                 in.Ask.MedicalHistory.Epidemiology.Place,
// 		BodyNote:              in.Examination.Body.Note,
// 		Height:                in.Examination.Body.Height,
// 		Weight:                in.Examination.Body.Weight,
// 		Cyclic:                in.Examination.Agencies.Cyclic,
// 		Respiratory:           in.Examination.Agencies.Respiratory,
// 		Digest:                in.Examination.Agencies.Digest,
// 		Kidney:                in.Examination.Agencies.Kidney,
// 		Nerve:                 in.Examination.Agencies.Nerve,
// 		Musculoskeletal:       in.Examination.Agencies.Musculoskeletal,
// 		Endocrine:             in.Examination.Agencies.Endocrine,
// 		Ent:                   in.Examination.Agencies.Ent,
// 		Nutrition:             in.Examination.Agencies.Nutrition,
// 		AgenciesOthers:        in.Examination.Agencies.Others,
// 		MainDisease:           in.Diagnostic.MainDisease,
// 		SideDisease:           in.Diagnostic.SideDisease,
// 		Distinguish:           in.Diagnostic.Distinguish,
// 		Prognosis:             in.Prognosis,
// 		Treatments:            in.Treatments,
// 		Subclinical:           in.Examination.Subclinical,
// 		SumDischarge:          in.Summary.Discharge,
// 		SumTreatDir:           in.Summary.TreatDir,
// 		Vaccination:           in.Vaccination,
// 	}

// 	if db.Model(&record).Debug().Where(&PatientMedical{
// 		UserUUID: in.UserUUID,
// 	}).Updates(&record).RowsAffected == 0 {
// 		db.Debug().Create(&record)
// 	}

// 	return in, nil
// }

// func (s *Storage) GetPatientMedical(user string) (patient.MedicalData, error) {
// 	var record PatientMedical
// 	var ret patient.MedicalData

// 	result := s.db.Debug().Where(&PatientMedical{
// 		UserUUID: user,
// 	}).First(&record)
// 	if result.Error != nil {
// 		return patient.MedicalData{}, fmt.Errorf("Failed to search user in database: %v : %v", user, result.Error)
// 	}

// 	data, err := s.GetPatientMgmt(user)
// 	if err != nil {
// 		return patient.MedicalData{}, fmt.Errorf("Failed to search user mgmt in database: %v : %v", user, err)
// 	}

// 	ret = patient.MedicalData{
// 		UserUUID:     user,
// 		Reason:       record.Reason,
// 		Token:        "",
// 		ReceivedTime: data.ReceivedTime,
// 		Ask: patient.AskPatient{
// 			PathologicalProcess: record.PathologicalProcess,
// 			MedicalHistory: patient.MedicalHistory{
// 				Self:     record.Self,
// 				SelfNote: record.SelfNote,
// 				Characteristics: patient.Characteristics{
// 					Allergy:   record.Allergy,
// 					Medicine:  record.Medicine,
// 					Cigarette: record.Cigarette,
// 					Alcohol:   record.Alcohol,
// 					Drug:      record.Drug,
// 					Others:    record.CharacteristicsOthers,
// 				},
// 				Family: record.Family,
// 				Epidemiology: patient.Epidemiology{
// 					Environment:  record.Environment,
// 					AcuteIllness: record.AcuteIllness,
// 					Place:        record.Place,
// 				},
// 			},
// 		},
// 		Examination: patient.Examination{
// 			Body: patient.Body{
// 				Note:   record.BodyNote,
// 				Height: record.Height,
// 				Weight: record.Weight,
// 			},
// 			Agencies: patient.Agencies{
// 				Cyclic:          record.Cyclic,
// 				Respiratory:     record.Respiratory,
// 				Digest:          record.Digest,
// 				Kidney:          record.Kidney,
// 				Nerve:           record.Nerve,
// 				Musculoskeletal: record.Musculoskeletal,
// 				Endocrine:       record.Endocrine,
// 				Ent:             record.Ent,
// 				Nutrition:       record.Nutrition,
// 				Others:          record.AgenciesOthers,
// 			},
// 			Subclinical: record.Subclinical,
// 		},
// 		Diagnostic: patient.Diagnostic{
// 			MainDisease: record.MainDisease,
// 			SideDisease: record.SideDisease,
// 			Distinguish: record.Distinguish,
// 		},
// 		Summary: patient.Summary{
// 			Discharge: record.SumDischarge,
// 			TreatDir:  record.SumTreatDir,
// 		},
// 		Prognosis:   record.Prognosis,
// 		Treatments:  record.Treatments,
// 		Vaccination: record.Vaccination,
// 	}

// 	return ret, nil
// }
