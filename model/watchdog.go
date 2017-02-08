package model

type Watchdog struct {
	Wid       int64  `xorm:"not null pk autoincr INT(11)"`
	Uid       int    `xorm:"not null default 0 index INT(10)"`
	Type      string `xorm:"not null default '' index VARCHAR(64)"`
	Message   string `xorm:"not null LONGTEXT"`
	Variables []byte `xorm:"not null LONGBLOB"`
	Severity  int    `xorm:"not null default 0 index TINYINT(3)"`
	Link      string `xorm:"TEXT"`
	Location  string `xorm:"not null TEXT"`
	Referer   string `xorm:"TEXT"`
	Hostname  string `xorm:"not null default '' VARCHAR(128)"`
	Timestamp int    `xorm:"not null default 0 INT(11)"`
}

// GetWatchdogsCount Watchdogs' Count
func GetWatchdogsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Watchdog{})
	return total, err
}

// GetWatchdogCountViaWid Get Watchdog via Wid
func GetWatchdogCountViaWid(iWid int64) int64 {
	n, _ := Engine.Where("wid = ?", iWid).Count(&Watchdog{Wid: iWid})
	return n
}

// GetWatchdogCountViaUid Get Watchdog via Uid
func GetWatchdogCountViaUid(iUid int) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&Watchdog{Uid: iUid})
	return n
}

// GetWatchdogCountViaType Get Watchdog via Type
func GetWatchdogCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&Watchdog{Type: iType})
	return n
}

// GetWatchdogCountViaMessage Get Watchdog via Message
func GetWatchdogCountViaMessage(iMessage string) int64 {
	n, _ := Engine.Where("message = ?", iMessage).Count(&Watchdog{Message: iMessage})
	return n
}

// GetWatchdogCountViaVariables Get Watchdog via Variables
func GetWatchdogCountViaVariables(iVariables []byte) int64 {
	n, _ := Engine.Where("variables = ?", iVariables).Count(&Watchdog{Variables: iVariables})
	return n
}

// GetWatchdogCountViaSeverity Get Watchdog via Severity
func GetWatchdogCountViaSeverity(iSeverity int) int64 {
	n, _ := Engine.Where("severity = ?", iSeverity).Count(&Watchdog{Severity: iSeverity})
	return n
}

// GetWatchdogCountViaLink Get Watchdog via Link
func GetWatchdogCountViaLink(iLink string) int64 {
	n, _ := Engine.Where("link = ?", iLink).Count(&Watchdog{Link: iLink})
	return n
}

// GetWatchdogCountViaLocation Get Watchdog via Location
func GetWatchdogCountViaLocation(iLocation string) int64 {
	n, _ := Engine.Where("location = ?", iLocation).Count(&Watchdog{Location: iLocation})
	return n
}

// GetWatchdogCountViaReferer Get Watchdog via Referer
func GetWatchdogCountViaReferer(iReferer string) int64 {
	n, _ := Engine.Where("referer = ?", iReferer).Count(&Watchdog{Referer: iReferer})
	return n
}

// GetWatchdogCountViaHostname Get Watchdog via Hostname
func GetWatchdogCountViaHostname(iHostname string) int64 {
	n, _ := Engine.Where("hostname = ?", iHostname).Count(&Watchdog{Hostname: iHostname})
	return n
}

// GetWatchdogCountViaTimestamp Get Watchdog via Timestamp
func GetWatchdogCountViaTimestamp(iTimestamp int) int64 {
	n, _ := Engine.Where("timestamp = ?", iTimestamp).Count(&Watchdog{Timestamp: iTimestamp})
	return n
}

