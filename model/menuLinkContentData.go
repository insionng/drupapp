package model

type MenuLinkContentData struct {
	Id              int64  `xorm:"not null pk index(menu_link_content__id__default_langcode__langcode) INT(10)"`
	Bundle          string `xorm:"not null VARCHAR(32)"`
	Langcode        string `xorm:"not null pk index(menu_link_content__id__default_langcode__langcode) VARCHAR(12)"`
	Title           string `xorm:"VARCHAR(255)"`
	Description     string `xorm:"VARCHAR(255)"`
	MenuName        string `xorm:"VARCHAR(255)"`
	Link_uri        string `xorm:"index VARCHAR(2048)"`
	Link_title      string `xorm:"VARCHAR(255)"`
	Link_options    []byte `xorm:"LONGBLOB"`
	External        int    `xorm:"TINYINT(4)"`
	Rediscover      int    `xorm:"TINYINT(4)"`
	Weight          int    `xorm:"INT(11)"`
	Expanded        int    `xorm:"TINYINT(4)"`
	Enabled         int    `xorm:"TINYINT(4)"`
	Parent          string `xorm:"VARCHAR(255)"`
	Changed         int    `xorm:"INT(11)"`
	DefaultLangcode int    `xorm:"not null index(menu_link_content__id__default_langcode__langcode) TINYINT(4)"`
}

// GetMenuLinkContentDatasCount MenuLinkContentDatas' Count
func GetMenuLinkContentDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&MenuLinkContentData{})
	return total, err
}

// GetMenuLinkContentDataCountViaId Get MenuLinkContentData via Id
func GetMenuLinkContentDataCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&MenuLinkContentData{Id: iId})
	return n
}

// GetMenuLinkContentDataCountViaBundle Get MenuLinkContentData via Bundle
func GetMenuLinkContentDataCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&MenuLinkContentData{Bundle: iBundle})
	return n
}

// GetMenuLinkContentDataCountViaLangcode Get MenuLinkContentData via Langcode
func GetMenuLinkContentDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&MenuLinkContentData{Langcode: iLangcode})
	return n
}

// GetMenuLinkContentDataCountViaTitle Get MenuLinkContentData via Title
func GetMenuLinkContentDataCountViaTitle(iTitle string) int64 {
	n, _ := Engine.Where("title = ?", iTitle).Count(&MenuLinkContentData{Title: iTitle})
	return n
}

// GetMenuLinkContentDataCountViaDescription Get MenuLinkContentData via Description
func GetMenuLinkContentDataCountViaDescription(iDescription string) int64 {
	n, _ := Engine.Where("description = ?", iDescription).Count(&MenuLinkContentData{Description: iDescription})
	return n
}

// GetMenuLinkContentDataCountViaMenuName Get MenuLinkContentData via MenuName
func GetMenuLinkContentDataCountViaMenuName(iMenuName string) int64 {
	n, _ := Engine.Where("menu_name = ?", iMenuName).Count(&MenuLinkContentData{MenuName: iMenuName})
	return n
}

// GetMenuLinkContentDataCountViaLink_uri Get MenuLinkContentData via Link_uri
func GetMenuLinkContentDataCountViaLink_uri(iLink_uri string) int64 {
	n, _ := Engine.Where("link__uri = ?", iLink_uri).Count(&MenuLinkContentData{Link_uri: iLink_uri})
	return n
}

// GetMenuLinkContentDataCountViaLink_title Get MenuLinkContentData via Link_title
func GetMenuLinkContentDataCountViaLink_title(iLink_title string) int64 {
	n, _ := Engine.Where("link__title = ?", iLink_title).Count(&MenuLinkContentData{Link_title: iLink_title})
	return n
}

// GetMenuLinkContentDataCountViaLink_options Get MenuLinkContentData via Link_options
func GetMenuLinkContentDataCountViaLink_options(iLink_options []byte) int64 {
	n, _ := Engine.Where("link__options = ?", iLink_options).Count(&MenuLinkContentData{Link_options: iLink_options})
	return n
}

// GetMenuLinkContentDataCountViaExternal Get MenuLinkContentData via External
func GetMenuLinkContentDataCountViaExternal(iExternal int) int64 {
	n, _ := Engine.Where("external = ?", iExternal).Count(&MenuLinkContentData{External: iExternal})
	return n
}

// GetMenuLinkContentDataCountViaRediscover Get MenuLinkContentData via Rediscover
func GetMenuLinkContentDataCountViaRediscover(iRediscover int) int64 {
	n, _ := Engine.Where("rediscover = ?", iRediscover).Count(&MenuLinkContentData{Rediscover: iRediscover})
	return n
}

// GetMenuLinkContentDataCountViaWeight Get MenuLinkContentData via Weight
func GetMenuLinkContentDataCountViaWeight(iWeight int) int64 {
	n, _ := Engine.Where("weight = ?", iWeight).Count(&MenuLinkContentData{Weight: iWeight})
	return n
}

// GetMenuLinkContentDataCountViaExpanded Get MenuLinkContentData via Expanded
func GetMenuLinkContentDataCountViaExpanded(iExpanded int) int64 {
	n, _ := Engine.Where("expanded = ?", iExpanded).Count(&MenuLinkContentData{Expanded: iExpanded})
	return n
}

// GetMenuLinkContentDataCountViaEnabled Get MenuLinkContentData via Enabled
func GetMenuLinkContentDataCountViaEnabled(iEnabled int) int64 {
	n, _ := Engine.Where("enabled = ?", iEnabled).Count(&MenuLinkContentData{Enabled: iEnabled})
	return n
}

// GetMenuLinkContentDataCountViaParent Get MenuLinkContentData via Parent
func GetMenuLinkContentDataCountViaParent(iParent string) int64 {
	n, _ := Engine.Where("parent = ?", iParent).Count(&MenuLinkContentData{Parent: iParent})
	return n
}

// GetMenuLinkContentDataCountViaChanged Get MenuLinkContentData via Changed
func GetMenuLinkContentDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&MenuLinkContentData{Changed: iChanged})
	return n
}

