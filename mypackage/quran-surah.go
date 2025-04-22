package mypackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GenerateQuranSurah() {
	// read word-data.json
	wordData, err := os.ReadFile("data/word-data.json")
	if err != nil {
		log.Fatal(err)
	}
	type Word struct {
		Surah int `json:"surah"`
	}
	var resultWord map[string]Word
	json.Unmarshal(wordData, &resultWord)

	// read surah-data.json
	surahData, err2 := os.ReadFile("data/surah-data.json")
	if err2 != nil {
		log.Fatal(err2)
	}
	type Surah struct {
		Name            string `json:"name"`
		TotalAyah       int    `json:"nAyah"`
		RevelationOrder int    `json:"revelationOrder"`
		Type            string `json:"type"`
		AyahStart       int    `json:"start"`
		AyahEnd         int    `json:"end"`
	}
	var resultSurah map[string]Surah
	json.Unmarshal(surahData, &resultSurah)

	// read surah-data-indo.json
	surahDataIndo, err3 := os.ReadFile("data/surah-data-indo.json")
	if err3 != nil {
		log.Fatal(err3)
	}
	type SurahIndoItem struct {
		Id            int    `json:"id"`
		Name          string `json:"surat_name"`
		NameTranslate string `json:"surat_terjemahan"`
	}
	type SurahIndo struct {
		Msg  string             `json:"msg"`
		Data [114]SurahIndoItem `json:"data"`
	}
	var resultSurat SurahIndo
	json.Unmarshal(surahDataIndo, &resultSurat)

	// read surah-name-en.json
	surahNameEnData, err3 := os.ReadFile("data/surah-name-en.json")
	if err3 != nil {
		log.Fatal(err3)
	}
	type SurahNameEn struct {
		Name        string `json:"name"`
		Translation string `json:"translation"`
	}
	var resultSurahNameEn map[string]SurahNameEn
	json.Unmarshal(surahNameEnData, &resultSurahNameEn)

	// create data for sql
	type SurahDBItem struct {
		SurahId         int
		SurahName       string
		ArabicName      string
		IdnName         string
		EnName          string
		Type            string
		TotalAyah       int
		WordCount       int
		RevelationOrder int
		AyahStart       int
		AyahEnd         int
	}

	var surahDB []SurahDBItem = make([]SurahDBItem, 0)
	var surahIndex = 0
	var surahIndexStr = "0"

	for i := 1; i <= len(resultWord); i++ {
		if value := resultWord[strconv.Itoa(i)]; value.Surah != surahIndex {
			surahIndex = value.Surah
			surahIndexStr = strconv.Itoa(value.Surah)
			surahDB = append(surahDB,
				SurahDBItem{
					SurahId:         value.Surah,
					SurahName:       resultSurat.Data[value.Surah-1].Name,
					ArabicName:      resultSurah[surahIndexStr].Name,
					IdnName:         resultSurat.Data[value.Surah-1].NameTranslate,
					EnName:          resultSurahNameEn[surahIndexStr].Translation,
					Type:            resultSurah[surahIndexStr].Type,
					TotalAyah:       resultSurah[surahIndexStr].TotalAyah,
					WordCount:       0,
					RevelationOrder: resultSurah[surahIndexStr].RevelationOrder,
					AyahStart:       resultSurah[surahIndexStr].AyahStart,
					AyahEnd:         resultSurah[surahIndexStr].AyahEnd,
				})
		}
		surahDB[surahIndex-1].WordCount += 1
	}

	var query = "CREATE TABLE quran_surah( id INTEGER PRIMARY KEY, surahName VARCHAR(64), arabicName TEXT, idnName VARCHAR(128), enName VARCHAR(128),surahType VARCHAR(32), totalAyah INTEGER, wordCount INTEGER, revelationOrder INTEGER, ayahStart INTEGER, ayahEnd INTEGER, surahInfo TEXT);\n"

	for _, item := range surahDB {
		newQuery := fmt.Sprintf(`INSERT INTO quran_surah (id, surahName, arabicName, idnName, enName, surahType, totalAyah, wordCount, revelationOrder, ayahStart, ayahEnd, surahInfo ) VALUES (%s,"%s","%s","%s","%s","%s",%s,%s,%s,%s,%s,"");`, strconv.Itoa(item.SurahId), item.SurahName, item.ArabicName, item.IdnName, item.EnName, item.Type, strconv.Itoa(item.TotalAyah), strconv.Itoa(item.WordCount), strconv.Itoa(item.RevelationOrder), strconv.Itoa(item.AyahStart), strconv.Itoa(item.AyahEnd))
		query += newQuery + "\n"
	}

	err8 := os.WriteFile("sql/quran-surah.sql", []byte(query), 0777)
	if err8 != nil {
		log.Fatal(err8)
	}
	fmt.Println("quran-surah.sql written successfully.")
}
