package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type vpnsession struct {
	crasUsername              string
	crasGroup                 string
	crasSessionDuration       int64
	crasLocalAddress          string
	crasISPAddress            string
	crasClientVendorString    string
	crasClientVersionString   string
	crasClientOSVendorString  string
	crasClientOSVersionString string
	crasSessionInPkts         int64
	crasSessionOutPkts        int64
	crasSessionInDropPkts     int64
	crasSessionOutDropPkts    int64
	crasSessionInOctets       int64
	crasSessionOutOctets      int64
	crasSessionState          int64
	updateTime                int64
}

func init_database() error {
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

func insert_database(VpnSessionDB map[int]vpnsession) error {
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
	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	stmt, err := tx.Prepare(insert_sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	now := time.Now().Unix()
	for _, session := range VpnSessionDB {
		_, err = stmt.Exec(session.crasUsername,
			session.crasGroup,
			session.crasSessionDuration,
			session.crasLocalAddress,
			session.crasISPAddress,
			session.crasClientVendorString,
			session.crasClientVersionString,
			session.crasClientOSVendorString,
			session.crasClientOSVersionString,
			session.crasSessionInPkts,
			session.crasSessionOutPkts,
			session.crasSessionInDropPkts,
			session.crasSessionOutDropPkts,
			session.crasSessionInOctets,
			session.crasSessionOutOctets,
			session.crasSessionState,
			now,
		)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return err
}

func search_database(filter string, mode int) ([]vpnsession, error) {
	var (
		id      int
		session vpnsession
	)
	sessions := []vpnsession{}
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	switch mode {
	case 0:
		rows, err := db.Query("select * from sessions where crasUsername = ?", filter)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(&id,
				&session.crasUsername,
				&session.crasGroup,
				&session.crasSessionDuration,
				&session.crasLocalAddress,
				&session.crasISPAddress,
				&session.crasClientVendorString,
				&session.crasClientVersionString,
				&session.crasClientOSVendorString,
				&session.crasClientOSVersionString,
				&session.crasSessionInPkts,
				&session.crasSessionOutPkts,
				&session.crasSessionInDropPkts,
				&session.crasSessionOutDropPkts,
				&session.crasSessionInOctets,
				&session.crasSessionOutOctets,
				&session.crasSessionState,
				&session.updateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 1:
		rows, err := db.Query("select * from sessions where crasLocalAddress = ?", filter)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(&id,
				&session.crasUsername,
				&session.crasGroup,
				&session.crasSessionDuration,
				&session.crasLocalAddress,
				&session.crasISPAddress,
				&session.crasClientVendorString,
				&session.crasClientVersionString,
				&session.crasClientOSVendorString,
				&session.crasClientOSVersionString,
				&session.crasSessionInPkts,
				&session.crasSessionOutPkts,
				&session.crasSessionInDropPkts,
				&session.crasSessionOutDropPkts,
				&session.crasSessionInOctets,
				&session.crasSessionOutOctets,
				&session.crasSessionState,
				&session.updateTime,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			sessions = append(sessions, session)
		}
	case 2:
		rows, err := db.Query("select * from sessions where crasISPAddress = ?", filter)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(&id,
				&session.crasUsername,
				&session.crasGroup,
				&session.crasSessionDuration,
				&session.crasLocalAddress,
				&session.crasISPAddress,
				&session.crasClientVendorString,
				&session.crasClientVersionString,
				&session.crasClientOSVendorString,
				&session.crasClientOSVersionString,
				&session.crasSessionInPkts,
				&session.crasSessionOutPkts,
				&session.crasSessionInDropPkts,
				&session.crasSessionOutDropPkts,
				&session.crasSessionInOctets,
				&session.crasSessionOutOctets,
				&session.crasSessionState,
				&session.updateTime,
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

func searchall_database() ([]vpnsession, error) {
	var (
		id      int
		session vpnsession
	)
	sessions := []vpnsession{}
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from sessions")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&id,
			&session.crasUsername,
			&session.crasGroup,
			&session.crasSessionDuration,
			&session.crasLocalAddress,
			&session.crasISPAddress,
			&session.crasClientVendorString,
			&session.crasClientVersionString,
			&session.crasClientOSVendorString,
			&session.crasClientOSVersionString,
			&session.crasSessionInPkts,
			&session.crasSessionOutPkts,
			&session.crasSessionInDropPkts,
			&session.crasSessionOutDropPkts,
			&session.crasSessionInOctets,
			&session.crasSessionOutOctets,
			&session.crasSessionState,
			&session.updateTime,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		sessions = append(sessions, session)
	}

	return sessions, err
}

func deleteold_database(interval int64) error {
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
