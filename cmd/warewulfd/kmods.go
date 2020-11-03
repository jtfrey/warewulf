package main

import (
	"fmt"
	"log"
	"net/http"
)

func kmods(w http.ResponseWriter, req *http.Request) {

	node, err := getSanity(req)
	if err != nil {
		w.WriteHeader(404)
		log.Panicln(err)
		return
	}
	if node.KernelVersion != "" {
		fileName := fmt.Sprintf("%s/provision/kernels/kmods-%s.img", LocalStateDir, node.KernelVersion)

		err := sendFile(w, fileName, node.Fqdn)
		if err != nil {
			log.Println(err)
		}

	} else {
		w.WriteHeader(503)
		log.Printf("No kernel version set for node %s\n", node.Fqdn)
	}

	return
}