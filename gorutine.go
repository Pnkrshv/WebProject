package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type User struct {
	ID    int
	Email string
	Logs  []logItem
}

type logItem struct {
	Action    string
	timestamp time.Time
}

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity log:\n", u.ID, u.Email)
	for i, item := range u.Logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.Action, item.timestamp)
	}
	return out
}

func main() {
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)

	for _, user := range users {
		saveUserInfo(user)
	}
}

func saveUserInfo(user User) error {
	fmt.Printf("Writting file for user ID: %d\n", user.ID)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.ID)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	file.WriteString(user.getActivityInfo())
	return nil
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			ID:    i + 1,
			Email: fmt.Sprintf("user%d@mail.ru", i+1),
			Logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			Action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}
	return logs
}
