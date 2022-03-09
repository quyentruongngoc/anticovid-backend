package repository

import (
	"anti-corona-backend/package/api-process/status"
	"anti-corona-backend/package/constant"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type PatientStatus struct {
	BaseModel
	PatientUUID   string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Type          int
	Serious       bool
	SeriousDetail string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Audio         string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	// oxy trong mau
	OxyBlood int
	// nhip tim
	HeartBeat int
	// than nhiet
	Temperature float32
	// huyet ap group
	// nguong tren
	BloodPressureUp int
	// nguong duoi
	BloodPressureDown int
	// ho hap group
	// so lan tho/phut
	Respiratory int
	// kho tho
	Stuffy bool
	// tho co keo, rut lom
	BreathRetract bool
	// tho bung
	BellyBreath bool
	// tieu hoa
	// tieu chay
	Diarrhea bool
	// buon non/ non
	Nausea bool
	// dau bung
	Stomachache bool
	// tinh trang khac
	// ho khang
	DryCough bool
	// dau hong
	SoreThroat bool
	// nghet mui
	StuffyNose bool
	// so mui
	Snivel bool
	// met moi
	Tire bool
	// dau dau
	Headache bool
	// dau moi co
	MusclePain bool
	// mat vi giac, khuu giac
	LossTaste bool
	// bieu hien nghiem trong
	// moi tim tai
	PurpleLip bool
	// rat kho tho, phai gang suc
	VeryStuffy bool
	// dau tuc nguc
	ChestPain bool
	// noi lap
	Stutter bool
	// co giat
	Convulsion bool
	// kho moi mieng
	DryMouth bool
	// tieu it
	UrinaLess bool
	// record number
	RecordNumber int
	// tho oxy
	BreathOxy        bool
	OxyConcentration int
	OxyVolume        int
	OxyTimesPerDay   int
	OxyDuraPerTimes  int
	// y thuc
	Consciousness   int
	MachineSeverity int
	DoctorSeverity  int
	// new ho hap
	VeryStuffyNew int
	// kho chiu trong long nguc
	ChestDiscomfor bool
	// Van dong thay met hon moi khi
	TiredThanUsual bool
	// ho ra dam hong
	Hemoptisi bool
}

func (s *Storage) calculateSeverity(rec PatientStatus) int {
	var sev = 0
	isLv3 := false
	isLv4 := false

	// nhip tho
	if rec.Respiratory == 0 {
		return 0
	}
	if (rec.Respiratory >= 1) && (rec.Respiratory <= 8) {
		sev += 3
		isLv3 = true
	} else if (rec.Respiratory >= 9) && (rec.Respiratory <= 11) {
		sev += 1
	} else if (rec.Respiratory >= 25) && (rec.Respiratory <= 30) {
		sev += 2
	} else if rec.Respiratory > 30 {
		sev += 3
		isLv3 = true
	}

	// oxy
	// isBreathOxy := false
	// if (rec.OxyConcentration > 0) || (rec.OxyDuraPerTimes > 0) || (rec.OxyVolume > 0) || (rec.OxyTimesPerDay > 0) {
	// 	// breath oxy = yes
	// 	isBreathOxy = true
	// 	sev += 2
	// }

	userAcc, _ := s.DescribeUser(rec.PatientUUID)
	birthDate, _ := time.Parse(time.RFC3339, userAcc.Mgmt.Birthday)
	currentTime := time.Now()

	old := currentTime.Year() - birthDate.Year()
	gender := userAcc.Mgmt.Gender

	log.Printf("Quyen debug: old %v", old)
	log.Printf("Quyen debug: gender %v", gender)

	var (
		BloodPressureSmall int
		BloodPressureBig   int

		HeartBeatSmall int
		HeartBeatBig   int
	)

	if old < 20 {
		if gender {
			BloodPressureSmall = 85
			BloodPressureBig = 115
			HeartBeatSmall = 57
			HeartBeatBig = 97
		} else {
			BloodPressureSmall = 90
			BloodPressureBig = 120
			HeartBeatSmall = 58
			HeartBeatBig = 99
		}
	} else if (old >= 20) && (old <= 45) {
		if gender {
			BloodPressureSmall = 105
			BloodPressureBig = 135
			HeartBeatSmall = 52
			HeartBeatBig = 89
		} else {
			BloodPressureSmall = 110
			BloodPressureBig = 140
			HeartBeatSmall = 57
			HeartBeatBig = 95
		}
	} else if old > 45 {
		if gender {
			BloodPressureSmall = 115
			BloodPressureBig = 145
			HeartBeatSmall = 50
			HeartBeatBig = 91
		} else {
			BloodPressureSmall = 120
			BloodPressureBig = 150
			HeartBeatSmall = 56
			HeartBeatBig = 92
		}
	}

	// huyen ap nguong duoi (tam truong)
	if rec.BloodPressureDown == 0 {
		return 0
	}
	// huyet ap nguong tren (tam thu)
	if rec.BloodPressureUp == 0 {
		return 0
	}

	if (rec.BloodPressureUp - rec.BloodPressureDown) <= 20 {
		return 5
	}

	if (rec.BloodPressureUp >= 1) && (rec.BloodPressureUp < (BloodPressureSmall - 20)) {
		sev += 3
		isLv3 = true
	} else if (rec.BloodPressureUp >= (BloodPressureSmall - 20)) && (rec.BloodPressureUp <= (BloodPressureSmall - 10)) {
		sev += 2
	} else if (rec.BloodPressureUp > (BloodPressureSmall - 10)) && (rec.BloodPressureUp < BloodPressureSmall) {
		sev += 1
	} else if (rec.BloodPressureUp > BloodPressureBig) && (rec.BloodPressureUp < (BloodPressureBig + 30)) {
		sev += 1
	} else if (rec.BloodPressureUp >= (BloodPressureBig + 30)) && (rec.BloodPressureUp <= (BloodPressureBig + 60)) {
		sev += 2
	} else if rec.BloodPressureUp > (BloodPressureBig + 60) {
		sev += 3
		isLv3 = true
	}

	// mach (nhip tim)
	if rec.HeartBeat == 0 {
		return 0
	}
	if (rec.HeartBeat >= 1) && (rec.HeartBeat <= (HeartBeatSmall - 15)) {
		sev += 3
		isLv3 = true
	} else if (rec.HeartBeat > (HeartBeatSmall - 15)) && (rec.HeartBeat < HeartBeatSmall) {
		sev += 1
	} else if (rec.HeartBeat > HeartBeatBig) && (rec.HeartBeat <= (HeartBeatBig + 15)) {
		sev += 1
	} else if (rec.HeartBeat > (HeartBeatBig + 15)) && (rec.HeartBeat <= (HeartBeatBig + 30)) {
		sev += 2
	} else if rec.HeartBeat > (HeartBeatBig + 30) {
		sev += 3
		isLv3 = true
	}

	// y thuc
	if rec.Consciousness == 0 {
		return 0
	}
	if rec.Consciousness == 2 {
		sev += 3
		isLv3 = true
	} else if rec.Consciousness == 3 {
		sev += 3
		isLv4 = true
	} else if rec.Consciousness == 4 {
		return 5
	}
	// sev += (rec.Consciousness - 1)

	// nhiet do
	if rec.Temperature == 0 {
		return 0
	}
	if (rec.Temperature >= 1.0) && (rec.Temperature <= 35.0) {
		sev += 3
		isLv3 = true
	} else if (rec.Temperature >= 35.1) && (rec.HeartBeat <= 36.0) {
		sev += 1
	} else if (rec.Temperature >= 38.1) && (rec.Temperature <= 39.0) {
		sev += 1
	} else if rec.Temperature >= 39.1 {
		sev += 2
	}

	// SpO2
	if rec.OxyBlood == 0 {
		return 0
	}
	data, _ := s.DescribeUser(rec.PatientUUID)
	log.Printf("Quyen debug self data json: %+v", data.MedicalInfo.SelfHistory)
	selfData := PatientSelf{}
	json.Unmarshal([]byte(data.MedicalInfo.SelfHistory), &selfData)
	log.Printf("Quyen debug self data %+v", selfData)

	if checkSelfBackground(selfData) {
		log.Printf("Quyen debug checkSelfBackground TRUE")
		sev += 1
	}

	isCOPD, _ := strconv.Atoi(selfData.Bp)
	if isCOPD == 1 {
		if (rec.OxyBlood >= 1) && (rec.OxyBlood <= 83) {
			sev += 3
			isLv3 = true
		} else if (rec.OxyBlood >= 84) && (rec.OxyBlood <= 85) && (!rec.BreathOxy) {
			sev += 2
		} else if (rec.OxyBlood >= 86) && (rec.OxyBlood <= 87) && (!rec.BreathOxy) {
			sev += 1
		} else if (rec.OxyBlood >= 90) && (rec.OxyBlood <= 92) && (rec.BreathOxy) {
			sev += 1
		} else if (rec.OxyBlood >= 87) && (rec.OxyBlood <= 89) && (rec.BreathOxy) {
			sev += 2
		} else if (rec.OxyBlood <= 86) && (rec.BreathOxy) {
			sev += 3
			isLv3 = true
		}
	} else {
		if (rec.OxyBlood >= 1) && (rec.OxyBlood <= 89) {
			sev += 3
			isLv3 = true
		} else if (rec.OxyBlood >= 90) && (rec.OxyBlood <= 92) {
			sev += 2
		} else if (rec.OxyBlood > 92) && (rec.OxyBlood <= 94) {
			sev += 1
		}
	}

	// expression serious
	// if rec.VeryStuffy || rec.PurpleLip {
	// 	isLv4 = true
	// }

	if rec.TiredThanUsual || rec.ChestDiscomfor {
		isLv4 = true
	}

	if (rec.VeryStuffyNew >= 2) || rec.ChestPain || rec.Hemoptisi ||
		rec.UrinaLess || rec.PurpleLip || rec.Stutter || rec.Convulsion {
		return 5
	}

	// if rec.ChestPain {
	// 	sev += 3
	// 	isLv3 = true
	// }

	// if rec.Stutter || rec.Convulsion || rec.UrinaLess {
	// 	return 5
	// }

	log.Printf("Quyen debug Severity %+v", sev)

	ret := 1
	if (sev >= 1) && (sev <= 4) {
		ret = 2
	} else if (sev >= 5) && (sev <= 6) {
		ret = 4
	} else if sev >= 7 {
		ret = 5
	}

	if (isLv3) && (ret < 3) {
		ret = 3
	}

	if (isLv4) && (ret < 4) {
		ret = 4
	}

	log.Printf("Quyen debug Severity level %+v", ret)

	return ret
}

func (s *Storage) AddPatientStatus(in status.PatientStatus) (status.PatientStatus, error) {
	record := PatientStatus{
		PatientUUID:       in.UserUUID,
		Type:              in.Type,
		Serious:           in.Serious,
		SeriousDetail:     in.SeriousDetail,
		OxyBlood:          in.OxyBlood,
		HeartBeat:         in.HeartBeat,
		Temperature:       in.Temperature,
		BloodPressureUp:   in.BloodPressure.UpThrld,
		BloodPressureDown: in.BloodPressure.DownThrld,
		Respiratory:       in.Respiratory.PerMinute,
		Stuffy:            in.Respiratory.Stuffy,
		BreathRetract:     in.Respiratory.BreathRetract,
		BellyBreath:       in.Respiratory.BellyBreath,
		DryCough:          in.Others.DryCough,
		SoreThroat:        in.Others.SoreThroat,
		StuffyNose:        in.Others.StuffyNose,
		Headache:          in.Others.Headache,
		LossTaste:         in.Others.LossTaste,
		RecordNumber:      in.RecordNumber,
		Audio:             in.Audio,
		BreathOxy:         in.BreathOxy,
		Consciousness:     in.Consciousness,
		MachineSeverity:   in.MachineSeverity,
		DoctorSeverity:    in.DoctorSeverity,
		PurpleLip:         in.ExpSerious.PurpleLip,
		ChestPain:         in.ExpSerious.ChestPain,
		Stutter:           in.ExpSerious.Stutter,
		Convulsion:        in.ExpSerious.Convulsion,
		UrinaLess:         in.ExpSerious.UrinaLess,
		Nausea:            in.Others.Nausea,
		Stomachache:       in.Others.Stomachache,
		VeryStuffyNew:     in.ExpSerious.VeryStuffyNew,
	}
	record.MachineSeverity = s.calculateSeverity(record)

	result := s.db.Model(&record).Debug().Create(&record)
	if result.Error != nil {
		return status.PatientStatus{}, result.Error
	}

	// update severity to Account table
	temp := Account{
		Sevirity: record.MachineSeverity,
	}
	db.Model(&temp).Debug().Where(&Account{
		User: record.PatientUUID,
	}).Updates(&temp)

	return in, nil
}

func (s *Storage) UpdatePatientStatus(in status.PatientStatus) (status.PatientStatus, error) {
	record := PatientStatus{
		BaseModel: BaseModel{
			ID: in.ID,
		},
		PatientUUID:       in.UserUUID,
		Type:              in.Type,
		Serious:           in.Serious,
		SeriousDetail:     in.SeriousDetail,
		OxyBlood:          in.OxyBlood,
		HeartBeat:         in.HeartBeat,
		Temperature:       in.Temperature,
		BloodPressureUp:   in.BloodPressure.UpThrld,
		BloodPressureDown: in.BloodPressure.DownThrld,
		Respiratory:       in.Respiratory.PerMinute,
		Stuffy:            in.Respiratory.Stuffy,
		BreathRetract:     in.Respiratory.BreathRetract,
		BellyBreath:       in.Respiratory.BellyBreath,
		DryCough:          in.Others.DryCough,
		SoreThroat:        in.Others.SoreThroat,
		StuffyNose:        in.Others.StuffyNose,
		Headache:          in.Others.Headache,
		LossTaste:         in.Others.LossTaste,
		RecordNumber:      in.RecordNumber,
		Audio:             in.Audio,
		BreathOxy:         in.BreathOxy,
		Consciousness:     in.Consciousness,
		MachineSeverity:   in.MachineSeverity,
		DoctorSeverity:    in.DoctorSeverity,
		PurpleLip:         in.ExpSerious.PurpleLip,
		ChestPain:         in.ExpSerious.ChestPain,
		Stutter:           in.ExpSerious.Stutter,
		Convulsion:        in.ExpSerious.Convulsion,
		UrinaLess:         in.ExpSerious.UrinaLess,
		Nausea:            in.Others.Nausea,
		Stomachache:       in.Others.Stomachache,
		VeryStuffyNew:     in.ExpSerious.VeryStuffyNew,
	}

	db.Model(&record).Debug().Where(&PatientStatus{
		BaseModel: BaseModel{
			ID: in.ID,
		},
	}).Updates(&record)

	// Just use pointer but it cause alot of nil check, so I fixed here
	db.Model(&record).Debug().Where(&PatientStatus{
		BaseModel: BaseModel{
			ID: in.ID,
		},
	}).Updates(map[string]interface{}{
		"stuffy":           in.Respiratory.Stuffy,
		"breath_retract":   in.Respiratory.BreathRetract,
		"belly_breath":     in.Respiratory.BellyBreath,
		"serious":          in.Serious,
		"breath_oxy":       in.BreathOxy,
		"headache":         in.Others.Headache,
		"stomachache":      in.Others.Stomachache,
		"nausea":           in.Others.Nausea,
		"stuffy_nose":      in.Others.StuffyNose,
		"dry_cough":        in.Others.DryCough,
		"sore_throat":      in.Others.SoreThroat,
		"loss_taste":       in.Others.LossTaste,
		"chest_discomfor":  in.Others.ChestDiscomfor,
		"tired_than_usual": in.Others.TiredThanUsual,
		"chest_pain":       in.ExpSerious.ChestPain,
		"hemoptisi":        in.ExpSerious.Hemoptisi,
		"urina_less":       in.ExpSerious.UrinaLess,
		"purple_lip":       in.ExpSerious.PurpleLip,
		"stutter":          in.ExpSerious.Stutter,
		"convulsion":       in.ExpSerious.Convulsion,
	})

	return in, nil
}

func (s *Storage) getPatientStatusRecord(where PatientStatus) ([]PatientStatus, error) {
	var records []PatientStatus

	result := s.db.Debug().Where(&where).Find(&records)

	if result.Error != nil {
		return []PatientStatus{}, fmt.Errorf("Failed to search staff record in database: %v : %v", where, result.Error)
	}

	return records, nil
}

func (s *Storage) GetPatientStatus(in status.PatientStatus) ([]status.PatientStatus, error) {
	var records []PatientStatus
	var ret []status.PatientStatus

	where := PatientStatus{
		RecordNumber: in.RecordNumber,
		PatientUUID:  in.UserUUID,
	}

	records, err := s.getPatientStatusRecord(where)
	if err != nil {
		return []status.PatientStatus{}, err
	}
	// result := s.db.Debug().Where(&where).Find(&records)
	// if result.Error != nil {
	// 	return []status.PatientStatus{}, fmt.Errorf("Failed to search patient record in database: %v : %v", in, result.Error)
	// }

	for _, rec := range records {
		ins := status.PatientStatus{
			ID:              rec.ID,
			MachineSeverity: rec.MachineSeverity,
			DoctorSeverity:  rec.DoctorSeverity,
			UserUUID:        rec.PatientUUID,
			Token:           "",
			Type:            rec.Type,
			Serious:         rec.Serious,
			SeriousDetail:   rec.SeriousDetail,
			OxyBlood:        rec.OxyBlood,
			HeartBeat:       rec.HeartBeat,
			Temperature:     rec.Temperature,
			RecordNumber:    rec.RecordNumber,
			Audio:           rec.Audio,
			CreateAt:        rec.CreatedAt.Format(time.RFC3339),
			BloodPressure: status.BloodPressure{
				UpThrld:   rec.BloodPressureUp,
				DownThrld: rec.BloodPressureDown,
			},
			Respiratory: status.Respiratory{
				PerMinute:     rec.Respiratory,
				Stuffy:        rec.Stuffy,
				BreathRetract: rec.BreathRetract,
				BellyBreath:   rec.BellyBreath,
			},
			Others: status.Others{
				DryCough:       rec.DryCough,
				SoreThroat:     rec.SoreThroat,
				StuffyNose:     rec.StuffyNose,
				Headache:       rec.Headache,
				LossTaste:      rec.LossTaste,
				Stomachache:    rec.Stomachache,
				Nausea:         rec.Nausea,
				ChestDiscomfor: rec.ChestDiscomfor,
				TiredThanUsual: rec.TiredThanUsual,
			},
			Consciousness: rec.Consciousness,
			BreathOxy:     rec.BreathOxy,
			ExpSerious: status.ExpSerious{
				PurpleLip:     rec.PurpleLip,
				ChestPain:     rec.ChestPain,
				Stutter:       rec.Stutter,
				Convulsion:    rec.Convulsion,
				UrinaLess:     rec.UrinaLess,
				Hemoptisi:     rec.Hemoptisi,
				VeryStuffyNew: rec.VeryStuffyNew,
			},
		}
		ret = append(ret, ins)
	}

	return ret, nil
}

func (s *Storage) GetPatientRecordList(patient string) ([]status.StaffRecordList, error) {
	var ret []status.StaffRecordList

	for i := constant.PatientStatus7hMor; i <= constant.PatientStatusEmergency; i++ {
		where := PatientStatus{
			Type:        i,
			PatientUUID: patient,
		}

		records, err := s.getPatientStatusRecord(where)
		if err != nil {
			return []status.StaffRecordList{}, err
		}

		var tempList []int

		for _, rec := range records {
			ins := rec.RecordNumber
			tempList = append(tempList, ins)
		}

		tempInst := status.StaffRecordList{
			Type:          i,
			RecordNumbers: tempList,
		}

		ret = append(ret, tempInst)
	}

	return ret, nil
}
