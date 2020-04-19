package logs

import (
	"strconv"
	"strings"
)

var defaultLimit = 10

type SearchCriteria struct {
	conditions []string
	params     []interface{}
	limit      int
}

func (sc *SearchCriteria) AddCriteria(condition string, param interface{}) {
	if param == nil {
		return
	}

	sc.conditions = append(sc.conditions, condition)
	sc.params = append(sc.params, param)
}

// BuildQuery builds and returns final query
func (sc *SearchCriteria) BuildQuery(tableName string) string {
	where := ""
	if len(sc.conditions) > 0 {
		where = " WHERE " + strings.Join(sc.conditions, " AND ")
	}

	return "SELECT id, source, category, level, message, trace, payload, created_at FROM " + tableName + where + " ORDER BY created_at DESC LIMIT ?"
}

func (sc *SearchCriteria) GetConditionValues() []interface{} {
	return sc.params
}

func (sc *SearchCriteria) PopulateWithMap(params map[string][]string) {
	sc.limit = defaultLimit

	for k, v := range params {
		if k == "limit" {
			sc.limit, _ = strconv.Atoi(v[0])
			continue
		}

		sc.AddCriteria(k+"=?", v[0])
	}

	sc.params = append(sc.params, sc.limit)
}
