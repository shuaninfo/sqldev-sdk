package sqldev

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// 如需测试，请将下面的常量替换为自己的
const (
	ENDPOINT   = `http://127.0.0.1:8080/openapi/v1`
	APP_KEY    = `a9967fcfa669c40c5df7db31a33f70d4`
	APP_SECRET = `5414aabf-db54-4f45-b7e1-1d99aa10e47c`
)

var sqldev = NewSqldev(ENDPOINT, APP_KEY, APP_SECRET, &http.Client{})

func TestGetUserInfo(t *testing.T) {
	a, err := sqldev.GetUserInfo(17)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestGetUserPage(t *testing.T) {
	a, total, err := sqldev.GetUserPage(&UserPageForm{
		PageNo:    0,
		PageSize:  100,
		GroupName: "测试部门",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
	fmt.Printf("%+v\n", total)
}

func TestAddUser(t *testing.T) {
	a, err := sqldev.AddUser(&UserAddForm{
		Name:      "账号必填11122222233",
		UserName:  "姓名必填11222123121",
		Pwd:       "mimabitian111@#¥",
		Expire:    -1,
		Role:      9,
		Sex:       1,
		Email:     "bitian@youxiang.com",
		Telno:     "13966661234",
		GroupName: "自动分组11",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestUpdUser(t *testing.T) {
	err := sqldev.UpdUser(&UserUpdForm{
		Id:        17,
		GroupName: "自动创建！",
		State:     0,
		Expire:    -1,
		TelNo:     "",
		RolesStr:  "",
		UserName:  "名称222",
		Role:      9,
		Telno:     "",
		Email:     "",
		Sex:       0,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestRemoveUser(t *testing.T) {
	err := sqldev.RemoveUser(15)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestUpdateUserState(t *testing.T) {
	err := sqldev.StateUser(&UserStateForm{
		Id:    14,
		State: 1,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestGetInstanceInfo(t *testing.T) {
	a, err := sqldev.GetInstanceInfo("3ab520c26bbd77f4b24db89aaaed71ec")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	b, _ := json.Marshal(a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

func TestGetInstancePage(t *testing.T) {
	a, total, err := sqldev.GetInstancePage(&InstancePageForm{
		PageNo:   0,
		PageSize: 100,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
	fmt.Printf("%+v\n", total)
}

func TestAddInstance(t *testing.T) {
	a, err := sqldev.AddInstance(&InstanceAddForm{
		Name:        "测试数据源名称1",
		User:        "ceshi",
		Pass:        "ceshimima",
		Ip:          "127.0.0.1",
		Port:        "6666",
		Db:          "testdb",
		ConType:     "base",
		ServiceType: "",
		UserType:    "",
		OperType:    4,
		DbType:      "mysql",
		AsynExport:  0,
		AsynQuery:   0,
		AsynSql:     0,
		PageQuery:   0,
		Config:      "",
		PoolMin:     1,
		PoolMax:     5,
		MaxRow:      100000,
		SshUse:      false,
		SshServer:   "",
		SshPort:     "",
		SshUser:     "",
		SshPassword: "",
		SshAuth:     0,
		TxTimeout:   30,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestUpdInstance(t *testing.T) {
	a, err := sqldev.UpdateInstance(&InstanceUpdForm{
		ID:          "f7e55dc7c669f79983565bcb42bd092c",
		Name:        "测试数据源名称2",
		User:        "ceshi",
		Pass:        "ceshimima",
		Ip:          "127.0.0.1",
		Port:        "6666",
		Db:          "testdb2",
		ConType:     "base",
		ServiceType: "",
		UserType:    "",
		OperType:    4,
		DbType:      "mysql",
		AsynExport:  0,
		AsynQuery:   0,
		AsynSql:     0,
		PageQuery:   0,
		Config:      "",
		PoolMin:     1,
		PoolMax:     5,
		MaxRow:      100000,
		SshUse:      false,
		SshServer:   "",
		SshPort:     "",
		SshUser:     "",
		SshPassword: "",
		SshAuth:     0,
		TxTimeout:   30,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestRemoveInstance(t *testing.T) {
	err := sqldev.RemoveInstance("3e1e5f8f59fa052dfd340aef932b164e")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestStateInstance(t *testing.T) {
	err := sqldev.StateInstance(&InstanceStateForm{
		ID:    "0587584ba4ee9ede2bf8f3eccab8ec53",
		State: 0,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestGetInstanceSync(t *testing.T) {
	a, err := sqldev.GetInstanceSync()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
}

func TestQuery(t *testing.T) {
	a, err := sqldev.Query(&SqlQueryForm{
		Project: 1,
		Iid:     "03af2e960d8736f7724484c263b257f4",
		Sql: `
select 
	t.id, t.sample 
from 
	test_table_1 t where t.id =1;

update test_table_1 t set t.sample = '777' where t.id = 1;

select 
	t.id, t.sample 
from 
	test_table_1 t where t.id=1;`,
		Owner:     "testdb",
		Schema:    "",
		PageIndex: 0,
		PageSize:  100,
		NeedTotal: true,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestGetWorkInfo(t *testing.T) {
	a, err := sqldev.GetWorkInfo("16915775542572")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestGetInstanceWhitelistList(t *testing.T) {
	a, err := sqldev.GetInstanceWhitelistList("3ab520c26bbd77f4b24db89aaaed71ec")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
}

func TestAddInstanceWhitelist(t *testing.T) {
	a, err := sqldev.AddInstanceWhitelist(&InstanceWhitelistAddForm{
		DbId: "3ab520c26bbd77f4b24db89aaaed71ec",
		Ip:   "127.0.0.1",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestUpdateInstanceWhitelist(t *testing.T) {
	a, err := sqldev.UpdateInstanceWhitelist(&InstanceWhitelistUpdForm{
		Id: 9,
		Ip: "127.0.0.1",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestRemoveInstanceWhitelist(t *testing.T) {
	err := sqldev.RemoveInstanceWhitelist(13)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestWhitelistFlow(t *testing.T) {
	dbId := "3ab520c26bbd77f4b24db89aaaed71ec"
	whitelistIP1 := "127.0.0.1"
	whitelistIP2 := "192.168.1.1"

	fmt.Println("查询当前数据库实例的白名单列表...")
	// Check initial whitelist
	whitelist, err := sqldev.GetInstanceWhitelistList(dbId)
	if err != nil {
		fmt.Printf("查询白名单失败: %v\n", err)
		return
	}
	if len(whitelist) == 0 {
		fmt.Println("当前没有设置白名单。")
	} else {
		for _, w := range whitelist {
			fmt.Printf("白名单: %+v\n", w)
		}
	}

	// Add first whitelist entry
	fmt.Println("添加第一个白名单...")
	_, err = sqldev.AddInstanceWhitelist(&InstanceWhitelistAddForm{
		DbId: dbId,
		Ip:   whitelistIP1,
	})
	if err != nil {
		fmt.Printf("添加白名单失败: %v\n", err)
		return
	}
	fmt.Printf("添加白名单 %s 成功。\n", whitelistIP1)

	// Add second whitelist entry
	fmt.Println("添加第二个白名单...")
	_, err = sqldev.AddInstanceWhitelist(&InstanceWhitelistAddForm{
		DbId: dbId,
		Ip:   whitelistIP2,
	})
	if err != nil {
		fmt.Printf("添加白名单失败: %v\n", err)
		return
	}
	fmt.Printf("添加白名单 %s 成功。\n", whitelistIP2)

	// Verify the whitelist entries
	fmt.Println("验证当前的白名单...")
	whitelist, err = sqldev.GetInstanceWhitelistList(dbId)
	if err != nil {
		fmt.Printf("查询白名单失败: %v\n", err)
		return
	}
	for _, w := range whitelist {
		fmt.Printf("白名单: %+v\n", w)
	}

	// Assuming we know the ID of the whitelist entry to update
	whitelistEntryToUpdate := whitelist[0].ID
	fmt.Printf("更新ID为 %d 的白名单...\n", whitelistEntryToUpdate)
	_, err = sqldev.UpdateInstanceWhitelist(&InstanceWhitelistUpdForm{
		Id: whitelistEntryToUpdate,
		Ip: "127.0.0.2", // New IP for the existing whitelist entry
	})
	if err != nil {
		fmt.Printf("更新白名单失败: %v\n", err)
		return
	}
	fmt.Printf("更新白名单 %d 成功。\n", whitelistEntryToUpdate)

	// Verify the updated whitelist entry
	fmt.Println("验证更新后的白名单...")
	whitelist, err = sqldev.GetInstanceWhitelistList(dbId)
	if err != nil {
		fmt.Printf("查询白名单失败: %v\n", err)
		return
	}
	for _, w := range whitelist {
		fmt.Printf("白名单: %+v\n", w)
	}

	// Assuming we know the ID of the whitelist entry to remove
	whitelistEntryToRemove := whitelist[1].ID // This ID needs to be obtained from previous response
	fmt.Printf("删除ID为 %d 的白名单...\n", whitelistEntryToRemove)
	err = sqldev.RemoveInstanceWhitelist(whitelistEntryToRemove)
	if err != nil {
		fmt.Printf("删除白名单失败: %v\n", err)
		return
	}
	fmt.Printf("删除白名单 %d 成功。\n", whitelistEntryToRemove)

	// Final verification to ensure the entry is removed
	fmt.Println("最终验证以确保白名单已被删除...")
	whitelist, err = sqldev.GetInstanceWhitelistList(dbId)
	if err != nil {
		fmt.Printf("查询白名单失败: %v\n", err)
		return
	}
	for _, w := range whitelist {
		fmt.Printf("白名单: %+v\n", w)
	}
}
