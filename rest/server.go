package rest

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sskoredin/iiko_report/report"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/pkg/errors"
)

func (r Rest) Run() error {
	r.logger.Debug("Starting web server...")
	if err := r.config.Read(); err != nil {
		return err
	}
	r.logger.Debugf("config:%+v", r.config)
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
	r.logger.Infof("get report request from %s to %s", q.start, q.end)

	valid, err := q.isValid()
	if err != nil {
		fmt.Fprintf(w, "error check period start=%s end=%s. err: %s ", q.start, q.end, err.Error())
		return
	}
	if !valid {
		fmt.Fprintf(w, "invalid period start=%s end=%s", q.start, q.end)
		return
	}
	r.logger.Debug("making report...")
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
	router.HandleFunc("/report", r.getReport).Methods(http.MethodGet)
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/app/public")))

	return &http.Server{
		Addr:    r.config.ListenAddr(),
		Handler: router,
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("/app/public/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
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
