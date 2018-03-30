package api

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/smallnest/gosteem"
)

// refer to http://steem.esteem.ws/#/,
// http://steem.readthedocs.io/en/latest/steem.html,
// https://steemit.github.io/steemit-docs/?Javascript#introduction

var (
	MethodGetVersion = "get_version"

	MethodGetTrendingTags              = "get_trending_tags"
	MethodGetState                     = "get_state"
	MethodGetActiveWitnesses           = "get_active_witnesses"
	MethodGetBlockHeader               = "get_block_header"
	MethodGetBlock                     = "get_block"
	MethodGetOpsInBlock                = "get_ops_in_block"
	MethodGetConfig                    = "get_config"
	MethodGetDynamicGlobalProperties   = "get_dynamic_global_properties"
	MethodGetChainProperties           = "get_chain_properties"
	MethodGetCurrentMedianHistoryPrice = "get_current_median_history_price"
	MethodGetFeedHistory               = "get_feed_history"
	MethodGetWitnessSchedule           = "get_witness_schedule"
	MethodGetHardforkVersion           = "get_hardfork_version"
	MethodGetNextScheduledHardfork     = "get_next_scheduled_hardfork"
	MethodGetRewardFund                = "get_reward_fund"
	MethodGetKeyReferences             = "get_key_references"
	MethodGetAccounts                  = "get_accounts"
	// (get_account_references)
	MethodLookupAccountNames = "lookup_account_names"
	MethodLookupAccounts     = "lookup_accounts"

	MethodGetAccountCount = "get_account_count"

	MethodGetOwnerHistory = "get_owner_history"

	MethodGetRecoveryRequest            = "get_recovery_request"
	MethodGetEscrow                     = "get_escrow"
	MethodGetWithdrawRoutes             = "get_withdraw_routes"
	MethodGetAccountBandwidth           = "get_account_bandwidth"
	MethodGetSavingsWithdrawFrom        = "get_savings_withdraw_from"
	MethodGetSavingsWithdrawTo          = "get_savings_withdraw_to"
	MethodGetVestingDelegations         = "get_vesting_delegations"
	MethodGetExpiringVestingDelegations = "get_expiring_vesting_delegations"
	MethodGetWitnesses                  = "get_witnesses"
	MethodGetConversionRequests         = "get_conversion_requests"

	MethodGetWitnessByAccount           = "get_witness_by_account"
	MethodGetWitnessesByVote            = "get_witnesses_by_vote"
	MethodLookupWitnessAccounts         = "lookup_witness_accounts"
	MethodGetWitnessCount               = "get_witness_count"
	MethodGetOpenOrders                 = "get_open_orders"
	MethodGetTransactionHex             = "get_transaction_hex"
	MethodGetTransaction                = "get_transaction"
	MethodGetRequiredSignatures         = "get_required_signatures"
	MethodGetPotentialSignatures        = "get_potential_signatures"
	MethodVerifyAuthority               = "verify_authority"
	MethodVerifyAccountAuthority        = "verify_account_authority"
	MethodGetActiveVotes                = "get_active_votes"
	MethodGetAccountVotes               = "get_account_votes"
	MethodGetContent                    = "get_content"
	MethodGetContentReplies             = "get_content_replies"
	MethodGetTagsUsedByAuthor           = "get_tags_used_by_author"
	MethodGetPostDiscussionsByPayout    = "get_post_discussions_by_payout"
	MethodGetCommentDiscussionsByPayout = "get_comment_discussions_by_payout"

	// (get_discussions_by_trending)
	// (get_discussions_by_created)
	// (get_discussions_by_active)
	// (get_discussions_by_cashout)
	// (get_discussions_by_votes)
	// (get_discussions_by_children)
	// (get_discussions_by_hot)
	// (get_discussions_by_feed)
	// (get_discussions_by_blog)
	// (get_discussions_by_comments)
	// (get_discussions_by_promoted)
	// (get_replies_by_last_update)
	// (get_discussions_by_author_before_date)

	MethodGetAccountHistory = "get_account_history"

	MethodBroadcastTransaction            = "broadcast_transaction"
	MethodBroadcastTransactionSynchronous = "broadcast_transaction_synchronous"
	MethodBroadcastBlock                  = "broadcast_block"

	MethodGetFollowers   = "get_followers"
	MethodGetFollowing   = "get_following"
	MethodGetFollowCount = "get_follow_count"

	// (get_feed_entries)
	// (get_feed)

	// (get_blog_entries)
	// (get_blog)
	// (get_account_reputations)
	// (get_reblogged_by)
	// (get_blog_authors)

	// (get_ticker)
	// (get_volume)
	// (get_order_book)
	// (get_trade_history)
	// (get_recent_trades)
	// (get_market_history)
	// (get_market_history_buckets)
)

