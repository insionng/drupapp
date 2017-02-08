package model

type CacheEntity struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheEntitiesCount CacheEntitys' Count
func GetCacheEntitiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheEntity{})
	return total, err
}

// GetCacheEntityCountViaCid Get CacheEntity via Cid
func GetCacheEntityCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheEntity{Cid: iCid})
	return n
}

// GetCacheEntityCountViaData Get CacheEntity via Data
func GetCacheEntityCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheEntity{Data: iData})
	return n
}

// GetCacheEntityCountViaExpire Get CacheEntity via Expire
func GetCacheEntityCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheEntity{Expire: iExpire})
	return n
}

// GetCacheEntityCountViaCreated Get CacheEntity via Created
func GetCacheEntityCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheEntity{Created: iCreated})
	return n
}

// GetCacheEntityCountViaSerialized Get CacheEntity via Serialized
func GetCacheEntityCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheEntity{Serialized: iSerialized})
	return n
}

// GetCacheEntityCountViaTags Get CacheEntity via Tags
func GetCacheEntityCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheEntity{Tags: iTags})
	return n
}

// GetCacheEntityCountViaChecksum Get CacheEntity via Checksum
func GetCacheEntityCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheEntity{Checksum: iChecksum})
	return n
}

// GetCacheEntitiesByCidAndDataAndExpire Get CacheEntities via CidAndDataAndExpire
func GetCacheEntitiesByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndDataAndCreated Get CacheEntities via CidAndDataAndCreated
func GetCacheEntitiesByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndDataAndSerialized Get CacheEntities via CidAndDataAndSerialized
func GetCacheEntitiesByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndDataAndTags Get CacheEntities via CidAndDataAndTags
func GetCacheEntitiesByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndDataAndChecksum Get CacheEntities via CidAndDataAndChecksum
func GetCacheEntitiesByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndExpireAndCreated Get CacheEntities via CidAndExpireAndCreated
func GetCacheEntitiesByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndExpireAndSerialized Get CacheEntities via CidAndExpireAndSerialized
func GetCacheEntitiesByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndExpireAndTags Get CacheEntities via CidAndExpireAndTags
func GetCacheEntitiesByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndExpireAndChecksum Get CacheEntities via CidAndExpireAndChecksum
func GetCacheEntitiesByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndCreatedAndSerialized Get CacheEntities via CidAndCreatedAndSerialized
func GetCacheEntitiesByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndCreatedAndTags Get CacheEntities via CidAndCreatedAndTags
func GetCacheEntitiesByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndCreatedAndChecksum Get CacheEntities via CidAndCreatedAndChecksum
func GetCacheEntitiesByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndSerializedAndTags Get CacheEntities via CidAndSerializedAndTags
func GetCacheEntitiesByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndSerializedAndChecksum Get CacheEntities via CidAndSerializedAndChecksum
func GetCacheEntitiesByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndTagsAndChecksum Get CacheEntities via CidAndTagsAndChecksum
func GetCacheEntitiesByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndExpireAndCreated Get CacheEntities via DataAndExpireAndCreated
func GetCacheEntitiesByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndExpireAndSerialized Get CacheEntities via DataAndExpireAndSerialized
func GetCacheEntitiesByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndExpireAndTags Get CacheEntities via DataAndExpireAndTags
func GetCacheEntitiesByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndExpireAndChecksum Get CacheEntities via DataAndExpireAndChecksum
func GetCacheEntitiesByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndCreatedAndSerialized Get CacheEntities via DataAndCreatedAndSerialized
func GetCacheEntitiesByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndCreatedAndTags Get CacheEntities via DataAndCreatedAndTags
func GetCacheEntitiesByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndCreatedAndChecksum Get CacheEntities via DataAndCreatedAndChecksum
func GetCacheEntitiesByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndSerializedAndTags Get CacheEntities via DataAndSerializedAndTags
func GetCacheEntitiesByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndSerializedAndChecksum Get CacheEntities via DataAndSerializedAndChecksum
func GetCacheEntitiesByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndTagsAndChecksum Get CacheEntities via DataAndTagsAndChecksum
func GetCacheEntitiesByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndCreatedAndSerialized Get CacheEntities via ExpireAndCreatedAndSerialized
func GetCacheEntitiesByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndCreatedAndTags Get CacheEntities via ExpireAndCreatedAndTags
func GetCacheEntitiesByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndCreatedAndChecksum Get CacheEntities via ExpireAndCreatedAndChecksum
func GetCacheEntitiesByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndSerializedAndTags Get CacheEntities via ExpireAndSerializedAndTags
func GetCacheEntitiesByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndSerializedAndChecksum Get CacheEntities via ExpireAndSerializedAndChecksum
func GetCacheEntitiesByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndTagsAndChecksum Get CacheEntities via ExpireAndTagsAndChecksum
func GetCacheEntitiesByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndSerializedAndTags Get CacheEntities via CreatedAndSerializedAndTags
func GetCacheEntitiesByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndSerializedAndChecksum Get CacheEntities via CreatedAndSerializedAndChecksum
func GetCacheEntitiesByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndTagsAndChecksum Get CacheEntities via CreatedAndTagsAndChecksum
func GetCacheEntitiesByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesBySerializedAndTagsAndChecksum Get CacheEntities via SerializedAndTagsAndChecksum
func GetCacheEntitiesBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndData Get CacheEntities via CidAndData
func GetCacheEntitiesByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndExpire Get CacheEntities via CidAndExpire
func GetCacheEntitiesByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndCreated Get CacheEntities via CidAndCreated
func GetCacheEntitiesByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndSerialized Get CacheEntities via CidAndSerialized
func GetCacheEntitiesByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndTags Get CacheEntities via CidAndTags
func GetCacheEntitiesByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCidAndChecksum Get CacheEntities via CidAndChecksum
func GetCacheEntitiesByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndExpire Get CacheEntities via DataAndExpire
func GetCacheEntitiesByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndCreated Get CacheEntities via DataAndCreated
func GetCacheEntitiesByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndSerialized Get CacheEntities via DataAndSerialized
func GetCacheEntitiesByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndTags Get CacheEntities via DataAndTags
func GetCacheEntitiesByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByDataAndChecksum Get CacheEntities via DataAndChecksum
func GetCacheEntitiesByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndCreated Get CacheEntities via ExpireAndCreated
func GetCacheEntitiesByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndSerialized Get CacheEntities via ExpireAndSerialized
func GetCacheEntitiesByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndTags Get CacheEntities via ExpireAndTags
func GetCacheEntitiesByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByExpireAndChecksum Get CacheEntities via ExpireAndChecksum
func GetCacheEntitiesByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndSerialized Get CacheEntities via CreatedAndSerialized
func GetCacheEntitiesByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndTags Get CacheEntities via CreatedAndTags
func GetCacheEntitiesByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByCreatedAndChecksum Get CacheEntities via CreatedAndChecksum
func GetCacheEntitiesByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesBySerializedAndTags Get CacheEntities via SerializedAndTags
func GetCacheEntitiesBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesBySerializedAndChecksum Get CacheEntities via SerializedAndChecksum
func GetCacheEntitiesBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesByTagsAndChecksum Get CacheEntities via TagsAndChecksum
func GetCacheEntitiesByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntities Get CacheEntities via field
func GetCacheEntities(offset int, limit int, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaCid Get CacheEntitys via Cid
func GetCacheEntitiesViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaData Get CacheEntitys via Data
func GetCacheEntitiesViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaExpire Get CacheEntitys via Expire
func GetCacheEntitiesViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaCreated Get CacheEntitys via Created
func GetCacheEntitiesViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaSerialized Get CacheEntitys via Serialized
func GetCacheEntitiesViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaTags Get CacheEntitys via Tags
func GetCacheEntitiesViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// GetCacheEntitiesViaChecksum Get CacheEntitys via Checksum
func GetCacheEntitiesViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheEntity, error) {
	var _CacheEntity = new([]*CacheEntity)
	err := Engine.Table("cache_entity").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheEntity)
	return _CacheEntity, err
}

