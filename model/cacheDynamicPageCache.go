package model

type CacheDynamicPageCache struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheDynamicPageCachesCount CacheDynamicPageCaches' Count
func GetCacheDynamicPageCachesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheDynamicPageCache{})
	return total, err
}

// GetCacheDynamicPageCacheCountViaCid Get CacheDynamicPageCache via Cid
func GetCacheDynamicPageCacheCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheDynamicPageCache{Cid: iCid})
	return n
}

// GetCacheDynamicPageCacheCountViaData Get CacheDynamicPageCache via Data
func GetCacheDynamicPageCacheCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheDynamicPageCache{Data: iData})
	return n
}

// GetCacheDynamicPageCacheCountViaExpire Get CacheDynamicPageCache via Expire
func GetCacheDynamicPageCacheCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheDynamicPageCache{Expire: iExpire})
	return n
}

// GetCacheDynamicPageCacheCountViaCreated Get CacheDynamicPageCache via Created
func GetCacheDynamicPageCacheCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheDynamicPageCache{Created: iCreated})
	return n
}

// GetCacheDynamicPageCacheCountViaSerialized Get CacheDynamicPageCache via Serialized
func GetCacheDynamicPageCacheCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheDynamicPageCache{Serialized: iSerialized})
	return n
}

// GetCacheDynamicPageCacheCountViaTags Get CacheDynamicPageCache via Tags
func GetCacheDynamicPageCacheCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheDynamicPageCache{Tags: iTags})
	return n
}

// GetCacheDynamicPageCacheCountViaChecksum Get CacheDynamicPageCache via Checksum
func GetCacheDynamicPageCacheCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheDynamicPageCache{Checksum: iChecksum})
	return n
}

