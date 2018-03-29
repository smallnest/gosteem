package api

import "github.com/smallnest/gosteem"

// from https://github.com/steemit/steem/blob/master/libraries/plugins/apis/database_api/include/steem/plugins/database_api/database_api.hpp
// curl --silent https://api.steemit.com --data '{"jsonrpc":"2.0","id":0,"method":"get_config","params":[]}' | python -m json.tool |gojson
const (
// MethodGetConfig = "get_config"

// MethodGetDynamicGlobalProperties = "get_dynamic_global_properties"

// MethodGetWitnessSchedule = "get_witness_schedule"

// MethodGetFeedHistory = "get_feed_history"

// MethodGetActiveWitnesses = "get_active_witnesses"
)

type GetConfigResponse struct {
	IsTestNet                                   bool          `json:"IS_TEST_NET"`
	SbdSymbol                                   interface{}   `json:"SBD_SYMBOL"`
	STEEM100PERCENT                             int64         `json:"STEEM_100_PERCENT"`
	STEEM1PERCENT                               int64         `json:"STEEM_1_PERCENT"`
	SteemAccountRecoveryRequestExpirationPeriod string        `json:"STEEM_ACCOUNT_RECOVERY_REQUEST_EXPIRATION_PERIOD"`
	SteemActiveChallengeCooldown                string        `json:"STEEM_ACTIVE_CHALLENGE_COOLDOWN"`
	SteemActiveChallengeFee                     []interface{} `json:"STEEM_ACTIVE_CHALLENGE_FEE"`
	SteemAddressPrefix                          string        `json:"STEEM_ADDRESS_PREFIX"`
	SteemAprPercentMultiplyPerBlock             string        `json:"STEEM_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	SteemAprPercentMultiplyPerHour              string        `json:"STEEM_APR_PERCENT_MULTIPLY_PER_HOUR"`
	SteemAprPercentMultiplyPerRound             string        `json:"STEEM_APR_PERCENT_MULTIPLY_PER_ROUND"`
	SteemAprPercentShiftPerBlock                int64         `json:"STEEM_APR_PERCENT_SHIFT_PER_BLOCK"`
	SteemAprPercentShiftPerHour                 int64         `json:"STEEM_APR_PERCENT_SHIFT_PER_HOUR"`
	SteemAprPercentShiftPerRound                int64         `json:"STEEM_APR_PERCENT_SHIFT_PER_ROUND"`
	SteemBandwidthAverageWindowSeconds          int64         `json:"STEEM_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	SteemBandwidthPrecision                     int64         `json:"STEEM_BANDWIDTH_PRECISION"`
	SteemBlockchainHardforkVersion              string        `json:"STEEM_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemBlockchainPrecision                    int64         `json:"STEEM_BLOCKCHAIN_PRECISION"`
	SteemBlockchainPrecisionDigits              int64         `json:"STEEM_BLOCKCHAIN_PRECISION_DIGITS"`
	SteemBlockchainVersion                      string        `json:"STEEM_BLOCKCHAIN_VERSION"`
	SteemBlocksPerDay                           int64         `json:"STEEM_BLOCKS_PER_DAY"`
	SteemBlocksPerHour                          int64         `json:"STEEM_BLOCKS_PER_HOUR"`
	SteemBlocksPerYear                          int64         `json:"STEEM_BLOCKS_PER_YEAR"`
	SteemBlockInterval                          int64         `json:"STEEM_BLOCK_INTERVAL"`
	SteemCashoutWindowSeconds                   int64         `json:"STEEM_CASHOUT_WINDOW_SECONDS"`
	STEEMCASHOUTWINDOWSECONDSPREHF12            int64         `json:"STEEM_CASHOUT_WINDOW_SECONDS_PRE_HF12"`
	STEEMCASHOUTWINDOWSECONDSPREHF17            int64         `json:"STEEM_CASHOUT_WINDOW_SECONDS_PRE_HF17"`
	SteemChainID                                string        `json:"STEEM_CHAIN_ID"`
	SteemChainIDName                            string        `json:"STEEM_CHAIN_ID_NAME"`
	SteemCommentRewardFundName                  string        `json:"STEEM_COMMENT_REWARD_FUND_NAME"`
	SteemContentAprPercent                      int64         `json:"STEEM_CONTENT_APR_PERCENT"`
	STEEMCONTENTCONSTANTHF0                     string        `json:"STEEM_CONTENT_CONSTANT_HF0"`
	SteemContentRewardPercent                   int64         `json:"STEEM_CONTENT_REWARD_PERCENT"`
	SteemConversionDelay                        string        `json:"STEEM_CONVERSION_DELAY"`
	STEEMCONVERSIONDELAYPREHF16                 string        `json:"STEEM_CONVERSION_DELAY_PRE_HF_16"`
	SteemCreateAccountDelegationRatio           int64         `json:"STEEM_CREATE_ACCOUNT_DELEGATION_RATIO"`
	SteemCreateAccountDelegationTime            string        `json:"STEEM_CREATE_ACCOUNT_DELEGATION_TIME"`
	SteemCreateAccountWithSteemModifier         int64         `json:"STEEM_CREATE_ACCOUNT_WITH_STEEM_MODIFIER"`
	SteemCurateAprPercent                       int64         `json:"STEEM_CURATE_APR_PERCENT"`
	SteemDefaultSbdInterestRate                 int64         `json:"STEEM_DEFAULT_SBD_INTEREST_RATE"`
	SteemEnableSmt                              bool          `json:"STEEM_ENABLE_SMT"`
	SteemEquihashK                              int64         `json:"STEEM_EQUIHASH_K"`
	SteemEquihashN                              int64         `json:"STEEM_EQUIHASH_N"`
	SteemFeedHistoryWindow                      int64         `json:"STEEM_FEED_HISTORY_WINDOW"`
	STEEMFEEDHISTORYWINDOWPREHF16               int64         `json:"STEEM_FEED_HISTORY_WINDOW_PRE_HF_16"`
	SteemFeedIntervalBlocks                     int64         `json:"STEEM_FEED_INTERVAL_BLOCKS"`
	SteemGenesisTime                            string        `json:"STEEM_GENESIS_TIME"`
	SteemHardforkRequiredWitnesses              int64         `json:"STEEM_HARDFORK_REQUIRED_WITNESSES"`
	SteemInflationNarrowingPeriod               int64         `json:"STEEM_INFLATION_NARROWING_PERIOD"`
	SteemInflationRateStartPercent              int64         `json:"STEEM_INFLATION_RATE_START_PERCENT"`
	SteemInflationRateStopPercent               int64         `json:"STEEM_INFLATION_RATE_STOP_PERCENT"`
	SteemInitialVotePowerRate                   int64         `json:"STEEM_INITIAL_VOTE_POWER_RATE"`
	SteemInitMinerName                          string        `json:"STEEM_INIT_MINER_NAME"`
	SteemInitPublicKeyStr                       string        `json:"STEEM_INIT_PUBLIC_KEY_STR"`
	SteemInitSupply                             int64         `json:"STEEM_INIT_SUPPLY"`
	SteemInitTime                               string        `json:"STEEM_INIT_TIME"`
	SteemIrreversibleThreshold                  int64         `json:"STEEM_IRREVERSIBLE_THRESHOLD"`
	SteemLiquidityAprPercent                    int64         `json:"STEEM_LIQUIDITY_APR_PERCENT"`
	SteemLiquidityRewardBlocks                  int64         `json:"STEEM_LIQUIDITY_REWARD_BLOCKS"`
	SteemLiquidityRewardPeriodSec               int64         `json:"STEEM_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemLiquidityTimeoutSec                    string        `json:"STEEM_LIQUIDITY_TIMEOUT_SEC"`
	SteemMaxAccountNameLength                   int64         `json:"STEEM_MAX_ACCOUNT_NAME_LENGTH"`
	SteemMaxAccountWitnessVotes                 int64         `json:"STEEM_MAX_ACCOUNT_WITNESS_VOTES"`
	SteemMaxAssetWhitelistAuthorities           int64         `json:"STEEM_MAX_ASSET_WHITELIST_AUTHORITIES"`
	SteemMaxAuthorityMembership                 int64         `json:"STEEM_MAX_AUTHORITY_MEMBERSHIP"`
	SteemMaxBlockSize                           int64         `json:"STEEM_MAX_BLOCK_SIZE"`
	SteemMaxCashoutWindowSeconds                int64         `json:"STEEM_MAX_CASHOUT_WINDOW_SECONDS"`
	SteemMaxCommentDepth                        int64         `json:"STEEM_MAX_COMMENT_DEPTH"`
	STEEMMAXCOMMENTDEPTHPREHF17                 int64         `json:"STEEM_MAX_COMMENT_DEPTH_PRE_HF17"`
	SteemMaxFeedAgeSeconds                      int64         `json:"STEEM_MAX_FEED_AGE_SECONDS"`
	SteemMaxInstanceID                          string        `json:"STEEM_MAX_INSTANCE_ID"`
	SteemMaxMemoSize                            int64         `json:"STEEM_MAX_MEMO_SIZE"`
	STEEMMAXMINERWITNESSESHF0                   int64         `json:"STEEM_MAX_MINER_WITNESSES_HF0"`
	STEEMMAXMINERWITNESSESHF17                  int64         `json:"STEEM_MAX_MINER_WITNESSES_HF17"`
	SteemMaxPermlinkLength                      int64         `json:"STEEM_MAX_PERMLINK_LENGTH"`
	SteemMaxProxyRecursionDepth                 int64         `json:"STEEM_MAX_PROXY_RECURSION_DEPTH"`
	SteemMaxRationDecayRate                     int64         `json:"STEEM_MAX_RATION_DECAY_RATE"`
	SteemMaxReserveRatio                        int64         `json:"STEEM_MAX_RESERVE_RATIO"`
	STEEMMAXRUNNERWITNESSESHF0                  int64         `json:"STEEM_MAX_RUNNER_WITNESSES_HF0"`
	STEEMMAXRUNNERWITNESSESHF17                 int64         `json:"STEEM_MAX_RUNNER_WITNESSES_HF17"`
	SteemMaxSatoshis                            string        `json:"STEEM_MAX_SATOSHIS"`
	SteemMaxShareSupply                         string        `json:"STEEM_MAX_SHARE_SUPPLY"`
	SteemMaxSigCheckDepth                       int64         `json:"STEEM_MAX_SIG_CHECK_DEPTH"`
	SteemMaxTimeUntilExpiration                 int64         `json:"STEEM_MAX_TIME_UNTIL_EXPIRATION"`
	SteemMaxTransactionSize                     int64         `json:"STEEM_MAX_TRANSACTION_SIZE"`
	SteemMaxUndoHistory                         int64         `json:"STEEM_MAX_UNDO_HISTORY"`
	SteemMaxURLLength                           int64         `json:"STEEM_MAX_URL_LENGTH"`
	STEEMMAXVOTEDWITNESSESHF0                   int64         `json:"STEEM_MAX_VOTED_WITNESSES_HF0"`
	STEEMMAXVOTEDWITNESSESHF17                  int64         `json:"STEEM_MAX_VOTED_WITNESSES_HF17"`
	SteemMaxVoteChanges                         int64         `json:"STEEM_MAX_VOTE_CHANGES"`
	SteemMaxWithdrawRoutes                      int64         `json:"STEEM_MAX_WITHDRAW_ROUTES"`
	SteemMaxWitnesses                           int64         `json:"STEEM_MAX_WITNESSES"`
	SteemMaxWitnessURLLength                    int64         `json:"STEEM_MAX_WITNESS_URL_LENGTH"`
	SteemMinerAccount                           string        `json:"STEEM_MINER_ACCOUNT"`
	SteemMinerPayPercent                        int64         `json:"STEEM_MINER_PAY_PERCENT"`
	SteemMiningReward                           []interface{} `json:"STEEM_MINING_REWARD"`
	SteemMiningTime                             string        `json:"STEEM_MINING_TIME"`
	SteemMinAccountCreationFee                  int64         `json:"STEEM_MIN_ACCOUNT_CREATION_FEE"`
	SteemMinAccountNameLength                   int64         `json:"STEEM_MIN_ACCOUNT_NAME_LENGTH"`
	SteemMinBlockSize                           int64         `json:"STEEM_MIN_BLOCK_SIZE"`
	SteemMinBlockSizeLimit                      int64         `json:"STEEM_MIN_BLOCK_SIZE_LIMIT"`
	SteemMinContentReward                       []interface{} `json:"STEEM_MIN_CONTENT_REWARD"`
	SteemMinCurateReward                        []interface{} `json:"STEEM_MIN_CURATE_REWARD"`
	SteemMinFeeds                               int64         `json:"STEEM_MIN_FEEDS"`
	SteemMinLiquidityReward                     []interface{} `json:"STEEM_MIN_LIQUIDITY_REWARD"`
	SteemMinLiquidityRewardPeriodSec            int64         `json:"STEEM_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemMinPayoutSbd                           []interface{} `json:"STEEM_MIN_PAYOUT_SBD"`
	SteemMinPermlinkLength                      int64         `json:"STEEM_MIN_PERMLINK_LENGTH"`
	SteemMinPowReward                           []interface{} `json:"STEEM_MIN_POW_REWARD"`
	SteemMinProducerReward                      []interface{} `json:"STEEM_MIN_PRODUCER_REWARD"`
	SteemMinReplyInterval                       int64         `json:"STEEM_MIN_REPLY_INTERVAL"`
	SteemMinRootCommentInterval                 int64         `json:"STEEM_MIN_ROOT_COMMENT_INTERVAL"`
	SteemMinTransactionExpirationLimit          int64         `json:"STEEM_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	SteemMinTransactionSizeLimit                int64         `json:"STEEM_MIN_TRANSACTION_SIZE_LIMIT"`
	SteemMinUndoHistory                         int64         `json:"STEEM_MIN_UNDO_HISTORY"`
	SteemMinVoteIntervalSec                     int64         `json:"STEEM_MIN_VOTE_INTERVAL_SEC"`
	SteemNullAccount                            string        `json:"STEEM_NULL_ACCOUNT"`
	SteemNumInitMiners                          int64         `json:"STEEM_NUM_INIT_MINERS"`
	SteemOwnerAuthHistoryTrackingStartBlockNum  int64         `json:"STEEM_OWNER_AUTH_HISTORY_TRACKING_START_BLOCK_NUM"`
	SteemOwnerAuthRecoveryPeriod                string        `json:"STEEM_OWNER_AUTH_RECOVERY_PERIOD"`
	SteemOwnerChallengeCooldown                 string        `json:"STEEM_OWNER_CHALLENGE_COOLDOWN"`
	SteemOwnerChallengeFee                      []interface{} `json:"STEEM_OWNER_CHALLENGE_FEE"`
	SteemOwnerUpdateLimit                       int64         `json:"STEEM_OWNER_UPDATE_LIMIT"`
	SteemPostAverageWindow                      int64         `json:"STEEM_POST_AVERAGE_WINDOW"`
	SteemPostRewardFundName                     string        `json:"STEEM_POST_REWARD_FUND_NAME"`
	SteemPostWeightConstant                     int64         `json:"STEEM_POST_WEIGHT_CONSTANT"`
	SteemPowAprPercent                          int64         `json:"STEEM_POW_APR_PERCENT"`
	SteemProducerAprPercent                     int64         `json:"STEEM_PRODUCER_APR_PERCENT"`
	SteemProxyToSelfAccount                     string        `json:"STEEM_PROXY_TO_SELF_ACCOUNT"`
	STEEMRECENTRSHARESDECAYTIMEHF17             string        `json:"STEEM_RECENT_RSHARES_DECAY_TIME_HF17"`
	STEEMRECENTRSHARESDECAYTIMEHF19             string        `json:"STEEM_RECENT_RSHARES_DECAY_TIME_HF19"`
	SteemReducedVotePowerRate                   int64         `json:"STEEM_REDUCED_VOTE_POWER_RATE"`
	SteemReverseAuctionWindowSeconds            int64         `json:"STEEM_REVERSE_AUCTION_WINDOW_SECONDS"`
	SteemRootPostParent                         string        `json:"STEEM_ROOT_POST_PARENT"`
	SteemSavingsWithdrawRequestLimit            int64         `json:"STEEM_SAVINGS_WITHDRAW_REQUEST_LIMIT"`
	SteemSavingsWithdrawTime                    string        `json:"STEEM_SAVINGS_WITHDRAW_TIME"`
	SteemSbdInterestCompoundIntervalSec         int64         `json:"STEEM_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SteemSbdStartPercent                        int64         `json:"STEEM_SBD_START_PERCENT"`
	SteemSbdStopPercent                         int64         `json:"STEEM_SBD_STOP_PERCENT"`
	SteemSecondsPerYear                         int64         `json:"STEEM_SECONDS_PER_YEAR"`
	SteemSecondCashoutWindow                    int64         `json:"STEEM_SECOND_CASHOUT_WINDOW"`
	SteemSoftMaxBlockSize                       int64         `json:"STEEM_SOFT_MAX_BLOCK_SIZE"`
	SteemSoftMaxCommentDepth                    int64         `json:"STEEM_SOFT_MAX_COMMENT_DEPTH"`
	SteemStartMinerVotingBlock                  int64         `json:"STEEM_START_MINER_VOTING_BLOCK"`
	SteemStartVestingBlock                      int64         `json:"STEEM_START_VESTING_BLOCK"`
	SteemSymbol                                 interface{}   `json:"STEEM_SYMBOL"`
	SteemTempAccount                            string        `json:"STEEM_TEMP_ACCOUNT"`
	STEEMUPVOTELOCKOUTHF17                      string        `json:"STEEM_UPVOTE_LOCKOUT_HF17"`
	STEEMUPVOTELOCKOUTHF7                       int64         `json:"STEEM_UPVOTE_LOCKOUT_HF7"`
	SteemVestingFundPercent                     int64         `json:"STEEM_VESTING_FUND_PERCENT"`
	SteemVestingWithdrawIntervals               int64         `json:"STEEM_VESTING_WITHDRAW_INTERVALS"`
	STEEMVESTINGWITHDRAWINTERVALSPREHF16        int64         `json:"STEEM_VESTING_WITHDRAW_INTERVALS_PRE_HF_16"`
	SteemVestingWithdrawIntervalSeconds         int64         `json:"STEEM_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	SteemVirtualScheduleLapLength               string        `json:"STEEM_VIRTUAL_SCHEDULE_LAP_LENGTH"`
	STEEMVIRTUALSCHEDULELAPLENGTH2              string        `json:"STEEM_VIRTUAL_SCHEDULE_LAP_LENGTH2"`
	SteemVoteDustThreshold                      int64         `json:"STEEM_VOTE_DUST_THRESHOLD"`
	SteemVoteRegenerationSeconds                int64         `json:"STEEM_VOTE_REGENERATION_SECONDS"`
	VestsSymbol                                 interface{}   `json:"VESTS_SYMBOL"`
}

