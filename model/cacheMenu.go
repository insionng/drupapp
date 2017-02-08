package model

type CacheMenu struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheMenusCount CacheMenus' Count
func GetCacheMenusCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheMenu{})
	return total, err
}

// GetCacheMenuCountViaCid Get CacheMenu via Cid
func GetCacheMenuCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheMenu{Cid: iCid})
	return n
}

// GetCacheMenuCountViaData Get CacheMenu via Data
func GetCacheMenuCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheMenu{Data: iData})
	return n
}

// GetCacheMenuCountViaExpire Get CacheMenu via Expire
func GetCacheMenuCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheMenu{Expire: iExpire})
	return n
}

// GetCacheMenuCountViaCreated Get CacheMenu via Created
func GetCacheMenuCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheMenu{Created: iCreated})
	return n
}

// GetCacheMenuCountViaSerialized Get CacheMenu via Serialized
func GetCacheMenuCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheMenu{Serialized: iSerialized})
	return n
}

// GetCacheMenuCountViaTags Get CacheMenu via Tags
func GetCacheMenuCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheMenu{Tags: iTags})
	return n
}

// GetCacheMenuCountViaChecksum Get CacheMenu via Checksum
func GetCacheMenuCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheMenu{Checksum: iChecksum})
	return n
}

// GetCacheMenusByCidAndDataAndExpire Get CacheMenus via CidAndDataAndExpire
func GetCacheMenusByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndDataAndCreated Get CacheMenus via CidAndDataAndCreated
func GetCacheMenusByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndDataAndSerialized Get CacheMenus via CidAndDataAndSerialized
func GetCacheMenusByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndDataAndTags Get CacheMenus via CidAndDataAndTags
func GetCacheMenusByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndDataAndChecksum Get CacheMenus via CidAndDataAndChecksum
func GetCacheMenusByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndExpireAndCreated Get CacheMenus via CidAndExpireAndCreated
func GetCacheMenusByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndExpireAndSerialized Get CacheMenus via CidAndExpireAndSerialized
func GetCacheMenusByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndExpireAndTags Get CacheMenus via CidAndExpireAndTags
func GetCacheMenusByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndExpireAndChecksum Get CacheMenus via CidAndExpireAndChecksum
func GetCacheMenusByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndCreatedAndSerialized Get CacheMenus via CidAndCreatedAndSerialized
func GetCacheMenusByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndCreatedAndTags Get CacheMenus via CidAndCreatedAndTags
func GetCacheMenusByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndCreatedAndChecksum Get CacheMenus via CidAndCreatedAndChecksum
func GetCacheMenusByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndSerializedAndTags Get CacheMenus via CidAndSerializedAndTags
func GetCacheMenusByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndSerializedAndChecksum Get CacheMenus via CidAndSerializedAndChecksum
func GetCacheMenusByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndTagsAndChecksum Get CacheMenus via CidAndTagsAndChecksum
func GetCacheMenusByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndExpireAndCreated Get CacheMenus via DataAndExpireAndCreated
func GetCacheMenusByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndExpireAndSerialized Get CacheMenus via DataAndExpireAndSerialized
func GetCacheMenusByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndExpireAndTags Get CacheMenus via DataAndExpireAndTags
func GetCacheMenusByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndExpireAndChecksum Get CacheMenus via DataAndExpireAndChecksum
func GetCacheMenusByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndCreatedAndSerialized Get CacheMenus via DataAndCreatedAndSerialized
func GetCacheMenusByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndCreatedAndTags Get CacheMenus via DataAndCreatedAndTags
func GetCacheMenusByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndCreatedAndChecksum Get CacheMenus via DataAndCreatedAndChecksum
func GetCacheMenusByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndSerializedAndTags Get CacheMenus via DataAndSerializedAndTags
func GetCacheMenusByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndSerializedAndChecksum Get CacheMenus via DataAndSerializedAndChecksum
func GetCacheMenusByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndTagsAndChecksum Get CacheMenus via DataAndTagsAndChecksum
func GetCacheMenusByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndCreatedAndSerialized Get CacheMenus via ExpireAndCreatedAndSerialized
func GetCacheMenusByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndCreatedAndTags Get CacheMenus via ExpireAndCreatedAndTags
func GetCacheMenusByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndCreatedAndChecksum Get CacheMenus via ExpireAndCreatedAndChecksum
func GetCacheMenusByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndSerializedAndTags Get CacheMenus via ExpireAndSerializedAndTags
func GetCacheMenusByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndSerializedAndChecksum Get CacheMenus via ExpireAndSerializedAndChecksum
func GetCacheMenusByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndTagsAndChecksum Get CacheMenus via ExpireAndTagsAndChecksum
func GetCacheMenusByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndSerializedAndTags Get CacheMenus via CreatedAndSerializedAndTags
func GetCacheMenusByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndSerializedAndChecksum Get CacheMenus via CreatedAndSerializedAndChecksum
func GetCacheMenusByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndTagsAndChecksum Get CacheMenus via CreatedAndTagsAndChecksum
func GetCacheMenusByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusBySerializedAndTagsAndChecksum Get CacheMenus via SerializedAndTagsAndChecksum
func GetCacheMenusBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndData Get CacheMenus via CidAndData
func GetCacheMenusByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndExpire Get CacheMenus via CidAndExpire
func GetCacheMenusByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndCreated Get CacheMenus via CidAndCreated
func GetCacheMenusByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndSerialized Get CacheMenus via CidAndSerialized
func GetCacheMenusByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndTags Get CacheMenus via CidAndTags
func GetCacheMenusByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCidAndChecksum Get CacheMenus via CidAndChecksum
func GetCacheMenusByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndExpire Get CacheMenus via DataAndExpire
func GetCacheMenusByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndCreated Get CacheMenus via DataAndCreated
func GetCacheMenusByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndSerialized Get CacheMenus via DataAndSerialized
func GetCacheMenusByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndTags Get CacheMenus via DataAndTags
func GetCacheMenusByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByDataAndChecksum Get CacheMenus via DataAndChecksum
func GetCacheMenusByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndCreated Get CacheMenus via ExpireAndCreated
func GetCacheMenusByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndSerialized Get CacheMenus via ExpireAndSerialized
func GetCacheMenusByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndTags Get CacheMenus via ExpireAndTags
func GetCacheMenusByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByExpireAndChecksum Get CacheMenus via ExpireAndChecksum
func GetCacheMenusByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndSerialized Get CacheMenus via CreatedAndSerialized
func GetCacheMenusByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndTags Get CacheMenus via CreatedAndTags
func GetCacheMenusByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByCreatedAndChecksum Get CacheMenus via CreatedAndChecksum
func GetCacheMenusByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusBySerializedAndTags Get CacheMenus via SerializedAndTags
func GetCacheMenusBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusBySerializedAndChecksum Get CacheMenus via SerializedAndChecksum
func GetCacheMenusBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusByTagsAndChecksum Get CacheMenus via TagsAndChecksum
func GetCacheMenusByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenus Get CacheMenus via field
func GetCacheMenus(offset int, limit int, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaCid Get CacheMenus via Cid
func GetCacheMenusViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaData Get CacheMenus via Data
func GetCacheMenusViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaExpire Get CacheMenus via Expire
func GetCacheMenusViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaCreated Get CacheMenus via Created
func GetCacheMenusViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaSerialized Get CacheMenus via Serialized
func GetCacheMenusViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaTags Get CacheMenus via Tags
func GetCacheMenusViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// GetCacheMenusViaChecksum Get CacheMenus via Checksum
func GetCacheMenusViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheMenu, error) {
	var _CacheMenu = new([]*CacheMenu)
	err := Engine.Table("cache_menu").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheMenu)
	return _CacheMenu, err
}

