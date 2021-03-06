package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConfigFile contains information about the site
var ConfigFile string

// SiteInfo contains all the variables for the site
var SiteInfo SiteInformation

type BlogInformation interface {
}

// SiteInformation contains the information inside ConfigFile
type SiteInformation struct {
	Author            string `json:Author`
	BlogURL           string `json:BlogURL`
	BlogLanguage      string `json:BlogLanguage`
	BlogDescription   string `json:BlogDescription`
	DateFormat        string `json:DateFormat`
	Theme             Theme
	ThemeFolder       string `json:ThemeFolder`
	BlogTitle         string `json:BlogTitle`
	NumPostsFrontPage int    `json:NumPostsFrontPage`
	DataSource        DataSource
	Upload            Upload
	TempFolder        string       `json:TempFolder`
	DestFolder        string       `json:DestFolder`
	StaticPages       []StaticPage `json:StaticPages`
}

type Theme struct {
	Repository string `json:Repository`
	Type       string `json:Type`
}

type StaticPage struct {
	File       string `json:File`
	To         string `json:To`
	IsTemplate bool   `json:IsTemplate`
}

type DataSource struct {
	Type       string `json:Type`
	Repository string `json:Repository`
}

type Upload struct {
	URL      string `json:URL`
	Username string `json:Username`
	Password string `json:Password`
}

func NewSiteInformation(configFile string) (siteInfo SiteInformation) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("error accessing directory %s: %v", configFile, err)
	}
	siteInfo.ParseJSON(data)
	return fillDefaultValues(siteInfo)
}

func (c *SiteInformation) ParseJSON(b []byte) (err error) {
	return json.Unmarshal(b, &c)
}

func fillDefaultValues(siteInfo SiteInformation) SiteInformation {
	if siteInfo.TempFolder == "" {
		siteInfo.TempFolder = "./tmp"
	}
	if siteInfo.DestFolder == "" {
		siteInfo.DestFolder = "./public"
	}
	if siteInfo.ThemeFolder == "" {
		siteInfo.DestFolder = "./static"
	}
	if siteInfo.NumPostsFrontPage == 0 {
		siteInfo.NumPostsFrontPage = 10
	}
	return siteInfo
}
