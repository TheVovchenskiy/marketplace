package listOptions

import (
	"marketplace/pkg/utils"
	"net/url"
)

type AllowedValues struct {
	Required bool
	Values   []string
	Default  string
}

/*
Parse url query for allowed params. If some param is not present in query
the default value for this param is the first one in the slice under the key of the param.

If slice in `allowedParams` is empty then any value is allowed.
*/
func NewListOptions(query url.Values, allowedParams map[string][]string) (map[string]interface{}, error) {
	res := map[string]interface{}{}

	for key, values := range query {
		allowedVals, ok := allowedParams[key]
		if !ok {
			return nil, ErrInvalidQueryParam
		}
		if len(allowedVals) > 0 && !utils.In(values[0], allowedVals) {
			return nil, ErrInvalidQueryParam
		}
		if _, ok := res[key]; ok {
			return nil, ErrInvalidQueryParam
		}
		res[key] = values[0]
	}

	for allowedKey, allowedValues := range allowedParams {
		if _, ok := res[allowedKey]; !ok {
			if len(allowedValues) == 0 {
				return nil, ErrInvalidQueryParam
			}
			res[allowedKey] = allowedValues[0]
		}
	}

	return res, nil
}
