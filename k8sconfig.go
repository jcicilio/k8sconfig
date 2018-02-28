package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/spf13/pflag"
)

const PREFIX = "K8SCONFIG"
const NAME = "vName"
const ENVIRONMENTVALUE = "environment"
const DEFAULTVALUE = "default"
const COMMANDLINEVALUE = "command-line"
const CONFIGFILEVALUE = "configuration-file"
const CONFIGFILENAMEPREFIX = "k8sconfig"

const URLDEFAULTVALUE = "0.0.0.0:80"
const URLDEFAULTNAME = "url"

// WriteStandardHeaders ... Write standard headers for each page, these include
// the content type and the allow-origin.
func writeStandardHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/x-javascript")

	// Allow applications from another domain to use Api
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	writeStandardHeaders(w)
	w.WriteHeader(http.StatusOK)
	var buffer bytes.Buffer

	buffer.WriteString("{'configuration':true, 'value':'" + getValue() + "'}")
	w.Write(buffer.Bytes())
}

// Router ... Routes definition,  kept in distinct function to allow
// for compatibility with test writing
func Router() *mux.Router {
	// Setup route handlers
	r := mux.NewRouter()
	r.HandleFunc("/{id:(?i)config}", configHandler).Methods("GET")

	return r
}

func configSetup() {
	viper.SetEnvPrefix(PREFIX) // will be uppercased automatically
	viper.AutomaticEnv()

	viper.SetDefault(NAME, DEFAULTVALUE)
	viper.SetDefault(URLDEFAULTNAME, URLDEFAULTVALUE)

	viper.SetConfigName(CONFIGFILENAMEPREFIX) // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
}

func configReadConfigFile() {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("error reading in configuration file %v\n", CONFIGFILENAMEPREFIX)
	}
}

func configCommandLine(){
	pflag.String("url", viper.GetString(URLDEFAULTNAME), fmt.Sprintf("The base URL for the service. Default to %v", getConfigString(URLDEFAULTNAME)))
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func getValue() string {
	return viper.GetString(NAME)
}

func getConfigString(s string) string {
	return viper.GetString(s)
}

func main() {
	fmt.Println("starting k8sconfigtest... ")

	configSetup()
	configReadConfigFile()
	configCommandLine()

	fmt.Printf("on start test variable value is: %v\n", getValue())

	http.Handle("/", Router())

	// Start the http server
	fmt.Printf("listening on %v\n", getConfigString(URLDEFAULTNAME))
	log.Fatal(http.ListenAndServe(getConfigString(URLDEFAULTNAME), nil))
}
