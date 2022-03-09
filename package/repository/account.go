package repository

import (
	"anti-corona-backend/internal"
	"anti-corona-backend/package/api-process/account"
	"anti-corona-backend/package/constant"
	"fmt"
	"log"

	"github.com/google/uuid"
)

const (
	MAX_PATIENT_PER_DOCTOR = 50
	MAX_DOCTOR_PER_EXPERT  = 20
	AUTO_DOCTOR_PREFIX     = "dr-"
	AUTO_EXPERT_PREFIX     = "ex-"
	ITEM_PER_PAGE          = 20
)

type Account struct {
	BaseModel

	User         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	DoctorUUID   string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	ExpertUUID   string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Passwd       string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	Role         int
	Name         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Addr         string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Phone        string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Email        string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	IDCard       string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Birthday     string
	Gender       bool
	RelPhone     string `sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Sevirity     int
	Discharge    bool
	ReceivedTime string
	Subclinical  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	DiseaseDate  string
	SelfHistory  string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Vaccine      string `sql:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

var accountDf = []Account{
	{User: "system", Passwd: "2af8540c7f949ea2d46fc0259edbda911419b4f989000c985e9af4f44451d9e9", Role: constant.SystemRole},
	{User: "admin", Passwd: "b715e2b3ffe652c66c4e4645b9bb53509d50fce5c7c691abf25f80050d319f1e", Role: constant.AdminRole},
}

func (s *Storage) DescribeUser(user string) (account.Instance, error) {
	var record Account
	var ret account.Instance

	result := s.db.Debug().Where(&Account{
		User: user,
	}).First(&record)
	if result.Error != nil {
		return account.Instance{}, fmt.Errorf("Failed to search user in database: %+v : %+v", user, result.Error)
	}

	ret = account.Instance{
		ID: record.ID,
		Account: account.Account{
			User:   record.User,
			Role:   record.Role,
			Passwd: record.Passwd,
			Token:  "",
		},
		Mgmt: account.Mgmt{
			Name:         record.Name,
			Addr:         record.Addr,
			Phone:        record.Phone,
			Email:        record.Email,
			IDcard:       record.IDCard,
			Birthday:     record.Birthday,
			Gender:       record.Gender,
			RelPhone:     record.RelPhone,
			Sevirity:     record.Sevirity,
			Discharge:    record.Discharge,
			ReceivedTime: record.ReceivedTime,
		},
		MedicalInfo: account.MedicalInfo{
			Subclinical: record.Subclinical,
			DiseaseDate: record.DiseaseDate,
			SelfHistory: record.SelfHistory,
			Vaccine:     record.Vaccine,
			Doctor: account.AccountInfo{
				User: record.DoctorUUID,
			},
			Expert: account.AccountInfo{
				User: record.ExpertUUID,
			},
		},
	}

	return ret, nil
}

func (s *Storage) countDoctorActivePatient(user string) int {
	var records Account
	result := s.db.Debug().Model(&Account{}).
		Where("doctor_uuid = ? AND role = ?", user, constant.PatientRole).Find(&records)
	if result.Error != nil {
		return 0
	}

	log.Printf("Quyen debug: Number of patient %+v", result.RowsAffected)
	return int(result.RowsAffected)
}

func (s *Storage) getDoctorUUID() (string, error) {
	var records []Account
	ret := ""

	result := s.db.Debug().Model(&Account{}).Where(Account{
		Role: constant.AutoDoctorRole,
	}).Find(&records)

	if result.Error != nil {
		return "", result.Error
	}

	for _, rec := range records {
		count := s.countDoctorActivePatient(rec.User)
		if count < MAX_PATIENT_PER_DOCTOR {
			ret = rec.User
			return ret, nil
		}
	}

	uid := uuid.New().String()
	ret = AUTO_DOCTOR_PREFIX + uid
	s.addNewDoctor(ret, uid)

	return ret, nil
}

func (s *Storage) addNewDoctor(user string, pass string) error {
	passwd := internal.SHA256(pass)
	acc := Account{
		User:   user,
		Passwd: passwd,
		Role:   constant.AutoDoctorRole,
	}
	var err error
	acc.ExpertUUID, err = s.getExpertUUID()
	if err != nil {
		return err
	}

	result := s.db.Debug().Model(&acc).FirstOrCreate(&acc, &acc)

	return result.Error
}

func (s *Storage) countExpertActiveDoctor(user string) int {
	var records Account
	result := s.db.Debug().Model(&Account{}).
		Where("expert_uuid = ? AND role = ?", user, constant.AutoDoctorRole).
		Or("expert_uuid = ? AND role = ?", user, constant.DoctorRole).Find(&records)
	if result.Error != nil {
		return 0
	}

	log.Printf("Quyen debug: Number of doctor %+v", result.RowsAffected)
	return int(result.RowsAffected)
}

func (s *Storage) getExpertUUID() (string, error) {
	var records []Account
	ret := ""

	result := s.db.Debug().Model(&Account{}).Where(Account{
		Role: constant.AutoExpertRole,
	}).Find(&records)
	if result.Error != nil {
		return "", result.Error
	}

	for _, rec := range records {
		count := s.countExpertActiveDoctor(rec.User)
		log.Printf("Quyen debug: count doctor: %v", count)
		if count < MAX_DOCTOR_PER_EXPERT {
			ret = rec.User
			return ret, nil
		}
	}

	uid := uuid.New().String()
	ret = AUTO_EXPERT_PREFIX + uid
	err := s.addNewExpert(ret, uid)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (s *Storage) addNewExpert(user string, pass string) error {
	passwd := internal.SHA256(pass)
	acc := Account{
		User:   user,
		Passwd: passwd,
		Role:   constant.AutoExpertRole,
	}

	result := s.db.Debug().Model(&acc).FirstOrCreate(&acc, &acc)

	return result.Error
}

// func (s *Storage) DescribeAccountByRole(role uint) ([]account.Instance, error) {
// 	var records []Account
// 	var ret []account.Instance

// 	result := s.db.Debug().Where(&Account{
// 		Role:      role,
// 		Discharge: false,
// 	}).Find(&records)
// 	if result.Error != nil {
// 		return []account.Instance{}, fmt.Errorf("Failed to search role in database: %+v : %+v", role, result.Error)
// 	}

// 	for _, rec := range records {
// 		ins := account.Instance{
// 			ID:        rec.ID,
// 			User:      rec.User,
// 			Passwd:    rec.Passwd,
// 			Role:      rec.Role,
// 			Name:      rec.Name,
// 			Addr:      rec.Addr,
// 			Token:     "",
// 			Phone:     rec.Phone,
// 			Email:     rec.Email,
// 			IDCard:    rec.IDCard,
// 			Severity:  rec.Severity,
// 			Discharge: rec.Discharge,
// 		}
// 		ret = append(ret, ins)
// 	}

// 	return ret, nil
// }

func (s *Storage) CreateUser(in account.Instance, isCreate bool) error {
	checkAccount, _ := s.DescribeUser(in.Account.User)

	if isCreate {
		if checkAccount.Account.User == in.Account.User {
			return fmt.Errorf("Duplicate")
		}
	}

	user := Account{
		User:         in.Account.User,
		DoctorUUID:   in.MedicalInfo.Doctor.User,
		ExpertUUID:   in.MedicalInfo.Expert.User,
		Passwd:       in.Account.Passwd,
		Name:         in.Mgmt.Name,
		Addr:         in.Mgmt.Addr,
		Phone:        in.Mgmt.Phone,
		Email:        in.Mgmt.Email,
		IDCard:       in.Mgmt.IDcard,
		Gender:       in.Mgmt.Gender,
		RelPhone:     in.Mgmt.RelPhone,
		Discharge:    in.Mgmt.Discharge,
		Subclinical:  in.MedicalInfo.Subclinical,
		SelfHistory:  in.MedicalInfo.SelfHistory,
		Vaccine:      in.MedicalInfo.Vaccine,
		Birthday:     in.Mgmt.Birthday,
		ReceivedTime: in.Mgmt.ReceivedTime,
		DiseaseDate:  in.MedicalInfo.DiseaseDate,
	}

	if isCreate {
		// log.Printf("Quyen debug time birthday: %v\n", in.Mgmt.Birthday)
		// t, err := time.Parse(time.RFC3339, in.Mgmt.Birthday)
		// if err != nil {
		// 	user.Birthday = time.Now()
		// 	log.Printf("failed to parse time format: %v - %v", in.Mgmt.Birthday, err)
		// 	// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.Transfer.Time, err)
		// }
		// user.Birthday = t
		// log.Printf("Quyen debug time birthday parse: %v\n", user.Birthday.Format(time.RFC3339))

		// t, err = time.Parse(time.RFC3339, in.Mgmt.ReceivedTime)
		// if err != nil {
		// 	user.ReceivedTime = time.Now()
		// 	// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.Transfer.Time, err)
		// }
		// user.ReceivedTime = t
		// log.Printf("Quyen debug time recervied time parse: %v\n", user.ReceivedTime.Format(time.RFC3339))

		// t, err = time.Parse(time.RFC3339, in.MedicalInfo.DiseaseDate)
		// if err != nil {
		// 	user.DiseaseDate = time.Now()
		// 	// return patient.PatientMgmt{}, fmt.Errorf("failed to parse time format: %v - %v", in.Transfer.Time, err)
		// }
		// user.DiseaseDate = t
		// log.Printf("Quyen debug time DiseaseDate parse: %v\n", user.DiseaseDate.Format(time.RFC3339))

		// log.Printf("Quyen debug time user parse: %+v\n", user)
		var err error
		user.Role = in.Account.Role

		if user.Role == constant.PatientRole {
			if user.DoctorUUID == "" {
				user.DoctorUUID, err = s.getDoctorUUID()
				if err != nil {
					return err
				}
			} else if len(user.DoctorUUID) > 0 {
				doctorUser, err := s.DescribeUser(user.DoctorUUID)
				if err != nil {
					return err
				}
				if (doctorUser.Account.User == "") || (doctorUser.ID == 0) {
					return fmt.Errorf("Doctor UUID doesn't exist")
				}
			}
		}

		if (user.Role == constant.DoctorRole) || (user.Role == constant.AutoDoctorRole) {
			if user.ExpertUUID == "" {
				user.ExpertUUID, err = s.getExpertUUID()
				if err != nil {
					return err
				}
			}
		}
		user.Sevirity = 0

		result := s.db.Debug().FirstOrCreate(&user, &user)
		if result.Error != nil {
			return fmt.Errorf("Duplicate")
		}
	} else {
		role, _ := internal.GetTokenRole(in.Account.Token)
		if role > uint(checkAccount.Account.Role) {
			return fmt.Errorf("Cannot update account with bigger role")
		}
		result := db.Model(&user).Debug().Where(&Account{
			User: in.Account.User,
		}).Updates(&user)

		if result.Error != nil {
			return result.Error
		}

		// Just use pointer but it cause alot of nil check, so I fixed here
		db.Model(&user).Debug().Where(&Account{
			User: in.Account.User,
		}).Updates(map[string]interface{}{
			"gender": in.Mgmt.Gender,
		})
	}

	return nil
}

func (s *Storage) GetUserByDoctor(doctor string, search string, page int) (account.DescAccount, error) {
	var records []Account
	var ret account.DescAccount
	var dataArray []account.Instance

	where := Account{}
	where.DoctorUUID = doctor
	where.Role = constant.PatientRole

	search = "%" + search + "%"

	result := s.db.Debug().Model(&where).Where("doctor_uuid = ? AND role = ? AND user LIKE ? AND discharge = ?", where.DoctorUUID, where.Role, search, false).
		Or("doctor_uuid = ? AND role = ? AND name LIKE ? AND discharge = ?", where.DoctorUUID, where.Role, search, false).
		Or("doctor_uuid = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.DoctorUUID, where.Role, search, false).Find(&records)

	// result := s.db.Debug().Where(&where).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, fmt.Errorf("Failed to search doctor in database: %+v : %+v", doctor, result.Error)
	}

	ret.TotalRecord = int(result.RowsAffected)
	ret.TotalPage = ret.TotalRecord / ITEM_PER_PAGE
	if (ret.TotalRecord % ITEM_PER_PAGE) != 0 {
		ret.TotalPage += 1
	}
	ret.Page = page + 1

	offset := page * ITEM_PER_PAGE
	result = result.Order("sevirity desc, id").Offset(offset).Limit(ITEM_PER_PAGE).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, result.Error
	}

	for _, rec := range records {
		ins := account.Instance{
			ID: 0,
			Account: account.Account{
				User:   rec.User,
				Role:   rec.Role,
				Passwd: "",
				Token:  "",
			},
			Mgmt: account.Mgmt{
				Name:         rec.Name,
				Addr:         rec.Addr,
				Phone:        rec.Phone,
				Email:        rec.Email,
				IDcard:       rec.IDCard,
				Birthday:     rec.Birthday,
				Gender:       rec.Gender,
				RelPhone:     rec.RelPhone,
				Sevirity:     rec.Sevirity,
				Discharge:    rec.Discharge,
				ReceivedTime: rec.ReceivedTime,
			},
			MedicalInfo: account.MedicalInfo{
				Subclinical: rec.Subclinical,
				DiseaseDate: rec.DiseaseDate,
				SelfHistory: rec.SelfHistory,
				Vaccine:     rec.Vaccine,
				Doctor: account.AccountInfo{
					User: rec.DoctorUUID,
				},
				Expert: account.AccountInfo{
					User: rec.ExpertUUID,
				},
			},
		}

		dataArray = append(dataArray, ins)
	}
	ret.Data = dataArray

	return ret, nil
}

func (s *Storage) GetUserByExpert(expert string, search string, page int) (account.DescAccount, error) {
	var records []Account
	var ret account.DescAccount
	var dataArray []account.Instance

	where := Account{}
	where.ExpertUUID = expert

	search = "%" + search + "%"

	result := s.db.Debug().Model(&where).Where("expert_uuid = ? AND role = ? AND user LIKE ? AND discharge = ?", where.ExpertUUID, constant.DoctorRole, search, false).
		Or("expert_uuid = ? AND role = ? AND name LIKE ? AND discharge = ?", where.ExpertUUID, constant.DoctorRole, search, false).
		Or("expert_uuid = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.ExpertUUID, constant.DoctorRole, search, false).
		Or("expert_uuid = ? AND role = ? AND user LIKE ? AND discharge = ?", where.ExpertUUID, constant.AutoDoctorRole, search, false).
		Or("expert_uuid = ? AND role = ? AND name LIKE ? AND discharge = ?", where.ExpertUUID, constant.AutoDoctorRole, search, false).
		Or("expert_uuid = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.ExpertUUID, constant.AutoDoctorRole, search, false).Find(&records)

	// result := s.db.Debug().Where(&where).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, fmt.Errorf("Failed to search doctor in database: %+v : %+v", expert, result.Error)
	}

	ret.TotalRecord = int(result.RowsAffected)
	ret.TotalPage = ret.TotalRecord / ITEM_PER_PAGE
	if (ret.TotalRecord % ITEM_PER_PAGE) != 0 {
		ret.TotalPage += 1
	}
	ret.Page = page + 1

	offset := page * ITEM_PER_PAGE
	result = result.Order("id").Offset(offset).Limit(ITEM_PER_PAGE).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, result.Error
	}

	for _, rec := range records {
		ins := account.Instance{
			ID: 0,
			Account: account.Account{
				User:   rec.User,
				Role:   rec.Role,
				Passwd: "",
				Token:  "",
			},
			Mgmt: account.Mgmt{
				Name:         rec.Name,
				Addr:         rec.Addr,
				Phone:        rec.Phone,
				Email:        rec.Email,
				IDcard:       rec.IDCard,
				Birthday:     rec.Birthday,
				Gender:       rec.Gender,
				RelPhone:     rec.RelPhone,
				Sevirity:     rec.Sevirity,
				Discharge:    rec.Discharge,
				ReceivedTime: rec.ReceivedTime,
			},
			MedicalInfo: account.MedicalInfo{
				Subclinical: rec.Subclinical,
				DiseaseDate: rec.DiseaseDate,
				SelfHistory: rec.SelfHistory,
				Vaccine:     rec.Vaccine,
				Doctor: account.AccountInfo{
					User: rec.DoctorUUID,
				},
				Expert: account.AccountInfo{
					User: rec.ExpertUUID,
				},
			},
		}

		dataArray = append(dataArray, ins)
	}
	ret.Data = dataArray

	return ret, nil
}

func (s *Storage) GetUserByAdmin(admin string, search string, page int) (account.DescAccount, error) {
	var records []Account
	var ret account.DescAccount
	var dataArray []account.Instance

	where := Account{}
	where.ExpertUUID = admin

	search = "%" + search + "%"

	result := s.db.Debug().Model(&where).Where("role = ? AND user LIKE ? AND discharge = ?", constant.ExpertRole, search, false).
		Or("role = ? AND name LIKE ? AND discharge = ?", constant.ExpertRole, search, false).
		Or("role = ? AND phone LIKE ? AND discharge = ?", constant.ExpertRole, search, false).
		Or("role = ? AND user LIKE ? AND discharge = ?", constant.AutoExpertRole, search, false).
		Or("role = ? AND name LIKE ? AND discharge = ?", constant.AutoExpertRole, search, false).
		Or("role = ? AND phone LIKE ? AND discharge = ?", constant.AutoExpertRole, search, false).Find(&records)

	// result := s.db.Debug().Where(&where).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, fmt.Errorf("Failed to search admin in database: %+v : %+v", admin, result.Error)
	}

	ret.TotalRecord = int(result.RowsAffected)
	ret.TotalPage = ret.TotalRecord / ITEM_PER_PAGE
	if (ret.TotalRecord % ITEM_PER_PAGE) != 0 {
		ret.TotalPage += 1
	}
	ret.Page = page + 1

	offset := page * ITEM_PER_PAGE
	result = result.Order("id").Offset(offset).Limit(ITEM_PER_PAGE).Find(&records)
	if result.Error != nil {
		return account.DescAccount{}, result.Error
	}

	for _, rec := range records {
		ins := account.Instance{
			ID: 0,
			Account: account.Account{
				User:   rec.User,
				Role:   rec.Role,
				Passwd: "",
				Token:  "",
			},
			Mgmt: account.Mgmt{
				Name:         rec.Name,
				Addr:         rec.Addr,
				Phone:        rec.Phone,
				Email:        rec.Email,
				IDcard:       rec.IDCard,
				Birthday:     rec.Birthday,
				Gender:       rec.Gender,
				RelPhone:     rec.RelPhone,
				Sevirity:     rec.Sevirity,
				Discharge:    rec.Discharge,
				ReceivedTime: rec.ReceivedTime,
			},
			MedicalInfo: account.MedicalInfo{
				Subclinical: rec.Subclinical,
				DiseaseDate: rec.DiseaseDate,
				SelfHistory: rec.SelfHistory,
				Vaccine:     rec.Vaccine,
				Doctor: account.AccountInfo{
					User: rec.DoctorUUID,
				},
				Expert: account.AccountInfo{
					User: rec.ExpertUUID,
				},
			},
		}

		dataArray = append(dataArray, ins)
	}
	ret.Data = dataArray

	return ret, nil
}

