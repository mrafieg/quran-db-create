package mypackage

import "fmt"

func GenerateQuranText() {
	Copy("data/quran-text.sql", "sql/quran-text.sql")
	fmt.Println("quran-text.sql written successfully.")
}
