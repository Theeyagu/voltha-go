package model

import (
	"fmt"
	"github.com/opencord/voltha/protos/go/voltha"
	"reflect"
	"testing"
)

/*

 */
func Test_ChildType_01_CacheIsEmpty(t *testing.T) {
	if GetInstance().ChildrenFieldsCache != nil || len(GetInstance().ChildrenFieldsCache) > 0 {
		t.Errorf("GetInstance().ChildrenFieldsCache should be empty: %+v\n", GetInstance().ChildrenFieldsCache)
	}
	t.Logf("GetInstance().ChildrenFieldsCache is empty - %+v\n", GetInstance().ChildrenFieldsCache)
}

/*

 */
func Test_ChildType_02_Device_Proto_ChildrenFields(t *testing.T) {

	var cls *voltha.Device
	//cls = &voltha.Device{Id:"testing-Config-rev-id"}

	names := ChildrenFields(cls)

	tst := reflect.ValueOf(cls).Elem().FieldByName("ImageDownloads")

	fmt.Printf("############ Field by name : %+v\n", reflect.TypeOf(tst.Interface()))

	if names == nil || len(names) == 0 {
		t.Errorf("ChildrenFields failed to return names: %+v\n", names)
	}
	t.Logf("ChildrenFields returned names: %+v\n", names)
}

/*

 */
func Test_ChildType_03_CacheHasOneEntry(t *testing.T) {
	if GetInstance().ChildrenFieldsCache == nil || len(GetInstance().ChildrenFieldsCache) != 1 {
		t.Errorf("GetInstance().ChildrenFieldsCache should have one entry: %+v\n", GetInstance().ChildrenFieldsCache)
	}
	t.Logf("GetInstance().ChildrenFieldsCache has one entry: %+v\n", GetInstance().ChildrenFieldsCache)
}

/*

 */
func Test_ChildType_04_ChildrenFieldsCache_Keys(t *testing.T) {
	for k := range GetInstance().ChildrenFieldsCache {
		t.Logf("GetInstance().ChildrenFieldsCache Key:%+v\n", k)
	}
}
