package domain

import (
	"context"
)

// Shape entity represents the domain model for a custom drawn shape.
type Shape struct {
	ID         string // Unique identifier for the shape (UUID or Serial from PostgreSQL)
	UserID     string // The User ID associated with the shape
	GeoDataID  string // The ID of the associated GeoData (GeoJSON/KML file)
	Geometry   string // The geometry of the shape (in GeoJSON format)
	Attributes string // Additional attributes for the shape (metadata)
}

// ShapeRepository defines the interface for shape-related operations.
type ShapeRepository interface {
	Create(ctx context.Context, shape *Shape) error
	FetchAll(ctx context.Context, userID string) ([]Shape, error) // Fetch all shapes for a user
	GetByID(ctx context.Context, id string) (Shape, error) // Get a shape by its ID
	GetByGeoDataID(ctx context.Context, geoDataID string) ([]Shape, error) // Fetch shapes by associated GeoData ID
	Update(ctx context.Context, shape *Shape) error
	Delete(ctx context.Context, id string) error
}