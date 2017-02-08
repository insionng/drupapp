package model

type MenuLinkContent struct {
	Id       int64  `xorm:"not null pk autoincr INT(10)"`
	Bundle   string `xorm:"not null VARCHAR(32)"`
	Uuid     string `xorm:"not null unique VARCHAR(128)"`
	Langcode string `xorm:"not null VARCHAR(12)"`
}

// GetMenuLinkContentsCount MenuLinkContents' Count
func GetMenuLinkContentsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&MenuLinkContent{})
	return total, err
}

// GetMenuLinkContentCountViaId Get MenuLinkContent via Id
func GetMenuLinkContentCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&MenuLinkContent{Id: iId})
	return n
}

// GetMenuLinkContentCountViaBundle Get MenuLinkContent via Bundle
func GetMenuLinkContentCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&MenuLinkContent{Bundle: iBundle})
	return n
}

// GetMenuLinkContentCountViaUuid Get MenuLinkContent via Uuid
func GetMenuLinkContentCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&MenuLinkContent{Uuid: iUuid})
	return n
}

// GetMenuLinkContentCountViaLangcode Get MenuLinkContent via Langcode
func GetMenuLinkContentCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&MenuLinkContent{Langcode: iLangcode})
	return n
}

// GetMenuLinkContentsByIdAndBundleAndUuid Get MenuLinkContents via IdAndBundleAndUuid
func GetMenuLinkContentsByIdAndBundleAndUuid(offset int, limit int, Id_ int64, Bundle_ string, Uuid_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and bundle = ? and uuid = ?", Id_, Bundle_, Uuid_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByIdAndBundleAndLangcode Get MenuLinkContents via IdAndBundleAndLangcode
func GetMenuLinkContentsByIdAndBundleAndLangcode(offset int, limit int, Id_ int64, Bundle_ string, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and bundle = ? and langcode = ?", Id_, Bundle_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByIdAndUuidAndLangcode Get MenuLinkContents via IdAndUuidAndLangcode
func GetMenuLinkContentsByIdAndUuidAndLangcode(offset int, limit int, Id_ int64, Uuid_ string, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and uuid = ? and langcode = ?", Id_, Uuid_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByBundleAndUuidAndLangcode Get MenuLinkContents via BundleAndUuidAndLangcode
func GetMenuLinkContentsByBundleAndUuidAndLangcode(offset int, limit int, Bundle_ string, Uuid_ string, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("bundle = ? and uuid = ? and langcode = ?", Bundle_, Uuid_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByIdAndBundle Get MenuLinkContents via IdAndBundle
func GetMenuLinkContentsByIdAndBundle(offset int, limit int, Id_ int64, Bundle_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and bundle = ?", Id_, Bundle_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByIdAndUuid Get MenuLinkContents via IdAndUuid
func GetMenuLinkContentsByIdAndUuid(offset int, limit int, Id_ int64, Uuid_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and uuid = ?", Id_, Uuid_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByIdAndLangcode Get MenuLinkContents via IdAndLangcode
func GetMenuLinkContentsByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByBundleAndUuid Get MenuLinkContents via BundleAndUuid
func GetMenuLinkContentsByBundleAndUuid(offset int, limit int, Bundle_ string, Uuid_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("bundle = ? and uuid = ?", Bundle_, Uuid_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByBundleAndLangcode Get MenuLinkContents via BundleAndLangcode
func GetMenuLinkContentsByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsByUuidAndLangcode Get MenuLinkContents via UuidAndLangcode
func GetMenuLinkContentsByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContents Get MenuLinkContents via field
func GetMenuLinkContents(offset int, limit int, field string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Limit(limit, offset).Desc(field).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsViaId Get MenuLinkContents via Id
func GetMenuLinkContentsViaId(offset int, limit int, Id_ int64, field string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsViaBundle Get MenuLinkContents via Bundle
func GetMenuLinkContentsViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsViaUuid Get MenuLinkContents via Uuid
func GetMenuLinkContentsViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// GetMenuLinkContentsViaLangcode Get MenuLinkContents via Langcode
func GetMenuLinkContentsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*MenuLinkContent, error) {
	var _MenuLinkContent = new([]*MenuLinkContent)
	err := Engine.Table("menu_link_content").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_MenuLinkContent)
	return _MenuLinkContent, err
}

// HasMenuLinkContentViaId Has MenuLinkContent via Id
func HasMenuLinkContentViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(MenuLinkContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentViaBundle Has MenuLinkContent via Bundle
func HasMenuLinkContentViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(MenuLinkContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentViaUuid Has MenuLinkContent via Uuid
func HasMenuLinkContentViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(MenuLinkContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasMenuLinkContentViaLangcode Has MenuLinkContent via Langcode
func HasMenuLinkContentViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(MenuLinkContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetMenuLinkContentViaId Get MenuLinkContent via Id
func GetMenuLinkContentViaId(iId int64) (*MenuLinkContent, error) {
	var _MenuLinkContent = &MenuLinkContent{Id: iId}
	has, err := Engine.Get(_MenuLinkContent)
	if has {
		return _MenuLinkContent, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentViaBundle Get MenuLinkContent via Bundle
func GetMenuLinkContentViaBundle(iBundle string) (*MenuLinkContent, error) {
	var _MenuLinkContent = &MenuLinkContent{Bundle: iBundle}
	has, err := Engine.Get(_MenuLinkContent)
	if has {
		return _MenuLinkContent, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentViaUuid Get MenuLinkContent via Uuid
func GetMenuLinkContentViaUuid(iUuid string) (*MenuLinkContent, error) {
	var _MenuLinkContent = &MenuLinkContent{Uuid: iUuid}
	has, err := Engine.Get(_MenuLinkContent)
	if has {
		return _MenuLinkContent, err
	} else {
		return nil, err
	}
}

// GetMenuLinkContentViaLangcode Get MenuLinkContent via Langcode
func GetMenuLinkContentViaLangcode(iLangcode string) (*MenuLinkContent, error) {
	var _MenuLinkContent = &MenuLinkContent{Langcode: iLangcode}
	has, err := Engine.Get(_MenuLinkContent)
	if has {
		return _MenuLinkContent, err
	} else {
		return nil, err
	}
}

// SetMenuLinkContentViaId Set MenuLinkContent via Id
func SetMenuLinkContentViaId(iId int64, menu_link_content *MenuLinkContent) (int64, error) {
	menu_link_content.Id = iId
	return Engine.Insert(menu_link_content)
}

// SetMenuLinkContentViaBundle Set MenuLinkContent via Bundle
func SetMenuLinkContentViaBundle(iBundle string, menu_link_content *MenuLinkContent) (int64, error) {
	menu_link_content.Bundle = iBundle
	return Engine.Insert(menu_link_content)
}

// SetMenuLinkContentViaUuid Set MenuLinkContent via Uuid
func SetMenuLinkContentViaUuid(iUuid string, menu_link_content *MenuLinkContent) (int64, error) {
	menu_link_content.Uuid = iUuid
	return Engine.Insert(menu_link_content)
}

// SetMenuLinkContentViaLangcode Set MenuLinkContent via Langcode
func SetMenuLinkContentViaLangcode(iLangcode string, menu_link_content *MenuLinkContent) (int64, error) {
	menu_link_content.Langcode = iLangcode
	return Engine.Insert(menu_link_content)
}

// AddMenuLinkContent Add MenuLinkContent via all columns
func AddMenuLinkContent(iId int64, iBundle string, iUuid string, iLangcode string) error {
	_MenuLinkContent := &MenuLinkContent{Id: iId, Bundle: iBundle, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_MenuLinkContent); err != nil {
		return err
	} else {
		return nil
	}
}

// PostMenuLinkContent Post MenuLinkContent via iMenuLinkContent
func PostMenuLinkContent(iMenuLinkContent *MenuLinkContent) (int64, error) {
	_, err := Engine.Insert(iMenuLinkContent)
	return iMenuLinkContent.Id, err
}

// PutMenuLinkContent Put MenuLinkContent
func PutMenuLinkContent(iMenuLinkContent *MenuLinkContent) (int64, error) {
	_, err := Engine.Id(iMenuLinkContent.Id).Update(iMenuLinkContent)
	return iMenuLinkContent.Id, err
}

// PutMenuLinkContentViaId Put MenuLinkContent via Id
func PutMenuLinkContentViaId(Id_ int64, iMenuLinkContent *MenuLinkContent) (int64, error) {
	row, err := Engine.Update(iMenuLinkContent, &MenuLinkContent{Id: Id_})
	return row, err
}

// PutMenuLinkContentViaBundle Put MenuLinkContent via Bundle
func PutMenuLinkContentViaBundle(Bundle_ string, iMenuLinkContent *MenuLinkContent) (int64, error) {
	row, err := Engine.Update(iMenuLinkContent, &MenuLinkContent{Bundle: Bundle_})
	return row, err
}

// PutMenuLinkContentViaUuid Put MenuLinkContent via Uuid
func PutMenuLinkContentViaUuid(Uuid_ string, iMenuLinkContent *MenuLinkContent) (int64, error) {
	row, err := Engine.Update(iMenuLinkContent, &MenuLinkContent{Uuid: Uuid_})
	return row, err
}

// PutMenuLinkContentViaLangcode Put MenuLinkContent via Langcode
func PutMenuLinkContentViaLangcode(Langcode_ string, iMenuLinkContent *MenuLinkContent) (int64, error) {
	row, err := Engine.Update(iMenuLinkContent, &MenuLinkContent{Langcode: Langcode_})
	return row, err
}

// UpdateMenuLinkContentViaId via map[string]interface{}{}
func UpdateMenuLinkContentViaId(iId int64, iMenuLinkContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContent)).Where("id = ?", iId).Update(iMenuLinkContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentViaBundle via map[string]interface{}{}
func UpdateMenuLinkContentViaBundle(iBundle string, iMenuLinkContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContent)).Where("bundle = ?", iBundle).Update(iMenuLinkContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentViaUuid via map[string]interface{}{}
func UpdateMenuLinkContentViaUuid(iUuid string, iMenuLinkContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContent)).Where("uuid = ?", iUuid).Update(iMenuLinkContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateMenuLinkContentViaLangcode via map[string]interface{}{}
func UpdateMenuLinkContentViaLangcode(iLangcode string, iMenuLinkContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(MenuLinkContent)).Where("langcode = ?", iLangcode).Update(iMenuLinkContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteMenuLinkContent Delete MenuLinkContent
func DeleteMenuLinkContent(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(MenuLinkContent))
	return row, err
}

// DeleteMenuLinkContentViaId Delete MenuLinkContent via Id
func DeleteMenuLinkContentViaId(iId int64) (err error) {
	var has bool
	var _MenuLinkContent = &MenuLinkContent{Id: iId}
	if has, err = Engine.Get(_MenuLinkContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(MenuLinkContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentViaBundle Delete MenuLinkContent via Bundle
func DeleteMenuLinkContentViaBundle(iBundle string) (err error) {
	var has bool
	var _MenuLinkContent = &MenuLinkContent{Bundle: iBundle}
	if has, err = Engine.Get(_MenuLinkContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(MenuLinkContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentViaUuid Delete MenuLinkContent via Uuid
func DeleteMenuLinkContentViaUuid(iUuid string) (err error) {
	var has bool
	var _MenuLinkContent = &MenuLinkContent{Uuid: iUuid}
	if has, err = Engine.Get(_MenuLinkContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(MenuLinkContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteMenuLinkContentViaLangcode Delete MenuLinkContent via Langcode
func DeleteMenuLinkContentViaLangcode(iLangcode string) (err error) {
	var has bool
	var _MenuLinkContent = &MenuLinkContent{Langcode: iLangcode}
	if has, err = Engine.Get(_MenuLinkContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(MenuLinkContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
