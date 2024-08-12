package watermark

import (
	"context"

	"github.com/arslanoktay/WatermarkMicroservices.git/internal"
)

// With Go-Kit framework, a service always represent by interface
type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
