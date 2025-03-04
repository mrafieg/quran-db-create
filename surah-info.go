package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func generateSurahInfo() {
	// create scanner
	infoHtml, err := os.Open("data/surah-info.md")
	if err != nil {
		log.Fatal(err)
	}

	defer infoHtml.Close()
	scanner := bufio.NewScanner(infoHtml)

	// read lines
	var surahInfos [115]string
	var currentSurah int64
	var regexArabic, _ = regexp.Compile("[\u0600-\u06FF]")
	for scanner.Scan() {
		line := string(mdToHTML(scanner.Bytes()))
		// get surah id
		if strings.Index(line, "<h1>") == 0 {
			id := strings.Replace(line, "<h1>", "", 1)
			id = strings.Replace(id, "</h1>", "", 1)
			id = strings.TrimSpace(id)
			currentSurah, _ = strconv.ParseInt(id, 10, 64)
			continue
		}
		// check if line in arabic
		if match := regexArabic.FindAllString(line, -1); len(match) > 0 {
			line = strings.Replace(line, "<p>", `<p dir="rtl">`, 1)
		}
		// append line to approtiate surah description based on id
		surahInfos[currentSurah] += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// build query
	var query = ""
	for i, info := range surahInfos {
		if i > 0 {
			query += fmt.Sprintf(`UPDATE quran_surah SET surahInfo = "%s" WHERE id = %s;`, info, (strconv.Itoa(i))) + "\n"
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/surah-info.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("surah-info.sql written successfully.")
}
