package main

import (
	"jalan-surah-db-create/internal"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("sql", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	internal.GenerateQuranSurah()
	internal.GenerateQuranAyah()
	internal.GenerateAyahEnText()
	internal.GenerateSurahInfo()
	internal.GenerateSurahInfoEn()
	internal.GenerateAyahInfo()
	internal.GenerateAyahInfoEn()
	internal.GetAyahWordCount()
	internal.GetAyahSajda()
	internal.GetAyahRuku()
	internal.GetQuranJuz()
}