func GetConfig(client gosteem.Client) (*GetConfigResponse, error) {
	var resp = &GetConfigResponse{}
	err := client.Call(MethodGetConfig, nil, resp)
	return resp, err
}

type GetDynamicGlobalPropertiesResponse struct {
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
}

func GetDynamicGlobalProperties(client gosteem.Client) (*GetDynamicGlobalPropertiesResponse, error) {
	var resp = &GetDynamicGlobalPropertiesResponse{}
	err := client.Call(MethodGetDynamicGlobalProperties, nil, resp)
	return resp, err
}

type GetWitnessScheduleResponse struct {
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
}

func GetWitnessSchedule(client gosteem.Client) (*GetWitnessScheduleResponse, error) {
	var resp = &GetWitnessScheduleResponse{}
	err := client.Call(MethodGetWitnessSchedule, nil, resp)
	return resp, err
}

type GetFeedHistoryResponse struct {
	CurrentMedianHistory struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
	} `json:"current_median_history"`
	ID           int64 `json:"id"`
	PriceHistory []struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
	} `json:"price_history"`
}

func GetFeedHistory(client gosteem.Client) (*GetFeedHistoryResponse, error) {
	var resp = &GetFeedHistoryResponse{}
	err := client.Call(MethodGetFeedHistory, nil, resp)
	return resp, err
}

func GetActiveWitnesses(client gosteem.Client) ([]string, error) {
	var resp []string
	err := client.Call(MethodGetActiveWitnesses, nil, &resp)
	return resp, err
}
