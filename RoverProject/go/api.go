package mars

import (
	// "fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"

)

func MoveRover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	var xyCoordOfPlateau [2]int
    xyCoordOfPlateau[0], _ = strconv.Atoi(vars["xOfPlateau"])
	xyCoordOfPlateau[1], _ = strconv.Atoi(vars["yOfPlateau"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var rover Rover
	_ = json.Unmarshal([]byte(reqBody), &rover)
	roverDest, err := FindDestination(rover, xyCoordOfPlateau)
	if err!= nil{
		var apiResponse ApiResponse
		apiResponse.Code = 502
		apiResponse.Status = "Failure"
		apiResponse.Message = "The rover moves out of the plateau"
		json.NewEncoder(w).Encode(apiResponse)
	}else{
		json.NewEncoder(w).Encode(roverDest)
	}
	w.WriteHeader(http.StatusOK)
}