// GetCacheDynamicPageCachesByCidAndDataAndExpire Get CacheDynamicPageCaches via CidAndDataAndExpire
func GetCacheDynamicPageCachesByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndDataAndCreated Get CacheDynamicPageCaches via CidAndDataAndCreated
func GetCacheDynamicPageCachesByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndDataAndSerialized Get CacheDynamicPageCaches via CidAndDataAndSerialized
func GetCacheDynamicPageCachesByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndDataAndTags Get CacheDynamicPageCaches via CidAndDataAndTags
func GetCacheDynamicPageCachesByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndDataAndChecksum Get CacheDynamicPageCaches via CidAndDataAndChecksum
func GetCacheDynamicPageCachesByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndExpireAndCreated Get CacheDynamicPageCaches via CidAndExpireAndCreated
func GetCacheDynamicPageCachesByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndExpireAndSerialized Get CacheDynamicPageCaches via CidAndExpireAndSerialized
func GetCacheDynamicPageCachesByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndExpireAndTags Get CacheDynamicPageCaches via CidAndExpireAndTags
func GetCacheDynamicPageCachesByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndExpireAndChecksum Get CacheDynamicPageCaches via CidAndExpireAndChecksum
func GetCacheDynamicPageCachesByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndCreatedAndSerialized Get CacheDynamicPageCaches via CidAndCreatedAndSerialized
func GetCacheDynamicPageCachesByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndCreatedAndTags Get CacheDynamicPageCaches via CidAndCreatedAndTags
func GetCacheDynamicPageCachesByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndCreatedAndChecksum Get CacheDynamicPageCaches via CidAndCreatedAndChecksum
func GetCacheDynamicPageCachesByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndSerializedAndTags Get CacheDynamicPageCaches via CidAndSerializedAndTags
func GetCacheDynamicPageCachesByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndSerializedAndChecksum Get CacheDynamicPageCaches via CidAndSerializedAndChecksum
func GetCacheDynamicPageCachesByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndTagsAndChecksum Get CacheDynamicPageCaches via CidAndTagsAndChecksum
func GetCacheDynamicPageCachesByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndExpireAndCreated Get CacheDynamicPageCaches via DataAndExpireAndCreated
func GetCacheDynamicPageCachesByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndExpireAndSerialized Get CacheDynamicPageCaches via DataAndExpireAndSerialized
func GetCacheDynamicPageCachesByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndExpireAndTags Get CacheDynamicPageCaches via DataAndExpireAndTags
func GetCacheDynamicPageCachesByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndExpireAndChecksum Get CacheDynamicPageCaches via DataAndExpireAndChecksum
func GetCacheDynamicPageCachesByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndCreatedAndSerialized Get CacheDynamicPageCaches via DataAndCreatedAndSerialized
func GetCacheDynamicPageCachesByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndCreatedAndTags Get CacheDynamicPageCaches via DataAndCreatedAndTags
func GetCacheDynamicPageCachesByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndCreatedAndChecksum Get CacheDynamicPageCaches via DataAndCreatedAndChecksum
func GetCacheDynamicPageCachesByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndSerializedAndTags Get CacheDynamicPageCaches via DataAndSerializedAndTags
func GetCacheDynamicPageCachesByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndSerializedAndChecksum Get CacheDynamicPageCaches via DataAndSerializedAndChecksum
func GetCacheDynamicPageCachesByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndTagsAndChecksum Get CacheDynamicPageCaches via DataAndTagsAndChecksum
func GetCacheDynamicPageCachesByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndCreatedAndSerialized Get CacheDynamicPageCaches via ExpireAndCreatedAndSerialized
func GetCacheDynamicPageCachesByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndCreatedAndTags Get CacheDynamicPageCaches via ExpireAndCreatedAndTags
func GetCacheDynamicPageCachesByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndCreatedAndChecksum Get CacheDynamicPageCaches via ExpireAndCreatedAndChecksum
func GetCacheDynamicPageCachesByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndSerializedAndTags Get CacheDynamicPageCaches via ExpireAndSerializedAndTags
func GetCacheDynamicPageCachesByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndSerializedAndChecksum Get CacheDynamicPageCaches via ExpireAndSerializedAndChecksum
func GetCacheDynamicPageCachesByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndTagsAndChecksum Get CacheDynamicPageCaches via ExpireAndTagsAndChecksum
func GetCacheDynamicPageCachesByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndSerializedAndTags Get CacheDynamicPageCaches via CreatedAndSerializedAndTags
func GetCacheDynamicPageCachesByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndSerializedAndChecksum Get CacheDynamicPageCaches via CreatedAndSerializedAndChecksum
func GetCacheDynamicPageCachesByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndTagsAndChecksum Get CacheDynamicPageCaches via CreatedAndTagsAndChecksum
func GetCacheDynamicPageCachesByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesBySerializedAndTagsAndChecksum Get CacheDynamicPageCaches via SerializedAndTagsAndChecksum
func GetCacheDynamicPageCachesBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndData Get CacheDynamicPageCaches via CidAndData
func GetCacheDynamicPageCachesByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndExpire Get CacheDynamicPageCaches via CidAndExpire
func GetCacheDynamicPageCachesByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndCreated Get CacheDynamicPageCaches via CidAndCreated
func GetCacheDynamicPageCachesByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndSerialized Get CacheDynamicPageCaches via CidAndSerialized
func GetCacheDynamicPageCachesByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndTags Get CacheDynamicPageCaches via CidAndTags
func GetCacheDynamicPageCachesByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCidAndChecksum Get CacheDynamicPageCaches via CidAndChecksum
func GetCacheDynamicPageCachesByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndExpire Get CacheDynamicPageCaches via DataAndExpire
func GetCacheDynamicPageCachesByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndCreated Get CacheDynamicPageCaches via DataAndCreated
func GetCacheDynamicPageCachesByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndSerialized Get CacheDynamicPageCaches via DataAndSerialized
func GetCacheDynamicPageCachesByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndTags Get CacheDynamicPageCaches via DataAndTags
func GetCacheDynamicPageCachesByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByDataAndChecksum Get CacheDynamicPageCaches via DataAndChecksum
func GetCacheDynamicPageCachesByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndCreated Get CacheDynamicPageCaches via ExpireAndCreated
func GetCacheDynamicPageCachesByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndSerialized Get CacheDynamicPageCaches via ExpireAndSerialized
func GetCacheDynamicPageCachesByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndTags Get CacheDynamicPageCaches via ExpireAndTags
func GetCacheDynamicPageCachesByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByExpireAndChecksum Get CacheDynamicPageCaches via ExpireAndChecksum
func GetCacheDynamicPageCachesByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndSerialized Get CacheDynamicPageCaches via CreatedAndSerialized
func GetCacheDynamicPageCachesByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndTags Get CacheDynamicPageCaches via CreatedAndTags
func GetCacheDynamicPageCachesByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByCreatedAndChecksum Get CacheDynamicPageCaches via CreatedAndChecksum
func GetCacheDynamicPageCachesByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesBySerializedAndTags Get CacheDynamicPageCaches via SerializedAndTags
func GetCacheDynamicPageCachesBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesBySerializedAndChecksum Get CacheDynamicPageCaches via SerializedAndChecksum
func GetCacheDynamicPageCachesBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesByTagsAndChecksum Get CacheDynamicPageCaches via TagsAndChecksum
func GetCacheDynamicPageCachesByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCaches Get CacheDynamicPageCaches via field
func GetCacheDynamicPageCaches(offset int, limit int, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaCid Get CacheDynamicPageCaches via Cid
func GetCacheDynamicPageCachesViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaData Get CacheDynamicPageCaches via Data
func GetCacheDynamicPageCachesViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaExpire Get CacheDynamicPageCaches via Expire
func GetCacheDynamicPageCachesViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaCreated Get CacheDynamicPageCaches via Created
func GetCacheDynamicPageCachesViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaSerialized Get CacheDynamicPageCaches via Serialized
func GetCacheDynamicPageCachesViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaTags Get CacheDynamicPageCaches via Tags
func GetCacheDynamicPageCachesViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// GetCacheDynamicPageCachesViaChecksum Get CacheDynamicPageCaches via Checksum
func GetCacheDynamicPageCachesViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = new([]*CacheDynamicPageCache)
	err := Engine.Table("cache_dynamic_page_cache").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheDynamicPageCache)
	return _CacheDynamicPageCache, err
}

