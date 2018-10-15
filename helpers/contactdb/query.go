package contactdb

import (
	"strconv"
)

// QueryBuilder will build a HTTP path with additonal queries.
// Example: "/contactdb/lists?page=1&page_size=100&list_id=20"
func QueryBuilder(path string, queries map[string]string) string {
	queryStr := "?"
	for key, val := range queries {
		_, err := strconv.Atoi(val)
		if err != nil {
			queryStr += key + "=\"" + val + "\"&"
		} else {
			queryStr += key + "=" + val + "&"
		}

	}
	queryStr = queryStr[:len(queryStr)-1]
	// Will always remove the last character.
	// This is safe even if the map is empty, due to the string
	// containing a character already. Therefore there is no check.

	return path + queryStr
}