// func (s *Storage) DescribeAccountByCreator(creator string, search string, role int) ([]account.Instance, error) {
// 	var records []Account
// 	var ret []account.Instance

// 	where := Account{}
// 	where.CreateBy = creator

// 	if (role > constant.MinRole) && (role < constant.MaxRole) {
// 		where.Role = uint(role)
// 	}

// 	search = "%" + search + "%"

// 	result := s.db.Debug().Where("create_by = ? AND role = ? AND user LIKE ? AND discharge = ?", where.CreateBy, where.Role, search, false).
// 		Or("create_by = ? AND role = ? AND name LIKE ? AND discharge = ?", where.CreateBy, where.Role, search, false).
// 		Or("create_by = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.CreateBy, where.Role, search, false).Find(&records)

// 	// result := s.db.Debug().Where(&where).Find(&records)
// 	if result.Error != nil {
// 		return []account.Instance{}, fmt.Errorf("Failed to search create_by in database: %+v : %+v", creator, result.Error)
// 	}
// 	// result = result.Where("user LIKE ?", search).Or("name LIKE ?", search).Or("phone LIKE ?", search)
// 	// if result.Error != nil {
// 	// 	return []account.Instance{}, fmt.Errorf("Failed to search create_by in database: %v : %v", creator, result.Error)
// 	// }

