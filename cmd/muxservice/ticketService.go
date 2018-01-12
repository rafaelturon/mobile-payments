package muxservice

// Ticket store basic information regarding PoS mining
type Ticket struct {
	OwnMempool   int     `json:"ownMempool"`
	Immature     int     `json:"immature"`
	Live         int     `json:"live"`
	TotalSubsidy float32 `json:"totalSubsidy"`
}

// GetTickets return Ticket object
func GetTickets() (Ticket, error) {
	var ticket Ticket

	ticket.Immature = 2
	ticket.Live = 3
	ticket.OwnMempool = 1
	ticket.TotalSubsidy = 130.32

	//err := e.marshal(v, encOpts{escapeHTML: true})
	//if err != nil {
	//	return nil, err
	//}
	return ticket, nil
}
