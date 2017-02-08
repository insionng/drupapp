package model

type CommentEntityStatistics struct {
	EntityId             int64  `xorm:"not null pk default 0 INT(10)"`
	EntityType           string `xorm:"not null pk default 'node' VARCHAR(32)"`
	FieldName            string `xorm:"not null pk default '' VARCHAR(32)"`
	Cid                  int    `xorm:"not null default 0 INT(11)"`
	LastCommentTimestamp int    `xorm:"not null default 0 index INT(11)"`
	LastCommentName      string `xorm:"VARCHAR(60)"`
	LastCommentUid       int    `xorm:"not null default 0 index INT(10)"`
	CommentCount         int    `xorm:"not null default 0 index INT(10)"`
}

// GetCommentEntityStatisticsesCount CommentEntityStatisticss' Count
func GetCommentEntityStatisticsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CommentEntityStatistics{})
	return total, err
}

// GetCommentEntityStatisticsCountViaEntityId Get CommentEntityStatistics via EntityId
func GetCommentEntityStatisticsCountViaEntityId(iEntityId int64) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&CommentEntityStatistics{EntityId: iEntityId})
	return n
}

// GetCommentEntityStatisticsCountViaEntityType Get CommentEntityStatistics via EntityType
func GetCommentEntityStatisticsCountViaEntityType(iEntityType string) int64 {
	n, _ := Engine.Where("entity_type = ?", iEntityType).Count(&CommentEntityStatistics{EntityType: iEntityType})
	return n
}

// GetCommentEntityStatisticsCountViaFieldName Get CommentEntityStatistics via FieldName
func GetCommentEntityStatisticsCountViaFieldName(iFieldName string) int64 {
	n, _ := Engine.Where("field_name = ?", iFieldName).Count(&CommentEntityStatistics{FieldName: iFieldName})
	return n
}

// GetCommentEntityStatisticsCountViaCid Get CommentEntityStatistics via Cid
func GetCommentEntityStatisticsCountViaCid(iCid int) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CommentEntityStatistics{Cid: iCid})
	return n
}

// GetCommentEntityStatisticsCountViaLastCommentTimestamp Get CommentEntityStatistics via LastCommentTimestamp
func GetCommentEntityStatisticsCountViaLastCommentTimestamp(iLastCommentTimestamp int) int64 {
	n, _ := Engine.Where("last_comment_timestamp = ?", iLastCommentTimestamp).Count(&CommentEntityStatistics{LastCommentTimestamp: iLastCommentTimestamp})
	return n
}

// GetCommentEntityStatisticsCountViaLastCommentName Get CommentEntityStatistics via LastCommentName
func GetCommentEntityStatisticsCountViaLastCommentName(iLastCommentName string) int64 {
	n, _ := Engine.Where("last_comment_name = ?", iLastCommentName).Count(&CommentEntityStatistics{LastCommentName: iLastCommentName})
	return n
}

// GetCommentEntityStatisticsCountViaLastCommentUid Get CommentEntityStatistics via LastCommentUid
func GetCommentEntityStatisticsCountViaLastCommentUid(iLastCommentUid int) int64 {
	n, _ := Engine.Where("last_comment_uid = ?", iLastCommentUid).Count(&CommentEntityStatistics{LastCommentUid: iLastCommentUid})
	return n
}

