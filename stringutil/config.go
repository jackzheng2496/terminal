package stringutil

type Entertainer interface {
	FormattedOutput()
	UpdateTimeStamp(timestamp string)
	UpdateValue(value string)
	UpdateSubVal(value string)
}

func UpdateTimestamp(e Entertainer, timestamp string) {
	e.UpdateTimeStamp(timestamp)
}

func UpdateVal(e Entertainer, val string) {
	e.UpdateValue(val)
}

func UpdateSub(e Entertainer, val string) {
	e.UpdateSubVal(val)
}