// HasCacheEntityViaCid Has CacheEntity via Cid
func HasCacheEntityViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaData Has CacheEntity via Data
func HasCacheEntityViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaExpire Has CacheEntity via Expire
func HasCacheEntityViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaCreated Has CacheEntity via Created
func HasCacheEntityViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaSerialized Has CacheEntity via Serialized
func HasCacheEntityViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaTags Has CacheEntity via Tags
func HasCacheEntityViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheEntityViaChecksum Has CacheEntity via Checksum
func HasCacheEntityViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheEntity)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheEntityViaCid Get CacheEntity via Cid
func GetCacheEntityViaCid(iCid string) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Cid: iCid}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaData Get CacheEntity via Data
func GetCacheEntityViaData(iData []byte) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Data: iData}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaExpire Get CacheEntity via Expire
func GetCacheEntityViaExpire(iExpire int) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Expire: iExpire}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaCreated Get CacheEntity via Created
func GetCacheEntityViaCreated(iCreated string) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Created: iCreated}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaSerialized Get CacheEntity via Serialized
func GetCacheEntityViaSerialized(iSerialized int) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Serialized: iSerialized}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaTags Get CacheEntity via Tags
func GetCacheEntityViaTags(iTags string) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Tags: iTags}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// GetCacheEntityViaChecksum Get CacheEntity via Checksum
func GetCacheEntityViaChecksum(iChecksum string) (*CacheEntity, error) {
	var _CacheEntity = &CacheEntity{Checksum: iChecksum}
	has, err := Engine.Get(_CacheEntity)
	if has {
		return _CacheEntity, err
	} else {
		return nil, err
	}
}

