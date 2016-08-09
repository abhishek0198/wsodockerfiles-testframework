/*
 ------------------------------------------------------------------------

 Copyright 2016 WSO2, Inc. (http://wso2.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License

 ------------------------------------------------------------------------
*/

package common

import (
	"os/exec"
	"strings"	
)

// Function to clean docker image by its tag
func CleanDockerImage(tag string) {
	command := "docker images -q " + tag
	out, _ := exec.Command("/bin/bash", "-c", command).Output()

	if out != nil && len(out) != 0 {
		Logger.Println("Removing docker image " + tag)
		_, err := exec.Command("/bin/bash", "-c", "docker rmi "+tag).Output()

		if err == nil {
			Logger.Println("Successfully removed docker image")
		}
	}
}

// Function to Stop and Remove a docker container
func StopAndRemoveDockerContainer(productName string) {
	command := "docker ps -a | grep " + productName + " | awk '{print $1}'"
	out, err := exec.Command("/bin/bash", "-c", command).Output()

	if err != nil {
		Logger.Println("Error in getting docker container id")
		Logger.Printf("%s\n", err)
	} else if string(out) != "" {
		Logger.Printf("Stopping and removing docker container with id: %s", out)

		_, err1 := exec.Command("/bin/bash", "-c", "docker stop "+string(out)).Output()
		_, err2 := exec.Command("/bin/bash", "-c", "docker rm "+string(out)).Output()

		if err1 == nil && err2 == nil {
			Logger.Println("Successfully stopped and removed docker container")
		}
	}
}

// Function to copy WSO2 Carbon serer logs from running container to local
func CopyWSO2CarbonLogs(productName string, productVersion string) {
	containerID := GetDockerContainerID(productName)
	containerIP := GetDockerContainerIPUsingID(containerID)
	command := "docker cp " + containerID + ":/mnt/" + containerIP + "/" + productName + "-" +
		productVersion + "/repository/logs/ ./" + productName + productVersion + "logs"
	err := RunCommandAndGetError(command)
	if err != nil {
		Logger.Fatal("Unable to copy carbon server logs from docker container. Command: " + command + ". Error:" + err.Error())
	} else {
		Logger.Println("Successfully copied carbon server logs from docker container")
	}
}

// Function to get docker container id for WSO2 product
func GetDockerContainerID(productName string) string {
	if DockerContainerID == "" {
		command := "docker ps | grep " + productName + " | awk '{print $1}'"
		out, err := exec.Command("/bin/bash", "-c", command).Output()
		if err != nil {
			message := "Error in getting Docker container id. Command: " + command + ". Error: " + err.Error()
			Logger.Println(message)
			panic(message)
		}

		DockerContainerID = strings.Replace(string(out), "\n", "", 1)
		Logger.Println("Setting docker ID " + DockerContainerID)
	}

	return DockerContainerID
}

// Function to get docker container IP addrses
func GetDockerContainerIP(productName string) string {
	if DockerContainerIP == "" {
		if Testconfig.Carbon_Server_Ip != "" {
			DockerContainerIP = Testconfig.Carbon_Server_Ip
		} else {
			containerId := GetDockerContainerID(productName)
			containerIp := GetDockerContainerIPUsingID(containerId)
			
			DockerContainerIP = containerIp
		}
		
		Logger.Println("Setting docker IP " + DockerContainerIP)
	}

	return DockerContainerIP
}

func GetDockerContainerIPUsingID(id string) string {
	command := "docker inspect --format '{{ .NetworkSettings.IPAddress }}' " + id
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		message := "Error in getting docker container IP. Command: " + command + ", Error: " + err.Error()
		Logger.Println(message)
		panic(message)
	}
	return strings.Replace(string(out), "\n", "", 1)
}