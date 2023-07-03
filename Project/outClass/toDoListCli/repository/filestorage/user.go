package filestorage

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todolist/entity"
)

type FileStore struct {
	path              string
	serializationMode string
}

func New(serializationMode string) FileStore {
	serializationMode = checkFormatFile(serializationMode)
	path := "./" + "data." + serializationMode

	return FileStore{path: path, serializationMode: serializationMode}
}

func checkFormatFile(serializationMode string) string {
	var fileFormats = []string{"data.csv", "data.json", "data.xml", "data.txt"}
	var formatName string

	for _, v := range fileFormats {
		_, err := os.Stat(v)
		if !os.IsNotExist(err) {
			formatName = v
			break
		}
	}

	// datase not exist
	if formatName == "" {
		formatName = serializationMode

		// datase exist
	} else {
		formatName = strings.Split(formatName, ".")[1]

		// fmt.Printf("file with format %s is Exist!\nso Data is save with format: | %s |\n",
		// 	formatName, strings.ToUpper(strings.Split(formatName, ".")[1]))
	}

	return formatName
}

func (f FileStore) Save(newUser entity.User) error {
	//Choose format to save into file
	var data []byte

	switch f.serializationMode {
	case "json":
		if d, err := json.Marshal(&newUser); err != nil {

			return fmt.Errorf("could not convert data to json")
		} else {
			data = d
		}

	case "xml":
		if d, err := xml.Marshal(&newUser); err != nil {

			return fmt.Errorf("could not convert data to xml")
		} else {
			data = d
		}

	case "csv":
		data = []byte(fmt.Sprintf("%s,%d,%s", newUser.Email, newUser.UserID, newUser.Password))

	case "txt":
		data = []byte(fmt.Sprintf("id: %d, email: %s, password: %s", newUser.UserID, newUser.Email, newUser.Password))

	default:

		return fmt.Errorf("format to save file is invalid")
	}

	//write data to the file
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		return fmt.Errorf("could not open dataset")
	}

	defer file.Close()

	if _, err2 := file.WriteString(string(data) + "\n"); err2 != nil {

		return fmt.Errorf("could not write data to dataset")
	}

	return nil
}

func (f FileStore) Load() ([]entity.User, error) {
	var users []entity.User

	// open file
	file, err := os.Open(f.path)
	if err != nil {
		return []entity.User{}, fmt.Errorf("error when open Dataset")
	}
	defer file.Close()

	var data string
	var fetchUser entity.User
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var err error
		// delete "\n" of end of each line
		data = strings.Split(scanner.Text(), "\n")[0]

		switch f.serializationMode {
		case "json":
			err = json.Unmarshal([]byte(data), &fetchUser)
			if err != nil {

				return []entity.User{}, fmt.Errorf("datas in dataset is wrong format for JSON")
			}

		case "xml":
			err = xml.Unmarshal([]byte(data), &fetchUser)
			if err != nil {

				return []entity.User{}, fmt.Errorf("datas in dataset is wrong format for JSON")
			}

		case "csv":
			fetchUser.Email = strings.Split(data, ",")[0]
			fetchUser.UserID, err = strconv.Atoi(strings.Split(data, ",")[1])
			if err != nil {

				return []entity.User{}, fmt.Errorf("datas in dataset is wrong format for CSV")
			}
			fetchUser.Password = strings.Split(data, ",")[2]

		case "txt":
			fetchUser.UserID, err = strconv.Atoi(strings.Split(strings.Split(data, ",")[0], ": ")[1])
			if err != nil {
				return []entity.User{}, fmt.Errorf("datas in dataset is wrong format for TXT")

			}
			fetchUser.Email = strings.Split(strings.Split(data, ",")[1], ": ")[1]
			fetchUser.Password = strings.Split(strings.Split(data, ",")[2], ": ")[1]
		}

		users = append(users, fetchUser)
	}

	return users, nil
}

func (f FileStore) ReturnSerialandPath() (string, string) {
	return f.serializationMode, f.path
}
