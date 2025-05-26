package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GenerateAyahEnText() {
	// read quran-text-en.md
	file, err := os.Open("data/quran-text-en.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read lines
	var ayahTexts [6237]string
	var currentAyah int64
	for scanner.Scan() {
		line := scanner.Text()
		// get ayah id
		if strings.HasPrefix(line, "#") {
			currentAyah, _ = strconv.ParseInt(line[2:], 10, 64)
		} else {
			// append ayah text
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
			query += strings.TrimSpace(fmt.Sprintf(`UPDATE quran_ayahs SET en_text = "%s" WHERE id = %v;`, Addslashes(text), i)) + "\n"
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/3_ayah-text-en.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-text-en.sql written successfully.")
}
