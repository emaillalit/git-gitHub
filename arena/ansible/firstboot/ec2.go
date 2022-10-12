// Version Notes:
// The post data from user data is a form
// Keep handling other POST requests even error occurs
// chmod +x run_ansible.sh
// Add health check handler
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("/firstboot", firstbootHandler)
	http.HandleFunc("/", defaultHandler)
	fmt.Printf("Starting server for handling EC2 firstboot...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, `You are visiting Arena Firstboot Server!`)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func firstbootHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		panic(err)
	}

	instanceID := r.FormValue("instance_id")

	if instanceID == "" {
		fmt.Fprintln(w, "DID NOT RECEIVE AN INSTANCE ID!")
	} else {
		// fmt.Printf("Executing Firstboot for instance %s\n", instanceID)
		fmt.Fprintf(w, "Executing firstboot for instance %s\n", instanceID)

		//Execute shell
		err := os.Chmod("run_ansible.sh", 0755)
		if err != nil {
			fmt.Println(err)
		}

		var outInfo bytes.Buffer
		cmd := exec.Command("./run_ansible.sh", instanceID)

		cmd.Stdout = &outInfo

		err = cmd.Run()
		if err != nil {
			//log.Fatal(err)
			// fmt.Println(err)
			fmt.Fprintln(w, err)
		}

		// fmt.Println(outInfo.String())
		fmt.Fprintf(w, outInfo.String())
	}
}
