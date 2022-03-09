package account

import (
	"anti-corona-backend/internal"
	"anti-corona-backend/package/api-process/account"
	"anti-corona-backend/package/constant"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	TOKEN   = "token"
	PAGE    = "page"
	USER    = "user"
	CREATOR = "creator"
	ROLE    = "role"
	SEARCH  = "search"
	DOCTOR  = "doctor"
	EXPERT  = "expert"

	ITEM_PER_PAGE = 20
)

func Install(r *mux.Router, s account.Rest) {
	r.HandleFunc("/auth", authticateHandler(s)).Methods(http.MethodPost, http.MethodHead)

	// r.HandleFunc("/role/{token}", getRoleByToken(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/user/create", createUser(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/user/create", updateUser(s)).Methods(http.MethodPut, http.MethodHead)
	// r.HandleFunc("/user/{user}", getUser(s)).Methods(http.MethodGet, http.MethodHead)
	// r.HandleFunc("/user/creator/{page}", getUserByCreator(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/user/patient/clinic/{page}", getPatientByClinic(s)).Methods(http.MethodGet, http.MethodHead)
	// r.HandleFunc("/user/patient/staff/{page}", getPatientByStaff(s)).Methods(http.MethodGet, http.MethodHead)
	// r.HandleFunc("/user/patient/administ/{page}", getPatientByAdminist(s)).Methods(http.MethodGet, http.MethodHead)

	r.HandleFunc("/user/create", newCreateUser(s)).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/user/create", newUpdateUser(s)).Methods(http.MethodPut, http.MethodHead)
	r.HandleFunc("/user/{user}", getUser(s)).Methods(http.MethodGet, http.MethodHead)

	r.HandleFunc("/doctor/patient/{page}", getUserByDoctor(s)).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/expert/doctor/{page}", getUserByExpert(s)).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/admin/expert/{page}", getUserByAdmin(s)).Methods(http.MethodGet, http.MethodHead)
}

type errorRsp struct {
	Msg string
}

