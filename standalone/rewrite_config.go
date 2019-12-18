package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ports struct {
	node_id    int
	api        string
	remote_api string
}

func main() {
	nodePorts := []*ports{}

	file, err := os.Open("iptb_logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)

	log := string(contents)

	//separate combined log into a log for each node
	node_logs := strings.Split(log, "Daemon is ready")

	//separate the lines in each node log
	for i := 0; i < (len(node_logs) - 1); i++ {
		var api_port string
		var remote_api_port string

		//separate the lines on newline character
		lines := strings.Split(node_logs[i], "\n")
		for j := 0; j < (len(lines)); j++ {
			if strings.HasPrefix(lines[j], "API server listening") {
				//split to find port value
				parts := strings.Split(lines[j], "/")
				api_port = parts[(len(parts))-1]
				fmt.Println("api server port is: ", api_port, " for node: ", i)
			}
			if strings.HasPrefix(lines[j], "Remote API server listening") {
				//split to find port value
				parts := strings.Split(lines[j], "/")
				remote_api_port = parts[(len(parts))-1]
				fmt.Println("remote api server port is: ", remote_api_port, " for node: ", i)
			}
		}

		my_node := new(ports)
		my_node.node_id = i
		my_node.api = api_port
		my_node.remote_api = remote_api_port
		nodePorts = append(nodePorts, my_node)

	}

	//modify  each of the testbed config files
	for i := range nodePorts {
		local_node_id := fmt.Sprintf("%d", nodePorts[i].node_id)

		//get env var $HOME
		home_path := os.Getenv("HOME")
		path := home_path + "/testbed/testbeds/default/" + local_node_id + "/config"

		new_api_string := `    "API": "/ip4/127.0.0.1/tcp/` + nodePorts[i].api + `",`
		new_remote_api_string := `    "RemoteAPI": "/ip4/127.0.0.1/tcp/` + nodePorts[i].remote_api + `",`

		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		//first pass
		newContents := strings.Replace(string(read), `    "API": "/ip4/127.0.0.1/tcp/0",`, new_api_string, -1)
		//second pass
		newContents2 := strings.Replace(newContents, `    "RemoteAPI": "/ip4/127.0.0.1/tcp/0",`, new_remote_api_string, -1)

		err = ioutil.WriteFile(path, []byte(newContents2), 0)
		if err != nil {
			panic(err)
		}
	}

}