var tagMethods = map[string]string{
	"active":   "get_discussions_by_active",
	"blog":     "get_discussions_by_blog",
	"cashout":  "get_discussions_by_cashout",
	"children": "get_discussions_by_children",
	"comments": "get_discussions_by_comments",
	"created":  "get_discussions_by_created",
	"feed":     "get_discussions_by_feed",
	"hot":      "get_discussions_by_hot",
	"payout":   "get_discussions_by_payout",
	"promoted": "get_discussions_by_promoted",
	"trending": "get_discussions_by_trending",
	"update":   "get_discussions_by_update",
	"votes":    "get_discussions_by_votes",
}

type GetVersionResponse struct {
	BlockchainVersion string `json:"blockchain_version"`
	FcRevision        string `json:"fc_revision"`
	SteemRevision     string `json:"steem_revision"`
}

func GetVersion(client gosteem.Client) (*GetVersionResponse, error) {
	var resp = &GetVersionResponse{}
	err := client.Call(MethodGetVersion, nil, resp)
	return resp, err
}

type Tag struct {
	Comments     int64  `json:"comments"`
	Name         string `json:"name"`
	NetVotes     int64  `json:"net_votes"`
	TopPosts     int64  `json:"top_posts"`
	TotalPayouts string `json:"total_payouts"`
	Trending     string `json:"trending"`
}

func GetTrendingTags(client gosteem.Client, afterTag string, limit int) ([]Tag, error) {
	var resp []Tag
	err := client.Call(MethodGetTrendingTags, []interface{}{afterTag, limit}, &resp)
	return resp, err
}

type GetStateResponse struct {
	Accounts      struct{} `json:"accounts"`
	Content       struct{} `json:"content"`
	CurrentRoute  string   `json:"current_route"`
	DiscussionIdx struct{} `json:"discussion_idx"`
	Error         string   `json:"error"`
	FeedPrice     struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
	} `json:"feed_price"`
	Props struct {
		AverageBlockSize             int64  `json:"average_block_size"`
		ConfidentialSbdSupply        string `json:"confidential_sbd_supply"`
		ConfidentialSupply           string `json:"confidential_supply"`
		CurrentAslot                 int64  `json:"current_aslot"`
		CurrentReserveRatio          int64  `json:"current_reserve_ratio"`
		CurrentSbdSupply             string `json:"current_sbd_supply"`
		CurrentSupply                string `json:"current_supply"`
		CurrentWitness               string `json:"current_witness"`
		HeadBlockID                  string `json:"head_block_id"`
		HeadBlockNumber              int64  `json:"head_block_number"`
		LastIrreversibleBlockNum     int64  `json:"last_irreversible_block_num"`
		MaxVirtualBandwidth          string `json:"max_virtual_bandwidth"`
		MaximumBlockSize             int64  `json:"maximum_block_size"`
		NumPowWitnesses              int64  `json:"num_pow_witnesses"`
		ParticipationCount           int64  `json:"participation_count"`
		PendingRewardedVestingShares string `json:"pending_rewarded_vesting_shares"`
		PendingRewardedVestingSteem  string `json:"pending_rewarded_vesting_steem"`
		RecentSlotsFilled            string `json:"recent_slots_filled"`
		SbdInterestRate              int64  `json:"sbd_interest_rate"`
		SbdPrintRate                 int64  `json:"sbd_print_rate"`
		Time                         string `json:"time"`
		TotalPow                     int64  `json:"total_pow"`
		TotalRewardFundSteem         string `json:"total_reward_fund_steem"`
		TotalRewardShares2           string `json:"total_reward_shares2"`
		TotalVestingFundSteem        string `json:"total_vesting_fund_steem"`
		TotalVestingShares           string `json:"total_vesting_shares"`
		VirtualSupply                string `json:"virtual_supply"`
		VotePowerReserveRate         int64  `json:"vote_power_reserve_rate"`
	} `json:"props"`
	TagIdx struct {
		Trending []string `json:"trending"`
	} `json:"tag_idx"`
	Tags            struct{} `json:"tags"`
	WitnessSchedule struct {
		CurrentShuffledWitnesses  []string `json:"current_shuffled_witnesses"`
		CurrentVirtualTime        string   `json:"current_virtual_time"`
		HardforkRequiredWitnesses int64    `json:"hardfork_required_witnesses"`
		ID                        int64    `json:"id"`
		MajorityVersion           string   `json:"majority_version"`
		MaxMinerWitnesses         int64    `json:"max_miner_witnesses"`
		MaxRunnerWitnesses        int64    `json:"max_runner_witnesses"`
		MaxVotedWitnesses         int64    `json:"max_voted_witnesses"`
		MedianProps               struct {
			AccountCreationFee string `json:"account_creation_fee"`
			MaximumBlockSize   int64  `json:"maximum_block_size"`
			SbdInterestRate    int64  `json:"sbd_interest_rate"`
		} `json:"median_props"`
		MinerWeight                   int64 `json:"miner_weight"`
		NextShuffleBlockNum           int64 `json:"next_shuffle_block_num"`
		NumScheduledWitnesses         int64 `json:"num_scheduled_witnesses"`
		TimeshareWeight               int64 `json:"timeshare_weight"`
		Top19Weight                   int64 `json:"top19_weight"`
		WitnessPayNormalizationFactor int64 `json:"witness_pay_normalization_factor"`
	} `json:"witness_schedule"`
	Witnesses struct{} `json:"witnesses"`
}

