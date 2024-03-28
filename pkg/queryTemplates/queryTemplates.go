package queryTemplates

import (
	"fmt"
	"marketplace/pkg/price"
	"strings"
)

func GenerateAdQuery(
	pageNum int,
	resultsPerPage int,
	sortField string,
	sortOrder string,
	minPrice string,
	maxPrice string,
) (string, []interface{}, error) {
	args := []interface{}{}
	baseQuery := `SELECT
		id,
		author_id,
		"name",
		description,
		cents_price,
		picture_url,
		created_at
	FROM
		public.ad`

	var whereClause string
	if minPrice == "" {
		minPrice = "0.00"
	}
	if maxPrice == "" {
		minPriceValue, err := price.ToCents(minPrice)
		if err != nil {
			return "", nil, err
		}
		whereClause = fmt.Sprintf("WHERE cents_price >= $%d", len(args)+1)
		args = append(args, minPriceValue)
	} else {
		minPriceValue, err := price.ToCents(minPrice)
		if err != nil {
			return "", nil, err
		}
		maxPriceValue, err := price.ToCents(maxPrice)
		if err != nil {
			return "", nil, err
		}
		whereClause = fmt.Sprintf("WHERE cents_price >= $%d AND cents_price <= $%d", len(args)+1, len(args)+2)
		args = append(args, minPriceValue, maxPriceValue)
	}

	orderClause := fmt.Sprintf("ORDER BY %s %s", sortField, strings.ToUpper(sortOrder))

	limitClause := fmt.Sprintf("LIMIT $%d", len(args)+1)
	args = append(args, resultsPerPage)

	offsetCause := fmt.Sprintf("OFFSET $%d", len(args)+1)
	args = append(args, (pageNum-1)*resultsPerPage)

	res := strings.Join(
		[]string{
			baseQuery,
			whereClause,
			orderClause,
			limitClause,
			offsetCause,
		},
		"\n\t",
	)

	// fmt.Println(res)

	return res, args, nil
}
