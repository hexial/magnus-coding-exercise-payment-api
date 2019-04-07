package storage

import (
	"fmt"
	"time"
	"unicode"

	"github.com/rs/zerolog/log"
)

//
// GORMZerolog is an adapter for GORM to log to https://github.com/rs/zerolog
type GORMZerolog struct {
}

//
// Print is the actual method that is invoked from GORM when logging
func (l GORMZerolog) Print(v ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Panic: %+v", r)
		}
	}()
	if len(v) > 1 {
		if v[0] == "sql" {
			log.Debug().Str("source", l.getSource(v)).Int("affected", l.getAffected(v)).Strs("values", l.getFormattedValues(v)).Msg(l.getSQL(v))
		} else {
			str := ""
			for _, val := range v {
				str += fmt.Sprintf("[%v]", val)
			}
			log.Debug().Msgf("%s", str)
		}
	}
}

func (l GORMZerolog) isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func (l GORMZerolog) getFormattedValues(values []interface{}) []string {
	rawValues := values[4].([]interface{})
	formattedValues := make([]string, 0, len(rawValues))
	for _, value := range rawValues {
		switch v := value.(type) {
		case time.Time:
			formattedValues = append(formattedValues, fmt.Sprint(v))
		case []byte:
			if str := string(v); l.isPrintable(str) {
				formattedValues = append(formattedValues, fmt.Sprint(str))
			} else {
				formattedValues = append(formattedValues, "<binary>")
			}
		default:
			str := "NULL"
			if v != nil {
				str = fmt.Sprint(v)
			}
			formattedValues = append(formattedValues, str)
		}
	}
	return formattedValues
}

func (l GORMZerolog) getSource(values []interface{}) string {
	return fmt.Sprint(values[1])
}

func (l GORMZerolog) getDuration(values []interface{}) time.Duration {
	return values[2].(time.Duration)
}

func (l GORMZerolog) getSQL(values []interface{}) string {
	return values[3].(string)
}

func (l GORMZerolog) getAffected(values []interface{}) int {
	return int(values[5].(int64))
}
