package db

import (
	"fmt"
)

var (
	insert               = `INSERT INTO %s (%s) VALUES %s`
	getRecord            = `SELECT %s FROM %s WHERE %s`
	updateRecord         = `UPDATE %s SET %s WHERE %s`
	getLeftJoinRecord    = `SELECT %s  FROM %s LEFT JOIN %s  ON %s`
	updateOnConflict     = `INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s`
	getWheelOfLifeDetail = `SELECT d.id,d.icon, e.type,e.content FROM (
								SELECT b.id, b.icon from wheel_of_life as b LEFT JOIN 
									(select min(score), wheel_of_life_id from wheel_of_life_user_score where user_id='%s' group by wheel_of_life_id limit 1) 
										as a ON b.id = a.wheel_of_life_id where b.id=a.wheel_of_life_id) as d left join multi_language_content as e on d.id=e.parent_id where language='%s';`
	getInspirationThemeDetail = `SELECT c.inspiration_id, c.theme_type, c.theme_content, d.array_content as motivations  FROM 
									(SELECT a.id as inspiration_id, b.type as theme_type, b.content as theme_content FROM 
										(select i.id, i.theme_id from inspiration as i where wheel_of_life_id='%s') as a 
										LEFT JOIN multi_language_content as b on a.theme_id = b.parent_id where b.language='%s') as c 
									LEFT JOIN multi_language_content as d on c.inspiration_id=d.parent_id where d.language='%s';`
	AddDeviceId = `UPDATE profile SET device_ids = array_append(device_ids, '%s') where user_id='%s'`
	// table names
	USERS = "users"
)

func InsertQuery(tableName string, fields []string, returnParam string, rows int) (insertQuery string) {
	var (
		placeHolderRows, keys string
	)
	l := 0
	for j := 1; j <= rows; j++ {
		var (
			placeholders string
		)
		for i, field := range fields {

			comma := ""
			if i < (len(fields) - 1) {
				comma = ", "
			}
			if j == 1 {
				keys = fmt.Sprintf("%s%s%s", keys, field, comma)
			}
			placeholders = fmt.Sprintf("%s$%d%s", placeholders, i+1+l, comma)
		}
		p := fmt.Sprintf("(%s)", placeholders)
		placeHolderRows = placeHolderRows + p
		if j < rows {
			placeHolderRows = placeHolderRows + ","
		}
		l += 5
	}

	insertQuery = fmt.Sprintf(insert, tableName, keys, placeHolderRows)
	if returnParam != "" {
		insertQuery = insertQuery + " RETURNING " + returnParam
	}

	return insertQuery
}

func GetRecordQuery(tableName string, projections, conditions []string, conditionType string, paramLength int) (selectQuery string) {
	getRecordQuery := ""
	switch conditionType {
	case "=":
		getRecordQuery = fmt.Sprintf(getRecord, ConcateArray(projections), tableName, GetQueryCondition(conditions, 1))
		return getRecordQuery
	case "in":
		condition := fmt.Sprintf("parent_id in (%s)", GetQueryInCondition(paramLength))
		getRecordQuery = fmt.Sprintf(getRecord, ConcateArray(projections), tableName, condition)
		return getRecordQuery

	}
	return getRecordQuery
}

func UpdateRecordQuery(tableName string, updatedFields []string, queryFields []string, returnParam string) (updateQuery string) {

	updateRecordQuery := fmt.Sprintf(updateRecord, tableName, GetQueryCondition(updatedFields, 1), GetQueryCondition(queryFields, len(updatedFields)+1))
	if returnParam != "" {
		updateRecordQuery = updateRecordQuery + " RETURNING " + returnParam
	}
	return updateRecordQuery
}

func GetLeftJoinRecordQuery(projection, tables, whereCond []string, tableRel, conditionType string, paramLength int) (query string) {
	query = fmt.Sprintf(getLeftJoinRecord, ConcateArray(projection), tables[0], tables[1], tableRel)
	if len(whereCond) > 0 {
		switch conditionType {
		case "=":
			query = fmt.Sprintf("%s where %s", query, GetQueryCondition(whereCond, 1))
			return query
		case "in":
			query = fmt.Sprintf("%s where parent_id in (%s)", query, GetQueryInCondition(paramLength))
			return query
		default:
			return query
		}
	}
	return query
}

func GetOnConflictUpdate(tableName string, fields []string, conflictFields []string, conflictUpdateFields []string) (updateQuery string) {
	keys, placeholders := GenerateInsertValue(fields)
	updateQuery = fmt.Sprintf(updateOnConflict, tableName, keys, placeholders, ConcateArray(conflictFields), GetQueryCondition(conflictUpdateFields, len(fields)+1))
	return
}
