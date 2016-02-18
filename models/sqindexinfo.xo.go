// Package models contains the types for schema ''.
package models

// GENERATED BY XO. DO NOT EDIT.

// SqIndexinfo represents a row from ..
type SqIndexinfo struct {
	SeqNo      int    // seq_no
	Cid        int    // cid
	ColumnName string // column_name
}

// SqIndexinfosByIndex runs a custom query, returning results as SqIndexinfo.
func SqIndexinfosByIndex(db XODB, index string) ([]*SqIndexinfo, error) {
	var err error

	// sql query
	var sqlstr = `PRAGMA index_info(` + index + `)`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*SqIndexinfo{}
	for q.Next() {
		si := SqIndexinfo{}

		// scan
		err = q.Scan(&si.SeqNo, &si.Cid, &si.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &si)
	}

	return res, nil
}