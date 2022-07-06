package tools

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"ow-udp-tracker/schemes"
	"strconv"
	"strings"
)

type handleFunc func(payload schemes.Player, raw []byte)

type Worker struct {
	handle handleFunc
}

func NewWorker(handler handleFunc) *Worker {
	return &Worker{
		handle: handler,
	}
}

func (w *Worker) Start() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	deviceID := ""
	for _, device := range devices {
		if strings.Contains(device.Name, "Loopback") {
			deviceID = device.Name
			break
		}
	}

	if deviceID == "" {
		log.Panic("failed to find loopback device\n")
		return
	}

	go func() {
		if handle, err := pcap.OpenLive(deviceID, 1600, true, pcap.BlockForever); err != nil {
			panic(err)
		} else if err := handle.SetBPFFilter("udp and port 4242"); err != nil { // optional
			panic(err)
		} else {
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				applicationLayer := packet.ApplicationLayer()
				if applicationLayer != nil {
					w.handlePacket(applicationLayer.Payload())
				}
			}
		}
	}()
}

func (w *Worker) handlePacket(data []byte) {
	var f schemes.Player
	err := json.Unmarshal(data, &f)
	if err != nil {
		return
	}

	for idx := range f.Endorsement {
		f.Endorsement[idx].SID = Hex(f.Endorsement[idx].ID)
		f.Endorsement[idx].Name = schemes.EndorsementNames[f.Endorsement[idx].SID]
	}

	f.AvatarSID = Hex(f.Avatar)
	f.PlayerLevelFrameSID = Hex(f.PlayerLevelFrame)
	f.AccountId = FromHex(strings.Split(f.Account, ":")[0])
	f.SecondaryAccountId = FromHex(strings.Split(f.SecondaryAccount, ":")[0])

	if w.handle != nil {
		w.handle(f, data)
	}
}

func Hex(id int64) string {
	return fmt.Sprintf("%X", id)
}

func FromHex(id string) string {
	// I forgot the base jim...
	d, _ := strconv.ParseInt(id, 16, 64)

	return fmt.Sprintf("%d", d)
}
