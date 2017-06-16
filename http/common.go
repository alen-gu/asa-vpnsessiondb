package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"time"

	"github.com/alen-gu/asa-vpnsessiondb/g"
	"github.com/alen-gu/asa-vpnsessiondb/sqlite"
	"github.com/toolkits/file"
)

type Sresult struct {
	Success bool                `json:"success"`
	Result  []sqlite.VpnSession `json:"result"`
}
type Fresult struct {
	Success bool   `json:"success"`
	Msg     string `json:"err_msg"`
}

func configCommonRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var fresult Fresult
		offset := 0
		limit := 0
		var err error
		var b []byte
		offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
		if (r.URL.Query().Get("offset") != "" && err != nil) || offset < 0 {
			fresult.Msg = "offset error"
			b, _ = json.Marshal(fresult)
			w.Write(b)
			return
		}
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if (r.URL.Query().Get("limit") != "" && err != nil) || limit < 0 {
			fresult.Msg = "limit error"
			b, _ = json.Marshal(fresult)
			w.Write(b)
			return
		}
		//fmt.Println(offset, limit)
		var sresult Sresult
		sresult.Success = true
		sessions, err1 := sqlite.SearchAll_DB()
		if err1 != nil {
			fresult.Msg = "sqlite error"
			b, _ = json.Marshal(fresult)
			w.Write(b)
			return
		}
		if offset < len(sessions) {
			if limit == 0 || offset+limit > len(sessions) {
				sresult.Result = sessions[offset:]
			} else {
				sresult.Result = sessions[offset : offset+limit]
			}
			b, _ = json.Marshal(sresult)
			w.Write(b)
		} else {
			fresult.Msg = "offset range error"
			b, _ = json.Marshal(fresult)
			w.Write(b)
		}

	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.VERSION))

	})

	http.HandleFunc("/workdir", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s\n", file.SelfDir())))
	})

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, g.Config())
	})

	//gu
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var fresult Fresult
		var query string
		//var err error
		var b []byte
		crasUsername := r.URL.Query().Get("crasUsername")
		crasLocalAddress := r.URL.Query().Get("crasLocalAddress")
		crasISPAddress := r.URL.Query().Get("crasISPAddress")
		if crasUsername != "" {
			query = query + " and crasUsername = '" + crasUsername + "'"
		}
		if crasLocalAddress != "" {
			query = query + " and crasLocalAddress = '" + crasLocalAddress + "'"
		}
		if crasISPAddress != "" {
			query = query + " and crasISPAddress = '" + crasISPAddress + "'"
		}
		if query == "" {
			fresult.Msg = "无查询条件"
			b, _ = json.Marshal(fresult)
			w.Write(b)
			return
		}
		var sresult Sresult
		sresult.Success = true
		sessions, err1 := sqlite.Search_DB(query, 4)
		if err1 != nil {
			fresult.Msg = "sqlite error"
			b, _ = json.Marshal(fresult)
			w.Write(b)
			return
		}
		fmt.Println(sessions)
		sresult.Result = sessions
		b, _ = json.Marshal(sresult)
		w.Write(b)

	})
}
