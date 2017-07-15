package service

import (
	"github.com/RomanosTrechlis/blog-generator/config"
)

func MapSiteInformationToRequest(siteInfo *config.SiteInformation) (req *GenerateRequest) {
	staticPages := make([]*StaticPages, 0)
	for _, s := range siteInfo.StaticPages {
		sp := &StaticPages{
			File:       s.File,
			To:         s.To,
			IsTemplate: s.IsTemplate,
		}
		staticPages = append(staticPages, sp)
	}
	req = &GenerateRequest{
		Author:            siteInfo.Author,
		BlogURL:           siteInfo.BlogURL,
		BlogLanguage:      siteInfo.BlogLanguage,
		BlogDescription:   siteInfo.BlogDescription,
		DateFormat:        siteInfo.DateFormat,
		ThemeFolder:       siteInfo.ThemeFolder,
		BlogTitle:         siteInfo.BlogTitle,
		NumPostsFrontPage: int64(siteInfo.NumPostsFrontPage),
		TempFolder:        siteInfo.TempFolder,
		DestFolder:        siteInfo.DestFolder,
		StaticPages:       staticPages,
		DataSource: &DataSource{
			Repository: siteInfo.DataSource.Repository,
			Type:       siteInfo.DataSource.Type,
		},
		Theme: &Theme{
			Repository: siteInfo.Theme.Repository,
			Type:       siteInfo.Theme.Type,
		},
	}
	return req
}

func MapRequestToSiteInformation(req *GenerateRequest) (siteInfo config.SiteInformation) {
	staticPages := make([]config.StaticPage, 0)
	for _, s := range req.StaticPages {
		sp := config.StaticPage{
			File: s.File,
			To: s.To,
			IsTemplate: s.IsTemplate,
		}
		staticPages = append(staticPages, sp)
	}

	siteInfo = config.SiteInformation{
		Author: req.Author,
		BlogURL: req.BlogURL,
		BlogLanguage: req.BlogLanguage,
		BlogDescription: req.BlogDescription,
		DateFormat: req.DateFormat,
		ThemeFolder: req.ThemeFolder,
		BlogTitle: req.BlogTitle,
		NumPostsFrontPage: int(req.NumPostsFrontPage),
		TempFolder: req.TempFolder,
		DestFolder: req.DestFolder,
		StaticPages: staticPages,
		DataSource: config.DataSource{
			Repository: req.DataSource.Repository,
			Type: req.DataSource.Type,
		},
		Theme: config.Theme{
			Repository: req.Theme.Repository,
			Type: req.Theme.Type,
		},
	}
	return siteInfo
}
