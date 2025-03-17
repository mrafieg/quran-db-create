package mypackage

import "fmt"

func GenerateQuranAyah() {
	Copy("data/quran-ayah.sql", "sql/quran-ayah.sql")
	fmt.Println("quran-ayah.sql written successfully.")
}
