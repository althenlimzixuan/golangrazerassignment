package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"hello/ent"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/integrations/ocsql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron"
)

var emailCache *cache.Cache

func main() {

	arg := os.Args[1:]

	for _, s := range arg {
		fmt.Println(s)
	}

	InitSampleCompromisedEmail()

	TryCache()

	// timer = time.NewTicker(1 * time.Second)
	// go func() {
	// 	<-timer.C
	// 	timer = time.NewTicker(1 * time.Second)
	// 	go func() {
	// 		<-timer.C
	// 		TryCache()
	// 	}()
	// }()

	c := cron.New()
	c.AddFunc("*/60 * * * *", func() {
		TryCache()
	})
	c.Start()

	r := newRouter()
	// We can then pass our router (after declaring all our routes) to this method
	// (where previously, we were leaving the second argument as nil)
	http.ListenAndServe(":5000", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/verifyEmail", verifyHandler).Methods("POST")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Println("Hello World!")
	fmt.Fprintf(w, "Hello World!")
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))

	email := r.Form.Get("email")
	fmt.Println("Received ", email)

	// result := IsEmailCompromised("althen:!Althen900601@tcp(192.168.1.115:3306)/test", email)
	result := IsEmailCompromisedInCache(email)

	fmt.Println("Target Email", email, ". Comprimised: ", result)

	responseResult := "{\"result\":"
	if result {
		responseResult += "true}"
	} else {
		responseResult += "false}"
	}
	reponseResultInByte := []byte(responseResult)

	w.Write(reponseResultInByte)
}

func InitSampleCompromisedEmail() {

	sampleList := []string{"althenlim601@gmail.com", "zixuan_lim@yahoo.com"}

	// db, err := ent.Open("mysql", "althen:!Althen900601@tcp(192.168.1.115:3306)/test")
	db, err := ent.Open("mysql", "root:!Razer1234.@tcp(mysql:3306)/test")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	defer db.Close()

	ctx := context.Background()

	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	for _, s := range sampleList {
		record := IsEmailExist(ctx, db, s)
		if !record {
			AddCompromisedEmail(ctx, db, s)
		}
	}

}

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

// Open new connection and start stats recorder.
func Open(dsn string) *ent.Client {
	db := sql.OpenDB(connector{dsn})
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv))
}

func TryCache() {

	// Create a cache with a default expiration time of 1 minutes, and which
	// purges expired items every 2 minutes
	if emailCache != nil {
		emailCache.Flush()
	}

	emailCache = cache.New(1*time.Minute, 2*time.Minute)

	db, err := ent.Open("mysql", "root:!Razer1234.@tcp(mysql:3306)/test")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	defer db.Close()

	ctx := context.Background()

	emailList := ReadAllEmail(ctx, db)

	for _, s := range emailList {
		emailCache.Set(s.Email, s.Email, cache.DefaultExpiration)
		if test, found := emailCache.Get(s.Email); found {
			fmt.Println("Cached ", s.Email, test)
		}
	}

}

func IsEmailCompromisedInCache(email string) bool {
	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:

	if _, found := emailCache.Get(email); found {
		return true
	}

	return false
}
