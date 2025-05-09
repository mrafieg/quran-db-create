package mypackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetAyahWordCount() {
	// read word-data.json
	wordData, err := os.ReadFile("data/word-data.json")
	if err != nil {
		log.Fatal(err)
	}
	type Word struct {
		Surah int `json:"surah"`
		Ayah  int `json:"ayah"`
	}
	var resultWord map[string]Word
	json.Unmarshal(wordData, &resultWord)

	type Ayah struct {
		VerseId   int
		WordCount int
	}

	type Surah struct {
		Id    int
		Ayahs []Ayah
	}
	var surahs [115]Surah
	var currentSurah = 0
	var currentAyah = 0

	for i := 1; i <= len(resultWord); i++ {
		if value := resultWord[strconv.Itoa(i)]; value.Surah != currentSurah {
			currentSurah = value.Surah
			currentAyah = value.Ayah
			surahs[currentSurah] = Surah{
				Id:    value.Surah,
				Ayahs: make([]Ayah, 0),
			}
			surahs[currentSurah].Ayahs = append(surahs[currentSurah].Ayahs, Ayah{VerseId: value.Ayah, WordCount: 0})
		} else if value.Surah == currentSurah && value.Ayah != currentAyah {
			currentAyah = value.Ayah
			surahs[currentSurah].Ayahs = append(surahs[currentSurah].Ayahs, Ayah{VerseId: value.Ayah, WordCount: 0})
		}
		surahs[currentSurah].Ayahs[currentAyah-1].WordCount += 1
	}

	// build query
	var query = ""
	for i, surah := range surahs {
		if i > 0 {
			for _, ayah := range surah.Ayahs {
				query += fmt.Sprintf(`UPDATE quran_ayah SET wordCount = %s WHERE surahId = %s AND verseId = %s`, strconv.Itoa(ayah.WordCount), strconv.Itoa(surah.Id), strconv.Itoa(ayah.VerseId)) + ";\n"
			}
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/8_ayah-word-count.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-word-count.sql written successfully.")

}
