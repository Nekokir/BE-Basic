package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"github.com/yanzay/log"
	"encoding/json"
)

func responseJson(w http.ResponseWriter, data map[string]interface{}) {
	res_json, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred encoding response.", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res_json)
	return
}

func SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params)  {
	req.ParseForm()
	if (len(req.Form["name"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid name.",
		}
		responseJson(w, res)
		return
	}
	name := req.Form["name"][0]
	user, err := StoreUser(db, name)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred creating user. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"data":   user,
	}
	responseJson(w, res)
}

func EditName(w http.ResponseWriter, req *http.Request, ps httprouter.Params)  {
	req.ParseForm()
	if (len(req.Form["token"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid token.",
		}
		responseJson(w, res)
		return
	}
	token := req.Form["token"][0]
	if (len(req.Form["name"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid name.",
		}
		responseJson(w, res)
		return
	}
	name := req.Form["name"][0]
	user, err := UpdateUser(db, token, name)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred updating user. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"data":   user,
	}
	responseJson(w, res)
}

func SubmitQuiz(w http.ResponseWriter, req *http.Request, ps httprouter.Params)  {
	req.ParseForm()
	if (len(req.Form["token"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid token.",
		}
		responseJson(w, res)
		return
	}
	token := req.Form["token"][0]
	user, err := SubmitUser(db, token)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred submitting. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"msg": "SUCCESS",
		"data":   user,
	}
	responseJson(w, res)
}


func ShowNames(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	names, err := FindSubmittedUsers(db)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred querying names. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"data":   names,
	}
	responseJson(w, res)
}

func ShowDetails(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	if (len(req.Form["pwd"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid pwd.",
		}
		responseJson(w, res)
		return
	}
	key := req.Form["pwd"][0]
	if key != GlobCfg.PWD {
		res := map[string]interface{}{
			"result": false,
			"msg":   "Invalid pwd.",
		}
		responseJson(w, res)
	}
	names, err := FindDetailedUsers(db)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred querying names. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"data":   names,
	}
	responseJson(w, res)
}

func AnswerQuestion(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	if (len(req.Form["token"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid token.",
		}
		responseJson(w, res)
		return
	}
	token := req.Form["token"][0]
	user, err := FindUser(db, token)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	qid64, err := strconv.ParseUint(ps.ByName("qid"), 10, 32)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred parsing question id.", http.StatusInternalServerError)
		return
	}
	qid := uint(qid64)
	if (len(req.Form["answer"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid answer.",
		}
		responseJson(w, res)
		return
	}
	ans := req.Form["answer"][0]
	answer, err := StoreAnswer(db, user.ID, qid, ans)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred saving answer. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"msg":    "SUCCESS",
		"data":   answer,
	}
	responseJson(w, res)
}

func EditAnswer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	if (len(req.Form["token"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid token.",
		}
		responseJson(w, res)
		return
	}
	token := req.Form["token"][0]
	user, err := FindUser(db, token)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	qid64, err := strconv.ParseUint(ps.ByName("qid"), 10, 32)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred parsing question id.", http.StatusInternalServerError)
		return
	}
	qid := uint(qid64)
	if (len(req.Form["answer"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid answer.",
		}
		responseJson(w, res)
		return
	}
	ans := req.Form["answer"][0]
	answer, err := UpdateAnswer(db, user.ID, qid, ans)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred updating answer. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"msg":    "SUCCESS",
		"data":   answer,
	}
	responseJson(w, res)
}

func ShowAnswer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	if (len(req.Form["token"]) != 1) {
		res := map[string]interface{}{
			"result": false,
			"msg":    "Invalid token.",
		}
		responseJson(w, res)
		return
	}
	token := req.Form["token"][0]
	user, err := FindUser(db, token)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	qid64, err := strconv.ParseUint(ps.ByName("qid"), 10, 32)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred parsing question id.", http.StatusInternalServerError)
		return
	}
	qid := uint(qid64)
	answer, err := FindAnswer(db, user.ID, qid)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred querying answer. " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"result": true,
		"msg":    "SUCCESS",
		"data":   answer,
	}
	responseJson(w, res)
}