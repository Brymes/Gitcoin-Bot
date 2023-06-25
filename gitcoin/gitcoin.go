package gitcoin

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"log"
)

type GitcoinService struct {
	pageSize int
}

var UrlMap = map[int]string{
	//1: "https://bounties.gitcoin.co/api/v0.1/grants_clr/?page_size=200",
	1: "https://bounties.gitcoin.co/grants/cards_info?page=1&limit=20&me=false&sort_option=weighted_shuffle&collection_id=false&network=mainnet&state=active&profile=false&customer_name=false&sub_round_slug=false&idle=true&featured=true&round_type=false&tab=grants",
	2: "https://bounties.gitcoin.co/api/v0.1/bounties/slim/?network=mainnet&idx_status=open&applicants=ALL&offset=0&limit=50",
}

func (gs GitcoinService) GetBounties(logger *log.Logger) []*discordgo.MessageEmbed {
	var (
		response   []BountiesResponse
		allUpdates []*discordgo.MessageEmbed
	)
	resBody := makeRequest(UrlMap[2], logger)
	_ = json.Unmarshal(resBody, &response)

	for _, bounty := range response {
		allUpdates = append(allUpdates, bounty.createBountyEmbed())
	}

	return allUpdates
}
