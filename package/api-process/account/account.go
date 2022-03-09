package account

const ()

type Account struct {
	User   string `json:"user"`
	Role   int    `json:"role"`
	Passwd string `json:"passwd"`
	Token  string `json:"token"`
}

type Mgmt struct {
	Name         string `json:"name,omitempty"`
	Addr         string `json:"addr,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	IDcard       string `json:"idcard,omitempty"`
	Birthday     string `json:"birthday,omitempty"`
	Gender       bool   `json:"gender,omitempty"`
	RelPhone     string `json:"rel_phone,omitempty"`
	Sevirity     int    `json:"sevirity,omitempty"`
	Discharge    bool   `json:"discharge,omitempty"`
	ReceivedTime string `json:"received_time,omitempty"`
}

type AccountInfo struct {
	User  string `json:"user,omitempty"`
	Name  string `json:"name,omitempty"`
	Addr  string `json:"addr,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type MedicalInfo struct {
	Subclinical string      `json:"subclinical,omitempty"`
	DiseaseDate string      `json:"disease_date,omitempty"`
	SelfHistory string      `json:"self_history,omitempty"`
	Vaccine     string      `json:"vaccine,omitempty"`
	Doctor      AccountInfo `json:"doctor,omitempty"`
	Expert      AccountInfo `json:"expert,omitempty"`
}

type Instance struct {
	ID          uint        `json:"id"`
	Account     Account     `json:"account"`
	Mgmt        Mgmt        `json:"mgmt"`
	MedicalInfo MedicalInfo `json:"medical_info"`
}

type DescAccount struct {
	TotalPage   int        `json:"total_page,omitempty"`
	TotalRecord int        `json:"total_record,omitempty"`
	Page        int        `json:"page,omitempty"`
	Data        []Instance `json:"data,omitempty"`
}

type Rest interface {
	Create(Instance, bool) (Instance, error)
	Delete(Instance) (Instance, error)
	DescribeUser(string) (Instance, error)
	Authenticate(Account) (Account, error)
	GetUserByDoctor(doctor string, search string, page int) (DescAccount, error)
	GetUserByExpert(doctor string, search string, page int) (DescAccount, error)
	GetUserByAdmin(doctor string, search string, page int) (DescAccount, error)
	// GetRoleByToken(token string) (Instance, error)
	// DescribeByCreator(string, string, int) (DescAccount, error)
	// DescribeUser(string) (Instance, error)
	// DescribePatientByClinic(string, string) (DescAccount, error)
	// DescribePatientByStaff(string, string, uint) (DescAccount, error)
	// DescribePatientByAdminst(string, string) (DescAccount, error)
}
