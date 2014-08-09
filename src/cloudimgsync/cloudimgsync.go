package main

import (
	"os/exec"
	"log"
	"os"
)

type image struct{
	src string
	dst string
	num string
	fileName string
}


func (img image) copy(){
	cmd := exec.Command("juju", "run", "--machine=" + string(img.num), "sudo mkdir -p /var/cache/lxc/cloud-trusty")
	out, err := cmd.Output()
	if err != nil{
		log.Fatal(err)
	}
	log.Print(out)
	scp := exec.Command("juju", "scp", img.src, string(img.num) + "~/")
	scpOut, scpErr := scp.Output()
	if scpErr != nil{
		log.Fatal(scpErr)
	}
	log.Print(scpOut)
}

func (img image) mv(){
	cmd := exec.Command("juju", "run", "--machine=" + string(img.num), "sudo mv /home/ubuntu/" + img.fileName, img.dst)
	out, err := cmd.Output()
	if err != nil{
		log.Fatal(err)
	}
	log.Print(out)
}

func main() {
	cloudImage :=  image{src: os.Args[1], num: os.Args[2], dst: "/var/cache/lxc/cloud-trusty", fileName: "ubuntu-14.04-server-cloudimg-amd64-root.tar.gz"}
	cloudImage.copy()
	cloudImage.mv()
}



