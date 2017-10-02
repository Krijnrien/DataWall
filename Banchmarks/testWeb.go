package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/gocql/gocql"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"gopkg.in/mgo.v2"
	//"bytes"
	//"reflect"
)
var Session *gocql.Session
var y int = 40000;

type Device struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	UserType int     `json:"userType"`
	Hash     string  `json:"hash"`
}

/* Token below must be replaced with a valid one */
var personalAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSIsImtpZCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSJ9.eyJpc3MiOiJodHRwczovL2lkZW50aXR5LmZoaWN0Lm5sIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eS5maGljdC5ubC9yZXNvdXJjZXMiLCJleHAiOjE1MDY5NTMxMDIsIm5iZiI6MTUwNjk0NTkwMiwiY2xpZW50X2lkIjoiYXBpLWNsaWVudCIsInVybjpubC5maGljdDp0cnVzdGVkX2NsaWVudCI6InRydWUiLCJzY29wZSI6WyJvcGVuaWQiLCJwcm9maWxlIiwiZW1haWwiLCJmaGljdCIsImZoaWN0X3BlcnNvbmFsIiwiZmhpY3RfbG9jYXRpb24iXSwic3ViIjoiY2ViNTA3ZDMtZGMxOC00NDdiLTkxNTEtZjNiNDZlOGRhMTkzIiwiYXV0aF90aW1lIjoxNTA2OTQ1OTAxLCJpZHAiOiJmaGljdC1zc28iLCJyb2xlIjpbInVzZXIiLCJzdHVkZW50Il0sInVwbiI6IkkzNTg3MTdAZmhpY3QubmwiLCJuYW1lIjoiU2F2b3YsTWFydGluIE0uWi4iLCJlbWFpbCI6Im1hcnRpbi5zYXZvdkBzdHVkZW50LmZvbnR5cy5ubCIsInVybjpubC5maGljdDpzY2hlZHVsZSI6ImNsYXNzfERlbHRhIC8gRWkzUzQgLyBTTTQxIiwiZm9udHlzX3VwbiI6IjM1ODcxN0BzdHVkZW50LmZvbnR5cy5ubCIsImFtciI6WyJleHRlcm5hbCJdfQ.PoNMzUv0GQYnZchjyqMN32Jj-mSj5Xs2KtZXHtPpSc3J0ocP2eoE-iFzYJzZTrpp-ZwXAXAWCRnikXREgpJRotK3Eq32Tgd8hUasYiOOYbdPy-ySYWrhNN0wbvn9BPpOSYzXOkCJ34jkVo33XZsJYHRJKt1MqfIy9j6ZLyZv7R64U6V0JX1m7qWg0aeFUEqkJJPFLHjjgKmWUGwLn2-nX5UTjrDWufY2yZCxsDY92KDYYOkL-hW1JjcanSOx18bWJBHm4xF64ugrZvshdW4J8EpvX-XOV4fxL1bmRaCD0OUUdhqedYr_m3KQPVzg1C1r9aE_SXYUmKcOOP4lrSOwHA"

var url = "https://api.fhict.nl/location/devices"

type TokenSource struct {
	AccessToken string
}

var tokenSource = &TokenSource{
	AccessToken: personalAccessToken,
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}


func getInfo(t time.Time) {

	//Cassandra
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	resp, err := oauthClient.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "people"
	//MongoDB
 	/*MongoDB, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	}
	defer MongoDB.Close()*/

	//MySql
	dbMySQL, err := sql.Open("mysql", "root:@/datawall_db_test")
	defer dbMySQL.Close()
	stmt, err := dbMySQL.Prepare("INSERT INTO test (id,x, y, z) values(?,?,?,?)")

	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	Session, err := cluster.CreateSession()

	body, _ := ioutil.ReadAll(resp.Body)

	jsondata := string(body)

	var devices []Device // stores json in struct

	err = json.Unmarshal([]byte(jsondata), &devices)

	//Cassandra
	fmt.Println("Casandera DB Starting insert")
	fmt.Println("StartTime: %v", time.Now())
	for i := 0; i < len(devices); i++ {
		y=y+1
		if err := Session.Query("INSERT INTO people.people (id,x, y, z) VALUES (?,?,?,? );",y,float32(devices[i].X),float32(devices[i].Y),float32(devices[i].Z)).Exec(); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Endtime: ", time.Now())
	//MySQL
	fmt.Println("MySQL DB Starting insert")
	fmt.Println("StartTime: %v", time.Now())
	for i := 0; i < len(devices); i++ {
			y=y+1
			res,err:= stmt.Exec(y,float32(devices[i].X),float32(devices[i].Y),float32(devices[i].Z))
			if err != nil {
	 			fmt.Println(err)
				fmt.Println(res)
			}
	}
		fmt.Println("Endtime: ", time.Now())
		//MongoDB
		/*fmt.Println("Mongo DB Starting insert")
		fmt.Println("StartTime: %v", time.Now())
		for i := 0; i < len(devices); i++ {
					if err = MongoDB.DB("store").C("Devices").Insert(y,float32(devices[i].X),float32(devices[i].Y),float32(devices[i].Z));err != nil {
						fmt.Println(err)
				}
			}
		fmt.Println("Endtime: ", time.Now())*/
		dbMySQL.Close()
}

func UpdateWebPage(t time.Time) {
	http.HandleFunc("/", serveApi)
	http.ListenAndServe(":7000", nil)
}

func serveApi(w http.ResponseWriter, h *http.Request) {

	resp, err := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	jsondata := string(body)

	var devices []Device // stores json in struct

	err = json.Unmarshal([]byte(jsondata), &devices)
	if err != nil {
		fmt.Println(err)
	}

	structuredjson, err := json.MarshalIndent(devices, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(structuredjson))
}

func main() {
	//http.HandleFunc("/", serveApi)
//	http.ListenAndServe(":7000", nil)
	doEvery(20000*time.Millisecond, getInfo)

}
