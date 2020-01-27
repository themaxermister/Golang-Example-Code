package main

import (
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seller struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	UserID     string             `bson:"userid" json:"userid"`
	LastUpdate time.Time          `bson:"lastupdate" json:"lastupdate"`
	Profiles   []interface{}         `bson:"profiles" json:"profiles"`
}

type Juristic struct {
	Name        string    `bson:"name" json:"name"`
	Type        string    `bson:"type" json:"type"`
	TaxID       string    `bson:"vat" json:"vat"`
	DBD         string    `bson:"dbd" json:"dbd"`
	DMS         time.Time `bson:"directmarketstart" json:"directmarketstart"`
	DME         time.Time `bson:"directmarketend" json:"directmarketend"`
	Phone       string    `bson:"phone" json:"phone"`
	Address     string    `bson:"address" json:"address"`
	Subdistrict string    `bson:"subdistrict" json:"subdistrict"`
	District    string    `bson:"district" json:"district"`
	Province    string    `bson:"province" json:"province"`
	Postal      string    `bson:"postal" json:"postal" `
	BranchCode  string    `bson:"branchcode" json:"branchcode"`
	Email       string    `bson:"email json:"email"`
	EffDate     time.Time `bson:"effectivedate" json:"effectivedate"`
	Verify      bool      `bson:"verified" json:"verified"`
}

type Ordinary struct {
	Name        string    `bson:"name" json:"name"`
	Type        string    `bson:"type" json:"type"`
	TaxID       string    `bson:"vat" json:"vat"`
	Phone       string    `bson:"phone" json:"phone"`
	Address     string    `bson:"address" json:"address"`
	Subdistrict string    `bson:"subdistrict" json:"subdistrict"`
	District    string    `bson:"district" json:"district"`
	Province    string    `bson:"province" json:"province"`
	Postal      string    `bson:"postal" json:"postal" `
	Email       string    `bson:"email json:"email"`
	EffDate     time.Time `bson:"effectivedate" json:"effectivedate"`
	Verify      bool      `bson:"verified" json:"verified"`
}

func main(){
	var sell Seller
	sell.UserID = "223233"
	var ord Ordinary
	ord.Name = "Jone May"
	var jur Juristic
	jur.Name = "Gemma Lee"

	fmt.Println(sell)
	sell.Profiles = []interface{}{ord, jur}
	fmt.Println(sell)
}