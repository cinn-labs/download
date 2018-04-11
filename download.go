package download

import (
	"fmt"
	"os"
	"strings"

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
	ffn := strings.Replace(fn, " ", "_", -1)
	ffn = strings.Replace(ffn, "#", "", -1)
	df := DownloadFolder()
	if _, err := os.Stat(df); os.IsNotExist(err) {
		os.Mkdir(df, 0700)
	}
	folderPath := fmt.Sprintf("%s%s", df, uuid.NewV4().String())
	err := os.Mkdir(folderPath, 0700)
	filePath := fmt.Sprintf("%s/%s.%s", folderPath, ffn, fe)
	return filePath, err
}
