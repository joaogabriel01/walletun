package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"transaction/internal/core"
	"transaction/internal/queue/nats"
)

func main() {
	natsConn, err := nats.NewNatsConnection("nats://localhost:4222")
	natsPub := nats.NewNatsPublisher(natsConn)
	if err != nil {
		panic(err)
	}

	natsSub := nats.NewNatsSubscriber(natsConn)
	err = natsSub.Subscribe("transactions", func(data []byte) {
		var transaction core.Transaction
		err := json.Unmarshal(data, &transaction)
		if err != nil {
			panic(err)
		}
		log.Printf("Received transaction: %+v", transaction)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var transaction core.Transaction
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &transaction)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(transaction)
		if err != nil {
			http.Error(w, "Error encoding transaction", http.StatusInternalServerError)
			return
		}
		err = natsPub.Publish("transactions", data)
		if err != nil {
			http.Error(w, "Error publishing transaction", http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
