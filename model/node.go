package model

type Node struct {
	Nid      int64  `xorm:"not null pk autoincr INT(10)"`
	Vid      int    `xorm:"unique INT(10)"`
	Type     string `xorm:"not null index VARCHAR(32)"`
	Uuid     string `xorm:"not null unique VARCHAR(128)"`
	Langcode string `xorm:"not null VARCHAR(12)"`
}

// GetNodesCount Nodes' Count
func GetNodesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Node{})
	return total, err
}

// GetNodeCountViaNid Get Node via Nid
func GetNodeCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&Node{Nid: iNid})
	return n
}

// GetNodeCountViaVid Get Node via Vid
func GetNodeCountViaVid(iVid int) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&Node{Vid: iVid})
	return n
}

// GetNodeCountViaType Get Node via Type
func GetNodeCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&Node{Type: iType})
	return n
}

// GetNodeCountViaUuid Get Node via Uuid
func GetNodeCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&Node{Uuid: iUuid})
	return n
}

// GetNodeCountViaLangcode Get Node via Langcode
func GetNodeCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Node{Langcode: iLangcode})
	return n
}

// GetNodesByNidAndVidAndType Get Nodes via NidAndVidAndType
func GetNodesByNidAndVidAndType(offset int, limit int, Nid_ int64, Vid_ int, Type_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and vid = ? and type = ?", Nid_, Vid_, Type_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndVidAndUuid Get Nodes via NidAndVidAndUuid
func GetNodesByNidAndVidAndUuid(offset int, limit int, Nid_ int64, Vid_ int, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and vid = ? and uuid = ?", Nid_, Vid_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndVidAndLangcode Get Nodes via NidAndVidAndLangcode
func GetNodesByNidAndVidAndLangcode(offset int, limit int, Nid_ int64, Vid_ int, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and vid = ? and langcode = ?", Nid_, Vid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndTypeAndUuid Get Nodes via NidAndTypeAndUuid
func GetNodesByNidAndTypeAndUuid(offset int, limit int, Nid_ int64, Type_ string, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and type = ? and uuid = ?", Nid_, Type_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndTypeAndLangcode Get Nodes via NidAndTypeAndLangcode
func GetNodesByNidAndTypeAndLangcode(offset int, limit int, Nid_ int64, Type_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and type = ? and langcode = ?", Nid_, Type_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndUuidAndLangcode Get Nodes via NidAndUuidAndLangcode
func GetNodesByNidAndUuidAndLangcode(offset int, limit int, Nid_ int64, Uuid_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and uuid = ? and langcode = ?", Nid_, Uuid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndTypeAndUuid Get Nodes via VidAndTypeAndUuid
func GetNodesByVidAndTypeAndUuid(offset int, limit int, Vid_ int, Type_ string, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and type = ? and uuid = ?", Vid_, Type_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndTypeAndLangcode Get Nodes via VidAndTypeAndLangcode
func GetNodesByVidAndTypeAndLangcode(offset int, limit int, Vid_ int, Type_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and type = ? and langcode = ?", Vid_, Type_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndUuidAndLangcode Get Nodes via VidAndUuidAndLangcode
func GetNodesByVidAndUuidAndLangcode(offset int, limit int, Vid_ int, Uuid_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and uuid = ? and langcode = ?", Vid_, Uuid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByTypeAndUuidAndLangcode Get Nodes via TypeAndUuidAndLangcode
func GetNodesByTypeAndUuidAndLangcode(offset int, limit int, Type_ string, Uuid_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("type = ? and uuid = ? and langcode = ?", Type_, Uuid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndVid Get Nodes via NidAndVid
func GetNodesByNidAndVid(offset int, limit int, Nid_ int64, Vid_ int) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and vid = ?", Nid_, Vid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndType Get Nodes via NidAndType
func GetNodesByNidAndType(offset int, limit int, Nid_ int64, Type_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and type = ?", Nid_, Type_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndUuid Get Nodes via NidAndUuid
func GetNodesByNidAndUuid(offset int, limit int, Nid_ int64, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and uuid = ?", Nid_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByNidAndLangcode Get Nodes via NidAndLangcode
func GetNodesByNidAndLangcode(offset int, limit int, Nid_ int64, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ? and langcode = ?", Nid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndType Get Nodes via VidAndType
func GetNodesByVidAndType(offset int, limit int, Vid_ int, Type_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and type = ?", Vid_, Type_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndUuid Get Nodes via VidAndUuid
func GetNodesByVidAndUuid(offset int, limit int, Vid_ int, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and uuid = ?", Vid_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByVidAndLangcode Get Nodes via VidAndLangcode
func GetNodesByVidAndLangcode(offset int, limit int, Vid_ int, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByTypeAndUuid Get Nodes via TypeAndUuid
func GetNodesByTypeAndUuid(offset int, limit int, Type_ string, Uuid_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("type = ? and uuid = ?", Type_, Uuid_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByTypeAndLangcode Get Nodes via TypeAndLangcode
func GetNodesByTypeAndLangcode(offset int, limit int, Type_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("type = ? and langcode = ?", Type_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodesByUuidAndLangcode Get Nodes via UuidAndLangcode
func GetNodesByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_Node)
	return _Node, err
}

// GetNodes Get Nodes via field
func GetNodes(offset int, limit int, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// GetNodesViaNid Get Nodes via Nid
func GetNodesViaNid(offset int, limit int, Nid_ int64, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// GetNodesViaVid Get Nodes via Vid
func GetNodesViaVid(offset int, limit int, Vid_ int, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// GetNodesViaType Get Nodes via Type
func GetNodesViaType(offset int, limit int, Type_ string, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// GetNodesViaUuid Get Nodes via Uuid
func GetNodesViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// GetNodesViaLangcode Get Nodes via Langcode
func GetNodesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Node, error) {
	var _Node = new([]*Node)
	err := Engine.Table("node").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Node)
	return _Node, err
}

// HasNodeViaNid Has Node via Nid
func HasNodeViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(Node)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeViaVid Has Node via Vid
func HasNodeViaVid(iVid int) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(Node)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeViaType Has Node via Type
func HasNodeViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(Node)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeViaUuid Has Node via Uuid
func HasNodeViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(Node)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeViaLangcode Has Node via Langcode
func HasNodeViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Node)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeViaNid Get Node via Nid
func GetNodeViaNid(iNid int64) (*Node, error) {
	var _Node = &Node{Nid: iNid}
	has, err := Engine.Get(_Node)
	if has {
		return _Node, err
	} else {
		return nil, err
	}
}

// GetNodeViaVid Get Node via Vid
func GetNodeViaVid(iVid int) (*Node, error) {
	var _Node = &Node{Vid: iVid}
	has, err := Engine.Get(_Node)
	if has {
		return _Node, err
	} else {
		return nil, err
	}
}

// GetNodeViaType Get Node via Type
func GetNodeViaType(iType string) (*Node, error) {
	var _Node = &Node{Type: iType}
	has, err := Engine.Get(_Node)
	if has {
		return _Node, err
	} else {
		return nil, err
	}
}

// GetNodeViaUuid Get Node via Uuid
func GetNodeViaUuid(iUuid string) (*Node, error) {
	var _Node = &Node{Uuid: iUuid}
	has, err := Engine.Get(_Node)
	if has {
		return _Node, err
	} else {
		return nil, err
	}
}

// GetNodeViaLangcode Get Node via Langcode
func GetNodeViaLangcode(iLangcode string) (*Node, error) {
	var _Node = &Node{Langcode: iLangcode}
	has, err := Engine.Get(_Node)
	if has {
		return _Node, err
	} else {
		return nil, err
	}
}

// SetNodeViaNid Set Node via Nid
func SetNodeViaNid(iNid int64, node *Node) (int64, error) {
	node.Nid = iNid
	return Engine.Insert(node)
}

// SetNodeViaVid Set Node via Vid
func SetNodeViaVid(iVid int, node *Node) (int64, error) {
	node.Vid = iVid
	return Engine.Insert(node)
}

// SetNodeViaType Set Node via Type
func SetNodeViaType(iType string, node *Node) (int64, error) {
	node.Type = iType
	return Engine.Insert(node)
}

// SetNodeViaUuid Set Node via Uuid
func SetNodeViaUuid(iUuid string, node *Node) (int64, error) {
	node.Uuid = iUuid
	return Engine.Insert(node)
}

// SetNodeViaLangcode Set Node via Langcode
func SetNodeViaLangcode(iLangcode string, node *Node) (int64, error) {
	node.Langcode = iLangcode
	return Engine.Insert(node)
}

// AddNode Add Node via all columns
func AddNode(iNid int64, iVid int, iType string, iUuid string, iLangcode string) error {
	_Node := &Node{Nid: iNid, Vid: iVid, Type: iType, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_Node); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNode Post Node via iNode
func PostNode(iNode *Node) (int64, error) {
	_, err := Engine.Insert(iNode)
	return iNode.Nid, err
}

// PutNode Put Node
func PutNode(iNode *Node) (int64, error) {
	_, err := Engine.Id(iNode.Nid).Update(iNode)
	return iNode.Nid, err
}

// PutNodeViaNid Put Node via Nid
func PutNodeViaNid(Nid_ int64, iNode *Node) (int64, error) {
	row, err := Engine.Update(iNode, &Node{Nid: Nid_})
	return row, err
}

// PutNodeViaVid Put Node via Vid
func PutNodeViaVid(Vid_ int, iNode *Node) (int64, error) {
	row, err := Engine.Update(iNode, &Node{Vid: Vid_})
	return row, err
}

// PutNodeViaType Put Node via Type
func PutNodeViaType(Type_ string, iNode *Node) (int64, error) {
	row, err := Engine.Update(iNode, &Node{Type: Type_})
	return row, err
}

// PutNodeViaUuid Put Node via Uuid
func PutNodeViaUuid(Uuid_ string, iNode *Node) (int64, error) {
	row, err := Engine.Update(iNode, &Node{Uuid: Uuid_})
	return row, err
}

// PutNodeViaLangcode Put Node via Langcode
func PutNodeViaLangcode(Langcode_ string, iNode *Node) (int64, error) {
	row, err := Engine.Update(iNode, &Node{Langcode: Langcode_})
	return row, err
}

// UpdateNodeViaNid via map[string]interface{}{}
func UpdateNodeViaNid(iNid int64, iNodeMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node)).Where("nid = ?", iNid).Update(iNodeMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeViaVid via map[string]interface{}{}
func UpdateNodeViaVid(iVid int, iNodeMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node)).Where("vid = ?", iVid).Update(iNodeMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeViaType via map[string]interface{}{}
func UpdateNodeViaType(iType string, iNodeMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node)).Where("type = ?", iType).Update(iNodeMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeViaUuid via map[string]interface{}{}
func UpdateNodeViaUuid(iUuid string, iNodeMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node)).Where("uuid = ?", iUuid).Update(iNodeMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeViaLangcode via map[string]interface{}{}
func UpdateNodeViaLangcode(iLangcode string, iNodeMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node)).Where("langcode = ?", iLangcode).Update(iNodeMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNode Delete Node
func DeleteNode(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(Node))
	return row, err
}

// DeleteNodeViaNid Delete Node via Nid
func DeleteNodeViaNid(iNid int64) (err error) {
	var has bool
	var _Node = &Node{Nid: iNid}
	if has, err = Engine.Get(_Node); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(Node)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeViaVid Delete Node via Vid
func DeleteNodeViaVid(iVid int) (err error) {
	var has bool
	var _Node = &Node{Vid: iVid}
	if has, err = Engine.Get(_Node); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(Node)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeViaType Delete Node via Type
func DeleteNodeViaType(iType string) (err error) {
	var has bool
	var _Node = &Node{Type: iType}
	if has, err = Engine.Get(_Node); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(Node)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeViaUuid Delete Node via Uuid
func DeleteNodeViaUuid(iUuid string) (err error) {
	var has bool
	var _Node = &Node{Uuid: iUuid}
	if has, err = Engine.Get(_Node); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(Node)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeViaLangcode Delete Node via Langcode
func DeleteNodeViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Node = &Node{Langcode: iLangcode}
	if has, err = Engine.Get(_Node); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Node)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
