package lib

import (
	"flag"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Configuration struct {
	Address      string `json:"address"`
	Port         string `json:"port"`
	ReadTimeout  string `json:"readtimeout"`
	WriteTimeout string `json:"writetimeout"`
	Static       string `json:"static"`
	Database     struct {
		DbAddr string `json:"host"`
		DbPort string `json:"dbport"`
		DbName string `json:"dbname"`
	} `json:"database"`
}

type AddressDoc struct {
	_id              primitive.ObjectID `json:"_id" bson:"_id"`
	Subdistrict      string             `json:"subdistrict" bson:"subdistrict"`
	Subdistrict_en   string             `json:"subdistrict_en" bson:"subdistrict_en"`
	Subdistrict_code int32              `json:"subdistrict_code" bson:"subdistrict_code"`
	District         string             `json:"district" bson:"district"`
	District_en      string             `json:"district_en" bson:"district_en"`
	District_code    int32              `json:"district_code" bson:"district_code"`
	Province         string             `json:"province" bson:"province"`
	Province_en      string             `json:"province_en" bson:"province_en"`
	Province_code    int32              `json:"province_code" bson:"province_code"`
	Postal           int32              `json:"zipcode" bson:"zipcode"`
}

type AddressDetails struct {
	Main     string `bson:"main" json:"main"`
	Building string `bson:"building,omitempty" json:"building,omitempty"`
	Village  string `bson:"village,omitempty" json:"village,omitempty"`
	Moo      string `bson:"moo,omitempty" json:"moo,omitempty"`
	Soi      string `bson:"soi,omitempty" json:"soi,omitempty"`
	Road     string `bson:"road,omitempty" json:"road,omitempty"`
}

type Seller struct {
	Userid       string         `bson:"userid" json:"userid"`
	Username     string         `bson:"username" json:"username"`
	Name         string         `bson:"name" json:"name"`
	Type         string         `bson:"type" json:"type"`
	Vat          string         `bson:"vat" json:"vat"`
	TaxID        string         `bson:"taxid" json:"taxid"`
	Address      AddressDetails `bson:"address" json:"address"`
	Phone        string         `bson:"phone" json:"phone"`
	Email        []string       `bson:"email" json:"email"`
	Subdistrict  string         `bson:"subdistrict" json:"subdistrict"`
	District     string         `bson:"district" json:"district"`
	Province     string         `bson:"province" json:"province"`
	Postal       int32          `bson:"postal" json:"postal"`
	BranchCode   string         `bson:"branchcode,omitempty" json:"branchcode,omitempty"`
	DBD          string         `bson:"dbd,omitempty" json:"dbd,omitempty"`
	DMS          time.Time      `bson:"directmarketstart,omitempty" json:"directmarketstart,omitempty"`
	DME          time.Time      `bson:"directmarketend,omitempty" json:"directmarketend,omitempty"`
	EffDate      time.Time      `bson:"effectivedate" json:"effectivedate"`
	Status       string         `bson:"status" json:"status"`
	Reason       []string       `bson:"reason,omitempty" json:"reason,omitempty"`
	Default      bool           `bson:"default" json:"default"`
	CreateDate   time.Time      `bson:"createdate" json:"createdate"`
	LastUpdate   time.Time      `bson:"lastupdate" json:"lastupdate"`
	VerifiedTime time.Time      `bson:"verifiedtime,omitempty" json:"verifiedtime,omitempty"`
}

type FullProfile struct {
	Profileinfo     *Seller `json:"profileinfo" bson:"profileinfo"`
	Reviewed        int64   `json:"reviewed" bson:"reviewed"`
	ProcessStatus   bool    `json:"processstatus" bson:"processstatus"`
	Subdistrictcode int32   `json:"subdistrict_code" bson:"subdistrict_code"`
	Districtcode    int32   `json:"district_code" bson:"district_code"`
	Provincecode    int32   `json:"province_code" bson:"province_code"`
}

type StatusProfile struct {
	Userid       string    `json:"userid" bson:"userid"`
	UserName     string    `bson:"username" json:"username"`
	Taxid        string    `bson:"taxid" json:"taxid"`
	Email        string    `bson:"email" json:"email"`
	Status       string    `bson:"status" json:"status"`
	CreateDate   time.Time `bson:"createdate" json:"createdate"`
	VerifiedTime time.Time `bson:"verifiedtime,omitempty" json:"verifiedtime,omitempty"`
	Reason       []string  `bson:"reason,omitempty" json:"reason,omitempty"`
}

var Config = Configuration{}
var Client *mongo.Client
var TempDir = "./public/data/temp/"
var UploadDir = flag.String("d", "./public/data/temp", "Upload directory, defaults to './public/data/temp'")

var SellerProfile Seller

//var id = 1
//var Userid = "9260820"

/*var TempProfile L.ProfileApp
TempProfile.Name = "John Mary"
TempProfile.Type = "juristic"
TempProfile.Vat = "no"
TempProfile.TaxID = "7285398444458"
TempProfile.Address.Main = "137/89-90"
TempProfile.Address.Building = "เธซเธกเธนเนเธ"
TempProfile.Address.Village = "RANDOM STUFF"
TempProfile.Address.Soi = "RANDOM STUFF"
TempProfile.Address.Moo = "RANDOM STUFF"
TempProfile.Address.Road = "เธตเน 5"
TempProfile.Phone = "0775269943"
TempProfile.Email = []string{"random3@gmail.com", "random3@gmail.com"}
TempProfile.Subdistrict = "แขวงพระบรมมหาราชวัง"
TempProfile.District = "เขตพระนคร"
TempProfile.Province = "กรุงเทพมหานคร"
TempProfile.Postal = 10200
TempProfile.BranchCode = "00001"
TempProfile.Verify = false
TempProfile.Default = false
layout := "2006-01-02T15:04:05Z"
layout2 := "2006-01-02T15:04:05.999999-07:00"
TempProfile.CreateDate, err = time.Parse(layout, "2019-10-12T05:28:35Z")
TempProfile.LastUpdate, err = time.Parse(layout2, "2020-01-15T17:45:07.2083622+07:00")
TempProfile.DBD = "2154678904"
TempProfile.DMS = time.Now()
TempProfile.DME = time.Now()
TempProfile.EffDate, err = time.Parse(layout, "2019-10-13T00:00:00Z")
TempProfile.VerifiedTime = time.Time{}*/
