package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GenerateAyahInfo() {
	// create scanner
	infoHtml, err := os.Open("data/ayah-info.md")
	if err != nil {
		log.Fatal(err)
	}

	defer infoHtml.Close()
	scanner := bufio.NewScanner(infoHtml)

	// read lines
	var ayahInfos [6237]string
	var currentAyah int64
	var regexArabic, _ = regexp.Compile("[\u0600-\u06FF]")
	for scanner.Scan() {
		line := string(MdToHTML(scanner.Bytes()))
		// get ayah id
		if strings.Index(line, "<h1>") == 0 {
			id := strings.Replace(line, "<h1>", "", 1)
			id = strings.Replace(id, "</h1>", "", 1)
			id = strings.TrimSpace(id)
			currentAyah, _ = strconv.ParseInt(id, 10, 64)
			continue
		}
		// check if line in arabic
		if match := regexArabic.FindAllString(line, -1); len(match) > 0 {
			line = strings.Replace(line, "<p>", `<p dir="rtl">`, 1)
		}
		// append line to approtiate ayah description based on id
		ayahInfos[currentAyah] += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// build query
	var query = ""
	for i, info := range ayahInfos {
		if i > 0 {
			query += strings.TrimSpace(fmt.Sprintf(`UPDATE quran_ayahs SET idn_ayah_info = "%s" WHERE id = %s;`, Addslashes(info), (strconv.Itoa(i)))) + "\n"
		}
	}

	// write sql file
	err3 := os.WriteFile("sql/6_ayah-info.sql", []byte(query), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("ayah-info.sql written successfully.")
}
