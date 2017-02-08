package model

type ShortcutFieldData struct {
	Id              int64  `xorm:"not null pk index(shortcut__id__default_langcode__langcode) INT(10)"`
	ShortcutSet     string `xorm:"not null index VARCHAR(32)"`
	Langcode        string `xorm:"not null pk index(shortcut__id__default_langcode__langcode) VARCHAR(12)"`
	Title           string `xorm:"VARCHAR(255)"`
	Weight          int    `xorm:"INT(11)"`
	Link_uri        string `xorm:"index VARCHAR(2048)"`
	Link_title      string `xorm:"VARCHAR(255)"`
	Link_options    []byte `xorm:"LONGBLOB"`
	DefaultLangcode int    `xorm:"not null index(shortcut__id__default_langcode__langcode) TINYINT(4)"`
}

// GetShortcutFieldDatasCount ShortcutFieldDatas' Count
func GetShortcutFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&ShortcutFieldData{})
	return total, err
}

// GetShortcutFieldDataCountViaId Get ShortcutFieldData via Id
func GetShortcutFieldDataCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&ShortcutFieldData{Id: iId})
	return n
}

// GetShortcutFieldDataCountViaShortcutSet Get ShortcutFieldData via ShortcutSet
func GetShortcutFieldDataCountViaShortcutSet(iShortcutSet string) int64 {
	n, _ := Engine.Where("shortcut_set = ?", iShortcutSet).Count(&ShortcutFieldData{ShortcutSet: iShortcutSet})
	return n
}

// GetShortcutFieldDataCountViaLangcode Get ShortcutFieldData via Langcode
func GetShortcutFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&ShortcutFieldData{Langcode: iLangcode})
	return n
}

// GetShortcutFieldDataCountViaTitle Get ShortcutFieldData via Title
func GetShortcutFieldDataCountViaTitle(iTitle string) int64 {
	n, _ := Engine.Where("title = ?", iTitle).Count(&ShortcutFieldData{Title: iTitle})
	return n
}

// GetShortcutFieldDataCountViaWeight Get ShortcutFieldData via Weight
func GetShortcutFieldDataCountViaWeight(iWeight int) int64 {
	n, _ := Engine.Where("weight = ?", iWeight).Count(&ShortcutFieldData{Weight: iWeight})
	return n
}

// GetShortcutFieldDataCountViaLink_uri Get ShortcutFieldData via Link_uri
func GetShortcutFieldDataCountViaLink_uri(iLink_uri string) int64 {
	n, _ := Engine.Where("link__uri = ?", iLink_uri).Count(&ShortcutFieldData{Link_uri: iLink_uri})
	return n
}

// GetShortcutFieldDataCountViaLink_title Get ShortcutFieldData via Link_title
func GetShortcutFieldDataCountViaLink_title(iLink_title string) int64 {
	n, _ := Engine.Where("link__title = ?", iLink_title).Count(&ShortcutFieldData{Link_title: iLink_title})
	return n
}

// GetShortcutFieldDataCountViaLink_options Get ShortcutFieldData via Link_options
func GetShortcutFieldDataCountViaLink_options(iLink_options []byte) int64 {
	n, _ := Engine.Where("link__options = ?", iLink_options).Count(&ShortcutFieldData{Link_options: iLink_options})
	return n
}

