
package main

import (
	"fmt"
	"net"
	"time"
	"os/exec"


const host = "129.241.187.255"
const port = "20015"

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}
}

func main() {

	sendAdd,err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
	checkError(err)
	conn,err := net.ListenUDP("udp",sendAdd)
	checkError(err)

	msg := make(byte[], 8)

	var master bool = false

	for !master {
		conn.SetReadDeadline()

		n, ,err := conn.ReadFromUDp(msg) 

		if err == nil{
			//lagre msg
		}
		else{
			master = true
		}
	}

	//nå må



}
func newbackup(){
	cmd = exec.Command()
	cmd.Run()
}



	conn.Close()	
	
	fmt.Println("I am now Master")
	spawnMaster()
	conn, _ = net.DialUDP("udp", nil ,udpaddr)	
		
	for { 
		
		fmt.Println(currentNum)
		currentNum++
		binary.BigEndian.PutUint64(udpmessage, currentNum)
		_, _ = conn.Write(udpmessage)
		
		time.Sleep(time.Second)
	}