// SetCacheEntityViaCid Set CacheEntity via Cid
func SetCacheEntityViaCid(iCid string, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Cid = iCid
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaData Set CacheEntity via Data
func SetCacheEntityViaData(iData []byte, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Data = iData
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaExpire Set CacheEntity via Expire
func SetCacheEntityViaExpire(iExpire int, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Expire = iExpire
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaCreated Set CacheEntity via Created
func SetCacheEntityViaCreated(iCreated string, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Created = iCreated
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaSerialized Set CacheEntity via Serialized
func SetCacheEntityViaSerialized(iSerialized int, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Serialized = iSerialized
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaTags Set CacheEntity via Tags
func SetCacheEntityViaTags(iTags string, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Tags = iTags
	return Engine.Insert(cache_entity)
}

// SetCacheEntityViaChecksum Set CacheEntity via Checksum
func SetCacheEntityViaChecksum(iChecksum string, cache_entity *CacheEntity) (int64, error) {
	cache_entity.Checksum = iChecksum
	return Engine.Insert(cache_entity)
}

// AddCacheEntity Add CacheEntity via all columns
func AddCacheEntity(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheEntity := &CacheEntity{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheEntity); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheEntity Post CacheEntity via iCacheEntity
func PostCacheEntity(iCacheEntity *CacheEntity) (string, error) {
	_, err := Engine.Insert(iCacheEntity)
	return iCacheEntity.Cid, err
}

// PutCacheEntity Put CacheEntity
func PutCacheEntity(iCacheEntity *CacheEntity) (string, error) {
	_, err := Engine.Id(iCacheEntity.Cid).Update(iCacheEntity)
	return iCacheEntity.Cid, err
}

// PutCacheEntityViaCid Put CacheEntity via Cid
func PutCacheEntityViaCid(Cid_ string, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Cid: Cid_})
	return row, err
}

// PutCacheEntityViaData Put CacheEntity via Data
func PutCacheEntityViaData(Data_ []byte, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Data: Data_})
	return row, err
}

// PutCacheEntityViaExpire Put CacheEntity via Expire
func PutCacheEntityViaExpire(Expire_ int, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Expire: Expire_})
	return row, err
}

// PutCacheEntityViaCreated Put CacheEntity via Created
func PutCacheEntityViaCreated(Created_ string, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Created: Created_})
	return row, err
}

// PutCacheEntityViaSerialized Put CacheEntity via Serialized
func PutCacheEntityViaSerialized(Serialized_ int, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Serialized: Serialized_})
	return row, err
}

// PutCacheEntityViaTags Put CacheEntity via Tags
func PutCacheEntityViaTags(Tags_ string, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Tags: Tags_})
	return row, err
}

// PutCacheEntityViaChecksum Put CacheEntity via Checksum
func PutCacheEntityViaChecksum(Checksum_ string, iCacheEntity *CacheEntity) (int64, error) {
	row, err := Engine.Update(iCacheEntity, &CacheEntity{Checksum: Checksum_})
	return row, err
}

// UpdateCacheEntityViaCid via map[string]interface{}{}
func UpdateCacheEntityViaCid(iCid string, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("cid = ?", iCid).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaData via map[string]interface{}{}
func UpdateCacheEntityViaData(iData []byte, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("data = ?", iData).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaExpire via map[string]interface{}{}
func UpdateCacheEntityViaExpire(iExpire int, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("expire = ?", iExpire).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaCreated via map[string]interface{}{}
func UpdateCacheEntityViaCreated(iCreated string, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("created = ?", iCreated).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaSerialized via map[string]interface{}{}
func UpdateCacheEntityViaSerialized(iSerialized int, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("serialized = ?", iSerialized).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaTags via map[string]interface{}{}
func UpdateCacheEntityViaTags(iTags string, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("tags = ?", iTags).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheEntityViaChecksum via map[string]interface{}{}
func UpdateCacheEntityViaChecksum(iChecksum string, iCacheEntityMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheEntity)).Where("checksum = ?", iChecksum).Update(iCacheEntityMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheEntity Delete CacheEntity
func DeleteCacheEntity(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheEntity))
	return row, err
}

// DeleteCacheEntityViaCid Delete CacheEntity via Cid
func DeleteCacheEntityViaCid(iCid string) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Cid: iCid}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaData Delete CacheEntity via Data
func DeleteCacheEntityViaData(iData []byte) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Data: iData}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaExpire Delete CacheEntity via Expire
func DeleteCacheEntityViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Expire: iExpire}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaCreated Delete CacheEntity via Created
func DeleteCacheEntityViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Created: iCreated}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaSerialized Delete CacheEntity via Serialized
func DeleteCacheEntityViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaTags Delete CacheEntity via Tags
func DeleteCacheEntityViaTags(iTags string) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Tags: iTags}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheEntityViaChecksum Delete CacheEntity via Checksum
func DeleteCacheEntityViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheEntity = &CacheEntity{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheEntity); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheEntity)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
