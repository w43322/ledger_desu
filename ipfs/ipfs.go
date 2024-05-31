package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"log"
	"os"
)

func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}
func UploadIPFS(raw []byte) (string, error) {
	sh := shell.NewShell("localhost:5001")
	reader := bytes.NewReader(raw)
	// https://github.com/ipfs/go-ipfs-api/blob/master/add.go
	fileHash, err := sh.Add(reader)
	if err != nil {
		return "", err
	}
	fmt.Println(fileHash)
	return fileHash, nil
}
func WriteHash(writeJson string, cont interface{}) {
	//
	if distFile, err := os.OpenFile(writeJson, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		log.Println("create file failed", err)
	} else {
		enc := json.NewEncoder(distFile)
		if err1 := enc.Encode(cont); err1 != nil {
			log.Println("write failed", err1)
		} else {
			log.Println("write successful")
		}
	}
}
func test() {
	hashMap := make(map[int]string, 10000)
	for i := 0; i < 10000; i++ {
		file := fmt.Sprintf("./greencard/green_%d.gif", i)
		raw := Read(file)
		if raw != nil {
			hash, err := UploadIPFS(raw)
			if err != nil {
				log.Println("UploadIPFS err", err)
			} else {
				hashMap[i] = fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", hash, hash)
			}
			log.Println("hash", hash)
		}
	}
	WriteHash("hash.json", hashMap)
}
