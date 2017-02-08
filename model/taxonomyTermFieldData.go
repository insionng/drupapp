package model

type TaxonomyTermFieldData struct {
	Tid                int64  `xorm:"not null pk index(taxonomy_term__id__default_langcode__langcode) INT(10)"`
	Vid                string `xorm:"not null index(taxonomy_term__vid_name) index(taxonomy_term__tree) VARCHAR(32)"`
	Langcode           string `xorm:"not null pk index(taxonomy_term__id__default_langcode__langcode) VARCHAR(12)"`
	Name               string `xorm:"not null index(taxonomy_term__tree) index(taxonomy_term__vid_name) index VARCHAR(255)"`
	Description_value  string `xorm:"LONGTEXT"`
	Description_format string `xorm:"VARCHAR(255)"`
	Weight             int    `xorm:"not null index(taxonomy_term__tree) INT(11)"`
	Changed            int    `xorm:"INT(11)"`
	DefaultLangcode    int    `xorm:"not null index(taxonomy_term__id__default_langcode__langcode) TINYINT(4)"`
}

// GetTaxonomyTermFieldDatasCount TaxonomyTermFieldDatas' Count
func GetTaxonomyTermFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TaxonomyTermFieldData{})
	return total, err
}

// GetTaxonomyTermFieldDataCountViaTid Get TaxonomyTermFieldData via Tid
func GetTaxonomyTermFieldDataCountViaTid(iTid int64) int64 {
	n, _ := Engine.Where("tid = ?", iTid).Count(&TaxonomyTermFieldData{Tid: iTid})
	return n
}

// GetTaxonomyTermFieldDataCountViaVid Get TaxonomyTermFieldData via Vid
func GetTaxonomyTermFieldDataCountViaVid(iVid string) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&TaxonomyTermFieldData{Vid: iVid})
	return n
}

// GetTaxonomyTermFieldDataCountViaLangcode Get TaxonomyTermFieldData via Langcode
func GetTaxonomyTermFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&TaxonomyTermFieldData{Langcode: iLangcode})
	return n
}

// GetTaxonomyTermFieldDataCountViaName Get TaxonomyTermFieldData via Name
func GetTaxonomyTermFieldDataCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&TaxonomyTermFieldData{Name: iName})
	return n
}

// GetTaxonomyTermFieldDataCountViaDescription_value Get TaxonomyTermFieldData via Description_value
func GetTaxonomyTermFieldDataCountViaDescription_value(iDescription_value string) int64 {
	n, _ := Engine.Where("description__value = ?", iDescription_value).Count(&TaxonomyTermFieldData{Description_value: iDescription_value})
	return n
}

// GetTaxonomyTermFieldDataCountViaDescription_format Get TaxonomyTermFieldData via Description_format
func GetTaxonomyTermFieldDataCountViaDescription_format(iDescription_format string) int64 {
	n, _ := Engine.Where("description__format = ?", iDescription_format).Count(&TaxonomyTermFieldData{Description_format: iDescription_format})
	return n
}

// GetTaxonomyTermFieldDataCountViaWeight Get TaxonomyTermFieldData via Weight
func GetTaxonomyTermFieldDataCountViaWeight(iWeight int) int64 {
	n, _ := Engine.Where("weight = ?", iWeight).Count(&TaxonomyTermFieldData{Weight: iWeight})
	return n
}

// GetTaxonomyTermFieldDataCountViaChanged Get TaxonomyTermFieldData via Changed
func GetTaxonomyTermFieldDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&TaxonomyTermFieldData{Changed: iChanged})
	return n
}

