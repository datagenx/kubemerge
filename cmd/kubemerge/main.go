package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: kubemerge <config1> <config2> ....")
		os.Exit(0)
	}

	backup()
	merge(args)
}

func backup() {
	home, _ := os.UserHomeDir()
	kubeDir := home + "/" + ".kube"

	// creating ~/.kube directory
	os.Mkdir(kubeDir, 0755)

	file := kubeDir + "/" + "config"
	backup_file := file + ".bak"

	// checking if backup file exists, then adding UNix time with filename
	isFile, _ := isExists(backup_file)
	if isFile {
		backup_file = fmt.Sprintf("%s_%d.bak", file, time.Now().Local().Unix())
	}

	// Renaming the file
	err := os.Rename(file, backup_file)
	if err != nil {
		log.Fatalf("Unable to rename %s.", file)
	}

}

func merge(args []string) {

}

func isExists(file string) (bool, error) {
	if _, err := os.Stat(file); err == nil {
		log.Printf("%s exists.", file)
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		log.Printf("%s does not exists.", file)
		return false, nil

	} else {
		log.Println("Err:", err)
		return false, err
	}
}
