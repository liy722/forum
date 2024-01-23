package main

import (
	"database/sql"
	"encoding/json"
	_ "fmt"
	"regexp"
	"strings"

	_ "github.com/astaxie/session"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"log"
	"net/http"
	_ "regexp"
)

// Store the username of the currently logged-in user
var currentusername string

// Whether the reply list has a delete button, 1 for yes, 0 for no
var mydelete string = ""

// The ID of the currently selected main post
var currenttid string = ""

// Main post structure
type Thread struct {
	tid      string `json:""` // thread ID
	username string `json:""` // Username of the thread author
	topic    string `json:""` // Thread content
}

// Structure for reply information
type Revert struct {
	revert    string `json:""` // Reply content
	rusername string `json:""` // Username of the person who replied
	stid      string `json:""` // ID of the reply post
	isdelete  string `json:""` // Whether it has a delete button
}

// User structure
type User struct {
	username string `json:"username"` //username
	email    string `json:"email"`    //email
}

// Login response information, where response information refers to messages conveyed from the server to the client
// Regarding response information for login business
type LoginResponse struct {
	Success bool   `json:"success"` // Define a boolean variable named 'Success'
	Message string `json:"message"` // Define a string variable named 'Message'
}

func main() {
	cross() // Call the 'cross' function
}

func cross() {
	// 创建路由和处理器
	router := http.NewServeMux()
	router.HandleFunc("/login", handleLogin) // Frontend path to access the backend, along with the corresponding functions for the respective paths
	router.HandleFunc("/signup", handleSign)
	router.HandleFunc("/blog", handleBlog)
	router.HandleFunc("/createtopic", handleCreatetopic)
	router.HandleFunc("/deletetopic", handleDeleteTopic)
	router.HandleFunc("/subthread", handleSubThread)
	router.HandleFunc("/deletesubthread", handleDeleteSubThread)
	router.HandleFunc("/reply", handleReply)
	// Configure CORS to allow specific origins and headers
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Set the frontend's access origin

	// This is typically used for CORS (Cross-Origin Resource Sharing) settings to restrict which HTTP header fields can be sent from different origins to the server. This enhances security by preventing malicious or unauthorized requests.
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) // Set some allowed values for the headers of the HTTP request
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	// Apply CORS middleware and other middleware to the router
	handler := handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)

	http.HandleFunc("/login", handleLogin) // Frontend path to access the backend, along with the corresponding functions for the respective paths
	http.HandleFunc("/signup", handleSign)
	http.HandleFunc("/blog?topic=", handleBlog)
	http.HandleFunc("/createtopic", handleCreatetopic)
	http.HandleFunc("/subthread", handleSubThread)
	http.HandleFunc("/deletesubthread", handleDeleteSubThread)
	http.HandleFunc("/reply", handleReply)
	// Start the server

	err := http.ListenAndServe(":8080", handler) //// Backend port configuration

	if err != nil {
		panic(err)
	}
}

