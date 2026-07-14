package storage

import "context"

// CloudStorage is our interface. The Business Logic only talks to this.
type CloudStorage interface {
    Upload(ctx context.Context, bucket string, key string, data []byte) error
}

// S3Adapter implements the CloudStorage interface for AWS.
type S3Adapter struct { /* AWS SDK Client */ }

func (s *S3Adapter) Upload(ctx context.Context, b, k string, d []byte) error {
    // AWS specific implementation
    return nil
}

// GCSAdapter implements the CloudStorage interface for GCP.
type GCSAdapter struct { /* Google Cloud SDK Client */ }

func (g *GCSAdapter) Upload(ctx context.Context, b, k string, d []byte) error {
    // GCP specific implementation
    return nil
}

// Service uses the interface, totally unaware of the underlying Cloud Provider.
type FileProcessor struct {
    Storage CloudStorage
}
