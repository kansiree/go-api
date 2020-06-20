package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Masterdata []struct {
	CreatedDate string `json:"created_date"`
	FullName    string `json:"full_name"`
	ID          string `json:"id"`
}
type Masters []Masterdata

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the  !")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getMasterSystem", getMasterSystem)
	http.HandleFunc("/getMasterAircraft", getMasterAircraft)
	http.HandleFunc("/getMasterTechnicalOrder", getMasterTechnicalOrder)
	http.ListenAndServe(getPort(), nil)
}

func getMasterSystem(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_system")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonMasterdataFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func getMasterAircraft(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_aircraft")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonMasterdataFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func getMasterTechnicalOrder(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_technical_order")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonMasterdataFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}

func getJSON(sqlString string) (string, error) {
	db, err := sql.Open("mysql", "sz0debklevf8wjhf:gu2af8swu50tjc3k@tcp(u3r5w4ayhxzdrw87.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/t016ffukzsi0y5ie")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func convertStringToJsonMasterdataFormat(message string, format Masterdata) ([]byte, error) {
	errpare := json.Unmarshal([]byte(message), &format)
	if errpare != nil {
		return nil, errpare
	} else {
		json, err := json.Marshal(format)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
}