func GetState(client gosteem.Client, account string) (*GetStateResponse, error) {
	var resp = &GetStateResponse{}
	err := client.Call(MethodGetState, []interface{}{account}, resp)
	return resp, err
}

type BlockHeader struct {
	Extensions            []interface{} `json:"extensions"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Witness               string        `json:"witness"`
}

func GetBlockHeader(client gosteem.Client, blockID uint64) (*BlockHeader, error) {
	var resp = &BlockHeader{}
	err := client.Call(MethodGetBlockHeader, []interface{}{blockID}, resp)
	return resp, err
}

type Block struct {
	BlockID               string        `json:"block_id"`
	Extensions            []interface{} `json:"extensions"`
	Previous              string        `json:"previous"`
	SigningKey            string        `json:"signing_key"`
	Timestamp             string        `json:"timestamp"`
	TransactionIds        []string      `json:"transaction_ids"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Transactions          []struct {
		BlockNum       int64           `json:"block_num"`
		Expiration     string          `json:"expiration"`
		Extensions     []interface{}   `json:"extensions"`
		Operations     [][]interface{} `json:"operations"`
		RefBlockNum    int64           `json:"ref_block_num"`
		RefBlockPrefix int64           `json:"ref_block_prefix"`
		Signatures     []string        `json:"signatures"`
		TransactionID  string          `json:"transaction_id"`
		TransactionNum int64           `json:"transaction_num"`
	} `json:"transactions"`
	Witness          string `json:"witness"`
	WitnessSignature string `json:"witness_signature"`
}

func GetBlock(client gosteem.Client, blockID uint64) (*Block, error) {
	var resp = &Block{}
	err := client.Call(MethodGetBlock, []interface{}{blockID}, resp)
	return resp, err
}

type BlockOps struct {
	Block      int64         `json:"block"`
	Op         []interface{} `json:"op"`
	OpInTrx    int64         `json:"op_in_trx"`
	Timestamp  string        `json:"timestamp"`
	TrxID      string        `json:"trx_id"`
	TrxInBlock int64         `json:"trx_in_block"`
	VirtualOp  int64         `json:"virtual_op"`
}

func GetOpsInBlock(client gosteem.Client, blockID uint64, onlyVirtual bool) ([]BlockOps, error) {
	var resp = []BlockOps{}
	err := client.Call(MethodGetOpsInBlock, []interface{}{blockID, onlyVirtual}, &resp)
	return resp, err
}

type GetChainPropertiesResponse struct {
	AccountCreationFee string `json:"account_creation_fee"`
	MaximumBlockSize   int64  `json:"maximum_block_size"`
	SbdInterestRate    int64  `json:"sbd_interest_rate"`
}

func GetChainProperties(client gosteem.Client) (*GetChainPropertiesResponse, error) {
	var resp = &GetChainPropertiesResponse{}
	err := client.Call(MethodGetChainProperties, nil, &resp)
	return resp, err
}

func GetCurrentMedianHistoryPrice(client gosteem.Client) (base, quote string, err error) {
	var resp = make(map[string]string)
	err = client.Call(MethodGetCurrentMedianHistoryPrice, nil, &resp)
	return resp["base"], resp["quote"], err
}

func GetHardforkVersion(client gosteem.Client) (version string, err error) {
	err = client.Call(MethodGetHardforkVersion, nil, &version)
	return version, err
}