func writeRsp(w http.ResponseWriter, err error, rsp interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	if encError := json.NewEncoder(w).Encode(rsp); encError != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseReqDatas(r *http.Request) (account.Instance, error) {
	decoder := json.NewDecoder(r.Body)
	var instance account.Instance
	err := decoder.Decode(&instance)
	return instance, err
}

func parseAccoutDatas(r *http.Request) (account.Account, error) {
	decoder := json.NewDecoder(r.Body)
	var instance account.Account
	err := decoder.Decode(&instance)
	return instance, err
}

func authticateHandler(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		instance, err := parseAccoutDatas(r)
		if err != nil {
			log.Println("Authenticate failed: unable to parse user input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		instance, err = s.Authenticate(instance)
		if err == nil {
			instance.Passwd = ""
			writeRsp(w, nil, instance, http.StatusOK)
		} else {
			log.Printf("Failed to authenticate with data: %v - %v", instance, err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}

	}
}

// func getRoleByToken(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		token := mux.Vars(r)[TOKEN]
// 		log.Printf("Get role for token: %v\n", token)

// 		instance, err := s.GetRoleByToken(token)
// 		if err == nil {
// 			instance.ID = 0
// 			instance.Passwd = ""
// 			writeRsp(w, nil, instance, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get role: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func createUser(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		instance, err := parseReqDatas(r)
// 		if err != nil {
// 			log.Println("Create failed: unable to parse user inpu: ", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		log.Printf("Quyen debug body: %v\n", instance)

// 		if !internal.IsTokenExist(instance.Token) {
// 			log.Println("Create failed: unable to get token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		role, err := internal.GetTokenRole(instance.Token)
// 		if err != nil {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid role"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		if (role == constant.SystemRole) || (role == constant.AdminRole) || (role == constant.ClinicRole) || (role == constant.AdministrativeRole) {
// 			instance, err = s.Create(instance, true)
// 			if err == nil {
// 				writeRsp(w, nil, fmt.Sprintf("Create sucessfully"), http.StatusOK)
// 			} else {
// 				instance.Passwd = ""
// 				log.Printf("Failed to authenticate with data: %v - %v", instance, err)
// 				writeRsp(w, err, &errorRsp{
// 					Msg: fmt.Sprintf("%v", err),
// 				}, http.StatusNotFound)
// 			}
// 		} else {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid role"),
// 			}, http.StatusBadRequest)
// 		}

// 	}
// }

// func updateUser(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		instance, err := parseReqDatas(r)
// 		if err != nil {
// 			log.Println("update failed: unable to parse user inpu: ", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		if len(instance.User) == 0 {
// 			log.Println("user empty: ", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("User should not empty"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		log.Printf("Quyen debug body: %+v\n", instance)

// 		if !internal.IsTokenExist(instance.Token) {
// 			log.Println("update failed: unable to get token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		role, err := internal.GetTokenRole(instance.Token)
// 		if err != nil {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid role"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		if (role == constant.SystemRole) || (role == constant.AdminRole) || (role == constant.ClinicRole) || (role == constant.AdministrativeRole) {
// 			instance, err = s.Create(instance, false)
// 			if err == nil {
// 				writeRsp(w, nil, fmt.Sprintf("update sucessfully"), http.StatusOK)
// 			} else {
// 				instance.Passwd = ""
// 				log.Printf("Failed to update with data: %v - %v", instance, err)
// 				writeRsp(w, err, &errorRsp{
// 					Msg: fmt.Sprintf("%v", err),
// 				}, http.StatusNotFound)
// 			}
// 		} else {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid role"),
// 			}, http.StatusBadRequest)
// 		}
// 	}
// }

// func getUserByCreator(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		spage := mux.Vars(r)[PAGE]
// 		page, err := strconv.Atoi(spage)
// 		if err != nil {
// 			log.Println("Get user failed: unable to parse page input")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		log.Printf("Process data for page: %v\n", page)

// 		q := r.URL.Query()
// 		creator := q.Get(CREATOR)
// 		token := q.Get(TOKEN)
// 		srole := q.Get(ROLE)
// 		role, _ := strconv.Atoi(srole)
// 		search := q.Get(SEARCH)

// 		log.Printf("Get user create by %v\n", creator)

// 		if len(creator) == 0 {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("creator params should not empty"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		if !internal.IsTokenExist(token) {
// 			log.Println("Create failed: unable to verify token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		internal.UpdateToken(token)

// 		ret, err := s.DescribeByCreator(creator, search, role)
// 		if err == nil {
// 			ret.Page = page
// 			if page == 0 {
// 				page = 1
// 			}
// 			arrayLen := len(ret.Data)

// 			ret.TotalRecord = arrayLen
// 			ret.TotalPage = arrayLen / ITEM_PER_PAGE
// 			mod := arrayLen % ITEM_PER_PAGE
// 			if mod != 0 {
// 				ret.TotalPage += 1
// 			}
// 			log.Printf("Quyen debug: TotalPage: %v", ret.TotalPage)

// 			var start int
// 			var end int

// 			start = (page - 1) * ITEM_PER_PAGE
// 			if page == ret.TotalPage {
// 				end = arrayLen
// 			} else {
// 				end = (page-1)*ITEM_PER_PAGE + ITEM_PER_PAGE
// 			}

// 			if end > arrayLen {
// 				end = arrayLen
// 			}

// 			log.Printf("Quyen debug: arraylen: %v", arrayLen)
// 			log.Printf("Quyen debug: start: %v", start)
// 			log.Printf("Quyen debug: end: %v", end)

// 			ret.Data = ret.Data[start:end]

// 			internal.UpdateToken(token)
// 			writeRsp(w, nil, ret, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get role: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getUser(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		user := mux.Vars(r)[USER]
// 		log.Printf("Process data for user: %v\n", user)

// 		q := r.URL.Query()
// 		token := q.Get("token")

// 		if !internal.IsTokenExist(token) {
// 			log.Println("Create failed: unable to verify token")
// 			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		internal.UpdateToken(token)

// 		// tUser, err := internal.GetTokenUser(token)
// 		// if err != nil {
// 		// 	log.Println("Failed to get token user")
// 		// 	http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 		// 	return
// 		// }

// 		// if user != tUser {
// 		// 	log.Println("User and token not match")
// 		// 	http.Error(w, fmt.Sprintf("Invalid user"), http.StatusBadRequest)
// 		// 	return
// 		// }

// 		ret, err := s.DescribeUser(user)
// 		if err == nil {
// 			log.Printf("Quyen debug data return: %v\n", ret)
// 			writeRsp(w, nil, ret, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get user: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPatientByClinic(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		spage := mux.Vars(r)[PAGE]
// 		page, err := strconv.Atoi(spage)
// 		if err != nil {
// 			log.Println("Get user failed: unable to parse page input")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		log.Printf("Process data for page: %v\n", page)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)
// 		search := q.Get(SEARCH)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("get failed: unable to verify token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		user, err := internal.GetTokenUser(token)
// 		if err != nil {
// 			log.Printf("get failed: unable to get user from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		internal.UpdateToken(token)

// 		ret, err := s.DescribePatientByClinic(user, search)
// 		if err == nil {
// 			ret.Page = page
// 			if page == 0 {
// 				page = 1
// 			}
// 			arrayLen := len(ret.Data)

// 			ret.TotalRecord = arrayLen
// 			ret.TotalPage = arrayLen / ITEM_PER_PAGE
// 			mod := arrayLen % ITEM_PER_PAGE
// 			if mod != 0 {
// 				ret.TotalPage += 1
// 			}
// 			log.Printf("Quyen debug: TotalPage: %v", ret.TotalPage)

// 			var start int
// 			var end int

// 			start = (page - 1) * ITEM_PER_PAGE
// 			if page == ret.TotalPage {
// 				end = arrayLen
// 			} else {
// 				end = (page-1)*ITEM_PER_PAGE + ITEM_PER_PAGE
// 			}

// 			if end > arrayLen {
// 				end = arrayLen
// 			}

// 			log.Printf("Quyen debug: arraylen: %v", arrayLen)
// 			log.Printf("Quyen debug: start: %v", start)
// 			log.Printf("Quyen debug: end: %v", end)

// 			ret.Data = ret.Data[start:end]

// 			internal.UpdateToken(token)
// 			writeRsp(w, nil, ret, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get role: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPatientByStaff(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		spage := mux.Vars(r)[PAGE]
// 		page, err := strconv.Atoi(spage)
// 		if err != nil {
// 			log.Println("Get user failed: unable to parse page input")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		log.Printf("Process data for page: %v\n", page)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)
// 		search := q.Get(SEARCH)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("get failed: unable to verify token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		user, err := internal.GetTokenUser(token)
// 		if err != nil {
// 			log.Printf("get failed: unable to get user from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		internal.UpdateToken(token)

// 		role, err := internal.GetTokenRole(token)
// 		if err != nil {
// 			log.Printf("get failed: unable to get role from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 		}

// 		// data, err := s.DescribeUser(user)
// 		// if err != nil {
// 		// 	log.Printf("get failed: unable to describe user from token: %v\n", err)
// 		// 	writeRsp(w, err, &errorRsp{
// 		// 		Msg: fmt.Sprintf("%v", err),
// 		// 	}, http.StatusBadRequest)
// 		// }

// 		ret, err := s.DescribePatientByStaff(user, search, role)
// 		if err == nil {
// 			ret.Page = page
// 			if page == 0 {
// 				page = 1
// 			}
// 			arrayLen := len(ret.Data)

// 			ret.TotalRecord = arrayLen
// 			ret.TotalPage = arrayLen / ITEM_PER_PAGE
// 			mod := arrayLen % ITEM_PER_PAGE
// 			if mod != 0 {
// 				ret.TotalPage += 1
// 			}
// 			log.Printf("Quyen debug: TotalPage: %v", ret.TotalPage)

// 			var start int
// 			var end int

// 			start = (page - 1) * ITEM_PER_PAGE
// 			if page == ret.TotalPage {
// 				end = arrayLen
// 			} else {
// 				end = (page-1)*ITEM_PER_PAGE + ITEM_PER_PAGE
// 			}

// 			if end > arrayLen {
// 				end = arrayLen
// 			}

// 			log.Printf("Quyen debug: arraylen: %v", arrayLen)
// 			log.Printf("Quyen debug: start: %v", start)
// 			log.Printf("Quyen debug: end: %v", end)

// 			ret.Data = ret.Data[start:end]

// 			internal.UpdateToken(token)
// 			writeRsp(w, nil, ret, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get role: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPatientByAdminist(s account.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Println("Details:", r.Body)

// 		spage := mux.Vars(r)[PAGE]
// 		page, err := strconv.Atoi(spage)
// 		if err != nil {
// 			log.Println("Get user failed: unable to parse page input")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		log.Printf("Process data for page: %v\n", page)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)
// 		search := q.Get(SEARCH)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("get failed: unable to verify token")
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid token"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		role, err := internal.GetTokenRole(token)
// 		if err != nil {
// 			log.Printf("get failed: unable to get role from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		// if role == constant.AdministrativeRole {
// 		// 	writeRsp(w, err, &errorRsp{
// 		// 		Msg: fmt.Sprintf("Invalid role"),
// 		// 	}, http.StatusBadRequest)
// 		// 	return
// 		// }

// 		if (role < constant.DoctorRole) && (role > constant.AdministrativeRole) {
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("Invalid role"),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		user, err := internal.GetTokenUser(token)
// 		if err != nil {
// 			log.Printf("get failed: unable to get user from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		internal.UpdateToken(token)

// 		data, err := s.DescribeUser(user)
// 		if err != nil {
// 			log.Printf("get failed: unable to describe user from token: %v\n", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		ret, err := s.DescribePatientByAdminst(data.CreateBy, search)
// 		if err == nil {
// 			ret.Page = page
// 			if page == 0 {
// 				page = 1
// 			}
// 			arrayLen := len(ret.Data)

// 			ret.TotalRecord = arrayLen
// 			ret.TotalPage = arrayLen / ITEM_PER_PAGE
// 			mod := arrayLen % ITEM_PER_PAGE
// 			if mod != 0 {
// 				ret.TotalPage += 1
// 			}
// 			log.Printf("Quyen debug: TotalPage: %v", ret.TotalPage)

// 			var start int
// 			var end int

// 			start = (page - 1) * ITEM_PER_PAGE
// 			if page == ret.TotalPage {
// 				end = arrayLen
// 			} else {
// 				end = (page-1)*ITEM_PER_PAGE + ITEM_PER_PAGE
// 			}

// 			if end > arrayLen {
// 				end = arrayLen
// 			}

// 			log.Printf("Quyen debug: arraylen: %v", arrayLen)
// 			log.Printf("Quyen debug: start: %v", start)
// 			log.Printf("Quyen debug: end: %v", end)

// 			ret.Data = ret.Data[start:end]

// 			internal.UpdateToken(token)
// 			writeRsp(w, nil, ret, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get role: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

func newCreateUser(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %+v, method: %+v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		instance, err := parseReqDatas(r)
		if err != nil {
			log.Println("Create failed: unable to parse user input: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		log.Printf("Quyen debug body: %+v\n", instance)

		// if !internal.IsTokenExist(instance.Token) {
		// 	log.Println("Create failed: unable to get token")
		// 	writeRsp(w, err, &errorRsp{
		// 		Msg: fmt.Sprintf("Invalid token"),
		// 	}, http.StatusBadRequest)
		// 	return
		// }

		// role, err := internal.GetTokenRole(instance.Token)
		// if err != nil {
		// 	writeRsp(w, err, &errorRsp{
		// 		Msg: fmt.Sprintf("Invalid role"),
		// 	}, http.StatusBadRequest)
		// 	return
		// }

		// if (role == constant.SystemRole) || (role == constant.AdminRole) || (role == constant.ClinicRole) || (role == constant.AdministrativeRole) {
		// 	instance, err = s.Create(instance, true)
		// 	if err == nil {
		// 		writeRsp(w, nil, fmt.Sprintf("Create sucessfully"), http.StatusOK)
		// 	} else {
		// 		instance.Passwd = ""
		// 		log.Printf("Failed to authenticate with data: %v - %v", instance, err)
		// 		writeRsp(w, err, &errorRsp{
		// 			Msg: fmt.Sprintf("%v", err),
		// 		}, http.StatusNotFound)
		// 	}
		// } else {
		// 	writeRsp(w, err, &errorRsp{
		// 		Msg: fmt.Sprintf("Invalid role"),
		// 	}, http.StatusBadRequest)
		// }

		if (instance.Account.Role == constant.PatientRole) || (instance.Account.Role == constant.DoctorRole) || (instance.Account.Role == constant.ExpertRole) {
			instance, err = s.Create(instance, true)
			if err == nil {
				writeRsp(w, nil, fmt.Sprintf("Create sucessfully"), http.StatusOK)
			} else {
				instance.Account.Passwd = ""
				log.Printf("Failed to create with data: %+v - %+v", instance, err)
				writeRsp(w, err, &errorRsp{
					Msg: fmt.Sprintf("%v", err),
				}, http.StatusNotFound)
			}

		} else {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid role"),
			}, http.StatusBadRequest)
			return
		}

	}
}

func newUpdateUser(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %+v, method: %+v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		instance, err := parseReqDatas(r)
		if err != nil {
			log.Println("update failed: unable to parse user input: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		if len(instance.Account.User) == 0 {
			log.Println("user empty: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("User should not empty"),
			}, http.StatusBadRequest)
			return
		}

		log.Printf("Quyen debug body: %+v\n", instance)

		if !internal.IsTokenExist(instance.Account.Token) {
			log.Println("update failed: unable to get token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		role, err := internal.GetTokenRole(instance.Account.Token)
		if err != nil {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		// if role != constant.PatientRole {
		// 	writeRsp(w, err, &errorRsp{
		// 		Msg: fmt.Sprintf("Invalid role"),
		// 	}, http.StatusBadRequest)
		// 	return
		// }

		user, err := internal.GetTokenUser(instance.Account.Token)
		if err != nil {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		if role == constant.PatientRole {
			if user != instance.Account.User {
				writeRsp(w, err, &errorRsp{
					Msg: fmt.Sprintf("Invalid user"),
				}, http.StatusBadRequest)
				return
			}
		}

		instance, err = s.Create(instance, false)
		if err == nil {
			writeRsp(w, nil, fmt.Sprintf("update sucessfully"), http.StatusOK)
		} else {
			instance.Account.Passwd = ""
			log.Printf("Failed to update with data: %v - %v", instance, err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getUserByDoctor(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		spage := mux.Vars(r)[PAGE]
		page, err := strconv.Atoi(spage)
		if err != nil {
			log.Println("Get user failed: unable to parse page input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		log.Printf("Process data for page: %v\n", page)
		if page <= 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid page (page should > 0)"),
			}, http.StatusBadRequest)
			return
		}
		page = page - 1
		q := r.URL.Query()
		// creator := q.Get(CREATOR)
		token := q.Get(TOKEN)
		// srole := q.Get(ROLE)
		// role, _ := strconv.Atoi(srole)
		search := q.Get(SEARCH)
		doctor := q.Get(DOCTOR)

		log.Printf("Get user by %v\n", doctor)

		if len(doctor) == 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("doctor params should not empty"),
			}, http.StatusBadRequest)
			return
		}

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to verify token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		internal.UpdateToken(token)

		ret, err := s.GetUserByDoctor(doctor, search, page)
		if err == nil {
			internal.UpdateToken(token)
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get role: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getUserByExpert(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		spage := mux.Vars(r)[PAGE]
		page, err := strconv.Atoi(spage)
		if err != nil {
			log.Println("Get user failed: unable to parse page input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		log.Printf("Process data for page: %v\n", page)
		if page <= 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid page (page should > 0)"),
			}, http.StatusBadRequest)
			return
		}
		page = page - 1

		q := r.URL.Query()
		// creator := q.Get(CREATOR)
		token := q.Get(TOKEN)
		// srole := q.Get(ROLE)
		// role, _ := strconv.Atoi(srole)
		search := q.Get(SEARCH)
		expert := q.Get(EXPERT)

		log.Printf("Get user by %v\n", expert)

		if len(expert) == 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("expert params should not empty"),
			}, http.StatusBadRequest)
			return
		}

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to verify token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		internal.UpdateToken(token)

		ret, err := s.GetUserByExpert(expert, search, page)
		if err == nil {
			internal.UpdateToken(token)
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getUserByAdmin(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		spage := mux.Vars(r)[PAGE]
		page, err := strconv.Atoi(spage)
		if err != nil {
			log.Println("Get user failed: unable to parse page input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		log.Printf("Process data for page: %v\n", page)
		if page <= 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid page (page should > 0)"),
			}, http.StatusBadRequest)
			return
		}
		page = page - 1

		q := r.URL.Query()
		// creator := q.Get(CREATOR)
		token := q.Get(TOKEN)
		// srole := q.Get(ROLE)
		// role, _ := strconv.Atoi(srole)
		search := q.Get(SEARCH)
		expert := q.Get(EXPERT)

		log.Printf("Get user by %v\n", expert)

		if len(expert) == 0 {
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("expert params should not empty"),
			}, http.StatusBadRequest)
			return
		}

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to verify token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		internal.UpdateToken(token)

		ret, err := s.GetUserByAdmin(expert, search, page)
		if err == nil {
			internal.UpdateToken(token)
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getUser(s account.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Println("Details:", r.Body)

		user := mux.Vars(r)[USER]
		log.Printf("Process data for user: %v\n", user)

		q := r.URL.Query()
		token := q.Get("token")

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to verify token")
			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}
		internal.UpdateToken(token)

		// tUser, err := internal.GetTokenUser(token)
		// if err != nil {
		// 	log.Println("Failed to get token user")
		// 	http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
		// 	return
		// }

		// if user != tUser {
		// 	log.Println("User and token not match")
		// 	http.Error(w, fmt.Sprintf("Invalid user"), http.StatusBadRequest)
		// 	return
		// }

		ret, err := s.DescribeUser(user)
		if err == nil {
			log.Printf("Quyen debug data return: %v\n", ret)
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get user: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}
