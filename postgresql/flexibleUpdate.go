package main

import (
	"encoding/json"
	"log"
	"strconv"
)

func main() {
	// unmarshal data
	req := []byte(`{"name":"TEST1","phone_number":"+62857887878","id":1}`)
	var model map[string]interface{}
	_ = json.Unmarshal(req, &model)

	SQL, values := FlexibleUpdate1("customer", model, "id")

	log.Println("Print Query: ", SQL)
	log.Println("Print Query: ", values)
	// db.Conn.Exec(SQL, values...)
}

//FlexibleUpdate1 function to update
func FlexibleUpdate1(tableName string, payload map[string]interface{}, cond string) (SQL string, values []interface{}) {
	SQL = "UPDATE " + tableName + " SET"

	j := 0
	var condData interface{}
	for key, v := range payload {
		if key != cond {
			j++
			comma := ","
			if j == 1 {
				comma = ""
			}
			SQL = SQL + comma + " " + key + "=$" + strconv.Itoa(j)
			values = append(values, v)
		} else {
			condData = v
		}
	}

	//add where condition
	values = append(values, condData)
	SQL = SQL + " where id=$" + strconv.Itoa(j+1)

	return
}