func GetNextScheduledHardfork(client gosteem.Client) (hfVersion string, liveTime time.Time, err error) {
	var resp = make(map[string]string)
	err = client.Call(MethodGetNextScheduledHardfork, nil, &resp)

	t, err := time.Parse("2006-01-02T15:04:05", resp["live_time"])
	if err != nil {
		return "", time.Time{}, err
	}
	return resp["hf_version"], t, err
}

type GetRewardFundResponse struct {
	AuthorRewardCurve      string `json:"author_reward_curve"`
	ContentConstant        string `json:"content_constant"`
	CurationRewardCurve    string `json:"curation_reward_curve"`
	ID                     int64  `json:"id"`
	LastUpdate             string `json:"last_update"`
	Name                   string `json:"name"`
	PercentContentRewards  int64  `json:"percent_content_rewards"`
	PercentCurationRewards int64  `json:"percent_curation_rewards"`
	RecentClaims           string `json:"recent_claims"`
	RewardBalance          string `json:"reward_balance"`
}

func GetRewardFund(client gosteem.Client, fundName string) (*GetRewardFundResponse, error) {
	var resp = &GetRewardFundResponse{}
	err := client.Call(MethodGetRewardFund, []interface{}{fundName}, &resp)
	return resp, err
}

func GetAccountCount(client gosteem.Client) (uint64, error) {
	var count uint64
	err := client.Call(MethodGetAccountCount, nil, &count)
	return count, err
}

func GetOwnerHistory(client gosteem.Client, account string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetOwnerHistory, []interface{}{account}, &resp)
	return resp, err
}

func GetKeyReferences(client gosteem.Client, keys ...string) ([]string, error) {
	var resp []string
	err := client.Call(MethodGetKeyReferences, []interface{}{keys}, &resp)

	return resp, err
}

