package mypackage

import "fmt"

func GenerateQuranAyah() {
	Copy("data/quran-text.sql", "sql/quran-ayah.sql")
	fmt.Println("quran-ayah.sql written successfully.")
}
