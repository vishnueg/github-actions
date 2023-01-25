package utils

import (
	"encoding/json"
	datamodel "initSetupScripts/Datamodel"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func GetAssetInfo(macAddr string) []datamodel.Asset {
	configs := GetServerConfigs()
	secrets := GetServerSecrets()
	var bearer string = secrets["squirrelAdminApiToken"]
	var api = configs["queryAssetApi"]
	url := api + macAddr
	log.Println("Making a get request to following url: ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while initializing a new get request")
		log.Println(err)
		return nil
	}

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}

	log.Println("Making the GET request, to fetch asset info")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making the request to the following url: ", url)
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var asset []datamodel.Asset
	log.Println("Unmarshalling json response to struct")
	err = json.NewDecoder(resp.Body).Decode(&asset)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil
	}
	return asset
}

func CreateTarFile(parent, folder, name string) int {
	cmd := exec.Command("tar", "-C", parent, "-cvf", name, folder)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("!!!!!!! Error while tring to init a output pipe ", "err>>", err)
		return 1
	}
	stdErrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Println("!!!!!!! Error while tring to init a error pipe ", "err>>", err)
		return 1
	}
	err = cmd.Start()
	if err != nil {
		log.Println("!!!!!!! Error while tyring to start the cmd ", "err>>", err)
		return 1
	}
	cmdOp, err := io.ReadAll(stdoutPipe)
	if err != nil {
		log.Println("!!!!!!! Error while tring to process command output", "err>>", err)
		return 1
	}
	log.Println(string(cmdOp))
	cmdErr, err := io.ReadAll(stdErrPipe)
	if err != nil {
		log.Println("!!!!!!! Error while tring to process command Error", "err>>", err)
		return 1
	}
	if len(cmdErr) > 0 {
		log.Println(string(cmdErr))
		return 1
	}
	cmd.Wait()
	return 0
}

func CheckFileExistOrNot(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	return nil
}

func GetServerConfigs() map[string]string {
	data := make(map[string]string)
	configs, err := os.ReadFile(CONFIG_FOLDER + SERVER_CONFIG)
	if err != nil {
		log.Println("Error while trying to read server configs from: ", CONFIG_FOLDER+SERVER_CONFIG)
		log.Println("Error>>", err)
		return nil
	}

	err = yaml.Unmarshal(configs, &data)

	if err != nil {
		log.Println("Error while trying to unmarshall server configs ")
		log.Println("Error>>", err)
	}

	return data
}

func GetServerSecrets() map[string]string {
	data := make(map[string]string)
	configs, err := os.ReadFile(CONFIG_FOLDER + SECRETS_CONFIG)
	if err != nil {
		log.Println("Error while trying to read server configs from: ", CONFIG_FOLDER+SERVER_CONFIG)
		log.Println("Error>>", err)
		return nil
	}

	err = yaml.Unmarshal(configs, &data)

	if err != nil {
		log.Println("Error while trying to unmarshall server configs ")
		log.Println("Error>>", err)
	}

	return data
}