// GetShortcutFieldDataCountViaDefaultLangcode Get ShortcutFieldData via DefaultLangcode
func GetShortcutFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&ShortcutFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetShortcutFieldDatasByIdAndShortcutSetAndLangcode Get ShortcutFieldDatas via IdAndShortcutSetAndLangcode
func GetShortcutFieldDatasByIdAndShortcutSetAndLangcode(offset int, limit int, Id_ int64, ShortcutSet_ string, Langcode_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and langcode = ?", Id_, ShortcutSet_, Langcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndTitle Get ShortcutFieldDatas via IdAndShortcutSetAndTitle
func GetShortcutFieldDatasByIdAndShortcutSetAndTitle(offset int, limit int, Id_ int64, ShortcutSet_ string, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and title = ?", Id_, ShortcutSet_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndWeight Get ShortcutFieldDatas via IdAndShortcutSetAndWeight
func GetShortcutFieldDatasByIdAndShortcutSetAndWeight(offset int, limit int, Id_ int64, ShortcutSet_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and weight = ?", Id_, ShortcutSet_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndLink_uri Get ShortcutFieldDatas via IdAndShortcutSetAndLink_uri
func GetShortcutFieldDatasByIdAndShortcutSetAndLink_uri(offset int, limit int, Id_ int64, ShortcutSet_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and link_uri = ?", Id_, ShortcutSet_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndLink_title Get ShortcutFieldDatas via IdAndShortcutSetAndLink_title
func GetShortcutFieldDatasByIdAndShortcutSetAndLink_title(offset int, limit int, Id_ int64, ShortcutSet_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and link_title = ?", Id_, ShortcutSet_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndLink_options Get ShortcutFieldDatas via IdAndShortcutSetAndLink_options
func GetShortcutFieldDatasByIdAndShortcutSetAndLink_options(offset int, limit int, Id_ int64, ShortcutSet_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and link_options = ?", Id_, ShortcutSet_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSetAndDefaultLangcode Get ShortcutFieldDatas via IdAndShortcutSetAndDefaultLangcode
func GetShortcutFieldDatasByIdAndShortcutSetAndDefaultLangcode(offset int, limit int, Id_ int64, ShortcutSet_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ? and default_langcode = ?", Id_, ShortcutSet_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndTitle Get ShortcutFieldDatas via IdAndLangcodeAndTitle
func GetShortcutFieldDatasByIdAndLangcodeAndTitle(offset int, limit int, Id_ int64, Langcode_ string, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and title = ?", Id_, Langcode_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndWeight Get ShortcutFieldDatas via IdAndLangcodeAndWeight
func GetShortcutFieldDatasByIdAndLangcodeAndWeight(offset int, limit int, Id_ int64, Langcode_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and weight = ?", Id_, Langcode_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndLink_uri Get ShortcutFieldDatas via IdAndLangcodeAndLink_uri
func GetShortcutFieldDatasByIdAndLangcodeAndLink_uri(offset int, limit int, Id_ int64, Langcode_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and link_uri = ?", Id_, Langcode_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndLink_title Get ShortcutFieldDatas via IdAndLangcodeAndLink_title
func GetShortcutFieldDatasByIdAndLangcodeAndLink_title(offset int, limit int, Id_ int64, Langcode_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and link_title = ?", Id_, Langcode_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndLink_options Get ShortcutFieldDatas via IdAndLangcodeAndLink_options
func GetShortcutFieldDatasByIdAndLangcodeAndLink_options(offset int, limit int, Id_ int64, Langcode_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and link_options = ?", Id_, Langcode_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcodeAndDefaultLangcode Get ShortcutFieldDatas via IdAndLangcodeAndDefaultLangcode
func GetShortcutFieldDatasByIdAndLangcodeAndDefaultLangcode(offset int, limit int, Id_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ? and default_langcode = ?", Id_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitleAndWeight Get ShortcutFieldDatas via IdAndTitleAndWeight
func GetShortcutFieldDatasByIdAndTitleAndWeight(offset int, limit int, Id_ int64, Title_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ? and weight = ?", Id_, Title_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitleAndLink_uri Get ShortcutFieldDatas via IdAndTitleAndLink_uri
func GetShortcutFieldDatasByIdAndTitleAndLink_uri(offset int, limit int, Id_ int64, Title_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ? and link_uri = ?", Id_, Title_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitleAndLink_title Get ShortcutFieldDatas via IdAndTitleAndLink_title
func GetShortcutFieldDatasByIdAndTitleAndLink_title(offset int, limit int, Id_ int64, Title_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ? and link_title = ?", Id_, Title_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitleAndLink_options Get ShortcutFieldDatas via IdAndTitleAndLink_options
func GetShortcutFieldDatasByIdAndTitleAndLink_options(offset int, limit int, Id_ int64, Title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ? and link_options = ?", Id_, Title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitleAndDefaultLangcode Get ShortcutFieldDatas via IdAndTitleAndDefaultLangcode
func GetShortcutFieldDatasByIdAndTitleAndDefaultLangcode(offset int, limit int, Id_ int64, Title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ? and default_langcode = ?", Id_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndWeightAndLink_uri Get ShortcutFieldDatas via IdAndWeightAndLink_uri
func GetShortcutFieldDatasByIdAndWeightAndLink_uri(offset int, limit int, Id_ int64, Weight_ int, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and weight = ? and link_uri = ?", Id_, Weight_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndWeightAndLink_title Get ShortcutFieldDatas via IdAndWeightAndLink_title
func GetShortcutFieldDatasByIdAndWeightAndLink_title(offset int, limit int, Id_ int64, Weight_ int, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and weight = ? and link_title = ?", Id_, Weight_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndWeightAndLink_options Get ShortcutFieldDatas via IdAndWeightAndLink_options
func GetShortcutFieldDatasByIdAndWeightAndLink_options(offset int, limit int, Id_ int64, Weight_ int, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and weight = ? and link_options = ?", Id_, Weight_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndWeightAndDefaultLangcode Get ShortcutFieldDatas via IdAndWeightAndDefaultLangcode
func GetShortcutFieldDatasByIdAndWeightAndDefaultLangcode(offset int, limit int, Id_ int64, Weight_ int, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and weight = ? and default_langcode = ?", Id_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_uriAndLink_title Get ShortcutFieldDatas via IdAndLink_uriAndLink_title
func GetShortcutFieldDatasByIdAndLink_uriAndLink_title(offset int, limit int, Id_ int64, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_uri = ? and link_title = ?", Id_, Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_uriAndLink_options Get ShortcutFieldDatas via IdAndLink_uriAndLink_options
func GetShortcutFieldDatasByIdAndLink_uriAndLink_options(offset int, limit int, Id_ int64, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_uri = ? and link_options = ?", Id_, Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_uriAndDefaultLangcode Get ShortcutFieldDatas via IdAndLink_uriAndDefaultLangcode
func GetShortcutFieldDatasByIdAndLink_uriAndDefaultLangcode(offset int, limit int, Id_ int64, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_uri = ? and default_langcode = ?", Id_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_titleAndLink_options Get ShortcutFieldDatas via IdAndLink_titleAndLink_options
func GetShortcutFieldDatasByIdAndLink_titleAndLink_options(offset int, limit int, Id_ int64, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_title = ? and link_options = ?", Id_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via IdAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByIdAndLink_titleAndDefaultLangcode(offset int, limit int, Id_ int64, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_title = ? and default_langcode = ?", Id_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via IdAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByIdAndLink_optionsAndDefaultLangcode(offset int, limit int, Id_ int64, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_options = ? and default_langcode = ?", Id_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndTitle Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndTitle
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndTitle(offset int, limit int, ShortcutSet_ string, Langcode_ string, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and title = ?", ShortcutSet_, Langcode_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndWeight Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndWeight
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndWeight(offset int, limit int, ShortcutSet_ string, Langcode_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and weight = ?", ShortcutSet_, Langcode_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_uri Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndLink_uri
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_uri(offset int, limit int, ShortcutSet_ string, Langcode_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and link_uri = ?", ShortcutSet_, Langcode_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_title Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndLink_title
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_title(offset int, limit int, ShortcutSet_ string, Langcode_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and link_title = ?", ShortcutSet_, Langcode_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_options Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndLink_options
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndLink_options(offset int, limit int, ShortcutSet_ string, Langcode_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and link_options = ?", ShortcutSet_, Langcode_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcodeAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndLangcodeAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndLangcodeAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Langcode_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ? and default_langcode = ?", ShortcutSet_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitleAndWeight Get ShortcutFieldDatas via ShortcutSetAndTitleAndWeight
func GetShortcutFieldDatasByShortcutSetAndTitleAndWeight(offset int, limit int, ShortcutSet_ string, Title_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ? and weight = ?", ShortcutSet_, Title_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitleAndLink_uri Get ShortcutFieldDatas via ShortcutSetAndTitleAndLink_uri
func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_uri(offset int, limit int, ShortcutSet_ string, Title_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ? and link_uri = ?", ShortcutSet_, Title_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitleAndLink_title Get ShortcutFieldDatas via ShortcutSetAndTitleAndLink_title
func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_title(offset int, limit int, ShortcutSet_ string, Title_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ? and link_title = ?", ShortcutSet_, Title_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitleAndLink_options Get ShortcutFieldDatas via ShortcutSetAndTitleAndLink_options
func GetShortcutFieldDatasByShortcutSetAndTitleAndLink_options(offset int, limit int, ShortcutSet_ string, Title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ? and link_options = ?", ShortcutSet_, Title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitleAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndTitleAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndTitleAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ? and default_langcode = ?", ShortcutSet_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndWeightAndLink_uri Get ShortcutFieldDatas via ShortcutSetAndWeightAndLink_uri
func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_uri(offset int, limit int, ShortcutSet_ string, Weight_ int, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and weight = ? and link_uri = ?", ShortcutSet_, Weight_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndWeightAndLink_title Get ShortcutFieldDatas via ShortcutSetAndWeightAndLink_title
func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_title(offset int, limit int, ShortcutSet_ string, Weight_ int, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and weight = ? and link_title = ?", ShortcutSet_, Weight_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndWeightAndLink_options Get ShortcutFieldDatas via ShortcutSetAndWeightAndLink_options
func GetShortcutFieldDatasByShortcutSetAndWeightAndLink_options(offset int, limit int, ShortcutSet_ string, Weight_ int, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and weight = ? and link_options = ?", ShortcutSet_, Weight_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndWeightAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndWeightAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndWeightAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Weight_ int, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and weight = ? and default_langcode = ?", ShortcutSet_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_title Get ShortcutFieldDatas via ShortcutSetAndLink_uriAndLink_title
func GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_title(offset int, limit int, ShortcutSet_ string, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_uri = ? and link_title = ?", ShortcutSet_, Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_options Get ShortcutFieldDatas via ShortcutSetAndLink_uriAndLink_options
func GetShortcutFieldDatasByShortcutSetAndLink_uriAndLink_options(offset int, limit int, ShortcutSet_ string, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_uri = ? and link_options = ?", ShortcutSet_, Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_uriAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndLink_uriAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndLink_uriAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_uri = ? and default_langcode = ?", ShortcutSet_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_titleAndLink_options Get ShortcutFieldDatas via ShortcutSetAndLink_titleAndLink_options
func GetShortcutFieldDatasByShortcutSetAndLink_titleAndLink_options(offset int, limit int, ShortcutSet_ string, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_title = ? and link_options = ?", ShortcutSet_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndLink_titleAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_title = ? and default_langcode = ?", ShortcutSet_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndLink_optionsAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_options = ? and default_langcode = ?", ShortcutSet_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitleAndWeight Get ShortcutFieldDatas via LangcodeAndTitleAndWeight
func GetShortcutFieldDatasByLangcodeAndTitleAndWeight(offset int, limit int, Langcode_ string, Title_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ? and weight = ?", Langcode_, Title_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitleAndLink_uri Get ShortcutFieldDatas via LangcodeAndTitleAndLink_uri
func GetShortcutFieldDatasByLangcodeAndTitleAndLink_uri(offset int, limit int, Langcode_ string, Title_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ? and link_uri = ?", Langcode_, Title_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitleAndLink_title Get ShortcutFieldDatas via LangcodeAndTitleAndLink_title
func GetShortcutFieldDatasByLangcodeAndTitleAndLink_title(offset int, limit int, Langcode_ string, Title_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ? and link_title = ?", Langcode_, Title_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitleAndLink_options Get ShortcutFieldDatas via LangcodeAndTitleAndLink_options
func GetShortcutFieldDatasByLangcodeAndTitleAndLink_options(offset int, limit int, Langcode_ string, Title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ? and link_options = ?", Langcode_, Title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitleAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndTitleAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndTitleAndDefaultLangcode(offset int, limit int, Langcode_ string, Title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ? and default_langcode = ?", Langcode_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndWeightAndLink_uri Get ShortcutFieldDatas via LangcodeAndWeightAndLink_uri
func GetShortcutFieldDatasByLangcodeAndWeightAndLink_uri(offset int, limit int, Langcode_ string, Weight_ int, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and weight = ? and link_uri = ?", Langcode_, Weight_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndWeightAndLink_title Get ShortcutFieldDatas via LangcodeAndWeightAndLink_title
func GetShortcutFieldDatasByLangcodeAndWeightAndLink_title(offset int, limit int, Langcode_ string, Weight_ int, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and weight = ? and link_title = ?", Langcode_, Weight_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndWeightAndLink_options Get ShortcutFieldDatas via LangcodeAndWeightAndLink_options
func GetShortcutFieldDatasByLangcodeAndWeightAndLink_options(offset int, limit int, Langcode_ string, Weight_ int, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and weight = ? and link_options = ?", Langcode_, Weight_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndWeightAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndWeightAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndWeightAndDefaultLangcode(offset int, limit int, Langcode_ string, Weight_ int, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and weight = ? and default_langcode = ?", Langcode_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_title Get ShortcutFieldDatas via LangcodeAndLink_uriAndLink_title
func GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_title(offset int, limit int, Langcode_ string, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_uri = ? and link_title = ?", Langcode_, Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_options Get ShortcutFieldDatas via LangcodeAndLink_uriAndLink_options
func GetShortcutFieldDatasByLangcodeAndLink_uriAndLink_options(offset int, limit int, Langcode_ string, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_uri = ? and link_options = ?", Langcode_, Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_uriAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndLink_uriAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndLink_uriAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_uri = ? and default_langcode = ?", Langcode_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_titleAndLink_options Get ShortcutFieldDatas via LangcodeAndLink_titleAndLink_options
func GetShortcutFieldDatasByLangcodeAndLink_titleAndLink_options(offset int, limit int, Langcode_ string, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_title = ? and link_options = ?", Langcode_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndLink_titleAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_title = ? and default_langcode = ?", Langcode_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndLink_optionsAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_options = ? and default_langcode = ?", Langcode_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndWeightAndLink_uri Get ShortcutFieldDatas via TitleAndWeightAndLink_uri
func GetShortcutFieldDatasByTitleAndWeightAndLink_uri(offset int, limit int, Title_ string, Weight_ int, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and weight = ? and link_uri = ?", Title_, Weight_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndWeightAndLink_title Get ShortcutFieldDatas via TitleAndWeightAndLink_title
func GetShortcutFieldDatasByTitleAndWeightAndLink_title(offset int, limit int, Title_ string, Weight_ int, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and weight = ? and link_title = ?", Title_, Weight_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndWeightAndLink_options Get ShortcutFieldDatas via TitleAndWeightAndLink_options
func GetShortcutFieldDatasByTitleAndWeightAndLink_options(offset int, limit int, Title_ string, Weight_ int, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and weight = ? and link_options = ?", Title_, Weight_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndWeightAndDefaultLangcode Get ShortcutFieldDatas via TitleAndWeightAndDefaultLangcode
func GetShortcutFieldDatasByTitleAndWeightAndDefaultLangcode(offset int, limit int, Title_ string, Weight_ int, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and weight = ? and default_langcode = ?", Title_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_uriAndLink_title Get ShortcutFieldDatas via TitleAndLink_uriAndLink_title
func GetShortcutFieldDatasByTitleAndLink_uriAndLink_title(offset int, limit int, Title_ string, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_uri = ? and link_title = ?", Title_, Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_uriAndLink_options Get ShortcutFieldDatas via TitleAndLink_uriAndLink_options
func GetShortcutFieldDatasByTitleAndLink_uriAndLink_options(offset int, limit int, Title_ string, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_uri = ? and link_options = ?", Title_, Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_uriAndDefaultLangcode Get ShortcutFieldDatas via TitleAndLink_uriAndDefaultLangcode
func GetShortcutFieldDatasByTitleAndLink_uriAndDefaultLangcode(offset int, limit int, Title_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_uri = ? and default_langcode = ?", Title_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_titleAndLink_options Get ShortcutFieldDatas via TitleAndLink_titleAndLink_options
func GetShortcutFieldDatasByTitleAndLink_titleAndLink_options(offset int, limit int, Title_ string, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_title = ? and link_options = ?", Title_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via TitleAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByTitleAndLink_titleAndDefaultLangcode(offset int, limit int, Title_ string, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_title = ? and default_langcode = ?", Title_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via TitleAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByTitleAndLink_optionsAndDefaultLangcode(offset int, limit int, Title_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_options = ? and default_langcode = ?", Title_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_uriAndLink_title Get ShortcutFieldDatas via WeightAndLink_uriAndLink_title
func GetShortcutFieldDatasByWeightAndLink_uriAndLink_title(offset int, limit int, Weight_ int, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_uri = ? and link_title = ?", Weight_, Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_uriAndLink_options Get ShortcutFieldDatas via WeightAndLink_uriAndLink_options
func GetShortcutFieldDatasByWeightAndLink_uriAndLink_options(offset int, limit int, Weight_ int, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_uri = ? and link_options = ?", Weight_, Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_uriAndDefaultLangcode Get ShortcutFieldDatas via WeightAndLink_uriAndDefaultLangcode
func GetShortcutFieldDatasByWeightAndLink_uriAndDefaultLangcode(offset int, limit int, Weight_ int, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_uri = ? and default_langcode = ?", Weight_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_titleAndLink_options Get ShortcutFieldDatas via WeightAndLink_titleAndLink_options
func GetShortcutFieldDatasByWeightAndLink_titleAndLink_options(offset int, limit int, Weight_ int, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_title = ? and link_options = ?", Weight_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via WeightAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByWeightAndLink_titleAndDefaultLangcode(offset int, limit int, Weight_ int, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_title = ? and default_langcode = ?", Weight_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via WeightAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByWeightAndLink_optionsAndDefaultLangcode(offset int, limit int, Weight_ int, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_options = ? and default_langcode = ?", Weight_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndLink_titleAndLink_options Get ShortcutFieldDatas via Link_uriAndLink_titleAndLink_options
func GetShortcutFieldDatasByLink_uriAndLink_titleAndLink_options(offset int, limit int, Link_uri_ string, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and link_title = ? and link_options = ?", Link_uri_, Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndLink_titleAndDefaultLangcode Get ShortcutFieldDatas via Link_uriAndLink_titleAndDefaultLangcode
func GetShortcutFieldDatasByLink_uriAndLink_titleAndDefaultLangcode(offset int, limit int, Link_uri_ string, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and link_title = ? and default_langcode = ?", Link_uri_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via Link_uriAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByLink_uriAndLink_optionsAndDefaultLangcode(offset int, limit int, Link_uri_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and link_options = ? and default_langcode = ?", Link_uri_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_titleAndLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via Link_titleAndLink_optionsAndDefaultLangcode
func GetShortcutFieldDatasByLink_titleAndLink_optionsAndDefaultLangcode(offset int, limit int, Link_title_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_title = ? and link_options = ? and default_langcode = ?", Link_title_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndShortcutSet Get ShortcutFieldDatas via IdAndShortcutSet
func GetShortcutFieldDatasByIdAndShortcutSet(offset int, limit int, Id_ int64, ShortcutSet_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and shortcut_set = ?", Id_, ShortcutSet_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLangcode Get ShortcutFieldDatas via IdAndLangcode
func GetShortcutFieldDatasByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndTitle Get ShortcutFieldDatas via IdAndTitle
func GetShortcutFieldDatasByIdAndTitle(offset int, limit int, Id_ int64, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and title = ?", Id_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndWeight Get ShortcutFieldDatas via IdAndWeight
func GetShortcutFieldDatasByIdAndWeight(offset int, limit int, Id_ int64, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and weight = ?", Id_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_uri Get ShortcutFieldDatas via IdAndLink_uri
func GetShortcutFieldDatasByIdAndLink_uri(offset int, limit int, Id_ int64, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_uri = ?", Id_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_title Get ShortcutFieldDatas via IdAndLink_title
func GetShortcutFieldDatasByIdAndLink_title(offset int, limit int, Id_ int64, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_title = ?", Id_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndLink_options Get ShortcutFieldDatas via IdAndLink_options
func GetShortcutFieldDatasByIdAndLink_options(offset int, limit int, Id_ int64, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and link_options = ?", Id_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByIdAndDefaultLangcode Get ShortcutFieldDatas via IdAndDefaultLangcode
func GetShortcutFieldDatasByIdAndDefaultLangcode(offset int, limit int, Id_ int64, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ? and default_langcode = ?", Id_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLangcode Get ShortcutFieldDatas via ShortcutSetAndLangcode
func GetShortcutFieldDatasByShortcutSetAndLangcode(offset int, limit int, ShortcutSet_ string, Langcode_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and langcode = ?", ShortcutSet_, Langcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndTitle Get ShortcutFieldDatas via ShortcutSetAndTitle
func GetShortcutFieldDatasByShortcutSetAndTitle(offset int, limit int, ShortcutSet_ string, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and title = ?", ShortcutSet_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndWeight Get ShortcutFieldDatas via ShortcutSetAndWeight
func GetShortcutFieldDatasByShortcutSetAndWeight(offset int, limit int, ShortcutSet_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and weight = ?", ShortcutSet_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_uri Get ShortcutFieldDatas via ShortcutSetAndLink_uri
func GetShortcutFieldDatasByShortcutSetAndLink_uri(offset int, limit int, ShortcutSet_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_uri = ?", ShortcutSet_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_title Get ShortcutFieldDatas via ShortcutSetAndLink_title
func GetShortcutFieldDatasByShortcutSetAndLink_title(offset int, limit int, ShortcutSet_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_title = ?", ShortcutSet_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndLink_options Get ShortcutFieldDatas via ShortcutSetAndLink_options
func GetShortcutFieldDatasByShortcutSetAndLink_options(offset int, limit int, ShortcutSet_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and link_options = ?", ShortcutSet_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByShortcutSetAndDefaultLangcode Get ShortcutFieldDatas via ShortcutSetAndDefaultLangcode
func GetShortcutFieldDatasByShortcutSetAndDefaultLangcode(offset int, limit int, ShortcutSet_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ? and default_langcode = ?", ShortcutSet_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndTitle Get ShortcutFieldDatas via LangcodeAndTitle
func GetShortcutFieldDatasByLangcodeAndTitle(offset int, limit int, Langcode_ string, Title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and title = ?", Langcode_, Title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndWeight Get ShortcutFieldDatas via LangcodeAndWeight
func GetShortcutFieldDatasByLangcodeAndWeight(offset int, limit int, Langcode_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and weight = ?", Langcode_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_uri Get ShortcutFieldDatas via LangcodeAndLink_uri
func GetShortcutFieldDatasByLangcodeAndLink_uri(offset int, limit int, Langcode_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_uri = ?", Langcode_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_title Get ShortcutFieldDatas via LangcodeAndLink_title
func GetShortcutFieldDatasByLangcodeAndLink_title(offset int, limit int, Langcode_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_title = ?", Langcode_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndLink_options Get ShortcutFieldDatas via LangcodeAndLink_options
func GetShortcutFieldDatasByLangcodeAndLink_options(offset int, limit int, Langcode_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and link_options = ?", Langcode_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLangcodeAndDefaultLangcode Get ShortcutFieldDatas via LangcodeAndDefaultLangcode
func GetShortcutFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndWeight Get ShortcutFieldDatas via TitleAndWeight
func GetShortcutFieldDatasByTitleAndWeight(offset int, limit int, Title_ string, Weight_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and weight = ?", Title_, Weight_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_uri Get ShortcutFieldDatas via TitleAndLink_uri
func GetShortcutFieldDatasByTitleAndLink_uri(offset int, limit int, Title_ string, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_uri = ?", Title_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_title Get ShortcutFieldDatas via TitleAndLink_title
func GetShortcutFieldDatasByTitleAndLink_title(offset int, limit int, Title_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_title = ?", Title_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndLink_options Get ShortcutFieldDatas via TitleAndLink_options
func GetShortcutFieldDatasByTitleAndLink_options(offset int, limit int, Title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and link_options = ?", Title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByTitleAndDefaultLangcode Get ShortcutFieldDatas via TitleAndDefaultLangcode
func GetShortcutFieldDatasByTitleAndDefaultLangcode(offset int, limit int, Title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ? and default_langcode = ?", Title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_uri Get ShortcutFieldDatas via WeightAndLink_uri
func GetShortcutFieldDatasByWeightAndLink_uri(offset int, limit int, Weight_ int, Link_uri_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_uri = ?", Weight_, Link_uri_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_title Get ShortcutFieldDatas via WeightAndLink_title
func GetShortcutFieldDatasByWeightAndLink_title(offset int, limit int, Weight_ int, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_title = ?", Weight_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndLink_options Get ShortcutFieldDatas via WeightAndLink_options
func GetShortcutFieldDatasByWeightAndLink_options(offset int, limit int, Weight_ int, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and link_options = ?", Weight_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByWeightAndDefaultLangcode Get ShortcutFieldDatas via WeightAndDefaultLangcode
func GetShortcutFieldDatasByWeightAndDefaultLangcode(offset int, limit int, Weight_ int, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ? and default_langcode = ?", Weight_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndLink_title Get ShortcutFieldDatas via Link_uriAndLink_title
func GetShortcutFieldDatasByLink_uriAndLink_title(offset int, limit int, Link_uri_ string, Link_title_ string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and link_title = ?", Link_uri_, Link_title_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndLink_options Get ShortcutFieldDatas via Link_uriAndLink_options
func GetShortcutFieldDatasByLink_uriAndLink_options(offset int, limit int, Link_uri_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and link_options = ?", Link_uri_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_uriAndDefaultLangcode Get ShortcutFieldDatas via Link_uriAndDefaultLangcode
func GetShortcutFieldDatasByLink_uriAndDefaultLangcode(offset int, limit int, Link_uri_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_uri = ? and default_langcode = ?", Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_titleAndLink_options Get ShortcutFieldDatas via Link_titleAndLink_options
func GetShortcutFieldDatasByLink_titleAndLink_options(offset int, limit int, Link_title_ string, Link_options_ []byte) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_title = ? and link_options = ?", Link_title_, Link_options_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_titleAndDefaultLangcode Get ShortcutFieldDatas via Link_titleAndDefaultLangcode
func GetShortcutFieldDatasByLink_titleAndDefaultLangcode(offset int, limit int, Link_title_ string, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_title = ? and default_langcode = ?", Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasByLink_optionsAndDefaultLangcode Get ShortcutFieldDatas via Link_optionsAndDefaultLangcode
func GetShortcutFieldDatasByLink_optionsAndDefaultLangcode(offset int, limit int, Link_options_ []byte, DefaultLangcode_ int) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link_options = ? and default_langcode = ?", Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatas Get ShortcutFieldDatas via field
func GetShortcutFieldDatas(offset int, limit int, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaId Get ShortcutFieldDatas via Id
func GetShortcutFieldDatasViaId(offset int, limit int, Id_ int64, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaShortcutSet Get ShortcutFieldDatas via ShortcutSet
func GetShortcutFieldDatasViaShortcutSet(offset int, limit int, ShortcutSet_ string, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("shortcut_set = ?", ShortcutSet_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaLangcode Get ShortcutFieldDatas via Langcode
func GetShortcutFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaTitle Get ShortcutFieldDatas via Title
func GetShortcutFieldDatasViaTitle(offset int, limit int, Title_ string, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("title = ?", Title_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaWeight Get ShortcutFieldDatas via Weight
func GetShortcutFieldDatasViaWeight(offset int, limit int, Weight_ int, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("weight = ?", Weight_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaLink_uri Get ShortcutFieldDatas via Link_uri
func GetShortcutFieldDatasViaLink_uri(offset int, limit int, Link_uri_ string, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link__uri = ?", Link_uri_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaLink_title Get ShortcutFieldDatas via Link_title
func GetShortcutFieldDatasViaLink_title(offset int, limit int, Link_title_ string, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link__title = ?", Link_title_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaLink_options Get ShortcutFieldDatas via Link_options
func GetShortcutFieldDatasViaLink_options(offset int, limit int, Link_options_ []byte, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("link__options = ?", Link_options_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// GetShortcutFieldDatasViaDefaultLangcode Get ShortcutFieldDatas via DefaultLangcode
func GetShortcutFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*ShortcutFieldData, error) {
	var _ShortcutFieldData = new([]*ShortcutFieldData)
	err := Engine.Table("shortcut_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_ShortcutFieldData)
	return _ShortcutFieldData, err
}

// HasShortcutFieldDataViaId Has ShortcutFieldData via Id
func HasShortcutFieldDataViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaShortcutSet Has ShortcutFieldData via ShortcutSet
func HasShortcutFieldDataViaShortcutSet(iShortcutSet string) bool {
	if has, err := Engine.Where("shortcut_set = ?", iShortcutSet).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaLangcode Has ShortcutFieldData via Langcode
func HasShortcutFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaTitle Has ShortcutFieldData via Title
func HasShortcutFieldDataViaTitle(iTitle string) bool {
	if has, err := Engine.Where("title = ?", iTitle).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaWeight Has ShortcutFieldData via Weight
func HasShortcutFieldDataViaWeight(iWeight int) bool {
	if has, err := Engine.Where("weight = ?", iWeight).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaLink_uri Has ShortcutFieldData via Link_uri
func HasShortcutFieldDataViaLink_uri(iLink_uri string) bool {
	if has, err := Engine.Where("link__uri = ?", iLink_uri).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaLink_title Has ShortcutFieldData via Link_title
func HasShortcutFieldDataViaLink_title(iLink_title string) bool {
	if has, err := Engine.Where("link__title = ?", iLink_title).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaLink_options Has ShortcutFieldData via Link_options
func HasShortcutFieldDataViaLink_options(iLink_options []byte) bool {
	if has, err := Engine.Where("link__options = ?", iLink_options).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasShortcutFieldDataViaDefaultLangcode Has ShortcutFieldData via DefaultLangcode
func HasShortcutFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(ShortcutFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetShortcutFieldDataViaId Get ShortcutFieldData via Id
func GetShortcutFieldDataViaId(iId int64) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Id: iId}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaShortcutSet Get ShortcutFieldData via ShortcutSet
func GetShortcutFieldDataViaShortcutSet(iShortcutSet string) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{ShortcutSet: iShortcutSet}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaLangcode Get ShortcutFieldData via Langcode
func GetShortcutFieldDataViaLangcode(iLangcode string) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaTitle Get ShortcutFieldData via Title
func GetShortcutFieldDataViaTitle(iTitle string) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Title: iTitle}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaWeight Get ShortcutFieldData via Weight
func GetShortcutFieldDataViaWeight(iWeight int) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Weight: iWeight}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaLink_uri Get ShortcutFieldData via Link_uri
func GetShortcutFieldDataViaLink_uri(iLink_uri string) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Link_uri: iLink_uri}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaLink_title Get ShortcutFieldData via Link_title
func GetShortcutFieldDataViaLink_title(iLink_title string) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Link_title: iLink_title}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaLink_options Get ShortcutFieldData via Link_options
func GetShortcutFieldDataViaLink_options(iLink_options []byte) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{Link_options: iLink_options}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// GetShortcutFieldDataViaDefaultLangcode Get ShortcutFieldData via DefaultLangcode
func GetShortcutFieldDataViaDefaultLangcode(iDefaultLangcode int) (*ShortcutFieldData, error) {
	var _ShortcutFieldData = &ShortcutFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_ShortcutFieldData)
	if has {
		return _ShortcutFieldData, err
	} else {
		return nil, err
	}
}

// SetShortcutFieldDataViaId Set ShortcutFieldData via Id
func SetShortcutFieldDataViaId(iId int64, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Id = iId
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaShortcutSet Set ShortcutFieldData via ShortcutSet
func SetShortcutFieldDataViaShortcutSet(iShortcutSet string, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.ShortcutSet = iShortcutSet
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaLangcode Set ShortcutFieldData via Langcode
func SetShortcutFieldDataViaLangcode(iLangcode string, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Langcode = iLangcode
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaTitle Set ShortcutFieldData via Title
func SetShortcutFieldDataViaTitle(iTitle string, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Title = iTitle
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaWeight Set ShortcutFieldData via Weight
func SetShortcutFieldDataViaWeight(iWeight int, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Weight = iWeight
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaLink_uri Set ShortcutFieldData via Link_uri
func SetShortcutFieldDataViaLink_uri(iLink_uri string, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Link_uri = iLink_uri
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaLink_title Set ShortcutFieldData via Link_title
func SetShortcutFieldDataViaLink_title(iLink_title string, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Link_title = iLink_title
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaLink_options Set ShortcutFieldData via Link_options
func SetShortcutFieldDataViaLink_options(iLink_options []byte, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.Link_options = iLink_options
	return Engine.Insert(shortcut_field_data)
}

// SetShortcutFieldDataViaDefaultLangcode Set ShortcutFieldData via DefaultLangcode
func SetShortcutFieldDataViaDefaultLangcode(iDefaultLangcode int, shortcut_field_data *ShortcutFieldData) (int64, error) {
	shortcut_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(shortcut_field_data)
}

// AddShortcutFieldData Add ShortcutFieldData via all columns
func AddShortcutFieldData(iId int64, iShortcutSet string, iLangcode string, iTitle string, iWeight int, iLink_uri string, iLink_title string, iLink_options []byte, iDefaultLangcode int) error {
	_ShortcutFieldData := &ShortcutFieldData{Id: iId, ShortcutSet: iShortcutSet, Langcode: iLangcode, Title: iTitle, Weight: iWeight, Link_uri: iLink_uri, Link_title: iLink_title, Link_options: iLink_options, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_ShortcutFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostShortcutFieldData Post ShortcutFieldData via iShortcutFieldData
func PostShortcutFieldData(iShortcutFieldData *ShortcutFieldData) (int64, error) {
	_, err := Engine.Insert(iShortcutFieldData)
	return iShortcutFieldData.Id, err
}

// PutShortcutFieldData Put ShortcutFieldData
func PutShortcutFieldData(iShortcutFieldData *ShortcutFieldData) (int64, error) {
	_, err := Engine.Id(iShortcutFieldData.Id).Update(iShortcutFieldData)
	return iShortcutFieldData.Id, err
}

// PutShortcutFieldDataViaId Put ShortcutFieldData via Id
func PutShortcutFieldDataViaId(Id_ int64, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Id: Id_})
	return row, err
}

// PutShortcutFieldDataViaShortcutSet Put ShortcutFieldData via ShortcutSet
func PutShortcutFieldDataViaShortcutSet(ShortcutSet_ string, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{ShortcutSet: ShortcutSet_})
	return row, err
}

// PutShortcutFieldDataViaLangcode Put ShortcutFieldData via Langcode
func PutShortcutFieldDataViaLangcode(Langcode_ string, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Langcode: Langcode_})
	return row, err
}

// PutShortcutFieldDataViaTitle Put ShortcutFieldData via Title
func PutShortcutFieldDataViaTitle(Title_ string, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Title: Title_})
	return row, err
}

// PutShortcutFieldDataViaWeight Put ShortcutFieldData via Weight
func PutShortcutFieldDataViaWeight(Weight_ int, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Weight: Weight_})
	return row, err
}

// PutShortcutFieldDataViaLink_uri Put ShortcutFieldData via Link_uri
func PutShortcutFieldDataViaLink_uri(Link_uri_ string, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Link_uri: Link_uri_})
	return row, err
}

// PutShortcutFieldDataViaLink_title Put ShortcutFieldData via Link_title
func PutShortcutFieldDataViaLink_title(Link_title_ string, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Link_title: Link_title_})
	return row, err
}

// PutShortcutFieldDataViaLink_options Put ShortcutFieldData via Link_options
func PutShortcutFieldDataViaLink_options(Link_options_ []byte, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{Link_options: Link_options_})
	return row, err
}

// PutShortcutFieldDataViaDefaultLangcode Put ShortcutFieldData via DefaultLangcode
func PutShortcutFieldDataViaDefaultLangcode(DefaultLangcode_ int, iShortcutFieldData *ShortcutFieldData) (int64, error) {
	row, err := Engine.Update(iShortcutFieldData, &ShortcutFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateShortcutFieldDataViaId via map[string]interface{}{}
func UpdateShortcutFieldDataViaId(iId int64, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("id = ?", iId).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaShortcutSet via map[string]interface{}{}
func UpdateShortcutFieldDataViaShortcutSet(iShortcutSet string, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("shortcut_set = ?", iShortcutSet).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaLangcode via map[string]interface{}{}
func UpdateShortcutFieldDataViaLangcode(iLangcode string, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("langcode = ?", iLangcode).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaTitle via map[string]interface{}{}
func UpdateShortcutFieldDataViaTitle(iTitle string, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("title = ?", iTitle).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaWeight via map[string]interface{}{}
func UpdateShortcutFieldDataViaWeight(iWeight int, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("weight = ?", iWeight).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaLink_uri via map[string]interface{}{}
func UpdateShortcutFieldDataViaLink_uri(iLink_uri string, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("link__uri = ?", iLink_uri).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaLink_title via map[string]interface{}{}
func UpdateShortcutFieldDataViaLink_title(iLink_title string, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("link__title = ?", iLink_title).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaLink_options via map[string]interface{}{}
func UpdateShortcutFieldDataViaLink_options(iLink_options []byte, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("link__options = ?", iLink_options).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateShortcutFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateShortcutFieldDataViaDefaultLangcode(iDefaultLangcode int, iShortcutFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(ShortcutFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iShortcutFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteShortcutFieldData Delete ShortcutFieldData
func DeleteShortcutFieldData(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(ShortcutFieldData))
	return row, err
}

// DeleteShortcutFieldDataViaId Delete ShortcutFieldData via Id
func DeleteShortcutFieldDataViaId(iId int64) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Id: iId}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaShortcutSet Delete ShortcutFieldData via ShortcutSet
func DeleteShortcutFieldDataViaShortcutSet(iShortcutSet string) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{ShortcutSet: iShortcutSet}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("shortcut_set = ?", iShortcutSet).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaLangcode Delete ShortcutFieldData via Langcode
func DeleteShortcutFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaTitle Delete ShortcutFieldData via Title
func DeleteShortcutFieldDataViaTitle(iTitle string) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Title: iTitle}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("title = ?", iTitle).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaWeight Delete ShortcutFieldData via Weight
func DeleteShortcutFieldDataViaWeight(iWeight int) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Weight: iWeight}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("weight = ?", iWeight).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaLink_uri Delete ShortcutFieldData via Link_uri
func DeleteShortcutFieldDataViaLink_uri(iLink_uri string) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Link_uri: iLink_uri}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__uri = ?", iLink_uri).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaLink_title Delete ShortcutFieldData via Link_title
func DeleteShortcutFieldDataViaLink_title(iLink_title string) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Link_title: iLink_title}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__title = ?", iLink_title).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaLink_options Delete ShortcutFieldData via Link_options
func DeleteShortcutFieldDataViaLink_options(iLink_options []byte) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{Link_options: iLink_options}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__options = ?", iLink_options).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteShortcutFieldDataViaDefaultLangcode Delete ShortcutFieldData via DefaultLangcode
func DeleteShortcutFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _ShortcutFieldData = &ShortcutFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_ShortcutFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(ShortcutFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
