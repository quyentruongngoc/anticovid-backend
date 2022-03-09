package status

type BloodPressure struct {
	UpThrld   int `json:"up_thrld,omitempty"`
	DownThrld int `json:"down_thrld,omitempty"`
}

type Respiratory struct {
	PerMinute     int  `json:"per_minute,omitempty"`
	Stuffy        bool `json:"stuffy,omitempty"`
	BreathRetract bool `json:"breath_retract,omitempty"`
	BellyBreath   bool `json:"belly_breath,omitempty"`
}

type Others struct {
	Headache       bool `json:"headache,omitempty"`         // 1. dau dau
	Stomachache    bool `json:"stomachache,omitempty"`      // 2. dau bung/ tieu chay
	Nausea         bool `json:"nausea,omitempty"`           // 3. buon non
	StuffyNose     bool `json:"stuffy_nose,omitempty"`      // 4. nghet mui/ so mui
	DryCough       bool `json:"dry_cough,omitempty"`        // 5. ho khang
	SoreThroat     bool `json:"sore_throat,omitempty"`      // 6. dau hong
	LossTaste      bool `json:"loss_taste,omitempty"`       // 7. mat vi giac/ khuu giac
	ChestDiscomfor bool `json:"chest_discomfor,omitempty"`  // 8. kho chiu trong long nguc
	TiredThanUsual bool `json:"tired_than_usual,omitempty"` // 9. Van dong thay met hon moi khi (lv4)
}

type ExpSerious struct {
	ChestPain     bool `json:"chest_pain,omitempty"` // 10. dau ben trong long nguc
	Hemoptisi     bool `json:"hemoptisi,omitempty"`  // 11. ho ra dam hong
	UrinaLess     bool `json:"urina_less,omitempty"` // 12. khong tieu hoac tieu rat it trong 24h
	PurpleLip     bool `json:"purple_lip,omitempty"` // 13. moi tim tai
	Stutter       bool `json:"stutter,omitempty"`    // 14. Xuat hien noi lap
	Convulsion    bool `json:"convulsion,omitempty"` // 15. Co giat
	VeryStuffyNew int  `json:"very_stuffy_new"`
}

type Oxy struct {
	Concentration    int `json:"concentration,omitempty"`
	Volume           int `json:"oxy_volume,omitempty"`
	TimesPerDay      int `json:"times_per_day,omitempty"`
	DurationPerTimes int `json:"dura_per_times,omitempty"`
}

type PatientStatus struct {
	UserUUID        string        `json:"user_uuid,omitempty"`
	Token           string        `json:"token,omitempty"`
	Type            int           `json:"type,omitempty"`
	Serious         bool          `json:"serious,omitempty"`
	SeriousDetail   string        `json:"serious_detail,omitempty"`
	OxyBlood        int           `json:"oxy_blood,omitempty"`
	HeartBeat       int           `json:"heart_beat,omitempty"`
	Temperature     float32       `json:"temperature,omitempty"`
	RecordNumber    int           `json:"record_number,omitempty"`
	BloodPressure   BloodPressure `json:"blood_pressure,omitempty"`
	Respiratory     Respiratory   `json:"respiratory,omitempty"`
	Others          Others        `json:"others,omitempty"`
	ExpSerious      ExpSerious    `json:"exp_serious,omitempty"`
	Audio           string        `json:"audio,omitempty"`
	CreateAt        string        `json:"create_at,omitempty"`
	MachineSeverity int           `json:"machine_severity,omitempty"`
	DoctorSeverity  int           `json:"doctor_severity,omitempty"`
	ID              uint          `json:"id,omitempty"`
	BreathOxy       bool          `json:"breath_oxy,omitempty"`
	Consciousness   int           `json:"consciousness,omitempty"`
}

type StaffStatus struct {
	Note         string `json:"note,omitempty"`
	CreateAt     string `json:"create_at,omitempty"`
	RecordNumber int    `json:"record_number,omitempty"`
	UserUUID     string `json:"user_uuid,omitempty"`
	PatientUUID  string `json:"patient_uuid,omitempty"`
	Token        string `json:"token,omitempty"`
	Type         int    `json:"type,omitempty"`
	ParentID     int    `json:"parent_id,omitempty"`
	ID           int    `json:"id,omitempty"`
	Role         int    `json:"-"`
}

type RecordData struct {
	RecordNumber int `json:"record_number,omitempty"`
	RecordType   int `json:"type,omitempty"`
}

type StaffRecordList struct {
	Role          int          `json:"role,omitempty"`
	Type          int          `json:"type,omitempty"`
	RecordData    []RecordData `json:"record_numbers,omitempty"`
	RecordNumbers []int        `json:"record_number,omitempty"`
}

type Rest interface {
	AddStatusRecord(StaffStatus) (StaffStatus, error)
	GetStatusRecord(StaffStatus) ([]StaffStatus, error)
	AddPatientStatus(PatientStatus) (PatientStatus, error)
	UpdatePatientStatus(PatientStatus) (PatientStatus, error)
	GetPatientStatus(PatientStatus) ([]PatientStatus, error)
	GetStaffRecordList(string) ([]StaffRecordList, error)
	GetPatientRecordList(string) ([]StaffRecordList, error)
}
