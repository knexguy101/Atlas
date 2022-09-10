package globals

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strings"
)


const (
	ProfileFilePath         string = "storage/profiles.json"
	AccountFilePath         string = "storage/accounts.json"
	ProxiesFilePath			string = "storage/proxies.txt"
	TasksFilePath			string = "storage/tasks.json"
	StorageFolderPath		string = "storage"
)
var (
	appDataPath = ""
)

func LoadFiles() {
	adp, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	adp = path.Join(adp, "Atlas")
	appDataPath = adp
	SafeCreateFolder("")
	SafeCreateFolder(StorageFolderPath)

	SafeCheckFile(TasksFilePath)
	SafeCheckFile(ProxiesFilePath)
	SafeCheckFile(ProfileFilePath)
	SafeCheckFile(AccountFilePath)
}

func SafeWriteFile(p string, text []byte) error {
	file, err := os.OpenFile(path.Join(appDataPath, p), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Write(text)
	if err != nil {
		return err
	}

	return nil
}

func SafeCreateFolder(folder string) {
	if _, err := os.Stat(path.Join(appDataPath, folder)); os.IsNotExist(err) {
		mer := os.Mkdir(path.Join(appDataPath, folder), os.ModePerm)
		if mer != nil {
			log.Println(mer.Error())
		}
	}
}

func SafeCreateFile(taskFile string) (*os.File, error) {
	var _, taskErr = os.Stat(path.Join(appDataPath, taskFile))
	if os.IsNotExist(taskErr) {
		var file, err = os.Create(path.Join(appDataPath, taskFile))
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	file, err := os.OpenFile(path.Join(appDataPath, taskFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SafeCopyFile(ogFile string, newFile *os.File) error {
	var _, taskErr = os.Stat(path.Join(appDataPath, ogFile))
	if os.IsNotExist(taskErr) {
		return errors.New("ogFile does not exist")
	}

	input, err := os.ReadFile(path.Join(appDataPath, ogFile))
	if err != nil {
		return err
	}

	_, err = newFile.Write(input)
	if err != nil {
		return err
		// return
	}

	return nil
}

func SafeDeleteFile(file string) {
	err := os.Remove(path.Join(appDataPath, file))
	if err != nil {
		return
	}
}

func SafeCheckFile(taskFile string) error {
	var _, taskErr = os.Stat(path.Join(appDataPath, taskFile))
	if os.IsNotExist(taskErr) {
		var file, err = os.Create(path.Join(appDataPath, taskFile))
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func SafeFileExists(taskFile string) bool {
	var _, taskErr = os.Stat(path.Join(appDataPath, taskFile))
	return !os.IsNotExist(taskErr)
}

func SafeReadFile(taskFile string) ([]byte, error) {
	//SafeCheckFile(path.Join(appDataPath, taskFile))

	input, err := os.ReadFile(path.Join(appDataPath, taskFile))
	if err != nil {
		return nil, err
	}
	return input, nil
}

func SafeReadLines(taskFile string) ([]string, error) {
	taskFile = strings.ReplaceAll(taskFile, `\`, `/`)
	file, err := os.Open(path.Join(appDataPath, taskFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetAppData() string {
	return appDataPath
}
