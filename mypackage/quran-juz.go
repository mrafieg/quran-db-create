package mypackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetQuranJuz() {
	// read juz-data.json
	juzData, err := os.ReadFile("data/juz-data.json")
	if err != nil {
		log.Fatal(err)
	}
	type Juz struct {
		Start int `json:"start"`
		End   int `json:"end"`
	}
	var resultJuz map[string]Juz
	json.Unmarshal(juzData, &resultJuz)

	// build query
	query := "CREATE TABLE quran_juzs( id INTEGER PRIMARY KEY AUTO_INCREMENT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP(), updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP() , juz INTEGER, surah_id INTEGER, ayah_id INTEGER, FOREIGN kEY (surah_id) REFERENCES quran_surahs(id), FOREIGN kEY (ayah_id) REFERENCES quran_ayahs(id));\n"
	for juzId, juz := range resultJuz {
		query += fmt.Sprintf(`INSERT INTO quran_juzs (juz, ayah_id) SELECT %v as juz, id as ayah_id FROM quran_ayahs WHERE id BETWEEN %v AND %v`, juzId, juz.Start, juz.End) + ";\n"
		query += fmt.Sprintf(`INSERT INTO quran_juzs (juz, surah_id) SELECT %v as juz, surah_id FROM quran_ayahs WHERE id BETWEEN %v AND %v GROUP BY surah_id`, juzId, juz.Start, juz.End) + ";\n"
	}

	// write sql file
	err3 := os.WriteFile("sql/11_quran-juz.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("quran-juz.sql written successfully.")
}