// HasCacheDynamicPageCacheViaCid Has CacheDynamicPageCache via Cid
func HasCacheDynamicPageCacheViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaData Has CacheDynamicPageCache via Data
func HasCacheDynamicPageCacheViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaExpire Has CacheDynamicPageCache via Expire
func HasCacheDynamicPageCacheViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaCreated Has CacheDynamicPageCache via Created
func HasCacheDynamicPageCacheViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaSerialized Has CacheDynamicPageCache via Serialized
func HasCacheDynamicPageCacheViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaTags Has CacheDynamicPageCache via Tags
func HasCacheDynamicPageCacheViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDynamicPageCacheViaChecksum Has CacheDynamicPageCache via Checksum
func HasCacheDynamicPageCacheViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheDynamicPageCache)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheDynamicPageCacheViaCid Get CacheDynamicPageCache via Cid
func GetCacheDynamicPageCacheViaCid(iCid string) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Cid: iCid}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaData Get CacheDynamicPageCache via Data
func GetCacheDynamicPageCacheViaData(iData []byte) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Data: iData}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaExpire Get CacheDynamicPageCache via Expire
func GetCacheDynamicPageCacheViaExpire(iExpire int) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Expire: iExpire}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaCreated Get CacheDynamicPageCache via Created
func GetCacheDynamicPageCacheViaCreated(iCreated string) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Created: iCreated}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaSerialized Get CacheDynamicPageCache via Serialized
func GetCacheDynamicPageCacheViaSerialized(iSerialized int) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Serialized: iSerialized}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaTags Get CacheDynamicPageCache via Tags
func GetCacheDynamicPageCacheViaTags(iTags string) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Tags: iTags}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// GetCacheDynamicPageCacheViaChecksum Get CacheDynamicPageCache via Checksum
func GetCacheDynamicPageCacheViaChecksum(iChecksum string) (*CacheDynamicPageCache, error) {
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Checksum: iChecksum}
	has, err := Engine.Get(_CacheDynamicPageCache)
	if has {
		return _CacheDynamicPageCache, err
	} else {
		return nil, err
	}
}

