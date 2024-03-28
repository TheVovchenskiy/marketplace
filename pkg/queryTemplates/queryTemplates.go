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
					a.id,
					a.author_id,
					a."name",
					a.description,
					a.cents_price,
					a.picture_url,
					a.created_at
				FROM
					public.ad a`

	var whereClause string
	if minPrice == "" {
		minPrice = "0.00"
	}
	if maxPrice == "" {
		minPriceValue, err := price.ToCents(minPrice)
		if err != nil {
			return "", nil, err
		}
		whereClause = fmt.Sprintf("WHERE a.cents_price >= $%d", len(args)+1)
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
		whereClause = fmt.Sprintf("WHERE a.cents_price >= $%d AND a.cents_price <= $%d", len(args)+1, len(args)+2)
		args = append(args, minPriceValue, maxPriceValue)
	}

	orderClause := fmt.Sprintf("ORDER BY $%d %s", len(args)+1, strings.ToUpper(sortOrder))
	args = append(args, sortField)

	limitClause := fmt.Sprintf("LIMIT $%d", len(args)+1)
	args = append(args, resultsPerPage)

	offsetCause := fmt.Sprintf("OFFSET $%d", len(args)+1)
	args = append(args, (pageNum-1)*resultsPerPage)

	return strings.Join(
			[]string{
				baseQuery,
				whereClause,
				orderClause,
				limitClause,
				offsetCause,
			},
			"\n",
		),
		args,
		nil
}
