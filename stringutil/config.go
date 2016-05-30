package stringutil

type Entertainment interface {
	FormattedOutput()
	UpdateTimeStamp(timestamp string)
	UpdateValue(value string)
	UpdateSubVal(value string)
}

func UpdateTimestamp(e Entertainment, timestamp string) {
	e.UpdateTimeStamp(timestamp)
}

func UpdateVal(e Entertainment, val string) {
	e.UpdateValue(val)
}

func UpdateSub(e Entertainment, val string) {
	e.UpdateSubVal(val)
}
