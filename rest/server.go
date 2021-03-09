package rest

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/iiko_report/config"
	"github.com/sskoredin/iiko_report/logger"
	"github.com/sskoredin/iiko_report/report"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
)

type Rest struct {
	Configfile string
	logger     logger.Logger
	config     config.Rest
}

func New() Rest {
	return Rest{
		logger:     logger.New("api", logrus.DebugLevel),
		Configfile: config.FileName(),
	}
}
func (r *Rest) readConfig() error {
	var c config.Config
	if _, err := os.Stat(r.Configfile); err != nil {
		return err
	}
	v, err := ioutil.ReadFile(r.Configfile)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(v, &c); err != nil {
		return err
	}
	r.config = c.Rest
	return nil
}

func (r Rest) Run() error {
	r.logger.Debug("Starting web server...")
	if err := r.readConfig(); err != nil {
		return err
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := r.newServer()
	go r.graceFullShutdown(server, quit, done)

	r.logger.Info("Server is ready to handle requests at ", r.config.ListenAddr())

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrapf(err, "Could not listen on %s", r.config.ListenAddr())
	}

	<-done
	r.logger.Info("Server stopped")
	os.Exit(0)

	return nil
}

type query struct {
	start, end string
}

func (q query) isValid() (bool, error) {
	if len(q.start) == 0 && len(q.end) == 0 {
		return true, nil
	}
	if len(q.start) == 0 || len(q.end) == 0 {
		return false, nil
	}

	st, err := time.Parse("02.01.2006", q.start)
	if err != nil {
		return false, err
	}
	end, err := time.Parse("02.01.2006", q.end)
	if err != nil {
		return false, err
	}

	if v := end.Sub(st); end.Before(st) || v > (31*24*time.Hour) {
		return false, nil
	}
	return true, nil
}

func (r Rest) getReport(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	var q query
	q.start = values.Get("start")
	q.end = values.Get("end")
	valid, err := q.isValid()
	if err != nil {
		fmt.Fprintf(w, "error check period start=%s end=%s. err: %s ", q.start, q.end, err.Error())
	}
	if !valid {
		fmt.Fprintf(w, "invalid period start=%s end=%s", q.start, q.end)
		return
	}

	if err := report.MakeReportWithAttempts(q.start, q.end, 0); err != nil {
		response := fmt.Sprintf("Failed to make report, err:%s", err.Error())
		fmt.Fprint(w, response)
		return
	}

	response := fmt.Sprintf("Report from %s to %s was sended", q.start, q.end)
	fmt.Fprint(w, response)

}
func (r Rest) newServer() *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/", r.getReport).Methods(http.MethodGet)

	return &http.Server{
		Addr:    r.config.ListenAddr(),
		Handler: router,
	}
}

func (r Rest) graceFullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	r.logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		r.logger.Fatalf("Could not gracefully shutdown the server: %v", err)
	}

	close(done)
}
