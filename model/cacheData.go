package model

type CacheData struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheDatasCount CacheDatas' Count
func GetCacheDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheData{})
	return total, err
}

// GetCacheDataCountViaCid Get CacheData via Cid
func GetCacheDataCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheData{Cid: iCid})
	return n
}

// GetCacheDataCountViaData Get CacheData via Data
func GetCacheDataCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheData{Data: iData})
	return n
}

// GetCacheDataCountViaExpire Get CacheData via Expire
func GetCacheDataCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheData{Expire: iExpire})
	return n
}

// GetCacheDataCountViaCreated Get CacheData via Created
func GetCacheDataCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheData{Created: iCreated})
	return n
}

// GetCacheDataCountViaSerialized Get CacheData via Serialized
func GetCacheDataCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheData{Serialized: iSerialized})
	return n
}

// GetCacheDataCountViaTags Get CacheData via Tags
func GetCacheDataCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheData{Tags: iTags})
	return n
}

// GetCacheDataCountViaChecksum Get CacheData via Checksum
func GetCacheDataCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheData{Checksum: iChecksum})
	return n
}

// GetCacheDatasByCidAndDataAndExpire Get CacheDatas via CidAndDataAndExpire
func GetCacheDatasByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndDataAndCreated Get CacheDatas via CidAndDataAndCreated
func GetCacheDatasByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndDataAndSerialized Get CacheDatas via CidAndDataAndSerialized
func GetCacheDatasByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndDataAndTags Get CacheDatas via CidAndDataAndTags
func GetCacheDatasByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndDataAndChecksum Get CacheDatas via CidAndDataAndChecksum
func GetCacheDatasByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndExpireAndCreated Get CacheDatas via CidAndExpireAndCreated
func GetCacheDatasByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndExpireAndSerialized Get CacheDatas via CidAndExpireAndSerialized
func GetCacheDatasByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndExpireAndTags Get CacheDatas via CidAndExpireAndTags
func GetCacheDatasByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndExpireAndChecksum Get CacheDatas via CidAndExpireAndChecksum
func GetCacheDatasByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndCreatedAndSerialized Get CacheDatas via CidAndCreatedAndSerialized
func GetCacheDatasByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndCreatedAndTags Get CacheDatas via CidAndCreatedAndTags
func GetCacheDatasByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndCreatedAndChecksum Get CacheDatas via CidAndCreatedAndChecksum
func GetCacheDatasByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndSerializedAndTags Get CacheDatas via CidAndSerializedAndTags
func GetCacheDatasByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndSerializedAndChecksum Get CacheDatas via CidAndSerializedAndChecksum
func GetCacheDatasByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndTagsAndChecksum Get CacheDatas via CidAndTagsAndChecksum
func GetCacheDatasByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndExpireAndCreated Get CacheDatas via DataAndExpireAndCreated
func GetCacheDatasByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndExpireAndSerialized Get CacheDatas via DataAndExpireAndSerialized
func GetCacheDatasByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndExpireAndTags Get CacheDatas via DataAndExpireAndTags
func GetCacheDatasByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndExpireAndChecksum Get CacheDatas via DataAndExpireAndChecksum
func GetCacheDatasByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndCreatedAndSerialized Get CacheDatas via DataAndCreatedAndSerialized
func GetCacheDatasByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndCreatedAndTags Get CacheDatas via DataAndCreatedAndTags
func GetCacheDatasByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndCreatedAndChecksum Get CacheDatas via DataAndCreatedAndChecksum
func GetCacheDatasByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndSerializedAndTags Get CacheDatas via DataAndSerializedAndTags
func GetCacheDatasByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndSerializedAndChecksum Get CacheDatas via DataAndSerializedAndChecksum
func GetCacheDatasByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndTagsAndChecksum Get CacheDatas via DataAndTagsAndChecksum
func GetCacheDatasByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndCreatedAndSerialized Get CacheDatas via ExpireAndCreatedAndSerialized
func GetCacheDatasByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndCreatedAndTags Get CacheDatas via ExpireAndCreatedAndTags
func GetCacheDatasByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndCreatedAndChecksum Get CacheDatas via ExpireAndCreatedAndChecksum
func GetCacheDatasByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndSerializedAndTags Get CacheDatas via ExpireAndSerializedAndTags
func GetCacheDatasByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndSerializedAndChecksum Get CacheDatas via ExpireAndSerializedAndChecksum
func GetCacheDatasByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndTagsAndChecksum Get CacheDatas via ExpireAndTagsAndChecksum
func GetCacheDatasByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndSerializedAndTags Get CacheDatas via CreatedAndSerializedAndTags
func GetCacheDatasByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndSerializedAndChecksum Get CacheDatas via CreatedAndSerializedAndChecksum
func GetCacheDatasByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndTagsAndChecksum Get CacheDatas via CreatedAndTagsAndChecksum
func GetCacheDatasByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasBySerializedAndTagsAndChecksum Get CacheDatas via SerializedAndTagsAndChecksum
func GetCacheDatasBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndData Get CacheDatas via CidAndData
func GetCacheDatasByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndExpire Get CacheDatas via CidAndExpire
func GetCacheDatasByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndCreated Get CacheDatas via CidAndCreated
func GetCacheDatasByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndSerialized Get CacheDatas via CidAndSerialized
func GetCacheDatasByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndTags Get CacheDatas via CidAndTags
func GetCacheDatasByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCidAndChecksum Get CacheDatas via CidAndChecksum
func GetCacheDatasByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndExpire Get CacheDatas via DataAndExpire
func GetCacheDatasByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndCreated Get CacheDatas via DataAndCreated
func GetCacheDatasByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndSerialized Get CacheDatas via DataAndSerialized
func GetCacheDatasByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndTags Get CacheDatas via DataAndTags
func GetCacheDatasByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByDataAndChecksum Get CacheDatas via DataAndChecksum
func GetCacheDatasByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndCreated Get CacheDatas via ExpireAndCreated
func GetCacheDatasByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndSerialized Get CacheDatas via ExpireAndSerialized
func GetCacheDatasByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndTags Get CacheDatas via ExpireAndTags
func GetCacheDatasByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByExpireAndChecksum Get CacheDatas via ExpireAndChecksum
func GetCacheDatasByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndSerialized Get CacheDatas via CreatedAndSerialized
func GetCacheDatasByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndTags Get CacheDatas via CreatedAndTags
func GetCacheDatasByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByCreatedAndChecksum Get CacheDatas via CreatedAndChecksum
func GetCacheDatasByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasBySerializedAndTags Get CacheDatas via SerializedAndTags
func GetCacheDatasBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasBySerializedAndChecksum Get CacheDatas via SerializedAndChecksum
func GetCacheDatasBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasByTagsAndChecksum Get CacheDatas via TagsAndChecksum
func GetCacheDatasByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatas Get CacheDatas via field
func GetCacheDatas(offset int, limit int, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaCid Get CacheDatas via Cid
func GetCacheDatasViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaData Get CacheDatas via Data
func GetCacheDatasViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaExpire Get CacheDatas via Expire
func GetCacheDatasViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaCreated Get CacheDatas via Created
func GetCacheDatasViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaSerialized Get CacheDatas via Serialized
func GetCacheDatasViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaTags Get CacheDatas via Tags
func GetCacheDatasViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// GetCacheDatasViaChecksum Get CacheDatas via Checksum
func GetCacheDatasViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheData, error) {
	var _CacheData = new([]*CacheData)
	err := Engine.Table("cache_data").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheData)
	return _CacheData, err
}

