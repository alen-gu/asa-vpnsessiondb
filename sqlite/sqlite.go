package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type VpnSession struct {
	CrasIndex                 int64  `json:"crasIndex"`
	CrasUsername              string `json:"crasUsername"`
	CrasGroup                 string `json:"crasGroup"`
	CrasSessionDuration       int64  `json:"crasSessionDuration"`
	CrasLocalAddress          string `json:"crasLocalAddress"`
	CrasISPAddress            string `json:"crasISPAddress"`
	CrasClientVendorString    string `json:"crasClientVendorString"`
	CrasClientVersionString   string `json:"crasClientVersionString"`
	CrasClientOSVendorString  string `json:"crasClientOSVendorString"`
	CrasClientOSVersionString string `json:"crasClientOSVersionString"`
	CrasSessionInPkts         int64  `json:"crasSessionOutPkts"`
	CrasSessionOutPkts        int64  `json:"crasSessionOutPkts"`
	CrasSessionInDropPkts     int64  `json:"crasSessionInDropPkts"`
	CrasSessionOutDropPkts    int64  `json:"crasSessionOutDropPkts"`
	CrasSessionInOctets       int64  `json:"crasSessionInOctets"`
	CrasSessionOutOctets      int64  `json:"crasSessionOutOctets"`
	CrasSessionState          int64  `json:"crasSessionState"`
	UpdateTime                int64  `json:"updateTime"`
}

