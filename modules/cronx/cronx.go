package cronx

import "github.com/robfig/cron/v3"

var client = cron.New()

func init() {
	client.Start()
}

// AddJob 增加任务
func AddJob(spec string, cmd cron.Job) (cron.EntryID, error) {
	return client.AddJob(spec, cmd)
}

// AddFunc 增加任务
func AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return client.AddFunc(spec, cmd)
}

// EveryMinuteFunc 每分钟
func EveryMinuteFunc(cmd func()) (cron.EntryID, error) {
	return AddFunc("* * * * *", cmd)
}

// EveryTenMinuteFunc 每10分钟
func EveryTenMinuteFunc(cmd func()) (cron.EntryID, error) {
	return AddFunc("*/10 * * * *", cmd)
}

// EveryHalfHourFunc 每半小时
func EveryHalfHourFunc(cmd func()) (cron.EntryID, error) {
	return AddFunc("*/30 * * * *", cmd)
}

// EveryHourFunc 每小时
func EveryHourFunc(cmd func()) (cron.EntryID, error) {
	return AddFunc("0 * * * *", cmd)
}
