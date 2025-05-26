package mypackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetAyahSajda() {
	// read ayah-sajda.json
	sajdaData, err := os.ReadFile("data/ayah-sajda.json")
	if err != nil {
		log.Fatal(err)
	}
	type Sajda struct {
		Type string `json:"type"`
		Ayah int    `json:"ayah"`
	}
	var resultSajda map[string]Sajda
	json.Unmarshal(sajdaData, &resultSajda)

	// build query
	var query = ""
	for _, sajda := range resultSajda {
		query += fmt.Sprintf(`UPDATE quran_ayahs SET sajda = "%s" WHERE id = %s`, sajda.Type, strconv.Itoa(sajda.Ayah)) + ";\n"
	}

	// write sql file
	err3 := os.WriteFile("sql/9_ayah-sajda.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-sajda.sql written successfully.")

}
