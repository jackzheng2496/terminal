package stringutil

type Entertainment interface {
	FormattedOutput()
	UpdateTimeStamp(timestamp string)
	UpdateValue(value int)
	UpdateSubVal(value int)
}

func UpdateTimestamp(e Entertainment, timestamp string) {
	e.UpdateTimeStamp(timestamp)
}

func UpdateVal(e Entertainment, val int) {
	e.UpdateValue(val)
}

func UpdateSub(e Entertainment, val int) {
	e.UpdateSubVal(val)
}
