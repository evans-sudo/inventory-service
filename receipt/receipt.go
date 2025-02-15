package receipt


import (
	"os"
	"path/filepath"
)

var ReceiptDirectory string = filepath.Join("uploads")

type Receipt struct {
	ReceiptNamec string `json:"Name"`
	UploadDate   string `json:"UploadDate"`
}


func GetReceipts() ([]Receipt, error) {
	reciept := make([]Receipt, 0)
	files, err := os.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		reciept = append(reciept, Receipt{ReceiptNamec: file.Name(), UploadDate: info.ModTime().String()})
	}
	return reciept, nil
}