// GetWatchdogsByWidAndUidAndType Get Watchdogs via WidAndUidAndType
func GetWatchdogsByWidAndUidAndType(offset int, limit int, Wid_ int64, Uid_ int, Type_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and type = ?", Wid_, Uid_, Type_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndMessage Get Watchdogs via WidAndUidAndMessage
func GetWatchdogsByWidAndUidAndMessage(offset int, limit int, Wid_ int64, Uid_ int, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and message = ?", Wid_, Uid_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndVariables Get Watchdogs via WidAndUidAndVariables
func GetWatchdogsByWidAndUidAndVariables(offset int, limit int, Wid_ int64, Uid_ int, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and variables = ?", Wid_, Uid_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndSeverity Get Watchdogs via WidAndUidAndSeverity
func GetWatchdogsByWidAndUidAndSeverity(offset int, limit int, Wid_ int64, Uid_ int, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and severity = ?", Wid_, Uid_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndLink Get Watchdogs via WidAndUidAndLink
func GetWatchdogsByWidAndUidAndLink(offset int, limit int, Wid_ int64, Uid_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and link = ?", Wid_, Uid_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndLocation Get Watchdogs via WidAndUidAndLocation
func GetWatchdogsByWidAndUidAndLocation(offset int, limit int, Wid_ int64, Uid_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and location = ?", Wid_, Uid_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndReferer Get Watchdogs via WidAndUidAndReferer
func GetWatchdogsByWidAndUidAndReferer(offset int, limit int, Wid_ int64, Uid_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and referer = ?", Wid_, Uid_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndHostname Get Watchdogs via WidAndUidAndHostname
func GetWatchdogsByWidAndUidAndHostname(offset int, limit int, Wid_ int64, Uid_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and hostname = ?", Wid_, Uid_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUidAndTimestamp Get Watchdogs via WidAndUidAndTimestamp
func GetWatchdogsByWidAndUidAndTimestamp(offset int, limit int, Wid_ int64, Uid_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ? and timestamp = ?", Wid_, Uid_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndMessage Get Watchdogs via WidAndTypeAndMessage
func GetWatchdogsByWidAndTypeAndMessage(offset int, limit int, Wid_ int64, Type_ string, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and message = ?", Wid_, Type_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndVariables Get Watchdogs via WidAndTypeAndVariables
func GetWatchdogsByWidAndTypeAndVariables(offset int, limit int, Wid_ int64, Type_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and variables = ?", Wid_, Type_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndSeverity Get Watchdogs via WidAndTypeAndSeverity
func GetWatchdogsByWidAndTypeAndSeverity(offset int, limit int, Wid_ int64, Type_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and severity = ?", Wid_, Type_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndLink Get Watchdogs via WidAndTypeAndLink
func GetWatchdogsByWidAndTypeAndLink(offset int, limit int, Wid_ int64, Type_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and link = ?", Wid_, Type_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndLocation Get Watchdogs via WidAndTypeAndLocation
func GetWatchdogsByWidAndTypeAndLocation(offset int, limit int, Wid_ int64, Type_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and location = ?", Wid_, Type_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndReferer Get Watchdogs via WidAndTypeAndReferer
func GetWatchdogsByWidAndTypeAndReferer(offset int, limit int, Wid_ int64, Type_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and referer = ?", Wid_, Type_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndHostname Get Watchdogs via WidAndTypeAndHostname
func GetWatchdogsByWidAndTypeAndHostname(offset int, limit int, Wid_ int64, Type_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and hostname = ?", Wid_, Type_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTypeAndTimestamp Get Watchdogs via WidAndTypeAndTimestamp
func GetWatchdogsByWidAndTypeAndTimestamp(offset int, limit int, Wid_ int64, Type_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ? and timestamp = ?", Wid_, Type_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndVariables Get Watchdogs via WidAndMessageAndVariables
func GetWatchdogsByWidAndMessageAndVariables(offset int, limit int, Wid_ int64, Message_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and variables = ?", Wid_, Message_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndSeverity Get Watchdogs via WidAndMessageAndSeverity
func GetWatchdogsByWidAndMessageAndSeverity(offset int, limit int, Wid_ int64, Message_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and severity = ?", Wid_, Message_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndLink Get Watchdogs via WidAndMessageAndLink
func GetWatchdogsByWidAndMessageAndLink(offset int, limit int, Wid_ int64, Message_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and link = ?", Wid_, Message_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndLocation Get Watchdogs via WidAndMessageAndLocation
func GetWatchdogsByWidAndMessageAndLocation(offset int, limit int, Wid_ int64, Message_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and location = ?", Wid_, Message_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndReferer Get Watchdogs via WidAndMessageAndReferer
func GetWatchdogsByWidAndMessageAndReferer(offset int, limit int, Wid_ int64, Message_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and referer = ?", Wid_, Message_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndHostname Get Watchdogs via WidAndMessageAndHostname
func GetWatchdogsByWidAndMessageAndHostname(offset int, limit int, Wid_ int64, Message_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and hostname = ?", Wid_, Message_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessageAndTimestamp Get Watchdogs via WidAndMessageAndTimestamp
func GetWatchdogsByWidAndMessageAndTimestamp(offset int, limit int, Wid_ int64, Message_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ? and timestamp = ?", Wid_, Message_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndSeverity Get Watchdogs via WidAndVariablesAndSeverity
func GetWatchdogsByWidAndVariablesAndSeverity(offset int, limit int, Wid_ int64, Variables_ []byte, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and severity = ?", Wid_, Variables_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndLink Get Watchdogs via WidAndVariablesAndLink
func GetWatchdogsByWidAndVariablesAndLink(offset int, limit int, Wid_ int64, Variables_ []byte, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and link = ?", Wid_, Variables_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndLocation Get Watchdogs via WidAndVariablesAndLocation
func GetWatchdogsByWidAndVariablesAndLocation(offset int, limit int, Wid_ int64, Variables_ []byte, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and location = ?", Wid_, Variables_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndReferer Get Watchdogs via WidAndVariablesAndReferer
func GetWatchdogsByWidAndVariablesAndReferer(offset int, limit int, Wid_ int64, Variables_ []byte, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and referer = ?", Wid_, Variables_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndHostname Get Watchdogs via WidAndVariablesAndHostname
func GetWatchdogsByWidAndVariablesAndHostname(offset int, limit int, Wid_ int64, Variables_ []byte, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and hostname = ?", Wid_, Variables_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariablesAndTimestamp Get Watchdogs via WidAndVariablesAndTimestamp
func GetWatchdogsByWidAndVariablesAndTimestamp(offset int, limit int, Wid_ int64, Variables_ []byte, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ? and timestamp = ?", Wid_, Variables_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverityAndLink Get Watchdogs via WidAndSeverityAndLink
func GetWatchdogsByWidAndSeverityAndLink(offset int, limit int, Wid_ int64, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ? and link = ?", Wid_, Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverityAndLocation Get Watchdogs via WidAndSeverityAndLocation
func GetWatchdogsByWidAndSeverityAndLocation(offset int, limit int, Wid_ int64, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ? and location = ?", Wid_, Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverityAndReferer Get Watchdogs via WidAndSeverityAndReferer
func GetWatchdogsByWidAndSeverityAndReferer(offset int, limit int, Wid_ int64, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ? and referer = ?", Wid_, Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverityAndHostname Get Watchdogs via WidAndSeverityAndHostname
func GetWatchdogsByWidAndSeverityAndHostname(offset int, limit int, Wid_ int64, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ? and hostname = ?", Wid_, Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverityAndTimestamp Get Watchdogs via WidAndSeverityAndTimestamp
func GetWatchdogsByWidAndSeverityAndTimestamp(offset int, limit int, Wid_ int64, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ? and timestamp = ?", Wid_, Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLinkAndLocation Get Watchdogs via WidAndLinkAndLocation
func GetWatchdogsByWidAndLinkAndLocation(offset int, limit int, Wid_ int64, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and link = ? and location = ?", Wid_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLinkAndReferer Get Watchdogs via WidAndLinkAndReferer
func GetWatchdogsByWidAndLinkAndReferer(offset int, limit int, Wid_ int64, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and link = ? and referer = ?", Wid_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLinkAndHostname Get Watchdogs via WidAndLinkAndHostname
func GetWatchdogsByWidAndLinkAndHostname(offset int, limit int, Wid_ int64, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and link = ? and hostname = ?", Wid_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLinkAndTimestamp Get Watchdogs via WidAndLinkAndTimestamp
func GetWatchdogsByWidAndLinkAndTimestamp(offset int, limit int, Wid_ int64, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and link = ? and timestamp = ?", Wid_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLocationAndReferer Get Watchdogs via WidAndLocationAndReferer
func GetWatchdogsByWidAndLocationAndReferer(offset int, limit int, Wid_ int64, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and location = ? and referer = ?", Wid_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLocationAndHostname Get Watchdogs via WidAndLocationAndHostname
func GetWatchdogsByWidAndLocationAndHostname(offset int, limit int, Wid_ int64, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and location = ? and hostname = ?", Wid_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLocationAndTimestamp Get Watchdogs via WidAndLocationAndTimestamp
func GetWatchdogsByWidAndLocationAndTimestamp(offset int, limit int, Wid_ int64, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and location = ? and timestamp = ?", Wid_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndRefererAndHostname Get Watchdogs via WidAndRefererAndHostname
func GetWatchdogsByWidAndRefererAndHostname(offset int, limit int, Wid_ int64, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and referer = ? and hostname = ?", Wid_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndRefererAndTimestamp Get Watchdogs via WidAndRefererAndTimestamp
func GetWatchdogsByWidAndRefererAndTimestamp(offset int, limit int, Wid_ int64, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and referer = ? and timestamp = ?", Wid_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndHostnameAndTimestamp Get Watchdogs via WidAndHostnameAndTimestamp
func GetWatchdogsByWidAndHostnameAndTimestamp(offset int, limit int, Wid_ int64, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and hostname = ? and timestamp = ?", Wid_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndMessage Get Watchdogs via UidAndTypeAndMessage
func GetWatchdogsByUidAndTypeAndMessage(offset int, limit int, Uid_ int, Type_ string, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and message = ?", Uid_, Type_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndVariables Get Watchdogs via UidAndTypeAndVariables
func GetWatchdogsByUidAndTypeAndVariables(offset int, limit int, Uid_ int, Type_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and variables = ?", Uid_, Type_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndSeverity Get Watchdogs via UidAndTypeAndSeverity
func GetWatchdogsByUidAndTypeAndSeverity(offset int, limit int, Uid_ int, Type_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and severity = ?", Uid_, Type_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndLink Get Watchdogs via UidAndTypeAndLink
func GetWatchdogsByUidAndTypeAndLink(offset int, limit int, Uid_ int, Type_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and link = ?", Uid_, Type_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndLocation Get Watchdogs via UidAndTypeAndLocation
func GetWatchdogsByUidAndTypeAndLocation(offset int, limit int, Uid_ int, Type_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and location = ?", Uid_, Type_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndReferer Get Watchdogs via UidAndTypeAndReferer
func GetWatchdogsByUidAndTypeAndReferer(offset int, limit int, Uid_ int, Type_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and referer = ?", Uid_, Type_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndHostname Get Watchdogs via UidAndTypeAndHostname
func GetWatchdogsByUidAndTypeAndHostname(offset int, limit int, Uid_ int, Type_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and hostname = ?", Uid_, Type_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTypeAndTimestamp Get Watchdogs via UidAndTypeAndTimestamp
func GetWatchdogsByUidAndTypeAndTimestamp(offset int, limit int, Uid_ int, Type_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ? and timestamp = ?", Uid_, Type_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndVariables Get Watchdogs via UidAndMessageAndVariables
func GetWatchdogsByUidAndMessageAndVariables(offset int, limit int, Uid_ int, Message_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and variables = ?", Uid_, Message_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndSeverity Get Watchdogs via UidAndMessageAndSeverity
func GetWatchdogsByUidAndMessageAndSeverity(offset int, limit int, Uid_ int, Message_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and severity = ?", Uid_, Message_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndLink Get Watchdogs via UidAndMessageAndLink
func GetWatchdogsByUidAndMessageAndLink(offset int, limit int, Uid_ int, Message_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and link = ?", Uid_, Message_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndLocation Get Watchdogs via UidAndMessageAndLocation
func GetWatchdogsByUidAndMessageAndLocation(offset int, limit int, Uid_ int, Message_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and location = ?", Uid_, Message_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndReferer Get Watchdogs via UidAndMessageAndReferer
func GetWatchdogsByUidAndMessageAndReferer(offset int, limit int, Uid_ int, Message_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and referer = ?", Uid_, Message_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndHostname Get Watchdogs via UidAndMessageAndHostname
func GetWatchdogsByUidAndMessageAndHostname(offset int, limit int, Uid_ int, Message_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and hostname = ?", Uid_, Message_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessageAndTimestamp Get Watchdogs via UidAndMessageAndTimestamp
func GetWatchdogsByUidAndMessageAndTimestamp(offset int, limit int, Uid_ int, Message_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ? and timestamp = ?", Uid_, Message_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndSeverity Get Watchdogs via UidAndVariablesAndSeverity
func GetWatchdogsByUidAndVariablesAndSeverity(offset int, limit int, Uid_ int, Variables_ []byte, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and severity = ?", Uid_, Variables_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndLink Get Watchdogs via UidAndVariablesAndLink
func GetWatchdogsByUidAndVariablesAndLink(offset int, limit int, Uid_ int, Variables_ []byte, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and link = ?", Uid_, Variables_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndLocation Get Watchdogs via UidAndVariablesAndLocation
func GetWatchdogsByUidAndVariablesAndLocation(offset int, limit int, Uid_ int, Variables_ []byte, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and location = ?", Uid_, Variables_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndReferer Get Watchdogs via UidAndVariablesAndReferer
func GetWatchdogsByUidAndVariablesAndReferer(offset int, limit int, Uid_ int, Variables_ []byte, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and referer = ?", Uid_, Variables_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndHostname Get Watchdogs via UidAndVariablesAndHostname
func GetWatchdogsByUidAndVariablesAndHostname(offset int, limit int, Uid_ int, Variables_ []byte, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and hostname = ?", Uid_, Variables_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariablesAndTimestamp Get Watchdogs via UidAndVariablesAndTimestamp
func GetWatchdogsByUidAndVariablesAndTimestamp(offset int, limit int, Uid_ int, Variables_ []byte, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ? and timestamp = ?", Uid_, Variables_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverityAndLink Get Watchdogs via UidAndSeverityAndLink
func GetWatchdogsByUidAndSeverityAndLink(offset int, limit int, Uid_ int, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ? and link = ?", Uid_, Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverityAndLocation Get Watchdogs via UidAndSeverityAndLocation
func GetWatchdogsByUidAndSeverityAndLocation(offset int, limit int, Uid_ int, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ? and location = ?", Uid_, Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverityAndReferer Get Watchdogs via UidAndSeverityAndReferer
func GetWatchdogsByUidAndSeverityAndReferer(offset int, limit int, Uid_ int, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ? and referer = ?", Uid_, Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverityAndHostname Get Watchdogs via UidAndSeverityAndHostname
func GetWatchdogsByUidAndSeverityAndHostname(offset int, limit int, Uid_ int, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ? and hostname = ?", Uid_, Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverityAndTimestamp Get Watchdogs via UidAndSeverityAndTimestamp
func GetWatchdogsByUidAndSeverityAndTimestamp(offset int, limit int, Uid_ int, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ? and timestamp = ?", Uid_, Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLinkAndLocation Get Watchdogs via UidAndLinkAndLocation
func GetWatchdogsByUidAndLinkAndLocation(offset int, limit int, Uid_ int, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and link = ? and location = ?", Uid_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLinkAndReferer Get Watchdogs via UidAndLinkAndReferer
func GetWatchdogsByUidAndLinkAndReferer(offset int, limit int, Uid_ int, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and link = ? and referer = ?", Uid_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLinkAndHostname Get Watchdogs via UidAndLinkAndHostname
func GetWatchdogsByUidAndLinkAndHostname(offset int, limit int, Uid_ int, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and link = ? and hostname = ?", Uid_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLinkAndTimestamp Get Watchdogs via UidAndLinkAndTimestamp
func GetWatchdogsByUidAndLinkAndTimestamp(offset int, limit int, Uid_ int, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and link = ? and timestamp = ?", Uid_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLocationAndReferer Get Watchdogs via UidAndLocationAndReferer
func GetWatchdogsByUidAndLocationAndReferer(offset int, limit int, Uid_ int, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and location = ? and referer = ?", Uid_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLocationAndHostname Get Watchdogs via UidAndLocationAndHostname
func GetWatchdogsByUidAndLocationAndHostname(offset int, limit int, Uid_ int, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and location = ? and hostname = ?", Uid_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLocationAndTimestamp Get Watchdogs via UidAndLocationAndTimestamp
func GetWatchdogsByUidAndLocationAndTimestamp(offset int, limit int, Uid_ int, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and location = ? and timestamp = ?", Uid_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndRefererAndHostname Get Watchdogs via UidAndRefererAndHostname
func GetWatchdogsByUidAndRefererAndHostname(offset int, limit int, Uid_ int, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and referer = ? and hostname = ?", Uid_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndRefererAndTimestamp Get Watchdogs via UidAndRefererAndTimestamp
func GetWatchdogsByUidAndRefererAndTimestamp(offset int, limit int, Uid_ int, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and referer = ? and timestamp = ?", Uid_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndHostnameAndTimestamp Get Watchdogs via UidAndHostnameAndTimestamp
func GetWatchdogsByUidAndHostnameAndTimestamp(offset int, limit int, Uid_ int, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and hostname = ? and timestamp = ?", Uid_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndVariables Get Watchdogs via TypeAndMessageAndVariables
func GetWatchdogsByTypeAndMessageAndVariables(offset int, limit int, Type_ string, Message_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and variables = ?", Type_, Message_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndSeverity Get Watchdogs via TypeAndMessageAndSeverity
func GetWatchdogsByTypeAndMessageAndSeverity(offset int, limit int, Type_ string, Message_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and severity = ?", Type_, Message_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndLink Get Watchdogs via TypeAndMessageAndLink
func GetWatchdogsByTypeAndMessageAndLink(offset int, limit int, Type_ string, Message_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and link = ?", Type_, Message_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndLocation Get Watchdogs via TypeAndMessageAndLocation
func GetWatchdogsByTypeAndMessageAndLocation(offset int, limit int, Type_ string, Message_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and location = ?", Type_, Message_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndReferer Get Watchdogs via TypeAndMessageAndReferer
func GetWatchdogsByTypeAndMessageAndReferer(offset int, limit int, Type_ string, Message_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and referer = ?", Type_, Message_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndHostname Get Watchdogs via TypeAndMessageAndHostname
func GetWatchdogsByTypeAndMessageAndHostname(offset int, limit int, Type_ string, Message_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and hostname = ?", Type_, Message_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessageAndTimestamp Get Watchdogs via TypeAndMessageAndTimestamp
func GetWatchdogsByTypeAndMessageAndTimestamp(offset int, limit int, Type_ string, Message_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ? and timestamp = ?", Type_, Message_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndSeverity Get Watchdogs via TypeAndVariablesAndSeverity
func GetWatchdogsByTypeAndVariablesAndSeverity(offset int, limit int, Type_ string, Variables_ []byte, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and severity = ?", Type_, Variables_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndLink Get Watchdogs via TypeAndVariablesAndLink
func GetWatchdogsByTypeAndVariablesAndLink(offset int, limit int, Type_ string, Variables_ []byte, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and link = ?", Type_, Variables_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndLocation Get Watchdogs via TypeAndVariablesAndLocation
func GetWatchdogsByTypeAndVariablesAndLocation(offset int, limit int, Type_ string, Variables_ []byte, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and location = ?", Type_, Variables_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndReferer Get Watchdogs via TypeAndVariablesAndReferer
func GetWatchdogsByTypeAndVariablesAndReferer(offset int, limit int, Type_ string, Variables_ []byte, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and referer = ?", Type_, Variables_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndHostname Get Watchdogs via TypeAndVariablesAndHostname
func GetWatchdogsByTypeAndVariablesAndHostname(offset int, limit int, Type_ string, Variables_ []byte, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and hostname = ?", Type_, Variables_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariablesAndTimestamp Get Watchdogs via TypeAndVariablesAndTimestamp
func GetWatchdogsByTypeAndVariablesAndTimestamp(offset int, limit int, Type_ string, Variables_ []byte, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ? and timestamp = ?", Type_, Variables_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverityAndLink Get Watchdogs via TypeAndSeverityAndLink
func GetWatchdogsByTypeAndSeverityAndLink(offset int, limit int, Type_ string, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ? and link = ?", Type_, Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverityAndLocation Get Watchdogs via TypeAndSeverityAndLocation
func GetWatchdogsByTypeAndSeverityAndLocation(offset int, limit int, Type_ string, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ? and location = ?", Type_, Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverityAndReferer Get Watchdogs via TypeAndSeverityAndReferer
func GetWatchdogsByTypeAndSeverityAndReferer(offset int, limit int, Type_ string, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ? and referer = ?", Type_, Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverityAndHostname Get Watchdogs via TypeAndSeverityAndHostname
func GetWatchdogsByTypeAndSeverityAndHostname(offset int, limit int, Type_ string, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ? and hostname = ?", Type_, Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverityAndTimestamp Get Watchdogs via TypeAndSeverityAndTimestamp
func GetWatchdogsByTypeAndSeverityAndTimestamp(offset int, limit int, Type_ string, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ? and timestamp = ?", Type_, Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLinkAndLocation Get Watchdogs via TypeAndLinkAndLocation
func GetWatchdogsByTypeAndLinkAndLocation(offset int, limit int, Type_ string, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and link = ? and location = ?", Type_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLinkAndReferer Get Watchdogs via TypeAndLinkAndReferer
func GetWatchdogsByTypeAndLinkAndReferer(offset int, limit int, Type_ string, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and link = ? and referer = ?", Type_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLinkAndHostname Get Watchdogs via TypeAndLinkAndHostname
func GetWatchdogsByTypeAndLinkAndHostname(offset int, limit int, Type_ string, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and link = ? and hostname = ?", Type_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLinkAndTimestamp Get Watchdogs via TypeAndLinkAndTimestamp
func GetWatchdogsByTypeAndLinkAndTimestamp(offset int, limit int, Type_ string, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and link = ? and timestamp = ?", Type_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLocationAndReferer Get Watchdogs via TypeAndLocationAndReferer
func GetWatchdogsByTypeAndLocationAndReferer(offset int, limit int, Type_ string, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and location = ? and referer = ?", Type_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLocationAndHostname Get Watchdogs via TypeAndLocationAndHostname
func GetWatchdogsByTypeAndLocationAndHostname(offset int, limit int, Type_ string, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and location = ? and hostname = ?", Type_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLocationAndTimestamp Get Watchdogs via TypeAndLocationAndTimestamp
func GetWatchdogsByTypeAndLocationAndTimestamp(offset int, limit int, Type_ string, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and location = ? and timestamp = ?", Type_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndRefererAndHostname Get Watchdogs via TypeAndRefererAndHostname
func GetWatchdogsByTypeAndRefererAndHostname(offset int, limit int, Type_ string, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and referer = ? and hostname = ?", Type_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndRefererAndTimestamp Get Watchdogs via TypeAndRefererAndTimestamp
func GetWatchdogsByTypeAndRefererAndTimestamp(offset int, limit int, Type_ string, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and referer = ? and timestamp = ?", Type_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndHostnameAndTimestamp Get Watchdogs via TypeAndHostnameAndTimestamp
func GetWatchdogsByTypeAndHostnameAndTimestamp(offset int, limit int, Type_ string, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and hostname = ? and timestamp = ?", Type_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndSeverity Get Watchdogs via MessageAndVariablesAndSeverity
func GetWatchdogsByMessageAndVariablesAndSeverity(offset int, limit int, Message_ string, Variables_ []byte, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and severity = ?", Message_, Variables_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndLink Get Watchdogs via MessageAndVariablesAndLink
func GetWatchdogsByMessageAndVariablesAndLink(offset int, limit int, Message_ string, Variables_ []byte, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and link = ?", Message_, Variables_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndLocation Get Watchdogs via MessageAndVariablesAndLocation
func GetWatchdogsByMessageAndVariablesAndLocation(offset int, limit int, Message_ string, Variables_ []byte, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and location = ?", Message_, Variables_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndReferer Get Watchdogs via MessageAndVariablesAndReferer
func GetWatchdogsByMessageAndVariablesAndReferer(offset int, limit int, Message_ string, Variables_ []byte, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and referer = ?", Message_, Variables_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndHostname Get Watchdogs via MessageAndVariablesAndHostname
func GetWatchdogsByMessageAndVariablesAndHostname(offset int, limit int, Message_ string, Variables_ []byte, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and hostname = ?", Message_, Variables_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariablesAndTimestamp Get Watchdogs via MessageAndVariablesAndTimestamp
func GetWatchdogsByMessageAndVariablesAndTimestamp(offset int, limit int, Message_ string, Variables_ []byte, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ? and timestamp = ?", Message_, Variables_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverityAndLink Get Watchdogs via MessageAndSeverityAndLink
func GetWatchdogsByMessageAndSeverityAndLink(offset int, limit int, Message_ string, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ? and link = ?", Message_, Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverityAndLocation Get Watchdogs via MessageAndSeverityAndLocation
func GetWatchdogsByMessageAndSeverityAndLocation(offset int, limit int, Message_ string, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ? and location = ?", Message_, Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverityAndReferer Get Watchdogs via MessageAndSeverityAndReferer
func GetWatchdogsByMessageAndSeverityAndReferer(offset int, limit int, Message_ string, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ? and referer = ?", Message_, Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverityAndHostname Get Watchdogs via MessageAndSeverityAndHostname
func GetWatchdogsByMessageAndSeverityAndHostname(offset int, limit int, Message_ string, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ? and hostname = ?", Message_, Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverityAndTimestamp Get Watchdogs via MessageAndSeverityAndTimestamp
func GetWatchdogsByMessageAndSeverityAndTimestamp(offset int, limit int, Message_ string, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ? and timestamp = ?", Message_, Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLinkAndLocation Get Watchdogs via MessageAndLinkAndLocation
func GetWatchdogsByMessageAndLinkAndLocation(offset int, limit int, Message_ string, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and link = ? and location = ?", Message_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLinkAndReferer Get Watchdogs via MessageAndLinkAndReferer
func GetWatchdogsByMessageAndLinkAndReferer(offset int, limit int, Message_ string, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and link = ? and referer = ?", Message_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLinkAndHostname Get Watchdogs via MessageAndLinkAndHostname
func GetWatchdogsByMessageAndLinkAndHostname(offset int, limit int, Message_ string, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and link = ? and hostname = ?", Message_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLinkAndTimestamp Get Watchdogs via MessageAndLinkAndTimestamp
func GetWatchdogsByMessageAndLinkAndTimestamp(offset int, limit int, Message_ string, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and link = ? and timestamp = ?", Message_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLocationAndReferer Get Watchdogs via MessageAndLocationAndReferer
func GetWatchdogsByMessageAndLocationAndReferer(offset int, limit int, Message_ string, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and location = ? and referer = ?", Message_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLocationAndHostname Get Watchdogs via MessageAndLocationAndHostname
func GetWatchdogsByMessageAndLocationAndHostname(offset int, limit int, Message_ string, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and location = ? and hostname = ?", Message_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLocationAndTimestamp Get Watchdogs via MessageAndLocationAndTimestamp
func GetWatchdogsByMessageAndLocationAndTimestamp(offset int, limit int, Message_ string, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and location = ? and timestamp = ?", Message_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndRefererAndHostname Get Watchdogs via MessageAndRefererAndHostname
func GetWatchdogsByMessageAndRefererAndHostname(offset int, limit int, Message_ string, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and referer = ? and hostname = ?", Message_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndRefererAndTimestamp Get Watchdogs via MessageAndRefererAndTimestamp
func GetWatchdogsByMessageAndRefererAndTimestamp(offset int, limit int, Message_ string, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and referer = ? and timestamp = ?", Message_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndHostnameAndTimestamp Get Watchdogs via MessageAndHostnameAndTimestamp
func GetWatchdogsByMessageAndHostnameAndTimestamp(offset int, limit int, Message_ string, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and hostname = ? and timestamp = ?", Message_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverityAndLink Get Watchdogs via VariablesAndSeverityAndLink
func GetWatchdogsByVariablesAndSeverityAndLink(offset int, limit int, Variables_ []byte, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ? and link = ?", Variables_, Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverityAndLocation Get Watchdogs via VariablesAndSeverityAndLocation
func GetWatchdogsByVariablesAndSeverityAndLocation(offset int, limit int, Variables_ []byte, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ? and location = ?", Variables_, Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverityAndReferer Get Watchdogs via VariablesAndSeverityAndReferer
func GetWatchdogsByVariablesAndSeverityAndReferer(offset int, limit int, Variables_ []byte, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ? and referer = ?", Variables_, Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverityAndHostname Get Watchdogs via VariablesAndSeverityAndHostname
func GetWatchdogsByVariablesAndSeverityAndHostname(offset int, limit int, Variables_ []byte, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ? and hostname = ?", Variables_, Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverityAndTimestamp Get Watchdogs via VariablesAndSeverityAndTimestamp
func GetWatchdogsByVariablesAndSeverityAndTimestamp(offset int, limit int, Variables_ []byte, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ? and timestamp = ?", Variables_, Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLinkAndLocation Get Watchdogs via VariablesAndLinkAndLocation
func GetWatchdogsByVariablesAndLinkAndLocation(offset int, limit int, Variables_ []byte, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and link = ? and location = ?", Variables_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLinkAndReferer Get Watchdogs via VariablesAndLinkAndReferer
func GetWatchdogsByVariablesAndLinkAndReferer(offset int, limit int, Variables_ []byte, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and link = ? and referer = ?", Variables_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLinkAndHostname Get Watchdogs via VariablesAndLinkAndHostname
func GetWatchdogsByVariablesAndLinkAndHostname(offset int, limit int, Variables_ []byte, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and link = ? and hostname = ?", Variables_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLinkAndTimestamp Get Watchdogs via VariablesAndLinkAndTimestamp
func GetWatchdogsByVariablesAndLinkAndTimestamp(offset int, limit int, Variables_ []byte, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and link = ? and timestamp = ?", Variables_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLocationAndReferer Get Watchdogs via VariablesAndLocationAndReferer
func GetWatchdogsByVariablesAndLocationAndReferer(offset int, limit int, Variables_ []byte, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and location = ? and referer = ?", Variables_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLocationAndHostname Get Watchdogs via VariablesAndLocationAndHostname
func GetWatchdogsByVariablesAndLocationAndHostname(offset int, limit int, Variables_ []byte, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and location = ? and hostname = ?", Variables_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLocationAndTimestamp Get Watchdogs via VariablesAndLocationAndTimestamp
func GetWatchdogsByVariablesAndLocationAndTimestamp(offset int, limit int, Variables_ []byte, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and location = ? and timestamp = ?", Variables_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndRefererAndHostname Get Watchdogs via VariablesAndRefererAndHostname
func GetWatchdogsByVariablesAndRefererAndHostname(offset int, limit int, Variables_ []byte, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and referer = ? and hostname = ?", Variables_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndRefererAndTimestamp Get Watchdogs via VariablesAndRefererAndTimestamp
func GetWatchdogsByVariablesAndRefererAndTimestamp(offset int, limit int, Variables_ []byte, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and referer = ? and timestamp = ?", Variables_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndHostnameAndTimestamp Get Watchdogs via VariablesAndHostnameAndTimestamp
func GetWatchdogsByVariablesAndHostnameAndTimestamp(offset int, limit int, Variables_ []byte, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and hostname = ? and timestamp = ?", Variables_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLinkAndLocation Get Watchdogs via SeverityAndLinkAndLocation
func GetWatchdogsBySeverityAndLinkAndLocation(offset int, limit int, Severity_ int, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and link = ? and location = ?", Severity_, Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLinkAndReferer Get Watchdogs via SeverityAndLinkAndReferer
func GetWatchdogsBySeverityAndLinkAndReferer(offset int, limit int, Severity_ int, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and link = ? and referer = ?", Severity_, Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLinkAndHostname Get Watchdogs via SeverityAndLinkAndHostname
func GetWatchdogsBySeverityAndLinkAndHostname(offset int, limit int, Severity_ int, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and link = ? and hostname = ?", Severity_, Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLinkAndTimestamp Get Watchdogs via SeverityAndLinkAndTimestamp
func GetWatchdogsBySeverityAndLinkAndTimestamp(offset int, limit int, Severity_ int, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and link = ? and timestamp = ?", Severity_, Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLocationAndReferer Get Watchdogs via SeverityAndLocationAndReferer
func GetWatchdogsBySeverityAndLocationAndReferer(offset int, limit int, Severity_ int, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and location = ? and referer = ?", Severity_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLocationAndHostname Get Watchdogs via SeverityAndLocationAndHostname
func GetWatchdogsBySeverityAndLocationAndHostname(offset int, limit int, Severity_ int, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and location = ? and hostname = ?", Severity_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLocationAndTimestamp Get Watchdogs via SeverityAndLocationAndTimestamp
func GetWatchdogsBySeverityAndLocationAndTimestamp(offset int, limit int, Severity_ int, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and location = ? and timestamp = ?", Severity_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndRefererAndHostname Get Watchdogs via SeverityAndRefererAndHostname
func GetWatchdogsBySeverityAndRefererAndHostname(offset int, limit int, Severity_ int, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and referer = ? and hostname = ?", Severity_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndRefererAndTimestamp Get Watchdogs via SeverityAndRefererAndTimestamp
func GetWatchdogsBySeverityAndRefererAndTimestamp(offset int, limit int, Severity_ int, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and referer = ? and timestamp = ?", Severity_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndHostnameAndTimestamp Get Watchdogs via SeverityAndHostnameAndTimestamp
func GetWatchdogsBySeverityAndHostnameAndTimestamp(offset int, limit int, Severity_ int, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and hostname = ? and timestamp = ?", Severity_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndLocationAndReferer Get Watchdogs via LinkAndLocationAndReferer
func GetWatchdogsByLinkAndLocationAndReferer(offset int, limit int, Link_ string, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and location = ? and referer = ?", Link_, Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndLocationAndHostname Get Watchdogs via LinkAndLocationAndHostname
func GetWatchdogsByLinkAndLocationAndHostname(offset int, limit int, Link_ string, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and location = ? and hostname = ?", Link_, Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndLocationAndTimestamp Get Watchdogs via LinkAndLocationAndTimestamp
func GetWatchdogsByLinkAndLocationAndTimestamp(offset int, limit int, Link_ string, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and location = ? and timestamp = ?", Link_, Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndRefererAndHostname Get Watchdogs via LinkAndRefererAndHostname
func GetWatchdogsByLinkAndRefererAndHostname(offset int, limit int, Link_ string, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and referer = ? and hostname = ?", Link_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndRefererAndTimestamp Get Watchdogs via LinkAndRefererAndTimestamp
func GetWatchdogsByLinkAndRefererAndTimestamp(offset int, limit int, Link_ string, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and referer = ? and timestamp = ?", Link_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndHostnameAndTimestamp Get Watchdogs via LinkAndHostnameAndTimestamp
func GetWatchdogsByLinkAndHostnameAndTimestamp(offset int, limit int, Link_ string, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and hostname = ? and timestamp = ?", Link_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndRefererAndHostname Get Watchdogs via LocationAndRefererAndHostname
func GetWatchdogsByLocationAndRefererAndHostname(offset int, limit int, Location_ string, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and referer = ? and hostname = ?", Location_, Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndRefererAndTimestamp Get Watchdogs via LocationAndRefererAndTimestamp
func GetWatchdogsByLocationAndRefererAndTimestamp(offset int, limit int, Location_ string, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and referer = ? and timestamp = ?", Location_, Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndHostnameAndTimestamp Get Watchdogs via LocationAndHostnameAndTimestamp
func GetWatchdogsByLocationAndHostnameAndTimestamp(offset int, limit int, Location_ string, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and hostname = ? and timestamp = ?", Location_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByRefererAndHostnameAndTimestamp Get Watchdogs via RefererAndHostnameAndTimestamp
func GetWatchdogsByRefererAndHostnameAndTimestamp(offset int, limit int, Referer_ string, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("referer = ? and hostname = ? and timestamp = ?", Referer_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndUid Get Watchdogs via WidAndUid
func GetWatchdogsByWidAndUid(offset int, limit int, Wid_ int64, Uid_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and uid = ?", Wid_, Uid_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndType Get Watchdogs via WidAndType
func GetWatchdogsByWidAndType(offset int, limit int, Wid_ int64, Type_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and type = ?", Wid_, Type_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndMessage Get Watchdogs via WidAndMessage
func GetWatchdogsByWidAndMessage(offset int, limit int, Wid_ int64, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and message = ?", Wid_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndVariables Get Watchdogs via WidAndVariables
func GetWatchdogsByWidAndVariables(offset int, limit int, Wid_ int64, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and variables = ?", Wid_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndSeverity Get Watchdogs via WidAndSeverity
func GetWatchdogsByWidAndSeverity(offset int, limit int, Wid_ int64, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and severity = ?", Wid_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLink Get Watchdogs via WidAndLink
func GetWatchdogsByWidAndLink(offset int, limit int, Wid_ int64, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and link = ?", Wid_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndLocation Get Watchdogs via WidAndLocation
func GetWatchdogsByWidAndLocation(offset int, limit int, Wid_ int64, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and location = ?", Wid_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndReferer Get Watchdogs via WidAndReferer
func GetWatchdogsByWidAndReferer(offset int, limit int, Wid_ int64, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and referer = ?", Wid_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndHostname Get Watchdogs via WidAndHostname
func GetWatchdogsByWidAndHostname(offset int, limit int, Wid_ int64, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and hostname = ?", Wid_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByWidAndTimestamp Get Watchdogs via WidAndTimestamp
func GetWatchdogsByWidAndTimestamp(offset int, limit int, Wid_ int64, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ? and timestamp = ?", Wid_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndType Get Watchdogs via UidAndType
func GetWatchdogsByUidAndType(offset int, limit int, Uid_ int, Type_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and type = ?", Uid_, Type_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndMessage Get Watchdogs via UidAndMessage
func GetWatchdogsByUidAndMessage(offset int, limit int, Uid_ int, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and message = ?", Uid_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndVariables Get Watchdogs via UidAndVariables
func GetWatchdogsByUidAndVariables(offset int, limit int, Uid_ int, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and variables = ?", Uid_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndSeverity Get Watchdogs via UidAndSeverity
func GetWatchdogsByUidAndSeverity(offset int, limit int, Uid_ int, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and severity = ?", Uid_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLink Get Watchdogs via UidAndLink
func GetWatchdogsByUidAndLink(offset int, limit int, Uid_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and link = ?", Uid_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndLocation Get Watchdogs via UidAndLocation
func GetWatchdogsByUidAndLocation(offset int, limit int, Uid_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and location = ?", Uid_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndReferer Get Watchdogs via UidAndReferer
func GetWatchdogsByUidAndReferer(offset int, limit int, Uid_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and referer = ?", Uid_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndHostname Get Watchdogs via UidAndHostname
func GetWatchdogsByUidAndHostname(offset int, limit int, Uid_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and hostname = ?", Uid_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByUidAndTimestamp Get Watchdogs via UidAndTimestamp
func GetWatchdogsByUidAndTimestamp(offset int, limit int, Uid_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ? and timestamp = ?", Uid_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndMessage Get Watchdogs via TypeAndMessage
func GetWatchdogsByTypeAndMessage(offset int, limit int, Type_ string, Message_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and message = ?", Type_, Message_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndVariables Get Watchdogs via TypeAndVariables
func GetWatchdogsByTypeAndVariables(offset int, limit int, Type_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and variables = ?", Type_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndSeverity Get Watchdogs via TypeAndSeverity
func GetWatchdogsByTypeAndSeverity(offset int, limit int, Type_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and severity = ?", Type_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLink Get Watchdogs via TypeAndLink
func GetWatchdogsByTypeAndLink(offset int, limit int, Type_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and link = ?", Type_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndLocation Get Watchdogs via TypeAndLocation
func GetWatchdogsByTypeAndLocation(offset int, limit int, Type_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and location = ?", Type_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndReferer Get Watchdogs via TypeAndReferer
func GetWatchdogsByTypeAndReferer(offset int, limit int, Type_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and referer = ?", Type_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndHostname Get Watchdogs via TypeAndHostname
func GetWatchdogsByTypeAndHostname(offset int, limit int, Type_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and hostname = ?", Type_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByTypeAndTimestamp Get Watchdogs via TypeAndTimestamp
func GetWatchdogsByTypeAndTimestamp(offset int, limit int, Type_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ? and timestamp = ?", Type_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndVariables Get Watchdogs via MessageAndVariables
func GetWatchdogsByMessageAndVariables(offset int, limit int, Message_ string, Variables_ []byte) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and variables = ?", Message_, Variables_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndSeverity Get Watchdogs via MessageAndSeverity
func GetWatchdogsByMessageAndSeverity(offset int, limit int, Message_ string, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and severity = ?", Message_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLink Get Watchdogs via MessageAndLink
func GetWatchdogsByMessageAndLink(offset int, limit int, Message_ string, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and link = ?", Message_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndLocation Get Watchdogs via MessageAndLocation
func GetWatchdogsByMessageAndLocation(offset int, limit int, Message_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and location = ?", Message_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndReferer Get Watchdogs via MessageAndReferer
func GetWatchdogsByMessageAndReferer(offset int, limit int, Message_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and referer = ?", Message_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndHostname Get Watchdogs via MessageAndHostname
func GetWatchdogsByMessageAndHostname(offset int, limit int, Message_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and hostname = ?", Message_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByMessageAndTimestamp Get Watchdogs via MessageAndTimestamp
func GetWatchdogsByMessageAndTimestamp(offset int, limit int, Message_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ? and timestamp = ?", Message_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndSeverity Get Watchdogs via VariablesAndSeverity
func GetWatchdogsByVariablesAndSeverity(offset int, limit int, Variables_ []byte, Severity_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and severity = ?", Variables_, Severity_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLink Get Watchdogs via VariablesAndLink
func GetWatchdogsByVariablesAndLink(offset int, limit int, Variables_ []byte, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and link = ?", Variables_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndLocation Get Watchdogs via VariablesAndLocation
func GetWatchdogsByVariablesAndLocation(offset int, limit int, Variables_ []byte, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and location = ?", Variables_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndReferer Get Watchdogs via VariablesAndReferer
func GetWatchdogsByVariablesAndReferer(offset int, limit int, Variables_ []byte, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and referer = ?", Variables_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndHostname Get Watchdogs via VariablesAndHostname
func GetWatchdogsByVariablesAndHostname(offset int, limit int, Variables_ []byte, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and hostname = ?", Variables_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByVariablesAndTimestamp Get Watchdogs via VariablesAndTimestamp
func GetWatchdogsByVariablesAndTimestamp(offset int, limit int, Variables_ []byte, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ? and timestamp = ?", Variables_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLink Get Watchdogs via SeverityAndLink
func GetWatchdogsBySeverityAndLink(offset int, limit int, Severity_ int, Link_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and link = ?", Severity_, Link_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndLocation Get Watchdogs via SeverityAndLocation
func GetWatchdogsBySeverityAndLocation(offset int, limit int, Severity_ int, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and location = ?", Severity_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndReferer Get Watchdogs via SeverityAndReferer
func GetWatchdogsBySeverityAndReferer(offset int, limit int, Severity_ int, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and referer = ?", Severity_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndHostname Get Watchdogs via SeverityAndHostname
func GetWatchdogsBySeverityAndHostname(offset int, limit int, Severity_ int, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and hostname = ?", Severity_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsBySeverityAndTimestamp Get Watchdogs via SeverityAndTimestamp
func GetWatchdogsBySeverityAndTimestamp(offset int, limit int, Severity_ int, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ? and timestamp = ?", Severity_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndLocation Get Watchdogs via LinkAndLocation
func GetWatchdogsByLinkAndLocation(offset int, limit int, Link_ string, Location_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and location = ?", Link_, Location_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndReferer Get Watchdogs via LinkAndReferer
func GetWatchdogsByLinkAndReferer(offset int, limit int, Link_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and referer = ?", Link_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndHostname Get Watchdogs via LinkAndHostname
func GetWatchdogsByLinkAndHostname(offset int, limit int, Link_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and hostname = ?", Link_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLinkAndTimestamp Get Watchdogs via LinkAndTimestamp
func GetWatchdogsByLinkAndTimestamp(offset int, limit int, Link_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ? and timestamp = ?", Link_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndReferer Get Watchdogs via LocationAndReferer
func GetWatchdogsByLocationAndReferer(offset int, limit int, Location_ string, Referer_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and referer = ?", Location_, Referer_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndHostname Get Watchdogs via LocationAndHostname
func GetWatchdogsByLocationAndHostname(offset int, limit int, Location_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and hostname = ?", Location_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByLocationAndTimestamp Get Watchdogs via LocationAndTimestamp
func GetWatchdogsByLocationAndTimestamp(offset int, limit int, Location_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ? and timestamp = ?", Location_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByRefererAndHostname Get Watchdogs via RefererAndHostname
func GetWatchdogsByRefererAndHostname(offset int, limit int, Referer_ string, Hostname_ string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("referer = ? and hostname = ?", Referer_, Hostname_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByRefererAndTimestamp Get Watchdogs via RefererAndTimestamp
func GetWatchdogsByRefererAndTimestamp(offset int, limit int, Referer_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("referer = ? and timestamp = ?", Referer_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsByHostnameAndTimestamp Get Watchdogs via HostnameAndTimestamp
func GetWatchdogsByHostnameAndTimestamp(offset int, limit int, Hostname_ string, Timestamp_ int) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("hostname = ? and timestamp = ?", Hostname_, Timestamp_).Limit(limit, offset).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogs Get Watchdogs via field
func GetWatchdogs(offset int, limit int, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaWid Get Watchdogs via Wid
func GetWatchdogsViaWid(offset int, limit int, Wid_ int64, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("wid = ?", Wid_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaUid Get Watchdogs via Uid
func GetWatchdogsViaUid(offset int, limit int, Uid_ int, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaType Get Watchdogs via Type
func GetWatchdogsViaType(offset int, limit int, Type_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaMessage Get Watchdogs via Message
func GetWatchdogsViaMessage(offset int, limit int, Message_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("message = ?", Message_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaVariables Get Watchdogs via Variables
func GetWatchdogsViaVariables(offset int, limit int, Variables_ []byte, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("variables = ?", Variables_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaSeverity Get Watchdogs via Severity
func GetWatchdogsViaSeverity(offset int, limit int, Severity_ int, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("severity = ?", Severity_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaLink Get Watchdogs via Link
func GetWatchdogsViaLink(offset int, limit int, Link_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("link = ?", Link_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaLocation Get Watchdogs via Location
func GetWatchdogsViaLocation(offset int, limit int, Location_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("location = ?", Location_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaReferer Get Watchdogs via Referer
func GetWatchdogsViaReferer(offset int, limit int, Referer_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("referer = ?", Referer_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaHostname Get Watchdogs via Hostname
func GetWatchdogsViaHostname(offset int, limit int, Hostname_ string, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("hostname = ?", Hostname_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// GetWatchdogsViaTimestamp Get Watchdogs via Timestamp
func GetWatchdogsViaTimestamp(offset int, limit int, Timestamp_ int, field string) (*[]*Watchdog, error) {
	var _Watchdog = new([]*Watchdog)
	err := Engine.Table("watchdog").Where("timestamp = ?", Timestamp_).Limit(limit, offset).Desc(field).Find(_Watchdog)
	return _Watchdog, err
}

// HasWatchdogViaWid Has Watchdog via Wid
func HasWatchdogViaWid(iWid int64) bool {
	if has, err := Engine.Where("wid = ?", iWid).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaUid Has Watchdog via Uid
func HasWatchdogViaUid(iUid int) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaType Has Watchdog via Type
func HasWatchdogViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaMessage Has Watchdog via Message
func HasWatchdogViaMessage(iMessage string) bool {
	if has, err := Engine.Where("message = ?", iMessage).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaVariables Has Watchdog via Variables
func HasWatchdogViaVariables(iVariables []byte) bool {
	if has, err := Engine.Where("variables = ?", iVariables).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaSeverity Has Watchdog via Severity
func HasWatchdogViaSeverity(iSeverity int) bool {
	if has, err := Engine.Where("severity = ?", iSeverity).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaLink Has Watchdog via Link
func HasWatchdogViaLink(iLink string) bool {
	if has, err := Engine.Where("link = ?", iLink).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaLocation Has Watchdog via Location
func HasWatchdogViaLocation(iLocation string) bool {
	if has, err := Engine.Where("location = ?", iLocation).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaReferer Has Watchdog via Referer
func HasWatchdogViaReferer(iReferer string) bool {
	if has, err := Engine.Where("referer = ?", iReferer).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaHostname Has Watchdog via Hostname
func HasWatchdogViaHostname(iHostname string) bool {
	if has, err := Engine.Where("hostname = ?", iHostname).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasWatchdogViaTimestamp Has Watchdog via Timestamp
func HasWatchdogViaTimestamp(iTimestamp int) bool {
	if has, err := Engine.Where("timestamp = ?", iTimestamp).Get(new(Watchdog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetWatchdogViaWid Get Watchdog via Wid
func GetWatchdogViaWid(iWid int64) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Wid: iWid}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaUid Get Watchdog via Uid
func GetWatchdogViaUid(iUid int) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Uid: iUid}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaType Get Watchdog via Type
func GetWatchdogViaType(iType string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Type: iType}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaMessage Get Watchdog via Message
func GetWatchdogViaMessage(iMessage string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Message: iMessage}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaVariables Get Watchdog via Variables
func GetWatchdogViaVariables(iVariables []byte) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Variables: iVariables}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaSeverity Get Watchdog via Severity
func GetWatchdogViaSeverity(iSeverity int) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Severity: iSeverity}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaLink Get Watchdog via Link
func GetWatchdogViaLink(iLink string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Link: iLink}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaLocation Get Watchdog via Location
func GetWatchdogViaLocation(iLocation string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Location: iLocation}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaReferer Get Watchdog via Referer
func GetWatchdogViaReferer(iReferer string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Referer: iReferer}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaHostname Get Watchdog via Hostname
func GetWatchdogViaHostname(iHostname string) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Hostname: iHostname}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// GetWatchdogViaTimestamp Get Watchdog via Timestamp
func GetWatchdogViaTimestamp(iTimestamp int) (*Watchdog, error) {
	var _Watchdog = &Watchdog{Timestamp: iTimestamp}
	has, err := Engine.Get(_Watchdog)
	if has {
		return _Watchdog, err
	} else {
		return nil, err
	}
}

// SetWatchdogViaWid Set Watchdog via Wid
func SetWatchdogViaWid(iWid int64, watchdog *Watchdog) (int64, error) {
	watchdog.Wid = iWid
	return Engine.Insert(watchdog)
}

// SetWatchdogViaUid Set Watchdog via Uid
func SetWatchdogViaUid(iUid int, watchdog *Watchdog) (int64, error) {
	watchdog.Uid = iUid
	return Engine.Insert(watchdog)
}

// SetWatchdogViaType Set Watchdog via Type
func SetWatchdogViaType(iType string, watchdog *Watchdog) (int64, error) {
	watchdog.Type = iType
	return Engine.Insert(watchdog)
}

// SetWatchdogViaMessage Set Watchdog via Message
func SetWatchdogViaMessage(iMessage string, watchdog *Watchdog) (int64, error) {
	watchdog.Message = iMessage
	return Engine.Insert(watchdog)
}

// SetWatchdogViaVariables Set Watchdog via Variables
func SetWatchdogViaVariables(iVariables []byte, watchdog *Watchdog) (int64, error) {
	watchdog.Variables = iVariables
	return Engine.Insert(watchdog)
}

// SetWatchdogViaSeverity Set Watchdog via Severity
func SetWatchdogViaSeverity(iSeverity int, watchdog *Watchdog) (int64, error) {
	watchdog.Severity = iSeverity
	return Engine.Insert(watchdog)
}

// SetWatchdogViaLink Set Watchdog via Link
func SetWatchdogViaLink(iLink string, watchdog *Watchdog) (int64, error) {
	watchdog.Link = iLink
	return Engine.Insert(watchdog)
}

// SetWatchdogViaLocation Set Watchdog via Location
func SetWatchdogViaLocation(iLocation string, watchdog *Watchdog) (int64, error) {
	watchdog.Location = iLocation
	return Engine.Insert(watchdog)
}

// SetWatchdogViaReferer Set Watchdog via Referer
func SetWatchdogViaReferer(iReferer string, watchdog *Watchdog) (int64, error) {
	watchdog.Referer = iReferer
	return Engine.Insert(watchdog)
}

// SetWatchdogViaHostname Set Watchdog via Hostname
func SetWatchdogViaHostname(iHostname string, watchdog *Watchdog) (int64, error) {
	watchdog.Hostname = iHostname
	return Engine.Insert(watchdog)
}

// SetWatchdogViaTimestamp Set Watchdog via Timestamp
func SetWatchdogViaTimestamp(iTimestamp int, watchdog *Watchdog) (int64, error) {
	watchdog.Timestamp = iTimestamp
	return Engine.Insert(watchdog)
}

// AddWatchdog Add Watchdog via all columns
func AddWatchdog(iWid int64, iUid int, iType string, iMessage string, iVariables []byte, iSeverity int, iLink string, iLocation string, iReferer string, iHostname string, iTimestamp int) error {
	_Watchdog := &Watchdog{Wid: iWid, Uid: iUid, Type: iType, Message: iMessage, Variables: iVariables, Severity: iSeverity, Link: iLink, Location: iLocation, Referer: iReferer, Hostname: iHostname, Timestamp: iTimestamp}
	if _, err := Engine.Insert(_Watchdog); err != nil {
		return err
	} else {
		return nil
	}
}

// PostWatchdog Post Watchdog via iWatchdog
func PostWatchdog(iWatchdog *Watchdog) (int64, error) {
	_, err := Engine.Insert(iWatchdog)
	return iWatchdog.Wid, err
}

// PutWatchdog Put Watchdog
func PutWatchdog(iWatchdog *Watchdog) (int64, error) {
	_, err := Engine.Id(iWatchdog.Wid).Update(iWatchdog)
	return iWatchdog.Wid, err
}

// PutWatchdogViaWid Put Watchdog via Wid
func PutWatchdogViaWid(Wid_ int64, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Wid: Wid_})
	return row, err
}

// PutWatchdogViaUid Put Watchdog via Uid
func PutWatchdogViaUid(Uid_ int, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Uid: Uid_})
	return row, err
}

// PutWatchdogViaType Put Watchdog via Type
func PutWatchdogViaType(Type_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Type: Type_})
	return row, err
}

// PutWatchdogViaMessage Put Watchdog via Message
func PutWatchdogViaMessage(Message_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Message: Message_})
	return row, err
}

// PutWatchdogViaVariables Put Watchdog via Variables
func PutWatchdogViaVariables(Variables_ []byte, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Variables: Variables_})
	return row, err
}

// PutWatchdogViaSeverity Put Watchdog via Severity
func PutWatchdogViaSeverity(Severity_ int, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Severity: Severity_})
	return row, err
}

// PutWatchdogViaLink Put Watchdog via Link
func PutWatchdogViaLink(Link_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Link: Link_})
	return row, err
}

// PutWatchdogViaLocation Put Watchdog via Location
func PutWatchdogViaLocation(Location_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Location: Location_})
	return row, err
}

// PutWatchdogViaReferer Put Watchdog via Referer
func PutWatchdogViaReferer(Referer_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Referer: Referer_})
	return row, err
}

// PutWatchdogViaHostname Put Watchdog via Hostname
func PutWatchdogViaHostname(Hostname_ string, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Hostname: Hostname_})
	return row, err
}

// PutWatchdogViaTimestamp Put Watchdog via Timestamp
func PutWatchdogViaTimestamp(Timestamp_ int, iWatchdog *Watchdog) (int64, error) {
	row, err := Engine.Update(iWatchdog, &Watchdog{Timestamp: Timestamp_})
	return row, err
}

// UpdateWatchdogViaWid via map[string]interface{}{}
func UpdateWatchdogViaWid(iWid int64, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("wid = ?", iWid).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaUid via map[string]interface{}{}
func UpdateWatchdogViaUid(iUid int, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("uid = ?", iUid).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaType via map[string]interface{}{}
func UpdateWatchdogViaType(iType string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("type = ?", iType).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaMessage via map[string]interface{}{}
func UpdateWatchdogViaMessage(iMessage string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("message = ?", iMessage).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaVariables via map[string]interface{}{}
func UpdateWatchdogViaVariables(iVariables []byte, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("variables = ?", iVariables).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaSeverity via map[string]interface{}{}
func UpdateWatchdogViaSeverity(iSeverity int, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("severity = ?", iSeverity).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaLink via map[string]interface{}{}
func UpdateWatchdogViaLink(iLink string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("link = ?", iLink).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaLocation via map[string]interface{}{}
func UpdateWatchdogViaLocation(iLocation string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("location = ?", iLocation).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaReferer via map[string]interface{}{}
func UpdateWatchdogViaReferer(iReferer string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("referer = ?", iReferer).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaHostname via map[string]interface{}{}
func UpdateWatchdogViaHostname(iHostname string, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("hostname = ?", iHostname).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateWatchdogViaTimestamp via map[string]interface{}{}
func UpdateWatchdogViaTimestamp(iTimestamp int, iWatchdogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Watchdog)).Where("timestamp = ?", iTimestamp).Update(iWatchdogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteWatchdog Delete Watchdog
func DeleteWatchdog(iWid int64) (int64, error) {
	row, err := Engine.Id(iWid).Delete(new(Watchdog))
	return row, err
}

// DeleteWatchdogViaWid Delete Watchdog via Wid
func DeleteWatchdogViaWid(iWid int64) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Wid: iWid}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("wid = ?", iWid).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaUid Delete Watchdog via Uid
func DeleteWatchdogViaUid(iUid int) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Uid: iUid}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaType Delete Watchdog via Type
func DeleteWatchdogViaType(iType string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Type: iType}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaMessage Delete Watchdog via Message
func DeleteWatchdogViaMessage(iMessage string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Message: iMessage}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("message = ?", iMessage).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaVariables Delete Watchdog via Variables
func DeleteWatchdogViaVariables(iVariables []byte) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Variables: iVariables}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("variables = ?", iVariables).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaSeverity Delete Watchdog via Severity
func DeleteWatchdogViaSeverity(iSeverity int) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Severity: iSeverity}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("severity = ?", iSeverity).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaLink Delete Watchdog via Link
func DeleteWatchdogViaLink(iLink string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Link: iLink}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("link = ?", iLink).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaLocation Delete Watchdog via Location
func DeleteWatchdogViaLocation(iLocation string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Location: iLocation}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("location = ?", iLocation).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaReferer Delete Watchdog via Referer
func DeleteWatchdogViaReferer(iReferer string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Referer: iReferer}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("referer = ?", iReferer).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaHostname Delete Watchdog via Hostname
func DeleteWatchdogViaHostname(iHostname string) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Hostname: iHostname}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("hostname = ?", iHostname).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteWatchdogViaTimestamp Delete Watchdog via Timestamp
func DeleteWatchdogViaTimestamp(iTimestamp int) (err error) {
	var has bool
	var _Watchdog = &Watchdog{Timestamp: iTimestamp}
	if has, err = Engine.Get(_Watchdog); (has == true) && (err == nil) {
		if row, err := Engine.Where("timestamp = ?", iTimestamp).Delete(new(Watchdog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
