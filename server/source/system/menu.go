package system

import (
	"context"

	. "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []SysBaseMenu{
		// 首页 / 关于 / 个人信息 / 安全日志
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "首页", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 99, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 100, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "operation", Name: "operation", Component: "view/operationRecord/index.vue", Sort: 98, Meta: Meta{Title: "安全日志", Icon: "pie-chart"}},

		// 超级管理员
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 97, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "5", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "5", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 2, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "5", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 3, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "5", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 4, Meta: Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},

		// 服务器状态
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		// 系统工具
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 6, Meta: Meta{Title: "系统工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: true, ParentId: "11", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "自动化package", Icon: "folder"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "官方网站", Icon: "home-filled"}},
		// 插件系统
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 7, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "autoPlug", Name: "autoPlug", Component: "view/systemTools/autoPlug/autoPlug.vue", Sort: 2, Meta: Meta{Title: "插件模板", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "files"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "邮件插件", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "chatTable", Name: "chatTable", Component: "view/chatgpt/chatTable.vue", Sort: 6, Meta: Meta{Title: "万用表格", Icon: "chat-dot-square"}},
		// 示例文件
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "断点续传", Icon: "upload-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},

		// 测试管理
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "testMgr", Name: "testMgr", Component: "view/testMgr/index.vue", Sort: 2, Meta: Meta{Title: "测试管理", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "29", Path: "testCase", Name: "testCase", Component: "view/testCase/index.vue", Sort: 1, Meta: Meta{Title: "测试用例", Icon: "list"}},
		{MenuLevel: 0, Hidden: false, ParentId: "29", Path: "testSuite", Name: "testSuite", Component: "view/testSuite/index.vue", Sort: 2, Meta: Meta{Title: "测试套件", Icon: "list"}},
		{MenuLevel: 0, Hidden: false, ParentId: "29", Path: "testTask", Name: "testTask", Component: "view/testTask/index.vue", Sort: 3, Meta: Meta{Title: "测试任务", Icon: "list"}},
		{MenuLevel: 0, Hidden: true, ParentId: "29", Path: "testTaskDetail/:id", Name: "testTaskDetail", Component: "view/testTask/detail/index.vue", Sort: 4, Meta: Meta{Title: "任务详情-${id}", Icon: "list"}},
		{MenuLevel: 0, Hidden: false, ParentId: "29", Path: "testReport", Name: "testReport", Component: "view/testReport/index.vue", Sort: 5, Meta: Meta{Title: "测试报告", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: false, ParentId: "29", Path: "dataMgr", Name: "dataMgr", Component: "view/dataMgr/index.vue", Sort: 6, Meta: Meta{Title: "数据管理", Icon: "upload"}},
		// 实验室管理
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "labMgr", Name: "labMgr", Component: "view/labMgr/index.vue", Sort: 3, Meta: Meta{Title: "实验室管理", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "36", Path: "labHardware", Name: "labHardware", Component: "view/labMgr/hardware/index.vue", Sort: 1, Meta: Meta{Title: "物料管理", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "36", Path: "testEnv", Name: "testEnv", Component: "view/labMgr/testEnv/index.vue", Sort: 2, Meta: Meta{Title: "环境管理", Icon: "management"}},
		// 配置管理
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "cfgMgr", Name: "cfgMgr", Component: "view/cfgMgr/index.vue", Sort: 4, Meta: Meta{Title: "配置管理", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: "39", Path: "system", Name: "system", Component: "view/cfgMgr/system/index.vue", Sort: 1, Meta: Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "39", Path: "const", Name: "const", Component: "view/cfgMgr/const/sysDictionary.vue", Sort: 2, Meta: Meta{Title: "全局常量", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: true, ParentId: "39", Path: "constDetail/:id", Name: "constDetail", Component: "view/cfgMgr/const/sysDictionaryDetail.vue", Sort: 3, Meta: Meta{Title: "全局常量详情-${id}", Icon: "list", ActiveName: "dictionary"}},
		// 工单系统
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "order", Name: "order", Component: "view/order/index.vue", Sort: 5, Meta: Meta{Title: "工单系统", Icon: "management"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "autoPkg").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