// HasCacheDataViaCid Has CacheData via Cid
func HasCacheDataViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaData Has CacheData via Data
func HasCacheDataViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaExpire Has CacheData via Expire
func HasCacheDataViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaCreated Has CacheData via Created
func HasCacheDataViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaSerialized Has CacheData via Serialized
func HasCacheDataViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaTags Has CacheData via Tags
func HasCacheDataViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDataViaChecksum Has CacheData via Checksum
func HasCacheDataViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheDataViaCid Get CacheData via Cid
func GetCacheDataViaCid(iCid string) (*CacheData, error) {
	var _CacheData = &CacheData{Cid: iCid}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaData Get CacheData via Data
func GetCacheDataViaData(iData []byte) (*CacheData, error) {
	var _CacheData = &CacheData{Data: iData}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaExpire Get CacheData via Expire
func GetCacheDataViaExpire(iExpire int) (*CacheData, error) {
	var _CacheData = &CacheData{Expire: iExpire}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaCreated Get CacheData via Created
func GetCacheDataViaCreated(iCreated string) (*CacheData, error) {
	var _CacheData = &CacheData{Created: iCreated}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaSerialized Get CacheData via Serialized
func GetCacheDataViaSerialized(iSerialized int) (*CacheData, error) {
	var _CacheData = &CacheData{Serialized: iSerialized}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaTags Get CacheData via Tags
func GetCacheDataViaTags(iTags string) (*CacheData, error) {
	var _CacheData = &CacheData{Tags: iTags}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// GetCacheDataViaChecksum Get CacheData via Checksum
func GetCacheDataViaChecksum(iChecksum string) (*CacheData, error) {
	var _CacheData = &CacheData{Checksum: iChecksum}
	has, err := Engine.Get(_CacheData)
	if has {
		return _CacheData, err
	} else {
		return nil, err
	}
}

// SetCacheDataViaCid Set CacheData via Cid
func SetCacheDataViaCid(iCid string, cache_data *CacheData) (int64, error) {
	cache_data.Cid = iCid
	return Engine.Insert(cache_data)
}

// SetCacheDataViaData Set CacheData via Data
func SetCacheDataViaData(iData []byte, cache_data *CacheData) (int64, error) {
	cache_data.Data = iData
	return Engine.Insert(cache_data)
}

// SetCacheDataViaExpire Set CacheData via Expire
func SetCacheDataViaExpire(iExpire int, cache_data *CacheData) (int64, error) {
	cache_data.Expire = iExpire
	return Engine.Insert(cache_data)
}

// SetCacheDataViaCreated Set CacheData via Created
func SetCacheDataViaCreated(iCreated string, cache_data *CacheData) (int64, error) {
	cache_data.Created = iCreated
	return Engine.Insert(cache_data)
}

// SetCacheDataViaSerialized Set CacheData via Serialized
func SetCacheDataViaSerialized(iSerialized int, cache_data *CacheData) (int64, error) {
	cache_data.Serialized = iSerialized
	return Engine.Insert(cache_data)
}

// SetCacheDataViaTags Set CacheData via Tags
func SetCacheDataViaTags(iTags string, cache_data *CacheData) (int64, error) {
	cache_data.Tags = iTags
	return Engine.Insert(cache_data)
}

// SetCacheDataViaChecksum Set CacheData via Checksum
func SetCacheDataViaChecksum(iChecksum string, cache_data *CacheData) (int64, error) {
	cache_data.Checksum = iChecksum
	return Engine.Insert(cache_data)
}

// AddCacheData Add CacheData via all columns
func AddCacheData(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheData := &CacheData{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheData Post CacheData via iCacheData
func PostCacheData(iCacheData *CacheData) (string, error) {
	_, err := Engine.Insert(iCacheData)
	return iCacheData.Cid, err
}

// PutCacheData Put CacheData
func PutCacheData(iCacheData *CacheData) (string, error) {
	_, err := Engine.Id(iCacheData.Cid).Update(iCacheData)
	return iCacheData.Cid, err
}

// PutCacheDataViaCid Put CacheData via Cid
func PutCacheDataViaCid(Cid_ string, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Cid: Cid_})
	return row, err
}

// PutCacheDataViaData Put CacheData via Data
func PutCacheDataViaData(Data_ []byte, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Data: Data_})
	return row, err
}

// PutCacheDataViaExpire Put CacheData via Expire
func PutCacheDataViaExpire(Expire_ int, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Expire: Expire_})
	return row, err
}

// PutCacheDataViaCreated Put CacheData via Created
func PutCacheDataViaCreated(Created_ string, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Created: Created_})
	return row, err
}

