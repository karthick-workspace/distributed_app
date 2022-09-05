package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (f1 fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(f1), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		return 0, err
	}

	defer f.Close()
	return f.Write(data)
}

func Run(destrination string) {
	log = stlog.New(fileLog(destrination), "", stlog.LstdFlags)
}

func write(message string) {
	log.Printf("%v\n", message)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)

		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}
