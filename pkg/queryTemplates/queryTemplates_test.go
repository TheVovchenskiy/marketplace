package queryTemplates_test

import (
	"marketplace/pkg/price"
	"marketplace/pkg/queryTemplates"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAdQuery(t *testing.T) {
	tests := []struct {
		name           string
		pageNum        int
		resultsPerPage int
		sortField      string
		sortOrder      string
		minPrice       string
		maxPrice       string
		expectedQuery  string
		expectedError  error
	}{
		{
			"only min_price",
			1,
			10,
			"created_at",
			"desc",
			"0.00",
			"",
			`SELECT
		id,
		author_id,
		"name",
		description,
		cents_price,
		picture_url,
		created_at
	FROM
		public.ad
	WHERE cents_price >= $1
	ORDER BY created_at DESC
	LIMIT $2
	OFFSET $3`,
			nil,
		},
		{
			"min_price and max_price",
			1,
			10,
			"created_at",
			"desc",
			"0.00",
			"10.00",
			`SELECT
		id,
		author_id,
		"name",
		description,
		cents_price,
		picture_url,
		created_at
	FROM
		public.ad
	WHERE cents_price >= $1 AND cents_price <= $2
	ORDER BY created_at DESC
	LIMIT $3
	OFFSET $4`,
			nil,
		},
		{
			"invalid min_price",
			1,
			10,
			"created_at",
			"desc",
			"0",
			"10.00",
			"",
			price.ErrInvalidPriceFormat,
		},
		{
			"invalid max_price",
			1,
			10,
			"created_at",
			"desc",
			"0.00",
			"10",
			"",
			price.ErrInvalidPriceFormat,
		},
		{
			"invalid max_price",
			1,
			10,
			"created_at",
			"desc",
			"0.00",
			"10",
			"",
			price.ErrInvalidPriceFormat,
		},
		{
			"invalid min_price and max_price",
			1,
			10,
			"created_at",
			"desc",
			"0",
			"10",
			"",
			price.ErrInvalidPriceFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _, actualErr := queryTemplates.GenerateAdQuery(
				tt.pageNum,
				tt.resultsPerPage,
				tt.sortField,
				tt.sortOrder,
				tt.minPrice,
				tt.maxPrice,
			)
			assert.Equal(t, tt.expectedQuery, actual)
			assert.ErrorIs(t, actualErr, tt.expectedError, "errors must match")
		})
	}
}
