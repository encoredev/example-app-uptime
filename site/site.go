package site

import (
	"context"

	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Site describes a monitored site.
type Site struct {
	// ID is a unique ID for the site.
	ID int `json:"id"`
	// URL is the site's URL.
	URL string `json:"url"`
}

// AddParams are the parameters for adding a site to be monitored.
type AddParams struct {
	// URL is the URL of the site. If it doesn't contain a scheme
	// (like "http:" or "https:") it defaults to "https:".
	URL string `json:"url"`
}

// Add adds a new site to the list of monitored websites.
//
//encore:api public method=POST path=/site
func (s *Service) Add(ctx context.Context, p *AddParams) (*Site, error) {
	site := &Site{URL: p.URL}
	if err := s.db.Create(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}

type ListResponse struct {
	// Sites is the list of monitored sites.
	Sites []*Site `json:"sites"`
}

// Get gets a site by id.
//
//encore:api public method=GET path=/site/:siteID
func (s *Service) Get(ctx context.Context, siteID int) (*Site, error) {
	var site Site
	if err := s.db.Where("id = $1", siteID).First(&site).Error; err != nil {
		return nil, err
	}
	return &site, nil
}

// Delete deletes a site by id.
//
//encore:api public method=DELETE path=/site/:siteID
func (s *Service) Delete(ctx context.Context, siteID int) error {
	return s.db.Delete(&Site{ID: siteID}).Error
}

// List lists the monitored websites.
//
//encore:api public method=GET path=/site
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var sites []*Site
	if err := s.db.Find(&sites).Error; err != nil {
		return nil, err
	}
	return &ListResponse{Sites: sites}, nil
}

//encore:service
type Service struct {
	db *gorm.DB
}

var siteDB = sqldb.Named("site").Stdlib()

func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: siteDB,
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}
