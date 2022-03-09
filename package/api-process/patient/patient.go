package patient

import "anti-corona-backend/package/api-process/account"

type Patient struct {
	Gender       string `json:"gender,omitempty"`
	Ethnicity    string `json:"ethnicity,omitempty"`
	HIType       int    `json:"hi_type,omitempty" `
	HIExpire     string `json:"hi_expire,omitempty"`
	HINumber     string `json:"hi_number,omitempty"`
	GoogleMap    string `json:"google_map,omitempty"`
	Occupations  string `json:"occupations,omitempty"`
	Nationality  string `json:"nationality,omitempty"`
	BirthDay     string `json:"birthday,omitempty"`
	OtherContact string `json:"other_contact,omitempty"`
}

type Relative struct {
	Name         string `json:"name,omitempty"`
	Addr         string `json:"addr,omitempty"`
	Phone        string `json:"phone,omitempty"`
	GoogleMap    string `json:"google_map,omitempty"`
	OtherContact string `json:"other_contact,omitempty"`
}

type PatientInfo struct {
	UserUUID string   `json:"user_uuid,omitempty"`
	Token    string   `json:"token,omitempty"`
	Patient  Patient  `json:"patient,omitempty"`
	Relative Relative `json:"relative,omitempty"`
}

type Status struct {
	Time   string `json:"time,omitempty"`
	Note   string `json:"note,omitempty"`
	Status bool   `json:"status,omitempty"`
}

type PatientMgmt struct {
	UserUUID       string           `json:"user_uuid,omitempty"`
	Token          string           `json:"token,omitempty"`
	ReceivedTime   string           `json:"received_time,omitempty"`
	TreatmentPlace string           `json:"treatment_place,omitempty"`
	TreatmentMap   string           `json:"treatment_map,omitempty"`
	Doctor         account.Instance `json:"doctor,omitempty"`
	Administrative account.Instance `json:"administrative,omitempty"`
	Healthcare     account.Instance `json:"healthcare,omitempty"`
	Supply         account.Instance `json:"supply,omitempty"`
	Transfer       Status           `json:"transfer,omitempty"`
	StopTreatment  Status           `json:"stop_treatment,omitempty"`
	EndTreatment   Status           `json:"end_treatment,omitempty"`
	Dead           Status           `json:"dead,omitempty"`
}

type Characteristics struct {
	Allergy   string `json:"allergy,omitempty"`
	Medicine  string `json:"medicine,omitempty"`
	Cigarette string `json:"cigarette,omitempty"`
	Alcohol   string `json:"alcohol,omitempty"`
	Drug      string `json:"drug,omitempty"`
	Others    string `json:"others,omitempty"`
}

type Epidemiology struct {
	Environment  string `json:"environment,omitempty"`
	AcuteIllness string `json:"acute_illness,omitempty"`
	Place        string `json:"place,omitempty"`
}

type MedicalHistory struct {
	Self            string          `json:"self,omitempty"`
	SelfNote        string          `json:"self_note,omitempty"`
	Characteristics Characteristics `json:"characteristics,omitempty"`
	Family          string          `json:"family,omitempty"`
	Epidemiology    Epidemiology    `json:"epidemiology,omitempty"`
}

type AskPatient struct {
	PathologicalProcess string         `json:"pathological_process,omitempty"`
	MedicalHistory      MedicalHistory `json:"medical_history,omitempty"`
}

type Body struct {
	Note   string `json:"note,omitempty"`
	Height int    `json:"height,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

type Agencies struct {
	Cyclic          string `json:"cyclic,omitempty"`
	Respiratory     string `json:"respiratory,omitempty"`
	Digest          string `json:"digest,omitempty"`
	Kidney          string `json:"kidney,omitempty"`
	Nerve           string `json:"nerve,omitempty"`
	Musculoskeletal string `json:"musculoskeletal,omitempty"`
	Endocrine       string `json:"endocrine,omitempty"`
	Ent             string `json:"ent,omitempty"`
	Nutrition       string `json:"nutrition,omitempty"`
	Others          string `json:"others,omitempty"`
}

type TestData struct {
	Name   string `json:"name,omitempty"`
	Result string `json:"result,omitempty"`
	Image  string `json:"image,omitempty"`
}

type Subclinical struct {
	FastTest TestData `json:"fast_test,omitempty"`
	PCRTest  TestData `json:"pcr_test,omitempty"`
	Others   TestData `json:"others,omitempty"`
}

type Examination struct {
	Body        Body     `json:"body,omitempty"`
	Agencies    Agencies `json:"agencies,omitempty"`
	Subclinical string   `json:"subclinical,omitempty"`
	// Subclinical Subclinical `json:"subclinical,omitempty"`
}

type Diagnostic struct {
	MainDisease string `json:"main_disease,omitempty"`
	SideDisease string `json:"side_disease,omitempty"`
	Distinguish string `json:"distinguish,omitempty"`
}

type Summary struct {
	Discharge string `json:"discharge,omitempty"`
	TreatDir  string `json:"treat_dir,omitempty"`
}

type MedicalData struct {
	UserUUID     string      `json:"user_uuid,omitempty"`
	Token        string      `json:"token,omitempty"`
	Reason       string      `json:"reason,omitempty"`
	ReceivedTime string      `json:"received_time,omitempty"`
	Ask          AskPatient  `json:"ask,omitempty"`
	Examination  Examination `json:"examination,omitempty"`
	Diagnostic   Diagnostic  `json:"diagnostic,omitempty"`
	Summary      Summary     `json:"summary,omitempty"`
	Prognosis    string      `json:"prognosis,omitempty"`
	Treatments   string      `json:"treatments,omitempty"`
	Vaccination  string      `json:"vaccination,omitempty"`
}

type Rest interface {
	// UpdatePatientInfo(PatientInfo) (PatientInfo, error)
	// GetPatientInfo(string) (PatientInfo, error)
	// UpdatePatientMgmt(PatientMgmt) (PatientMgmt, error)
	// GetPatientMgmt(string) (PatientMgmt, error)
	// UpdatePatientMedical(MedicalData) (MedicalData, error)
	// GetPatientMedical(string) (MedicalData, error)
}