// 	for _, rec := range records {
// 		ins := account.Instance{
// 			ID:        0,
// 			User:      rec.User,
// 			Passwd:    "",
// 			Role:      rec.Role,
// 			Name:      rec.Name,
// 			Addr:      rec.Addr,
// 			Phone:     rec.Phone,
// 			Email:     rec.Email,
// 			IDCard:    rec.IDCard,
// 			Severity:  rec.Severity,
// 			Discharge: rec.Discharge,
// 		}
// 		ret = append(ret, ins)
// 	}

// 	return ret, nil
// }

// func (s *Storage) DescribeUser(user string) (account.Instance, error) {
// 	var record Account
// 	var ret account.Instance

// 	result := s.db.Debug().Where(&Account{
// 		User:      user,
// 		Discharge: false,
// 	}).First(&record)
// 	if result.Error != nil {
// 		return account.Instance{}, fmt.Errorf("Failed to search user in database: %+v : %+v", user, result.Error)
// 	}

// 	ret = account.Instance{
// 		ID:        0,
// 		User:      record.User,
// 		Passwd:    "",
// 		Role:      record.Role,
// 		Token:     "",
// 		Name:      record.Name,
// 		Phone:     record.Phone,
// 		Addr:      record.Addr,
// 		Email:     record.Email,
// 		IDCard:    record.IDCard,
// 		CreateBy:  record.CreateBy,
// 		Severity:  record.Severity,
// 		Discharge: record.Discharge,
// 	}

