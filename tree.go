package helper

import (
	"errors"
	"reflect"
)

type tree struct {
	// 节点id字段
	IDField string
	// 父节点id字段
	PIDField string
	// list的map形式
	ListMap []map[string]interface{}
	// list struct的field tag，默认json tag，
	FieldTagMap map[string]string
}

// New
// IDField 节点id字段
// PIDField 父节点字段
func NewTree(IDField string, PIDField string) *tree {
	t := &tree{}
	t.FieldTagMap = make(map[string]string)
	t.IDField = "ID"
	t.PIDField = "PID"
	if IDField != "" {
		t.IDField = IDField
	}
	if PIDField != "" {
		t.PIDField = PIDField
	}
	return t
}

// ListToTree 列表转树结构
// parentID 父级id
// list 列表数据
//
// 返回treeJson
func (tree *tree) ListToTree(parentID interface{}, list interface{}) (listMapTree []map[string]interface{}, err error) {
	var fieldNames []string
	// 判断类型是否为array或slice
	k := reflect.ValueOf(list).Kind().String()
	if k != "array" && k != "slice" {
		return nil, errors.New("list is not array or slice")
	}

	// 判断类型及字段是否存在
	t := reflect.TypeOf(list)
	tEl := t.Elem()

	if _, ok := tEl.FieldByName(tree.IDField); !ok {
		return nil, errors.New("IDfield does not exist")
	}
	if _, ok := tEl.FieldByName(tree.PIDField); !ok {
		return nil, errors.New("PIDField does not exist")
	}
	// 获取field->tag
	for i := 0; i < tEl.NumField(); i++ {
		// 获取去fieldName的jsonTag
		// 如果为空则使用fieldName
		key := ""
		if tEl.Field(i).Tag.Get("json") != "" {
			key = tEl.Field(i).Tag.Get("json")
		} else {
			key = tEl.Field(i).Name
		}
		tree.FieldTagMap[tEl.Field(i).Name] = key
		// 获取fieldName
		fieldNames = append(fieldNames, tEl.Field(i).Name)
	}
	// 获取值
	v := reflect.ValueOf(list)
	for i := 0; i < v.Len(); i++ {
		// 计算structMap
		structMap := make(map[string]interface{})
		for _, fieldName := range fieldNames {
			structMap[tree.FieldTagMap[fieldName]] = v.Index(i).FieldByName(fieldName).Interface()
		}
		// 计算listMap
		tree.ListMap = append(tree.ListMap, structMap)
	}
	// 递归处理
	listMapTree = tree.recursionElement(tree.ListMap, parentID)
	return listMapTree, nil
}

// recursionElement 递归元素
// 先从跟pid开始找其子集
func (tree *tree) recursionElement(
	listMap []map[string]interface{},
	parentID interface{},
) []map[string]interface{} {
	listMapNew := make([]map[string]interface{}, 0)
	// 创建副本，防止影响原始listMap
	listMapCopy := make([]map[string]interface{}, len(listMap))
	copy(listMapCopy, listMap)
	// 遍历查找pid的children
	// listMap的key，使用的是jsonTag
	for index, el := range listMap {
		if el[tree.FieldTagMap[tree.PIDField]] == parentID {
			listMapNew = append(listMapNew, el)
			// 过滤掉已经属于pid的元素
			var lastList []map[string]interface{}
			if index < len(listMapCopy)-1 {
				lastList = listMapCopy[index+1:]
			}
			listMapCopy = append(listMapCopy[0:index], lastList...)
			// 如果是pid的子集，递归判断其剩余
			children := tree.recursionElement(listMapCopy, el[tree.FieldTagMap[tree.IDField]])
			if len(children) > 0 {
				listMap[index]["children"] = children
			}
		}
	}
	return listMapNew
}
