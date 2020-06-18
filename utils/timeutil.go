package utils

import "time"

// AllDays get all days from timespan
func AllDays(sTiemStr, eTimeStr string) []string {
	sTime, err := time.Parse("2006-01-02 15:04:05", sTiemStr)
	if err != nil {
		return nil
	}
	eTime, err := time.Parse("2006-01-02 15:04:05", eTimeStr)
	if err != nil {
		return nil
	}
	if sTime.After(eTime) {
		return nil
	}
	timeArr := make([]string, 0)
	timeArr = append(timeArr, sTiemStr)
	for i := 1; ; i++ {
		temp, _ := time.Parse("2006-01-02 15:04:05", sTime.Add(24*time.Duration(i)*time.Hour).Format("2006-01-02")+" 00:00:00")
		if temp.After(eTime) {
			break
		}
		timeArr = append(timeArr, temp.Format("2006-01-02 15:04:05"))
	}
	return timeArr
}

// MyTimer with duration
// duration : h,m,s,ms,us,ns
func MyTimer(timerFunc func(), duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			timerFunc()
		}
	}
}
