package model

type CacheContainer struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheContainersCount CacheContainers' Count
func GetCacheContainersCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheContainer{})
	return total, err
}

// GetCacheContainerCountViaCid Get CacheContainer via Cid
func GetCacheContainerCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheContainer{Cid: iCid})
	return n
}

// GetCacheContainerCountViaData Get CacheContainer via Data
func GetCacheContainerCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheContainer{Data: iData})
	return n
}

// GetCacheContainerCountViaExpire Get CacheContainer via Expire
func GetCacheContainerCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheContainer{Expire: iExpire})
	return n
}

// GetCacheContainerCountViaCreated Get CacheContainer via Created
func GetCacheContainerCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheContainer{Created: iCreated})
	return n
}

// GetCacheContainerCountViaSerialized Get CacheContainer via Serialized
func GetCacheContainerCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheContainer{Serialized: iSerialized})
	return n
}

// GetCacheContainerCountViaTags Get CacheContainer via Tags
func GetCacheContainerCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheContainer{Tags: iTags})
	return n
}

// GetCacheContainerCountViaChecksum Get CacheContainer via Checksum
func GetCacheContainerCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheContainer{Checksum: iChecksum})
	return n
}

// GetCacheContainersByCidAndDataAndExpire Get CacheContainers via CidAndDataAndExpire
func GetCacheContainersByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndDataAndCreated Get CacheContainers via CidAndDataAndCreated
func GetCacheContainersByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndDataAndSerialized Get CacheContainers via CidAndDataAndSerialized
func GetCacheContainersByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndDataAndTags Get CacheContainers via CidAndDataAndTags
func GetCacheContainersByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndDataAndChecksum Get CacheContainers via CidAndDataAndChecksum
func GetCacheContainersByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndExpireAndCreated Get CacheContainers via CidAndExpireAndCreated
func GetCacheContainersByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndExpireAndSerialized Get CacheContainers via CidAndExpireAndSerialized
func GetCacheContainersByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndExpireAndTags Get CacheContainers via CidAndExpireAndTags
func GetCacheContainersByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndExpireAndChecksum Get CacheContainers via CidAndExpireAndChecksum
func GetCacheContainersByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndCreatedAndSerialized Get CacheContainers via CidAndCreatedAndSerialized
func GetCacheContainersByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndCreatedAndTags Get CacheContainers via CidAndCreatedAndTags
func GetCacheContainersByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndCreatedAndChecksum Get CacheContainers via CidAndCreatedAndChecksum
func GetCacheContainersByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndSerializedAndTags Get CacheContainers via CidAndSerializedAndTags
func GetCacheContainersByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndSerializedAndChecksum Get CacheContainers via CidAndSerializedAndChecksum
func GetCacheContainersByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndTagsAndChecksum Get CacheContainers via CidAndTagsAndChecksum
func GetCacheContainersByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndExpireAndCreated Get CacheContainers via DataAndExpireAndCreated
func GetCacheContainersByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndExpireAndSerialized Get CacheContainers via DataAndExpireAndSerialized
func GetCacheContainersByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndExpireAndTags Get CacheContainers via DataAndExpireAndTags
func GetCacheContainersByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndExpireAndChecksum Get CacheContainers via DataAndExpireAndChecksum
func GetCacheContainersByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndCreatedAndSerialized Get CacheContainers via DataAndCreatedAndSerialized
func GetCacheContainersByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndCreatedAndTags Get CacheContainers via DataAndCreatedAndTags
func GetCacheContainersByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndCreatedAndChecksum Get CacheContainers via DataAndCreatedAndChecksum
func GetCacheContainersByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndSerializedAndTags Get CacheContainers via DataAndSerializedAndTags
func GetCacheContainersByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndSerializedAndChecksum Get CacheContainers via DataAndSerializedAndChecksum
func GetCacheContainersByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndTagsAndChecksum Get CacheContainers via DataAndTagsAndChecksum
func GetCacheContainersByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndCreatedAndSerialized Get CacheContainers via ExpireAndCreatedAndSerialized
func GetCacheContainersByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndCreatedAndTags Get CacheContainers via ExpireAndCreatedAndTags
func GetCacheContainersByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndCreatedAndChecksum Get CacheContainers via ExpireAndCreatedAndChecksum
func GetCacheContainersByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndSerializedAndTags Get CacheContainers via ExpireAndSerializedAndTags
func GetCacheContainersByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndSerializedAndChecksum Get CacheContainers via ExpireAndSerializedAndChecksum
func GetCacheContainersByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndTagsAndChecksum Get CacheContainers via ExpireAndTagsAndChecksum
func GetCacheContainersByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndSerializedAndTags Get CacheContainers via CreatedAndSerializedAndTags
func GetCacheContainersByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndSerializedAndChecksum Get CacheContainers via CreatedAndSerializedAndChecksum
func GetCacheContainersByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndTagsAndChecksum Get CacheContainers via CreatedAndTagsAndChecksum
func GetCacheContainersByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersBySerializedAndTagsAndChecksum Get CacheContainers via SerializedAndTagsAndChecksum
func GetCacheContainersBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndData Get CacheContainers via CidAndData
func GetCacheContainersByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndExpire Get CacheContainers via CidAndExpire
func GetCacheContainersByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndCreated Get CacheContainers via CidAndCreated
func GetCacheContainersByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndSerialized Get CacheContainers via CidAndSerialized
func GetCacheContainersByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndTags Get CacheContainers via CidAndTags
func GetCacheContainersByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCidAndChecksum Get CacheContainers via CidAndChecksum
func GetCacheContainersByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndExpire Get CacheContainers via DataAndExpire
func GetCacheContainersByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndCreated Get CacheContainers via DataAndCreated
func GetCacheContainersByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndSerialized Get CacheContainers via DataAndSerialized
func GetCacheContainersByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndTags Get CacheContainers via DataAndTags
func GetCacheContainersByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByDataAndChecksum Get CacheContainers via DataAndChecksum
func GetCacheContainersByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndCreated Get CacheContainers via ExpireAndCreated
func GetCacheContainersByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndSerialized Get CacheContainers via ExpireAndSerialized
func GetCacheContainersByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndTags Get CacheContainers via ExpireAndTags
func GetCacheContainersByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByExpireAndChecksum Get CacheContainers via ExpireAndChecksum
func GetCacheContainersByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndSerialized Get CacheContainers via CreatedAndSerialized
func GetCacheContainersByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndTags Get CacheContainers via CreatedAndTags
func GetCacheContainersByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByCreatedAndChecksum Get CacheContainers via CreatedAndChecksum
func GetCacheContainersByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersBySerializedAndTags Get CacheContainers via SerializedAndTags
func GetCacheContainersBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersBySerializedAndChecksum Get CacheContainers via SerializedAndChecksum
func GetCacheContainersBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersByTagsAndChecksum Get CacheContainers via TagsAndChecksum
func GetCacheContainersByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainers Get CacheContainers via field
func GetCacheContainers(offset int, limit int, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaCid Get CacheContainers via Cid
func GetCacheContainersViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaData Get CacheContainers via Data
func GetCacheContainersViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaExpire Get CacheContainers via Expire
func GetCacheContainersViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaCreated Get CacheContainers via Created
func GetCacheContainersViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaSerialized Get CacheContainers via Serialized
func GetCacheContainersViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaTags Get CacheContainers via Tags
func GetCacheContainersViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// GetCacheContainersViaChecksum Get CacheContainers via Checksum
func GetCacheContainersViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheContainer, error) {
	var _CacheContainer = new([]*CacheContainer)
	err := Engine.Table("cache_container").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheContainer)
	return _CacheContainer, err
}

