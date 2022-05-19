package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sosedoff/ansible-vault-go"
	"gopkg.in/yaml.v2"
)

var vaultFile string
var passwordFile string
var listKeys bool
var getAllKeys bool
var getKey string
var putKey string
var putValue string

func init() {
	flag.StringVar(&vaultFile, "vaultFile", "", "Ansible vault file.")
	flag.StringVar(&vaultFile, "V", "", "Short form of -vaultFile")
	flag.StringVar(&passwordFile, "passwordFile", "", "File containing the vault password.")
	flag.StringVar(&passwordFile, "P", "", "Short form of -passwordFile")

	flag.BoolVar(&listKeys, "list", false, "Show a full list of keys, but without their values.")
	flag.BoolVar(&getAllKeys, "view", false, "Show the full contents of the vault.")
	flag.StringVar(&getKey, "get", "", "Get a key.")
	flag.StringVar(&putKey, "put", "", "Put a key. This modifies the vault.")
	flag.StringVar(&putValue, "putval", "", "Value for -put.")

	flag.Parse()
}

// TODO: refactor to combine duplicate code among getAll(), get(), and put().

func getAll(vaultFile string, password string) {
	yaml_str, err := vault.DecryptFile(vaultFile, password)
	if err != nil {
		log.Fatalf("error decrypting file: %s\n", err)
	}
	fmt.Printf("%s", yaml_str)
}

func list(vaultFile string, password string) {
	yaml_str, err := vault.DecryptFile(vaultFile, password)
	yaml_bin := []byte(yaml_str)
	if err != nil {
		log.Fatalf("error decrypting file: %s\n", err)
	}

	m := make(map[string]string)
	err = yaml.Unmarshal(yaml_bin, &m)
	if err != nil {
		log.Fatalf("error unmarshaling yaml %s\n", err)
	}

    for k, _ := range(m) {
        fmt.Printf("%s\n", k)
    }
}


func get(vaultFile string, password string, key string) {
	yaml_str, err := vault.DecryptFile(vaultFile, password)
	yaml_bin := []byte(yaml_str)
	if err != nil {
		log.Fatalf("error decrypting file: %s\n", err)
	}

	m := make(map[string]string)
	err = yaml.Unmarshal(yaml_bin, &m)
	if err != nil {
		log.Fatalf("error unmarshaling yaml %s\n", err)
	}

	fmt.Printf("%s\n", m[key])
}

func put(vaultFile string, password string, key string, value string) {
	// Read.
	yaml_str, err := vault.DecryptFile(vaultFile, password)
	yaml_bin := []byte(yaml_str)
	if err != nil {
		log.Fatalf("error decrypting file: %s\n", err)
	}

	m := make(map[string]string)
	err = yaml.Unmarshal(yaml_bin, &m)
	if err != nil {
		log.Fatalf("error unmarshaling yaml: %s\n", err)
	}

	// Set value.
	m[key] = value

	// Write back.
	yaml_bin, err = yaml.Marshal(&m)
	yaml_str = string(yaml_bin)
	if err != nil {
		log.Fatalf("while marshaling map %s", err)
	}

	err = vault.EncryptFile(vaultFile, yaml_str, password)
	if err != nil {
		log.Fatalf("error encrypting file: %s\n", err)
	}
}

func main() {
	if vaultFile == "" || passwordFile == "" {
		flag.Usage()
		log.Fatalf("Error, both -vaultFile and -passwordFile options are required.\n")
	}

	var password string = ""
	bytes, err := os.ReadFile(passwordFile)
	if err != nil {
		log.Fatalf("could not read password file: %s\n", err)
	}
	password = strings.TrimSpace(string(bytes))

	if getAllKeys {
		getAll(vaultFile, password)
	}

	if listKeys {
		list(vaultFile, password)
	}


	if getKey != "" {
		get(vaultFile, password, getKey)
	}

	if putKey != "" {
		// It's fine if putValue is "".
		put(vaultFile, password, putKey, putValue)
	}
}
