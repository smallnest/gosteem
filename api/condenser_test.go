package api

import (
	"testing"
	"time"
)

func TestGetAccountCount(t *testing.T) {
	count, err := GetAccountCount(client)
	if err != nil {
		t.Fatal(err)
	}

	if count < 10000 {
		t.Logf("%s: expected > 10000 but got %d", MethodGetAccountCount, count)
	}

	t.Logf("%s: %d", MethodGetAccountCount, count)
}

func TestGetOwnerHistory(t *testing.T) {
	resp, err := GetOwnerHistory(client, "steemit")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %s", MethodGetOwnerHistory, resp)
}

func TestGetVersion(t *testing.T) {
	resp, err := GetVersion(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %s", MethodGetVersion, resp)
}

func TestGetTrendingTags(t *testing.T) {
	resp, err := GetTrendingTags(client, "steem", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetTrendingTags, resp)
}

func TestGetState(t *testing.T) {
	resp, err := GetState(client, "steem")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetState, resp)
}

func TestGetBlock(t *testing.T) {
	resp, err := GetBlock(client, 7595656)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetBlock, resp)
}

func TestGetBlockHeader(t *testing.T) {
	resp, err := GetBlockHeader(client, 7595656)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetBlockHeader, resp)
}

func TestGetOpsInBlock(t *testing.T) {
	resp, err := GetOpsInBlock(client, 7595656, false)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetOpsInBlock, resp)
}

func TestGetChainProperties(t *testing.T) {
	resp, err := GetChainProperties(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetChainProperties, resp)
}

func TestGetCurrentMedianHistoryPrice(t *testing.T) {
	base, quote, err := GetCurrentMedianHistoryPrice(client)
	if err != nil {
		t.Fatal(err)
	}

	if base == "" || quote == "" {
		t.Fatalf("got base: %s, quote: %s", base, quote)
	}

	t.Logf("%s: %s, %s", MethodGetCurrentMedianHistoryPrice, base, quote)
}

func TestGetHardforkVersion(t *testing.T) {
	version, err := GetHardforkVersion(client)
	if err != nil {
		t.Fatal(err)
	}

	if version == "" {
		t.Fatalf("got empty version")
	}

	t.Logf("%s: %s", MethodGetHardforkVersion, version)
}

func TestGetNextScheduledHardfork(t *testing.T) {
	hfVersion, liveTime, err := GetNextScheduledHardfork(client)
	if err != nil {
		t.Fatal(err)
	}

	if hfVersion == "" || liveTime == (time.Time{}) {
		t.Fatalf("got %s, %v", hfVersion, liveTime)
	}

	t.Logf("%s: %s, %v", MethodGetNextScheduledHardfork, hfVersion, liveTime)
}

func TestGetRewardFund(t *testing.T) {
	resp, err := GetRewardFund(client, "post")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetRewardFund, resp)
}

func TestGetKeyReferences(t *testing.T) {
	accounts, err := GetKeyReferences(client, "STM5Sr42NHoEFuFwRcgM7fAqqeqsVHFJgqpDW68MdUSaVpKGFTmC2")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetKeyReferences, accounts)
}

func TestGetAccounts(t *testing.T) {
	accounts, err := GetAccounts(client, "steemit", "good-karma")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetAccounts, accounts)
}

func TestLookupAccountNames(t *testing.T) {
	accounts, err := LookupAccountNames(client, "steemit", "good-karma")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodLookupAccountNames, accounts)
}

func TestLookupAccounts(t *testing.T) {
	accounts, err := LookupAccounts(client, "steem", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodLookupAccounts, accounts)
}

func TestGetWithdrawRoutes(t *testing.T) {
	resp, err := GetWithdrawRoutes(client, "clayop", "all")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWithdrawRoutes, resp)
}

func TestGetAccountBandwidth(t *testing.T) {
	resp, err := GetAccountBandwidth(client, "clayop", "forum")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetAccountBandwidth, resp)
}

func TestGetWitnesses(t *testing.T) {
	resp, err := GetWitnesses(client, "231", "12286")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnesses, resp)
}

func TestGetWitnessByAccount(t *testing.T) {
	resp, err := GetWitnessByAccount(client, "steemit")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnesses, resp)
}

func TestGetWitnessesByVote(t *testing.T) {
	resp, err := GetWitnessesByVote(client, "steemit", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnesses, resp)
}

func TestLookupWitnessAccounts(t *testing.T) {
	resp, err := LookupWitnessAccounts(client, "steem", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnesses, resp)
}

func TestGetWitnessCount(t *testing.T) {
	resp, err := GetWitnessCount(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnessCount, resp)
}

func TestGetActiveVotes(t *testing.T) {
	resp, err := GetActiveVotes(client, "smooth", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnessCount, resp)
}

func TestGetAccountVotes(t *testing.T) {
	resp, err := GetAccountVotes(client, "smooth")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnessCount, resp)
}

func TestGetContent(t *testing.T) {
	resp, err := GetContent(client, "smooth", "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetContent, resp)
}

func TestGetFollowers(t *testing.T) {
	resp, err := GetFollowers(client, "good-karma", "fabien", "post", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetFollowers, resp)
}

func TestGetFollowCount(t *testing.T) {
	c1, c2, err := GetFollowCount(client, "good-karma")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %d, %d", MethodGetFollowCount, c1, c2)
}

func TestGetAccountReputations(t *testing.T) {
	resp, err := GetAccountReputations(client, "good-karma", 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetAccountReputations, resp)
}

func TestGetTicker(t *testing.T) {
	resp, err := GetTicker(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetTicker, resp)
}

func TestGetVolume(t *testing.T) {
	v1, v2, err := GetVolume(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %s, %s", MethodGetVolume, v1, v2)
}
