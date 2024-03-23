package helper

type CountMissionApproval struct {
	TotalCount    int64 `json:"total_count"`
	CountPending  int64 `json:"count_pending"`
	CountApproved int64 `json:"count_approved"`
	CountRejected int64 `json:"count_rejected"`
}

type CountMission struct {
	TotalCount   int64 `json:"total_count"`
	CountActive  int64 `json:"count_active"`
	CountExpired int64 `json:"count_expired"`
}

type CountExchangeVoucher struct {
	TotalCount   int64 `json:"total_count"`
	CountNewest  int64 `json:"count_newest"`
	CountProcess int64 `json:"count_process"`
	CountDone    int64 `json:"count_done"`
}

type CountPrompt struct {
	TotalCount       int64 `json:"total_count"`
	CountOrganic     int64 `json:"count_organic"`
	CountAnorganic   int64 `json:"count_anorganic"`
	CountInformation int64 `json:"count_information"`
	CountLimitation  int64 `json:"count_limitation"`
}
