package muxservice

import "github.com/davecgh/go-spew/spew"

// Ticket store basic information regarding PoS mining
type Ticket struct {
	OwnMempool   uint32  `json:"ownMempool"`
	Immature     uint32  `json:"immature"`
	Live         uint32  `json:"live"`
	TotalSubsidy float64 `json:"totalSubsidy"`
}

// GetTickets return Ticket object
func GetTickets() (Ticket, error) {
	var ticket Ticket

	stakeinfo, err := client.GetStakeInfo()
	if err != nil {
		logger.Errorf("Error %v", err)
		return ticket, err
	}

	logger.Tracef("StakeInfo:\n%v", spew.Sdump(stakeinfo))
	ticket.Immature = stakeinfo.Immature
	ticket.Live = stakeinfo.Live
	ticket.OwnMempool = stakeinfo.OwnMempoolTix
	ticket.TotalSubsidy = stakeinfo.TotalSubsidy

	return ticket, nil
}
