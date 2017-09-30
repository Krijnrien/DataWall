package cassandra

import (
	"fmt"
	"time"
	"log"
)


func RunTests() {
	_session = session()

	routineStressTest()
}


/**
 *  NOT WRITING COMMENTS HERE AS THIS CODE WILL BE REMOVED IN FUTURE
 */

const (
	gophers = 20
	entries = 1000000
)

func concurrentStressTest(){
	var entries int = 100000

	fmt.Println("StartTime: %v", time.Now())
	for i := 0; i < entries; i++ {
		//fmt.Println("inserting")
		if err := _session.Query("INSERT INTO locations (loc_x, loc_y, loc_z, user_hash, createdAt) VALUES (?, 33421, 33, 'oh cmon', ?);", float32(i), time.Now()).Exec(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("StopTime: %v\n", time.Now())
}

func routineStressTest(){
	fmt.Println("StartTime: ", time.Now())

	for i := 0; i < gophers; i++ {
		go gopher(i)
	}
	fmt.Println("Endtime: ", time.Now())
}

func gopher(gopher_id int) {
	fmt.Printf("Gopher Id: %v \t|| StartTime: %v\n",gopher_id, time.Now())

	for i := 0; i < entries; i++ {
		if err := _session.Query("INSERT INTO locations (loc_x, loc_y, loc_z, user_hash, createdAt) VALUES (?, 33421, 33, 'oh cmon', ?);", float32(i), time.Now()).Exec(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Gopher Id: %v \t|| StopTime: %v\n",gopher_id, time.Now())
}

func printAllTest(){
	var user_hash float32

	iter := _session.Query(`SELECT loc_x FROM locations`).Iter()
	for iter.Scan(&user_hash) {
		fmt.Println("Result: ", user_hash)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}