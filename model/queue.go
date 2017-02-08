package model

type Queue struct {
	ItemId  int64  `xorm:"not null pk autoincr INT(10)"`
	Name    string `xorm:"not null default '' index(name_created) VARCHAR(255)"`
	Data    []byte `xorm:"LONGBLOB"`
	Expire  int    `xorm:"not null default 0 index INT(11)"`
	Created int    `xorm:"not null default 0 index(name_created) INT(11)"`
}

// GetQueuesCount Queues' Count
func GetQueuesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Queue{})
	return total, err
}

// GetQueueCountViaItemId Get Queue via ItemId
func GetQueueCountViaItemId(iItemId int64) int64 {
	n, _ := Engine.Where("item_id = ?", iItemId).Count(&Queue{ItemId: iItemId})
	return n
}

// GetQueueCountViaName Get Queue via Name
func GetQueueCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Queue{Name: iName})
	return n
}

// GetQueueCountViaData Get Queue via Data
func GetQueueCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&Queue{Data: iData})
	return n
}

// GetQueueCountViaExpire Get Queue via Expire
func GetQueueCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&Queue{Expire: iExpire})
	return n
}

// GetQueueCountViaCreated Get Queue via Created
func GetQueueCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&Queue{Created: iCreated})
	return n
}

// GetQueuesByItemIdAndNameAndData Get Queues via ItemIdAndNameAndData
func GetQueuesByItemIdAndNameAndData(offset int, limit int, ItemId_ int64, Name_ string, Data_ []byte) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and name = ? and data = ?", ItemId_, Name_, Data_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndNameAndExpire Get Queues via ItemIdAndNameAndExpire
func GetQueuesByItemIdAndNameAndExpire(offset int, limit int, ItemId_ int64, Name_ string, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and name = ? and expire = ?", ItemId_, Name_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndNameAndCreated Get Queues via ItemIdAndNameAndCreated
func GetQueuesByItemIdAndNameAndCreated(offset int, limit int, ItemId_ int64, Name_ string, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and name = ? and created = ?", ItemId_, Name_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndDataAndExpire Get Queues via ItemIdAndDataAndExpire
func GetQueuesByItemIdAndDataAndExpire(offset int, limit int, ItemId_ int64, Data_ []byte, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and data = ? and expire = ?", ItemId_, Data_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndDataAndCreated Get Queues via ItemIdAndDataAndCreated
func GetQueuesByItemIdAndDataAndCreated(offset int, limit int, ItemId_ int64, Data_ []byte, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and data = ? and created = ?", ItemId_, Data_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndExpireAndCreated Get Queues via ItemIdAndExpireAndCreated
func GetQueuesByItemIdAndExpireAndCreated(offset int, limit int, ItemId_ int64, Expire_ int, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and expire = ? and created = ?", ItemId_, Expire_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndDataAndExpire Get Queues via NameAndDataAndExpire
func GetQueuesByNameAndDataAndExpire(offset int, limit int, Name_ string, Data_ []byte, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and data = ? and expire = ?", Name_, Data_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndDataAndCreated Get Queues via NameAndDataAndCreated
func GetQueuesByNameAndDataAndCreated(offset int, limit int, Name_ string, Data_ []byte, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and data = ? and created = ?", Name_, Data_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndExpireAndCreated Get Queues via NameAndExpireAndCreated
func GetQueuesByNameAndExpireAndCreated(offset int, limit int, Name_ string, Expire_ int, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and expire = ? and created = ?", Name_, Expire_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByDataAndExpireAndCreated Get Queues via DataAndExpireAndCreated
func GetQueuesByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndName Get Queues via ItemIdAndName
func GetQueuesByItemIdAndName(offset int, limit int, ItemId_ int64, Name_ string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and name = ?", ItemId_, Name_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndData Get Queues via ItemIdAndData
func GetQueuesByItemIdAndData(offset int, limit int, ItemId_ int64, Data_ []byte) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and data = ?", ItemId_, Data_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndExpire Get Queues via ItemIdAndExpire
func GetQueuesByItemIdAndExpire(offset int, limit int, ItemId_ int64, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and expire = ?", ItemId_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByItemIdAndCreated Get Queues via ItemIdAndCreated
func GetQueuesByItemIdAndCreated(offset int, limit int, ItemId_ int64, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ? and created = ?", ItemId_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndData Get Queues via NameAndData
func GetQueuesByNameAndData(offset int, limit int, Name_ string, Data_ []byte) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and data = ?", Name_, Data_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndExpire Get Queues via NameAndExpire
func GetQueuesByNameAndExpire(offset int, limit int, Name_ string, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and expire = ?", Name_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByNameAndCreated Get Queues via NameAndCreated
func GetQueuesByNameAndCreated(offset int, limit int, Name_ string, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ? and created = ?", Name_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByDataAndExpire Get Queues via DataAndExpire
func GetQueuesByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByDataAndCreated Get Queues via DataAndCreated
func GetQueuesByDataAndCreated(offset int, limit int, Data_ []byte, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueuesByExpireAndCreated Get Queues via ExpireAndCreated
func GetQueuesByExpireAndCreated(offset int, limit int, Expire_ int, Created_ int) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_Queue)
	return _Queue, err
}

// GetQueues Get Queues via field
func GetQueues(offset int, limit int, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// GetQueuesViaItemId Get Queues via ItemId
func GetQueuesViaItemId(offset int, limit int, ItemId_ int64, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("item_id = ?", ItemId_).Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// GetQueuesViaName Get Queues via Name
func GetQueuesViaName(offset int, limit int, Name_ string, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// GetQueuesViaData Get Queues via Data
func GetQueuesViaData(offset int, limit int, Data_ []byte, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// GetQueuesViaExpire Get Queues via Expire
func GetQueuesViaExpire(offset int, limit int, Expire_ int, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// GetQueuesViaCreated Get Queues via Created
func GetQueuesViaCreated(offset int, limit int, Created_ int, field string) (*[]*Queue, error) {
	var _Queue = new([]*Queue)
	err := Engine.Table("queue").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_Queue)
	return _Queue, err
}

// HasQueueViaItemId Has Queue via ItemId
func HasQueueViaItemId(iItemId int64) bool {
	if has, err := Engine.Where("item_id = ?", iItemId).Get(new(Queue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasQueueViaName Has Queue via Name
func HasQueueViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Queue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasQueueViaData Has Queue via Data
func HasQueueViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(Queue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasQueueViaExpire Has Queue via Expire
func HasQueueViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(Queue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasQueueViaCreated Has Queue via Created
func HasQueueViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(Queue)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetQueueViaItemId Get Queue via ItemId
func GetQueueViaItemId(iItemId int64) (*Queue, error) {
	var _Queue = &Queue{ItemId: iItemId}
	has, err := Engine.Get(_Queue)
	if has {
		return _Queue, err
	} else {
		return nil, err
	}
}

// GetQueueViaName Get Queue via Name
func GetQueueViaName(iName string) (*Queue, error) {
	var _Queue = &Queue{Name: iName}
	has, err := Engine.Get(_Queue)
	if has {
		return _Queue, err
	} else {
		return nil, err
	}
}

// GetQueueViaData Get Queue via Data
func GetQueueViaData(iData []byte) (*Queue, error) {
	var _Queue = &Queue{Data: iData}
	has, err := Engine.Get(_Queue)
	if has {
		return _Queue, err
	} else {
		return nil, err
	}
}

// GetQueueViaExpire Get Queue via Expire
func GetQueueViaExpire(iExpire int) (*Queue, error) {
	var _Queue = &Queue{Expire: iExpire}
	has, err := Engine.Get(_Queue)
	if has {
		return _Queue, err
	} else {
		return nil, err
	}
}

// GetQueueViaCreated Get Queue via Created
func GetQueueViaCreated(iCreated int) (*Queue, error) {
	var _Queue = &Queue{Created: iCreated}
	has, err := Engine.Get(_Queue)
	if has {
		return _Queue, err
	} else {
		return nil, err
	}
}

// SetQueueViaItemId Set Queue via ItemId
func SetQueueViaItemId(iItemId int64, queue *Queue) (int64, error) {
	queue.ItemId = iItemId
	return Engine.Insert(queue)
}

// SetQueueViaName Set Queue via Name
func SetQueueViaName(iName string, queue *Queue) (int64, error) {
	queue.Name = iName
	return Engine.Insert(queue)
}

// SetQueueViaData Set Queue via Data
func SetQueueViaData(iData []byte, queue *Queue) (int64, error) {
	queue.Data = iData
	return Engine.Insert(queue)
}

// SetQueueViaExpire Set Queue via Expire
func SetQueueViaExpire(iExpire int, queue *Queue) (int64, error) {
	queue.Expire = iExpire
	return Engine.Insert(queue)
}

// SetQueueViaCreated Set Queue via Created
func SetQueueViaCreated(iCreated int, queue *Queue) (int64, error) {
	queue.Created = iCreated
	return Engine.Insert(queue)
}

// AddQueue Add Queue via all columns
func AddQueue(iItemId int64, iName string, iData []byte, iExpire int, iCreated int) error {
	_Queue := &Queue{ItemId: iItemId, Name: iName, Data: iData, Expire: iExpire, Created: iCreated}
	if _, err := Engine.Insert(_Queue); err != nil {
		return err
	} else {
		return nil
	}
}

// PostQueue Post Queue via iQueue
func PostQueue(iQueue *Queue) (int64, error) {
	_, err := Engine.Insert(iQueue)
	return iQueue.ItemId, err
}

// PutQueue Put Queue
func PutQueue(iQueue *Queue) (int64, error) {
	_, err := Engine.Id(iQueue.ItemId).Update(iQueue)
	return iQueue.ItemId, err
}

// PutQueueViaItemId Put Queue via ItemId
func PutQueueViaItemId(ItemId_ int64, iQueue *Queue) (int64, error) {
	row, err := Engine.Update(iQueue, &Queue{ItemId: ItemId_})
	return row, err
}

// PutQueueViaName Put Queue via Name
func PutQueueViaName(Name_ string, iQueue *Queue) (int64, error) {
	row, err := Engine.Update(iQueue, &Queue{Name: Name_})
	return row, err
}

// PutQueueViaData Put Queue via Data
func PutQueueViaData(Data_ []byte, iQueue *Queue) (int64, error) {
	row, err := Engine.Update(iQueue, &Queue{Data: Data_})
	return row, err
}

// PutQueueViaExpire Put Queue via Expire
func PutQueueViaExpire(Expire_ int, iQueue *Queue) (int64, error) {
	row, err := Engine.Update(iQueue, &Queue{Expire: Expire_})
	return row, err
}

// PutQueueViaCreated Put Queue via Created
func PutQueueViaCreated(Created_ int, iQueue *Queue) (int64, error) {
	row, err := Engine.Update(iQueue, &Queue{Created: Created_})
	return row, err
}

// UpdateQueueViaItemId via map[string]interface{}{}
func UpdateQueueViaItemId(iItemId int64, iQueueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Queue)).Where("item_id = ?", iItemId).Update(iQueueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateQueueViaName via map[string]interface{}{}
func UpdateQueueViaName(iName string, iQueueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Queue)).Where("name = ?", iName).Update(iQueueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateQueueViaData via map[string]interface{}{}
func UpdateQueueViaData(iData []byte, iQueueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Queue)).Where("data = ?", iData).Update(iQueueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateQueueViaExpire via map[string]interface{}{}
func UpdateQueueViaExpire(iExpire int, iQueueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Queue)).Where("expire = ?", iExpire).Update(iQueueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateQueueViaCreated via map[string]interface{}{}
func UpdateQueueViaCreated(iCreated int, iQueueMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Queue)).Where("created = ?", iCreated).Update(iQueueMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteQueue Delete Queue
func DeleteQueue(iItemId int64) (int64, error) {
	row, err := Engine.Id(iItemId).Delete(new(Queue))
	return row, err
}

// DeleteQueueViaItemId Delete Queue via ItemId
func DeleteQueueViaItemId(iItemId int64) (err error) {
	var has bool
	var _Queue = &Queue{ItemId: iItemId}
	if has, err = Engine.Get(_Queue); (has == true) && (err == nil) {
		if row, err := Engine.Where("item_id = ?", iItemId).Delete(new(Queue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteQueueViaName Delete Queue via Name
func DeleteQueueViaName(iName string) (err error) {
	var has bool
	var _Queue = &Queue{Name: iName}
	if has, err = Engine.Get(_Queue); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(Queue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteQueueViaData Delete Queue via Data
func DeleteQueueViaData(iData []byte) (err error) {
	var has bool
	var _Queue = &Queue{Data: iData}
	if has, err = Engine.Get(_Queue); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(Queue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteQueueViaExpire Delete Queue via Expire
func DeleteQueueViaExpire(iExpire int) (err error) {
	var has bool
	var _Queue = &Queue{Expire: iExpire}
	if has, err = Engine.Get(_Queue); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(Queue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteQueueViaCreated Delete Queue via Created
func DeleteQueueViaCreated(iCreated int) (err error) {
	var has bool
	var _Queue = &Queue{Created: iCreated}
	if has, err = Engine.Get(_Queue); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(Queue)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
