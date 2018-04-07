package download

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
)

const TMP_DIR = "tmp/"

func DownloadFolder() string {
	df := os.Getenv("DOWNLOAD_PATH")
	if df == "" {
		df = TMP_DIR
	}
	return df
}

func GenerateFolder(fn string, fe string) (string, error) {
	df := DownloadFolder()
	if _, err := os.Stat(df); os.IsNotExist(err) {
		os.Mkdir(df, 0700)
	}
	folderPath := fmt.Sprintf("%s%s", df, uuid.NewV4().String())
	err := os.Mkdir(folderPath, 0700)
	filePath := fmt.Sprintf("%s/%s.%s", folderPath, fn, fe)
	return filePath, err
}
