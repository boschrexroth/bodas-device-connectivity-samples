package frontend

import (
	"fmt"
	"net/http"

	"boschrexroth.com/can-demo-golang/pkg/can"
	ecan "go.einride.tech/can"
)

func StartServer() {
	fmt.Println("Start server...")
	http.HandleFunc("/api", api)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		panic(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("can_id")
		payload := r.FormValue("can_payload")

		sendFrame(id, fmt.Sprintf("%016s", payload))

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func sendFrame(id string, payload string) {
	frame := ecan.Frame{}

	jsonFrame := []byte(`{"id":` + id + `,"data":"` + payload + `"}`)
	frame.UnmarshalJSON(jsonFrame)
	//fmt.Printf("Send frame on can1: \t\t%v\n", frame)

	canDevice := can.NewCanDevice("can1")

	canDevice.CanSend(frame)
}
