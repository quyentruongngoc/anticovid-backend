package status

import (
	"anti-corona-backend/internal"
	"anti-corona-backend/package/api-process/status"
	"anti-corona-backend/package/constant"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	ROLE    = "role"
	RECNUM  = "recnum"
	TOKEN   = "token"
	PATIENT = "patient"
	ID      = "id"
)

func Install(r *mux.Router, s status.Rest) {
	r.HandleFunc("/staff/status/{role}", addStatusRecord(s)).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/staff/status/{role}/{recnum}", getStatusRecord(s)).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/staff/list/status", getStatusList(s)).Methods(http.MethodGet, http.MethodHead)

	r.HandleFunc("/patient/status", addPatientStatus(s)).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/patient/status/{id}", updatePatientStatus(s)).Methods(http.MethodPut, http.MethodHead)
	r.HandleFunc("/patient/status/{recnum}", getPatientStatus(s)).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/patient/list/status", getPatientStatusList(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/patient/info/{patient}", getPartientInfo(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/patient/mgmt", updatePatientManagerment(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/patient/mgmt/{patient}", getPatientManagerment(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/patien/medical", updatePatientMedical(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/patien/medical/{patient}", getPatientMedical(s)).Methods(http.MethodGet, http.MethodHead)
}

type errorRsp struct {
	Msg string
}

func parseStaffReqDatas(r *http.Request) (status.StaffStatus, error) {
	decoder := json.NewDecoder(r.Body)
	var instance status.StaffStatus
	err := decoder.Decode(&instance)
	return instance, err
}

func parsePatientReqDatas(r *http.Request) (status.PatientStatus, error) {
	decoder := json.NewDecoder(r.Body)
	var instance status.PatientStatus
	err := decoder.Decode(&instance)
	return instance, err
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

func addStatusRecord(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		role := mux.Vars(r)[ROLE]
		log.Printf("Add status record for : %v\n", role)

		instance, err := parseStaffReqDatas(r)
		if err != nil {
			log.Println("Add failed: unable to parse user input: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		log.Printf("Quyen debug body: %v\n", instance)

		if !internal.IsTokenExist(instance.Token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		switch role {
		case "doctor":
			// instance.Role = constant.DoctorRole
			if instance.Type == constant.CommandExpert {
				instance.Role = constant.ExpertRole
			} else if instance.Type == constant.CommandPatient {
				instance.Role = constant.PatientRole
			}
		case "expert":
			instance.Role = constant.ExpertRole
			instance.Type = constant.ReportExpert
		case "patient":
			instance.Role = constant.PatientRole
			instance.Type = constant.ReportPatient
		default:
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid role variable"),
			}, http.StatusBadRequest)
			return
		}

		instance, err = s.AddStatusRecord(instance)
		if err == nil {
			writeRsp(w, nil, fmt.Sprintf("Add staff record sucessfully"), http.StatusOK)
		} else {
			log.Printf("Failed to add record with data: %v - %v", instance, err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getStatusRecord(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		role := mux.Vars(r)[ROLE]
		log.Printf("Get status record for: %v\n", role)

		snum := mux.Vars(r)[RECNUM]
		log.Printf("Get record number: %v\n", snum)
		num, err := strconv.Atoi(snum)
		if err != nil {
			log.Println("Get user failed: unable to parse page input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Unable to parse page input: %v", err),
			}, http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		token := q.Get(TOKEN)

		q = r.URL.Query()
		patient := q.Get(PATIENT)
		log.Printf("Patient: %v\n", patient)

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		instance := status.StaffStatus{
			RecordNumber: num,
			PatientUUID:  patient,
		}

		switch role {
		case "doctor":
			instance.Role = constant.DoctorRole
		case "expert":
			instance.Role = constant.ExpertRole
		case "patient":
			instance.Role = constant.PatientRole
		default:
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid role variable"),
			}, http.StatusBadRequest)
			return
		}

		ret, err := s.GetStatusRecord(instance)
		if err == nil {
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient info: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func addPatientStatus(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		instance, err := parsePatientReqDatas(r)
		if err != nil {
			log.Println("Add failed: unable to parse user input: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		log.Printf("Quyen debug body: %v\n", instance)

		if !internal.IsTokenExist(instance.Token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		instance, err = s.AddPatientStatus(instance)
		if err == nil {
			writeRsp(w, nil, fmt.Sprintf("Add patient record sucessfully"), http.StatusOK)
		} else {
			log.Printf("Failed to add record with data: %v - %v", instance, err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func updatePatientStatus(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		instance, err := parsePatientReqDatas(r)
		if err != nil {
			log.Println("Add failed: unable to parse user input: ", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}

		log.Printf("Quyen debug body: %v\n", instance)

		sid := mux.Vars(r)[ID]
		id, err := strconv.Atoi(sid)
		if err != nil {
			log.Println("Update patient status failed: unable to parse id input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusBadRequest)
			return
		}
		log.Printf("Process data for id: %v\n", id)
		instance.ID = uint(id)

		if !internal.IsTokenExist(instance.Token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		instance, err = s.UpdatePatientStatus(instance)
		if err == nil {
			writeRsp(w, nil, fmt.Sprintf("Add patient record sucessfully"), http.StatusOK)
		} else {
			log.Printf("Failed to add record with data: %v - %v", instance, err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getPatientStatus(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		snum := mux.Vars(r)[RECNUM]
		log.Printf("Get record number: %v\n", snum)
		num, err := strconv.Atoi(snum)
		if err != nil {
			log.Println("Get user failed: unable to parse page input")
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("Unable to parse page input: %v", err),
			}, http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		token := q.Get(TOKEN)

		q = r.URL.Query()
		patient := q.Get(PATIENT)
		log.Printf("Patient: %v\n", patient)

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		instance := status.PatientStatus{
			RecordNumber: num,
			UserUUID:     patient,
		}

		ret, err := s.GetPatientStatus(instance)
		if err == nil {
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient info: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getStatusList(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		q := r.URL.Query()
		token := q.Get(TOKEN)
		patient := q.Get(PATIENT)
		log.Printf("Patient: %v\n", patient)

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		ret, err := s.GetStaffRecordList(patient)
		if err == nil {
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient info: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}

func getPatientStatusList(s status.Rest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
		log.Printf("Details: %v\n", r.Body)

		q := r.URL.Query()
		token := q.Get(TOKEN)

		q = r.URL.Query()
		patient := q.Get(PATIENT)
		log.Printf("Patient: %v\n", patient)

		if !internal.IsTokenExist(token) {
			log.Println("Create failed: unable to get token")
			writeRsp(w, fmt.Errorf("Invalid token"), &errorRsp{
				Msg: fmt.Sprintf("Invalid token"),
			}, http.StatusBadRequest)
			return
		}

		ret, err := s.GetPatientRecordList(patient)
		if err == nil {
			writeRsp(w, nil, ret, http.StatusOK)
		} else {
			log.Printf("Failed to get patient info: %v", err)
			writeRsp(w, err, &errorRsp{
				Msg: fmt.Sprintf("%v", err),
			}, http.StatusNotFound)
		}
	}
}
