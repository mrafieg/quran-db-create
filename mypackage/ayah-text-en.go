package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GenerateAyahEnText() {
	file, err := os.Open("data/quran-text-en.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ayahTexts [6237]string
	var currentAyah int64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			currentAyah, _ = strconv.ParseInt(string(line[2:]), 10, 64)
		} else {
			ayahTexts[currentAyah] += line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// build query
	var query = ""
	for i, text := range ayahTexts {
		if i > 0 {
			query += strings.TrimSpace(fmt.Sprintf(`UPDATE quran_ayah SET enText = "%s" WHERE id = %s;`, Addslashes(text), (strconv.Itoa(i-1)))) + "\n"
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/ayah-text-en.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-text-en.sql written successfully.")
}