// HasCacheContainerViaCid Has CacheContainer via Cid
func HasCacheContainerViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaData Has CacheContainer via Data
func HasCacheContainerViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaExpire Has CacheContainer via Expire
func HasCacheContainerViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaCreated Has CacheContainer via Created
func HasCacheContainerViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaSerialized Has CacheContainer via Serialized
func HasCacheContainerViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaTags Has CacheContainer via Tags
func HasCacheContainerViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheContainerViaChecksum Has CacheContainer via Checksum
func HasCacheContainerViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheContainer)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheContainerViaCid Get CacheContainer via Cid
func GetCacheContainerViaCid(iCid string) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Cid: iCid}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaData Get CacheContainer via Data
func GetCacheContainerViaData(iData []byte) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Data: iData}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaExpire Get CacheContainer via Expire
func GetCacheContainerViaExpire(iExpire int) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Expire: iExpire}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaCreated Get CacheContainer via Created
func GetCacheContainerViaCreated(iCreated string) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Created: iCreated}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaSerialized Get CacheContainer via Serialized
func GetCacheContainerViaSerialized(iSerialized int) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Serialized: iSerialized}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaTags Get CacheContainer via Tags
func GetCacheContainerViaTags(iTags string) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Tags: iTags}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// GetCacheContainerViaChecksum Get CacheContainer via Checksum
func GetCacheContainerViaChecksum(iChecksum string) (*CacheContainer, error) {
	var _CacheContainer = &CacheContainer{Checksum: iChecksum}
	has, err := Engine.Get(_CacheContainer)
	if has {
		return _CacheContainer, err
	} else {
		return nil, err
	}
}

