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
	mypackage.GenerateQuranSurah()
	mypackage.GenerateQuranAyah()
	mypackage.GenerateSurahInfo()
	mypackage.GenerateAyahInfo()
	mypackage.GetAyahWordCount()
}
