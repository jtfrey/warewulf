package main

import (
	"fmt"
	"github.com/hpcng/warewulf/internal/pkg/vnfs"
	"os"
	"path"
	"sync"
)

func vnfsLocalBuild(vnfsPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	v := vnfs.New(vnfsPath)

	if _, err := os.Stat(vnfsPath); err == nil {
		// TODO: Build VNFS to temporary file and move to real location when complete atomically
		// TODO: Check time stamps of sourcedir and build file to see if we need to rebuild or skip

		//vnfsDestination := fmt.Sprintf("%s/provision/vnfs/%s.img.gz", LocalStateDir, path.Base(vnfsPath))

		fmt.Printf("Building local Container: %s\n", vnfsPath)

		err := os.MkdirAll(path.Dir(v.Image()), 0755)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}

		err = buildVnfs(vnfsPath, v.Image())
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}

		err = buildLinks(v, vnfsPath)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}

	} else {
		fmt.Printf("SKIPPING VNFS:  (bad path) %s\n", vnfsPath)
	}
}