// SetCacheContainerViaCid Set CacheContainer via Cid
func SetCacheContainerViaCid(iCid string, cache_container *CacheContainer) (int64, error) {
	cache_container.Cid = iCid
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaData Set CacheContainer via Data
func SetCacheContainerViaData(iData []byte, cache_container *CacheContainer) (int64, error) {
	cache_container.Data = iData
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaExpire Set CacheContainer via Expire
func SetCacheContainerViaExpire(iExpire int, cache_container *CacheContainer) (int64, error) {
	cache_container.Expire = iExpire
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaCreated Set CacheContainer via Created
func SetCacheContainerViaCreated(iCreated string, cache_container *CacheContainer) (int64, error) {
	cache_container.Created = iCreated
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaSerialized Set CacheContainer via Serialized
func SetCacheContainerViaSerialized(iSerialized int, cache_container *CacheContainer) (int64, error) {
	cache_container.Serialized = iSerialized
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaTags Set CacheContainer via Tags
func SetCacheContainerViaTags(iTags string, cache_container *CacheContainer) (int64, error) {
	cache_container.Tags = iTags
	return Engine.Insert(cache_container)
}

// SetCacheContainerViaChecksum Set CacheContainer via Checksum
func SetCacheContainerViaChecksum(iChecksum string, cache_container *CacheContainer) (int64, error) {
	cache_container.Checksum = iChecksum
	return Engine.Insert(cache_container)
}

// AddCacheContainer Add CacheContainer via all columns
func AddCacheContainer(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheContainer := &CacheContainer{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheContainer); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheContainer Post CacheContainer via iCacheContainer
func PostCacheContainer(iCacheContainer *CacheContainer) (string, error) {
	_, err := Engine.Insert(iCacheContainer)
	return iCacheContainer.Cid, err
}

// PutCacheContainer Put CacheContainer
func PutCacheContainer(iCacheContainer *CacheContainer) (string, error) {
	_, err := Engine.Id(iCacheContainer.Cid).Update(iCacheContainer)
	return iCacheContainer.Cid, err
}

// PutCacheContainerViaCid Put CacheContainer via Cid
func PutCacheContainerViaCid(Cid_ string, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Cid: Cid_})
	return row, err
}

// PutCacheContainerViaData Put CacheContainer via Data
func PutCacheContainerViaData(Data_ []byte, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Data: Data_})
	return row, err
}

// PutCacheContainerViaExpire Put CacheContainer via Expire
func PutCacheContainerViaExpire(Expire_ int, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Expire: Expire_})
	return row, err
}

// PutCacheContainerViaCreated Put CacheContainer via Created
func PutCacheContainerViaCreated(Created_ string, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Created: Created_})
	return row, err
}

// PutCacheContainerViaSerialized Put CacheContainer via Serialized
func PutCacheContainerViaSerialized(Serialized_ int, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Serialized: Serialized_})
	return row, err
}

// PutCacheContainerViaTags Put CacheContainer via Tags
func PutCacheContainerViaTags(Tags_ string, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Tags: Tags_})
	return row, err
}

// PutCacheContainerViaChecksum Put CacheContainer via Checksum
func PutCacheContainerViaChecksum(Checksum_ string, iCacheContainer *CacheContainer) (int64, error) {
	row, err := Engine.Update(iCacheContainer, &CacheContainer{Checksum: Checksum_})
	return row, err
}

// UpdateCacheContainerViaCid via map[string]interface{}{}
func UpdateCacheContainerViaCid(iCid string, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("cid = ?", iCid).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaData via map[string]interface{}{}
func UpdateCacheContainerViaData(iData []byte, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("data = ?", iData).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaExpire via map[string]interface{}{}
func UpdateCacheContainerViaExpire(iExpire int, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("expire = ?", iExpire).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaCreated via map[string]interface{}{}
func UpdateCacheContainerViaCreated(iCreated string, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("created = ?", iCreated).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaSerialized via map[string]interface{}{}
func UpdateCacheContainerViaSerialized(iSerialized int, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("serialized = ?", iSerialized).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaTags via map[string]interface{}{}
func UpdateCacheContainerViaTags(iTags string, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("tags = ?", iTags).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheContainerViaChecksum via map[string]interface{}{}
func UpdateCacheContainerViaChecksum(iChecksum string, iCacheContainerMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheContainer)).Where("checksum = ?", iChecksum).Update(iCacheContainerMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheContainer Delete CacheContainer
func DeleteCacheContainer(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheContainer))
	return row, err
}

// DeleteCacheContainerViaCid Delete CacheContainer via Cid
func DeleteCacheContainerViaCid(iCid string) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Cid: iCid}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaData Delete CacheContainer via Data
func DeleteCacheContainerViaData(iData []byte) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Data: iData}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaExpire Delete CacheContainer via Expire
func DeleteCacheContainerViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Expire: iExpire}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaCreated Delete CacheContainer via Created
func DeleteCacheContainerViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Created: iCreated}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaSerialized Delete CacheContainer via Serialized
func DeleteCacheContainerViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaTags Delete CacheContainer via Tags
func DeleteCacheContainerViaTags(iTags string) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Tags: iTags}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheContainerViaChecksum Delete CacheContainer via Checksum
func DeleteCacheContainerViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheContainer = &CacheContainer{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheContainer); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheContainer)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