// 	return ret, nil
// }

// func (s *Storage) DescribePatientByClinic(clinic string, search string) ([]account.Instance, error) {
// 	var records []Account
// 	var ret []account.Instance
// 	where := Account{}
// 	where.ClinicID = clinic

// 	search = "%" + search + "%"

// 	result := s.db.Debug().Where("clinic_id = ? AND role = ? AND user LIKE ? AND discharge = ?", where.ClinicID, constant.PatientRole, search, false).
// 		Or("clinic_id = ? AND role = ? AND name LIKE ? AND discharge = ?", where.ClinicID, constant.PatientRole, search, false).
// 		Or("clinic_id = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.ClinicID, constant.PatientRole, search, false).Find(&records)

// 	// result := s.db.Debug().Where(&where).Find(&records)
// 	if result.Error != nil {
// 		return []account.Instance{}, fmt.Errorf("Failed to search create_by in database: %+v : %+v", clinic, result.Error)
// 	}

// 	for _, rec := range records {
// 		ins := account.Instance{
// 			ID:        0,
// 			User:      rec.User,
// 			Passwd:    "",
// 			Role:      rec.Role,
// 			Name:      rec.Name,
// 			Addr:      rec.Addr,
// 			Phone:     rec.Phone,
// 			Email:     rec.Email,
// 			IDCard:    rec.IDCard,
// 			Severity:  rec.Severity,
// 			Discharge: rec.Discharge,
// 		}
// 		ret = append(ret, ins)
// 	}