// HasCacheMenuViaCid Has CacheMenu via Cid
func HasCacheMenuViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaData Has CacheMenu via Data
func HasCacheMenuViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaExpire Has CacheMenu via Expire
func HasCacheMenuViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaCreated Has CacheMenu via Created
func HasCacheMenuViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaSerialized Has CacheMenu via Serialized
func HasCacheMenuViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaTags Has CacheMenu via Tags
func HasCacheMenuViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheMenuViaChecksum Has CacheMenu via Checksum
func HasCacheMenuViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheMenu)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheMenuViaCid Get CacheMenu via Cid
func GetCacheMenuViaCid(iCid string) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Cid: iCid}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaData Get CacheMenu via Data
func GetCacheMenuViaData(iData []byte) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Data: iData}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaExpire Get CacheMenu via Expire
func GetCacheMenuViaExpire(iExpire int) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Expire: iExpire}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaCreated Get CacheMenu via Created
func GetCacheMenuViaCreated(iCreated string) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Created: iCreated}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaSerialized Get CacheMenu via Serialized
func GetCacheMenuViaSerialized(iSerialized int) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Serialized: iSerialized}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaTags Get CacheMenu via Tags
func GetCacheMenuViaTags(iTags string) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Tags: iTags}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// GetCacheMenuViaChecksum Get CacheMenu via Checksum
func GetCacheMenuViaChecksum(iChecksum string) (*CacheMenu, error) {
	var _CacheMenu = &CacheMenu{Checksum: iChecksum}
	has, err := Engine.Get(_CacheMenu)
	if has {
		return _CacheMenu, err
	} else {
		return nil, err
	}
}

