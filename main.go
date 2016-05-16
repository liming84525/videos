package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	orm "github.com/jinzhu/gorm"
)

type Video struct {
	orm.Model
	Title    string `json:"videoname,omitempty"`
	Id       string `json:"cloudname,omitempty"`
	ImageUrl string `json:"imageurl,omitempty"`
	VideoUrl string `json:"videourl,omitempty"`
}

type Series struct {
	orm.Model
	Id     string  `json:"subject_id,omitempty"`
	Title  string  `json:"subject,omitempty"`
	Videos []Video `json:"details,omitempty"`
}

var (
	obj struct {
		S Series `json:"videoinfo,omitempty"`
	}
	db *orm.DB
)

func init() {
	bytes, err := ioutil.ReadFile("video.json")
	obj = struct {
		S series `json:"videoinfo,omitempty"`
	}{}
	if err != nil {
		log.Println(err)
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		log.Println(err)
	}
	log.Printf("result is : %+v", obj)
	db, err = orm.Open("mysql", "root:@tcp(127.0.0.1:3306)/weixin?charset=utf8mb4")
	if err != nil {
		log.Println(err)
	}
	db.LogMode(true)
	if !db.HasTable(&Series{}) {
		db.DropTableIfExists(&Series{})
		db.CreateTable(&Series{})
	}
}

func main() {

	log.Println(obj)
}
