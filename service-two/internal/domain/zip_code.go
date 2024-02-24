package domain

import (
	"github.com/k-vanio/observabilidade-open-telemetry/service-two/internal/dto"
)

type ZipCode interface {
	Search(request dto.SearchRequest) dto.SearchResponse
}
