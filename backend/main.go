package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"

	// "gorm.io/driver/clickhouse"
	// "gorm.io/gorm"

	waLog "go.mau.fi/whatsmeow/util/log"
)

var click_db_config_path string = "./config/database.yaml"

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

// func testGetProjectItems() {

// }

func main() {

	// var database_55_cred map[string]config.DatabaseCredentials = config.Import_yaml_config(click_db_config_path)
	// var dsn string = fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s", database_55_cred["CLICKHOUSE.55"].User, database_55_cred["CLICKHOUSE.55"].Passwd, database_55_cred["CLICKHOUSE.55"].Host, database_55_cred["CLICKHOUSE.55"].Port, database_55_cred["CLICKHOUSE.55"].DB)
	// db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var project_items_result []models.ProjectItems

	// db.Table("imas.project_items").Where("s_date between ? AND ?", time.Now().Add(time.Duration(-10)*time.Minute), time.Now()).Limit(1000).Find(&project_items_result)

	dbLog := waLog.Stdout("Database", "DEBUG", true)

	connection_string := "file:examplestore.db?_foreign_keys=on"

	container, err := sqlstore.New("sqlite3", connection_string, dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// echo 2@... | qrencode -t ansiutf8
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	}

	err = client.Connect()
	if err != nil {
		panic(err)
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