// SetCacheMenuViaCid Set CacheMenu via Cid
func SetCacheMenuViaCid(iCid string, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Cid = iCid
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaData Set CacheMenu via Data
func SetCacheMenuViaData(iData []byte, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Data = iData
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaExpire Set CacheMenu via Expire
func SetCacheMenuViaExpire(iExpire int, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Expire = iExpire
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaCreated Set CacheMenu via Created
func SetCacheMenuViaCreated(iCreated string, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Created = iCreated
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaSerialized Set CacheMenu via Serialized
func SetCacheMenuViaSerialized(iSerialized int, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Serialized = iSerialized
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaTags Set CacheMenu via Tags
func SetCacheMenuViaTags(iTags string, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Tags = iTags
	return Engine.Insert(cache_menu)
}

// SetCacheMenuViaChecksum Set CacheMenu via Checksum
func SetCacheMenuViaChecksum(iChecksum string, cache_menu *CacheMenu) (int64, error) {
	cache_menu.Checksum = iChecksum
	return Engine.Insert(cache_menu)
}

// AddCacheMenu Add CacheMenu via all columns
func AddCacheMenu(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheMenu := &CacheMenu{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheMenu); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheMenu Post CacheMenu via iCacheMenu
func PostCacheMenu(iCacheMenu *CacheMenu) (string, error) {
	_, err := Engine.Insert(iCacheMenu)
	return iCacheMenu.Cid, err
}

// PutCacheMenu Put CacheMenu
func PutCacheMenu(iCacheMenu *CacheMenu) (string, error) {
	_, err := Engine.Id(iCacheMenu.Cid).Update(iCacheMenu)
	return iCacheMenu.Cid, err
}

// PutCacheMenuViaCid Put CacheMenu via Cid
func PutCacheMenuViaCid(Cid_ string, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Cid: Cid_})
	return row, err
}

// PutCacheMenuViaData Put CacheMenu via Data
func PutCacheMenuViaData(Data_ []byte, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Data: Data_})
	return row, err
}

// PutCacheMenuViaExpire Put CacheMenu via Expire
func PutCacheMenuViaExpire(Expire_ int, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Expire: Expire_})
	return row, err
}

// PutCacheMenuViaCreated Put CacheMenu via Created
func PutCacheMenuViaCreated(Created_ string, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Created: Created_})
	return row, err
}

// PutCacheMenuViaSerialized Put CacheMenu via Serialized
func PutCacheMenuViaSerialized(Serialized_ int, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Serialized: Serialized_})
	return row, err
}

// PutCacheMenuViaTags Put CacheMenu via Tags
func PutCacheMenuViaTags(Tags_ string, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Tags: Tags_})
	return row, err
}

// PutCacheMenuViaChecksum Put CacheMenu via Checksum
func PutCacheMenuViaChecksum(Checksum_ string, iCacheMenu *CacheMenu) (int64, error) {
	row, err := Engine.Update(iCacheMenu, &CacheMenu{Checksum: Checksum_})
	return row, err
}

// UpdateCacheMenuViaCid via map[string]interface{}{}
func UpdateCacheMenuViaCid(iCid string, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("cid = ?", iCid).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaData via map[string]interface{}{}
func UpdateCacheMenuViaData(iData []byte, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("data = ?", iData).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaExpire via map[string]interface{}{}
func UpdateCacheMenuViaExpire(iExpire int, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("expire = ?", iExpire).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaCreated via map[string]interface{}{}
func UpdateCacheMenuViaCreated(iCreated string, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("created = ?", iCreated).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaSerialized via map[string]interface{}{}
func UpdateCacheMenuViaSerialized(iSerialized int, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("serialized = ?", iSerialized).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaTags via map[string]interface{}{}
func UpdateCacheMenuViaTags(iTags string, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("tags = ?", iTags).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheMenuViaChecksum via map[string]interface{}{}
func UpdateCacheMenuViaChecksum(iChecksum string, iCacheMenuMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheMenu)).Where("checksum = ?", iChecksum).Update(iCacheMenuMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheMenu Delete CacheMenu
func DeleteCacheMenu(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheMenu))
	return row, err
}

// DeleteCacheMenuViaCid Delete CacheMenu via Cid
func DeleteCacheMenuViaCid(iCid string) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Cid: iCid}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaData Delete CacheMenu via Data
func DeleteCacheMenuViaData(iData []byte) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Data: iData}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaExpire Delete CacheMenu via Expire
func DeleteCacheMenuViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Expire: iExpire}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaCreated Delete CacheMenu via Created
func DeleteCacheMenuViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Created: iCreated}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaSerialized Delete CacheMenu via Serialized
func DeleteCacheMenuViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaTags Delete CacheMenu via Tags
func DeleteCacheMenuViaTags(iTags string) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Tags: iTags}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheMenuViaChecksum Delete CacheMenu via Checksum
func DeleteCacheMenuViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheMenu = &CacheMenu{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheMenu); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheMenu)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
