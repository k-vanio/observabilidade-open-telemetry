package domain

import (
	"context"

	"github.com/k-vanio/observabilidade-open-telemetry/service-two/internal/dto"
)

type ZipCode interface {
	Search(ctx context.Context, request dto.SearchRequest) dto.SearchResponse
}
