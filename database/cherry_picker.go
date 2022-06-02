package database

// Session model of the aggregated data of cherry picker of all regions
type Session struct {
	PublicKey          string  `db:"public_key"`
	Chain              string  `db:"chain"`
	SessionKey         string  `db:"session_key"`
	Address            string  `db:"address"`
	SessionHeight      int     `db:"session_height"`
	TotalSuccess       int     `db:"total_success"`
	TotalFailure       int     `db:"total_failure"`
	AverageSuccessTime float64 `db:"avg_success_time"`
	Failure            bool    `db:"failure"`
}

// Region model of data of cherry picker for a single region
type Region struct {
	PublicKey                 string    `db:"public_key"`
	Chain                     string    `db:"chain"`
	SessionKey                string    `db:"session_key"`
	Region                    string    `db:"region"`
	Address                   string    `db:"address"`
	SessionHeight             int       `db:"session_height"`
	TotalSuccess              int       `db:"total_success"`
	TotalFailure              int       `db:"total_failure"`
	MedianSuccessLatency      []float32 `db:"median_success_latency"`
	WeightedSuccessLatency    []float32 `db:"weighted_success_latency"`
	P90Latency                []float32 `db:"p_90_latency"`
	SuccessRate               []float32 `db:"success_rate"`
	Attempts                  []int     `db:"attempts"`
	AvgSuccessLatency         float32   `db:"avg_success_latency"`
	AvgWeightedSuccessLatency float32   `db:"avg_weighted_success_latency"`
	Failure                   bool      `db:"failure"`
}

// CherryPickerStorage represents the operations where the cherry picker data is stored
type CherryPickerStorage interface {
	GetNodeSessionHeight(publicKey string, sessionHeight int) (*Session, error)
	GetNodeSessionFromKey(publicKey, sessionKey string) (*Session, error)
}
