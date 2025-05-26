package internal

import (
	"fmt"
)

func GenerateQuranAyah() {
	Copy("data/quran-text.sql", "sql/2_quran-ayah.sql")
	fmt.Println("quran-ayah.sql written successfully.")
}
