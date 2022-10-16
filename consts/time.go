package consts

import "time"

var (
	CSTSH = time.FixedZone("CST", 8*3600)
)

const (
	YYYY           = "2006"
	YYYYMM         = "2006-01"
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHH     = "2006-01-02 15"
	YYYYMMDDHHMM   = "2006-01-02 15:04"
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
)