// GetCommentEntityStatisticsCountViaCommentCount Get CommentEntityStatistics via CommentCount
func GetCommentEntityStatisticsCountViaCommentCount(iCommentCount int) int64 {
	n, _ := Engine.Where("comment_count = ?", iCommentCount).Count(&CommentEntityStatistics{CommentCount: iCommentCount})
	return n
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndFieldName Get CommentEntityStatisticses via EntityIdAndEntityTypeAndFieldName
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndFieldName(offset int, limit int, EntityId_ int64, EntityType_ string, FieldName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and field_name = ?", EntityId_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCid Get CommentEntityStatisticses via EntityIdAndEntityTypeAndCid
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCid(offset int, limit int, EntityId_ int64, EntityType_ string, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and cid = ?", EntityId_, EntityType_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentTimestamp Get CommentEntityStatisticses via EntityIdAndEntityTypeAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentTimestamp(offset int, limit int, EntityId_ int64, EntityType_ string, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and last_comment_timestamp = ?", EntityId_, EntityType_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentName Get CommentEntityStatisticses via EntityIdAndEntityTypeAndLastCommentName
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentName(offset int, limit int, EntityId_ int64, EntityType_ string, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and last_comment_name = ?", EntityId_, EntityType_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndEntityTypeAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentUid(offset int, limit int, EntityId_ int64, EntityType_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and last_comment_uid = ?", EntityId_, EntityType_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCommentCount Get CommentEntityStatisticses via EntityIdAndEntityTypeAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCommentCount(offset int, limit int, EntityId_ int64, EntityType_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ? and comment_count = ?", EntityId_, EntityType_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCid Get CommentEntityStatisticses via EntityIdAndFieldNameAndCid
func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCid(offset int, limit int, EntityId_ int64, FieldName_ string, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ? and cid = ?", EntityId_, FieldName_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentTimestamp Get CommentEntityStatisticses via EntityIdAndFieldNameAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentTimestamp(offset int, limit int, EntityId_ int64, FieldName_ string, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ? and last_comment_timestamp = ?", EntityId_, FieldName_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentName Get CommentEntityStatisticses via EntityIdAndFieldNameAndLastCommentName
func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentName(offset int, limit int, EntityId_ int64, FieldName_ string, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ? and last_comment_name = ?", EntityId_, FieldName_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndFieldNameAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentUid(offset int, limit int, EntityId_ int64, FieldName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ? and last_comment_uid = ?", EntityId_, FieldName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCommentCount Get CommentEntityStatisticses via EntityIdAndFieldNameAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCommentCount(offset int, limit int, EntityId_ int64, FieldName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ? and comment_count = ?", EntityId_, FieldName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentTimestamp Get CommentEntityStatisticses via EntityIdAndCidAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentTimestamp(offset int, limit int, EntityId_ int64, Cid_ int, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and cid = ? and last_comment_timestamp = ?", EntityId_, Cid_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentName Get CommentEntityStatisticses via EntityIdAndCidAndLastCommentName
func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentName(offset int, limit int, EntityId_ int64, Cid_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and cid = ? and last_comment_name = ?", EntityId_, Cid_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndCidAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentUid(offset int, limit int, EntityId_ int64, Cid_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and cid = ? and last_comment_uid = ?", EntityId_, Cid_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCidAndCommentCount Get CommentEntityStatisticses via EntityIdAndCidAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndCidAndCommentCount(offset int, limit int, EntityId_ int64, Cid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and cid = ? and comment_count = ?", EntityId_, Cid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentName Get CommentEntityStatisticses via EntityIdAndLastCommentTimestampAndLastCommentName
func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentName(offset int, limit int, EntityId_ int64, LastCommentTimestamp_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_timestamp = ? and last_comment_name = ?", EntityId_, LastCommentTimestamp_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndLastCommentTimestampAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentUid(offset int, limit int, EntityId_ int64, LastCommentTimestamp_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_timestamp = ? and last_comment_uid = ?", EntityId_, LastCommentTimestamp_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndCommentCount Get CommentEntityStatisticses via EntityIdAndLastCommentTimestampAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndCommentCount(offset int, limit int, EntityId_ int64, LastCommentTimestamp_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_timestamp = ? and comment_count = ?", EntityId_, LastCommentTimestamp_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndLastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndLastCommentUid(offset int, limit int, EntityId_ int64, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_name = ? and last_comment_uid = ?", EntityId_, LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndCommentCount Get CommentEntityStatisticses via EntityIdAndLastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndCommentCount(offset int, limit int, EntityId_ int64, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_name = ? and comment_count = ?", EntityId_, LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via EntityIdAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndLastCommentUidAndCommentCount(offset int, limit int, EntityId_ int64, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_uid = ? and comment_count = ?", EntityId_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCid Get CommentEntityStatisticses via EntityTypeAndFieldNameAndCid
func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCid(offset int, limit int, EntityType_ string, FieldName_ string, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ? and cid = ?", EntityType_, FieldName_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentTimestamp Get CommentEntityStatisticses via EntityTypeAndFieldNameAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentTimestamp(offset int, limit int, EntityType_ string, FieldName_ string, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ? and last_comment_timestamp = ?", EntityType_, FieldName_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentName Get CommentEntityStatisticses via EntityTypeAndFieldNameAndLastCommentName
func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentName(offset int, limit int, EntityType_ string, FieldName_ string, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ? and last_comment_name = ?", EntityType_, FieldName_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentUid Get CommentEntityStatisticses via EntityTypeAndFieldNameAndLastCommentUid
func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentUid(offset int, limit int, EntityType_ string, FieldName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ? and last_comment_uid = ?", EntityType_, FieldName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCommentCount Get CommentEntityStatisticses via EntityTypeAndFieldNameAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCommentCount(offset int, limit int, EntityType_ string, FieldName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ? and comment_count = ?", EntityType_, FieldName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentTimestamp Get CommentEntityStatisticses via EntityTypeAndCidAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentTimestamp(offset int, limit int, EntityType_ string, Cid_ int, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and cid = ? and last_comment_timestamp = ?", EntityType_, Cid_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentName Get CommentEntityStatisticses via EntityTypeAndCidAndLastCommentName
func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentName(offset int, limit int, EntityType_ string, Cid_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and cid = ? and last_comment_name = ?", EntityType_, Cid_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentUid Get CommentEntityStatisticses via EntityTypeAndCidAndLastCommentUid
func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentUid(offset int, limit int, EntityType_ string, Cid_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and cid = ? and last_comment_uid = ?", EntityType_, Cid_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCidAndCommentCount Get CommentEntityStatisticses via EntityTypeAndCidAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndCidAndCommentCount(offset int, limit int, EntityType_ string, Cid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and cid = ? and comment_count = ?", EntityType_, Cid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentName Get CommentEntityStatisticses via EntityTypeAndLastCommentTimestampAndLastCommentName
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentName(offset int, limit int, EntityType_ string, LastCommentTimestamp_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_timestamp = ? and last_comment_name = ?", EntityType_, LastCommentTimestamp_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentUid Get CommentEntityStatisticses via EntityTypeAndLastCommentTimestampAndLastCommentUid
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentUid(offset int, limit int, EntityType_ string, LastCommentTimestamp_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_timestamp = ? and last_comment_uid = ?", EntityType_, LastCommentTimestamp_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndCommentCount Get CommentEntityStatisticses via EntityTypeAndLastCommentTimestampAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndCommentCount(offset int, limit int, EntityType_ string, LastCommentTimestamp_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_timestamp = ? and comment_count = ?", EntityType_, LastCommentTimestamp_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via EntityTypeAndLastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndLastCommentUid(offset int, limit int, EntityType_ string, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_name = ? and last_comment_uid = ?", EntityType_, LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndCommentCount Get CommentEntityStatisticses via EntityTypeAndLastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndCommentCount(offset int, limit int, EntityType_ string, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_name = ? and comment_count = ?", EntityType_, LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via EntityTypeAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidAndCommentCount(offset int, limit int, EntityType_ string, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_uid = ? and comment_count = ?", EntityType_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentTimestamp Get CommentEntityStatisticses via FieldNameAndCidAndLastCommentTimestamp
func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentTimestamp(offset int, limit int, FieldName_ string, Cid_ int, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and cid = ? and last_comment_timestamp = ?", FieldName_, Cid_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentName Get CommentEntityStatisticses via FieldNameAndCidAndLastCommentName
func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentName(offset int, limit int, FieldName_ string, Cid_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and cid = ? and last_comment_name = ?", FieldName_, Cid_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentUid Get CommentEntityStatisticses via FieldNameAndCidAndLastCommentUid
func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentUid(offset int, limit int, FieldName_ string, Cid_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and cid = ? and last_comment_uid = ?", FieldName_, Cid_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCidAndCommentCount Get CommentEntityStatisticses via FieldNameAndCidAndCommentCount
func GetCommentEntityStatisticsesByFieldNameAndCidAndCommentCount(offset int, limit int, FieldName_ string, Cid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and cid = ? and comment_count = ?", FieldName_, Cid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentName Get CommentEntityStatisticses via FieldNameAndLastCommentTimestampAndLastCommentName
func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentName(offset int, limit int, FieldName_ string, LastCommentTimestamp_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_timestamp = ? and last_comment_name = ?", FieldName_, LastCommentTimestamp_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentUid Get CommentEntityStatisticses via FieldNameAndLastCommentTimestampAndLastCommentUid
func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentUid(offset int, limit int, FieldName_ string, LastCommentTimestamp_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_timestamp = ? and last_comment_uid = ?", FieldName_, LastCommentTimestamp_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndCommentCount Get CommentEntityStatisticses via FieldNameAndLastCommentTimestampAndCommentCount
func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndCommentCount(offset int, limit int, FieldName_ string, LastCommentTimestamp_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_timestamp = ? and comment_count = ?", FieldName_, LastCommentTimestamp_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via FieldNameAndLastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndLastCommentUid(offset int, limit int, FieldName_ string, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_name = ? and last_comment_uid = ?", FieldName_, LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndCommentCount Get CommentEntityStatisticses via FieldNameAndLastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndCommentCount(offset int, limit int, FieldName_ string, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_name = ? and comment_count = ?", FieldName_, LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via FieldNameAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByFieldNameAndLastCommentUidAndCommentCount(offset int, limit int, FieldName_ string, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_uid = ? and comment_count = ?", FieldName_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentName Get CommentEntityStatisticses via CidAndLastCommentTimestampAndLastCommentName
func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentName(offset int, limit int, Cid_ int, LastCommentTimestamp_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_timestamp = ? and last_comment_name = ?", Cid_, LastCommentTimestamp_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentUid Get CommentEntityStatisticses via CidAndLastCommentTimestampAndLastCommentUid
func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentUid(offset int, limit int, Cid_ int, LastCommentTimestamp_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_timestamp = ? and last_comment_uid = ?", Cid_, LastCommentTimestamp_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndCommentCount Get CommentEntityStatisticses via CidAndLastCommentTimestampAndCommentCount
func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndCommentCount(offset int, limit int, Cid_ int, LastCommentTimestamp_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_timestamp = ? and comment_count = ?", Cid_, LastCommentTimestamp_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via CidAndLastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByCidAndLastCommentNameAndLastCommentUid(offset int, limit int, Cid_ int, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_name = ? and last_comment_uid = ?", Cid_, LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentNameAndCommentCount Get CommentEntityStatisticses via CidAndLastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByCidAndLastCommentNameAndCommentCount(offset int, limit int, Cid_ int, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_name = ? and comment_count = ?", Cid_, LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via CidAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByCidAndLastCommentUidAndCommentCount(offset int, limit int, Cid_ int, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_uid = ? and comment_count = ?", Cid_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via LastCommentTimestampAndLastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndLastCommentUid(offset int, limit int, LastCommentTimestamp_ int, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and last_comment_name = ? and last_comment_uid = ?", LastCommentTimestamp_, LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndCommentCount Get CommentEntityStatisticses via LastCommentTimestampAndLastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndCommentCount(offset int, limit int, LastCommentTimestamp_ int, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and last_comment_name = ? and comment_count = ?", LastCommentTimestamp_, LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via LastCommentTimestampAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidAndCommentCount(offset int, limit int, LastCommentTimestamp_ int, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and last_comment_uid = ? and comment_count = ?", LastCommentTimestamp_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidAndCommentCount Get CommentEntityStatisticses via LastCommentNameAndLastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidAndCommentCount(offset int, limit int, LastCommentName_ string, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_name = ? and last_comment_uid = ? and comment_count = ?", LastCommentName_, LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndEntityType Get CommentEntityStatisticses via EntityIdAndEntityType
func GetCommentEntityStatisticsesByEntityIdAndEntityType(offset int, limit int, EntityId_ int64, EntityType_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and entity_type = ?", EntityId_, EntityType_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndFieldName Get CommentEntityStatisticses via EntityIdAndFieldName
func GetCommentEntityStatisticsesByEntityIdAndFieldName(offset int, limit int, EntityId_ int64, FieldName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and field_name = ?", EntityId_, FieldName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCid Get CommentEntityStatisticses via EntityIdAndCid
func GetCommentEntityStatisticsesByEntityIdAndCid(offset int, limit int, EntityId_ int64, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and cid = ?", EntityId_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestamp Get CommentEntityStatisticses via EntityIdAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestamp(offset int, limit int, EntityId_ int64, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_timestamp = ?", EntityId_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentName Get CommentEntityStatisticses via EntityIdAndLastCommentName
func GetCommentEntityStatisticsesByEntityIdAndLastCommentName(offset int, limit int, EntityId_ int64, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_name = ?", EntityId_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndLastCommentUid Get CommentEntityStatisticses via EntityIdAndLastCommentUid
func GetCommentEntityStatisticsesByEntityIdAndLastCommentUid(offset int, limit int, EntityId_ int64, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and last_comment_uid = ?", EntityId_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityIdAndCommentCount Get CommentEntityStatisticses via EntityIdAndCommentCount
func GetCommentEntityStatisticsesByEntityIdAndCommentCount(offset int, limit int, EntityId_ int64, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ? and comment_count = ?", EntityId_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndFieldName Get CommentEntityStatisticses via EntityTypeAndFieldName
func GetCommentEntityStatisticsesByEntityTypeAndFieldName(offset int, limit int, EntityType_ string, FieldName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and field_name = ?", EntityType_, FieldName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCid Get CommentEntityStatisticses via EntityTypeAndCid
func GetCommentEntityStatisticsesByEntityTypeAndCid(offset int, limit int, EntityType_ string, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and cid = ?", EntityType_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestamp Get CommentEntityStatisticses via EntityTypeAndLastCommentTimestamp
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestamp(offset int, limit int, EntityType_ string, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_timestamp = ?", EntityType_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentName Get CommentEntityStatisticses via EntityTypeAndLastCommentName
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentName(offset int, limit int, EntityType_ string, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_name = ?", EntityType_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndLastCommentUid Get CommentEntityStatisticses via EntityTypeAndLastCommentUid
func GetCommentEntityStatisticsesByEntityTypeAndLastCommentUid(offset int, limit int, EntityType_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and last_comment_uid = ?", EntityType_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByEntityTypeAndCommentCount Get CommentEntityStatisticses via EntityTypeAndCommentCount
func GetCommentEntityStatisticsesByEntityTypeAndCommentCount(offset int, limit int, EntityType_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ? and comment_count = ?", EntityType_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCid Get CommentEntityStatisticses via FieldNameAndCid
func GetCommentEntityStatisticsesByFieldNameAndCid(offset int, limit int, FieldName_ string, Cid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and cid = ?", FieldName_, Cid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestamp Get CommentEntityStatisticses via FieldNameAndLastCommentTimestamp
func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestamp(offset int, limit int, FieldName_ string, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_timestamp = ?", FieldName_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentName Get CommentEntityStatisticses via FieldNameAndLastCommentName
func GetCommentEntityStatisticsesByFieldNameAndLastCommentName(offset int, limit int, FieldName_ string, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_name = ?", FieldName_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndLastCommentUid Get CommentEntityStatisticses via FieldNameAndLastCommentUid
func GetCommentEntityStatisticsesByFieldNameAndLastCommentUid(offset int, limit int, FieldName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and last_comment_uid = ?", FieldName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByFieldNameAndCommentCount Get CommentEntityStatisticses via FieldNameAndCommentCount
func GetCommentEntityStatisticsesByFieldNameAndCommentCount(offset int, limit int, FieldName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ? and comment_count = ?", FieldName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentTimestamp Get CommentEntityStatisticses via CidAndLastCommentTimestamp
func GetCommentEntityStatisticsesByCidAndLastCommentTimestamp(offset int, limit int, Cid_ int, LastCommentTimestamp_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_timestamp = ?", Cid_, LastCommentTimestamp_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentName Get CommentEntityStatisticses via CidAndLastCommentName
func GetCommentEntityStatisticsesByCidAndLastCommentName(offset int, limit int, Cid_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_name = ?", Cid_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndLastCommentUid Get CommentEntityStatisticses via CidAndLastCommentUid
func GetCommentEntityStatisticsesByCidAndLastCommentUid(offset int, limit int, Cid_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and last_comment_uid = ?", Cid_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByCidAndCommentCount Get CommentEntityStatisticses via CidAndCommentCount
func GetCommentEntityStatisticsesByCidAndCommentCount(offset int, limit int, Cid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ? and comment_count = ?", Cid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentName Get CommentEntityStatisticses via LastCommentTimestampAndLastCommentName
func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentName(offset int, limit int, LastCommentTimestamp_ int, LastCommentName_ string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and last_comment_name = ?", LastCommentTimestamp_, LastCommentName_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUid Get CommentEntityStatisticses via LastCommentTimestampAndLastCommentUid
func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUid(offset int, limit int, LastCommentTimestamp_ int, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and last_comment_uid = ?", LastCommentTimestamp_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentTimestampAndCommentCount Get CommentEntityStatisticses via LastCommentTimestampAndCommentCount
func GetCommentEntityStatisticsesByLastCommentTimestampAndCommentCount(offset int, limit int, LastCommentTimestamp_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ? and comment_count = ?", LastCommentTimestamp_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUid Get CommentEntityStatisticses via LastCommentNameAndLastCommentUid
func GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUid(offset int, limit int, LastCommentName_ string, LastCommentUid_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_name = ? and last_comment_uid = ?", LastCommentName_, LastCommentUid_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentNameAndCommentCount Get CommentEntityStatisticses via LastCommentNameAndCommentCount
func GetCommentEntityStatisticsesByLastCommentNameAndCommentCount(offset int, limit int, LastCommentName_ string, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_name = ? and comment_count = ?", LastCommentName_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesByLastCommentUidAndCommentCount Get CommentEntityStatisticses via LastCommentUidAndCommentCount
func GetCommentEntityStatisticsesByLastCommentUidAndCommentCount(offset int, limit int, LastCommentUid_ int, CommentCount_ int) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_uid = ? and comment_count = ?", LastCommentUid_, CommentCount_).Limit(limit, offset).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticses Get CommentEntityStatisticses via field
func GetCommentEntityStatisticses(offset int, limit int, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaEntityId Get CommentEntityStatisticss via EntityId
func GetCommentEntityStatisticsesViaEntityId(offset int, limit int, EntityId_ int64, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaEntityType Get CommentEntityStatisticss via EntityType
func GetCommentEntityStatisticsesViaEntityType(offset int, limit int, EntityType_ string, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("entity_type = ?", EntityType_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaFieldName Get CommentEntityStatisticss via FieldName
func GetCommentEntityStatisticsesViaFieldName(offset int, limit int, FieldName_ string, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("field_name = ?", FieldName_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaCid Get CommentEntityStatisticss via Cid
func GetCommentEntityStatisticsesViaCid(offset int, limit int, Cid_ int, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaLastCommentTimestamp Get CommentEntityStatisticss via LastCommentTimestamp
func GetCommentEntityStatisticsesViaLastCommentTimestamp(offset int, limit int, LastCommentTimestamp_ int, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_timestamp = ?", LastCommentTimestamp_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaLastCommentName Get CommentEntityStatisticss via LastCommentName
func GetCommentEntityStatisticsesViaLastCommentName(offset int, limit int, LastCommentName_ string, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_name = ?", LastCommentName_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaLastCommentUid Get CommentEntityStatisticss via LastCommentUid
func GetCommentEntityStatisticsesViaLastCommentUid(offset int, limit int, LastCommentUid_ int, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("last_comment_uid = ?", LastCommentUid_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// GetCommentEntityStatisticsesViaCommentCount Get CommentEntityStatisticss via CommentCount
func GetCommentEntityStatisticsesViaCommentCount(offset int, limit int, CommentCount_ int, field string) (*[]*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = new([]*CommentEntityStatistics)
	err := Engine.Table("comment_entity_statistics").Where("comment_count = ?", CommentCount_).Limit(limit, offset).Desc(field).Find(_CommentEntityStatistics)
	return _CommentEntityStatistics, err
}

// HasCommentEntityStatisticsViaEntityId Has CommentEntityStatistics via EntityId
func HasCommentEntityStatisticsViaEntityId(iEntityId int64) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaEntityType Has CommentEntityStatistics via EntityType
func HasCommentEntityStatisticsViaEntityType(iEntityType string) bool {
	if has, err := Engine.Where("entity_type = ?", iEntityType).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaFieldName Has CommentEntityStatistics via FieldName
func HasCommentEntityStatisticsViaFieldName(iFieldName string) bool {
	if has, err := Engine.Where("field_name = ?", iFieldName).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaCid Has CommentEntityStatistics via Cid
func HasCommentEntityStatisticsViaCid(iCid int) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaLastCommentTimestamp Has CommentEntityStatistics via LastCommentTimestamp
func HasCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp int) bool {
	if has, err := Engine.Where("last_comment_timestamp = ?", iLastCommentTimestamp).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaLastCommentName Has CommentEntityStatistics via LastCommentName
func HasCommentEntityStatisticsViaLastCommentName(iLastCommentName string) bool {
	if has, err := Engine.Where("last_comment_name = ?", iLastCommentName).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaLastCommentUid Has CommentEntityStatistics via LastCommentUid
func HasCommentEntityStatisticsViaLastCommentUid(iLastCommentUid int) bool {
	if has, err := Engine.Where("last_comment_uid = ?", iLastCommentUid).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentEntityStatisticsViaCommentCount Has CommentEntityStatistics via CommentCount
func HasCommentEntityStatisticsViaCommentCount(iCommentCount int) bool {
	if has, err := Engine.Where("comment_count = ?", iCommentCount).Get(new(CommentEntityStatistics)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentEntityStatisticsViaEntityId Get CommentEntityStatistics via EntityId
func GetCommentEntityStatisticsViaEntityId(iEntityId int64) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{EntityId: iEntityId}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaEntityType Get CommentEntityStatistics via EntityType
func GetCommentEntityStatisticsViaEntityType(iEntityType string) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{EntityType: iEntityType}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaFieldName Get CommentEntityStatistics via FieldName
func GetCommentEntityStatisticsViaFieldName(iFieldName string) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{FieldName: iFieldName}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaCid Get CommentEntityStatistics via Cid
func GetCommentEntityStatisticsViaCid(iCid int) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{Cid: iCid}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaLastCommentTimestamp Get CommentEntityStatistics via LastCommentTimestamp
func GetCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp int) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentTimestamp: iLastCommentTimestamp}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaLastCommentName Get CommentEntityStatistics via LastCommentName
func GetCommentEntityStatisticsViaLastCommentName(iLastCommentName string) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentName: iLastCommentName}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaLastCommentUid Get CommentEntityStatistics via LastCommentUid
func GetCommentEntityStatisticsViaLastCommentUid(iLastCommentUid int) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentUid: iLastCommentUid}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// GetCommentEntityStatisticsViaCommentCount Get CommentEntityStatistics via CommentCount
func GetCommentEntityStatisticsViaCommentCount(iCommentCount int) (*CommentEntityStatistics, error) {
	var _CommentEntityStatistics = &CommentEntityStatistics{CommentCount: iCommentCount}
	has, err := Engine.Get(_CommentEntityStatistics)
	if has {
		return _CommentEntityStatistics, err
	} else {
		return nil, err
	}
}

// SetCommentEntityStatisticsViaEntityId Set CommentEntityStatistics via EntityId
func SetCommentEntityStatisticsViaEntityId(iEntityId int64, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.EntityId = iEntityId
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaEntityType Set CommentEntityStatistics via EntityType
func SetCommentEntityStatisticsViaEntityType(iEntityType string, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.EntityType = iEntityType
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaFieldName Set CommentEntityStatistics via FieldName
func SetCommentEntityStatisticsViaFieldName(iFieldName string, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.FieldName = iFieldName
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaCid Set CommentEntityStatistics via Cid
func SetCommentEntityStatisticsViaCid(iCid int, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.Cid = iCid
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaLastCommentTimestamp Set CommentEntityStatistics via LastCommentTimestamp
func SetCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp int, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.LastCommentTimestamp = iLastCommentTimestamp
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaLastCommentName Set CommentEntityStatistics via LastCommentName
func SetCommentEntityStatisticsViaLastCommentName(iLastCommentName string, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.LastCommentName = iLastCommentName
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaLastCommentUid Set CommentEntityStatistics via LastCommentUid
func SetCommentEntityStatisticsViaLastCommentUid(iLastCommentUid int, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.LastCommentUid = iLastCommentUid
	return Engine.Insert(comment_entity_statistics)
}

// SetCommentEntityStatisticsViaCommentCount Set CommentEntityStatistics via CommentCount
func SetCommentEntityStatisticsViaCommentCount(iCommentCount int, comment_entity_statistics *CommentEntityStatistics) (int64, error) {
	comment_entity_statistics.CommentCount = iCommentCount
	return Engine.Insert(comment_entity_statistics)
}

// AddCommentEntityStatistics Add CommentEntityStatistics via all columns
func AddCommentEntityStatistics(iEntityId int64, iEntityType string, iFieldName string, iCid int, iLastCommentTimestamp int, iLastCommentName string, iLastCommentUid int, iCommentCount int) error {
	_CommentEntityStatistics := &CommentEntityStatistics{EntityId: iEntityId, EntityType: iEntityType, FieldName: iFieldName, Cid: iCid, LastCommentTimestamp: iLastCommentTimestamp, LastCommentName: iLastCommentName, LastCommentUid: iLastCommentUid, CommentCount: iCommentCount}
	if _, err := Engine.Insert(_CommentEntityStatistics); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCommentEntityStatistics Post CommentEntityStatistics via iCommentEntityStatistics
func PostCommentEntityStatistics(iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	_, err := Engine.Insert(iCommentEntityStatistics)
	return iCommentEntityStatistics.EntityId, err
}

// PutCommentEntityStatistics Put CommentEntityStatistics
func PutCommentEntityStatistics(iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	_, err := Engine.Id(iCommentEntityStatistics.EntityId).Update(iCommentEntityStatistics)
	return iCommentEntityStatistics.EntityId, err
}

// PutCommentEntityStatisticsViaEntityId Put CommentEntityStatistics via EntityId
func PutCommentEntityStatisticsViaEntityId(EntityId_ int64, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{EntityId: EntityId_})
	return row, err
}

// PutCommentEntityStatisticsViaEntityType Put CommentEntityStatistics via EntityType
func PutCommentEntityStatisticsViaEntityType(EntityType_ string, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{EntityType: EntityType_})
	return row, err
}

// PutCommentEntityStatisticsViaFieldName Put CommentEntityStatistics via FieldName
func PutCommentEntityStatisticsViaFieldName(FieldName_ string, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{FieldName: FieldName_})
	return row, err
}

// PutCommentEntityStatisticsViaCid Put CommentEntityStatistics via Cid
func PutCommentEntityStatisticsViaCid(Cid_ int, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{Cid: Cid_})
	return row, err
}

// PutCommentEntityStatisticsViaLastCommentTimestamp Put CommentEntityStatistics via LastCommentTimestamp
func PutCommentEntityStatisticsViaLastCommentTimestamp(LastCommentTimestamp_ int, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{LastCommentTimestamp: LastCommentTimestamp_})
	return row, err
}

// PutCommentEntityStatisticsViaLastCommentName Put CommentEntityStatistics via LastCommentName
func PutCommentEntityStatisticsViaLastCommentName(LastCommentName_ string, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{LastCommentName: LastCommentName_})
	return row, err
}

// PutCommentEntityStatisticsViaLastCommentUid Put CommentEntityStatistics via LastCommentUid
func PutCommentEntityStatisticsViaLastCommentUid(LastCommentUid_ int, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{LastCommentUid: LastCommentUid_})
	return row, err
}

// PutCommentEntityStatisticsViaCommentCount Put CommentEntityStatistics via CommentCount
func PutCommentEntityStatisticsViaCommentCount(CommentCount_ int, iCommentEntityStatistics *CommentEntityStatistics) (int64, error) {
	row, err := Engine.Update(iCommentEntityStatistics, &CommentEntityStatistics{CommentCount: CommentCount_})
	return row, err
}

// UpdateCommentEntityStatisticsViaEntityId via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaEntityId(iEntityId int64, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("entity_id = ?", iEntityId).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaEntityType via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaEntityType(iEntityType string, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("entity_type = ?", iEntityType).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaFieldName via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaFieldName(iFieldName string, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("field_name = ?", iFieldName).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaCid via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaCid(iCid int, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("cid = ?", iCid).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaLastCommentTimestamp via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp int, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("last_comment_timestamp = ?", iLastCommentTimestamp).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaLastCommentName via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaLastCommentName(iLastCommentName string, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("last_comment_name = ?", iLastCommentName).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaLastCommentUid via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaLastCommentUid(iLastCommentUid int, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("last_comment_uid = ?", iLastCommentUid).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentEntityStatisticsViaCommentCount via map[string]interface{}{}
func UpdateCommentEntityStatisticsViaCommentCount(iCommentCount int, iCommentEntityStatisticsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentEntityStatistics)).Where("comment_count = ?", iCommentCount).Update(iCommentEntityStatisticsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCommentEntityStatistics Delete CommentEntityStatistics
func DeleteCommentEntityStatistics(iEntityId int64) (int64, error) {
	row, err := Engine.Id(iEntityId).Delete(new(CommentEntityStatistics))
	return row, err
}

// DeleteCommentEntityStatisticsViaEntityId Delete CommentEntityStatistics via EntityId
func DeleteCommentEntityStatisticsViaEntityId(iEntityId int64) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{EntityId: iEntityId}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaEntityType Delete CommentEntityStatistics via EntityType
func DeleteCommentEntityStatisticsViaEntityType(iEntityType string) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{EntityType: iEntityType}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_type = ?", iEntityType).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaFieldName Delete CommentEntityStatistics via FieldName
func DeleteCommentEntityStatisticsViaFieldName(iFieldName string) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{FieldName: iFieldName}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_name = ?", iFieldName).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaCid Delete CommentEntityStatistics via Cid
func DeleteCommentEntityStatisticsViaCid(iCid int) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{Cid: iCid}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaLastCommentTimestamp Delete CommentEntityStatistics via LastCommentTimestamp
func DeleteCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp int) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentTimestamp: iLastCommentTimestamp}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("last_comment_timestamp = ?", iLastCommentTimestamp).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaLastCommentName Delete CommentEntityStatistics via LastCommentName
func DeleteCommentEntityStatisticsViaLastCommentName(iLastCommentName string) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentName: iLastCommentName}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("last_comment_name = ?", iLastCommentName).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaLastCommentUid Delete CommentEntityStatistics via LastCommentUid
func DeleteCommentEntityStatisticsViaLastCommentUid(iLastCommentUid int) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{LastCommentUid: iLastCommentUid}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("last_comment_uid = ?", iLastCommentUid).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentEntityStatisticsViaCommentCount Delete CommentEntityStatistics via CommentCount
func DeleteCommentEntityStatisticsViaCommentCount(iCommentCount int) (err error) {
	var has bool
	var _CommentEntityStatistics = &CommentEntityStatistics{CommentCount: iCommentCount}
	if has, err = Engine.Get(_CommentEntityStatistics); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_count = ?", iCommentCount).Delete(new(CommentEntityStatistics)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
