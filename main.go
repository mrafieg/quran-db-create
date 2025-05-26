package main

import (
	"jalan-surah-db-create/mypackage"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("sql", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	// mypackage.GenerateQuranSurah()
	mypackage.GenerateQuranAyah()
	mypackage.GenerateAyahEnText()
	mypackage.GenerateSurahInfo()
	mypackage.GenerateSurahInfoEn()
	mypackage.GenerateAyahInfo()
	mypackage.GenerateAyahInfoEn()
	mypackage.GetAyahWordCount()
	mypackage.GetAyahSajda()
	mypackage.GetAyahRuku()
	mypackage.GetQuranJuz()
}
