package stringutil

type Entertainer interface {
	FormattedOutput() string
	LongOutput() string
	UpdateTimeStamp(timestamp string)
	UpdateValue(value string)
	UpdateSubVal(value string)
	UpdateProducer(value string)
	SaveOutput() string
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

func GetData(e Entertainer) string {
	return e.LongOutput()
}

func UpdatePublisher(e Entertainer, val string) {
	e.UpdateProducer(val)
}
