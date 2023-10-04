package file

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func GetUploadsDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}

	uploads := filepath.Join(dir, "uploads")
	// uploads := filepath.Join(dir, "uploads/"+time.Now().Format(time.DateOnly))
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uploads, nil
}

func GetFileList() ([]string, error) {
	dir, _ := GetUploadsDir()
	return filepath.Glob(dir)
}

func RemoveLogFile(fileName string) error {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}

	return os.Remove(filepath.Join(dir, fileName))
}

type HistoryTiem struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Time    int64  `json:"time"`
	IP      string `json:"ip"`
}

func getHistoryPath() string {

	dir, _ := GetUploadsDir()
	path := filepath.Join(dir, "history.json") // 文件路径
	checkHistoryFile(path)
	log.Println(path)
	return path
}

func checkHistoryFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		emptyData := []byte("[]")
		os.WriteFile(path, emptyData, os.ModePerm)
		os.Create(path)
	}
}

func getHistoryFile() (*os.File, error) {
	filePath := getHistoryPath() // 文件路径
	_, err := os.Stat(filePath)
	if err != nil {
		// return
		os.Create(filePath)
	}

	return os.Open(filePath)
}

func GetHistory() ([]HistoryTiem, error) {
	var list []HistoryTiem
	path := getHistoryPath()
	data, _ := os.ReadFile(path)
	err := json.Unmarshal(data, &list)
	if err != nil {
		log.Println("解析JSON数据时出错:", err)
	}
	return list, nil
}

func WhiteHistory(item HistoryTiem) error {
	list, err := GetHistory()
	path := getHistoryPath()

	if err != nil {
		return err
	}

	list = append(list, item)
	log.Print("新数据：", list)

	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		log.Print("写入失败", err)
	}

	return err
}