func Init_DB() error {
	err := os.Remove("./session.db")
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStmt := `
	create table sessions (
	id integer not null primary key, 
	crasIndex integer,
	crasUsername text,
	crasGroup text,
	crasSessionDuration integer,
	crasLocalAddress text,
	crasISPAddress text,
	crasClientVendorString text,
	crasClientVersionString text,
	crasClientOSVendorString text,
	crasClientOSVersionString text,
	crasSessionInPkts integer,
	crasSessionOutPkts integer,
	crasSessionInDropPkts integer,
	crasSessionOutDropPkts integer,
	crasSessionInOctets integer,
	crasSessionOutOctets integer,
	crasSessionState integer,
	updateTime integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return err
}

func InsertOne_DB(session VpnSession) error {
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	insert_sql := `insert into sessions(
	crasIndex,
	crasUsername,
	crasGroup,
	crasSessionDuration,
	crasLocalAddress,
	crasISPAddress,
	crasClientVendorString,
	crasClientVersionString,
	crasClientOSVendorString,
	crasClientOSVersionString,
	crasSessionInPkts,
	crasSessionOutPkts,
	crasSessionInDropPkts,
	crasSessionOutDropPkts,
	crasSessionInOctets,
	crasSessionOutOctets,
	crasSessionState,
	updateTime) 
	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	stmt, err := tx.Prepare(insert_sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	now := time.Now().Unix()
	_, err = stmt.Exec(session.CrasIndex,
		session.CrasUsername,
		session.CrasGroup,
		session.CrasSessionDuration,
		session.CrasLocalAddress,
		session.CrasISPAddress,
		session.CrasClientVendorString,
		session.CrasClientVersionString,
		session.CrasClientOSVendorString,
		session.CrasClientOSVersionString,
		session.CrasSessionInPkts,
		session.CrasSessionOutPkts,
		session.CrasSessionInDropPkts,
		session.CrasSessionOutDropPkts,
		session.CrasSessionInOctets,
		session.CrasSessionOutOctets,
		session.CrasSessionState,
		now,
	)
	if err != nil {
		return err
	}

	tx.Commit()
	return err
}

func Insert_DB(VpnSessionDB map[int]VpnSession) error {
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	insert_sql := `insert into sessions(
	crasIndex,
	crasUsername,
	crasGroup,
	crasSessionDuration,
	crasLocalAddress,
	crasISPAddress,
	crasClientVendorString,
	crasClientVersionString,
	crasClientOSVendorString,
	crasClientOSVersionString,
	crasSessionInPkts,
	crasSessionOutPkts,
	crasSessionInDropPkts,
	crasSessionOutDropPkts,
	crasSessionInOctets,
	crasSessionOutOctets,
	crasSessionState,
	updateTime) 
	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	stmt, err := tx.Prepare(insert_sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	now := time.Now().Unix()
	for _, session := range VpnSessionDB {
		_, err = stmt.Exec(session.CrasIndex,
			session.CrasUsername,
			session.CrasGroup,
			session.CrasSessionDuration,
			session.CrasLocalAddress,
			session.CrasISPAddress,
			session.CrasClientVendorString,
			session.CrasClientVersionString,
			session.CrasClientOSVendorString,
			session.CrasClientOSVersionString,
			session.CrasSessionInPkts,
			session.CrasSessionOutPkts,
			session.CrasSessionInDropPkts,
			session.CrasSessionOutDropPkts,
			session.CrasSessionInOctets,
			session.CrasSessionOutOctets,
			session.CrasSessionState,
			now,
		)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return err
}

func Search_DB(filter string, mode int) ([]VpnSession, error) {
	var (
		id      int
		session VpnSession
	)
	sessions := []VpnSession{}
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	switch mode {
	case 1:
		rows, err := db.Query("select * from sessions where crasUsername = ?", filter)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id,
				&session.CrasIndex,
				&session.CrasUsername,
				&session.CrasGroup,
				&session.CrasSessionDuration,
				&session.CrasLocalAddress,
				&session.CrasISPAddress,
				&session.CrasClientVendorString,
				&session.CrasClientVersionString,
				&session.CrasClientOSVendorString,
				&session.CrasClientOSVersionString,
				&session.CrasSessionInPkts,
				&session.CrasSessionOutPkts,
				&session.CrasSessionInDropPkts,
				&session.CrasSessionOutDropPkts,
				&session.CrasSessionInOctets,
				&session.CrasSessionOutOctets,
				&session.CrasSessionState,
				&session.UpdateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 2:
		rows, err := db.Query("select * from sessions where crasLocalAddress = ?", filter)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id,
				&session.CrasIndex,
				&session.CrasUsername,
				&session.CrasGroup,
				&session.CrasSessionDuration,
				&session.CrasLocalAddress,
				&session.CrasISPAddress,
				&session.CrasClientVendorString,
				&session.CrasClientVersionString,
				&session.CrasClientOSVendorString,
				&session.CrasClientOSVersionString,
				&session.CrasSessionInPkts,
				&session.CrasSessionOutPkts,
				&session.CrasSessionInDropPkts,
				&session.CrasSessionOutDropPkts,
				&session.CrasSessionInOctets,
				&session.CrasSessionOutOctets,
				&session.CrasSessionState,
				&session.UpdateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 3:
		rows, err := db.Query("select * from sessions where crasISPAddress = ?", filter)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id,
				&session.CrasIndex,
				&session.CrasUsername,
				&session.CrasGroup,
				&session.CrasSessionDuration,
				&session.CrasLocalAddress,
				&session.CrasISPAddress,
				&session.CrasClientVendorString,
				&session.CrasClientVersionString,
				&session.CrasClientOSVendorString,
				&session.CrasClientOSVersionString,
				&session.CrasSessionInPkts,
				&session.CrasSessionOutPkts,
				&session.CrasSessionInDropPkts,
				&session.CrasSessionOutDropPkts,
				&session.CrasSessionInOctets,
				&session.CrasSessionOutOctets,
				&session.CrasSessionState,
				&session.UpdateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 0:
		rows, err := db.Query("select * from sessions where crasIndex = ?", filter)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id,
				&session.CrasIndex,
				&session.CrasUsername,
				&session.CrasGroup,
				&session.CrasSessionDuration,
				&session.CrasLocalAddress,
				&session.CrasISPAddress,
				&session.CrasClientVendorString,
				&session.CrasClientVersionString,
				&session.CrasClientOSVendorString,
				&session.CrasClientOSVersionString,
				&session.CrasSessionInPkts,
				&session.CrasSessionOutPkts,
				&session.CrasSessionInDropPkts,
				&session.CrasSessionOutDropPkts,
				&session.CrasSessionInOctets,
				&session.CrasSessionOutOctets,
				&session.CrasSessionState,
				&session.UpdateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 4:
		query := "select * from sessions where 1=1 "
		query = query + filter
		log.Println(query)
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id,
				&session.CrasIndex,
				&session.CrasUsername,
				&session.CrasGroup,
				&session.CrasSessionDuration,
				&session.CrasLocalAddress,
				&session.CrasISPAddress,
				&session.CrasClientVendorString,
				&session.CrasClientVersionString,
				&session.CrasClientOSVendorString,
				&session.CrasClientOSVersionString,
				&session.CrasSessionInPkts,
				&session.CrasSessionOutPkts,
				&session.CrasSessionInDropPkts,
				&session.CrasSessionOutDropPkts,
				&session.CrasSessionInOctets,
				&session.CrasSessionOutOctets,
				&session.CrasSessionState,
				&session.UpdateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	default:
		return nil, errors.New("searchmode error: must in 0-2")
	}
	return sessions, err
}

func SearchByIndex_DB(filter int64) (bool, error) {
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return false, err
	}
	defer db.Close()
	rows, err := db.Query("select * from sessions where crasIndex = ?", filter)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	result := rows.Next()
	return result, err
}

func SearchAll_DB() ([]VpnSession, error) {
	var (
		id      int
		session VpnSession
	)
	sessions := []VpnSession{}
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,
			&session.CrasIndex,
			&session.CrasUsername,
			&session.CrasGroup,
			&session.CrasSessionDuration,
			&session.CrasLocalAddress,
			&session.CrasISPAddress,
			&session.CrasClientVendorString,
			&session.CrasClientVersionString,
			&session.CrasClientOSVendorString,
			&session.CrasClientOSVersionString,
			&session.CrasSessionInPkts,
			&session.CrasSessionOutPkts,
			&session.CrasSessionInDropPkts,
			&session.CrasSessionOutDropPkts,
			&session.CrasSessionInOctets,
			&session.CrasSessionOutOctets,
			&session.CrasSessionState,
			&session.UpdateTime,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		sessions = append(sessions, session)
	}

	return sessions, err
}

func DeleteOld_DB(interval int64) error {
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer db.Close()
	expiredTime := time.Now().Unix() - 2*interval
	stmt, err := db.Prepare("delete from sessions where updateTime < ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(expiredTime)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(affect)
	return err

}

func UpdateTime_DB(filter int64) error {
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer db.Close()
	now := time.Now().Unix()
	stmt, err := db.Prepare("update sessions set updateTime = ? where crasIndex = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(now, filter)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(affect)
	return err
}