// GetMenuLinkContentDataCountViaDefaultLangcode Get MenuLinkContentData via DefaultLangcode
func GetMenuLinkContentDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&MenuLinkContentData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetMenuLinkContentDatasByIdAndBundleAndLangcode Get MenuLinkContentDatas via IdAndBundleAndLangcode
func GetMenuLinkContentDatasByIdAndBundleAndLangcode(offset int, limit int, Id_ int64, Bundle_ string, Langcode_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and langcode = ?", Id_, Bundle_, Langcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndTitle Get MenuLinkContentDatas via IdAndBundleAndTitle
func GetMenuLinkContentDatasByIdAndBundleAndTitle(offset int, limit int, Id_ int64, Bundle_ string, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and title = ?", Id_, Bundle_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndDescription Get MenuLinkContentDatas via IdAndBundleAndDescription
func GetMenuLinkContentDatasByIdAndBundleAndDescription(offset int, limit int, Id_ int64, Bundle_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and description = ?", Id_, Bundle_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndMenuName Get MenuLinkContentDatas via IdAndBundleAndMenuName
func GetMenuLinkContentDatasByIdAndBundleAndMenuName(offset int, limit int, Id_ int64, Bundle_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and menu_name = ?", Id_, Bundle_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndLink_uri Get MenuLinkContentDatas via IdAndBundleAndLink_uri
func GetMenuLinkContentDatasByIdAndBundleAndLink_uri(offset int, limit int, Id_ int64, Bundle_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and link_uri = ?", Id_, Bundle_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndLink_title Get MenuLinkContentDatas via IdAndBundleAndLink_title
func GetMenuLinkContentDatasByIdAndBundleAndLink_title(offset int, limit int, Id_ int64, Bundle_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and link_title = ?", Id_, Bundle_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndLink_options Get MenuLinkContentDatas via IdAndBundleAndLink_options
func GetMenuLinkContentDatasByIdAndBundleAndLink_options(offset int, limit int, Id_ int64, Bundle_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and link_options = ?", Id_, Bundle_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndExternal Get MenuLinkContentDatas via IdAndBundleAndExternal
func GetMenuLinkContentDatasByIdAndBundleAndExternal(offset int, limit int, Id_ int64, Bundle_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and external = ?", Id_, Bundle_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndRediscover Get MenuLinkContentDatas via IdAndBundleAndRediscover
func GetMenuLinkContentDatasByIdAndBundleAndRediscover(offset int, limit int, Id_ int64, Bundle_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and rediscover = ?", Id_, Bundle_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndWeight Get MenuLinkContentDatas via IdAndBundleAndWeight
func GetMenuLinkContentDatasByIdAndBundleAndWeight(offset int, limit int, Id_ int64, Bundle_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and weight = ?", Id_, Bundle_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndExpanded Get MenuLinkContentDatas via IdAndBundleAndExpanded
func GetMenuLinkContentDatasByIdAndBundleAndExpanded(offset int, limit int, Id_ int64, Bundle_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and expanded = ?", Id_, Bundle_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndEnabled Get MenuLinkContentDatas via IdAndBundleAndEnabled
func GetMenuLinkContentDatasByIdAndBundleAndEnabled(offset int, limit int, Id_ int64, Bundle_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and enabled = ?", Id_, Bundle_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndParent Get MenuLinkContentDatas via IdAndBundleAndParent
func GetMenuLinkContentDatasByIdAndBundleAndParent(offset int, limit int, Id_ int64, Bundle_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and parent = ?", Id_, Bundle_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndChanged Get MenuLinkContentDatas via IdAndBundleAndChanged
func GetMenuLinkContentDatasByIdAndBundleAndChanged(offset int, limit int, Id_ int64, Bundle_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and changed = ?", Id_, Bundle_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundleAndDefaultLangcode Get MenuLinkContentDatas via IdAndBundleAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndBundleAndDefaultLangcode(offset int, limit int, Id_ int64, Bundle_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ? and default_langcode = ?", Id_, Bundle_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndTitle Get MenuLinkContentDatas via IdAndLangcodeAndTitle
func GetMenuLinkContentDatasByIdAndLangcodeAndTitle(offset int, limit int, Id_ int64, Langcode_ string, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and title = ?", Id_, Langcode_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndDescription Get MenuLinkContentDatas via IdAndLangcodeAndDescription
func GetMenuLinkContentDatasByIdAndLangcodeAndDescription(offset int, limit int, Id_ int64, Langcode_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and description = ?", Id_, Langcode_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndMenuName Get MenuLinkContentDatas via IdAndLangcodeAndMenuName
func GetMenuLinkContentDatasByIdAndLangcodeAndMenuName(offset int, limit int, Id_ int64, Langcode_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and menu_name = ?", Id_, Langcode_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndLink_uri Get MenuLinkContentDatas via IdAndLangcodeAndLink_uri
func GetMenuLinkContentDatasByIdAndLangcodeAndLink_uri(offset int, limit int, Id_ int64, Langcode_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and link_uri = ?", Id_, Langcode_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndLink_title Get MenuLinkContentDatas via IdAndLangcodeAndLink_title
func GetMenuLinkContentDatasByIdAndLangcodeAndLink_title(offset int, limit int, Id_ int64, Langcode_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and link_title = ?", Id_, Langcode_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndLink_options Get MenuLinkContentDatas via IdAndLangcodeAndLink_options
func GetMenuLinkContentDatasByIdAndLangcodeAndLink_options(offset int, limit int, Id_ int64, Langcode_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and link_options = ?", Id_, Langcode_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndExternal Get MenuLinkContentDatas via IdAndLangcodeAndExternal
func GetMenuLinkContentDatasByIdAndLangcodeAndExternal(offset int, limit int, Id_ int64, Langcode_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and external = ?", Id_, Langcode_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndRediscover Get MenuLinkContentDatas via IdAndLangcodeAndRediscover
func GetMenuLinkContentDatasByIdAndLangcodeAndRediscover(offset int, limit int, Id_ int64, Langcode_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and rediscover = ?", Id_, Langcode_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndWeight Get MenuLinkContentDatas via IdAndLangcodeAndWeight
func GetMenuLinkContentDatasByIdAndLangcodeAndWeight(offset int, limit int, Id_ int64, Langcode_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and weight = ?", Id_, Langcode_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndExpanded Get MenuLinkContentDatas via IdAndLangcodeAndExpanded
func GetMenuLinkContentDatasByIdAndLangcodeAndExpanded(offset int, limit int, Id_ int64, Langcode_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and expanded = ?", Id_, Langcode_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndEnabled Get MenuLinkContentDatas via IdAndLangcodeAndEnabled
func GetMenuLinkContentDatasByIdAndLangcodeAndEnabled(offset int, limit int, Id_ int64, Langcode_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and enabled = ?", Id_, Langcode_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndParent Get MenuLinkContentDatas via IdAndLangcodeAndParent
func GetMenuLinkContentDatasByIdAndLangcodeAndParent(offset int, limit int, Id_ int64, Langcode_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and parent = ?", Id_, Langcode_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndChanged Get MenuLinkContentDatas via IdAndLangcodeAndChanged
func GetMenuLinkContentDatasByIdAndLangcodeAndChanged(offset int, limit int, Id_ int64, Langcode_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and changed = ?", Id_, Langcode_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcodeAndDefaultLangcode Get MenuLinkContentDatas via IdAndLangcodeAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndLangcodeAndDefaultLangcode(offset int, limit int, Id_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ? and default_langcode = ?", Id_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndDescription Get MenuLinkContentDatas via IdAndTitleAndDescription
func GetMenuLinkContentDatasByIdAndTitleAndDescription(offset int, limit int, Id_ int64, Title_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and description = ?", Id_, Title_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndMenuName Get MenuLinkContentDatas via IdAndTitleAndMenuName
func GetMenuLinkContentDatasByIdAndTitleAndMenuName(offset int, limit int, Id_ int64, Title_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and menu_name = ?", Id_, Title_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndLink_uri Get MenuLinkContentDatas via IdAndTitleAndLink_uri
func GetMenuLinkContentDatasByIdAndTitleAndLink_uri(offset int, limit int, Id_ int64, Title_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and link_uri = ?", Id_, Title_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndLink_title Get MenuLinkContentDatas via IdAndTitleAndLink_title
func GetMenuLinkContentDatasByIdAndTitleAndLink_title(offset int, limit int, Id_ int64, Title_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and link_title = ?", Id_, Title_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndLink_options Get MenuLinkContentDatas via IdAndTitleAndLink_options
func GetMenuLinkContentDatasByIdAndTitleAndLink_options(offset int, limit int, Id_ int64, Title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and link_options = ?", Id_, Title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndExternal Get MenuLinkContentDatas via IdAndTitleAndExternal
func GetMenuLinkContentDatasByIdAndTitleAndExternal(offset int, limit int, Id_ int64, Title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and external = ?", Id_, Title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndRediscover Get MenuLinkContentDatas via IdAndTitleAndRediscover
func GetMenuLinkContentDatasByIdAndTitleAndRediscover(offset int, limit int, Id_ int64, Title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and rediscover = ?", Id_, Title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndWeight Get MenuLinkContentDatas via IdAndTitleAndWeight
func GetMenuLinkContentDatasByIdAndTitleAndWeight(offset int, limit int, Id_ int64, Title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and weight = ?", Id_, Title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndExpanded Get MenuLinkContentDatas via IdAndTitleAndExpanded
func GetMenuLinkContentDatasByIdAndTitleAndExpanded(offset int, limit int, Id_ int64, Title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and expanded = ?", Id_, Title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndEnabled Get MenuLinkContentDatas via IdAndTitleAndEnabled
func GetMenuLinkContentDatasByIdAndTitleAndEnabled(offset int, limit int, Id_ int64, Title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and enabled = ?", Id_, Title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndParent Get MenuLinkContentDatas via IdAndTitleAndParent
func GetMenuLinkContentDatasByIdAndTitleAndParent(offset int, limit int, Id_ int64, Title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and parent = ?", Id_, Title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndChanged Get MenuLinkContentDatas via IdAndTitleAndChanged
func GetMenuLinkContentDatasByIdAndTitleAndChanged(offset int, limit int, Id_ int64, Title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and changed = ?", Id_, Title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitleAndDefaultLangcode Get MenuLinkContentDatas via IdAndTitleAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndTitleAndDefaultLangcode(offset int, limit int, Id_ int64, Title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ? and default_langcode = ?", Id_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndMenuName Get MenuLinkContentDatas via IdAndDescriptionAndMenuName
func GetMenuLinkContentDatasByIdAndDescriptionAndMenuName(offset int, limit int, Id_ int64, Description_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and menu_name = ?", Id_, Description_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndLink_uri Get MenuLinkContentDatas via IdAndDescriptionAndLink_uri
func GetMenuLinkContentDatasByIdAndDescriptionAndLink_uri(offset int, limit int, Id_ int64, Description_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and link_uri = ?", Id_, Description_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndLink_title Get MenuLinkContentDatas via IdAndDescriptionAndLink_title
func GetMenuLinkContentDatasByIdAndDescriptionAndLink_title(offset int, limit int, Id_ int64, Description_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and link_title = ?", Id_, Description_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndLink_options Get MenuLinkContentDatas via IdAndDescriptionAndLink_options
func GetMenuLinkContentDatasByIdAndDescriptionAndLink_options(offset int, limit int, Id_ int64, Description_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and link_options = ?", Id_, Description_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndExternal Get MenuLinkContentDatas via IdAndDescriptionAndExternal
func GetMenuLinkContentDatasByIdAndDescriptionAndExternal(offset int, limit int, Id_ int64, Description_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and external = ?", Id_, Description_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndRediscover Get MenuLinkContentDatas via IdAndDescriptionAndRediscover
func GetMenuLinkContentDatasByIdAndDescriptionAndRediscover(offset int, limit int, Id_ int64, Description_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and rediscover = ?", Id_, Description_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndWeight Get MenuLinkContentDatas via IdAndDescriptionAndWeight
func GetMenuLinkContentDatasByIdAndDescriptionAndWeight(offset int, limit int, Id_ int64, Description_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and weight = ?", Id_, Description_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndExpanded Get MenuLinkContentDatas via IdAndDescriptionAndExpanded
func GetMenuLinkContentDatasByIdAndDescriptionAndExpanded(offset int, limit int, Id_ int64, Description_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and expanded = ?", Id_, Description_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndEnabled Get MenuLinkContentDatas via IdAndDescriptionAndEnabled
func GetMenuLinkContentDatasByIdAndDescriptionAndEnabled(offset int, limit int, Id_ int64, Description_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and enabled = ?", Id_, Description_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndParent Get MenuLinkContentDatas via IdAndDescriptionAndParent
func GetMenuLinkContentDatasByIdAndDescriptionAndParent(offset int, limit int, Id_ int64, Description_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and parent = ?", Id_, Description_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndChanged Get MenuLinkContentDatas via IdAndDescriptionAndChanged
func GetMenuLinkContentDatasByIdAndDescriptionAndChanged(offset int, limit int, Id_ int64, Description_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and changed = ?", Id_, Description_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescriptionAndDefaultLangcode Get MenuLinkContentDatas via IdAndDescriptionAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndDescriptionAndDefaultLangcode(offset int, limit int, Id_ int64, Description_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ? and default_langcode = ?", Id_, Description_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndLink_uri Get MenuLinkContentDatas via IdAndMenuNameAndLink_uri
func GetMenuLinkContentDatasByIdAndMenuNameAndLink_uri(offset int, limit int, Id_ int64, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and link_uri = ?", Id_, MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndLink_title Get MenuLinkContentDatas via IdAndMenuNameAndLink_title
func GetMenuLinkContentDatasByIdAndMenuNameAndLink_title(offset int, limit int, Id_ int64, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and link_title = ?", Id_, MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndLink_options Get MenuLinkContentDatas via IdAndMenuNameAndLink_options
func GetMenuLinkContentDatasByIdAndMenuNameAndLink_options(offset int, limit int, Id_ int64, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and link_options = ?", Id_, MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndExternal Get MenuLinkContentDatas via IdAndMenuNameAndExternal
func GetMenuLinkContentDatasByIdAndMenuNameAndExternal(offset int, limit int, Id_ int64, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and external = ?", Id_, MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndRediscover Get MenuLinkContentDatas via IdAndMenuNameAndRediscover
func GetMenuLinkContentDatasByIdAndMenuNameAndRediscover(offset int, limit int, Id_ int64, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and rediscover = ?", Id_, MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndWeight Get MenuLinkContentDatas via IdAndMenuNameAndWeight
func GetMenuLinkContentDatasByIdAndMenuNameAndWeight(offset int, limit int, Id_ int64, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and weight = ?", Id_, MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndExpanded Get MenuLinkContentDatas via IdAndMenuNameAndExpanded
func GetMenuLinkContentDatasByIdAndMenuNameAndExpanded(offset int, limit int, Id_ int64, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and expanded = ?", Id_, MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndEnabled Get MenuLinkContentDatas via IdAndMenuNameAndEnabled
func GetMenuLinkContentDatasByIdAndMenuNameAndEnabled(offset int, limit int, Id_ int64, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and enabled = ?", Id_, MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndParent Get MenuLinkContentDatas via IdAndMenuNameAndParent
func GetMenuLinkContentDatasByIdAndMenuNameAndParent(offset int, limit int, Id_ int64, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and parent = ?", Id_, MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndChanged Get MenuLinkContentDatas via IdAndMenuNameAndChanged
func GetMenuLinkContentDatasByIdAndMenuNameAndChanged(offset int, limit int, Id_ int64, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and changed = ?", Id_, MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuNameAndDefaultLangcode Get MenuLinkContentDatas via IdAndMenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndMenuNameAndDefaultLangcode(offset int, limit int, Id_ int64, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ? and default_langcode = ?", Id_, MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndLink_title Get MenuLinkContentDatas via IdAndLink_uriAndLink_title
func GetMenuLinkContentDatasByIdAndLink_uriAndLink_title(offset int, limit int, Id_ int64, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and link_title = ?", Id_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndLink_options Get MenuLinkContentDatas via IdAndLink_uriAndLink_options
func GetMenuLinkContentDatasByIdAndLink_uriAndLink_options(offset int, limit int, Id_ int64, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and link_options = ?", Id_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndExternal Get MenuLinkContentDatas via IdAndLink_uriAndExternal
func GetMenuLinkContentDatasByIdAndLink_uriAndExternal(offset int, limit int, Id_ int64, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and external = ?", Id_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndRediscover Get MenuLinkContentDatas via IdAndLink_uriAndRediscover
func GetMenuLinkContentDatasByIdAndLink_uriAndRediscover(offset int, limit int, Id_ int64, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and rediscover = ?", Id_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndWeight Get MenuLinkContentDatas via IdAndLink_uriAndWeight
func GetMenuLinkContentDatasByIdAndLink_uriAndWeight(offset int, limit int, Id_ int64, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and weight = ?", Id_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndExpanded Get MenuLinkContentDatas via IdAndLink_uriAndExpanded
func GetMenuLinkContentDatasByIdAndLink_uriAndExpanded(offset int, limit int, Id_ int64, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and expanded = ?", Id_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndEnabled Get MenuLinkContentDatas via IdAndLink_uriAndEnabled
func GetMenuLinkContentDatasByIdAndLink_uriAndEnabled(offset int, limit int, Id_ int64, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and enabled = ?", Id_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndParent Get MenuLinkContentDatas via IdAndLink_uriAndParent
func GetMenuLinkContentDatasByIdAndLink_uriAndParent(offset int, limit int, Id_ int64, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and parent = ?", Id_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndChanged Get MenuLinkContentDatas via IdAndLink_uriAndChanged
func GetMenuLinkContentDatasByIdAndLink_uriAndChanged(offset int, limit int, Id_ int64, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and changed = ?", Id_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via IdAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndLink_uriAndDefaultLangcode(offset int, limit int, Id_ int64, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ? and default_langcode = ?", Id_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndLink_options Get MenuLinkContentDatas via IdAndLink_titleAndLink_options
func GetMenuLinkContentDatasByIdAndLink_titleAndLink_options(offset int, limit int, Id_ int64, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and link_options = ?", Id_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndExternal Get MenuLinkContentDatas via IdAndLink_titleAndExternal
func GetMenuLinkContentDatasByIdAndLink_titleAndExternal(offset int, limit int, Id_ int64, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and external = ?", Id_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndRediscover Get MenuLinkContentDatas via IdAndLink_titleAndRediscover
func GetMenuLinkContentDatasByIdAndLink_titleAndRediscover(offset int, limit int, Id_ int64, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and rediscover = ?", Id_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndWeight Get MenuLinkContentDatas via IdAndLink_titleAndWeight
func GetMenuLinkContentDatasByIdAndLink_titleAndWeight(offset int, limit int, Id_ int64, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and weight = ?", Id_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndExpanded Get MenuLinkContentDatas via IdAndLink_titleAndExpanded
func GetMenuLinkContentDatasByIdAndLink_titleAndExpanded(offset int, limit int, Id_ int64, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and expanded = ?", Id_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndEnabled Get MenuLinkContentDatas via IdAndLink_titleAndEnabled
func GetMenuLinkContentDatasByIdAndLink_titleAndEnabled(offset int, limit int, Id_ int64, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and enabled = ?", Id_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndParent Get MenuLinkContentDatas via IdAndLink_titleAndParent
func GetMenuLinkContentDatasByIdAndLink_titleAndParent(offset int, limit int, Id_ int64, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and parent = ?", Id_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndChanged Get MenuLinkContentDatas via IdAndLink_titleAndChanged
func GetMenuLinkContentDatasByIdAndLink_titleAndChanged(offset int, limit int, Id_ int64, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and changed = ?", Id_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via IdAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndLink_titleAndDefaultLangcode(offset int, limit int, Id_ int64, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ? and default_langcode = ?", Id_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndExternal Get MenuLinkContentDatas via IdAndLink_optionsAndExternal
func GetMenuLinkContentDatasByIdAndLink_optionsAndExternal(offset int, limit int, Id_ int64, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and external = ?", Id_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndRediscover Get MenuLinkContentDatas via IdAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByIdAndLink_optionsAndRediscover(offset int, limit int, Id_ int64, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and rediscover = ?", Id_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndWeight Get MenuLinkContentDatas via IdAndLink_optionsAndWeight
func GetMenuLinkContentDatasByIdAndLink_optionsAndWeight(offset int, limit int, Id_ int64, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and weight = ?", Id_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndExpanded Get MenuLinkContentDatas via IdAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByIdAndLink_optionsAndExpanded(offset int, limit int, Id_ int64, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and expanded = ?", Id_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndEnabled Get MenuLinkContentDatas via IdAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByIdAndLink_optionsAndEnabled(offset int, limit int, Id_ int64, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and enabled = ?", Id_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndParent Get MenuLinkContentDatas via IdAndLink_optionsAndParent
func GetMenuLinkContentDatasByIdAndLink_optionsAndParent(offset int, limit int, Id_ int64, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and parent = ?", Id_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndChanged Get MenuLinkContentDatas via IdAndLink_optionsAndChanged
func GetMenuLinkContentDatasByIdAndLink_optionsAndChanged(offset int, limit int, Id_ int64, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and changed = ?", Id_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via IdAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndLink_optionsAndDefaultLangcode(offset int, limit int, Id_ int64, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ? and default_langcode = ?", Id_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndRediscover Get MenuLinkContentDatas via IdAndExternalAndRediscover
func GetMenuLinkContentDatasByIdAndExternalAndRediscover(offset int, limit int, Id_ int64, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and rediscover = ?", Id_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndWeight Get MenuLinkContentDatas via IdAndExternalAndWeight
func GetMenuLinkContentDatasByIdAndExternalAndWeight(offset int, limit int, Id_ int64, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and weight = ?", Id_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndExpanded Get MenuLinkContentDatas via IdAndExternalAndExpanded
func GetMenuLinkContentDatasByIdAndExternalAndExpanded(offset int, limit int, Id_ int64, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and expanded = ?", Id_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndEnabled Get MenuLinkContentDatas via IdAndExternalAndEnabled
func GetMenuLinkContentDatasByIdAndExternalAndEnabled(offset int, limit int, Id_ int64, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and enabled = ?", Id_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndParent Get MenuLinkContentDatas via IdAndExternalAndParent
func GetMenuLinkContentDatasByIdAndExternalAndParent(offset int, limit int, Id_ int64, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and parent = ?", Id_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndChanged Get MenuLinkContentDatas via IdAndExternalAndChanged
func GetMenuLinkContentDatasByIdAndExternalAndChanged(offset int, limit int, Id_ int64, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and changed = ?", Id_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternalAndDefaultLangcode Get MenuLinkContentDatas via IdAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndExternalAndDefaultLangcode(offset int, limit int, Id_ int64, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ? and default_langcode = ?", Id_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndWeight Get MenuLinkContentDatas via IdAndRediscoverAndWeight
func GetMenuLinkContentDatasByIdAndRediscoverAndWeight(offset int, limit int, Id_ int64, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and weight = ?", Id_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndExpanded Get MenuLinkContentDatas via IdAndRediscoverAndExpanded
func GetMenuLinkContentDatasByIdAndRediscoverAndExpanded(offset int, limit int, Id_ int64, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and expanded = ?", Id_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndEnabled Get MenuLinkContentDatas via IdAndRediscoverAndEnabled
func GetMenuLinkContentDatasByIdAndRediscoverAndEnabled(offset int, limit int, Id_ int64, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and enabled = ?", Id_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndParent Get MenuLinkContentDatas via IdAndRediscoverAndParent
func GetMenuLinkContentDatasByIdAndRediscoverAndParent(offset int, limit int, Id_ int64, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and parent = ?", Id_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndChanged Get MenuLinkContentDatas via IdAndRediscoverAndChanged
func GetMenuLinkContentDatasByIdAndRediscoverAndChanged(offset int, limit int, Id_ int64, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and changed = ?", Id_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via IdAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndRediscoverAndDefaultLangcode(offset int, limit int, Id_ int64, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ? and default_langcode = ?", Id_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeightAndExpanded Get MenuLinkContentDatas via IdAndWeightAndExpanded
func GetMenuLinkContentDatasByIdAndWeightAndExpanded(offset int, limit int, Id_ int64, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ? and expanded = ?", Id_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeightAndEnabled Get MenuLinkContentDatas via IdAndWeightAndEnabled
func GetMenuLinkContentDatasByIdAndWeightAndEnabled(offset int, limit int, Id_ int64, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ? and enabled = ?", Id_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeightAndParent Get MenuLinkContentDatas via IdAndWeightAndParent
func GetMenuLinkContentDatasByIdAndWeightAndParent(offset int, limit int, Id_ int64, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ? and parent = ?", Id_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeightAndChanged Get MenuLinkContentDatas via IdAndWeightAndChanged
func GetMenuLinkContentDatasByIdAndWeightAndChanged(offset int, limit int, Id_ int64, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ? and changed = ?", Id_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeightAndDefaultLangcode Get MenuLinkContentDatas via IdAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndWeightAndDefaultLangcode(offset int, limit int, Id_ int64, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ? and default_langcode = ?", Id_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExpandedAndEnabled Get MenuLinkContentDatas via IdAndExpandedAndEnabled
func GetMenuLinkContentDatasByIdAndExpandedAndEnabled(offset int, limit int, Id_ int64, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and expanded = ? and enabled = ?", Id_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExpandedAndParent Get MenuLinkContentDatas via IdAndExpandedAndParent
func GetMenuLinkContentDatasByIdAndExpandedAndParent(offset int, limit int, Id_ int64, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and expanded = ? and parent = ?", Id_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExpandedAndChanged Get MenuLinkContentDatas via IdAndExpandedAndChanged
func GetMenuLinkContentDatasByIdAndExpandedAndChanged(offset int, limit int, Id_ int64, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and expanded = ? and changed = ?", Id_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via IdAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndExpandedAndDefaultLangcode(offset int, limit int, Id_ int64, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and expanded = ? and default_langcode = ?", Id_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndEnabledAndParent Get MenuLinkContentDatas via IdAndEnabledAndParent
func GetMenuLinkContentDatasByIdAndEnabledAndParent(offset int, limit int, Id_ int64, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and enabled = ? and parent = ?", Id_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndEnabledAndChanged Get MenuLinkContentDatas via IdAndEnabledAndChanged
func GetMenuLinkContentDatasByIdAndEnabledAndChanged(offset int, limit int, Id_ int64, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and enabled = ? and changed = ?", Id_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via IdAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndEnabledAndDefaultLangcode(offset int, limit int, Id_ int64, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and enabled = ? and default_langcode = ?", Id_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndParentAndChanged Get MenuLinkContentDatas via IdAndParentAndChanged
func GetMenuLinkContentDatasByIdAndParentAndChanged(offset int, limit int, Id_ int64, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and parent = ? and changed = ?", Id_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndParentAndDefaultLangcode Get MenuLinkContentDatas via IdAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndParentAndDefaultLangcode(offset int, limit int, Id_ int64, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and parent = ? and default_langcode = ?", Id_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndChangedAndDefaultLangcode Get MenuLinkContentDatas via IdAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndChangedAndDefaultLangcode(offset int, limit int, Id_ int64, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and changed = ? and default_langcode = ?", Id_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndTitle Get MenuLinkContentDatas via BundleAndLangcodeAndTitle
func GetMenuLinkContentDatasByBundleAndLangcodeAndTitle(offset int, limit int, Bundle_ string, Langcode_ string, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and title = ?", Bundle_, Langcode_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndDescription Get MenuLinkContentDatas via BundleAndLangcodeAndDescription
func GetMenuLinkContentDatasByBundleAndLangcodeAndDescription(offset int, limit int, Bundle_ string, Langcode_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and description = ?", Bundle_, Langcode_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndMenuName Get MenuLinkContentDatas via BundleAndLangcodeAndMenuName
func GetMenuLinkContentDatasByBundleAndLangcodeAndMenuName(offset int, limit int, Bundle_ string, Langcode_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and menu_name = ?", Bundle_, Langcode_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndLink_uri Get MenuLinkContentDatas via BundleAndLangcodeAndLink_uri
func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_uri(offset int, limit int, Bundle_ string, Langcode_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and link_uri = ?", Bundle_, Langcode_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndLink_title Get MenuLinkContentDatas via BundleAndLangcodeAndLink_title
func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_title(offset int, limit int, Bundle_ string, Langcode_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and link_title = ?", Bundle_, Langcode_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndLink_options Get MenuLinkContentDatas via BundleAndLangcodeAndLink_options
func GetMenuLinkContentDatasByBundleAndLangcodeAndLink_options(offset int, limit int, Bundle_ string, Langcode_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and link_options = ?", Bundle_, Langcode_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndExternal Get MenuLinkContentDatas via BundleAndLangcodeAndExternal
func GetMenuLinkContentDatasByBundleAndLangcodeAndExternal(offset int, limit int, Bundle_ string, Langcode_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and external = ?", Bundle_, Langcode_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndRediscover Get MenuLinkContentDatas via BundleAndLangcodeAndRediscover
func GetMenuLinkContentDatasByBundleAndLangcodeAndRediscover(offset int, limit int, Bundle_ string, Langcode_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and rediscover = ?", Bundle_, Langcode_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndWeight Get MenuLinkContentDatas via BundleAndLangcodeAndWeight
func GetMenuLinkContentDatasByBundleAndLangcodeAndWeight(offset int, limit int, Bundle_ string, Langcode_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and weight = ?", Bundle_, Langcode_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndExpanded Get MenuLinkContentDatas via BundleAndLangcodeAndExpanded
func GetMenuLinkContentDatasByBundleAndLangcodeAndExpanded(offset int, limit int, Bundle_ string, Langcode_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and expanded = ?", Bundle_, Langcode_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndEnabled Get MenuLinkContentDatas via BundleAndLangcodeAndEnabled
func GetMenuLinkContentDatasByBundleAndLangcodeAndEnabled(offset int, limit int, Bundle_ string, Langcode_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and enabled = ?", Bundle_, Langcode_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndParent Get MenuLinkContentDatas via BundleAndLangcodeAndParent
func GetMenuLinkContentDatasByBundleAndLangcodeAndParent(offset int, limit int, Bundle_ string, Langcode_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and parent = ?", Bundle_, Langcode_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndChanged Get MenuLinkContentDatas via BundleAndLangcodeAndChanged
func GetMenuLinkContentDatasByBundleAndLangcodeAndChanged(offset int, limit int, Bundle_ string, Langcode_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and changed = ?", Bundle_, Langcode_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcodeAndDefaultLangcode Get MenuLinkContentDatas via BundleAndLangcodeAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndLangcodeAndDefaultLangcode(offset int, limit int, Bundle_ string, Langcode_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ? and default_langcode = ?", Bundle_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndDescription Get MenuLinkContentDatas via BundleAndTitleAndDescription
func GetMenuLinkContentDatasByBundleAndTitleAndDescription(offset int, limit int, Bundle_ string, Title_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and description = ?", Bundle_, Title_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndMenuName Get MenuLinkContentDatas via BundleAndTitleAndMenuName
func GetMenuLinkContentDatasByBundleAndTitleAndMenuName(offset int, limit int, Bundle_ string, Title_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and menu_name = ?", Bundle_, Title_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndLink_uri Get MenuLinkContentDatas via BundleAndTitleAndLink_uri
func GetMenuLinkContentDatasByBundleAndTitleAndLink_uri(offset int, limit int, Bundle_ string, Title_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and link_uri = ?", Bundle_, Title_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndLink_title Get MenuLinkContentDatas via BundleAndTitleAndLink_title
func GetMenuLinkContentDatasByBundleAndTitleAndLink_title(offset int, limit int, Bundle_ string, Title_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and link_title = ?", Bundle_, Title_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndLink_options Get MenuLinkContentDatas via BundleAndTitleAndLink_options
func GetMenuLinkContentDatasByBundleAndTitleAndLink_options(offset int, limit int, Bundle_ string, Title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and link_options = ?", Bundle_, Title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndExternal Get MenuLinkContentDatas via BundleAndTitleAndExternal
func GetMenuLinkContentDatasByBundleAndTitleAndExternal(offset int, limit int, Bundle_ string, Title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and external = ?", Bundle_, Title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndRediscover Get MenuLinkContentDatas via BundleAndTitleAndRediscover
func GetMenuLinkContentDatasByBundleAndTitleAndRediscover(offset int, limit int, Bundle_ string, Title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and rediscover = ?", Bundle_, Title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndWeight Get MenuLinkContentDatas via BundleAndTitleAndWeight
func GetMenuLinkContentDatasByBundleAndTitleAndWeight(offset int, limit int, Bundle_ string, Title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and weight = ?", Bundle_, Title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndExpanded Get MenuLinkContentDatas via BundleAndTitleAndExpanded
func GetMenuLinkContentDatasByBundleAndTitleAndExpanded(offset int, limit int, Bundle_ string, Title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and expanded = ?", Bundle_, Title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndEnabled Get MenuLinkContentDatas via BundleAndTitleAndEnabled
func GetMenuLinkContentDatasByBundleAndTitleAndEnabled(offset int, limit int, Bundle_ string, Title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and enabled = ?", Bundle_, Title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndParent Get MenuLinkContentDatas via BundleAndTitleAndParent
func GetMenuLinkContentDatasByBundleAndTitleAndParent(offset int, limit int, Bundle_ string, Title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and parent = ?", Bundle_, Title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndChanged Get MenuLinkContentDatas via BundleAndTitleAndChanged
func GetMenuLinkContentDatasByBundleAndTitleAndChanged(offset int, limit int, Bundle_ string, Title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and changed = ?", Bundle_, Title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitleAndDefaultLangcode Get MenuLinkContentDatas via BundleAndTitleAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndTitleAndDefaultLangcode(offset int, limit int, Bundle_ string, Title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ? and default_langcode = ?", Bundle_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndMenuName Get MenuLinkContentDatas via BundleAndDescriptionAndMenuName
func GetMenuLinkContentDatasByBundleAndDescriptionAndMenuName(offset int, limit int, Bundle_ string, Description_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and menu_name = ?", Bundle_, Description_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndLink_uri Get MenuLinkContentDatas via BundleAndDescriptionAndLink_uri
func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_uri(offset int, limit int, Bundle_ string, Description_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and link_uri = ?", Bundle_, Description_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndLink_title Get MenuLinkContentDatas via BundleAndDescriptionAndLink_title
func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_title(offset int, limit int, Bundle_ string, Description_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and link_title = ?", Bundle_, Description_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndLink_options Get MenuLinkContentDatas via BundleAndDescriptionAndLink_options
func GetMenuLinkContentDatasByBundleAndDescriptionAndLink_options(offset int, limit int, Bundle_ string, Description_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and link_options = ?", Bundle_, Description_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndExternal Get MenuLinkContentDatas via BundleAndDescriptionAndExternal
func GetMenuLinkContentDatasByBundleAndDescriptionAndExternal(offset int, limit int, Bundle_ string, Description_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and external = ?", Bundle_, Description_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndRediscover Get MenuLinkContentDatas via BundleAndDescriptionAndRediscover
func GetMenuLinkContentDatasByBundleAndDescriptionAndRediscover(offset int, limit int, Bundle_ string, Description_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and rediscover = ?", Bundle_, Description_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndWeight Get MenuLinkContentDatas via BundleAndDescriptionAndWeight
func GetMenuLinkContentDatasByBundleAndDescriptionAndWeight(offset int, limit int, Bundle_ string, Description_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and weight = ?", Bundle_, Description_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndExpanded Get MenuLinkContentDatas via BundleAndDescriptionAndExpanded
func GetMenuLinkContentDatasByBundleAndDescriptionAndExpanded(offset int, limit int, Bundle_ string, Description_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and expanded = ?", Bundle_, Description_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndEnabled Get MenuLinkContentDatas via BundleAndDescriptionAndEnabled
func GetMenuLinkContentDatasByBundleAndDescriptionAndEnabled(offset int, limit int, Bundle_ string, Description_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and enabled = ?", Bundle_, Description_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndParent Get MenuLinkContentDatas via BundleAndDescriptionAndParent
func GetMenuLinkContentDatasByBundleAndDescriptionAndParent(offset int, limit int, Bundle_ string, Description_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and parent = ?", Bundle_, Description_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndChanged Get MenuLinkContentDatas via BundleAndDescriptionAndChanged
func GetMenuLinkContentDatasByBundleAndDescriptionAndChanged(offset int, limit int, Bundle_ string, Description_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and changed = ?", Bundle_, Description_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescriptionAndDefaultLangcode Get MenuLinkContentDatas via BundleAndDescriptionAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndDescriptionAndDefaultLangcode(offset int, limit int, Bundle_ string, Description_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ? and default_langcode = ?", Bundle_, Description_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndLink_uri Get MenuLinkContentDatas via BundleAndMenuNameAndLink_uri
func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_uri(offset int, limit int, Bundle_ string, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and link_uri = ?", Bundle_, MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndLink_title Get MenuLinkContentDatas via BundleAndMenuNameAndLink_title
func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_title(offset int, limit int, Bundle_ string, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and link_title = ?", Bundle_, MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndLink_options Get MenuLinkContentDatas via BundleAndMenuNameAndLink_options
func GetMenuLinkContentDatasByBundleAndMenuNameAndLink_options(offset int, limit int, Bundle_ string, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and link_options = ?", Bundle_, MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndExternal Get MenuLinkContentDatas via BundleAndMenuNameAndExternal
func GetMenuLinkContentDatasByBundleAndMenuNameAndExternal(offset int, limit int, Bundle_ string, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and external = ?", Bundle_, MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndRediscover Get MenuLinkContentDatas via BundleAndMenuNameAndRediscover
func GetMenuLinkContentDatasByBundleAndMenuNameAndRediscover(offset int, limit int, Bundle_ string, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and rediscover = ?", Bundle_, MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndWeight Get MenuLinkContentDatas via BundleAndMenuNameAndWeight
func GetMenuLinkContentDatasByBundleAndMenuNameAndWeight(offset int, limit int, Bundle_ string, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and weight = ?", Bundle_, MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndExpanded Get MenuLinkContentDatas via BundleAndMenuNameAndExpanded
func GetMenuLinkContentDatasByBundleAndMenuNameAndExpanded(offset int, limit int, Bundle_ string, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and expanded = ?", Bundle_, MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndEnabled Get MenuLinkContentDatas via BundleAndMenuNameAndEnabled
func GetMenuLinkContentDatasByBundleAndMenuNameAndEnabled(offset int, limit int, Bundle_ string, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and enabled = ?", Bundle_, MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndParent Get MenuLinkContentDatas via BundleAndMenuNameAndParent
func GetMenuLinkContentDatasByBundleAndMenuNameAndParent(offset int, limit int, Bundle_ string, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and parent = ?", Bundle_, MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndChanged Get MenuLinkContentDatas via BundleAndMenuNameAndChanged
func GetMenuLinkContentDatasByBundleAndMenuNameAndChanged(offset int, limit int, Bundle_ string, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and changed = ?", Bundle_, MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuNameAndDefaultLangcode Get MenuLinkContentDatas via BundleAndMenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndMenuNameAndDefaultLangcode(offset int, limit int, Bundle_ string, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ? and default_langcode = ?", Bundle_, MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndLink_title Get MenuLinkContentDatas via BundleAndLink_uriAndLink_title
func GetMenuLinkContentDatasByBundleAndLink_uriAndLink_title(offset int, limit int, Bundle_ string, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and link_title = ?", Bundle_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndLink_options Get MenuLinkContentDatas via BundleAndLink_uriAndLink_options
func GetMenuLinkContentDatasByBundleAndLink_uriAndLink_options(offset int, limit int, Bundle_ string, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and link_options = ?", Bundle_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndExternal Get MenuLinkContentDatas via BundleAndLink_uriAndExternal
func GetMenuLinkContentDatasByBundleAndLink_uriAndExternal(offset int, limit int, Bundle_ string, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and external = ?", Bundle_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndRediscover Get MenuLinkContentDatas via BundleAndLink_uriAndRediscover
func GetMenuLinkContentDatasByBundleAndLink_uriAndRediscover(offset int, limit int, Bundle_ string, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and rediscover = ?", Bundle_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndWeight Get MenuLinkContentDatas via BundleAndLink_uriAndWeight
func GetMenuLinkContentDatasByBundleAndLink_uriAndWeight(offset int, limit int, Bundle_ string, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and weight = ?", Bundle_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndExpanded Get MenuLinkContentDatas via BundleAndLink_uriAndExpanded
func GetMenuLinkContentDatasByBundleAndLink_uriAndExpanded(offset int, limit int, Bundle_ string, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and expanded = ?", Bundle_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndEnabled Get MenuLinkContentDatas via BundleAndLink_uriAndEnabled
func GetMenuLinkContentDatasByBundleAndLink_uriAndEnabled(offset int, limit int, Bundle_ string, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and enabled = ?", Bundle_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndParent Get MenuLinkContentDatas via BundleAndLink_uriAndParent
func GetMenuLinkContentDatasByBundleAndLink_uriAndParent(offset int, limit int, Bundle_ string, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and parent = ?", Bundle_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndChanged Get MenuLinkContentDatas via BundleAndLink_uriAndChanged
func GetMenuLinkContentDatasByBundleAndLink_uriAndChanged(offset int, limit int, Bundle_ string, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and changed = ?", Bundle_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via BundleAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndLink_uriAndDefaultLangcode(offset int, limit int, Bundle_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ? and default_langcode = ?", Bundle_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndLink_options Get MenuLinkContentDatas via BundleAndLink_titleAndLink_options
func GetMenuLinkContentDatasByBundleAndLink_titleAndLink_options(offset int, limit int, Bundle_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and link_options = ?", Bundle_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndExternal Get MenuLinkContentDatas via BundleAndLink_titleAndExternal
func GetMenuLinkContentDatasByBundleAndLink_titleAndExternal(offset int, limit int, Bundle_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and external = ?", Bundle_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndRediscover Get MenuLinkContentDatas via BundleAndLink_titleAndRediscover
func GetMenuLinkContentDatasByBundleAndLink_titleAndRediscover(offset int, limit int, Bundle_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and rediscover = ?", Bundle_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndWeight Get MenuLinkContentDatas via BundleAndLink_titleAndWeight
func GetMenuLinkContentDatasByBundleAndLink_titleAndWeight(offset int, limit int, Bundle_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and weight = ?", Bundle_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndExpanded Get MenuLinkContentDatas via BundleAndLink_titleAndExpanded
func GetMenuLinkContentDatasByBundleAndLink_titleAndExpanded(offset int, limit int, Bundle_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and expanded = ?", Bundle_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndEnabled Get MenuLinkContentDatas via BundleAndLink_titleAndEnabled
func GetMenuLinkContentDatasByBundleAndLink_titleAndEnabled(offset int, limit int, Bundle_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and enabled = ?", Bundle_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndParent Get MenuLinkContentDatas via BundleAndLink_titleAndParent
func GetMenuLinkContentDatasByBundleAndLink_titleAndParent(offset int, limit int, Bundle_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and parent = ?", Bundle_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndChanged Get MenuLinkContentDatas via BundleAndLink_titleAndChanged
func GetMenuLinkContentDatasByBundleAndLink_titleAndChanged(offset int, limit int, Bundle_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and changed = ?", Bundle_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via BundleAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndLink_titleAndDefaultLangcode(offset int, limit int, Bundle_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ? and default_langcode = ?", Bundle_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndExternal Get MenuLinkContentDatas via BundleAndLink_optionsAndExternal
func GetMenuLinkContentDatasByBundleAndLink_optionsAndExternal(offset int, limit int, Bundle_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and external = ?", Bundle_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndRediscover Get MenuLinkContentDatas via BundleAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByBundleAndLink_optionsAndRediscover(offset int, limit int, Bundle_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and rediscover = ?", Bundle_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndWeight Get MenuLinkContentDatas via BundleAndLink_optionsAndWeight
func GetMenuLinkContentDatasByBundleAndLink_optionsAndWeight(offset int, limit int, Bundle_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and weight = ?", Bundle_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndExpanded Get MenuLinkContentDatas via BundleAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByBundleAndLink_optionsAndExpanded(offset int, limit int, Bundle_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and expanded = ?", Bundle_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndEnabled Get MenuLinkContentDatas via BundleAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByBundleAndLink_optionsAndEnabled(offset int, limit int, Bundle_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and enabled = ?", Bundle_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndParent Get MenuLinkContentDatas via BundleAndLink_optionsAndParent
func GetMenuLinkContentDatasByBundleAndLink_optionsAndParent(offset int, limit int, Bundle_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and parent = ?", Bundle_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndChanged Get MenuLinkContentDatas via BundleAndLink_optionsAndChanged
func GetMenuLinkContentDatasByBundleAndLink_optionsAndChanged(offset int, limit int, Bundle_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and changed = ?", Bundle_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via BundleAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndLink_optionsAndDefaultLangcode(offset int, limit int, Bundle_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ? and default_langcode = ?", Bundle_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndRediscover Get MenuLinkContentDatas via BundleAndExternalAndRediscover
func GetMenuLinkContentDatasByBundleAndExternalAndRediscover(offset int, limit int, Bundle_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and rediscover = ?", Bundle_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndWeight Get MenuLinkContentDatas via BundleAndExternalAndWeight
func GetMenuLinkContentDatasByBundleAndExternalAndWeight(offset int, limit int, Bundle_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and weight = ?", Bundle_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndExpanded Get MenuLinkContentDatas via BundleAndExternalAndExpanded
func GetMenuLinkContentDatasByBundleAndExternalAndExpanded(offset int, limit int, Bundle_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and expanded = ?", Bundle_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndEnabled Get MenuLinkContentDatas via BundleAndExternalAndEnabled
func GetMenuLinkContentDatasByBundleAndExternalAndEnabled(offset int, limit int, Bundle_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and enabled = ?", Bundle_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndParent Get MenuLinkContentDatas via BundleAndExternalAndParent
func GetMenuLinkContentDatasByBundleAndExternalAndParent(offset int, limit int, Bundle_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and parent = ?", Bundle_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndChanged Get MenuLinkContentDatas via BundleAndExternalAndChanged
func GetMenuLinkContentDatasByBundleAndExternalAndChanged(offset int, limit int, Bundle_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and changed = ?", Bundle_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternalAndDefaultLangcode Get MenuLinkContentDatas via BundleAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndExternalAndDefaultLangcode(offset int, limit int, Bundle_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ? and default_langcode = ?", Bundle_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndWeight Get MenuLinkContentDatas via BundleAndRediscoverAndWeight
func GetMenuLinkContentDatasByBundleAndRediscoverAndWeight(offset int, limit int, Bundle_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and weight = ?", Bundle_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndExpanded Get MenuLinkContentDatas via BundleAndRediscoverAndExpanded
func GetMenuLinkContentDatasByBundleAndRediscoverAndExpanded(offset int, limit int, Bundle_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and expanded = ?", Bundle_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndEnabled Get MenuLinkContentDatas via BundleAndRediscoverAndEnabled
func GetMenuLinkContentDatasByBundleAndRediscoverAndEnabled(offset int, limit int, Bundle_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and enabled = ?", Bundle_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndParent Get MenuLinkContentDatas via BundleAndRediscoverAndParent
func GetMenuLinkContentDatasByBundleAndRediscoverAndParent(offset int, limit int, Bundle_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and parent = ?", Bundle_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndChanged Get MenuLinkContentDatas via BundleAndRediscoverAndChanged
func GetMenuLinkContentDatasByBundleAndRediscoverAndChanged(offset int, limit int, Bundle_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and changed = ?", Bundle_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via BundleAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndRediscoverAndDefaultLangcode(offset int, limit int, Bundle_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ? and default_langcode = ?", Bundle_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeightAndExpanded Get MenuLinkContentDatas via BundleAndWeightAndExpanded
func GetMenuLinkContentDatasByBundleAndWeightAndExpanded(offset int, limit int, Bundle_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ? and expanded = ?", Bundle_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeightAndEnabled Get MenuLinkContentDatas via BundleAndWeightAndEnabled
func GetMenuLinkContentDatasByBundleAndWeightAndEnabled(offset int, limit int, Bundle_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ? and enabled = ?", Bundle_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeightAndParent Get MenuLinkContentDatas via BundleAndWeightAndParent
func GetMenuLinkContentDatasByBundleAndWeightAndParent(offset int, limit int, Bundle_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ? and parent = ?", Bundle_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeightAndChanged Get MenuLinkContentDatas via BundleAndWeightAndChanged
func GetMenuLinkContentDatasByBundleAndWeightAndChanged(offset int, limit int, Bundle_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ? and changed = ?", Bundle_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeightAndDefaultLangcode Get MenuLinkContentDatas via BundleAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndWeightAndDefaultLangcode(offset int, limit int, Bundle_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ? and default_langcode = ?", Bundle_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExpandedAndEnabled Get MenuLinkContentDatas via BundleAndExpandedAndEnabled
func GetMenuLinkContentDatasByBundleAndExpandedAndEnabled(offset int, limit int, Bundle_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and expanded = ? and enabled = ?", Bundle_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExpandedAndParent Get MenuLinkContentDatas via BundleAndExpandedAndParent
func GetMenuLinkContentDatasByBundleAndExpandedAndParent(offset int, limit int, Bundle_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and expanded = ? and parent = ?", Bundle_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExpandedAndChanged Get MenuLinkContentDatas via BundleAndExpandedAndChanged
func GetMenuLinkContentDatasByBundleAndExpandedAndChanged(offset int, limit int, Bundle_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and expanded = ? and changed = ?", Bundle_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via BundleAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndExpandedAndDefaultLangcode(offset int, limit int, Bundle_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and expanded = ? and default_langcode = ?", Bundle_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndEnabledAndParent Get MenuLinkContentDatas via BundleAndEnabledAndParent
func GetMenuLinkContentDatasByBundleAndEnabledAndParent(offset int, limit int, Bundle_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and enabled = ? and parent = ?", Bundle_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndEnabledAndChanged Get MenuLinkContentDatas via BundleAndEnabledAndChanged
func GetMenuLinkContentDatasByBundleAndEnabledAndChanged(offset int, limit int, Bundle_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and enabled = ? and changed = ?", Bundle_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via BundleAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndEnabledAndDefaultLangcode(offset int, limit int, Bundle_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and enabled = ? and default_langcode = ?", Bundle_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndParentAndChanged Get MenuLinkContentDatas via BundleAndParentAndChanged
func GetMenuLinkContentDatasByBundleAndParentAndChanged(offset int, limit int, Bundle_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and parent = ? and changed = ?", Bundle_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndParentAndDefaultLangcode Get MenuLinkContentDatas via BundleAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndParentAndDefaultLangcode(offset int, limit int, Bundle_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and parent = ? and default_langcode = ?", Bundle_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndChangedAndDefaultLangcode Get MenuLinkContentDatas via BundleAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndChangedAndDefaultLangcode(offset int, limit int, Bundle_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and changed = ? and default_langcode = ?", Bundle_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndDescription Get MenuLinkContentDatas via LangcodeAndTitleAndDescription
func GetMenuLinkContentDatasByLangcodeAndTitleAndDescription(offset int, limit int, Langcode_ string, Title_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and description = ?", Langcode_, Title_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndMenuName Get MenuLinkContentDatas via LangcodeAndTitleAndMenuName
func GetMenuLinkContentDatasByLangcodeAndTitleAndMenuName(offset int, limit int, Langcode_ string, Title_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and menu_name = ?", Langcode_, Title_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndLink_uri Get MenuLinkContentDatas via LangcodeAndTitleAndLink_uri
func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_uri(offset int, limit int, Langcode_ string, Title_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and link_uri = ?", Langcode_, Title_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndLink_title Get MenuLinkContentDatas via LangcodeAndTitleAndLink_title
func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_title(offset int, limit int, Langcode_ string, Title_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and link_title = ?", Langcode_, Title_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndLink_options Get MenuLinkContentDatas via LangcodeAndTitleAndLink_options
func GetMenuLinkContentDatasByLangcodeAndTitleAndLink_options(offset int, limit int, Langcode_ string, Title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and link_options = ?", Langcode_, Title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndExternal Get MenuLinkContentDatas via LangcodeAndTitleAndExternal
func GetMenuLinkContentDatasByLangcodeAndTitleAndExternal(offset int, limit int, Langcode_ string, Title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and external = ?", Langcode_, Title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndRediscover Get MenuLinkContentDatas via LangcodeAndTitleAndRediscover
func GetMenuLinkContentDatasByLangcodeAndTitleAndRediscover(offset int, limit int, Langcode_ string, Title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and rediscover = ?", Langcode_, Title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndWeight Get MenuLinkContentDatas via LangcodeAndTitleAndWeight
func GetMenuLinkContentDatasByLangcodeAndTitleAndWeight(offset int, limit int, Langcode_ string, Title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and weight = ?", Langcode_, Title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndExpanded Get MenuLinkContentDatas via LangcodeAndTitleAndExpanded
func GetMenuLinkContentDatasByLangcodeAndTitleAndExpanded(offset int, limit int, Langcode_ string, Title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and expanded = ?", Langcode_, Title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndEnabled Get MenuLinkContentDatas via LangcodeAndTitleAndEnabled
func GetMenuLinkContentDatasByLangcodeAndTitleAndEnabled(offset int, limit int, Langcode_ string, Title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and enabled = ?", Langcode_, Title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndParent Get MenuLinkContentDatas via LangcodeAndTitleAndParent
func GetMenuLinkContentDatasByLangcodeAndTitleAndParent(offset int, limit int, Langcode_ string, Title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and parent = ?", Langcode_, Title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndChanged Get MenuLinkContentDatas via LangcodeAndTitleAndChanged
func GetMenuLinkContentDatasByLangcodeAndTitleAndChanged(offset int, limit int, Langcode_ string, Title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and changed = ?", Langcode_, Title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitleAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndTitleAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndTitleAndDefaultLangcode(offset int, limit int, Langcode_ string, Title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ? and default_langcode = ?", Langcode_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndMenuName Get MenuLinkContentDatas via LangcodeAndDescriptionAndMenuName
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndMenuName(offset int, limit int, Langcode_ string, Description_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and menu_name = ?", Langcode_, Description_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_uri Get MenuLinkContentDatas via LangcodeAndDescriptionAndLink_uri
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_uri(offset int, limit int, Langcode_ string, Description_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and link_uri = ?", Langcode_, Description_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_title Get MenuLinkContentDatas via LangcodeAndDescriptionAndLink_title
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_title(offset int, limit int, Langcode_ string, Description_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and link_title = ?", Langcode_, Description_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_options Get MenuLinkContentDatas via LangcodeAndDescriptionAndLink_options
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndLink_options(offset int, limit int, Langcode_ string, Description_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and link_options = ?", Langcode_, Description_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndExternal Get MenuLinkContentDatas via LangcodeAndDescriptionAndExternal
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndExternal(offset int, limit int, Langcode_ string, Description_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and external = ?", Langcode_, Description_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndRediscover Get MenuLinkContentDatas via LangcodeAndDescriptionAndRediscover
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndRediscover(offset int, limit int, Langcode_ string, Description_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and rediscover = ?", Langcode_, Description_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndWeight Get MenuLinkContentDatas via LangcodeAndDescriptionAndWeight
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndWeight(offset int, limit int, Langcode_ string, Description_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and weight = ?", Langcode_, Description_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndExpanded Get MenuLinkContentDatas via LangcodeAndDescriptionAndExpanded
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndExpanded(offset int, limit int, Langcode_ string, Description_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and expanded = ?", Langcode_, Description_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndEnabled Get MenuLinkContentDatas via LangcodeAndDescriptionAndEnabled
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndEnabled(offset int, limit int, Langcode_ string, Description_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and enabled = ?", Langcode_, Description_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndParent Get MenuLinkContentDatas via LangcodeAndDescriptionAndParent
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndParent(offset int, limit int, Langcode_ string, Description_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and parent = ?", Langcode_, Description_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndChanged Get MenuLinkContentDatas via LangcodeAndDescriptionAndChanged
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndChanged(offset int, limit int, Langcode_ string, Description_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and changed = ?", Langcode_, Description_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescriptionAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndDescriptionAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndDescriptionAndDefaultLangcode(offset int, limit int, Langcode_ string, Description_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ? and default_langcode = ?", Langcode_, Description_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_uri Get MenuLinkContentDatas via LangcodeAndMenuNameAndLink_uri
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_uri(offset int, limit int, Langcode_ string, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and link_uri = ?", Langcode_, MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_title Get MenuLinkContentDatas via LangcodeAndMenuNameAndLink_title
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_title(offset int, limit int, Langcode_ string, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and link_title = ?", Langcode_, MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_options Get MenuLinkContentDatas via LangcodeAndMenuNameAndLink_options
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndLink_options(offset int, limit int, Langcode_ string, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and link_options = ?", Langcode_, MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndExternal Get MenuLinkContentDatas via LangcodeAndMenuNameAndExternal
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndExternal(offset int, limit int, Langcode_ string, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and external = ?", Langcode_, MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndRediscover Get MenuLinkContentDatas via LangcodeAndMenuNameAndRediscover
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndRediscover(offset int, limit int, Langcode_ string, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and rediscover = ?", Langcode_, MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndWeight Get MenuLinkContentDatas via LangcodeAndMenuNameAndWeight
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndWeight(offset int, limit int, Langcode_ string, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and weight = ?", Langcode_, MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndExpanded Get MenuLinkContentDatas via LangcodeAndMenuNameAndExpanded
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndExpanded(offset int, limit int, Langcode_ string, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and expanded = ?", Langcode_, MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndEnabled Get MenuLinkContentDatas via LangcodeAndMenuNameAndEnabled
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndEnabled(offset int, limit int, Langcode_ string, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and enabled = ?", Langcode_, MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndParent Get MenuLinkContentDatas via LangcodeAndMenuNameAndParent
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndParent(offset int, limit int, Langcode_ string, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and parent = ?", Langcode_, MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndChanged Get MenuLinkContentDatas via LangcodeAndMenuNameAndChanged
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndChanged(offset int, limit int, Langcode_ string, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and changed = ?", Langcode_, MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuNameAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndMenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndMenuNameAndDefaultLangcode(offset int, limit int, Langcode_ string, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ? and default_langcode = ?", Langcode_, MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_title Get MenuLinkContentDatas via LangcodeAndLink_uriAndLink_title
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_title(offset int, limit int, Langcode_ string, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and link_title = ?", Langcode_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_options Get MenuLinkContentDatas via LangcodeAndLink_uriAndLink_options
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndLink_options(offset int, limit int, Langcode_ string, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and link_options = ?", Langcode_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndExternal Get MenuLinkContentDatas via LangcodeAndLink_uriAndExternal
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndExternal(offset int, limit int, Langcode_ string, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and external = ?", Langcode_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndRediscover Get MenuLinkContentDatas via LangcodeAndLink_uriAndRediscover
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndRediscover(offset int, limit int, Langcode_ string, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and rediscover = ?", Langcode_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndWeight Get MenuLinkContentDatas via LangcodeAndLink_uriAndWeight
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndWeight(offset int, limit int, Langcode_ string, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and weight = ?", Langcode_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndExpanded Get MenuLinkContentDatas via LangcodeAndLink_uriAndExpanded
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndExpanded(offset int, limit int, Langcode_ string, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and expanded = ?", Langcode_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndEnabled Get MenuLinkContentDatas via LangcodeAndLink_uriAndEnabled
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndEnabled(offset int, limit int, Langcode_ string, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and enabled = ?", Langcode_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndParent Get MenuLinkContentDatas via LangcodeAndLink_uriAndParent
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndParent(offset int, limit int, Langcode_ string, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and parent = ?", Langcode_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndChanged Get MenuLinkContentDatas via LangcodeAndLink_uriAndChanged
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndChanged(offset int, limit int, Langcode_ string, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and changed = ?", Langcode_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndLink_uriAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ? and default_langcode = ?", Langcode_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndLink_options Get MenuLinkContentDatas via LangcodeAndLink_titleAndLink_options
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndLink_options(offset int, limit int, Langcode_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and link_options = ?", Langcode_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndExternal Get MenuLinkContentDatas via LangcodeAndLink_titleAndExternal
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndExternal(offset int, limit int, Langcode_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and external = ?", Langcode_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndRediscover Get MenuLinkContentDatas via LangcodeAndLink_titleAndRediscover
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndRediscover(offset int, limit int, Langcode_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and rediscover = ?", Langcode_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndWeight Get MenuLinkContentDatas via LangcodeAndLink_titleAndWeight
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndWeight(offset int, limit int, Langcode_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and weight = ?", Langcode_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndExpanded Get MenuLinkContentDatas via LangcodeAndLink_titleAndExpanded
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndExpanded(offset int, limit int, Langcode_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and expanded = ?", Langcode_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndEnabled Get MenuLinkContentDatas via LangcodeAndLink_titleAndEnabled
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndEnabled(offset int, limit int, Langcode_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and enabled = ?", Langcode_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndParent Get MenuLinkContentDatas via LangcodeAndLink_titleAndParent
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndParent(offset int, limit int, Langcode_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and parent = ?", Langcode_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndChanged Get MenuLinkContentDatas via LangcodeAndLink_titleAndChanged
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndChanged(offset int, limit int, Langcode_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and changed = ?", Langcode_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndLink_titleAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ? and default_langcode = ?", Langcode_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExternal Get MenuLinkContentDatas via LangcodeAndLink_optionsAndExternal
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExternal(offset int, limit int, Langcode_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and external = ?", Langcode_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndRediscover Get MenuLinkContentDatas via LangcodeAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndRediscover(offset int, limit int, Langcode_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and rediscover = ?", Langcode_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndWeight Get MenuLinkContentDatas via LangcodeAndLink_optionsAndWeight
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndWeight(offset int, limit int, Langcode_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and weight = ?", Langcode_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExpanded Get MenuLinkContentDatas via LangcodeAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndExpanded(offset int, limit int, Langcode_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and expanded = ?", Langcode_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndEnabled Get MenuLinkContentDatas via LangcodeAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndEnabled(offset int, limit int, Langcode_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and enabled = ?", Langcode_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndParent Get MenuLinkContentDatas via LangcodeAndLink_optionsAndParent
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndParent(offset int, limit int, Langcode_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and parent = ?", Langcode_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndChanged Get MenuLinkContentDatas via LangcodeAndLink_optionsAndChanged
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndChanged(offset int, limit int, Langcode_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and changed = ?", Langcode_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndLink_optionsAndDefaultLangcode(offset int, limit int, Langcode_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ? and default_langcode = ?", Langcode_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndRediscover Get MenuLinkContentDatas via LangcodeAndExternalAndRediscover
func GetMenuLinkContentDatasByLangcodeAndExternalAndRediscover(offset int, limit int, Langcode_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and rediscover = ?", Langcode_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndWeight Get MenuLinkContentDatas via LangcodeAndExternalAndWeight
func GetMenuLinkContentDatasByLangcodeAndExternalAndWeight(offset int, limit int, Langcode_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and weight = ?", Langcode_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndExpanded Get MenuLinkContentDatas via LangcodeAndExternalAndExpanded
func GetMenuLinkContentDatasByLangcodeAndExternalAndExpanded(offset int, limit int, Langcode_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and expanded = ?", Langcode_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndEnabled Get MenuLinkContentDatas via LangcodeAndExternalAndEnabled
func GetMenuLinkContentDatasByLangcodeAndExternalAndEnabled(offset int, limit int, Langcode_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and enabled = ?", Langcode_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndParent Get MenuLinkContentDatas via LangcodeAndExternalAndParent
func GetMenuLinkContentDatasByLangcodeAndExternalAndParent(offset int, limit int, Langcode_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and parent = ?", Langcode_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndChanged Get MenuLinkContentDatas via LangcodeAndExternalAndChanged
func GetMenuLinkContentDatasByLangcodeAndExternalAndChanged(offset int, limit int, Langcode_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and changed = ?", Langcode_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternalAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndExternalAndDefaultLangcode(offset int, limit int, Langcode_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ? and default_langcode = ?", Langcode_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndWeight Get MenuLinkContentDatas via LangcodeAndRediscoverAndWeight
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndWeight(offset int, limit int, Langcode_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and weight = ?", Langcode_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndExpanded Get MenuLinkContentDatas via LangcodeAndRediscoverAndExpanded
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndExpanded(offset int, limit int, Langcode_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and expanded = ?", Langcode_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndEnabled Get MenuLinkContentDatas via LangcodeAndRediscoverAndEnabled
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndEnabled(offset int, limit int, Langcode_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and enabled = ?", Langcode_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndParent Get MenuLinkContentDatas via LangcodeAndRediscoverAndParent
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndParent(offset int, limit int, Langcode_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and parent = ?", Langcode_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndChanged Get MenuLinkContentDatas via LangcodeAndRediscoverAndChanged
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndChanged(offset int, limit int, Langcode_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and changed = ?", Langcode_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndRediscoverAndDefaultLangcode(offset int, limit int, Langcode_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ? and default_langcode = ?", Langcode_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeightAndExpanded Get MenuLinkContentDatas via LangcodeAndWeightAndExpanded
func GetMenuLinkContentDatasByLangcodeAndWeightAndExpanded(offset int, limit int, Langcode_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ? and expanded = ?", Langcode_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeightAndEnabled Get MenuLinkContentDatas via LangcodeAndWeightAndEnabled
func GetMenuLinkContentDatasByLangcodeAndWeightAndEnabled(offset int, limit int, Langcode_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ? and enabled = ?", Langcode_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeightAndParent Get MenuLinkContentDatas via LangcodeAndWeightAndParent
func GetMenuLinkContentDatasByLangcodeAndWeightAndParent(offset int, limit int, Langcode_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ? and parent = ?", Langcode_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeightAndChanged Get MenuLinkContentDatas via LangcodeAndWeightAndChanged
func GetMenuLinkContentDatasByLangcodeAndWeightAndChanged(offset int, limit int, Langcode_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ? and changed = ?", Langcode_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeightAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndWeightAndDefaultLangcode(offset int, limit int, Langcode_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ? and default_langcode = ?", Langcode_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExpandedAndEnabled Get MenuLinkContentDatas via LangcodeAndExpandedAndEnabled
func GetMenuLinkContentDatasByLangcodeAndExpandedAndEnabled(offset int, limit int, Langcode_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and expanded = ? and enabled = ?", Langcode_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExpandedAndParent Get MenuLinkContentDatas via LangcodeAndExpandedAndParent
func GetMenuLinkContentDatasByLangcodeAndExpandedAndParent(offset int, limit int, Langcode_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and expanded = ? and parent = ?", Langcode_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExpandedAndChanged Get MenuLinkContentDatas via LangcodeAndExpandedAndChanged
func GetMenuLinkContentDatasByLangcodeAndExpandedAndChanged(offset int, limit int, Langcode_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and expanded = ? and changed = ?", Langcode_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndExpandedAndDefaultLangcode(offset int, limit int, Langcode_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and expanded = ? and default_langcode = ?", Langcode_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndEnabledAndParent Get MenuLinkContentDatas via LangcodeAndEnabledAndParent
func GetMenuLinkContentDatasByLangcodeAndEnabledAndParent(offset int, limit int, Langcode_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and enabled = ? and parent = ?", Langcode_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndEnabledAndChanged Get MenuLinkContentDatas via LangcodeAndEnabledAndChanged
func GetMenuLinkContentDatasByLangcodeAndEnabledAndChanged(offset int, limit int, Langcode_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and enabled = ? and changed = ?", Langcode_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndEnabledAndDefaultLangcode(offset int, limit int, Langcode_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and enabled = ? and default_langcode = ?", Langcode_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndParentAndChanged Get MenuLinkContentDatas via LangcodeAndParentAndChanged
func GetMenuLinkContentDatasByLangcodeAndParentAndChanged(offset int, limit int, Langcode_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and parent = ? and changed = ?", Langcode_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndParentAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndParentAndDefaultLangcode(offset int, limit int, Langcode_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and parent = ? and default_langcode = ?", Langcode_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndChangedAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndMenuName Get MenuLinkContentDatas via TitleAndDescriptionAndMenuName
func GetMenuLinkContentDatasByTitleAndDescriptionAndMenuName(offset int, limit int, Title_ string, Description_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and menu_name = ?", Title_, Description_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndLink_uri Get MenuLinkContentDatas via TitleAndDescriptionAndLink_uri
func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_uri(offset int, limit int, Title_ string, Description_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and link_uri = ?", Title_, Description_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndLink_title Get MenuLinkContentDatas via TitleAndDescriptionAndLink_title
func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_title(offset int, limit int, Title_ string, Description_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and link_title = ?", Title_, Description_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndLink_options Get MenuLinkContentDatas via TitleAndDescriptionAndLink_options
func GetMenuLinkContentDatasByTitleAndDescriptionAndLink_options(offset int, limit int, Title_ string, Description_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and link_options = ?", Title_, Description_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndExternal Get MenuLinkContentDatas via TitleAndDescriptionAndExternal
func GetMenuLinkContentDatasByTitleAndDescriptionAndExternal(offset int, limit int, Title_ string, Description_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and external = ?", Title_, Description_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndRediscover Get MenuLinkContentDatas via TitleAndDescriptionAndRediscover
func GetMenuLinkContentDatasByTitleAndDescriptionAndRediscover(offset int, limit int, Title_ string, Description_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and rediscover = ?", Title_, Description_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndWeight Get MenuLinkContentDatas via TitleAndDescriptionAndWeight
func GetMenuLinkContentDatasByTitleAndDescriptionAndWeight(offset int, limit int, Title_ string, Description_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and weight = ?", Title_, Description_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndExpanded Get MenuLinkContentDatas via TitleAndDescriptionAndExpanded
func GetMenuLinkContentDatasByTitleAndDescriptionAndExpanded(offset int, limit int, Title_ string, Description_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and expanded = ?", Title_, Description_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndEnabled Get MenuLinkContentDatas via TitleAndDescriptionAndEnabled
func GetMenuLinkContentDatasByTitleAndDescriptionAndEnabled(offset int, limit int, Title_ string, Description_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and enabled = ?", Title_, Description_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndParent Get MenuLinkContentDatas via TitleAndDescriptionAndParent
func GetMenuLinkContentDatasByTitleAndDescriptionAndParent(offset int, limit int, Title_ string, Description_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and parent = ?", Title_, Description_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndChanged Get MenuLinkContentDatas via TitleAndDescriptionAndChanged
func GetMenuLinkContentDatasByTitleAndDescriptionAndChanged(offset int, limit int, Title_ string, Description_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and changed = ?", Title_, Description_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescriptionAndDefaultLangcode Get MenuLinkContentDatas via TitleAndDescriptionAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndDescriptionAndDefaultLangcode(offset int, limit int, Title_ string, Description_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ? and default_langcode = ?", Title_, Description_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndLink_uri Get MenuLinkContentDatas via TitleAndMenuNameAndLink_uri
func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_uri(offset int, limit int, Title_ string, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and link_uri = ?", Title_, MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndLink_title Get MenuLinkContentDatas via TitleAndMenuNameAndLink_title
func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_title(offset int, limit int, Title_ string, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and link_title = ?", Title_, MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndLink_options Get MenuLinkContentDatas via TitleAndMenuNameAndLink_options
func GetMenuLinkContentDatasByTitleAndMenuNameAndLink_options(offset int, limit int, Title_ string, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and link_options = ?", Title_, MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndExternal Get MenuLinkContentDatas via TitleAndMenuNameAndExternal
func GetMenuLinkContentDatasByTitleAndMenuNameAndExternal(offset int, limit int, Title_ string, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and external = ?", Title_, MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndRediscover Get MenuLinkContentDatas via TitleAndMenuNameAndRediscover
func GetMenuLinkContentDatasByTitleAndMenuNameAndRediscover(offset int, limit int, Title_ string, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and rediscover = ?", Title_, MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndWeight Get MenuLinkContentDatas via TitleAndMenuNameAndWeight
func GetMenuLinkContentDatasByTitleAndMenuNameAndWeight(offset int, limit int, Title_ string, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and weight = ?", Title_, MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndExpanded Get MenuLinkContentDatas via TitleAndMenuNameAndExpanded
func GetMenuLinkContentDatasByTitleAndMenuNameAndExpanded(offset int, limit int, Title_ string, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and expanded = ?", Title_, MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndEnabled Get MenuLinkContentDatas via TitleAndMenuNameAndEnabled
func GetMenuLinkContentDatasByTitleAndMenuNameAndEnabled(offset int, limit int, Title_ string, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and enabled = ?", Title_, MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndParent Get MenuLinkContentDatas via TitleAndMenuNameAndParent
func GetMenuLinkContentDatasByTitleAndMenuNameAndParent(offset int, limit int, Title_ string, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and parent = ?", Title_, MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndChanged Get MenuLinkContentDatas via TitleAndMenuNameAndChanged
func GetMenuLinkContentDatasByTitleAndMenuNameAndChanged(offset int, limit int, Title_ string, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and changed = ?", Title_, MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuNameAndDefaultLangcode Get MenuLinkContentDatas via TitleAndMenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndMenuNameAndDefaultLangcode(offset int, limit int, Title_ string, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ? and default_langcode = ?", Title_, MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndLink_title Get MenuLinkContentDatas via TitleAndLink_uriAndLink_title
func GetMenuLinkContentDatasByTitleAndLink_uriAndLink_title(offset int, limit int, Title_ string, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and link_title = ?", Title_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndLink_options Get MenuLinkContentDatas via TitleAndLink_uriAndLink_options
func GetMenuLinkContentDatasByTitleAndLink_uriAndLink_options(offset int, limit int, Title_ string, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and link_options = ?", Title_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndExternal Get MenuLinkContentDatas via TitleAndLink_uriAndExternal
func GetMenuLinkContentDatasByTitleAndLink_uriAndExternal(offset int, limit int, Title_ string, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and external = ?", Title_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndRediscover Get MenuLinkContentDatas via TitleAndLink_uriAndRediscover
func GetMenuLinkContentDatasByTitleAndLink_uriAndRediscover(offset int, limit int, Title_ string, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and rediscover = ?", Title_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndWeight Get MenuLinkContentDatas via TitleAndLink_uriAndWeight
func GetMenuLinkContentDatasByTitleAndLink_uriAndWeight(offset int, limit int, Title_ string, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and weight = ?", Title_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndExpanded Get MenuLinkContentDatas via TitleAndLink_uriAndExpanded
func GetMenuLinkContentDatasByTitleAndLink_uriAndExpanded(offset int, limit int, Title_ string, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and expanded = ?", Title_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndEnabled Get MenuLinkContentDatas via TitleAndLink_uriAndEnabled
func GetMenuLinkContentDatasByTitleAndLink_uriAndEnabled(offset int, limit int, Title_ string, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and enabled = ?", Title_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndParent Get MenuLinkContentDatas via TitleAndLink_uriAndParent
func GetMenuLinkContentDatasByTitleAndLink_uriAndParent(offset int, limit int, Title_ string, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and parent = ?", Title_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndChanged Get MenuLinkContentDatas via TitleAndLink_uriAndChanged
func GetMenuLinkContentDatasByTitleAndLink_uriAndChanged(offset int, limit int, Title_ string, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and changed = ?", Title_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via TitleAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndLink_uriAndDefaultLangcode(offset int, limit int, Title_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ? and default_langcode = ?", Title_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndLink_options Get MenuLinkContentDatas via TitleAndLink_titleAndLink_options
func GetMenuLinkContentDatasByTitleAndLink_titleAndLink_options(offset int, limit int, Title_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and link_options = ?", Title_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndExternal Get MenuLinkContentDatas via TitleAndLink_titleAndExternal
func GetMenuLinkContentDatasByTitleAndLink_titleAndExternal(offset int, limit int, Title_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and external = ?", Title_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndRediscover Get MenuLinkContentDatas via TitleAndLink_titleAndRediscover
func GetMenuLinkContentDatasByTitleAndLink_titleAndRediscover(offset int, limit int, Title_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and rediscover = ?", Title_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndWeight Get MenuLinkContentDatas via TitleAndLink_titleAndWeight
func GetMenuLinkContentDatasByTitleAndLink_titleAndWeight(offset int, limit int, Title_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and weight = ?", Title_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndExpanded Get MenuLinkContentDatas via TitleAndLink_titleAndExpanded
func GetMenuLinkContentDatasByTitleAndLink_titleAndExpanded(offset int, limit int, Title_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and expanded = ?", Title_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndEnabled Get MenuLinkContentDatas via TitleAndLink_titleAndEnabled
func GetMenuLinkContentDatasByTitleAndLink_titleAndEnabled(offset int, limit int, Title_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and enabled = ?", Title_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndParent Get MenuLinkContentDatas via TitleAndLink_titleAndParent
func GetMenuLinkContentDatasByTitleAndLink_titleAndParent(offset int, limit int, Title_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and parent = ?", Title_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndChanged Get MenuLinkContentDatas via TitleAndLink_titleAndChanged
func GetMenuLinkContentDatasByTitleAndLink_titleAndChanged(offset int, limit int, Title_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and changed = ?", Title_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via TitleAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndLink_titleAndDefaultLangcode(offset int, limit int, Title_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ? and default_langcode = ?", Title_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndExternal Get MenuLinkContentDatas via TitleAndLink_optionsAndExternal
func GetMenuLinkContentDatasByTitleAndLink_optionsAndExternal(offset int, limit int, Title_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and external = ?", Title_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndRediscover Get MenuLinkContentDatas via TitleAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByTitleAndLink_optionsAndRediscover(offset int, limit int, Title_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and rediscover = ?", Title_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndWeight Get MenuLinkContentDatas via TitleAndLink_optionsAndWeight
func GetMenuLinkContentDatasByTitleAndLink_optionsAndWeight(offset int, limit int, Title_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and weight = ?", Title_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndExpanded Get MenuLinkContentDatas via TitleAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByTitleAndLink_optionsAndExpanded(offset int, limit int, Title_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and expanded = ?", Title_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndEnabled Get MenuLinkContentDatas via TitleAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByTitleAndLink_optionsAndEnabled(offset int, limit int, Title_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and enabled = ?", Title_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndParent Get MenuLinkContentDatas via TitleAndLink_optionsAndParent
func GetMenuLinkContentDatasByTitleAndLink_optionsAndParent(offset int, limit int, Title_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and parent = ?", Title_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndChanged Get MenuLinkContentDatas via TitleAndLink_optionsAndChanged
func GetMenuLinkContentDatasByTitleAndLink_optionsAndChanged(offset int, limit int, Title_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and changed = ?", Title_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via TitleAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndLink_optionsAndDefaultLangcode(offset int, limit int, Title_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ? and default_langcode = ?", Title_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndRediscover Get MenuLinkContentDatas via TitleAndExternalAndRediscover
func GetMenuLinkContentDatasByTitleAndExternalAndRediscover(offset int, limit int, Title_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and rediscover = ?", Title_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndWeight Get MenuLinkContentDatas via TitleAndExternalAndWeight
func GetMenuLinkContentDatasByTitleAndExternalAndWeight(offset int, limit int, Title_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and weight = ?", Title_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndExpanded Get MenuLinkContentDatas via TitleAndExternalAndExpanded
func GetMenuLinkContentDatasByTitleAndExternalAndExpanded(offset int, limit int, Title_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and expanded = ?", Title_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndEnabled Get MenuLinkContentDatas via TitleAndExternalAndEnabled
func GetMenuLinkContentDatasByTitleAndExternalAndEnabled(offset int, limit int, Title_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and enabled = ?", Title_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndParent Get MenuLinkContentDatas via TitleAndExternalAndParent
func GetMenuLinkContentDatasByTitleAndExternalAndParent(offset int, limit int, Title_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and parent = ?", Title_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndChanged Get MenuLinkContentDatas via TitleAndExternalAndChanged
func GetMenuLinkContentDatasByTitleAndExternalAndChanged(offset int, limit int, Title_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and changed = ?", Title_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternalAndDefaultLangcode Get MenuLinkContentDatas via TitleAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndExternalAndDefaultLangcode(offset int, limit int, Title_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ? and default_langcode = ?", Title_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndWeight Get MenuLinkContentDatas via TitleAndRediscoverAndWeight
func GetMenuLinkContentDatasByTitleAndRediscoverAndWeight(offset int, limit int, Title_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and weight = ?", Title_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndExpanded Get MenuLinkContentDatas via TitleAndRediscoverAndExpanded
func GetMenuLinkContentDatasByTitleAndRediscoverAndExpanded(offset int, limit int, Title_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and expanded = ?", Title_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndEnabled Get MenuLinkContentDatas via TitleAndRediscoverAndEnabled
func GetMenuLinkContentDatasByTitleAndRediscoverAndEnabled(offset int, limit int, Title_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and enabled = ?", Title_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndParent Get MenuLinkContentDatas via TitleAndRediscoverAndParent
func GetMenuLinkContentDatasByTitleAndRediscoverAndParent(offset int, limit int, Title_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and parent = ?", Title_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndChanged Get MenuLinkContentDatas via TitleAndRediscoverAndChanged
func GetMenuLinkContentDatasByTitleAndRediscoverAndChanged(offset int, limit int, Title_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and changed = ?", Title_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via TitleAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndRediscoverAndDefaultLangcode(offset int, limit int, Title_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ? and default_langcode = ?", Title_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeightAndExpanded Get MenuLinkContentDatas via TitleAndWeightAndExpanded
func GetMenuLinkContentDatasByTitleAndWeightAndExpanded(offset int, limit int, Title_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ? and expanded = ?", Title_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeightAndEnabled Get MenuLinkContentDatas via TitleAndWeightAndEnabled
func GetMenuLinkContentDatasByTitleAndWeightAndEnabled(offset int, limit int, Title_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ? and enabled = ?", Title_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeightAndParent Get MenuLinkContentDatas via TitleAndWeightAndParent
func GetMenuLinkContentDatasByTitleAndWeightAndParent(offset int, limit int, Title_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ? and parent = ?", Title_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeightAndChanged Get MenuLinkContentDatas via TitleAndWeightAndChanged
func GetMenuLinkContentDatasByTitleAndWeightAndChanged(offset int, limit int, Title_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ? and changed = ?", Title_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeightAndDefaultLangcode Get MenuLinkContentDatas via TitleAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndWeightAndDefaultLangcode(offset int, limit int, Title_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ? and default_langcode = ?", Title_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExpandedAndEnabled Get MenuLinkContentDatas via TitleAndExpandedAndEnabled
func GetMenuLinkContentDatasByTitleAndExpandedAndEnabled(offset int, limit int, Title_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and expanded = ? and enabled = ?", Title_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExpandedAndParent Get MenuLinkContentDatas via TitleAndExpandedAndParent
func GetMenuLinkContentDatasByTitleAndExpandedAndParent(offset int, limit int, Title_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and expanded = ? and parent = ?", Title_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExpandedAndChanged Get MenuLinkContentDatas via TitleAndExpandedAndChanged
func GetMenuLinkContentDatasByTitleAndExpandedAndChanged(offset int, limit int, Title_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and expanded = ? and changed = ?", Title_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via TitleAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndExpandedAndDefaultLangcode(offset int, limit int, Title_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and expanded = ? and default_langcode = ?", Title_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndEnabledAndParent Get MenuLinkContentDatas via TitleAndEnabledAndParent
func GetMenuLinkContentDatasByTitleAndEnabledAndParent(offset int, limit int, Title_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and enabled = ? and parent = ?", Title_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndEnabledAndChanged Get MenuLinkContentDatas via TitleAndEnabledAndChanged
func GetMenuLinkContentDatasByTitleAndEnabledAndChanged(offset int, limit int, Title_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and enabled = ? and changed = ?", Title_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via TitleAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndEnabledAndDefaultLangcode(offset int, limit int, Title_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and enabled = ? and default_langcode = ?", Title_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndParentAndChanged Get MenuLinkContentDatas via TitleAndParentAndChanged
func GetMenuLinkContentDatasByTitleAndParentAndChanged(offset int, limit int, Title_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and parent = ? and changed = ?", Title_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndParentAndDefaultLangcode Get MenuLinkContentDatas via TitleAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndParentAndDefaultLangcode(offset int, limit int, Title_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and parent = ? and default_langcode = ?", Title_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndChangedAndDefaultLangcode Get MenuLinkContentDatas via TitleAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndChangedAndDefaultLangcode(offset int, limit int, Title_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and changed = ? and default_langcode = ?", Title_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_uri Get MenuLinkContentDatas via DescriptionAndMenuNameAndLink_uri
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_uri(offset int, limit int, Description_ string, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and link_uri = ?", Description_, MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_title Get MenuLinkContentDatas via DescriptionAndMenuNameAndLink_title
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_title(offset int, limit int, Description_ string, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and link_title = ?", Description_, MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_options Get MenuLinkContentDatas via DescriptionAndMenuNameAndLink_options
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndLink_options(offset int, limit int, Description_ string, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and link_options = ?", Description_, MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndExternal Get MenuLinkContentDatas via DescriptionAndMenuNameAndExternal
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndExternal(offset int, limit int, Description_ string, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and external = ?", Description_, MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndRediscover Get MenuLinkContentDatas via DescriptionAndMenuNameAndRediscover
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndRediscover(offset int, limit int, Description_ string, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and rediscover = ?", Description_, MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndWeight Get MenuLinkContentDatas via DescriptionAndMenuNameAndWeight
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndWeight(offset int, limit int, Description_ string, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and weight = ?", Description_, MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndExpanded Get MenuLinkContentDatas via DescriptionAndMenuNameAndExpanded
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndExpanded(offset int, limit int, Description_ string, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and expanded = ?", Description_, MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndEnabled Get MenuLinkContentDatas via DescriptionAndMenuNameAndEnabled
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndEnabled(offset int, limit int, Description_ string, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and enabled = ?", Description_, MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndParent Get MenuLinkContentDatas via DescriptionAndMenuNameAndParent
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndParent(offset int, limit int, Description_ string, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and parent = ?", Description_, MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndChanged Get MenuLinkContentDatas via DescriptionAndMenuNameAndChanged
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndChanged(offset int, limit int, Description_ string, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and changed = ?", Description_, MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuNameAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndMenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndMenuNameAndDefaultLangcode(offset int, limit int, Description_ string, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ? and default_langcode = ?", Description_, MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_title Get MenuLinkContentDatas via DescriptionAndLink_uriAndLink_title
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_title(offset int, limit int, Description_ string, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and link_title = ?", Description_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_options Get MenuLinkContentDatas via DescriptionAndLink_uriAndLink_options
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndLink_options(offset int, limit int, Description_ string, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and link_options = ?", Description_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndExternal Get MenuLinkContentDatas via DescriptionAndLink_uriAndExternal
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndExternal(offset int, limit int, Description_ string, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and external = ?", Description_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndRediscover Get MenuLinkContentDatas via DescriptionAndLink_uriAndRediscover
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndRediscover(offset int, limit int, Description_ string, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and rediscover = ?", Description_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndWeight Get MenuLinkContentDatas via DescriptionAndLink_uriAndWeight
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndWeight(offset int, limit int, Description_ string, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and weight = ?", Description_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndExpanded Get MenuLinkContentDatas via DescriptionAndLink_uriAndExpanded
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndExpanded(offset int, limit int, Description_ string, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and expanded = ?", Description_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndEnabled Get MenuLinkContentDatas via DescriptionAndLink_uriAndEnabled
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndEnabled(offset int, limit int, Description_ string, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and enabled = ?", Description_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndParent Get MenuLinkContentDatas via DescriptionAndLink_uriAndParent
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndParent(offset int, limit int, Description_ string, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and parent = ?", Description_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndChanged Get MenuLinkContentDatas via DescriptionAndLink_uriAndChanged
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndChanged(offset int, limit int, Description_ string, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and changed = ?", Description_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndLink_uriAndDefaultLangcode(offset int, limit int, Description_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ? and default_langcode = ?", Description_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndLink_options Get MenuLinkContentDatas via DescriptionAndLink_titleAndLink_options
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndLink_options(offset int, limit int, Description_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and link_options = ?", Description_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndExternal Get MenuLinkContentDatas via DescriptionAndLink_titleAndExternal
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndExternal(offset int, limit int, Description_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and external = ?", Description_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndRediscover Get MenuLinkContentDatas via DescriptionAndLink_titleAndRediscover
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndRediscover(offset int, limit int, Description_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and rediscover = ?", Description_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndWeight Get MenuLinkContentDatas via DescriptionAndLink_titleAndWeight
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndWeight(offset int, limit int, Description_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and weight = ?", Description_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndExpanded Get MenuLinkContentDatas via DescriptionAndLink_titleAndExpanded
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndExpanded(offset int, limit int, Description_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and expanded = ?", Description_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndEnabled Get MenuLinkContentDatas via DescriptionAndLink_titleAndEnabled
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndEnabled(offset int, limit int, Description_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and enabled = ?", Description_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndParent Get MenuLinkContentDatas via DescriptionAndLink_titleAndParent
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndParent(offset int, limit int, Description_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and parent = ?", Description_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndChanged Get MenuLinkContentDatas via DescriptionAndLink_titleAndChanged
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndChanged(offset int, limit int, Description_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and changed = ?", Description_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndLink_titleAndDefaultLangcode(offset int, limit int, Description_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ? and default_langcode = ?", Description_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExternal Get MenuLinkContentDatas via DescriptionAndLink_optionsAndExternal
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExternal(offset int, limit int, Description_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and external = ?", Description_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndRediscover Get MenuLinkContentDatas via DescriptionAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndRediscover(offset int, limit int, Description_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and rediscover = ?", Description_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndWeight Get MenuLinkContentDatas via DescriptionAndLink_optionsAndWeight
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndWeight(offset int, limit int, Description_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and weight = ?", Description_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExpanded Get MenuLinkContentDatas via DescriptionAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndExpanded(offset int, limit int, Description_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and expanded = ?", Description_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndEnabled Get MenuLinkContentDatas via DescriptionAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndEnabled(offset int, limit int, Description_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and enabled = ?", Description_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndParent Get MenuLinkContentDatas via DescriptionAndLink_optionsAndParent
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndParent(offset int, limit int, Description_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and parent = ?", Description_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndChanged Get MenuLinkContentDatas via DescriptionAndLink_optionsAndChanged
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndChanged(offset int, limit int, Description_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and changed = ?", Description_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndLink_optionsAndDefaultLangcode(offset int, limit int, Description_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ? and default_langcode = ?", Description_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndRediscover Get MenuLinkContentDatas via DescriptionAndExternalAndRediscover
func GetMenuLinkContentDatasByDescriptionAndExternalAndRediscover(offset int, limit int, Description_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and rediscover = ?", Description_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndWeight Get MenuLinkContentDatas via DescriptionAndExternalAndWeight
func GetMenuLinkContentDatasByDescriptionAndExternalAndWeight(offset int, limit int, Description_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and weight = ?", Description_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndExpanded Get MenuLinkContentDatas via DescriptionAndExternalAndExpanded
func GetMenuLinkContentDatasByDescriptionAndExternalAndExpanded(offset int, limit int, Description_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and expanded = ?", Description_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndEnabled Get MenuLinkContentDatas via DescriptionAndExternalAndEnabled
func GetMenuLinkContentDatasByDescriptionAndExternalAndEnabled(offset int, limit int, Description_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and enabled = ?", Description_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndParent Get MenuLinkContentDatas via DescriptionAndExternalAndParent
func GetMenuLinkContentDatasByDescriptionAndExternalAndParent(offset int, limit int, Description_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and parent = ?", Description_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndChanged Get MenuLinkContentDatas via DescriptionAndExternalAndChanged
func GetMenuLinkContentDatasByDescriptionAndExternalAndChanged(offset int, limit int, Description_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and changed = ?", Description_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternalAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndExternalAndDefaultLangcode(offset int, limit int, Description_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ? and default_langcode = ?", Description_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndWeight Get MenuLinkContentDatas via DescriptionAndRediscoverAndWeight
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndWeight(offset int, limit int, Description_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and weight = ?", Description_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndExpanded Get MenuLinkContentDatas via DescriptionAndRediscoverAndExpanded
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndExpanded(offset int, limit int, Description_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and expanded = ?", Description_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndEnabled Get MenuLinkContentDatas via DescriptionAndRediscoverAndEnabled
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndEnabled(offset int, limit int, Description_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and enabled = ?", Description_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndParent Get MenuLinkContentDatas via DescriptionAndRediscoverAndParent
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndParent(offset int, limit int, Description_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and parent = ?", Description_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndChanged Get MenuLinkContentDatas via DescriptionAndRediscoverAndChanged
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndChanged(offset int, limit int, Description_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and changed = ?", Description_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndRediscoverAndDefaultLangcode(offset int, limit int, Description_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ? and default_langcode = ?", Description_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeightAndExpanded Get MenuLinkContentDatas via DescriptionAndWeightAndExpanded
func GetMenuLinkContentDatasByDescriptionAndWeightAndExpanded(offset int, limit int, Description_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ? and expanded = ?", Description_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeightAndEnabled Get MenuLinkContentDatas via DescriptionAndWeightAndEnabled
func GetMenuLinkContentDatasByDescriptionAndWeightAndEnabled(offset int, limit int, Description_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ? and enabled = ?", Description_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeightAndParent Get MenuLinkContentDatas via DescriptionAndWeightAndParent
func GetMenuLinkContentDatasByDescriptionAndWeightAndParent(offset int, limit int, Description_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ? and parent = ?", Description_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeightAndChanged Get MenuLinkContentDatas via DescriptionAndWeightAndChanged
func GetMenuLinkContentDatasByDescriptionAndWeightAndChanged(offset int, limit int, Description_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ? and changed = ?", Description_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeightAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndWeightAndDefaultLangcode(offset int, limit int, Description_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ? and default_langcode = ?", Description_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExpandedAndEnabled Get MenuLinkContentDatas via DescriptionAndExpandedAndEnabled
func GetMenuLinkContentDatasByDescriptionAndExpandedAndEnabled(offset int, limit int, Description_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and expanded = ? and enabled = ?", Description_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExpandedAndParent Get MenuLinkContentDatas via DescriptionAndExpandedAndParent
func GetMenuLinkContentDatasByDescriptionAndExpandedAndParent(offset int, limit int, Description_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and expanded = ? and parent = ?", Description_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExpandedAndChanged Get MenuLinkContentDatas via DescriptionAndExpandedAndChanged
func GetMenuLinkContentDatasByDescriptionAndExpandedAndChanged(offset int, limit int, Description_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and expanded = ? and changed = ?", Description_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndExpandedAndDefaultLangcode(offset int, limit int, Description_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and expanded = ? and default_langcode = ?", Description_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndEnabledAndParent Get MenuLinkContentDatas via DescriptionAndEnabledAndParent
func GetMenuLinkContentDatasByDescriptionAndEnabledAndParent(offset int, limit int, Description_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and enabled = ? and parent = ?", Description_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndEnabledAndChanged Get MenuLinkContentDatas via DescriptionAndEnabledAndChanged
func GetMenuLinkContentDatasByDescriptionAndEnabledAndChanged(offset int, limit int, Description_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and enabled = ? and changed = ?", Description_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndEnabledAndDefaultLangcode(offset int, limit int, Description_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and enabled = ? and default_langcode = ?", Description_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndParentAndChanged Get MenuLinkContentDatas via DescriptionAndParentAndChanged
func GetMenuLinkContentDatasByDescriptionAndParentAndChanged(offset int, limit int, Description_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and parent = ? and changed = ?", Description_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndParentAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndParentAndDefaultLangcode(offset int, limit int, Description_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and parent = ? and default_langcode = ?", Description_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndChangedAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndChangedAndDefaultLangcode(offset int, limit int, Description_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and changed = ? and default_langcode = ?", Description_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_title Get MenuLinkContentDatas via MenuNameAndLink_uriAndLink_title
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_title(offset int, limit int, MenuName_ string, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and link_title = ?", MenuName_, Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_options Get MenuLinkContentDatas via MenuNameAndLink_uriAndLink_options
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndLink_options(offset int, limit int, MenuName_ string, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and link_options = ?", MenuName_, Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndExternal Get MenuLinkContentDatas via MenuNameAndLink_uriAndExternal
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndExternal(offset int, limit int, MenuName_ string, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and external = ?", MenuName_, Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndRediscover Get MenuLinkContentDatas via MenuNameAndLink_uriAndRediscover
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndRediscover(offset int, limit int, MenuName_ string, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and rediscover = ?", MenuName_, Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndWeight Get MenuLinkContentDatas via MenuNameAndLink_uriAndWeight
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndWeight(offset int, limit int, MenuName_ string, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and weight = ?", MenuName_, Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndExpanded Get MenuLinkContentDatas via MenuNameAndLink_uriAndExpanded
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndExpanded(offset int, limit int, MenuName_ string, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and expanded = ?", MenuName_, Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndEnabled Get MenuLinkContentDatas via MenuNameAndLink_uriAndEnabled
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndEnabled(offset int, limit int, MenuName_ string, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and enabled = ?", MenuName_, Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndParent Get MenuLinkContentDatas via MenuNameAndLink_uriAndParent
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndParent(offset int, limit int, MenuName_ string, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and parent = ?", MenuName_, Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndChanged Get MenuLinkContentDatas via MenuNameAndLink_uriAndChanged
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndChanged(offset int, limit int, MenuName_ string, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and changed = ?", MenuName_, Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uriAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndLink_uriAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndLink_uriAndDefaultLangcode(offset int, limit int, MenuName_ string, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ? and default_langcode = ?", MenuName_, Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndLink_options Get MenuLinkContentDatas via MenuNameAndLink_titleAndLink_options
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndLink_options(offset int, limit int, MenuName_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and link_options = ?", MenuName_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndExternal Get MenuLinkContentDatas via MenuNameAndLink_titleAndExternal
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndExternal(offset int, limit int, MenuName_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and external = ?", MenuName_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndRediscover Get MenuLinkContentDatas via MenuNameAndLink_titleAndRediscover
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndRediscover(offset int, limit int, MenuName_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and rediscover = ?", MenuName_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndWeight Get MenuLinkContentDatas via MenuNameAndLink_titleAndWeight
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndWeight(offset int, limit int, MenuName_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and weight = ?", MenuName_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndExpanded Get MenuLinkContentDatas via MenuNameAndLink_titleAndExpanded
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndExpanded(offset int, limit int, MenuName_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and expanded = ?", MenuName_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndEnabled Get MenuLinkContentDatas via MenuNameAndLink_titleAndEnabled
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndEnabled(offset int, limit int, MenuName_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and enabled = ?", MenuName_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndParent Get MenuLinkContentDatas via MenuNameAndLink_titleAndParent
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndParent(offset int, limit int, MenuName_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and parent = ?", MenuName_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndChanged Get MenuLinkContentDatas via MenuNameAndLink_titleAndChanged
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndChanged(offset int, limit int, MenuName_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and changed = ?", MenuName_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndLink_titleAndDefaultLangcode(offset int, limit int, MenuName_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ? and default_langcode = ?", MenuName_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExternal Get MenuLinkContentDatas via MenuNameAndLink_optionsAndExternal
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExternal(offset int, limit int, MenuName_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and external = ?", MenuName_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndRediscover Get MenuLinkContentDatas via MenuNameAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndRediscover(offset int, limit int, MenuName_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and rediscover = ?", MenuName_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndWeight Get MenuLinkContentDatas via MenuNameAndLink_optionsAndWeight
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndWeight(offset int, limit int, MenuName_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and weight = ?", MenuName_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExpanded Get MenuLinkContentDatas via MenuNameAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndExpanded(offset int, limit int, MenuName_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and expanded = ?", MenuName_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndEnabled Get MenuLinkContentDatas via MenuNameAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndEnabled(offset int, limit int, MenuName_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and enabled = ?", MenuName_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndParent Get MenuLinkContentDatas via MenuNameAndLink_optionsAndParent
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndParent(offset int, limit int, MenuName_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and parent = ?", MenuName_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndChanged Get MenuLinkContentDatas via MenuNameAndLink_optionsAndChanged
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndChanged(offset int, limit int, MenuName_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and changed = ?", MenuName_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndLink_optionsAndDefaultLangcode(offset int, limit int, MenuName_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ? and default_langcode = ?", MenuName_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndRediscover Get MenuLinkContentDatas via MenuNameAndExternalAndRediscover
func GetMenuLinkContentDatasByMenuNameAndExternalAndRediscover(offset int, limit int, MenuName_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and rediscover = ?", MenuName_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndWeight Get MenuLinkContentDatas via MenuNameAndExternalAndWeight
func GetMenuLinkContentDatasByMenuNameAndExternalAndWeight(offset int, limit int, MenuName_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and weight = ?", MenuName_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndExpanded Get MenuLinkContentDatas via MenuNameAndExternalAndExpanded
func GetMenuLinkContentDatasByMenuNameAndExternalAndExpanded(offset int, limit int, MenuName_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and expanded = ?", MenuName_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndEnabled Get MenuLinkContentDatas via MenuNameAndExternalAndEnabled
func GetMenuLinkContentDatasByMenuNameAndExternalAndEnabled(offset int, limit int, MenuName_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and enabled = ?", MenuName_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndParent Get MenuLinkContentDatas via MenuNameAndExternalAndParent
func GetMenuLinkContentDatasByMenuNameAndExternalAndParent(offset int, limit int, MenuName_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and parent = ?", MenuName_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndChanged Get MenuLinkContentDatas via MenuNameAndExternalAndChanged
func GetMenuLinkContentDatasByMenuNameAndExternalAndChanged(offset int, limit int, MenuName_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and changed = ?", MenuName_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternalAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndExternalAndDefaultLangcode(offset int, limit int, MenuName_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ? and default_langcode = ?", MenuName_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndWeight Get MenuLinkContentDatas via MenuNameAndRediscoverAndWeight
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndWeight(offset int, limit int, MenuName_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and weight = ?", MenuName_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndExpanded Get MenuLinkContentDatas via MenuNameAndRediscoverAndExpanded
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndExpanded(offset int, limit int, MenuName_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and expanded = ?", MenuName_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndEnabled Get MenuLinkContentDatas via MenuNameAndRediscoverAndEnabled
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndEnabled(offset int, limit int, MenuName_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and enabled = ?", MenuName_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndParent Get MenuLinkContentDatas via MenuNameAndRediscoverAndParent
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndParent(offset int, limit int, MenuName_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and parent = ?", MenuName_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndChanged Get MenuLinkContentDatas via MenuNameAndRediscoverAndChanged
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndChanged(offset int, limit int, MenuName_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and changed = ?", MenuName_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndRediscoverAndDefaultLangcode(offset int, limit int, MenuName_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ? and default_langcode = ?", MenuName_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeightAndExpanded Get MenuLinkContentDatas via MenuNameAndWeightAndExpanded
func GetMenuLinkContentDatasByMenuNameAndWeightAndExpanded(offset int, limit int, MenuName_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ? and expanded = ?", MenuName_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeightAndEnabled Get MenuLinkContentDatas via MenuNameAndWeightAndEnabled
func GetMenuLinkContentDatasByMenuNameAndWeightAndEnabled(offset int, limit int, MenuName_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ? and enabled = ?", MenuName_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeightAndParent Get MenuLinkContentDatas via MenuNameAndWeightAndParent
func GetMenuLinkContentDatasByMenuNameAndWeightAndParent(offset int, limit int, MenuName_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ? and parent = ?", MenuName_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeightAndChanged Get MenuLinkContentDatas via MenuNameAndWeightAndChanged
func GetMenuLinkContentDatasByMenuNameAndWeightAndChanged(offset int, limit int, MenuName_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ? and changed = ?", MenuName_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeightAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndWeightAndDefaultLangcode(offset int, limit int, MenuName_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ? and default_langcode = ?", MenuName_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExpandedAndEnabled Get MenuLinkContentDatas via MenuNameAndExpandedAndEnabled
func GetMenuLinkContentDatasByMenuNameAndExpandedAndEnabled(offset int, limit int, MenuName_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and expanded = ? and enabled = ?", MenuName_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExpandedAndParent Get MenuLinkContentDatas via MenuNameAndExpandedAndParent
func GetMenuLinkContentDatasByMenuNameAndExpandedAndParent(offset int, limit int, MenuName_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and expanded = ? and parent = ?", MenuName_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExpandedAndChanged Get MenuLinkContentDatas via MenuNameAndExpandedAndChanged
func GetMenuLinkContentDatasByMenuNameAndExpandedAndChanged(offset int, limit int, MenuName_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and expanded = ? and changed = ?", MenuName_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndExpandedAndDefaultLangcode(offset int, limit int, MenuName_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and expanded = ? and default_langcode = ?", MenuName_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndEnabledAndParent Get MenuLinkContentDatas via MenuNameAndEnabledAndParent
func GetMenuLinkContentDatasByMenuNameAndEnabledAndParent(offset int, limit int, MenuName_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and enabled = ? and parent = ?", MenuName_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndEnabledAndChanged Get MenuLinkContentDatas via MenuNameAndEnabledAndChanged
func GetMenuLinkContentDatasByMenuNameAndEnabledAndChanged(offset int, limit int, MenuName_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and enabled = ? and changed = ?", MenuName_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndEnabledAndDefaultLangcode(offset int, limit int, MenuName_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and enabled = ? and default_langcode = ?", MenuName_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndParentAndChanged Get MenuLinkContentDatas via MenuNameAndParentAndChanged
func GetMenuLinkContentDatasByMenuNameAndParentAndChanged(offset int, limit int, MenuName_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and parent = ? and changed = ?", MenuName_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndParentAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndParentAndDefaultLangcode(offset int, limit int, MenuName_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and parent = ? and default_langcode = ?", MenuName_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndChangedAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndChangedAndDefaultLangcode(offset int, limit int, MenuName_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and changed = ? and default_langcode = ?", MenuName_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndLink_options Get MenuLinkContentDatas via Link_uriAndLink_titleAndLink_options
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndLink_options(offset int, limit int, Link_uri_ string, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and link_options = ?", Link_uri_, Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndExternal Get MenuLinkContentDatas via Link_uriAndLink_titleAndExternal
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndExternal(offset int, limit int, Link_uri_ string, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and external = ?", Link_uri_, Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndRediscover Get MenuLinkContentDatas via Link_uriAndLink_titleAndRediscover
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndRediscover(offset int, limit int, Link_uri_ string, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and rediscover = ?", Link_uri_, Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndWeight Get MenuLinkContentDatas via Link_uriAndLink_titleAndWeight
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndWeight(offset int, limit int, Link_uri_ string, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and weight = ?", Link_uri_, Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndExpanded Get MenuLinkContentDatas via Link_uriAndLink_titleAndExpanded
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndExpanded(offset int, limit int, Link_uri_ string, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and expanded = ?", Link_uri_, Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndEnabled Get MenuLinkContentDatas via Link_uriAndLink_titleAndEnabled
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndEnabled(offset int, limit int, Link_uri_ string, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and enabled = ?", Link_uri_, Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndParent Get MenuLinkContentDatas via Link_uriAndLink_titleAndParent
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndParent(offset int, limit int, Link_uri_ string, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and parent = ?", Link_uri_, Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndChanged Get MenuLinkContentDatas via Link_uriAndLink_titleAndChanged
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndChanged(offset int, limit int, Link_uri_ string, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and changed = ?", Link_uri_, Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_titleAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndLink_titleAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndLink_titleAndDefaultLangcode(offset int, limit int, Link_uri_ string, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ? and default_langcode = ?", Link_uri_, Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExternal Get MenuLinkContentDatas via Link_uriAndLink_optionsAndExternal
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExternal(offset int, limit int, Link_uri_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and external = ?", Link_uri_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndRediscover Get MenuLinkContentDatas via Link_uriAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndRediscover(offset int, limit int, Link_uri_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and rediscover = ?", Link_uri_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndWeight Get MenuLinkContentDatas via Link_uriAndLink_optionsAndWeight
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndWeight(offset int, limit int, Link_uri_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and weight = ?", Link_uri_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExpanded Get MenuLinkContentDatas via Link_uriAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndExpanded(offset int, limit int, Link_uri_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and expanded = ?", Link_uri_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndEnabled Get MenuLinkContentDatas via Link_uriAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndEnabled(offset int, limit int, Link_uri_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and enabled = ?", Link_uri_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndParent Get MenuLinkContentDatas via Link_uriAndLink_optionsAndParent
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndParent(offset int, limit int, Link_uri_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and parent = ?", Link_uri_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndChanged Get MenuLinkContentDatas via Link_uriAndLink_optionsAndChanged
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndChanged(offset int, limit int, Link_uri_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and changed = ?", Link_uri_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndLink_optionsAndDefaultLangcode(offset int, limit int, Link_uri_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ? and default_langcode = ?", Link_uri_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndRediscover Get MenuLinkContentDatas via Link_uriAndExternalAndRediscover
func GetMenuLinkContentDatasByLink_uriAndExternalAndRediscover(offset int, limit int, Link_uri_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and rediscover = ?", Link_uri_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndWeight Get MenuLinkContentDatas via Link_uriAndExternalAndWeight
func GetMenuLinkContentDatasByLink_uriAndExternalAndWeight(offset int, limit int, Link_uri_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and weight = ?", Link_uri_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndExpanded Get MenuLinkContentDatas via Link_uriAndExternalAndExpanded
func GetMenuLinkContentDatasByLink_uriAndExternalAndExpanded(offset int, limit int, Link_uri_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and expanded = ?", Link_uri_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndEnabled Get MenuLinkContentDatas via Link_uriAndExternalAndEnabled
func GetMenuLinkContentDatasByLink_uriAndExternalAndEnabled(offset int, limit int, Link_uri_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and enabled = ?", Link_uri_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndParent Get MenuLinkContentDatas via Link_uriAndExternalAndParent
func GetMenuLinkContentDatasByLink_uriAndExternalAndParent(offset int, limit int, Link_uri_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and parent = ?", Link_uri_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndChanged Get MenuLinkContentDatas via Link_uriAndExternalAndChanged
func GetMenuLinkContentDatasByLink_uriAndExternalAndChanged(offset int, limit int, Link_uri_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and changed = ?", Link_uri_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternalAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndExternalAndDefaultLangcode(offset int, limit int, Link_uri_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ? and default_langcode = ?", Link_uri_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndWeight Get MenuLinkContentDatas via Link_uriAndRediscoverAndWeight
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndWeight(offset int, limit int, Link_uri_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and weight = ?", Link_uri_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndExpanded Get MenuLinkContentDatas via Link_uriAndRediscoverAndExpanded
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndExpanded(offset int, limit int, Link_uri_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and expanded = ?", Link_uri_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndEnabled Get MenuLinkContentDatas via Link_uriAndRediscoverAndEnabled
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndEnabled(offset int, limit int, Link_uri_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and enabled = ?", Link_uri_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndParent Get MenuLinkContentDatas via Link_uriAndRediscoverAndParent
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndParent(offset int, limit int, Link_uri_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and parent = ?", Link_uri_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndChanged Get MenuLinkContentDatas via Link_uriAndRediscoverAndChanged
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndChanged(offset int, limit int, Link_uri_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and changed = ?", Link_uri_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndRediscoverAndDefaultLangcode(offset int, limit int, Link_uri_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ? and default_langcode = ?", Link_uri_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeightAndExpanded Get MenuLinkContentDatas via Link_uriAndWeightAndExpanded
func GetMenuLinkContentDatasByLink_uriAndWeightAndExpanded(offset int, limit int, Link_uri_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ? and expanded = ?", Link_uri_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeightAndEnabled Get MenuLinkContentDatas via Link_uriAndWeightAndEnabled
func GetMenuLinkContentDatasByLink_uriAndWeightAndEnabled(offset int, limit int, Link_uri_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ? and enabled = ?", Link_uri_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeightAndParent Get MenuLinkContentDatas via Link_uriAndWeightAndParent
func GetMenuLinkContentDatasByLink_uriAndWeightAndParent(offset int, limit int, Link_uri_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ? and parent = ?", Link_uri_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeightAndChanged Get MenuLinkContentDatas via Link_uriAndWeightAndChanged
func GetMenuLinkContentDatasByLink_uriAndWeightAndChanged(offset int, limit int, Link_uri_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ? and changed = ?", Link_uri_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeightAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndWeightAndDefaultLangcode(offset int, limit int, Link_uri_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ? and default_langcode = ?", Link_uri_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExpandedAndEnabled Get MenuLinkContentDatas via Link_uriAndExpandedAndEnabled
func GetMenuLinkContentDatasByLink_uriAndExpandedAndEnabled(offset int, limit int, Link_uri_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and expanded = ? and enabled = ?", Link_uri_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExpandedAndParent Get MenuLinkContentDatas via Link_uriAndExpandedAndParent
func GetMenuLinkContentDatasByLink_uriAndExpandedAndParent(offset int, limit int, Link_uri_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and expanded = ? and parent = ?", Link_uri_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExpandedAndChanged Get MenuLinkContentDatas via Link_uriAndExpandedAndChanged
func GetMenuLinkContentDatasByLink_uriAndExpandedAndChanged(offset int, limit int, Link_uri_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and expanded = ? and changed = ?", Link_uri_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndExpandedAndDefaultLangcode(offset int, limit int, Link_uri_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and expanded = ? and default_langcode = ?", Link_uri_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndEnabledAndParent Get MenuLinkContentDatas via Link_uriAndEnabledAndParent
func GetMenuLinkContentDatasByLink_uriAndEnabledAndParent(offset int, limit int, Link_uri_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and enabled = ? and parent = ?", Link_uri_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndEnabledAndChanged Get MenuLinkContentDatas via Link_uriAndEnabledAndChanged
func GetMenuLinkContentDatasByLink_uriAndEnabledAndChanged(offset int, limit int, Link_uri_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and enabled = ? and changed = ?", Link_uri_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndEnabledAndDefaultLangcode(offset int, limit int, Link_uri_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and enabled = ? and default_langcode = ?", Link_uri_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndParentAndChanged Get MenuLinkContentDatas via Link_uriAndParentAndChanged
func GetMenuLinkContentDatasByLink_uriAndParentAndChanged(offset int, limit int, Link_uri_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and parent = ? and changed = ?", Link_uri_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndParentAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndParentAndDefaultLangcode(offset int, limit int, Link_uri_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and parent = ? and default_langcode = ?", Link_uri_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndChangedAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndChangedAndDefaultLangcode(offset int, limit int, Link_uri_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and changed = ? and default_langcode = ?", Link_uri_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExternal Get MenuLinkContentDatas via Link_titleAndLink_optionsAndExternal
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExternal(offset int, limit int, Link_title_ string, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and external = ?", Link_title_, Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndRediscover Get MenuLinkContentDatas via Link_titleAndLink_optionsAndRediscover
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndRediscover(offset int, limit int, Link_title_ string, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and rediscover = ?", Link_title_, Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndWeight Get MenuLinkContentDatas via Link_titleAndLink_optionsAndWeight
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndWeight(offset int, limit int, Link_title_ string, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and weight = ?", Link_title_, Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExpanded Get MenuLinkContentDatas via Link_titleAndLink_optionsAndExpanded
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndExpanded(offset int, limit int, Link_title_ string, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and expanded = ?", Link_title_, Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndEnabled Get MenuLinkContentDatas via Link_titleAndLink_optionsAndEnabled
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndEnabled(offset int, limit int, Link_title_ string, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and enabled = ?", Link_title_, Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndParent Get MenuLinkContentDatas via Link_titleAndLink_optionsAndParent
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndParent(offset int, limit int, Link_title_ string, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and parent = ?", Link_title_, Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndChanged Get MenuLinkContentDatas via Link_titleAndLink_optionsAndChanged
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndChanged(offset int, limit int, Link_title_ string, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and changed = ?", Link_title_, Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndLink_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndLink_optionsAndDefaultLangcode(offset int, limit int, Link_title_ string, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ? and default_langcode = ?", Link_title_, Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndRediscover Get MenuLinkContentDatas via Link_titleAndExternalAndRediscover
func GetMenuLinkContentDatasByLink_titleAndExternalAndRediscover(offset int, limit int, Link_title_ string, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and rediscover = ?", Link_title_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndWeight Get MenuLinkContentDatas via Link_titleAndExternalAndWeight
func GetMenuLinkContentDatasByLink_titleAndExternalAndWeight(offset int, limit int, Link_title_ string, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and weight = ?", Link_title_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndExpanded Get MenuLinkContentDatas via Link_titleAndExternalAndExpanded
func GetMenuLinkContentDatasByLink_titleAndExternalAndExpanded(offset int, limit int, Link_title_ string, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and expanded = ?", Link_title_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndEnabled Get MenuLinkContentDatas via Link_titleAndExternalAndEnabled
func GetMenuLinkContentDatasByLink_titleAndExternalAndEnabled(offset int, limit int, Link_title_ string, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and enabled = ?", Link_title_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndParent Get MenuLinkContentDatas via Link_titleAndExternalAndParent
func GetMenuLinkContentDatasByLink_titleAndExternalAndParent(offset int, limit int, Link_title_ string, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and parent = ?", Link_title_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndChanged Get MenuLinkContentDatas via Link_titleAndExternalAndChanged
func GetMenuLinkContentDatasByLink_titleAndExternalAndChanged(offset int, limit int, Link_title_ string, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and changed = ?", Link_title_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternalAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndExternalAndDefaultLangcode(offset int, limit int, Link_title_ string, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ? and default_langcode = ?", Link_title_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndWeight Get MenuLinkContentDatas via Link_titleAndRediscoverAndWeight
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndWeight(offset int, limit int, Link_title_ string, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and weight = ?", Link_title_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndExpanded Get MenuLinkContentDatas via Link_titleAndRediscoverAndExpanded
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndExpanded(offset int, limit int, Link_title_ string, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and expanded = ?", Link_title_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndEnabled Get MenuLinkContentDatas via Link_titleAndRediscoverAndEnabled
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndEnabled(offset int, limit int, Link_title_ string, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and enabled = ?", Link_title_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndParent Get MenuLinkContentDatas via Link_titleAndRediscoverAndParent
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndParent(offset int, limit int, Link_title_ string, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and parent = ?", Link_title_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndChanged Get MenuLinkContentDatas via Link_titleAndRediscoverAndChanged
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndChanged(offset int, limit int, Link_title_ string, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and changed = ?", Link_title_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndRediscoverAndDefaultLangcode(offset int, limit int, Link_title_ string, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ? and default_langcode = ?", Link_title_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeightAndExpanded Get MenuLinkContentDatas via Link_titleAndWeightAndExpanded
func GetMenuLinkContentDatasByLink_titleAndWeightAndExpanded(offset int, limit int, Link_title_ string, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ? and expanded = ?", Link_title_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeightAndEnabled Get MenuLinkContentDatas via Link_titleAndWeightAndEnabled
func GetMenuLinkContentDatasByLink_titleAndWeightAndEnabled(offset int, limit int, Link_title_ string, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ? and enabled = ?", Link_title_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeightAndParent Get MenuLinkContentDatas via Link_titleAndWeightAndParent
func GetMenuLinkContentDatasByLink_titleAndWeightAndParent(offset int, limit int, Link_title_ string, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ? and parent = ?", Link_title_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeightAndChanged Get MenuLinkContentDatas via Link_titleAndWeightAndChanged
func GetMenuLinkContentDatasByLink_titleAndWeightAndChanged(offset int, limit int, Link_title_ string, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ? and changed = ?", Link_title_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeightAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndWeightAndDefaultLangcode(offset int, limit int, Link_title_ string, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ? and default_langcode = ?", Link_title_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExpandedAndEnabled Get MenuLinkContentDatas via Link_titleAndExpandedAndEnabled
func GetMenuLinkContentDatasByLink_titleAndExpandedAndEnabled(offset int, limit int, Link_title_ string, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and expanded = ? and enabled = ?", Link_title_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExpandedAndParent Get MenuLinkContentDatas via Link_titleAndExpandedAndParent
func GetMenuLinkContentDatasByLink_titleAndExpandedAndParent(offset int, limit int, Link_title_ string, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and expanded = ? and parent = ?", Link_title_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExpandedAndChanged Get MenuLinkContentDatas via Link_titleAndExpandedAndChanged
func GetMenuLinkContentDatasByLink_titleAndExpandedAndChanged(offset int, limit int, Link_title_ string, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and expanded = ? and changed = ?", Link_title_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndExpandedAndDefaultLangcode(offset int, limit int, Link_title_ string, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and expanded = ? and default_langcode = ?", Link_title_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndEnabledAndParent Get MenuLinkContentDatas via Link_titleAndEnabledAndParent
func GetMenuLinkContentDatasByLink_titleAndEnabledAndParent(offset int, limit int, Link_title_ string, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and enabled = ? and parent = ?", Link_title_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndEnabledAndChanged Get MenuLinkContentDatas via Link_titleAndEnabledAndChanged
func GetMenuLinkContentDatasByLink_titleAndEnabledAndChanged(offset int, limit int, Link_title_ string, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and enabled = ? and changed = ?", Link_title_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndEnabledAndDefaultLangcode(offset int, limit int, Link_title_ string, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and enabled = ? and default_langcode = ?", Link_title_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndParentAndChanged Get MenuLinkContentDatas via Link_titleAndParentAndChanged
func GetMenuLinkContentDatasByLink_titleAndParentAndChanged(offset int, limit int, Link_title_ string, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and parent = ? and changed = ?", Link_title_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndParentAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndParentAndDefaultLangcode(offset int, limit int, Link_title_ string, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and parent = ? and default_langcode = ?", Link_title_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndChangedAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndChangedAndDefaultLangcode(offset int, limit int, Link_title_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and changed = ? and default_langcode = ?", Link_title_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndRediscover Get MenuLinkContentDatas via Link_optionsAndExternalAndRediscover
func GetMenuLinkContentDatasByLink_optionsAndExternalAndRediscover(offset int, limit int, Link_options_ []byte, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and rediscover = ?", Link_options_, External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndWeight Get MenuLinkContentDatas via Link_optionsAndExternalAndWeight
func GetMenuLinkContentDatasByLink_optionsAndExternalAndWeight(offset int, limit int, Link_options_ []byte, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and weight = ?", Link_options_, External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndExpanded Get MenuLinkContentDatas via Link_optionsAndExternalAndExpanded
func GetMenuLinkContentDatasByLink_optionsAndExternalAndExpanded(offset int, limit int, Link_options_ []byte, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and expanded = ?", Link_options_, External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndEnabled Get MenuLinkContentDatas via Link_optionsAndExternalAndEnabled
func GetMenuLinkContentDatasByLink_optionsAndExternalAndEnabled(offset int, limit int, Link_options_ []byte, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and enabled = ?", Link_options_, External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndParent Get MenuLinkContentDatas via Link_optionsAndExternalAndParent
func GetMenuLinkContentDatasByLink_optionsAndExternalAndParent(offset int, limit int, Link_options_ []byte, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and parent = ?", Link_options_, External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndChanged Get MenuLinkContentDatas via Link_optionsAndExternalAndChanged
func GetMenuLinkContentDatasByLink_optionsAndExternalAndChanged(offset int, limit int, Link_options_ []byte, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and changed = ?", Link_options_, External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternalAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndExternalAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndExternalAndDefaultLangcode(offset int, limit int, Link_options_ []byte, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ? and default_langcode = ?", Link_options_, External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndWeight Get MenuLinkContentDatas via Link_optionsAndRediscoverAndWeight
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndWeight(offset int, limit int, Link_options_ []byte, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and weight = ?", Link_options_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndExpanded Get MenuLinkContentDatas via Link_optionsAndRediscoverAndExpanded
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndExpanded(offset int, limit int, Link_options_ []byte, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and expanded = ?", Link_options_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndEnabled Get MenuLinkContentDatas via Link_optionsAndRediscoverAndEnabled
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndEnabled(offset int, limit int, Link_options_ []byte, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and enabled = ?", Link_options_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndParent Get MenuLinkContentDatas via Link_optionsAndRediscoverAndParent
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndParent(offset int, limit int, Link_options_ []byte, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and parent = ?", Link_options_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndChanged Get MenuLinkContentDatas via Link_optionsAndRediscoverAndChanged
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndChanged(offset int, limit int, Link_options_ []byte, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and changed = ?", Link_options_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndRediscoverAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ? and default_langcode = ?", Link_options_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeightAndExpanded Get MenuLinkContentDatas via Link_optionsAndWeightAndExpanded
func GetMenuLinkContentDatasByLink_optionsAndWeightAndExpanded(offset int, limit int, Link_options_ []byte, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ? and expanded = ?", Link_options_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeightAndEnabled Get MenuLinkContentDatas via Link_optionsAndWeightAndEnabled
func GetMenuLinkContentDatasByLink_optionsAndWeightAndEnabled(offset int, limit int, Link_options_ []byte, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ? and enabled = ?", Link_options_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeightAndParent Get MenuLinkContentDatas via Link_optionsAndWeightAndParent
func GetMenuLinkContentDatasByLink_optionsAndWeightAndParent(offset int, limit int, Link_options_ []byte, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ? and parent = ?", Link_options_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeightAndChanged Get MenuLinkContentDatas via Link_optionsAndWeightAndChanged
func GetMenuLinkContentDatasByLink_optionsAndWeightAndChanged(offset int, limit int, Link_options_ []byte, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ? and changed = ?", Link_options_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeightAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndWeightAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ? and default_langcode = ?", Link_options_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExpandedAndEnabled Get MenuLinkContentDatas via Link_optionsAndExpandedAndEnabled
func GetMenuLinkContentDatasByLink_optionsAndExpandedAndEnabled(offset int, limit int, Link_options_ []byte, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and expanded = ? and enabled = ?", Link_options_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExpandedAndParent Get MenuLinkContentDatas via Link_optionsAndExpandedAndParent
func GetMenuLinkContentDatasByLink_optionsAndExpandedAndParent(offset int, limit int, Link_options_ []byte, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and expanded = ? and parent = ?", Link_options_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExpandedAndChanged Get MenuLinkContentDatas via Link_optionsAndExpandedAndChanged
func GetMenuLinkContentDatasByLink_optionsAndExpandedAndChanged(offset int, limit int, Link_options_ []byte, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and expanded = ? and changed = ?", Link_options_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndExpandedAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and expanded = ? and default_langcode = ?", Link_options_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndEnabledAndParent Get MenuLinkContentDatas via Link_optionsAndEnabledAndParent
func GetMenuLinkContentDatasByLink_optionsAndEnabledAndParent(offset int, limit int, Link_options_ []byte, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and enabled = ? and parent = ?", Link_options_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndEnabledAndChanged Get MenuLinkContentDatas via Link_optionsAndEnabledAndChanged
func GetMenuLinkContentDatasByLink_optionsAndEnabledAndChanged(offset int, limit int, Link_options_ []byte, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and enabled = ? and changed = ?", Link_options_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndEnabledAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and enabled = ? and default_langcode = ?", Link_options_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndParentAndChanged Get MenuLinkContentDatas via Link_optionsAndParentAndChanged
func GetMenuLinkContentDatasByLink_optionsAndParentAndChanged(offset int, limit int, Link_options_ []byte, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and parent = ? and changed = ?", Link_options_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndParentAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndParentAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and parent = ? and default_langcode = ?", Link_options_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndChangedAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndChangedAndDefaultLangcode(offset int, limit int, Link_options_ []byte, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and changed = ? and default_langcode = ?", Link_options_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndWeight Get MenuLinkContentDatas via ExternalAndRediscoverAndWeight
func GetMenuLinkContentDatasByExternalAndRediscoverAndWeight(offset int, limit int, External_ int, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and weight = ?", External_, Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndExpanded Get MenuLinkContentDatas via ExternalAndRediscoverAndExpanded
func GetMenuLinkContentDatasByExternalAndRediscoverAndExpanded(offset int, limit int, External_ int, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and expanded = ?", External_, Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndEnabled Get MenuLinkContentDatas via ExternalAndRediscoverAndEnabled
func GetMenuLinkContentDatasByExternalAndRediscoverAndEnabled(offset int, limit int, External_ int, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and enabled = ?", External_, Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndParent Get MenuLinkContentDatas via ExternalAndRediscoverAndParent
func GetMenuLinkContentDatasByExternalAndRediscoverAndParent(offset int, limit int, External_ int, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and parent = ?", External_, Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndChanged Get MenuLinkContentDatas via ExternalAndRediscoverAndChanged
func GetMenuLinkContentDatasByExternalAndRediscoverAndChanged(offset int, limit int, External_ int, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and changed = ?", External_, Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscoverAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndRediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndRediscoverAndDefaultLangcode(offset int, limit int, External_ int, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ? and default_langcode = ?", External_, Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeightAndExpanded Get MenuLinkContentDatas via ExternalAndWeightAndExpanded
func GetMenuLinkContentDatasByExternalAndWeightAndExpanded(offset int, limit int, External_ int, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ? and expanded = ?", External_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeightAndEnabled Get MenuLinkContentDatas via ExternalAndWeightAndEnabled
func GetMenuLinkContentDatasByExternalAndWeightAndEnabled(offset int, limit int, External_ int, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ? and enabled = ?", External_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeightAndParent Get MenuLinkContentDatas via ExternalAndWeightAndParent
func GetMenuLinkContentDatasByExternalAndWeightAndParent(offset int, limit int, External_ int, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ? and parent = ?", External_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeightAndChanged Get MenuLinkContentDatas via ExternalAndWeightAndChanged
func GetMenuLinkContentDatasByExternalAndWeightAndChanged(offset int, limit int, External_ int, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ? and changed = ?", External_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeightAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndWeightAndDefaultLangcode(offset int, limit int, External_ int, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ? and default_langcode = ?", External_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndExpandedAndEnabled Get MenuLinkContentDatas via ExternalAndExpandedAndEnabled
func GetMenuLinkContentDatasByExternalAndExpandedAndEnabled(offset int, limit int, External_ int, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and expanded = ? and enabled = ?", External_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndExpandedAndParent Get MenuLinkContentDatas via ExternalAndExpandedAndParent
func GetMenuLinkContentDatasByExternalAndExpandedAndParent(offset int, limit int, External_ int, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and expanded = ? and parent = ?", External_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndExpandedAndChanged Get MenuLinkContentDatas via ExternalAndExpandedAndChanged
func GetMenuLinkContentDatasByExternalAndExpandedAndChanged(offset int, limit int, External_ int, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and expanded = ? and changed = ?", External_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndExpandedAndDefaultLangcode(offset int, limit int, External_ int, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and expanded = ? and default_langcode = ?", External_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndEnabledAndParent Get MenuLinkContentDatas via ExternalAndEnabledAndParent
func GetMenuLinkContentDatasByExternalAndEnabledAndParent(offset int, limit int, External_ int, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and enabled = ? and parent = ?", External_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndEnabledAndChanged Get MenuLinkContentDatas via ExternalAndEnabledAndChanged
func GetMenuLinkContentDatasByExternalAndEnabledAndChanged(offset int, limit int, External_ int, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and enabled = ? and changed = ?", External_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndEnabledAndDefaultLangcode(offset int, limit int, External_ int, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and enabled = ? and default_langcode = ?", External_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndParentAndChanged Get MenuLinkContentDatas via ExternalAndParentAndChanged
func GetMenuLinkContentDatasByExternalAndParentAndChanged(offset int, limit int, External_ int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and parent = ? and changed = ?", External_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndParentAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndParentAndDefaultLangcode(offset int, limit int, External_ int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and parent = ? and default_langcode = ?", External_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndChangedAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndChangedAndDefaultLangcode(offset int, limit int, External_ int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and changed = ? and default_langcode = ?", External_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeightAndExpanded Get MenuLinkContentDatas via RediscoverAndWeightAndExpanded
func GetMenuLinkContentDatasByRediscoverAndWeightAndExpanded(offset int, limit int, Rediscover_ int, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ? and expanded = ?", Rediscover_, Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeightAndEnabled Get MenuLinkContentDatas via RediscoverAndWeightAndEnabled
func GetMenuLinkContentDatasByRediscoverAndWeightAndEnabled(offset int, limit int, Rediscover_ int, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ? and enabled = ?", Rediscover_, Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeightAndParent Get MenuLinkContentDatas via RediscoverAndWeightAndParent
func GetMenuLinkContentDatasByRediscoverAndWeightAndParent(offset int, limit int, Rediscover_ int, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ? and parent = ?", Rediscover_, Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeightAndChanged Get MenuLinkContentDatas via RediscoverAndWeightAndChanged
func GetMenuLinkContentDatasByRediscoverAndWeightAndChanged(offset int, limit int, Rediscover_ int, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ? and changed = ?", Rediscover_, Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeightAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndWeightAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndWeightAndDefaultLangcode(offset int, limit int, Rediscover_ int, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ? and default_langcode = ?", Rediscover_, Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndExpandedAndEnabled Get MenuLinkContentDatas via RediscoverAndExpandedAndEnabled
func GetMenuLinkContentDatasByRediscoverAndExpandedAndEnabled(offset int, limit int, Rediscover_ int, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and expanded = ? and enabled = ?", Rediscover_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndExpandedAndParent Get MenuLinkContentDatas via RediscoverAndExpandedAndParent
func GetMenuLinkContentDatasByRediscoverAndExpandedAndParent(offset int, limit int, Rediscover_ int, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and expanded = ? and parent = ?", Rediscover_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndExpandedAndChanged Get MenuLinkContentDatas via RediscoverAndExpandedAndChanged
func GetMenuLinkContentDatasByRediscoverAndExpandedAndChanged(offset int, limit int, Rediscover_ int, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and expanded = ? and changed = ?", Rediscover_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndExpandedAndDefaultLangcode(offset int, limit int, Rediscover_ int, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and expanded = ? and default_langcode = ?", Rediscover_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndEnabledAndParent Get MenuLinkContentDatas via RediscoverAndEnabledAndParent
func GetMenuLinkContentDatasByRediscoverAndEnabledAndParent(offset int, limit int, Rediscover_ int, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and enabled = ? and parent = ?", Rediscover_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndEnabledAndChanged Get MenuLinkContentDatas via RediscoverAndEnabledAndChanged
func GetMenuLinkContentDatasByRediscoverAndEnabledAndChanged(offset int, limit int, Rediscover_ int, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and enabled = ? and changed = ?", Rediscover_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndEnabledAndDefaultLangcode(offset int, limit int, Rediscover_ int, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and enabled = ? and default_langcode = ?", Rediscover_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndParentAndChanged Get MenuLinkContentDatas via RediscoverAndParentAndChanged
func GetMenuLinkContentDatasByRediscoverAndParentAndChanged(offset int, limit int, Rediscover_ int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and parent = ? and changed = ?", Rediscover_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndParentAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndParentAndDefaultLangcode(offset int, limit int, Rediscover_ int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and parent = ? and default_langcode = ?", Rediscover_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndChangedAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndChangedAndDefaultLangcode(offset int, limit int, Rediscover_ int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and changed = ? and default_langcode = ?", Rediscover_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndExpandedAndEnabled Get MenuLinkContentDatas via WeightAndExpandedAndEnabled
func GetMenuLinkContentDatasByWeightAndExpandedAndEnabled(offset int, limit int, Weight_ int, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and expanded = ? and enabled = ?", Weight_, Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndExpandedAndParent Get MenuLinkContentDatas via WeightAndExpandedAndParent
func GetMenuLinkContentDatasByWeightAndExpandedAndParent(offset int, limit int, Weight_ int, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and expanded = ? and parent = ?", Weight_, Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndExpandedAndChanged Get MenuLinkContentDatas via WeightAndExpandedAndChanged
func GetMenuLinkContentDatasByWeightAndExpandedAndChanged(offset int, limit int, Weight_ int, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and expanded = ? and changed = ?", Weight_, Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndExpandedAndDefaultLangcode Get MenuLinkContentDatas via WeightAndExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByWeightAndExpandedAndDefaultLangcode(offset int, limit int, Weight_ int, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and expanded = ? and default_langcode = ?", Weight_, Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndEnabledAndParent Get MenuLinkContentDatas via WeightAndEnabledAndParent
func GetMenuLinkContentDatasByWeightAndEnabledAndParent(offset int, limit int, Weight_ int, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and enabled = ? and parent = ?", Weight_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndEnabledAndChanged Get MenuLinkContentDatas via WeightAndEnabledAndChanged
func GetMenuLinkContentDatasByWeightAndEnabledAndChanged(offset int, limit int, Weight_ int, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and enabled = ? and changed = ?", Weight_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via WeightAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByWeightAndEnabledAndDefaultLangcode(offset int, limit int, Weight_ int, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and enabled = ? and default_langcode = ?", Weight_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndParentAndChanged Get MenuLinkContentDatas via WeightAndParentAndChanged
func GetMenuLinkContentDatasByWeightAndParentAndChanged(offset int, limit int, Weight_ int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and parent = ? and changed = ?", Weight_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndParentAndDefaultLangcode Get MenuLinkContentDatas via WeightAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByWeightAndParentAndDefaultLangcode(offset int, limit int, Weight_ int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and parent = ? and default_langcode = ?", Weight_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndChangedAndDefaultLangcode Get MenuLinkContentDatas via WeightAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByWeightAndChangedAndDefaultLangcode(offset int, limit int, Weight_ int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and changed = ? and default_langcode = ?", Weight_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndEnabledAndParent Get MenuLinkContentDatas via ExpandedAndEnabledAndParent
func GetMenuLinkContentDatasByExpandedAndEnabledAndParent(offset int, limit int, Expanded_ int, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and enabled = ? and parent = ?", Expanded_, Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndEnabledAndChanged Get MenuLinkContentDatas via ExpandedAndEnabledAndChanged
func GetMenuLinkContentDatasByExpandedAndEnabledAndChanged(offset int, limit int, Expanded_ int, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and enabled = ? and changed = ?", Expanded_, Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndEnabledAndDefaultLangcode Get MenuLinkContentDatas via ExpandedAndEnabledAndDefaultLangcode
func GetMenuLinkContentDatasByExpandedAndEnabledAndDefaultLangcode(offset int, limit int, Expanded_ int, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and enabled = ? and default_langcode = ?", Expanded_, Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndParentAndChanged Get MenuLinkContentDatas via ExpandedAndParentAndChanged
func GetMenuLinkContentDatasByExpandedAndParentAndChanged(offset int, limit int, Expanded_ int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and parent = ? and changed = ?", Expanded_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndParentAndDefaultLangcode Get MenuLinkContentDatas via ExpandedAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByExpandedAndParentAndDefaultLangcode(offset int, limit int, Expanded_ int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and parent = ? and default_langcode = ?", Expanded_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndChangedAndDefaultLangcode Get MenuLinkContentDatas via ExpandedAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByExpandedAndChangedAndDefaultLangcode(offset int, limit int, Expanded_ int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and changed = ? and default_langcode = ?", Expanded_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndParentAndChanged Get MenuLinkContentDatas via EnabledAndParentAndChanged
func GetMenuLinkContentDatasByEnabledAndParentAndChanged(offset int, limit int, Enabled_ int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and parent = ? and changed = ?", Enabled_, Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndParentAndDefaultLangcode Get MenuLinkContentDatas via EnabledAndParentAndDefaultLangcode
func GetMenuLinkContentDatasByEnabledAndParentAndDefaultLangcode(offset int, limit int, Enabled_ int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and parent = ? and default_langcode = ?", Enabled_, Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndChangedAndDefaultLangcode Get MenuLinkContentDatas via EnabledAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByEnabledAndChangedAndDefaultLangcode(offset int, limit int, Enabled_ int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and changed = ? and default_langcode = ?", Enabled_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByParentAndChangedAndDefaultLangcode Get MenuLinkContentDatas via ParentAndChangedAndDefaultLangcode
func GetMenuLinkContentDatasByParentAndChangedAndDefaultLangcode(offset int, limit int, Parent_ string, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("parent = ? and changed = ? and default_langcode = ?", Parent_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndBundle Get MenuLinkContentDatas via IdAndBundle
func GetMenuLinkContentDatasByIdAndBundle(offset int, limit int, Id_ int64, Bundle_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and bundle = ?", Id_, Bundle_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLangcode Get MenuLinkContentDatas via IdAndLangcode
func GetMenuLinkContentDatasByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndTitle Get MenuLinkContentDatas via IdAndTitle
func GetMenuLinkContentDatasByIdAndTitle(offset int, limit int, Id_ int64, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and title = ?", Id_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDescription Get MenuLinkContentDatas via IdAndDescription
func GetMenuLinkContentDatasByIdAndDescription(offset int, limit int, Id_ int64, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and description = ?", Id_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndMenuName Get MenuLinkContentDatas via IdAndMenuName
func GetMenuLinkContentDatasByIdAndMenuName(offset int, limit int, Id_ int64, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and menu_name = ?", Id_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_uri Get MenuLinkContentDatas via IdAndLink_uri
func GetMenuLinkContentDatasByIdAndLink_uri(offset int, limit int, Id_ int64, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_uri = ?", Id_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_title Get MenuLinkContentDatas via IdAndLink_title
func GetMenuLinkContentDatasByIdAndLink_title(offset int, limit int, Id_ int64, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_title = ?", Id_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndLink_options Get MenuLinkContentDatas via IdAndLink_options
func GetMenuLinkContentDatasByIdAndLink_options(offset int, limit int, Id_ int64, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and link_options = ?", Id_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExternal Get MenuLinkContentDatas via IdAndExternal
func GetMenuLinkContentDatasByIdAndExternal(offset int, limit int, Id_ int64, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and external = ?", Id_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndRediscover Get MenuLinkContentDatas via IdAndRediscover
func GetMenuLinkContentDatasByIdAndRediscover(offset int, limit int, Id_ int64, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and rediscover = ?", Id_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndWeight Get MenuLinkContentDatas via IdAndWeight
func GetMenuLinkContentDatasByIdAndWeight(offset int, limit int, Id_ int64, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and weight = ?", Id_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndExpanded Get MenuLinkContentDatas via IdAndExpanded
func GetMenuLinkContentDatasByIdAndExpanded(offset int, limit int, Id_ int64, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and expanded = ?", Id_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndEnabled Get MenuLinkContentDatas via IdAndEnabled
func GetMenuLinkContentDatasByIdAndEnabled(offset int, limit int, Id_ int64, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and enabled = ?", Id_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndParent Get MenuLinkContentDatas via IdAndParent
func GetMenuLinkContentDatasByIdAndParent(offset int, limit int, Id_ int64, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and parent = ?", Id_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndChanged Get MenuLinkContentDatas via IdAndChanged
func GetMenuLinkContentDatasByIdAndChanged(offset int, limit int, Id_ int64, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and changed = ?", Id_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByIdAndDefaultLangcode Get MenuLinkContentDatas via IdAndDefaultLangcode
func GetMenuLinkContentDatasByIdAndDefaultLangcode(offset int, limit int, Id_ int64, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ? and default_langcode = ?", Id_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLangcode Get MenuLinkContentDatas via BundleAndLangcode
func GetMenuLinkContentDatasByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndTitle Get MenuLinkContentDatas via BundleAndTitle
func GetMenuLinkContentDatasByBundleAndTitle(offset int, limit int, Bundle_ string, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and title = ?", Bundle_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDescription Get MenuLinkContentDatas via BundleAndDescription
func GetMenuLinkContentDatasByBundleAndDescription(offset int, limit int, Bundle_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and description = ?", Bundle_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndMenuName Get MenuLinkContentDatas via BundleAndMenuName
func GetMenuLinkContentDatasByBundleAndMenuName(offset int, limit int, Bundle_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and menu_name = ?", Bundle_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_uri Get MenuLinkContentDatas via BundleAndLink_uri
func GetMenuLinkContentDatasByBundleAndLink_uri(offset int, limit int, Bundle_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_uri = ?", Bundle_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_title Get MenuLinkContentDatas via BundleAndLink_title
func GetMenuLinkContentDatasByBundleAndLink_title(offset int, limit int, Bundle_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_title = ?", Bundle_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndLink_options Get MenuLinkContentDatas via BundleAndLink_options
func GetMenuLinkContentDatasByBundleAndLink_options(offset int, limit int, Bundle_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and link_options = ?", Bundle_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExternal Get MenuLinkContentDatas via BundleAndExternal
func GetMenuLinkContentDatasByBundleAndExternal(offset int, limit int, Bundle_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and external = ?", Bundle_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndRediscover Get MenuLinkContentDatas via BundleAndRediscover
func GetMenuLinkContentDatasByBundleAndRediscover(offset int, limit int, Bundle_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and rediscover = ?", Bundle_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndWeight Get MenuLinkContentDatas via BundleAndWeight
func GetMenuLinkContentDatasByBundleAndWeight(offset int, limit int, Bundle_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and weight = ?", Bundle_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndExpanded Get MenuLinkContentDatas via BundleAndExpanded
func GetMenuLinkContentDatasByBundleAndExpanded(offset int, limit int, Bundle_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and expanded = ?", Bundle_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndEnabled Get MenuLinkContentDatas via BundleAndEnabled
func GetMenuLinkContentDatasByBundleAndEnabled(offset int, limit int, Bundle_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and enabled = ?", Bundle_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndParent Get MenuLinkContentDatas via BundleAndParent
func GetMenuLinkContentDatasByBundleAndParent(offset int, limit int, Bundle_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and parent = ?", Bundle_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndChanged Get MenuLinkContentDatas via BundleAndChanged
func GetMenuLinkContentDatasByBundleAndChanged(offset int, limit int, Bundle_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and changed = ?", Bundle_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByBundleAndDefaultLangcode Get MenuLinkContentDatas via BundleAndDefaultLangcode
func GetMenuLinkContentDatasByBundleAndDefaultLangcode(offset int, limit int, Bundle_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ? and default_langcode = ?", Bundle_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndTitle Get MenuLinkContentDatas via LangcodeAndTitle
func GetMenuLinkContentDatasByLangcodeAndTitle(offset int, limit int, Langcode_ string, Title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and title = ?", Langcode_, Title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDescription Get MenuLinkContentDatas via LangcodeAndDescription
func GetMenuLinkContentDatasByLangcodeAndDescription(offset int, limit int, Langcode_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and description = ?", Langcode_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndMenuName Get MenuLinkContentDatas via LangcodeAndMenuName
func GetMenuLinkContentDatasByLangcodeAndMenuName(offset int, limit int, Langcode_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and menu_name = ?", Langcode_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_uri Get MenuLinkContentDatas via LangcodeAndLink_uri
func GetMenuLinkContentDatasByLangcodeAndLink_uri(offset int, limit int, Langcode_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_uri = ?", Langcode_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_title Get MenuLinkContentDatas via LangcodeAndLink_title
func GetMenuLinkContentDatasByLangcodeAndLink_title(offset int, limit int, Langcode_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_title = ?", Langcode_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndLink_options Get MenuLinkContentDatas via LangcodeAndLink_options
func GetMenuLinkContentDatasByLangcodeAndLink_options(offset int, limit int, Langcode_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and link_options = ?", Langcode_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExternal Get MenuLinkContentDatas via LangcodeAndExternal
func GetMenuLinkContentDatasByLangcodeAndExternal(offset int, limit int, Langcode_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and external = ?", Langcode_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndRediscover Get MenuLinkContentDatas via LangcodeAndRediscover
func GetMenuLinkContentDatasByLangcodeAndRediscover(offset int, limit int, Langcode_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and rediscover = ?", Langcode_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndWeight Get MenuLinkContentDatas via LangcodeAndWeight
func GetMenuLinkContentDatasByLangcodeAndWeight(offset int, limit int, Langcode_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and weight = ?", Langcode_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndExpanded Get MenuLinkContentDatas via LangcodeAndExpanded
func GetMenuLinkContentDatasByLangcodeAndExpanded(offset int, limit int, Langcode_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and expanded = ?", Langcode_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndEnabled Get MenuLinkContentDatas via LangcodeAndEnabled
func GetMenuLinkContentDatasByLangcodeAndEnabled(offset int, limit int, Langcode_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and enabled = ?", Langcode_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndParent Get MenuLinkContentDatas via LangcodeAndParent
func GetMenuLinkContentDatasByLangcodeAndParent(offset int, limit int, Langcode_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and parent = ?", Langcode_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndChanged Get MenuLinkContentDatas via LangcodeAndChanged
func GetMenuLinkContentDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLangcodeAndDefaultLangcode Get MenuLinkContentDatas via LangcodeAndDefaultLangcode
func GetMenuLinkContentDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDescription Get MenuLinkContentDatas via TitleAndDescription
func GetMenuLinkContentDatasByTitleAndDescription(offset int, limit int, Title_ string, Description_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and description = ?", Title_, Description_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndMenuName Get MenuLinkContentDatas via TitleAndMenuName
func GetMenuLinkContentDatasByTitleAndMenuName(offset int, limit int, Title_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and menu_name = ?", Title_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_uri Get MenuLinkContentDatas via TitleAndLink_uri
func GetMenuLinkContentDatasByTitleAndLink_uri(offset int, limit int, Title_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_uri = ?", Title_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_title Get MenuLinkContentDatas via TitleAndLink_title
func GetMenuLinkContentDatasByTitleAndLink_title(offset int, limit int, Title_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_title = ?", Title_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndLink_options Get MenuLinkContentDatas via TitleAndLink_options
func GetMenuLinkContentDatasByTitleAndLink_options(offset int, limit int, Title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and link_options = ?", Title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExternal Get MenuLinkContentDatas via TitleAndExternal
func GetMenuLinkContentDatasByTitleAndExternal(offset int, limit int, Title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and external = ?", Title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndRediscover Get MenuLinkContentDatas via TitleAndRediscover
func GetMenuLinkContentDatasByTitleAndRediscover(offset int, limit int, Title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and rediscover = ?", Title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndWeight Get MenuLinkContentDatas via TitleAndWeight
func GetMenuLinkContentDatasByTitleAndWeight(offset int, limit int, Title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and weight = ?", Title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndExpanded Get MenuLinkContentDatas via TitleAndExpanded
func GetMenuLinkContentDatasByTitleAndExpanded(offset int, limit int, Title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and expanded = ?", Title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndEnabled Get MenuLinkContentDatas via TitleAndEnabled
func GetMenuLinkContentDatasByTitleAndEnabled(offset int, limit int, Title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and enabled = ?", Title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndParent Get MenuLinkContentDatas via TitleAndParent
func GetMenuLinkContentDatasByTitleAndParent(offset int, limit int, Title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and parent = ?", Title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndChanged Get MenuLinkContentDatas via TitleAndChanged
func GetMenuLinkContentDatasByTitleAndChanged(offset int, limit int, Title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and changed = ?", Title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByTitleAndDefaultLangcode Get MenuLinkContentDatas via TitleAndDefaultLangcode
func GetMenuLinkContentDatasByTitleAndDefaultLangcode(offset int, limit int, Title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ? and default_langcode = ?", Title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndMenuName Get MenuLinkContentDatas via DescriptionAndMenuName
func GetMenuLinkContentDatasByDescriptionAndMenuName(offset int, limit int, Description_ string, MenuName_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and menu_name = ?", Description_, MenuName_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_uri Get MenuLinkContentDatas via DescriptionAndLink_uri
func GetMenuLinkContentDatasByDescriptionAndLink_uri(offset int, limit int, Description_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_uri = ?", Description_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_title Get MenuLinkContentDatas via DescriptionAndLink_title
func GetMenuLinkContentDatasByDescriptionAndLink_title(offset int, limit int, Description_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_title = ?", Description_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndLink_options Get MenuLinkContentDatas via DescriptionAndLink_options
func GetMenuLinkContentDatasByDescriptionAndLink_options(offset int, limit int, Description_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and link_options = ?", Description_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExternal Get MenuLinkContentDatas via DescriptionAndExternal
func GetMenuLinkContentDatasByDescriptionAndExternal(offset int, limit int, Description_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and external = ?", Description_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndRediscover Get MenuLinkContentDatas via DescriptionAndRediscover
func GetMenuLinkContentDatasByDescriptionAndRediscover(offset int, limit int, Description_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and rediscover = ?", Description_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndWeight Get MenuLinkContentDatas via DescriptionAndWeight
func GetMenuLinkContentDatasByDescriptionAndWeight(offset int, limit int, Description_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and weight = ?", Description_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndExpanded Get MenuLinkContentDatas via DescriptionAndExpanded
func GetMenuLinkContentDatasByDescriptionAndExpanded(offset int, limit int, Description_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and expanded = ?", Description_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndEnabled Get MenuLinkContentDatas via DescriptionAndEnabled
func GetMenuLinkContentDatasByDescriptionAndEnabled(offset int, limit int, Description_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and enabled = ?", Description_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndParent Get MenuLinkContentDatas via DescriptionAndParent
func GetMenuLinkContentDatasByDescriptionAndParent(offset int, limit int, Description_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and parent = ?", Description_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndChanged Get MenuLinkContentDatas via DescriptionAndChanged
func GetMenuLinkContentDatasByDescriptionAndChanged(offset int, limit int, Description_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and changed = ?", Description_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByDescriptionAndDefaultLangcode Get MenuLinkContentDatas via DescriptionAndDefaultLangcode
func GetMenuLinkContentDatasByDescriptionAndDefaultLangcode(offset int, limit int, Description_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ? and default_langcode = ?", Description_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_uri Get MenuLinkContentDatas via MenuNameAndLink_uri
func GetMenuLinkContentDatasByMenuNameAndLink_uri(offset int, limit int, MenuName_ string, Link_uri_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_uri = ?", MenuName_, Link_uri_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_title Get MenuLinkContentDatas via MenuNameAndLink_title
func GetMenuLinkContentDatasByMenuNameAndLink_title(offset int, limit int, MenuName_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_title = ?", MenuName_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndLink_options Get MenuLinkContentDatas via MenuNameAndLink_options
func GetMenuLinkContentDatasByMenuNameAndLink_options(offset int, limit int, MenuName_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and link_options = ?", MenuName_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExternal Get MenuLinkContentDatas via MenuNameAndExternal
func GetMenuLinkContentDatasByMenuNameAndExternal(offset int, limit int, MenuName_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and external = ?", MenuName_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndRediscover Get MenuLinkContentDatas via MenuNameAndRediscover
func GetMenuLinkContentDatasByMenuNameAndRediscover(offset int, limit int, MenuName_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and rediscover = ?", MenuName_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndWeight Get MenuLinkContentDatas via MenuNameAndWeight
func GetMenuLinkContentDatasByMenuNameAndWeight(offset int, limit int, MenuName_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and weight = ?", MenuName_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndExpanded Get MenuLinkContentDatas via MenuNameAndExpanded
func GetMenuLinkContentDatasByMenuNameAndExpanded(offset int, limit int, MenuName_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and expanded = ?", MenuName_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndEnabled Get MenuLinkContentDatas via MenuNameAndEnabled
func GetMenuLinkContentDatasByMenuNameAndEnabled(offset int, limit int, MenuName_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and enabled = ?", MenuName_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndParent Get MenuLinkContentDatas via MenuNameAndParent
func GetMenuLinkContentDatasByMenuNameAndParent(offset int, limit int, MenuName_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and parent = ?", MenuName_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndChanged Get MenuLinkContentDatas via MenuNameAndChanged
func GetMenuLinkContentDatasByMenuNameAndChanged(offset int, limit int, MenuName_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and changed = ?", MenuName_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByMenuNameAndDefaultLangcode Get MenuLinkContentDatas via MenuNameAndDefaultLangcode
func GetMenuLinkContentDatasByMenuNameAndDefaultLangcode(offset int, limit int, MenuName_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ? and default_langcode = ?", MenuName_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_title Get MenuLinkContentDatas via Link_uriAndLink_title
func GetMenuLinkContentDatasByLink_uriAndLink_title(offset int, limit int, Link_uri_ string, Link_title_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_title = ?", Link_uri_, Link_title_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndLink_options Get MenuLinkContentDatas via Link_uriAndLink_options
func GetMenuLinkContentDatasByLink_uriAndLink_options(offset int, limit int, Link_uri_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and link_options = ?", Link_uri_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExternal Get MenuLinkContentDatas via Link_uriAndExternal
func GetMenuLinkContentDatasByLink_uriAndExternal(offset int, limit int, Link_uri_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and external = ?", Link_uri_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndRediscover Get MenuLinkContentDatas via Link_uriAndRediscover
func GetMenuLinkContentDatasByLink_uriAndRediscover(offset int, limit int, Link_uri_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and rediscover = ?", Link_uri_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndWeight Get MenuLinkContentDatas via Link_uriAndWeight
func GetMenuLinkContentDatasByLink_uriAndWeight(offset int, limit int, Link_uri_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and weight = ?", Link_uri_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndExpanded Get MenuLinkContentDatas via Link_uriAndExpanded
func GetMenuLinkContentDatasByLink_uriAndExpanded(offset int, limit int, Link_uri_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and expanded = ?", Link_uri_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndEnabled Get MenuLinkContentDatas via Link_uriAndEnabled
func GetMenuLinkContentDatasByLink_uriAndEnabled(offset int, limit int, Link_uri_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and enabled = ?", Link_uri_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndParent Get MenuLinkContentDatas via Link_uriAndParent
func GetMenuLinkContentDatasByLink_uriAndParent(offset int, limit int, Link_uri_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and parent = ?", Link_uri_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndChanged Get MenuLinkContentDatas via Link_uriAndChanged
func GetMenuLinkContentDatasByLink_uriAndChanged(offset int, limit int, Link_uri_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and changed = ?", Link_uri_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_uriAndDefaultLangcode Get MenuLinkContentDatas via Link_uriAndDefaultLangcode
func GetMenuLinkContentDatasByLink_uriAndDefaultLangcode(offset int, limit int, Link_uri_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_uri = ? and default_langcode = ?", Link_uri_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndLink_options Get MenuLinkContentDatas via Link_titleAndLink_options
func GetMenuLinkContentDatasByLink_titleAndLink_options(offset int, limit int, Link_title_ string, Link_options_ []byte) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and link_options = ?", Link_title_, Link_options_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExternal Get MenuLinkContentDatas via Link_titleAndExternal
func GetMenuLinkContentDatasByLink_titleAndExternal(offset int, limit int, Link_title_ string, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and external = ?", Link_title_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndRediscover Get MenuLinkContentDatas via Link_titleAndRediscover
func GetMenuLinkContentDatasByLink_titleAndRediscover(offset int, limit int, Link_title_ string, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and rediscover = ?", Link_title_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndWeight Get MenuLinkContentDatas via Link_titleAndWeight
func GetMenuLinkContentDatasByLink_titleAndWeight(offset int, limit int, Link_title_ string, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and weight = ?", Link_title_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndExpanded Get MenuLinkContentDatas via Link_titleAndExpanded
func GetMenuLinkContentDatasByLink_titleAndExpanded(offset int, limit int, Link_title_ string, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and expanded = ?", Link_title_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndEnabled Get MenuLinkContentDatas via Link_titleAndEnabled
func GetMenuLinkContentDatasByLink_titleAndEnabled(offset int, limit int, Link_title_ string, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and enabled = ?", Link_title_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndParent Get MenuLinkContentDatas via Link_titleAndParent
func GetMenuLinkContentDatasByLink_titleAndParent(offset int, limit int, Link_title_ string, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and parent = ?", Link_title_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndChanged Get MenuLinkContentDatas via Link_titleAndChanged
func GetMenuLinkContentDatasByLink_titleAndChanged(offset int, limit int, Link_title_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and changed = ?", Link_title_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_titleAndDefaultLangcode Get MenuLinkContentDatas via Link_titleAndDefaultLangcode
func GetMenuLinkContentDatasByLink_titleAndDefaultLangcode(offset int, limit int, Link_title_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_title = ? and default_langcode = ?", Link_title_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExternal Get MenuLinkContentDatas via Link_optionsAndExternal
func GetMenuLinkContentDatasByLink_optionsAndExternal(offset int, limit int, Link_options_ []byte, External_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and external = ?", Link_options_, External_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndRediscover Get MenuLinkContentDatas via Link_optionsAndRediscover
func GetMenuLinkContentDatasByLink_optionsAndRediscover(offset int, limit int, Link_options_ []byte, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and rediscover = ?", Link_options_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndWeight Get MenuLinkContentDatas via Link_optionsAndWeight
func GetMenuLinkContentDatasByLink_optionsAndWeight(offset int, limit int, Link_options_ []byte, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and weight = ?", Link_options_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndExpanded Get MenuLinkContentDatas via Link_optionsAndExpanded
func GetMenuLinkContentDatasByLink_optionsAndExpanded(offset int, limit int, Link_options_ []byte, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and expanded = ?", Link_options_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndEnabled Get MenuLinkContentDatas via Link_optionsAndEnabled
func GetMenuLinkContentDatasByLink_optionsAndEnabled(offset int, limit int, Link_options_ []byte, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and enabled = ?", Link_options_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndParent Get MenuLinkContentDatas via Link_optionsAndParent
func GetMenuLinkContentDatasByLink_optionsAndParent(offset int, limit int, Link_options_ []byte, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and parent = ?", Link_options_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndChanged Get MenuLinkContentDatas via Link_optionsAndChanged
func GetMenuLinkContentDatasByLink_optionsAndChanged(offset int, limit int, Link_options_ []byte, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and changed = ?", Link_options_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByLink_optionsAndDefaultLangcode Get MenuLinkContentDatas via Link_optionsAndDefaultLangcode
func GetMenuLinkContentDatasByLink_optionsAndDefaultLangcode(offset int, limit int, Link_options_ []byte, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link_options = ? and default_langcode = ?", Link_options_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndRediscover Get MenuLinkContentDatas via ExternalAndRediscover
func GetMenuLinkContentDatasByExternalAndRediscover(offset int, limit int, External_ int, Rediscover_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and rediscover = ?", External_, Rediscover_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndWeight Get MenuLinkContentDatas via ExternalAndWeight
func GetMenuLinkContentDatasByExternalAndWeight(offset int, limit int, External_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and weight = ?", External_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndExpanded Get MenuLinkContentDatas via ExternalAndExpanded
func GetMenuLinkContentDatasByExternalAndExpanded(offset int, limit int, External_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and expanded = ?", External_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndEnabled Get MenuLinkContentDatas via ExternalAndEnabled
func GetMenuLinkContentDatasByExternalAndEnabled(offset int, limit int, External_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and enabled = ?", External_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndParent Get MenuLinkContentDatas via ExternalAndParent
func GetMenuLinkContentDatasByExternalAndParent(offset int, limit int, External_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and parent = ?", External_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndChanged Get MenuLinkContentDatas via ExternalAndChanged
func GetMenuLinkContentDatasByExternalAndChanged(offset int, limit int, External_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and changed = ?", External_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExternalAndDefaultLangcode Get MenuLinkContentDatas via ExternalAndDefaultLangcode
func GetMenuLinkContentDatasByExternalAndDefaultLangcode(offset int, limit int, External_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ? and default_langcode = ?", External_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndWeight Get MenuLinkContentDatas via RediscoverAndWeight
func GetMenuLinkContentDatasByRediscoverAndWeight(offset int, limit int, Rediscover_ int, Weight_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and weight = ?", Rediscover_, Weight_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndExpanded Get MenuLinkContentDatas via RediscoverAndExpanded
func GetMenuLinkContentDatasByRediscoverAndExpanded(offset int, limit int, Rediscover_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and expanded = ?", Rediscover_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndEnabled Get MenuLinkContentDatas via RediscoverAndEnabled
func GetMenuLinkContentDatasByRediscoverAndEnabled(offset int, limit int, Rediscover_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and enabled = ?", Rediscover_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndParent Get MenuLinkContentDatas via RediscoverAndParent
func GetMenuLinkContentDatasByRediscoverAndParent(offset int, limit int, Rediscover_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and parent = ?", Rediscover_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndChanged Get MenuLinkContentDatas via RediscoverAndChanged
func GetMenuLinkContentDatasByRediscoverAndChanged(offset int, limit int, Rediscover_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and changed = ?", Rediscover_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByRediscoverAndDefaultLangcode Get MenuLinkContentDatas via RediscoverAndDefaultLangcode
func GetMenuLinkContentDatasByRediscoverAndDefaultLangcode(offset int, limit int, Rediscover_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ? and default_langcode = ?", Rediscover_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndExpanded Get MenuLinkContentDatas via WeightAndExpanded
func GetMenuLinkContentDatasByWeightAndExpanded(offset int, limit int, Weight_ int, Expanded_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and expanded = ?", Weight_, Expanded_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndEnabled Get MenuLinkContentDatas via WeightAndEnabled
func GetMenuLinkContentDatasByWeightAndEnabled(offset int, limit int, Weight_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and enabled = ?", Weight_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndParent Get MenuLinkContentDatas via WeightAndParent
func GetMenuLinkContentDatasByWeightAndParent(offset int, limit int, Weight_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and parent = ?", Weight_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndChanged Get MenuLinkContentDatas via WeightAndChanged
func GetMenuLinkContentDatasByWeightAndChanged(offset int, limit int, Weight_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and changed = ?", Weight_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByWeightAndDefaultLangcode Get MenuLinkContentDatas via WeightAndDefaultLangcode
func GetMenuLinkContentDatasByWeightAndDefaultLangcode(offset int, limit int, Weight_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ? and default_langcode = ?", Weight_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndEnabled Get MenuLinkContentDatas via ExpandedAndEnabled
func GetMenuLinkContentDatasByExpandedAndEnabled(offset int, limit int, Expanded_ int, Enabled_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and enabled = ?", Expanded_, Enabled_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndParent Get MenuLinkContentDatas via ExpandedAndParent
func GetMenuLinkContentDatasByExpandedAndParent(offset int, limit int, Expanded_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and parent = ?", Expanded_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndChanged Get MenuLinkContentDatas via ExpandedAndChanged
func GetMenuLinkContentDatasByExpandedAndChanged(offset int, limit int, Expanded_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and changed = ?", Expanded_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByExpandedAndDefaultLangcode Get MenuLinkContentDatas via ExpandedAndDefaultLangcode
func GetMenuLinkContentDatasByExpandedAndDefaultLangcode(offset int, limit int, Expanded_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ? and default_langcode = ?", Expanded_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndParent Get MenuLinkContentDatas via EnabledAndParent
func GetMenuLinkContentDatasByEnabledAndParent(offset int, limit int, Enabled_ int, Parent_ string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and parent = ?", Enabled_, Parent_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndChanged Get MenuLinkContentDatas via EnabledAndChanged
func GetMenuLinkContentDatasByEnabledAndChanged(offset int, limit int, Enabled_ int, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and changed = ?", Enabled_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByEnabledAndDefaultLangcode Get MenuLinkContentDatas via EnabledAndDefaultLangcode
func GetMenuLinkContentDatasByEnabledAndDefaultLangcode(offset int, limit int, Enabled_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ? and default_langcode = ?", Enabled_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByParentAndChanged Get MenuLinkContentDatas via ParentAndChanged
func GetMenuLinkContentDatasByParentAndChanged(offset int, limit int, Parent_ string, Changed_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("parent = ? and changed = ?", Parent_, Changed_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByParentAndDefaultLangcode Get MenuLinkContentDatas via ParentAndDefaultLangcode
func GetMenuLinkContentDatasByParentAndDefaultLangcode(offset int, limit int, Parent_ string, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("parent = ? and default_langcode = ?", Parent_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasByChangedAndDefaultLangcode Get MenuLinkContentDatas via ChangedAndDefaultLangcode
func GetMenuLinkContentDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatas Get MenuLinkContentDatas via field
func GetMenuLinkContentDatas(offset int, limit int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaId Get MenuLinkContentDatas via Id
func GetMenuLinkContentDatasViaId(offset int, limit int, Id_ int64, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaBundle Get MenuLinkContentDatas via Bundle
func GetMenuLinkContentDatasViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaLangcode Get MenuLinkContentDatas via Langcode
func GetMenuLinkContentDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaTitle Get MenuLinkContentDatas via Title
func GetMenuLinkContentDatasViaTitle(offset int, limit int, Title_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("title = ?", Title_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaDescription Get MenuLinkContentDatas via Description
func GetMenuLinkContentDatasViaDescription(offset int, limit int, Description_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("description = ?", Description_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaMenuName Get MenuLinkContentDatas via MenuName
func GetMenuLinkContentDatasViaMenuName(offset int, limit int, MenuName_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("menu_name = ?", MenuName_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaLink_uri Get MenuLinkContentDatas via Link_uri
func GetMenuLinkContentDatasViaLink_uri(offset int, limit int, Link_uri_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link__uri = ?", Link_uri_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaLink_title Get MenuLinkContentDatas via Link_title
func GetMenuLinkContentDatasViaLink_title(offset int, limit int, Link_title_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link__title = ?", Link_title_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaLink_options Get MenuLinkContentDatas via Link_options
func GetMenuLinkContentDatasViaLink_options(offset int, limit int, Link_options_ []byte, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("link__options = ?", Link_options_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaExternal Get MenuLinkContentDatas via External
func GetMenuLinkContentDatasViaExternal(offset int, limit int, External_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("external = ?", External_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaRediscover Get MenuLinkContentDatas via Rediscover
func GetMenuLinkContentDatasViaRediscover(offset int, limit int, Rediscover_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("rediscover = ?", Rediscover_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaWeight Get MenuLinkContentDatas via Weight
func GetMenuLinkContentDatasViaWeight(offset int, limit int, Weight_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("weight = ?", Weight_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaExpanded Get MenuLinkContentDatas via Expanded
func GetMenuLinkContentDatasViaExpanded(offset int, limit int, Expanded_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("expanded = ?", Expanded_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaEnabled Get MenuLinkContentDatas via Enabled
func GetMenuLinkContentDatasViaEnabled(offset int, limit int, Enabled_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("enabled = ?", Enabled_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaParent Get MenuLinkContentDatas via Parent
func GetMenuLinkContentDatasViaParent(offset int, limit int, Parent_ string, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("parent = ?", Parent_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaChanged Get MenuLinkContentDatas via Changed
func GetMenuLinkContentDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// GetMenuLinkContentDatasViaDefaultLangcode Get MenuLinkContentDatas via DefaultLangcode
func GetMenuLinkContentDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*MenuLinkContentData, error) {
	var _MenuLinkContentData = new([]*MenuLinkContentData)
	err := Engine.Table("menu_link_content_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_MenuLinkContentData)
	return _MenuLinkContentData, err
}

// HasMenuLinkContentDataViaId Has MenuLinkContentData via Id
func HasMenuLinkContentDataViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaBundle Has MenuLinkContentData via Bundle
func HasMenuLinkContentDataViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaLangcode Has MenuLinkContentData via Langcode
func HasMenuLinkContentDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaTitle Has MenuLinkContentData via Title
func HasMenuLinkContentDataViaTitle(iTitle string) bool {
	if has, err := Engine.Where("title = ?", iTitle).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaDescription Has MenuLinkContentData via Description
func HasMenuLinkContentDataViaDescription(iDescription string) bool {
	if has, err := Engine.Where("description = ?", iDescription).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaMenuName Has MenuLinkContentData via MenuName
func HasMenuLinkContentDataViaMenuName(iMenuName string) bool {
	if has, err := Engine.Where("menu_name = ?", iMenuName).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaLink_uri Has MenuLinkContentData via Link_uri
func HasMenuLinkContentDataViaLink_uri(iLink_uri string) bool {
	if has, err := Engine.Where("link__uri = ?", iLink_uri).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaLink_title Has MenuLinkContentData via Link_title
func HasMenuLinkContentDataViaLink_title(iLink_title string) bool {
	if has, err := Engine.Where("link__title = ?", iLink_title).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaLink_options Has MenuLinkContentData via Link_options
func HasMenuLinkContentDataViaLink_options(iLink_options []byte) bool {
	if has, err := Engine.Where("link__options = ?", iLink_options).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaExternal Has MenuLinkContentData via External
func HasMenuLinkContentDataViaExternal(iExternal int) bool {
	if has, err := Engine.Where("external = ?", iExternal).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaRediscover Has MenuLinkContentData via Rediscover
func HasMenuLinkContentDataViaRediscover(iRediscover int) bool {
	if has, err := Engine.Where("rediscover = ?", iRediscover).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaWeight Has MenuLinkContentData via Weight
func HasMenuLinkContentDataViaWeight(iWeight int) bool {
	if has, err := Engine.Where("weight = ?", iWeight).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaExpanded Has MenuLinkContentData via Expanded
func HasMenuLinkContentDataViaExpanded(iExpanded int) bool {
	if has, err := Engine.Where("expanded = ?", iExpanded).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaEnabled Has MenuLinkContentData via Enabled
func HasMenuLinkContentDataViaEnabled(iEnabled int) bool {
	if has, err := Engine.Where("enabled = ?", iEnabled).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaParent Has MenuLinkContentData via Parent
func HasMenuLinkContentDataViaParent(iParent string) bool {
	if has, err := Engine.Where("parent = ?", iParent).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaChanged Has MenuLinkContentData via Changed
func HasMenuLinkContentDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentDataViaDefaultLangcode Has MenuLinkContentData via DefaultLangcode
func HasMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(MenuLinkContentData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetMenuLinkContentDataViaId Get MenuLinkContentData via Id
func GetMenuLinkContentDataViaId(iId int64) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Id: iId}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaBundle Get MenuLinkContentData via Bundle
func GetMenuLinkContentDataViaBundle(iBundle string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Bundle: iBundle}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaLangcode Get MenuLinkContentData via Langcode
func GetMenuLinkContentDataViaLangcode(iLangcode string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Langcode: iLangcode}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaTitle Get MenuLinkContentData via Title
func GetMenuLinkContentDataViaTitle(iTitle string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Title: iTitle}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaDescription Get MenuLinkContentData via Description
func GetMenuLinkContentDataViaDescription(iDescription string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Description: iDescription}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaMenuName Get MenuLinkContentData via MenuName
func GetMenuLinkContentDataViaMenuName(iMenuName string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{MenuName: iMenuName}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaLink_uri Get MenuLinkContentData via Link_uri
func GetMenuLinkContentDataViaLink_uri(iLink_uri string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Link_uri: iLink_uri}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaLink_title Get MenuLinkContentData via Link_title
func GetMenuLinkContentDataViaLink_title(iLink_title string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Link_title: iLink_title}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaLink_options Get MenuLinkContentData via Link_options
func GetMenuLinkContentDataViaLink_options(iLink_options []byte) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Link_options: iLink_options}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaExternal Get MenuLinkContentData via External
func GetMenuLinkContentDataViaExternal(iExternal int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{External: iExternal}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaRediscover Get MenuLinkContentData via Rediscover
func GetMenuLinkContentDataViaRediscover(iRediscover int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Rediscover: iRediscover}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaWeight Get MenuLinkContentData via Weight
func GetMenuLinkContentDataViaWeight(iWeight int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Weight: iWeight}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaExpanded Get MenuLinkContentData via Expanded
func GetMenuLinkContentDataViaExpanded(iExpanded int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Expanded: iExpanded}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaEnabled Get MenuLinkContentData via Enabled
func GetMenuLinkContentDataViaEnabled(iEnabled int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Enabled: iEnabled}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaParent Get MenuLinkContentData via Parent
func GetMenuLinkContentDataViaParent(iParent string) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Parent: iParent}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaChanged Get MenuLinkContentData via Changed
func GetMenuLinkContentDataViaChanged(iChanged int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{Changed: iChanged}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentDataViaDefaultLangcode Get MenuLinkContentData via DefaultLangcode
func GetMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode int) (*MenuLinkContentData, error) {
	var _MenuLinkContentData = &MenuLinkContentData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_MenuLinkContentData)
	if has {
		return _MenuLinkContentData, err
	} else {
		return nil, err
	}
}

// SetMenuLinkContentDataViaId Set MenuLinkContentData via Id
func SetMenuLinkContentDataViaId(iId int64, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Id = iId
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaBundle Set MenuLinkContentData via Bundle
func SetMenuLinkContentDataViaBundle(iBundle string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Bundle = iBundle
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaLangcode Set MenuLinkContentData via Langcode
func SetMenuLinkContentDataViaLangcode(iLangcode string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Langcode = iLangcode
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaTitle Set MenuLinkContentData via Title
func SetMenuLinkContentDataViaTitle(iTitle string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Title = iTitle
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaDescription Set MenuLinkContentData via Description
func SetMenuLinkContentDataViaDescription(iDescription string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Description = iDescription
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaMenuName Set MenuLinkContentData via MenuName
func SetMenuLinkContentDataViaMenuName(iMenuName string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.MenuName = iMenuName
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaLink_uri Set MenuLinkContentData via Link_uri
func SetMenuLinkContentDataViaLink_uri(iLink_uri string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Link_uri = iLink_uri
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaLink_title Set MenuLinkContentData via Link_title
func SetMenuLinkContentDataViaLink_title(iLink_title string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Link_title = iLink_title
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaLink_options Set MenuLinkContentData via Link_options
func SetMenuLinkContentDataViaLink_options(iLink_options []byte, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Link_options = iLink_options
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaExternal Set MenuLinkContentData via External
func SetMenuLinkContentDataViaExternal(iExternal int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.External = iExternal
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaRediscover Set MenuLinkContentData via Rediscover
func SetMenuLinkContentDataViaRediscover(iRediscover int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Rediscover = iRediscover
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaWeight Set MenuLinkContentData via Weight
func SetMenuLinkContentDataViaWeight(iWeight int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Weight = iWeight
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaExpanded Set MenuLinkContentData via Expanded
func SetMenuLinkContentDataViaExpanded(iExpanded int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Expanded = iExpanded
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaEnabled Set MenuLinkContentData via Enabled
func SetMenuLinkContentDataViaEnabled(iEnabled int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Enabled = iEnabled
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaParent Set MenuLinkContentData via Parent
func SetMenuLinkContentDataViaParent(iParent string, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Parent = iParent
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaChanged Set MenuLinkContentData via Changed
func SetMenuLinkContentDataViaChanged(iChanged int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.Changed = iChanged
	return Engine.Insert(menu_link_content_data)
}

// SetMenuLinkContentDataViaDefaultLangcode Set MenuLinkContentData via DefaultLangcode
func SetMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode int, menu_link_content_data *MenuLinkContentData) (int64, error) {
	menu_link_content_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(menu_link_content_data)
}

// AddMenuLinkContentData Add MenuLinkContentData via all columns
func AddMenuLinkContentData(iId int64, iBundle string, iLangcode string, iTitle string, iDescription string, iMenuName string, iLink_uri string, iLink_title string, iLink_options []byte, iExternal int, iRediscover int, iWeight int, iExpanded int, iEnabled int, iParent string, iChanged int, iDefaultLangcode int) error {
	_MenuLinkContentData := &MenuLinkContentData{Id: iId, Bundle: iBundle, Langcode: iLangcode, Title: iTitle, Description: iDescription, MenuName: iMenuName, Link_uri: iLink_uri, Link_title: iLink_title, Link_options: iLink_options, External: iExternal, Rediscover: iRediscover, Weight: iWeight, Expanded: iExpanded, Enabled: iEnabled, Parent: iParent, Changed: iChanged, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_MenuLinkContentData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostMenuLinkContentData Post MenuLinkContentData via iMenuLinkContentData
func PostMenuLinkContentData(iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	_, err := Engine.Insert(iMenuLinkContentData)
	return iMenuLinkContentData.Id, err
}

// PutMenuLinkContentData Put MenuLinkContentData
func PutMenuLinkContentData(iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	_, err := Engine.Id(iMenuLinkContentData.Id).Update(iMenuLinkContentData)
	return iMenuLinkContentData.Id, err
}

// PutMenuLinkContentDataViaId Put MenuLinkContentData via Id
func PutMenuLinkContentDataViaId(Id_ int64, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Id: Id_})
	return row, err
}

// PutMenuLinkContentDataViaBundle Put MenuLinkContentData via Bundle
func PutMenuLinkContentDataViaBundle(Bundle_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Bundle: Bundle_})
	return row, err
}

// PutMenuLinkContentDataViaLangcode Put MenuLinkContentData via Langcode
func PutMenuLinkContentDataViaLangcode(Langcode_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Langcode: Langcode_})
	return row, err
}

// PutMenuLinkContentDataViaTitle Put MenuLinkContentData via Title
func PutMenuLinkContentDataViaTitle(Title_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Title: Title_})
	return row, err
}

// PutMenuLinkContentDataViaDescription Put MenuLinkContentData via Description
func PutMenuLinkContentDataViaDescription(Description_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Description: Description_})
	return row, err
}

// PutMenuLinkContentDataViaMenuName Put MenuLinkContentData via MenuName
func PutMenuLinkContentDataViaMenuName(MenuName_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{MenuName: MenuName_})
	return row, err
}

// PutMenuLinkContentDataViaLink_uri Put MenuLinkContentData via Link_uri
func PutMenuLinkContentDataViaLink_uri(Link_uri_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Link_uri: Link_uri_})
	return row, err
}

// PutMenuLinkContentDataViaLink_title Put MenuLinkContentData via Link_title
func PutMenuLinkContentDataViaLink_title(Link_title_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Link_title: Link_title_})
	return row, err
}

// PutMenuLinkContentDataViaLink_options Put MenuLinkContentData via Link_options
func PutMenuLinkContentDataViaLink_options(Link_options_ []byte, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Link_options: Link_options_})
	return row, err
}

// PutMenuLinkContentDataViaExternal Put MenuLinkContentData via External
func PutMenuLinkContentDataViaExternal(External_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{External: External_})
	return row, err
}

// PutMenuLinkContentDataViaRediscover Put MenuLinkContentData via Rediscover
func PutMenuLinkContentDataViaRediscover(Rediscover_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Rediscover: Rediscover_})
	return row, err
}

// PutMenuLinkContentDataViaWeight Put MenuLinkContentData via Weight
func PutMenuLinkContentDataViaWeight(Weight_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Weight: Weight_})
	return row, err
}

// PutMenuLinkContentDataViaExpanded Put MenuLinkContentData via Expanded
func PutMenuLinkContentDataViaExpanded(Expanded_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Expanded: Expanded_})
	return row, err
}

// PutMenuLinkContentDataViaEnabled Put MenuLinkContentData via Enabled
func PutMenuLinkContentDataViaEnabled(Enabled_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Enabled: Enabled_})
	return row, err
}

// PutMenuLinkContentDataViaParent Put MenuLinkContentData via Parent
func PutMenuLinkContentDataViaParent(Parent_ string, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Parent: Parent_})
	return row, err
}

// PutMenuLinkContentDataViaChanged Put MenuLinkContentData via Changed
func PutMenuLinkContentDataViaChanged(Changed_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{Changed: Changed_})
	return row, err
}

// PutMenuLinkContentDataViaDefaultLangcode Put MenuLinkContentData via DefaultLangcode
func PutMenuLinkContentDataViaDefaultLangcode(DefaultLangcode_ int, iMenuLinkContentData *MenuLinkContentData) (int64, error) {
	row, err := Engine.Update(iMenuLinkContentData, &MenuLinkContentData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateMenuLinkContentDataViaId via map[string]interface{}{}
func UpdateMenuLinkContentDataViaId(iId int64, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("id = ?", iId).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaBundle via map[string]interface{}{}
func UpdateMenuLinkContentDataViaBundle(iBundle string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("bundle = ?", iBundle).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaLangcode via map[string]interface{}{}
func UpdateMenuLinkContentDataViaLangcode(iLangcode string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("langcode = ?", iLangcode).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaTitle via map[string]interface{}{}
func UpdateMenuLinkContentDataViaTitle(iTitle string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("title = ?", iTitle).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaDescription via map[string]interface{}{}
func UpdateMenuLinkContentDataViaDescription(iDescription string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("description = ?", iDescription).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaMenuName via map[string]interface{}{}
func UpdateMenuLinkContentDataViaMenuName(iMenuName string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("menu_name = ?", iMenuName).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaLink_uri via map[string]interface{}{}
func UpdateMenuLinkContentDataViaLink_uri(iLink_uri string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("link__uri = ?", iLink_uri).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaLink_title via map[string]interface{}{}
func UpdateMenuLinkContentDataViaLink_title(iLink_title string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("link__title = ?", iLink_title).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaLink_options via map[string]interface{}{}
func UpdateMenuLinkContentDataViaLink_options(iLink_options []byte, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("link__options = ?", iLink_options).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaExternal via map[string]interface{}{}
func UpdateMenuLinkContentDataViaExternal(iExternal int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("external = ?", iExternal).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaRediscover via map[string]interface{}{}
func UpdateMenuLinkContentDataViaRediscover(iRediscover int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("rediscover = ?", iRediscover).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaWeight via map[string]interface{}{}
func UpdateMenuLinkContentDataViaWeight(iWeight int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("weight = ?", iWeight).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaExpanded via map[string]interface{}{}
func UpdateMenuLinkContentDataViaExpanded(iExpanded int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("expanded = ?", iExpanded).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaEnabled via map[string]interface{}{}
func UpdateMenuLinkContentDataViaEnabled(iEnabled int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("enabled = ?", iEnabled).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaParent via map[string]interface{}{}
func UpdateMenuLinkContentDataViaParent(iParent string, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("parent = ?", iParent).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaChanged via map[string]interface{}{}
func UpdateMenuLinkContentDataViaChanged(iChanged int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("changed = ?", iChanged).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentDataViaDefaultLangcode via map[string]interface{}{}
func UpdateMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode int, iMenuLinkContentDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContentData)).Where("default_langcode = ?", iDefaultLangcode).Update(iMenuLinkContentDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteMenuLinkContentData Delete MenuLinkContentData
func DeleteMenuLinkContentData(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(MenuLinkContentData))
	return row, err
}

// DeleteMenuLinkContentDataViaId Delete MenuLinkContentData via Id
func DeleteMenuLinkContentDataViaId(iId int64) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Id: iId}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaBundle Delete MenuLinkContentData via Bundle
func DeleteMenuLinkContentDataViaBundle(iBundle string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Bundle: iBundle}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaLangcode Delete MenuLinkContentData via Langcode
func DeleteMenuLinkContentDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Langcode: iLangcode}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaTitle Delete MenuLinkContentData via Title
func DeleteMenuLinkContentDataViaTitle(iTitle string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Title: iTitle}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("title = ?", iTitle).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaDescription Delete MenuLinkContentData via Description
func DeleteMenuLinkContentDataViaDescription(iDescription string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Description: iDescription}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("description = ?", iDescription).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaMenuName Delete MenuLinkContentData via MenuName
func DeleteMenuLinkContentDataViaMenuName(iMenuName string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{MenuName: iMenuName}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("menu_name = ?", iMenuName).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaLink_uri Delete MenuLinkContentData via Link_uri
func DeleteMenuLinkContentDataViaLink_uri(iLink_uri string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Link_uri: iLink_uri}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__uri = ?", iLink_uri).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaLink_title Delete MenuLinkContentData via Link_title
func DeleteMenuLinkContentDataViaLink_title(iLink_title string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Link_title: iLink_title}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__title = ?", iLink_title).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaLink_options Delete MenuLinkContentData via Link_options
func DeleteMenuLinkContentDataViaLink_options(iLink_options []byte) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Link_options: iLink_options}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("link__options = ?", iLink_options).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaExternal Delete MenuLinkContentData via External
func DeleteMenuLinkContentDataViaExternal(iExternal int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{External: iExternal}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("external = ?", iExternal).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaRediscover Delete MenuLinkContentData via Rediscover
func DeleteMenuLinkContentDataViaRediscover(iRediscover int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Rediscover: iRediscover}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("rediscover = ?", iRediscover).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaWeight Delete MenuLinkContentData via Weight
func DeleteMenuLinkContentDataViaWeight(iWeight int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Weight: iWeight}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("weight = ?", iWeight).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaExpanded Delete MenuLinkContentData via Expanded
func DeleteMenuLinkContentDataViaExpanded(iExpanded int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Expanded: iExpanded}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("expanded = ?", iExpanded).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaEnabled Delete MenuLinkContentData via Enabled
func DeleteMenuLinkContentDataViaEnabled(iEnabled int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Enabled: iEnabled}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("enabled = ?", iEnabled).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaParent Delete MenuLinkContentData via Parent
func DeleteMenuLinkContentDataViaParent(iParent string) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Parent: iParent}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("parent = ?", iParent).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaChanged Delete MenuLinkContentData via Changed
func DeleteMenuLinkContentDataViaChanged(iChanged int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{Changed: iChanged}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentDataViaDefaultLangcode Delete MenuLinkContentData via DefaultLangcode
func DeleteMenuLinkContentDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _MenuLinkContentData = &MenuLinkContentData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_MenuLinkContentData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(MenuLinkContentData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
