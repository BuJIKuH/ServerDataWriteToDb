package config

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
	"time"
)

type ServerDbValue struct {
	Index      int    `json:"index"`
	ConfigName string `json:"config_name"`
	Data       string `json:"data"`
}

func ResponseServer() error {

	// Считываем содержимое ключа
	confServ := NewServer()

	key, err := os.ReadFile(confServ.KeyPath)
	if err != nil {
		log.Fatalf("Не удалось считать ключ: %s", err)
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(confServ.Password))
	if err != nil {
		log.Fatalf("Не удалось считать ключ потому, что %s", err)
	}
	// Создаем конфиг для подключения к SSH серверу
	sshConfig := &ssh.ClientConfig{
		User: confServ.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Подключаемся к SSH серверу
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", confServ.Server, confServ.Port), sshConfig)
	if err != nil {
		log.Fatalf("Не удалось подключиться к SSH серверу: %s", err)
	}
	defer conn.Close()

	// Создаем новую сессию SSH
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Не удалось создать сессию SSH: %s", err)
	}
	defer session.Close()

	// Выполняем команду на удаленном сервере
	out, err := session.Output("ls")
	if err != nil {
		log.Fatalf("Ошибка выполнения команды: %s", err)
	}

	scannerLines := strings.Split(string(out), "\n")

	InsertToDatabase(scannerLines)

	return nil

}

func InsertToDatabase(data []string) {
	dt := time.Now()
	configs := make([]ServerDbValue, 0)

	for i, line := range data {
		if line != "" {
			scannerLine := strings.Split(line, " ")
			newConfig := ServerDbValue{
				Index:      i + 1,
				ConfigName: scannerLine[0],
				Data:       dt.Format("02:01:2006 15:04:05"),
			}
			configs = append(configs, newConfig)
		}

	}
	if configs != nil {
		DataBase.DB.Exec("DELETE FROM server_db_values")
		defer DataBase.DB.Create(configs)
	}
	log.Fatal("Данные не получены")

}
