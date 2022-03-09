package account

import (
	"anti-corona-backend/internal"
	"anti-corona-backend/package/api-process/account"
	"fmt"
	"log"
	"time"
)

type Repository interface {
	// DescribeAccount(string) (account.Instance, error)
	// DescribeAccountByCreator(string, string, int) ([]account.Instance, error)
	CreateUser(account.Instance, bool) error
	DescribeUser(string) (account.Instance, error)
	GetUserByDoctor(string, string, int) (account.DescAccount, error)
	GetUserByExpert(string, string, int) (account.DescAccount, error)
	GetUserByAdmin(string, string, int) (account.DescAccount, error)
	// DescribePatientByClinic(string, string) ([]account.Instance, error)
	// DescribePatientByStaff(string, string, uint) ([]account.Instance, error)
	// DescribePatientByAdminst(string, string) ([]account.Instance, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(in account.Instance, isCreate bool) (account.Instance, error) {
	if (len(in.Account.Passwd) == 0) && (isCreate) {
		return account.Instance{}, fmt.Errorf("Password should not empty")
	}

	if len(in.Account.Passwd) > 0 {
		in.Account.Passwd = internal.SHA256(in.Account.Passwd)
	} else {
		in.Account.Passwd = ""
	}

	err := s.repo.CreateUser(in, isCreate)
	if err != nil {
		return account.Instance{}, err
	}

	return in, nil
}

func (s *Service) Delete(in account.Instance) (account.Instance, error) {
	return account.Instance{}, nil
}

func (s *Service) Authenticate(in account.Account) (account.Account, error) {
	instance, err := s.repo.DescribeUser(in.User)
	if err != nil {
		log.Printf("Failed to get data from database: %v", err)
		return account.Account{}, err
	}
	hashPw := internal.SHA256(in.Passwd)

	if instance.Account.Passwd == hashPw {
		now := time.Now().Nanosecond()
		temp := fmt.Sprintf("%v-%v", instance.Account.User, now)
		log.Printf("Quyen debug Token before hash: %v\n", temp)
		instance.Account.Token = internal.SHA256(temp)
		log.Printf("Generate new Token: %v\n", instance.Account.Token)
		// update token to list
		internal.AddToTokenList(instance.Account.Token, uint(instance.Account.Role), instance.Account.User)

		return account.Account{
			User:   instance.Account.User,
			Role:   instance.Account.Role,
			Passwd: "",
			Token:  instance.Account.Token,
		}, nil
	}

	return account.Account{}, fmt.Errorf("Password not correct")
}

func (s *Service) GetUserByDoctor(doctor string, search string, page int) (account.DescAccount, error) {
	return s.repo.GetUserByDoctor(doctor, search, page)
}

func (s *Service) GetUserByExpert(expert string, search string, page int) (account.DescAccount, error) {
	return s.repo.GetUserByExpert(expert, search, page)
}

func (s *Service) GetUserByAdmin(admin string, search string, page int) (account.DescAccount, error) {
	return s.repo.GetUserByAdmin(admin, search, page)
}

func (s *Service) DescribeUser(user string) (account.Instance, error) {
	instance, err := s.repo.DescribeUser(user)
	if err != nil {
		return account.Instance{}, err
	}

	instance.Account.Passwd = ""
	instance.Account.Token = ""

	return instance, nil
}

// func (s *Service) GetRoleByToken(token string) (account.Instance, error) {
// 	var instance account.Instance

// 	role, err := internal.GetTokenRole(token)
// 	if err != nil {
// 		return account.Instance{}, err
// 	}

// 	user, err := internal.GetTokenUser(token)
// 	if err != nil {
// 		return account.Instance{}, err
// 	}

// 	instance.Role = role
// 	instance.User = user

// 	return instance, nil
// }

// func (s *Service) DescribeByCreator(creator string, search string, role int) (account.DescAccount, error) {
// 	insts, err := s.repo.DescribeAccountByCreator(creator, search, role)
// 	if err != nil {
// 		return account.DescAccount{}, err
// 	}
// 	ret := account.DescAccount{
// 		Data: insts,
// 	}

// 	return ret, nil
// }

// func (s *Service) DescribeUser(user string) (account.Instance, error) {
// 	instance, err := s.repo.DescribeUser(user)
// 	if err != nil {
// 		return account.Instance{}, err
// 	}

// 	return instance, nil
// }

// func (s *Service) DescribePatientByClinic(clinic string, search string) (account.DescAccount, error) {
// 	insts, err := s.repo.DescribePatientByClinic(clinic, search)
// 	if err != nil {
// 		return account.DescAccount{}, err
// 	}
// 	ret := account.DescAccount{
// 		Data: insts,
// 	}

// 	return ret, nil
// }

// func (s *Service) DescribePatientByStaff(user string, search string, role uint) (account.DescAccount, error) {
// 	insts, err := s.repo.DescribePatientByStaff(user, search, role)
// 	if err != nil {
// 		return account.DescAccount{}, err
// 	}
// 	ret := account.DescAccount{
// 		Data: insts,
// 	}

// 	return ret, nil
// }

// func (s *Service) DescribePatientByAdminst(user string, search string) (account.DescAccount, error) {
// 	insts, err := s.repo.DescribePatientByAdminst(user, search)
// 	if err != nil {
// 		return account.DescAccount{}, err
// 	}
// 	ret := account.DescAccount{
// 		Data: insts,
// 	}

// 	return ret, nil
// }
