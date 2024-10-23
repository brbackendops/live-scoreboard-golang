package match

import (
	"encoding/json"
	"net/http"
	"score/database/types"
	"strconv"

	"github.com/valyala/fasthttp"
)

// get single match
func (mc *MatchController) GetMatchController(req *fasthttp.RequestCtx) {

	id := req.UserValue("id")

	did, err := strconv.Atoi(id.(string))
	if err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := mc.matchService.GetMatchService(did)
	if err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	req.SetStatusCode(http.StatusOK)
	req.Success("application/json", res)
}

// get many matches
func (mc *MatchController) GetMatchesController(req *fasthttp.RequestCtx) {

	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	req.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	// req.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	data, err := mc.matchService.GetMatchesService()
	if err != nil {
		req.Error(err.Error(), http.StatusBadRequest)
		return
	}

	res := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	jdata, err := json.Marshal(&res)
	if err != nil {
		req.Error(err.Error(), http.StatusBadRequest)
		return
	}

	req.SetStatusCode(http.StatusOK)
	req.Success("application/json", jdata)
}

// create match
func (mc *MatchController) CreateMatchController(req *fasthttp.RequestCtx) {
	var payload types.MatchCreate

	body := req.PostBody()

	if err := json.Unmarshal(body, &payload); err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := mc.matchService.CreateMatchService(&payload)
	if err != nil {
		req.Error(err.Error(), http.StatusBadRequest)
		return
	}

	res := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	jdata, err := json.Marshal(&res)
	if err != nil {
		req.Error(err.Error(), http.StatusBadRequest)
		return
	}

	req.SetStatusCode(http.StatusCreated)
	req.Success("application/json", jdata)

}

// update match
func (mc *MatchController) UpdateMatchController(req *fasthttp.RequestCtx) {
	var payload types.MatchUpdate

	body := req.PostBody()

	if err := json.Unmarshal(body, &payload); err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	// publish latest score updates to redis channel match_updates
	if err := mc.rClient.Publish(req, "match_updates", body).Err(); err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := mc.matchService.UpdateMatchService(&payload)
	if err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	result, err := json.Marshal(res)
	if err != nil {
		req.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	req.SetStatusCode(http.StatusOK)
	req.Success("application/json", result)

}
