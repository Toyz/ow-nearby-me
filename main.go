package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"ow-udp-tracker/schemes"
	"ow-udp-tracker/tools"
	"strings"
	"syscall"
	"time"
)

func main() {
	mode := flag.String("mode", "send", "the mode to run (send, receive)")
	flag.Parse()

	switch strings.ToLower(*mode) {
	case "send":
		go sendNearByPacket()
		break
	case "receive":
		startWorker()
		break
	default:
		println("Invalid mode:", *mode)
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func startWorker() {
	worker := tools.NewWorker(handle)
	worker.Start()
}

func handle(player schemes.Player, raw []byte) {
	println("Got packet at:", time.Now().Format("2006-01-02 15:04:05"))
	data, err := player.MarshalIndent()
	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func sendNearByPacket() {
	sendPacket := time.NewTimer(5 * time.Second)
	defer sendPacket.Stop()

	raddr, err := net.ResolveUDPAddr("udp", "224.0.0.5:4242")
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		ips, err := getLocalIps()
		if err != nil {
			log.Println(err)
			return
		}

		socket, err := net.ListenMulticastUDP("udp4", nil, raddr)
		if err != nil {
			log.Println(err)
			return
		}

		p := make([]byte, 2048)
		for {
			_, remoteaddr, err := socket.ReadFromUDP(p)
			myOwn, err := isMyOwnIP(ips, remoteaddr.IP)
			if err != nil {
				log.Println(err)
				continue
			}
			if myOwn {
				continue
			}

			fmt.Printf("Got(%v) %s \n", remoteaddr, p)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
		}
	}()

	for {
		select {
		case <-sendPacket.C:
			payload, err := ioutil.ReadFile("payload.json")
			if err != nil {
				log.Println(err)
				return
			}

			buffer := new(bytes.Buffer)
			if err := json.Compact(buffer, payload); err != nil {
				fmt.Println(err)
			}

			_, err = conn.Write(buffer.Bytes())
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("Send(%v) %s \n", raddr, buffer.Bytes())

			sendPacket.Reset(5 * time.Second)
		}
	}
}

func getLocalIps() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]net.IP, 0)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}

		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil {
				continue
			}

			if ip.IsGlobalUnicast() {
				ips = append(ips, ip)
			}
		}
	}

	return ips, nil
}

func isMyOwnIP(ips []net.IP, incoming net.IP) (bool, error) {
	for _, ip := range ips {
		if ip.Equal(incoming) {
			return true, nil
		}
	}

	return false, nil
}
