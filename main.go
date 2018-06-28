package main

import (
	"net/http"
	"flag"
	"log"
	"os"
)

type Config struct {
	RootDir string
	Port    string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From PC :)"))
}

func exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func main() {
	fPort := flag.String("port", "9410", "port for server")
	fRootDir := flag.String("rootDir", "", "root directory for server")
	flag.Parse()

	if *fRootDir == "" {
		log.Fatalf("Have not set root dir!")
		return
	}
	if !exists(*fRootDir) {
		log.Fatalf("Dir not exists")
		return
	}
	log.Println(*fRootDir)

	http.HandleFunc("/", indexHandler)
	http.Handle("/view/", http.StripPrefix("/view", http.FileServer(http.Dir(*fRootDir))))
	http.ListenAndServe("0.0.0.0:" + *fPort, nil)
}
