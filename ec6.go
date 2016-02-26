package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

const add = "129.241.187.255:20015"

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}
}

func main() {

	sendAdd, err := net.ResolveUDPAddr("udp", add)
	checkError(err)
	conn, err := net.ListenUDP("udp", sendAdd)
	checkError(err)

	msg := make([]byte, 8)
	var currentNum uint64 = 0
	var master bool = false
	fmt.Printf("Backup \n")
	for !master {
		t := time.Now()
		d := time.Second * 3
		conn.SetReadDeadline(t.Add(d))

		number, _, err := conn.ReadFromUDP(msg)

		if err == nil {
			currentNum = binary.BigEndian.Uint64(msg[0:number]) //BigEndian stores the MSB of a word at a particular memory address
		} else {
			master = true
		}
	}
	conn.Close()
	//nå må dette bli master
	//og kaller ny backup

	newbackup()
	fmt.Printf("Master \n")

	conn, err = net.DialUDP("udp", nil, sendAdd)
	//evig løkke
	for {
		fmt.Println(currentNum)
		currentNum++

		binary.BigEndian.PutUint64(msg, currentNum)
		conn.Write(msg)
		time.Sleep(time.Second)
	}

}
func newbackup() { //Her får vi opp en ny terminal som backup
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run ec6.go")

	cmd.Run()

}
