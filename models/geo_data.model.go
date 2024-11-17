package models

import (
	"time"

	"github.com/google/uuid"
)

// GeoData represents the model for geospatial data.
type GeoData struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID as the primary key
	UserID     uuid.UUID `gorm:"type:uuid;not null"`                              // Foreign key reference to the User
	Title      string    `gorm:"type:varchar(255);not null"`                      // Title for the geospatial data
	Description string   `gorm:"type:text"`                                       // Optional description
	FileType   string    `gorm:"type:varchar(50);not null"`                       // File type (e.g., GeoJSON, KML)
	FilePath   string    `gorm:"type:varchar(500);not null"`                      // File path to the uploaded file
	Data       string    `gorm:"type:jsonb;not null"`                             // GeoJSON data stored as JSONB in PostgreSQL
	CreatedAt  time.Time `gorm:"not null"`                                        // Timestamp for when the record was created
	UpdatedAt  time.Time `gorm:"not null"`                                        // Timestamp for when the record was last updated
}


// GeoDataInput represents the required fields for creating or updating GeoData.
type GeoDataInput struct {
	Title       string `json:"title" binding:"required"`                  // Title is required
	Description string `json:"description"`                               // Optional description
	FileType    string `json:"file_type" binding:"required,oneof=GeoJSON KML"` // Must be either GeoJSON or KML
	FilePath    string `json:"file_path" binding:"required"`              // File path is required
	Data        string `json:"data" binding:"required"`                   // GeoJSON data is required
}

// GeoDataResponse represents the safe fields returned in API responses.
type GeoDataResponse struct {
	ID          uuid.UUID `json:"id"`                                    // GeoData ID
	Title       string    `json:"title"`                                 // Title of the geospatial data
	Description string    `json:"description"`                           // Description of the geospatial data
	FileType    string    `json:"file_type"`                             // File type (GeoJSON or KML)
	FilePath    string    `json:"file_path"`                             // Path to the uploaded file
	Data        string    `json:"data"`                                  // GeoJSON data
	CreatedAt   time.Time `json:"created_at"`                            // Timestamp for record creation
	UpdatedAt   time.Time `json:"updated_at"`                            // Timestamp for last record update
}