// SetCacheDynamicPageCacheViaCid Set CacheDynamicPageCache via Cid
func SetCacheDynamicPageCacheViaCid(iCid string, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Cid = iCid
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaData Set CacheDynamicPageCache via Data
func SetCacheDynamicPageCacheViaData(iData []byte, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Data = iData
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaExpire Set CacheDynamicPageCache via Expire
func SetCacheDynamicPageCacheViaExpire(iExpire int, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Expire = iExpire
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaCreated Set CacheDynamicPageCache via Created
func SetCacheDynamicPageCacheViaCreated(iCreated string, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Created = iCreated
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaSerialized Set CacheDynamicPageCache via Serialized
func SetCacheDynamicPageCacheViaSerialized(iSerialized int, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Serialized = iSerialized
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaTags Set CacheDynamicPageCache via Tags
func SetCacheDynamicPageCacheViaTags(iTags string, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Tags = iTags
	return Engine.Insert(cache_dynamic_page_cache)
}

// SetCacheDynamicPageCacheViaChecksum Set CacheDynamicPageCache via Checksum
func SetCacheDynamicPageCacheViaChecksum(iChecksum string, cache_dynamic_page_cache *CacheDynamicPageCache) (int64, error) {
	cache_dynamic_page_cache.Checksum = iChecksum
	return Engine.Insert(cache_dynamic_page_cache)
}

// AddCacheDynamicPageCache Add CacheDynamicPageCache via all columns
func AddCacheDynamicPageCache(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheDynamicPageCache := &CacheDynamicPageCache{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheDynamicPageCache); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheDynamicPageCache Post CacheDynamicPageCache via iCacheDynamicPageCache
func PostCacheDynamicPageCache(iCacheDynamicPageCache *CacheDynamicPageCache) (string, error) {
	_, err := Engine.Insert(iCacheDynamicPageCache)
	return iCacheDynamicPageCache.Cid, err
}

// PutCacheDynamicPageCache Put CacheDynamicPageCache
func PutCacheDynamicPageCache(iCacheDynamicPageCache *CacheDynamicPageCache) (string, error) {
	_, err := Engine.Id(iCacheDynamicPageCache.Cid).Update(iCacheDynamicPageCache)
	return iCacheDynamicPageCache.Cid, err
}

// PutCacheDynamicPageCacheViaCid Put CacheDynamicPageCache via Cid
func PutCacheDynamicPageCacheViaCid(Cid_ string, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Cid: Cid_})
	return row, err
}

// PutCacheDynamicPageCacheViaData Put CacheDynamicPageCache via Data
func PutCacheDynamicPageCacheViaData(Data_ []byte, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Data: Data_})
	return row, err
}

// PutCacheDynamicPageCacheViaExpire Put CacheDynamicPageCache via Expire
func PutCacheDynamicPageCacheViaExpire(Expire_ int, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Expire: Expire_})
	return row, err
}

// PutCacheDynamicPageCacheViaCreated Put CacheDynamicPageCache via Created
func PutCacheDynamicPageCacheViaCreated(Created_ string, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Created: Created_})
	return row, err
}

// PutCacheDynamicPageCacheViaSerialized Put CacheDynamicPageCache via Serialized
func PutCacheDynamicPageCacheViaSerialized(Serialized_ int, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Serialized: Serialized_})
	return row, err
}

// PutCacheDynamicPageCacheViaTags Put CacheDynamicPageCache via Tags
func PutCacheDynamicPageCacheViaTags(Tags_ string, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Tags: Tags_})
	return row, err
}

// PutCacheDynamicPageCacheViaChecksum Put CacheDynamicPageCache via Checksum
func PutCacheDynamicPageCacheViaChecksum(Checksum_ string, iCacheDynamicPageCache *CacheDynamicPageCache) (int64, error) {
	row, err := Engine.Update(iCacheDynamicPageCache, &CacheDynamicPageCache{Checksum: Checksum_})
	return row, err
}

// UpdateCacheDynamicPageCacheViaCid via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaCid(iCid string, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("cid = ?", iCid).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaData via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaData(iData []byte, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("data = ?", iData).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaExpire via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaExpire(iExpire int, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("expire = ?", iExpire).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaCreated via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaCreated(iCreated string, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("created = ?", iCreated).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaSerialized via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaSerialized(iSerialized int, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("serialized = ?", iSerialized).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaTags via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaTags(iTags string, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("tags = ?", iTags).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDynamicPageCacheViaChecksum via map[string]interface{}{}
func UpdateCacheDynamicPageCacheViaChecksum(iChecksum string, iCacheDynamicPageCacheMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDynamicPageCache)).Where("checksum = ?", iChecksum).Update(iCacheDynamicPageCacheMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheDynamicPageCache Delete CacheDynamicPageCache
func DeleteCacheDynamicPageCache(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheDynamicPageCache))
	return row, err
}

// DeleteCacheDynamicPageCacheViaCid Delete CacheDynamicPageCache via Cid
func DeleteCacheDynamicPageCacheViaCid(iCid string) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Cid: iCid}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaData Delete CacheDynamicPageCache via Data
func DeleteCacheDynamicPageCacheViaData(iData []byte) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Data: iData}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaExpire Delete CacheDynamicPageCache via Expire
func DeleteCacheDynamicPageCacheViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Expire: iExpire}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaCreated Delete CacheDynamicPageCache via Created
func DeleteCacheDynamicPageCacheViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Created: iCreated}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaSerialized Delete CacheDynamicPageCache via Serialized
func DeleteCacheDynamicPageCacheViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaTags Delete CacheDynamicPageCache via Tags
func DeleteCacheDynamicPageCacheViaTags(iTags string) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Tags: iTags}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDynamicPageCacheViaChecksum Delete CacheDynamicPageCache via Checksum
func DeleteCacheDynamicPageCacheViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheDynamicPageCache = &CacheDynamicPageCache{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheDynamicPageCache); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheDynamicPageCache)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