// PutCacheDataViaSerialized Put CacheData via Serialized
func PutCacheDataViaSerialized(Serialized_ int, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Serialized: Serialized_})
	return row, err
}

// PutCacheDataViaTags Put CacheData via Tags
func PutCacheDataViaTags(Tags_ string, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Tags: Tags_})
	return row, err
}

// PutCacheDataViaChecksum Put CacheData via Checksum
func PutCacheDataViaChecksum(Checksum_ string, iCacheData *CacheData) (int64, error) {
	row, err := Engine.Update(iCacheData, &CacheData{Checksum: Checksum_})
	return row, err
}

// UpdateCacheDataViaCid via map[string]interface{}{}
func UpdateCacheDataViaCid(iCid string, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("cid = ?", iCid).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaData via map[string]interface{}{}
func UpdateCacheDataViaData(iData []byte, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("data = ?", iData).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaExpire via map[string]interface{}{}
func UpdateCacheDataViaExpire(iExpire int, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("expire = ?", iExpire).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaCreated via map[string]interface{}{}
func UpdateCacheDataViaCreated(iCreated string, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("created = ?", iCreated).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaSerialized via map[string]interface{}{}
func UpdateCacheDataViaSerialized(iSerialized int, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("serialized = ?", iSerialized).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaTags via map[string]interface{}{}
func UpdateCacheDataViaTags(iTags string, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("tags = ?", iTags).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDataViaChecksum via map[string]interface{}{}
func UpdateCacheDataViaChecksum(iChecksum string, iCacheDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheData)).Where("checksum = ?", iChecksum).Update(iCacheDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheData Delete CacheData
func DeleteCacheData(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheData))
	return row, err
}

// DeleteCacheDataViaCid Delete CacheData via Cid
func DeleteCacheDataViaCid(iCid string) (err error) {
	var has bool
	var _CacheData = &CacheData{Cid: iCid}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaData Delete CacheData via Data
func DeleteCacheDataViaData(iData []byte) (err error) {
	var has bool
	var _CacheData = &CacheData{Data: iData}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaExpire Delete CacheData via Expire
func DeleteCacheDataViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheData = &CacheData{Expire: iExpire}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaCreated Delete CacheData via Created
func DeleteCacheDataViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheData = &CacheData{Created: iCreated}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaSerialized Delete CacheData via Serialized
func DeleteCacheDataViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheData = &CacheData{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaTags Delete CacheData via Tags
func DeleteCacheDataViaTags(iTags string) (err error) {
	var has bool
	var _CacheData = &CacheData{Tags: iTags}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDataViaChecksum Delete CacheData via Checksum
func DeleteCacheDataViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheData = &CacheData{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheData); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