type Account struct {
	Active struct {
		AccountAuths    []interface{}   `json:"account_auths"`
		KeyAuths        [][]interface{} `json:"key_auths"`
		WeightThreshold int64           `json:"weight_threshold"`
	} `json:"active"`
	AverageBandwidth          interface{}   `json:"average_bandwidth"`
	AverageMarketBandwidth    int64         `json:"average_market_bandwidth"`
	Balance                   string        `json:"balance"`
	CanVote                   bool          `json:"can_vote"`
	CommentCount              int64         `json:"comment_count"`
	Created                   string        `json:"created"`
	CurationRewards           int64         `json:"curation_rewards"`
	DelegatedVestingShares    string        `json:"delegated_vesting_shares"`
	GuestBloggers             []interface{} `json:"guest_bloggers"`
	ID                        int64         `json:"id"`
	JSONMetadata              string        `json:"json_metadata"`
	LastAccountRecovery       string        `json:"last_account_recovery"`
	LastAccountUpdate         string        `json:"last_account_update"`
	LastBandwidthUpdate       string        `json:"last_bandwidth_update"`
	LastMarketBandwidthUpdate string        `json:"last_market_bandwidth_update"`
	LastOwnerUpdate           string        `json:"last_owner_update"`
	LastPost                  string        `json:"last_post"`
	LastRootPost              string        `json:"last_root_post"`
	LastVoteTime              string        `json:"last_vote_time"`
	LifetimeBandwidth         string        `json:"lifetime_bandwidth"`
	LifetimeMarketBandwidth   interface{}   `json:"lifetime_market_bandwidth"`
	LifetimeVoteCount         int64         `json:"lifetime_vote_count"`
	MarketHistory             []interface{} `json:"market_history"`
	MemoKey                   string        `json:"memo_key"`
	Mined                     bool          `json:"mined"`
	Name                      string        `json:"name"`
	NextVestingWithdrawal     string        `json:"next_vesting_withdrawal"`
	OtherHistory              []interface{} `json:"other_history"`
	Owner                     struct {
		AccountAuths    []interface{}   `json:"account_auths"`
		KeyAuths        [][]interface{} `json:"key_auths"`
		WeightThreshold int64           `json:"weight_threshold"`
	} `json:"owner"`
	PostCount   int64         `json:"post_count"`
	PostHistory []interface{} `json:"post_history"`
	Posting     struct {
		AccountAuths    []interface{}   `json:"account_auths"`
		KeyAuths        [][]interface{} `json:"key_auths"`
		WeightThreshold int64           `json:"weight_threshold"`
	} `json:"posting"`
	PostingRewards                int64         `json:"posting_rewards"`
	ProxiedVsfVotes               []interface{} `json:"proxied_vsf_votes"` //maybe mix string or int64
	Proxy                         string        `json:"proxy"`
	ReceivedVestingShares         string        `json:"received_vesting_shares"`
	RecoveryAccount               string        `json:"recovery_account"`
	Reputation                    interface{}   `json:"reputation"`
	ResetAccount                  string        `json:"reset_account"`
	RewardSbdBalance              string        `json:"reward_sbd_balance"`
	RewardSteemBalance            string        `json:"reward_steem_balance"`
	RewardVestingBalance          string        `json:"reward_vesting_balance"`
	RewardVestingSteem            string        `json:"reward_vesting_steem"`
	SavingsBalance                string        `json:"savings_balance"`
	SavingsSbdBalance             string        `json:"savings_sbd_balance"`
	SavingsSbdLastInterestPayment string        `json:"savings_sbd_last_interest_payment"`
	SavingsSbdSeconds             string        `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   string        `json:"savings_sbd_seconds_last_update"`
	SavingsWithdrawRequests       int64         `json:"savings_withdraw_requests"`
	SbdBalance                    string        `json:"sbd_balance"`
	SbdLastInterestPayment        string        `json:"sbd_last_interest_payment"`
	SbdSeconds                    string        `json:"sbd_seconds"`
	SbdSecondsLastUpdate          string        `json:"sbd_seconds_last_update"`
	TagsUsage                     []interface{} `json:"tags_usage"`
	ToWithdraw                    interface{}   `json:"to_withdraw"` //maybe string or int64
	TransferHistory               []interface{} `json:"transfer_history"`
	VestingBalance                string        `json:"vesting_balance"`
	VestingShares                 string        `json:"vesting_shares"`
	VestingWithdrawRate           string        `json:"vesting_withdraw_rate"`
	VoteHistory                   []interface{} `json:"vote_history"`
	VotingPower                   int64         `json:"voting_power"`
	WithdrawRoutes                int64         `json:"withdraw_routes"`
	Withdrawn                     interface{}   `json:"withdrawn"` ////maybe string or int64
	WitnessVotes                  []interface{} `json:"witness_votes"`
	WitnessesVotedFor             int64         `json:"witnesses_voted_for"`
}

func GetAccounts(client gosteem.Client, accounts ...string) ([]Account, error) {
	var resp []Account
	err := client.Call(MethodGetAccounts, []interface{}{accounts}, &resp)

	return resp, err
}

func LookupAccountNames(client gosteem.Client, accounts ...string) ([]Account, error) {
	var resp []Account
	err := client.Call(MethodLookupAccountNames, []interface{}{accounts}, &resp)

	return resp, err
}

func LookupAccounts(client gosteem.Client, lowerBoundName string, limit int) ([]string, error) {
	var resp []string
	err := client.Call(MethodLookupAccounts, []interface{}{lowerBoundName, limit}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetRecoveryRequest(client gosteem.Client, account string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetRecoveryRequest, []interface{}{account}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetEscrow(client gosteem.Client, account string, escrowID int) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetEscrow, []interface{}{account, escrowID}, &resp)

	return resp, err
}

type WithdrawRoute struct {
	AutoVest    bool   `json:"auto_vest"`
	FromAccount string `json:"from_account"`
	ID          int64  `json:"id"`
	Percent     int64  `json:"percent"`
	ToAccount   string `json:"to_account"`
}

// withdrawRouteType is incoming,outgoing,all
func GetWithdrawRoutes(client gosteem.Client, account string, withdrawRouteType string) ([]WithdrawRoute, error) {
	var resp []WithdrawRoute
	err := client.Call(MethodGetWithdrawRoutes, []interface{}{account, withdrawRouteType}, &resp)

	return resp, err
}

type GetAccountBandwidthResponse struct {
	Account             string `json:"account"`
	AverageBandwidth    string `json:"average_bandwidth"`
	ID                  int64  `json:"id"`
	LastBandwidthUpdate string `json:"last_bandwidth_update"`
	LifetimeBandwidth   string `json:"lifetime_bandwidth"`
	Type                string `json:"type"`
}

// post, forum, market
func GetAccountBandwidth(client gosteem.Client, account string, bandwithType string) (*GetAccountBandwidthResponse, error) {
	var resp = &GetAccountBandwidthResponse{}
	err := client.Call(MethodGetAccountBandwidth, []interface{}{account, bandwithType}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetSavingsWithdrawFrom(client gosteem.Client, account string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetSavingsWithdrawFrom, []interface{}{account}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetSavingsWithdrawTo(client gosteem.Client, account string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetSavingsWithdrawTo, []interface{}{account}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetVestingDelegations(client gosteem.Client, account, fromAccount string, limit int) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetVestingDelegations, []interface{}{account, fromAccount, limit}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetExpiringVestingDelegations(client gosteem.Client, account, start string, limit int) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetExpiringVestingDelegations, []interface{}{account, start, limit}, &resp)

	return resp, err
}

type Witnesses struct {
	Created               string `json:"created"`
	HardforkTimeVote      string `json:"hardfork_time_vote"`
	HardforkVersionVote   string `json:"hardfork_version_vote"`
	ID                    int64  `json:"id"`
	LastAslot             int64  `json:"last_aslot"`
	LastConfirmedBlockNum int64  `json:"last_confirmed_block_num"`
	LastSbdExchangeUpdate string `json:"last_sbd_exchange_update"`
	LastWork              string `json:"last_work"`
	Owner                 string `json:"owner"`
	PowWorker             int64  `json:"pow_worker"`
	Props                 struct {
		AccountCreationFee string `json:"account_creation_fee"`
		MaximumBlockSize   int64  `json:"maximum_block_size"`
		SbdInterestRate    int64  `json:"sbd_interest_rate"`
	} `json:"props"`
	RunningVersion  string `json:"running_version"`
	SbdExchangeRate struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
	} `json:"sbd_exchange_rate"`
	SigningKey           string `json:"signing_key"`
	TotalMissed          int64  `json:"total_missed"`
	URL                  string `json:"url"`
	VirtualLastUpdate    string `json:"virtual_last_update"`
	VirtualPosition      string `json:"virtual_position"`
	VirtualScheduledTime string `json:"virtual_scheduled_time"`
	Votes                string `json:"votes"`
}

func GetWitnesses(client gosteem.Client, witnesses ...string) ([]Witnesses, error) {
	var resp []Witnesses
	err := client.Call(MethodGetWitnesses, []interface{}{witnesses}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetConversionRequests(client gosteem.Client, account string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetConversionRequests, []interface{}{account}, &resp)

	return resp, err
}

func GetWitnessByAccount(client gosteem.Client, account string) (*Witnesses, error) {
	var resp Witnesses
	err := client.Call(MethodGetWitnessByAccount, []interface{}{account}, &resp)

	return &resp, err
}

func GetWitnessesByVote(client gosteem.Client, fromWitness string, limit int) ([]Witnesses, error) {
	var resp []Witnesses
	err := client.Call(MethodGetWitnessesByVote, []interface{}{fromWitness, limit}, &resp)

	return resp, err
}

func LookupWitnessAccounts(client gosteem.Client, lowerBoundName string, limit int) ([]string, error) {
	var resp []string
	err := client.Call(MethodLookupWitnessAccounts, []interface{}{lowerBoundName, limit}, &resp)

	return resp, err
}

func GetWitnessCount(client gosteem.Client) (uint64, error) {
	var resp uint64
	err := client.Call(MethodGetWitnessCount, nil, &resp)

	return resp, err
}

// TODO: don't know response format
func GetOpenOrders(client gosteem.Client, owner string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetOpenOrders, []interface{}{owner}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetTransactionHex(client gosteem.Client, trx string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetTransactionHex, []interface{}{trx}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetTransaction(client gosteem.Client, trxID string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetTransaction, []interface{}{trxID}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetRequiredSignatures(client gosteem.Client, trx string, availableKeys ...string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetRequiredSignatures, []interface{}{trx, availableKeys}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetPotentialSignatures(client gosteem.Client, trx string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetPotentialSignatures, []interface{}{trx}, &resp)

	return resp, err
}

// TODO: don't know response format
func VerifyAuthority(client gosteem.Client, trx string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodVerifyAuthority, []interface{}{trx}, &resp)

	return resp, err
}

// TODO: don't know response format
func VerifyAccountAuthority(client gosteem.Client, nameOrID string, signers ...string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodVerifyAccountAuthority, []interface{}{nameOrID, signers}, &resp)

	return resp, err
}

type ActiveVote struct {
	Percent    int64       `json:"percent"`
	Reputation interface{} `json:"reputation"` //string or 0
	Rshares    interface{} `json:"rshares"`    //string or 0
	Time       string      `json:"time"`
	Voter      string      `json:"voter"`
	Weight     interface{} `json:"weight"` //string or 0
}

func GetActiveVotes(client gosteem.Client, author string, permlink string) ([]ActiveVote, error) {
	var resp []ActiveVote
	err := client.Call(MethodGetActiveVotes, []interface{}{author, permlink}, &resp)

	return resp, err
}

type Vote struct {
	Authorperm string `json:"authorperm"`
	Percent    int64  `json:"percent"`
	Rshares    string `json:"rshares"`
	Time       string `json:"time"`
	Weight     int64  `json:"weight"`
}

func GetAccountVotes(client gosteem.Client, voter string) ([]Vote, error) {
	var resp []Vote
	err := client.Call(MethodGetAccountVotes, []interface{}{voter}, &resp)

	return resp, err
}

type Discussion struct {
	AbsRshares  interface{} `json:"abs_rshares"`
	Active      string      `json:"active"`
	ActiveVotes []struct {
		Percent    int64  `json:"percent"`
		Reputation int64  `json:"reputation"`
		Rshares    int64  `json:"rshares"`
		Time       string `json:"time"`
		Voter      string `json:"voter"`
		Weight     int64  `json:"weight"`
	} `json:"active_votes"`
	AllowCurationRewards bool   `json:"allow_curation_rewards"`
	AllowReplies         bool   `json:"allow_replies"`
	AllowVotes           bool   `json:"allow_votes"`
	Author               string `json:"author"`
	AuthorReputation     string `json:"author_reputation"`
	AuthorRewards        int64  `json:"author_rewards"`
	Beneficiaries        []struct {
		Account string `json:"account"`
		Weight  int64  `json:"weight"`
	} `json:"beneficiaries"`
	Body                    string        `json:"body"`
	BodyLength              int64         `json:"body_length"`
	CashoutTime             string        `json:"cashout_time"`
	Category                string        `json:"category"`
	Children                int64         `json:"children"`
	ChildrenAbsRshares      interface{}   `json:"children_abs_rshares"`
	Created                 string        `json:"created"`
	CuratorPayoutValue      string        `json:"curator_payout_value"`
	Depth                   int64         `json:"depth"`
	ID                      int64         `json:"id"`
	JSONMetadata            string        `json:"json_metadata"`
	LastPayout              string        `json:"last_payout"`
	LastUpdate              string        `json:"last_update"`
	MaxAcceptedPayout       string        `json:"max_accepted_payout"`
	MaxCashoutTime          string        `json:"max_cashout_time"`
	NetRshares              interface{}   `json:"net_rshares"`
	NetVotes                int64         `json:"net_votes"`
	ParentAuthor            string        `json:"parent_author"`
	ParentPermlink          string        `json:"parent_permlink"`
	PendingPayoutValue      string        `json:"pending_payout_value"`
	PercentSteemDollars     int64         `json:"percent_steem_dollars"`
	Permlink                string        `json:"permlink"`
	Promoted                string        `json:"promoted"`
	RebloggedBy             []interface{} `json:"reblogged_by"`
	Replies                 []interface{} `json:"replies"`
	RewardWeight            int64         `json:"reward_weight"`
	RootAuthor              string        `json:"root_author"`
	RootPermlink            string        `json:"root_permlink"`
	RootTitle               string        `json:"root_title"`
	Title                   string        `json:"title"`
	TotalPayoutValue        string        `json:"total_payout_value"`
	TotalPendingPayoutValue string        `json:"total_pending_payout_value"`
	TotalVoteWeight         int64         `json:"total_vote_weight"`
	URL                     string        `json:"url"`
	VoteRshares             interface{}   `json:"vote_rshares"`
}

// method is active、blog、cashout、children、comments、created、feed、hot、payout、promoted、trending、update、votes
func GetDiscussionsBy(client gosteem.Client, method string, params string) ([]Discussion, error) {
	var resp []Discussion

	if !strings.HasPrefix(method, "get_discussions_by_") {
		method = "get_discussions_by_" + method
	}

	err := client.Call(method, []interface{}{params}, &resp)

	return resp, err
}

type Content struct {
	AbsRshares              int64            `json:"abs_rshares"`
	Active                  string           `json:"active"`
	ActiveVotes             []ActiveVote     `json:"active_votes"`
	AllowCurationRewards    bool             `json:"allow_curation_rewards"`
	AllowReplies            bool             `json:"allow_replies"`
	AllowVotes              bool             `json:"allow_votes"`
	Author                  string           `json:"author"`
	AuthorReputation        string           `json:"author_reputation"`
	AuthorRewards           int64            `json:"author_rewards"`
	Beneficiaries           *ArrayRawMessage `json:"beneficiaries"`
	Body                    string           `json:"body"`
	BodyLength              int64            `json:"body_length"`
	CashoutTime             string           `json:"cashout_time"`
	Category                string           `json:"category"`
	Children                int64            `json:"children"`
	ChildrenAbsRshares      int64            `json:"children_abs_rshares"`
	Created                 string           `json:"created"`
	CuratorPayoutValue      string           `json:"curator_payout_value"`
	Depth                   int64            `json:"depth"`
	ID                      int64            `json:"id"`
	JSONMetadata            string           `json:"json_metadata"`
	LastPayout              string           `json:"last_payout"`
	LastUpdate              string           `json:"last_update"`
	MaxAcceptedPayout       string           `json:"max_accepted_payout"`
	MaxCashoutTime          string           `json:"max_cashout_time"`
	NetRshares              int64            `json:"net_rshares"`
	NetVotes                int64            `json:"net_votes"`
	ParentAuthor            string           `json:"parent_author"`
	ParentPermlink          string           `json:"parent_permlink"`
	PendingPayoutValue      string           `json:"pending_payout_value"`
	PercentSteemDollars     int64            `json:"percent_steem_dollars"`
	Permlink                string           `json:"permlink"`
	Promoted                string           `json:"promoted"`
	RebloggedBy             *ArrayRawMessage `json:"reblogged_by"`
	Replies                 *ArrayRawMessage `json:"replies"`
	RewardWeight            int64            `json:"reward_weight"`
	RootAuthor              string           `json:"root_author"`
	RootPermlink            string           `json:"root_permlink"`
	RootTitle               string           `json:"root_title"`
	Title                   string           `json:"title"`
	TotalPayoutValue        string           `json:"total_payout_value"`
	TotalPendingPayoutValue string           `json:"total_pending_payout_value"`
	TotalVoteWeight         int64            `json:"total_vote_weight"`
	URL                     string           `json:"url"`
	VoteRshares             int64            `json:"vote_rshares"`
}

func GetContent(client gosteem.Client, author, permlink string) (*Content, error) {
	var resp = &Content{}
	err := client.Call(MethodGetContent, []interface{}{author, permlink}, resp)

	return resp, err
}

// TODO: don't know response format
func GetContentReplies(client gosteem.Client, author, permlink string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetContentReplies, []interface{}{author, permlink}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetTagsUsedByAuthor(client gosteem.Client, author string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetTagsUsedByAuthor, []interface{}{author}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetPostDiscussionsByPayout(client gosteem.Client, query string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetPostDiscussionsByPayout, []interface{}{query}, &resp)

	return resp, err
}

// TODO: don't know response format
func GetCommentDiscussionsByPayout(client gosteem.Client, query string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetCommentDiscussionsByPayout, []interface{}{query}, &resp)

	return resp, err
}

// TODO: don't know response format
func BroadcastTransaction(client gosteem.Client, trx string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodBroadcastTransaction, []interface{}{trx}, &resp)

	return resp, err
}

// TODO: don't know response format
func BroadcastTransactionSynchronous(client gosteem.Client, trx string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodBroadcastTransactionSynchronous, []interface{}{trx}, &resp)

	return resp, err
}

// TODO: don't know response format
func BroadcastBlock(client gosteem.Client, b string) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodBroadcastBlock, []interface{}{b}, &resp)

	return resp, err
}

type Follow struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

func GetFollowers(client gosteem.Client, following, startFollower, followType string, limit int) ([]Follow, error) {
	var resp []Follow
	err := client.Call(MethodGetFollowers, []interface{}{following, startFollower, followType, limit}, &resp)

	return resp, err
}

func GetFollowing(client gosteem.Client, following, startFollower, followType string, limit int) ([]Follow, error) {
	var resp []Follow
	err := client.Call(MethodGetFollowing, []interface{}{following, startFollower, followType, limit}, &resp)

	return resp, err
}

type FollowCount struct {
	Account        string `json:"account"`
	FollowerCount  uint64 `json:"follower_count"`
	FollowingCount uint64 `json:"following_count"`
}

func GetFollowCount(client gosteem.Client, account string) (followerCount uint64, followingCount uint64, err error) {
	var resp FollowCount
	err = client.Call(MethodGetFollowCount, []interface{}{account}, &resp)

	return resp.FollowerCount, resp.FollowingCount, err
}

// (get_feed_entries)
// (get_feed)

// (get_blog_entries)
// (get_blog)
// (get_account_reputations)
// (get_reblogged_by)
// (get_blog_authors)

// (get_ticker)
// (get_volume)
// (get_order_book)
// (get_trade_history)
// (get_recent_trades)
// (get_market_history)
// (get_market_history_buckets)
