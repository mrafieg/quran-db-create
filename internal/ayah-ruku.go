package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetAyahRuku() {
	// read ayah-ruku.json
	rukuData, err := os.ReadFile("data/ayah-ruku.json")
	if err != nil {
		log.Fatal(err)
	}
	type Ruku struct {
		Start int `json:"start"`
		End   int `json:"end"`
	}
	var resultRuku map[int]Ruku
	json.Unmarshal(rukuData, &resultRuku)

	// build query
	var query = ""
	for _, ruku := range resultRuku {
		for id := ruku.Start; id <= ruku.End; id++ {
			query += fmt.Sprintf(`UPDATE quran_ayahs SET ruku = "%v:%v" WHERE id = %v`, ruku.Start, ruku.End, id) + ";\n"
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/10_ayah-ruku.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-ruku.sql written successfully.")

}
