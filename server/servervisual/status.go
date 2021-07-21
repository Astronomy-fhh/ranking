package servervisual

import (
	"net/http"
	"ranking/db"
	config2 "ranking/server/config"
	"strconv"
)

var status *Status
type Status struct {}


func InitStatusServer()  {
	status = &Status{}
	http.HandleFunc("/printDbVals", status.printDbVals)
	go http.ListenAndServe(config2.SConfig.StatusHttpAddr, nil)
}

func (s Status) printDbVals(w http.ResponseWriter,r *http.Request)  {
	all := db.Db.AllObjs()
	str := "printDbVals:" + strconv.FormatUint(uint64(len(all)), 10) + "\n"
	for key, vals := range all {
		str += key + ":"+ strconv.FormatUint(uint64(len(vals)), 10) + "\n"
		for _, val := range vals {
			str += "\t" + val.Member + "\t" + strconv.FormatInt(val.Score, 10) + "\n"
		}
		str += "\n"
	}
	_, _ = w.Write([]byte(str))
}