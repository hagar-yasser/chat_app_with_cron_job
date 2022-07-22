// package main

// // 1
// import (
// //  "fmt"
//  "time"
//  "database/sql"
//  "os"
// //  "strings"

//  _ "github.com/go-co-op/gocron"
//  _ "github.com/go-sql-driver/mysql"
 
// )

// var db *sql.DB

// func dbConnect() error {
// 	//config
// 	cfg := mysql.Config{
// 		User:                 os.Getenv("DBUSER"),
// 		Passwd:               os.Getenv("DBPASS"),
// 		Net:                  "tcp",
// 		Addr:                 os.Getenv("DBHOST") ,
// 		DBName:               os.Getenv("DBNAME"),
// 		AllowNativePasswords: true,
// 	}

// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())

// 	return err
// }

// // 2
// func hello(){
// //  message := fmt.Sprintf("Hi, %v", name)
// //  fmt.Println(message)
// q := "UPDATE applications a "+
//     "INNER JOIN "+
// 		"( "+
// 			"SELECT COUNT(*) AS count , c.applications_id AS applications_id "+
// 			"FROM chats c INNER JOIN "+
// 				"( "+
// 				"SELECT id FROM applications WHERE changed = TRUE "+
// 				") q1 "+
// 				"ON c.applications_id = q1.id "+
// 				"GROUP BY c.applications_id "+
// 		") q2 "+
// 	"ON a.id = q2.applications_id "+
// 	"SET a.chats_count = q2.count"
// _, err := db.Exec(q)


// q2 := "UPDATE chats c "+
//     "INNER JOIN "+
// 		"( "+
// 			"SELECT COUNT(*) AS count , m.chats_id AS chats_id "+
// 			"FROM messages m INNER JOIN "+
// 				"( "+
// 				"SELECT id FROM chats WHERE changed = TRUE "+
// 				") q1 "+
// 				"ON m.chats_id = q1.id "+
// 				"GROUP BY m.chats_id "+
// 		") q2 "+
// 	"ON c.id = q2.chats_id "+
// 	"SET a.messages_count = q2.count"
// _, err2 := db.Exec(q2)




// }

// func runCronJobs() {
//  // 3
//  s := gocron.NewScheduler(time.UTC)

//  // 4
//  s.Every(1).Minutes().Do(func() {
//   hello()
//  })

//  // 5
//  s.StartBlocking()
// }

// // 6
// func main() {
//  runCronJobs()
// }

package main

// 1
import (
 "fmt"
 "time"
 "database/sql"
 "os"

 "github.com/go-co-op/gocron"
 "github.com/go-sql-driver/mysql"
)
//  var db *sql.DB

 func dbConnect() *sql.DB {
	//config
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}
	var db *sql.DB
	var err error
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
        fmt.Println("problem in connection")
    }

	return db
}

// 2
func hello(name string, db *sql.DB) {
 message := fmt.Sprintf("Hi, %v", name)
 fmt.Println(message)
 q := "UPDATE applications a "+
    "INNER JOIN "+
		"( "+
			"SELECT COUNT(*) AS count , c.applications_id AS applications_id "+
			"FROM chats c INNER JOIN "+
				"( "+
				"SELECT id FROM applications WHERE has_different_count = TRUE "+
				") q1 "+
				"ON c.applications_id = q1.id "+
				"GROUP BY c.applications_id "+
		") q2 "+
	"ON a.id = q2.applications_id "+
	"SET a.chats_count = q2.count"
_, err := db.Exec(q)


q2 := "UPDATE chats c "+
    "INNER JOIN "+
		"( "+
			"SELECT COUNT(*) AS count , m.chats_id AS chats_id "+
			"FROM messages m INNER JOIN "+
				"( "+
				"SELECT id FROM chats WHERE has_different_count = TRUE "+
				") q1 "+
				"ON m.chats_id = q1.id "+
				"GROUP BY m.chats_id "+
		") q2 "+
	"ON c.id = q2.chats_id "+
	"SET c.messages_count = q2.count"
_, err2 := db.Exec(q2)
fmt.Println(err)
fmt.Println(err2)
}

func runCronJobs(db *sql.DB) {
 // 3
 s := gocron.NewScheduler(time.UTC)
 

 // 4
 s.Every(20).Minutes().Do(func() {
  hello("Updated chats_count and messages_count",db)
 })

 // 5
 s.StartBlocking()
}

// 6
func main() {
 var db *sql.DB
 db = dbConnect()
 runCronJobs(db)
}
