package patient

import (
	"anti-corona-backend/package/api-process/patient"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PATIENT = "patient"
	TOKEN   = "token"
)

func Install(r *mux.Router, s patient.Rest) {
	// r.HandleFunc("/patient/info", updatePatientInfo(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/patient/info/{patient}", getPartientInfo(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/patient/mgmt", updatePatientManagerment(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/patient/mgmt/{patient}", getPatientManagerment(s)).Methods(http.MethodGet, http.MethodHead)

	// r.HandleFunc("/patient/medical", updatePatientMedical(s)).Methods(http.MethodPost, http.MethodHead)
	// r.HandleFunc("/patient/medical/{patient}", getPatientMedical(s)).Methods(http.MethodGet, http.MethodHead)
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

func parseReqInfoDatas(r *http.Request) (patient.PatientInfo, error) {
	decoder := json.NewDecoder(r.Body)
	var instance patient.PatientInfo
	err := decoder.Decode(&instance)
	return instance, err
}

func parseReqMgmtDatas(r *http.Request) (patient.PatientMgmt, error) {
	decoder := json.NewDecoder(r.Body)
	var instance patient.PatientMgmt
	err := decoder.Decode(&instance)
	return instance, err
}

func parseReqMedicalDatas(r *http.Request) (patient.MedicalData, error) {
	decoder := json.NewDecoder(r.Body)
	var instance patient.MedicalData
	err := decoder.Decode(&instance)
	return instance, err
}

// func updatePatientInfo(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		instance, err := parseReqInfoDatas(r)
// 		if err != nil {
// 			log.Println("Create failed: unable to parse user input: ", err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		log.Printf("Quyen debug body: %v\n", instance)

// 		if !internal.IsTokenExist(instance.Token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err = s.UpdatePatientInfo(instance)
// 		if err == nil {
// 			writeRsp(w, nil, fmt.Sprintf("Create/Update sucessfully"), http.StatusOK)
// 		} else {
// 			log.Printf("Failed to create/update patient with data: %v - %v", instance, err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPartientInfo(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		patient := mux.Vars(r)[PATIENT]
// 		log.Printf("Get patient info for: %v\n", patient)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err := s.GetPatientInfo(patient)
// 		if err == nil {
// 			writeRsp(w, nil, instance, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get patient info: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func updatePatientManagerment(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		instance, err := parseReqMgmtDatas(r)
// 		if err != nil {
// 			log.Println("Create failed: unable to parse user input: ", err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		log.Printf("Quyen debug body: %+v\n", instance)

// 		if !internal.IsTokenExist(instance.Token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err = s.UpdatePatientMgmt(instance)
// 		if err == nil {
// 			writeRsp(w, nil, fmt.Sprintf("Create/Update sucessfully"), http.StatusOK)
// 		} else {
// 			log.Printf("Failed to create/update patient with data: %v - %v", instance, err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPatientManagerment(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		patient := mux.Vars(r)[PATIENT]
// 		log.Printf("Get patient info for: %v\n", patient)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err := s.GetPatientMgmt(patient)
// 		if err == nil {
// 			writeRsp(w, nil, instance, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get patient managerment: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func updatePatientMedical(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		instance, err := parseReqMedicalDatas(r)
// 		if err != nil {
// 			log.Println("Update failed: unable to parse user input: ", err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		log.Printf("Quyen debug body: %+v\n", instance)

// 		if !internal.IsTokenExist(instance.Token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err = s.UpdatePatientMedical(instance)
// 		if err == nil {
// 			writeRsp(w, nil, fmt.Sprintf("Create/Update sucessfully"), http.StatusOK)
// 		} else {
// 			log.Printf("Failed to create/update patient medical with data: %v - %v", instance, err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }

// func getPatientMedical(s patient.Rest) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Received request: %v, method: %v\n", r.URL.String(), r.Method)
// 		log.Printf("Details: %v\n", r.Body)

// 		patient := mux.Vars(r)[PATIENT]
// 		log.Printf("Get patient medical for: %v\n", patient)

// 		q := r.URL.Query()
// 		token := q.Get(TOKEN)

// 		if !internal.IsTokenExist(token) {
// 			log.Println("Create failed: unable to get token")
// 			http.Error(w, fmt.Sprintf("Invalid token"), http.StatusBadRequest)
// 			return
// 		}

// 		instance, err := s.GetPatientMedical(patient)
// 		if err == nil {
// 			writeRsp(w, nil, instance, http.StatusOK)
// 		} else {
// 			log.Printf("Failed to get patient medical: %v", err)
// 			writeRsp(w, err, &errorRsp{
// 				Msg: fmt.Sprintf("%v", err),
// 			}, http.StatusNotFound)
// 		}
// 	}
// }
