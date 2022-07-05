# Overwatch UDP Broadcasting

## What is this?
This repo allows you to view/track the UDP broadcast packet of nearby players from Overwatch.

## What does the packet do?
The UDP broadcast packet is what Overwatch uses to check if other players are nearby (same network).  
The actual packet is just a UDP broadcast on port `4242` and is just JSON.

## Can I view the packet in Wireshark?
Yes, you can view the packet in Wireshark. You just need to use the rule `udp.port == 4242` on your `loopback` host. 

## What is in this packet?
The data contained in the packet is pretty minimal. It just contains the following:
- `version`: protocol version
- `build`: build number of the client
- `account`: the current logged in account seperated by `:`
- `secondary_account`: (UNKNOWN) currently same as `account`
- `avatar`: the current avatar of the user (Hex)
- `level`: the current level of the user
- `pframe`: the level frame of the user (Hex)
- `elevel`: the current endorsement level of the user
- `endors`: the breakdown of the endorsement level of the user
  - `endors.id`: the id of the endorsement (Hex)
  - `endors.count`: the count of the endorsement level

Below is a list of all the actual `endors.id` from hex to name (this is in go):
```go
EndorsementNames := map[string]string{
	"D80000000003944": "Shotcaller",
	"D80000000003945": "Good Teammate",
	"D80000000003946": "Sportsmanship",
}
```

## Notes
- The packet is sent out every 5 seconds.
- It's always a javascript payload and is not encrypted.
- It's always sent out on port `4242`.