// 	return ret, nil
// }

// func (s *Storage) DescribePatientByStaff(staff string, search string, role uint) ([]account.Instance, error) {
// 	var ret []account.Instance

// 	data, err := s.descPatientMgmtByStaff(staff, search, role)
// 	if err != nil {
// 		return []account.Instance{}, err
// 	}

// 	for _, rec := range data {
// 		acc, err := s.DescribeAccount(rec.PatientUUID)
// 		if err != nil {
// 			continue
// 		}
// 		if acc.Discharge || (acc.User == "") {
// 			continue
// 		}
// 		acc.Passwd = ""
// 		ret = append(ret, acc)
// 	}

// 	sort.SliceStable(ret, func(i, j int) bool {
// 		return ret[i].Severity > ret[j].Severity
// 	})

// 	// log.Printf("Quyen Debug: return slice after sort: %+v", ret)

// 	return ret, nil
// }

// func (s *Storage) DescribePatientByAdminst(creator string, search string) ([]account.Instance, error) {
// 	var records []Account
// 	var ret []account.Instance
// 	where := Account{}
// 	where.CreateBy = creator

// 	search = "%" + search + "%"

// 	result := s.db.Debug().Where("clinic_id = ? AND role = ? AND user LIKE ? AND discharge = ?", where.CreateBy, constant.PatientRole, search, false).
// 		Or("clinic_id = ? AND role = ? AND name LIKE ? AND discharge = ?", where.CreateBy, constant.PatientRole, search, false).
// 		Or("clinic_id = ? AND role = ? AND phone LIKE ? AND discharge = ?", where.CreateBy, constant.PatientRole, search, false).
// 		Order("severity desc, id").
// 		Find(&records)

// 	// result := s.db.Debug().Where(&where).Find(&records)
// 	if result.Error != nil {
// 		return []account.Instance{}, fmt.Errorf("Failed to search create_by in database: %+v : %+v", creator, result.Error)
// 	}

// 	for _, rec := range records {
// 		ins := account.Instance{
// 			ID:        0,
// 			User:      rec.User,
// 			Passwd:    "",
// 			Role:      rec.Role,
// 			Name:      rec.Name,
// 			Addr:      rec.Addr,
// 			Phone:     rec.Phone,
// 			Email:     rec.Email,
// 			IDCard:    rec.IDCard,
// 			Severity:  rec.Severity,
// 			Discharge: rec.Discharge,
// 		}
// 		ret = append(ret, ins)
// 	}

// 	return ret, nil
// }