// GetTaxonomyTermFieldDataCountViaDefaultLangcode Get TaxonomyTermFieldData via DefaultLangcode
func GetTaxonomyTermFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&TaxonomyTermFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetTaxonomyTermFieldDatasByTidAndVidAndLangcode Get TaxonomyTermFieldDatas via TidAndVidAndLangcode
func GetTaxonomyTermFieldDatasByTidAndVidAndLangcode(offset int, limit int, Tid_ int64, Vid_ string, Langcode_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and langcode = ?", Tid_, Vid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndName Get TaxonomyTermFieldDatas via TidAndVidAndName
func GetTaxonomyTermFieldDatasByTidAndVidAndName(offset int, limit int, Tid_ int64, Vid_ string, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and name = ?", Tid_, Vid_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndDescription_value Get TaxonomyTermFieldDatas via TidAndVidAndDescription_value
func GetTaxonomyTermFieldDatasByTidAndVidAndDescription_value(offset int, limit int, Tid_ int64, Vid_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and description_value = ?", Tid_, Vid_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndDescription_format Get TaxonomyTermFieldDatas via TidAndVidAndDescription_format
func GetTaxonomyTermFieldDatasByTidAndVidAndDescription_format(offset int, limit int, Tid_ int64, Vid_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and description_format = ?", Tid_, Vid_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndWeight Get TaxonomyTermFieldDatas via TidAndVidAndWeight
func GetTaxonomyTermFieldDatasByTidAndVidAndWeight(offset int, limit int, Tid_ int64, Vid_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and weight = ?", Tid_, Vid_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndChanged Get TaxonomyTermFieldDatas via TidAndVidAndChanged
func GetTaxonomyTermFieldDatasByTidAndVidAndChanged(offset int, limit int, Tid_ int64, Vid_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and changed = ?", Tid_, Vid_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVidAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndVidAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndVidAndDefaultLangcode(offset int, limit int, Tid_ int64, Vid_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ? and default_langcode = ?", Tid_, Vid_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndName Get TaxonomyTermFieldDatas via TidAndLangcodeAndName
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndName(offset int, limit int, Tid_ int64, Langcode_ string, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and name = ?", Tid_, Langcode_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_value Get TaxonomyTermFieldDatas via TidAndLangcodeAndDescription_value
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_value(offset int, limit int, Tid_ int64, Langcode_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and description_value = ?", Tid_, Langcode_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_format Get TaxonomyTermFieldDatas via TidAndLangcodeAndDescription_format
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDescription_format(offset int, limit int, Tid_ int64, Langcode_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and description_format = ?", Tid_, Langcode_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndWeight Get TaxonomyTermFieldDatas via TidAndLangcodeAndWeight
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndWeight(offset int, limit int, Tid_ int64, Langcode_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and weight = ?", Tid_, Langcode_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndChanged Get TaxonomyTermFieldDatas via TidAndLangcodeAndChanged
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndChanged(offset int, limit int, Tid_ int64, Langcode_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and changed = ?", Tid_, Langcode_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcodeAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndLangcodeAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndLangcodeAndDefaultLangcode(offset int, limit int, Tid_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ? and default_langcode = ?", Tid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndNameAndDescription_value Get TaxonomyTermFieldDatas via TidAndNameAndDescription_value
func GetTaxonomyTermFieldDatasByTidAndNameAndDescription_value(offset int, limit int, Tid_ int64, Name_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ? and description_value = ?", Tid_, Name_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndNameAndDescription_format Get TaxonomyTermFieldDatas via TidAndNameAndDescription_format
func GetTaxonomyTermFieldDatasByTidAndNameAndDescription_format(offset int, limit int, Tid_ int64, Name_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ? and description_format = ?", Tid_, Name_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndNameAndWeight Get TaxonomyTermFieldDatas via TidAndNameAndWeight
func GetTaxonomyTermFieldDatasByTidAndNameAndWeight(offset int, limit int, Tid_ int64, Name_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ? and weight = ?", Tid_, Name_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndNameAndChanged Get TaxonomyTermFieldDatas via TidAndNameAndChanged
func GetTaxonomyTermFieldDatasByTidAndNameAndChanged(offset int, limit int, Tid_ int64, Name_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ? and changed = ?", Tid_, Name_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndNameAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndNameAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndNameAndDefaultLangcode(offset int, limit int, Tid_ int64, Name_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ? and default_langcode = ?", Tid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDescription_format Get TaxonomyTermFieldDatas via TidAndDescription_valueAndDescription_format
func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDescription_format(offset int, limit int, Tid_ int64, Description_value_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_value = ? and description_format = ?", Tid_, Description_value_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_valueAndWeight Get TaxonomyTermFieldDatas via TidAndDescription_valueAndWeight
func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndWeight(offset int, limit int, Tid_ int64, Description_value_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_value = ? and weight = ?", Tid_, Description_value_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_valueAndChanged Get TaxonomyTermFieldDatas via TidAndDescription_valueAndChanged
func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndChanged(offset int, limit int, Tid_ int64, Description_value_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_value = ? and changed = ?", Tid_, Description_value_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndDescription_valueAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndDescription_valueAndDefaultLangcode(offset int, limit int, Tid_ int64, Description_value_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_value = ? and default_langcode = ?", Tid_, Description_value_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_formatAndWeight Get TaxonomyTermFieldDatas via TidAndDescription_formatAndWeight
func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndWeight(offset int, limit int, Tid_ int64, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_format = ? and weight = ?", Tid_, Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_formatAndChanged Get TaxonomyTermFieldDatas via TidAndDescription_formatAndChanged
func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndChanged(offset int, limit int, Tid_ int64, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_format = ? and changed = ?", Tid_, Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndDescription_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndDescription_formatAndDefaultLangcode(offset int, limit int, Tid_ int64, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_format = ? and default_langcode = ?", Tid_, Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndWeightAndChanged Get TaxonomyTermFieldDatas via TidAndWeightAndChanged
func GetTaxonomyTermFieldDatasByTidAndWeightAndChanged(offset int, limit int, Tid_ int64, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and weight = ? and changed = ?", Tid_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndWeightAndDefaultLangcode(offset int, limit int, Tid_ int64, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and weight = ? and default_langcode = ?", Tid_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndChangedAndDefaultLangcode(offset int, limit int, Tid_ int64, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and changed = ? and default_langcode = ?", Tid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndName Get TaxonomyTermFieldDatas via VidAndLangcodeAndName
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndName(offset int, limit int, Vid_ string, Langcode_ string, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and name = ?", Vid_, Langcode_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_value Get TaxonomyTermFieldDatas via VidAndLangcodeAndDescription_value
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_value(offset int, limit int, Vid_ string, Langcode_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and description_value = ?", Vid_, Langcode_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_format Get TaxonomyTermFieldDatas via VidAndLangcodeAndDescription_format
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDescription_format(offset int, limit int, Vid_ string, Langcode_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and description_format = ?", Vid_, Langcode_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndWeight Get TaxonomyTermFieldDatas via VidAndLangcodeAndWeight
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndWeight(offset int, limit int, Vid_ string, Langcode_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and weight = ?", Vid_, Langcode_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndChanged Get TaxonomyTermFieldDatas via VidAndLangcodeAndChanged
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndChanged(offset int, limit int, Vid_ string, Langcode_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and changed = ?", Vid_, Langcode_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcodeAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndLangcodeAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndLangcodeAndDefaultLangcode(offset int, limit int, Vid_ string, Langcode_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ? and default_langcode = ?", Vid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndNameAndDescription_value Get TaxonomyTermFieldDatas via VidAndNameAndDescription_value
func GetTaxonomyTermFieldDatasByVidAndNameAndDescription_value(offset int, limit int, Vid_ string, Name_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ? and description_value = ?", Vid_, Name_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndNameAndDescription_format Get TaxonomyTermFieldDatas via VidAndNameAndDescription_format
func GetTaxonomyTermFieldDatasByVidAndNameAndDescription_format(offset int, limit int, Vid_ string, Name_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ? and description_format = ?", Vid_, Name_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndNameAndWeight Get TaxonomyTermFieldDatas via VidAndNameAndWeight
func GetTaxonomyTermFieldDatasByVidAndNameAndWeight(offset int, limit int, Vid_ string, Name_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ? and weight = ?", Vid_, Name_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndNameAndChanged Get TaxonomyTermFieldDatas via VidAndNameAndChanged
func GetTaxonomyTermFieldDatasByVidAndNameAndChanged(offset int, limit int, Vid_ string, Name_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ? and changed = ?", Vid_, Name_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndNameAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndNameAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndNameAndDefaultLangcode(offset int, limit int, Vid_ string, Name_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ? and default_langcode = ?", Vid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDescription_format Get TaxonomyTermFieldDatas via VidAndDescription_valueAndDescription_format
func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDescription_format(offset int, limit int, Vid_ string, Description_value_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_value = ? and description_format = ?", Vid_, Description_value_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_valueAndWeight Get TaxonomyTermFieldDatas via VidAndDescription_valueAndWeight
func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndWeight(offset int, limit int, Vid_ string, Description_value_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_value = ? and weight = ?", Vid_, Description_value_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_valueAndChanged Get TaxonomyTermFieldDatas via VidAndDescription_valueAndChanged
func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndChanged(offset int, limit int, Vid_ string, Description_value_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_value = ? and changed = ?", Vid_, Description_value_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndDescription_valueAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndDescription_valueAndDefaultLangcode(offset int, limit int, Vid_ string, Description_value_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_value = ? and default_langcode = ?", Vid_, Description_value_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_formatAndWeight Get TaxonomyTermFieldDatas via VidAndDescription_formatAndWeight
func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndWeight(offset int, limit int, Vid_ string, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_format = ? and weight = ?", Vid_, Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_formatAndChanged Get TaxonomyTermFieldDatas via VidAndDescription_formatAndChanged
func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndChanged(offset int, limit int, Vid_ string, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_format = ? and changed = ?", Vid_, Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndDescription_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndDescription_formatAndDefaultLangcode(offset int, limit int, Vid_ string, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_format = ? and default_langcode = ?", Vid_, Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndWeightAndChanged Get TaxonomyTermFieldDatas via VidAndWeightAndChanged
func GetTaxonomyTermFieldDatasByVidAndWeightAndChanged(offset int, limit int, Vid_ string, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and weight = ? and changed = ?", Vid_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndWeightAndDefaultLangcode(offset int, limit int, Vid_ string, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and weight = ? and default_langcode = ?", Vid_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndChangedAndDefaultLangcode(offset int, limit int, Vid_ string, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and changed = ? and default_langcode = ?", Vid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_value Get TaxonomyTermFieldDatas via LangcodeAndNameAndDescription_value
func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_value(offset int, limit int, Langcode_ string, Name_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ? and description_value = ?", Langcode_, Name_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_format Get TaxonomyTermFieldDatas via LangcodeAndNameAndDescription_format
func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDescription_format(offset int, limit int, Langcode_ string, Name_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ? and description_format = ?", Langcode_, Name_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndNameAndWeight Get TaxonomyTermFieldDatas via LangcodeAndNameAndWeight
func GetTaxonomyTermFieldDatasByLangcodeAndNameAndWeight(offset int, limit int, Langcode_ string, Name_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ? and weight = ?", Langcode_, Name_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndNameAndChanged Get TaxonomyTermFieldDatas via LangcodeAndNameAndChanged
func GetTaxonomyTermFieldDatasByLangcodeAndNameAndChanged(offset int, limit int, Langcode_ string, Name_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ? and changed = ?", Langcode_, Name_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndNameAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndNameAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndNameAndDefaultLangcode(offset int, limit int, Langcode_ string, Name_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ? and default_langcode = ?", Langcode_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDescription_format Get TaxonomyTermFieldDatas via LangcodeAndDescription_valueAndDescription_format
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDescription_format(offset int, limit int, Langcode_ string, Description_value_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_value = ? and description_format = ?", Langcode_, Description_value_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndWeight Get TaxonomyTermFieldDatas via LangcodeAndDescription_valueAndWeight
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndWeight(offset int, limit int, Langcode_ string, Description_value_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_value = ? and weight = ?", Langcode_, Description_value_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndChanged Get TaxonomyTermFieldDatas via LangcodeAndDescription_valueAndChanged
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndChanged(offset int, limit int, Langcode_ string, Description_value_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_value = ? and changed = ?", Langcode_, Description_value_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndDescription_valueAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_valueAndDefaultLangcode(offset int, limit int, Langcode_ string, Description_value_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_value = ? and default_langcode = ?", Langcode_, Description_value_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndWeight Get TaxonomyTermFieldDatas via LangcodeAndDescription_formatAndWeight
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndWeight(offset int, limit int, Langcode_ string, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_format = ? and weight = ?", Langcode_, Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndChanged Get TaxonomyTermFieldDatas via LangcodeAndDescription_formatAndChanged
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndChanged(offset int, limit int, Langcode_ string, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_format = ? and changed = ?", Langcode_, Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndDescription_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_formatAndDefaultLangcode(offset int, limit int, Langcode_ string, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_format = ? and default_langcode = ?", Langcode_, Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndWeightAndChanged Get TaxonomyTermFieldDatas via LangcodeAndWeightAndChanged
func GetTaxonomyTermFieldDatasByLangcodeAndWeightAndChanged(offset int, limit int, Langcode_ string, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and weight = ? and changed = ?", Langcode_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndWeightAndDefaultLangcode(offset int, limit int, Langcode_ string, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and weight = ? and default_langcode = ?", Langcode_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDescription_format Get TaxonomyTermFieldDatas via NameAndDescription_valueAndDescription_format
func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDescription_format(offset int, limit int, Name_ string, Description_value_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_value = ? and description_format = ?", Name_, Description_value_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_valueAndWeight Get TaxonomyTermFieldDatas via NameAndDescription_valueAndWeight
func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndWeight(offset int, limit int, Name_ string, Description_value_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_value = ? and weight = ?", Name_, Description_value_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_valueAndChanged Get TaxonomyTermFieldDatas via NameAndDescription_valueAndChanged
func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndChanged(offset int, limit int, Name_ string, Description_value_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_value = ? and changed = ?", Name_, Description_value_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDefaultLangcode Get TaxonomyTermFieldDatas via NameAndDescription_valueAndDefaultLangcode
func GetTaxonomyTermFieldDatasByNameAndDescription_valueAndDefaultLangcode(offset int, limit int, Name_ string, Description_value_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_value = ? and default_langcode = ?", Name_, Description_value_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_formatAndWeight Get TaxonomyTermFieldDatas via NameAndDescription_formatAndWeight
func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndWeight(offset int, limit int, Name_ string, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_format = ? and weight = ?", Name_, Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_formatAndChanged Get TaxonomyTermFieldDatas via NameAndDescription_formatAndChanged
func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndChanged(offset int, limit int, Name_ string, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_format = ? and changed = ?", Name_, Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via NameAndDescription_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByNameAndDescription_formatAndDefaultLangcode(offset int, limit int, Name_ string, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_format = ? and default_langcode = ?", Name_, Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndWeightAndChanged Get TaxonomyTermFieldDatas via NameAndWeightAndChanged
func GetTaxonomyTermFieldDatasByNameAndWeightAndChanged(offset int, limit int, Name_ string, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and weight = ? and changed = ?", Name_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via NameAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByNameAndWeightAndDefaultLangcode(offset int, limit int, Name_ string, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and weight = ? and default_langcode = ?", Name_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via NameAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByNameAndChangedAndDefaultLangcode(offset int, limit int, Name_ string, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and changed = ? and default_langcode = ?", Name_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndWeight Get TaxonomyTermFieldDatas via Description_valueAndDescription_formatAndWeight
func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndWeight(offset int, limit int, Description_value_ string, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and description_format = ? and weight = ?", Description_value_, Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndChanged Get TaxonomyTermFieldDatas via Description_valueAndDescription_formatAndChanged
func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndChanged(offset int, limit int, Description_value_ string, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and description_format = ? and changed = ?", Description_value_, Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_valueAndDescription_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_formatAndDefaultLangcode(offset int, limit int, Description_value_ string, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and description_format = ? and default_langcode = ?", Description_value_, Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndChanged Get TaxonomyTermFieldDatas via Description_valueAndWeightAndChanged
func GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndChanged(offset int, limit int, Description_value_ string, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and weight = ? and changed = ?", Description_value_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_valueAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_valueAndWeightAndDefaultLangcode(offset int, limit int, Description_value_ string, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and weight = ? and default_langcode = ?", Description_value_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_valueAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_valueAndChangedAndDefaultLangcode(offset int, limit int, Description_value_ string, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and changed = ? and default_langcode = ?", Description_value_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndChanged Get TaxonomyTermFieldDatas via Description_formatAndWeightAndChanged
func GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndChanged(offset int, limit int, Description_format_ string, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and weight = ? and changed = ?", Description_format_, Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_formatAndWeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_formatAndWeightAndDefaultLangcode(offset int, limit int, Description_format_ string, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and weight = ? and default_langcode = ?", Description_format_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_formatAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_formatAndChangedAndDefaultLangcode(offset int, limit int, Description_format_ string, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and changed = ? and default_langcode = ?", Description_format_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByWeightAndChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via WeightAndChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByWeightAndChangedAndDefaultLangcode(offset int, limit int, Weight_ int, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("weight = ? and changed = ? and default_langcode = ?", Weight_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndVid Get TaxonomyTermFieldDatas via TidAndVid
func GetTaxonomyTermFieldDatasByTidAndVid(offset int, limit int, Tid_ int64, Vid_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and vid = ?", Tid_, Vid_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndLangcode Get TaxonomyTermFieldDatas via TidAndLangcode
func GetTaxonomyTermFieldDatasByTidAndLangcode(offset int, limit int, Tid_ int64, Langcode_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and langcode = ?", Tid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndName Get TaxonomyTermFieldDatas via TidAndName
func GetTaxonomyTermFieldDatasByTidAndName(offset int, limit int, Tid_ int64, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and name = ?", Tid_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_value Get TaxonomyTermFieldDatas via TidAndDescription_value
func GetTaxonomyTermFieldDatasByTidAndDescription_value(offset int, limit int, Tid_ int64, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_value = ?", Tid_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDescription_format Get TaxonomyTermFieldDatas via TidAndDescription_format
func GetTaxonomyTermFieldDatasByTidAndDescription_format(offset int, limit int, Tid_ int64, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and description_format = ?", Tid_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndWeight Get TaxonomyTermFieldDatas via TidAndWeight
func GetTaxonomyTermFieldDatasByTidAndWeight(offset int, limit int, Tid_ int64, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and weight = ?", Tid_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndChanged Get TaxonomyTermFieldDatas via TidAndChanged
func GetTaxonomyTermFieldDatasByTidAndChanged(offset int, limit int, Tid_ int64, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and changed = ?", Tid_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByTidAndDefaultLangcode Get TaxonomyTermFieldDatas via TidAndDefaultLangcode
func GetTaxonomyTermFieldDatasByTidAndDefaultLangcode(offset int, limit int, Tid_ int64, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ? and default_langcode = ?", Tid_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndLangcode Get TaxonomyTermFieldDatas via VidAndLangcode
func GetTaxonomyTermFieldDatasByVidAndLangcode(offset int, limit int, Vid_ string, Langcode_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndName Get TaxonomyTermFieldDatas via VidAndName
func GetTaxonomyTermFieldDatasByVidAndName(offset int, limit int, Vid_ string, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and name = ?", Vid_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_value Get TaxonomyTermFieldDatas via VidAndDescription_value
func GetTaxonomyTermFieldDatasByVidAndDescription_value(offset int, limit int, Vid_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_value = ?", Vid_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDescription_format Get TaxonomyTermFieldDatas via VidAndDescription_format
func GetTaxonomyTermFieldDatasByVidAndDescription_format(offset int, limit int, Vid_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and description_format = ?", Vid_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndWeight Get TaxonomyTermFieldDatas via VidAndWeight
func GetTaxonomyTermFieldDatasByVidAndWeight(offset int, limit int, Vid_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and weight = ?", Vid_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndChanged Get TaxonomyTermFieldDatas via VidAndChanged
func GetTaxonomyTermFieldDatasByVidAndChanged(offset int, limit int, Vid_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and changed = ?", Vid_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByVidAndDefaultLangcode Get TaxonomyTermFieldDatas via VidAndDefaultLangcode
func GetTaxonomyTermFieldDatasByVidAndDefaultLangcode(offset int, limit int, Vid_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ? and default_langcode = ?", Vid_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndName Get TaxonomyTermFieldDatas via LangcodeAndName
func GetTaxonomyTermFieldDatasByLangcodeAndName(offset int, limit int, Langcode_ string, Name_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and name = ?", Langcode_, Name_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_value Get TaxonomyTermFieldDatas via LangcodeAndDescription_value
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_value(offset int, limit int, Langcode_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_value = ?", Langcode_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDescription_format Get TaxonomyTermFieldDatas via LangcodeAndDescription_format
func GetTaxonomyTermFieldDatasByLangcodeAndDescription_format(offset int, limit int, Langcode_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and description_format = ?", Langcode_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndWeight Get TaxonomyTermFieldDatas via LangcodeAndWeight
func GetTaxonomyTermFieldDatasByLangcodeAndWeight(offset int, limit int, Langcode_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and weight = ?", Langcode_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndChanged Get TaxonomyTermFieldDatas via LangcodeAndChanged
func GetTaxonomyTermFieldDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByLangcodeAndDefaultLangcode Get TaxonomyTermFieldDatas via LangcodeAndDefaultLangcode
func GetTaxonomyTermFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_value Get TaxonomyTermFieldDatas via NameAndDescription_value
func GetTaxonomyTermFieldDatasByNameAndDescription_value(offset int, limit int, Name_ string, Description_value_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_value = ?", Name_, Description_value_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDescription_format Get TaxonomyTermFieldDatas via NameAndDescription_format
func GetTaxonomyTermFieldDatasByNameAndDescription_format(offset int, limit int, Name_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and description_format = ?", Name_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndWeight Get TaxonomyTermFieldDatas via NameAndWeight
func GetTaxonomyTermFieldDatasByNameAndWeight(offset int, limit int, Name_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and weight = ?", Name_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndChanged Get TaxonomyTermFieldDatas via NameAndChanged
func GetTaxonomyTermFieldDatasByNameAndChanged(offset int, limit int, Name_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and changed = ?", Name_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByNameAndDefaultLangcode Get TaxonomyTermFieldDatas via NameAndDefaultLangcode
func GetTaxonomyTermFieldDatasByNameAndDefaultLangcode(offset int, limit int, Name_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ? and default_langcode = ?", Name_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndDescription_format Get TaxonomyTermFieldDatas via Description_valueAndDescription_format
func GetTaxonomyTermFieldDatasByDescription_valueAndDescription_format(offset int, limit int, Description_value_ string, Description_format_ string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and description_format = ?", Description_value_, Description_format_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndWeight Get TaxonomyTermFieldDatas via Description_valueAndWeight
func GetTaxonomyTermFieldDatasByDescription_valueAndWeight(offset int, limit int, Description_value_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and weight = ?", Description_value_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndChanged Get TaxonomyTermFieldDatas via Description_valueAndChanged
func GetTaxonomyTermFieldDatasByDescription_valueAndChanged(offset int, limit int, Description_value_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and changed = ?", Description_value_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_valueAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_valueAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_valueAndDefaultLangcode(offset int, limit int, Description_value_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_value = ? and default_langcode = ?", Description_value_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndWeight Get TaxonomyTermFieldDatas via Description_formatAndWeight
func GetTaxonomyTermFieldDatasByDescription_formatAndWeight(offset int, limit int, Description_format_ string, Weight_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and weight = ?", Description_format_, Weight_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndChanged Get TaxonomyTermFieldDatas via Description_formatAndChanged
func GetTaxonomyTermFieldDatasByDescription_formatAndChanged(offset int, limit int, Description_format_ string, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and changed = ?", Description_format_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByDescription_formatAndDefaultLangcode Get TaxonomyTermFieldDatas via Description_formatAndDefaultLangcode
func GetTaxonomyTermFieldDatasByDescription_formatAndDefaultLangcode(offset int, limit int, Description_format_ string, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description_format = ? and default_langcode = ?", Description_format_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByWeightAndChanged Get TaxonomyTermFieldDatas via WeightAndChanged
func GetTaxonomyTermFieldDatasByWeightAndChanged(offset int, limit int, Weight_ int, Changed_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("weight = ? and changed = ?", Weight_, Changed_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByWeightAndDefaultLangcode Get TaxonomyTermFieldDatas via WeightAndDefaultLangcode
func GetTaxonomyTermFieldDatasByWeightAndDefaultLangcode(offset int, limit int, Weight_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("weight = ? and default_langcode = ?", Weight_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasByChangedAndDefaultLangcode Get TaxonomyTermFieldDatas via ChangedAndDefaultLangcode
func GetTaxonomyTermFieldDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatas Get TaxonomyTermFieldDatas via field
func GetTaxonomyTermFieldDatas(offset int, limit int, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaTid Get TaxonomyTermFieldDatas via Tid
func GetTaxonomyTermFieldDatasViaTid(offset int, limit int, Tid_ int64, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("tid = ?", Tid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaVid Get TaxonomyTermFieldDatas via Vid
func GetTaxonomyTermFieldDatasViaVid(offset int, limit int, Vid_ string, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaLangcode Get TaxonomyTermFieldDatas via Langcode
func GetTaxonomyTermFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaName Get TaxonomyTermFieldDatas via Name
func GetTaxonomyTermFieldDatasViaName(offset int, limit int, Name_ string, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaDescription_value Get TaxonomyTermFieldDatas via Description_value
func GetTaxonomyTermFieldDatasViaDescription_value(offset int, limit int, Description_value_ string, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description__value = ?", Description_value_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaDescription_format Get TaxonomyTermFieldDatas via Description_format
func GetTaxonomyTermFieldDatasViaDescription_format(offset int, limit int, Description_format_ string, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("description__format = ?", Description_format_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaWeight Get TaxonomyTermFieldDatas via Weight
func GetTaxonomyTermFieldDatasViaWeight(offset int, limit int, Weight_ int, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("weight = ?", Weight_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaChanged Get TaxonomyTermFieldDatas via Changed
func GetTaxonomyTermFieldDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// GetTaxonomyTermFieldDatasViaDefaultLangcode Get TaxonomyTermFieldDatas via DefaultLangcode
func GetTaxonomyTermFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = new([]*TaxonomyTermFieldData)
	err := Engine.Table("taxonomy_term_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermFieldData)
	return _TaxonomyTermFieldData, err
}

// HasTaxonomyTermFieldDataViaTid Has TaxonomyTermFieldData via Tid
func HasTaxonomyTermFieldDataViaTid(iTid int64) bool {
	if has, err := Engine.Where("tid = ?", iTid).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaVid Has TaxonomyTermFieldData via Vid
func HasTaxonomyTermFieldDataViaVid(iVid string) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaLangcode Has TaxonomyTermFieldData via Langcode
func HasTaxonomyTermFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaName Has TaxonomyTermFieldData via Name
func HasTaxonomyTermFieldDataViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaDescription_value Has TaxonomyTermFieldData via Description_value
func HasTaxonomyTermFieldDataViaDescription_value(iDescription_value string) bool {
	if has, err := Engine.Where("description__value = ?", iDescription_value).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaDescription_format Has TaxonomyTermFieldData via Description_format
func HasTaxonomyTermFieldDataViaDescription_format(iDescription_format string) bool {
	if has, err := Engine.Where("description__format = ?", iDescription_format).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaWeight Has TaxonomyTermFieldData via Weight
func HasTaxonomyTermFieldDataViaWeight(iWeight int) bool {
	if has, err := Engine.Where("weight = ?", iWeight).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaChanged Has TaxonomyTermFieldData via Changed
func HasTaxonomyTermFieldDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermFieldDataViaDefaultLangcode Has TaxonomyTermFieldData via DefaultLangcode
func HasTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(TaxonomyTermFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTaxonomyTermFieldDataViaTid Get TaxonomyTermFieldData via Tid
func GetTaxonomyTermFieldDataViaTid(iTid int64) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Tid: iTid}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaVid Get TaxonomyTermFieldData via Vid
func GetTaxonomyTermFieldDataViaVid(iVid string) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Vid: iVid}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaLangcode Get TaxonomyTermFieldData via Langcode
func GetTaxonomyTermFieldDataViaLangcode(iLangcode string) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaName Get TaxonomyTermFieldData via Name
func GetTaxonomyTermFieldDataViaName(iName string) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Name: iName}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaDescription_value Get TaxonomyTermFieldData via Description_value
func GetTaxonomyTermFieldDataViaDescription_value(iDescription_value string) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Description_value: iDescription_value}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaDescription_format Get TaxonomyTermFieldData via Description_format
func GetTaxonomyTermFieldDataViaDescription_format(iDescription_format string) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Description_format: iDescription_format}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaWeight Get TaxonomyTermFieldData via Weight
func GetTaxonomyTermFieldDataViaWeight(iWeight int) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Weight: iWeight}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaChanged Get TaxonomyTermFieldData via Changed
func GetTaxonomyTermFieldDataViaChanged(iChanged int) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Changed: iChanged}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermFieldDataViaDefaultLangcode Get TaxonomyTermFieldData via DefaultLangcode
func GetTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode int) (*TaxonomyTermFieldData, error) {
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_TaxonomyTermFieldData)
	if has {
		return _TaxonomyTermFieldData, err
	} else {
		return nil, err
	}
}

// SetTaxonomyTermFieldDataViaTid Set TaxonomyTermFieldData via Tid
func SetTaxonomyTermFieldDataViaTid(iTid int64, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Tid = iTid
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaVid Set TaxonomyTermFieldData via Vid
func SetTaxonomyTermFieldDataViaVid(iVid string, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Vid = iVid
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaLangcode Set TaxonomyTermFieldData via Langcode
func SetTaxonomyTermFieldDataViaLangcode(iLangcode string, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Langcode = iLangcode
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaName Set TaxonomyTermFieldData via Name
func SetTaxonomyTermFieldDataViaName(iName string, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Name = iName
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaDescription_value Set TaxonomyTermFieldData via Description_value
func SetTaxonomyTermFieldDataViaDescription_value(iDescription_value string, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Description_value = iDescription_value
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaDescription_format Set TaxonomyTermFieldData via Description_format
func SetTaxonomyTermFieldDataViaDescription_format(iDescription_format string, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Description_format = iDescription_format
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaWeight Set TaxonomyTermFieldData via Weight
func SetTaxonomyTermFieldDataViaWeight(iWeight int, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Weight = iWeight
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaChanged Set TaxonomyTermFieldData via Changed
func SetTaxonomyTermFieldDataViaChanged(iChanged int, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.Changed = iChanged
	return Engine.Insert(taxonomy_term_field_data)
}

// SetTaxonomyTermFieldDataViaDefaultLangcode Set TaxonomyTermFieldData via DefaultLangcode
func SetTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode int, taxonomy_term_field_data *TaxonomyTermFieldData) (int64, error) {
	taxonomy_term_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(taxonomy_term_field_data)
}

// AddTaxonomyTermFieldData Add TaxonomyTermFieldData via all columns
func AddTaxonomyTermFieldData(iTid int64, iVid string, iLangcode string, iName string, iDescription_value string, iDescription_format string, iWeight int, iChanged int, iDefaultLangcode int) error {
	_TaxonomyTermFieldData := &TaxonomyTermFieldData{Tid: iTid, Vid: iVid, Langcode: iLangcode, Name: iName, Description_value: iDescription_value, Description_format: iDescription_format, Weight: iWeight, Changed: iChanged, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_TaxonomyTermFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTaxonomyTermFieldData Post TaxonomyTermFieldData via iTaxonomyTermFieldData
func PostTaxonomyTermFieldData(iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	_, err := Engine.Insert(iTaxonomyTermFieldData)
	return iTaxonomyTermFieldData.Tid, err
}

// PutTaxonomyTermFieldData Put TaxonomyTermFieldData
func PutTaxonomyTermFieldData(iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	_, err := Engine.Id(iTaxonomyTermFieldData.Tid).Update(iTaxonomyTermFieldData)
	return iTaxonomyTermFieldData.Tid, err
}

// PutTaxonomyTermFieldDataViaTid Put TaxonomyTermFieldData via Tid
func PutTaxonomyTermFieldDataViaTid(Tid_ int64, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Tid: Tid_})
	return row, err
}

// PutTaxonomyTermFieldDataViaVid Put TaxonomyTermFieldData via Vid
func PutTaxonomyTermFieldDataViaVid(Vid_ string, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Vid: Vid_})
	return row, err
}

// PutTaxonomyTermFieldDataViaLangcode Put TaxonomyTermFieldData via Langcode
func PutTaxonomyTermFieldDataViaLangcode(Langcode_ string, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Langcode: Langcode_})
	return row, err
}

// PutTaxonomyTermFieldDataViaName Put TaxonomyTermFieldData via Name
func PutTaxonomyTermFieldDataViaName(Name_ string, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Name: Name_})
	return row, err
}

// PutTaxonomyTermFieldDataViaDescription_value Put TaxonomyTermFieldData via Description_value
func PutTaxonomyTermFieldDataViaDescription_value(Description_value_ string, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Description_value: Description_value_})
	return row, err
}

// PutTaxonomyTermFieldDataViaDescription_format Put TaxonomyTermFieldData via Description_format
func PutTaxonomyTermFieldDataViaDescription_format(Description_format_ string, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Description_format: Description_format_})
	return row, err
}

// PutTaxonomyTermFieldDataViaWeight Put TaxonomyTermFieldData via Weight
func PutTaxonomyTermFieldDataViaWeight(Weight_ int, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Weight: Weight_})
	return row, err
}

// PutTaxonomyTermFieldDataViaChanged Put TaxonomyTermFieldData via Changed
func PutTaxonomyTermFieldDataViaChanged(Changed_ int, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{Changed: Changed_})
	return row, err
}

// PutTaxonomyTermFieldDataViaDefaultLangcode Put TaxonomyTermFieldData via DefaultLangcode
func PutTaxonomyTermFieldDataViaDefaultLangcode(DefaultLangcode_ int, iTaxonomyTermFieldData *TaxonomyTermFieldData) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermFieldData, &TaxonomyTermFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateTaxonomyTermFieldDataViaTid via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaTid(iTid int64, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("tid = ?", iTid).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaVid via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaVid(iVid string, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("vid = ?", iVid).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaLangcode via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaLangcode(iLangcode string, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("langcode = ?", iLangcode).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaName via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaName(iName string, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("name = ?", iName).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaDescription_value via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaDescription_value(iDescription_value string, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("description__value = ?", iDescription_value).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaDescription_format via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaDescription_format(iDescription_format string, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("description__format = ?", iDescription_format).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaWeight via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaWeight(iWeight int, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("weight = ?", iWeight).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaChanged via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaChanged(iChanged int, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("changed = ?", iChanged).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode int, iTaxonomyTermFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iTaxonomyTermFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTaxonomyTermFieldData Delete TaxonomyTermFieldData
func DeleteTaxonomyTermFieldData(iTid int64) (int64, error) {
	row, err := Engine.Id(iTid).Delete(new(TaxonomyTermFieldData))
	return row, err
}

// DeleteTaxonomyTermFieldDataViaTid Delete TaxonomyTermFieldData via Tid
func DeleteTaxonomyTermFieldDataViaTid(iTid int64) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Tid: iTid}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("tid = ?", iTid).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaVid Delete TaxonomyTermFieldData via Vid
func DeleteTaxonomyTermFieldDataViaVid(iVid string) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Vid: iVid}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaLangcode Delete TaxonomyTermFieldData via Langcode
func DeleteTaxonomyTermFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaName Delete TaxonomyTermFieldData via Name
func DeleteTaxonomyTermFieldDataViaName(iName string) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Name: iName}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaDescription_value Delete TaxonomyTermFieldData via Description_value
func DeleteTaxonomyTermFieldDataViaDescription_value(iDescription_value string) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Description_value: iDescription_value}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("description__value = ?", iDescription_value).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaDescription_format Delete TaxonomyTermFieldData via Description_format
func DeleteTaxonomyTermFieldDataViaDescription_format(iDescription_format string) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Description_format: iDescription_format}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("description__format = ?", iDescription_format).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaWeight Delete TaxonomyTermFieldData via Weight
func DeleteTaxonomyTermFieldDataViaWeight(iWeight int) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Weight: iWeight}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("weight = ?", iWeight).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaChanged Delete TaxonomyTermFieldData via Changed
func DeleteTaxonomyTermFieldDataViaChanged(iChanged int) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{Changed: iChanged}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermFieldDataViaDefaultLangcode Delete TaxonomyTermFieldData via DefaultLangcode
func DeleteTaxonomyTermFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _TaxonomyTermFieldData = &TaxonomyTermFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_TaxonomyTermFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(TaxonomyTermFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
