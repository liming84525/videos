package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	orm "github.com/jinzhu/gorm"
)

type Video struct {
	Title    string `json:"videoname,omitempty" gorm:"type:varchar(100);not null"`
	Id       string `json:"cloudname,omitempty" gorm:"index;type:varchar(100);not null"`
	ImageUrl string `json:"imageurl,omitempty" gorm:"type:varchar(100);not null"`
	VideoUrl string `json:"videourl,omitempty" gorm:"type:varchar(100);not null"`
}

type Series struct {
	Id     string  `json:"subject_id,omitempty" gorm:"primary_key;type:varchar(100);not null"`
	Title  string  `json:"subject,omitempty" gorm:"type:varchar(100);not null"`
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
		S Series `json:"videoinfo,omitempty"`
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
	if !db.HasTable(&Video{}) {
		db.CreateTable(&Video{})
	}
	if !db.HasTable(&Series{}) {
		db.CreateTable(&Series{})
	}
}

func main() {
	serie := obj.S
	for _, video := range serie.Videos {
		db.Save(&video)
	}
	db.Save(&serie)
}
