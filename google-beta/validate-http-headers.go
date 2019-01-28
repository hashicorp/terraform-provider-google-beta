package google

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateHttpHeaders() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		headers := i.(map[string]interface{})
		if _, ok := headers["Content-Length"]; ok {
			es = append(es, fmt.Errorf("Cannot set the Content-Length header on %s", k))
			return
		}
		for key := range headers {
			match, _ := regexp.MatchString("(X-Google-|X-AppEngine-).*", key)
			if match {
				es = append(es, fmt.Errorf("Cannot set the %s header on %s", key, k))
				return
			}
		}

		return
	}
}
