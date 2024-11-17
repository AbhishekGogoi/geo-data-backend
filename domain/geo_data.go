package domain

import (
	"context"
)

// GeoData entity represents the domain model for geospatial data (GeoJSON or KML file).
type GeoData struct {
	ID       string // Unique identifier for GeoData (UUID or Serial from PostgreSQL)
	UserID   string // The User ID associated with the uploaded file
	FileName string // Name of the file (GeoJSON/KML)
	FileType string // File type (GeoJSON/KML)
	FilePath string // Path to the stored file (e.g., local or cloud storage)
}

// GeoDataRepository defines the interface for geo-data-related operations.
type GeoDataRepository interface {
	Create(ctx context.Context, geoData *GeoData) error
	FetchAll(ctx context.Context) ([]GeoData, error)
	GetByID(ctx context.Context, id string) (GeoData, error)
	GetByUserID(ctx context.Context, userID string) ([]GeoData, error) // Get files uploaded by a user
	Update(ctx context.Context, geoData *GeoData) error
	Delete(ctx context.Context, id string) error
}