// Handle login business
func handleLogin(w http.ResponseWriter, r *http.Request) { // Response writer, * indicates a pointer, request

	str := r.RequestURI                  //  /login?username=liy // Frontend access path
	index := strings.Index(str, "=") + 1 // Index position of the character after the equal sign, the result is 16, which is the index position of 'l'
	currentusername = str[index:]        // Including all characters after the index value, "liy"
	db, err1 := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")

	if err1 != nil {
		log.Fatal(err1)
	}

	success := true // Success status
	message := "Login success!"
	response := LoginResponse{Success: success, Message: message}
	//rows, err := db.Query("SELECT username FROM users WHERE username = substring(?,1+locate('=',?))",
	//r.RequestURI, r.RequestURI)
	rows, err := db.Query("SELECT username FROM users WHERE username = '" + currentusername + "'")
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Close the query result set at the end of the function

	// Check if there are results
	if rows.Next() == false {

		// If there are no results, update the response
		response = LoginResponse{Message: "login failed!"}
	}

	// Close the database connection at the end of the function
	defer db.Close()

	// Convert the response structure to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func handleSign(w http.ResponseWriter, r *http.Request) {

	db, err1 := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")
	if err1 != nil {
		log.Fatal(err1)
	}

	success := true
	message := "Signup success!"
	response := LoginResponse{Success: success, Message: message}
	//println(r.RequestURI)    signup?info=1234|lii
	rows, err := db.Query("SELECT username FROM users WHERE username = substring(?,1+locate('|',?))",
		r.RequestURI, r.RequestURI)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Close the query result set at the end of the function

	// Check if there are results
	if rows.Next() == true {

		// If there are no results, update the response
		response = LoginResponse{Message: "Signup failed!"}
	} else {
		str := r.RequestURI
		re := regexp.MustCompile(`=([^|]+)\|`)
		match := re.FindStringSubmatch(str)
		println(len(match)) // Use regular expressions for string matching
		println(match[0])
		println(match[1])
		if len(match) > 1 {
			result := match[1]
			_, err = db.Exec("insert into users (uid,username,email)values(null,substring(?,1+locate('|',?)),?)",

				r.RequestURI,
				r.RequestURI,
				result)
			if err != nil {
				http.Error(w, "Failed to register user", http.StatusInternalServerError)
				return
			}
		} else {
			response = LoginResponse{Message: "Signup failed!"}
		}

		response = LoginResponse{Message: "Signup success!"}
	}

	// Close the database connection at the end of the function
	defer db.Close()

	// Convert the response structure to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func handleBlog(w http.ResponseWriter, r *http.Request) {

	db, err1 := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")
	var threads []string

	if err1 != nil {
		log.Fatal(err1)
	}
	println(r.RequestURI) //   /blog?topic=     /blog?topic=6
	rows, err := db.Query("select users.username as username,topic,tid from users inner join thread on users.uid=thread.uid where topic "+
		"like concat('%',substring(?,1+locate('=',?)),'%') order by tid desc", r.RequestURI, r.RequestURI)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Close the query result set at the end of the function

	// Check if there are results
	for rows.Next() == true {
		var thread Thread
		rows.Scan(&thread.username, &thread.topic, &thread.tid)
		threads = append(threads, thread.username+","+thread.topic+","+currentusername+","+thread.tid)

	}
	if threads == nil {
		threads = append(threads, ",")
	}
	// Close the database connection at the end of the function
	defer db.Close()

	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(threads)
	//fmt.Println(string(responseJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
func handleCreatetopic(w http.ResponseWriter, r *http.Request) {

	var str = r.RequestURI
	index := strings.Index(str, "=") + 1

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")

	_, err = db.Exec("insert into thread values(null,(select uid from users where username='" + currentusername + "'),'" + strings.Replace(str[index:], "%20", " ", -1) + "')")
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

}

func handleDeleteTopic(w http.ResponseWriter, r *http.Request) {
	success := true
	message := "delete topic success!"
	response := LoginResponse{Success: success, Message: message}

	var str = r.RequestURI
	index := strings.Index(str, "=") + 1

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")

	_, err = db.Exec("delete from subthread where tid=" + str[index:])
	_, err = db.Exec("delete from thread where  tid=" + str[index:])
	//println("delete from subthread where tid=" + str[index:])
	//println("delete from thread where  tid=" + str[index:])
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
func handleSubThread(w http.ResponseWriter, r *http.Request) {

	str := r.RequestURI
	re := regexp.MustCompile("=([^=]+)$") // match[1] retrieves the value of TID
	match := re.FindStringSubmatch(str)
	db, err1 := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")
	sql := ""
	var reverts []string

	if err1 != nil {
		log.Fatal(err1)
	}
	if len(match) > 1 {
		sql = "select revert,(select username from users where uid=ruid)as rusername,stid from subthread where tid=" + match[1]
		currenttid = match[1]
	} else if len(match) <= 1 {
		sql = "select revert,(select username from users where uid=ruid)as rusername,stid from subthread where tid=-1"
	}

	rows, err := db.Query(sql)

	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	sql = "select * from thread where tid=" + match[1] + " and uid=(select uid from users where username='" + currentusername + "')"
	rows1, err1 := db.Query(sql)

	if err1 != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	if rows1.Next() == true { // If data is retrieved, it indicates that we have permission for deletion, so set mydelete to 1
		mydelete = "1"
	} else {
		mydelete = "0"
	}

	defer rows.Close()  // Close the query result set at the end of the function
	defer rows1.Close() // Close the query result set at the end of the function
	// Check if there are results
	for rows.Next() == true {
		var revert Revert
		rows.Scan(&revert.revert, &revert.rusername, &revert.stid)
		reverts = append(reverts, revert.revert+","+revert.rusername+","+revert.stid+","+mydelete)

	}

	if reverts == nil {
		reverts = append(reverts, ",")
	}
	// Close the database connection at the end of the function
	defer db.Close()

	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(reverts)
	//fmt.Println(string(responseJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func handleDeleteSubThread(w http.ResponseWriter, r *http.Request) {
	success := true
	message := "delete reply success!"
	response := LoginResponse{Success: success, Message: message}

	var str = r.RequestURI // Request path
	index := strings.Index(str, "=") + 1

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")

	_, err = db.Exec("delete from subthread where stid=" + str[index:])

	if err != nil {
		http.Error(w, "Failed to  delete reply", http.StatusInternalServerError)
		return
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func handleReply(w http.ResponseWriter, r *http.Request) {
	success := true
	message := "delete reply success!"
	response := LoginResponse{Success: success, Message: message}

	var str = r.RequestURI
	index := strings.Index(str, "=") + 1
	//println(str[index:])    reply

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs")
	sql := "insert into subthread values(null," + currenttid + ",'" + strings.Replace(str[index:], "%20", " ", -1) + "',(select uid from users where username='" + currentusername + "'),(select uid from thread where tid=" + currenttid + "))"
	println(sql)
	_, err = db.Exec(sql)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	sql = "select (select username from users where uid=thread.uid limit 0,1) as username,topic,tid from thread where tid=" + currenttid
	rows, err := db.Query(sql)

	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	if rows.Next() == true {
		var thread Thread
		rows.Scan(&thread.username, &thread.topic, &thread.tid)
		response = LoginResponse{Success: success, Message: "?username=" + thread.username + "&topic=" + thread.topic + "&tid=" + thread.tid}
	}

	defer rows.Close() // Close the query result set at the end of the function

	responseJSON, err := json.Marshal(response)
	//fmt.Println(string(responseJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
