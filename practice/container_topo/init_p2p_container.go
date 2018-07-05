package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os/exec"
	"io/ioutil"
	"os"
)

var (
	dockername1 string
	dockername2 string
	docker1Ip string
	docker2Ip string
	cmd string
)

// 获取命令行参数
/*
func argsParser() {
	app := cli.NewApp()

	app.Name = "init_network"
	app.Usage = "./init_p2p_container <docker_name1> <docker_name2>"
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			dockername1 = c.Args().Get(0)
			dockername2 = c.Args().Get(1)
		}
		return nil
	}
	app.Run(os.Args)
}
*/
func argsParser(){
	app := cli.NewApp()
	app.Name = `./init_p2p_container -c1 <docker1_Name> -n1 <docker1_IP> -c2 <docker2_Name> -n2 <docker2_IP> init
	 ./init_p2p_container -c1 <docker1_Name> -c2 <docker2_Name> remove
`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "container1Name, c1",
			Value:       "docker1",
			Usage:       "Input container1 name.",
			Destination: &dockername1,
		},
		cli.StringFlag{
			Name:        "container1Ip, n1",
			Value:       "192.168.10.1/24",
			Usage:       "Input container1 IP address.",
			Destination: &docker1Ip,
		},
		cli.StringFlag{
			Name:        "container2, c2",
			Value:       "docker2",
			Usage:       "Input container2 name.",
			Destination: &dockername2,
		},
		cli.StringFlag{
			Name:        "container2Ip, n2",
			Value:       "192.168.10.2/24",
			Usage:       "Input container2 IP address.",
			Destination: &docker2Ip,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cmd = c.Args().Get(0)
			if cmd == "init" {
				fmt.Println("Container1_Name:", dockername1, "    Container1_IP:", docker1Ip)
				fmt.Println("Container2_Name:", dockername2, "    Container2_IP:", docker2Ip)
				fmt.Println("Command:", cmd)
				initNetwork()
			}else if cmd == "remove"{
				fmt.Println("Container1_Name:", dockername1)
				fmt.Println("Container2_Name:", dockername2)
				fmt.Println("Command:", cmd)
				removeNetwork()
			}else {
				fmt.Println("Usage: ./init_p2p_container -c1 <docker1_Name> -n1 <docker1_IP> -c2 <docker2_Name> -n2 init")
			}
		}
		return nil
	}

	app.Run(os.Args)
}

// 执行系统命令并打印输出
func runCmd(s string) (str string) {
	//cmd := exec.Command(s)
	cmd := exec.Command("/bin/bash", "-c", s)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:can not obtain stdout pipe for command:", cmd)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err.", err)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("Readall stdout:", err.Error())
		return
	}
	if len(bytes) >0 {
		fmt.Printf("Stdout: %s\n", bytes)
	}
	return string(bytes)
}

//判断文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//创建容器，并构建点对点连接
func initNetwork() {
	//run containers
	RunDocker1Cmd := fmt.Sprintf(
		"docker run -d --name %s --net=none ubuntu:ip tail -f /dev/null", dockername1)
	runCmd(RunDocker1Cmd)
	RunDocker2Cmd := fmt.Sprintf(
		"docker run -d --name %s --net=none ubuntu:ip tail -f /dev/null", dockername2)
	runCmd(RunDocker2Cmd)

	getDocker1IidCmd := fmt.Sprintf(`docker inspect --format "{{.State.Pid}}" %s`, dockername1)
	docker1Id := runCmd(getDocker1IidCmd)[0:5]
	getDocker2IidCmd := fmt.Sprintf(`docker inspect --format "{{.State.Pid}}" %s`, dockername2)
	docker2Id := runCmd(getDocker2IidCmd)[0:5]

	if exists("/var/run/netns") != true {
		runCmd("mkdir /var/run/netns")
	}

	//create container netns dir
	docker1NsDir := fmt.Sprintf("/var/run/netns/%s", docker1Id)
	if exists(docker1NsDir) != true {
		createDocker1NsDirCmd := fmt.Sprintf("ln -s /proc/%s/ns/net /var/run/netns/%s",
			docker1Id,
			docker1Id)
		runCmd(createDocker1NsDirCmd)
	}
	docker2NsDir := fmt.Sprintf("/var/run/netns/%s", docker2Id)
	if exists(docker2NsDir) != true {
		createDocker2NsDirCmd := fmt.Sprintf("ln -s /proc/%s/ns/net /var/run/netns/%s",
			docker2Id,
			docker2Id)
		runCmd(createDocker2NsDirCmd)
	}


	veth1 := fmt.Sprintf("v%s%s",
		string(dockername1[len(dockername1)-1]),
		string(dockername2[len(dockername2)-1]))
	veth2 := fmt.Sprintf("v%s%s",
		string(dockername2[len(dockername2)-1]),
		string(dockername1[len(dockername1)-1]))


	// create virtual link patch
	linkAddCmd := fmt.Sprintf("ip link add %s type veth peer name %s", veth1, veth2)
	runCmd(linkAddCmd)


	//set docker1 ip netns
	linkSet1Cmd := fmt.Sprintf("ip link set %s netns %s", veth1, docker1Id)
	runCmd(linkSet1Cmd)
	linkset1NsCmd := fmt.Sprintf("ip netns exec %s ip link set %s up", docker1Id, veth1)
	runCmd(linkset1NsCmd)
	linkSet1IPCmd := fmt.Sprintf("ip netns exec %s ip addr add %s dev %s",
		docker1Id, docker1Ip, veth1)
	runCmd(linkSet1IPCmd)


	//set docker2 ip netns
	linkSet2Cmd := fmt.Sprintf("ip link set %s netns %s", veth2, docker2Id)
	runCmd(linkSet2Cmd)
	linkset2NsCmd := fmt.Sprintf("ip netns exec %s ip link set %s up", docker2Id, veth2)
	runCmd(linkset2NsCmd)
	linkSet2IPCmd := fmt.Sprintf("ip netns exec %s ip addr add %s dev %s",
		docker2Id, docker2Ip, veth2)
	runCmd(linkSet2IPCmd)
}

func removeNetwork()  {
	getDocker1IidCmd := fmt.Sprintf(`docker inspect --format "{{.State.Pid}}" %s`, dockername1)
	docker1Id := runCmd(getDocker1IidCmd)[0:5]
	getDocker2IidCmd := fmt.Sprintf(`docker inspect --format "{{.State.Pid}}" %s`, dockername2)
	docker2Id := runCmd(getDocker2IidCmd)[0:5]

	RunRemoveNetns1Cmd := fmt.Sprintf("ip netns delete %s", docker1Id)
	RunRemoveNetns2Cmd := fmt.Sprintf("ip netns delete %s", docker2Id)
	runCmd(RunRemoveNetns1Cmd)
	runCmd(RunRemoveNetns2Cmd)

	RunDocker1Cmd := fmt.Sprintf(
		"docker rm -f %s", dockername1)
	runCmd(RunDocker1Cmd)
	RunDocker2Cmd := fmt.Sprintf(
		"docker rm -f %s", dockername2)
	runCmd(RunDocker2Cmd)
}

func main() {
	argsParser()
}
