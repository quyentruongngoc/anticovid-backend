package constant

const (
	// SystemRole     = "system"
	// AdminRole      = "admin"
	// ClinicRole     = "clinic"
	// DoctorRole     = "doctor"
	// SupplyRole     = "supply"
	// HealthcareRole = "healthcare"
	// PatientRole    = "patient"
	MinRole            = 0
	SystemRole         = 1 // system
	AdminRole          = 2 // admin
	AutoExpertRole     = 3 // clinic
	DoctorRole         = 4 // doctor
	ExpertRole         = 5 // supply
	AutoDoctorRole     = 6 // healthcare
	AdministrativeRole = 7 // administrative
	PatientRole        = 8 // patient
	MaxRole            = 9

	PatientStatus7hMor     = 1
	PatientStatus7hEve     = 2
	PatientStatusEmergency = 3

	CommandAutoDoctor = 1
	CommandExpert     = 2
	CommandPatient    = 3
	ReportAutoDoctor  = 4
	ReportExpert      = 5
	ReportPatient     = 6
)
