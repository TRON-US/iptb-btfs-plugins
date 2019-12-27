package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type ports struct {
	node_id    int
	swarm      string
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
		var swarm_port string
		var api_port string
		var remote_api_port string

		//separate the lines on newline character
		lines := strings.Split(node_logs[i], "\n")
		for j := 0; j < (len(lines)); j++ {
			if strings.HasPrefix(lines[j], "Swarm listening on /ip4/") {
				//split to find port value
				parts := strings.Split(lines[j], "/")
				if len(parts) > 0 {
					swarm_port = parsePort(parts[len(parts)-1])
					fmt.Println("swarm port is: ", swarm_port, " for node: ", i)
				} else {
					fmt.Println("Unable to parse Swarm port")
				}
			}
			if strings.HasPrefix(lines[j], "API server listening") {
				//split to find port value
				parts := strings.Split(lines[j], "/")
				if len(parts) > 0 {
					api_port = parsePort(parts[len(parts)-1])
					fmt.Println("api server port is: ", api_port, " for node: ", i)
				} else {
					fmt.Println("Unable to parse API port")
				}
			}
			if strings.HasPrefix(lines[j], "Remote API server listening") {
				//split to find port value
				parts := strings.Split(lines[j], "/")
				if len(parts) > 0 {
					remote_api_port = parsePort(parts[len(parts)-1])
					fmt.Println("remote api server port is: ", remote_api_port, " for node: ", i)
				} else {
					fmt.Println("Unable to parse Remote API port")
				}
			}
		}

		my_node := new(ports)
		my_node.node_id = i
		my_node.swarm = swarm_port
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

		//allow connections from external nodes
		new_api_string := `    "API": "/ip4/0.0.0.0/tcp/` + nodePorts[i].api + `",`
		new_remote_api_string := `    "RemoteAPI": "/ip4/0.0.0.0/tcp/` + nodePorts[i].remote_api + `",`
		new_swarm_string := `      "/ip4/0.0.0.0/tcp/` + nodePorts[i].swarm + `"`

		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		//first pass
		newContents := strings.Replace(string(read), `    "API": "/ip4/127.0.0.1/tcp/0",`, new_api_string, -1)
		//second pass
		newContents2 := strings.Replace(newContents, `    "RemoteAPI": "/ip4/127.0.0.1/tcp/0",`, new_remote_api_string, -1)
		//third pass to set the swarm port do this last since there should only be one match 127.0.0.1/tcp/0 left
		newContents3 := strings.Replace(newContents2, `      "/ip4/127.0.0.1/tcp/0"`, new_swarm_string, -1)

		err = ioutil.WriteFile(path, []byte(newContents3), 0)
		if err != nil {
			panic(err)
		}
	}

}

func parsePort(port string) string {
	_, err := strconv.Atoi(port)
	if err != nil {
		panic("The port value cannot be read")
	}
	return port
}
