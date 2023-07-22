// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"shira-chan-dev/ent/schema","Package":"shira-chan-dev/ent","Schemas":[{"name":"Order","config":{"Table":""},"edges":[{"name":"requester","type":"User","ref_name":"requested","unique":true,"inverse":true,"comment":"需求者"},{"name":"receiver","type":"User","annotations":{"EntGQL":{"OrderField":"RECEIVER_COUNT","RelayConnection":true}},"comment":"接单者"}],"fields":[{"name":"created_at","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"CREAT_AT","Type":"Int64"}},"comment":"创建时间"},{"name":"updated_at","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"UPDATED_AT","Type":"Int64"}},"comment":"更新时间"},{"name":"title","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":100,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TITLE"}},"comment":"标题"},{"name":"content","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":15000,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"内容"},{"name":"contact","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":20,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"联系方式"},{"name":"type","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"default":true,"default_value":"other","default_kind":24,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}},"comment":"故障类别"},{"name":"is_closed","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"IS_CLOSED"}},"comment":"是否被关闭"},{"name":"is_finished","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"IS_FINISHED"}},"comment":"是否完成"},{"name":"evaluation","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"validators":2,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"EVALUATION"}},"comment":"评分"},{"name":"hope_at","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"HOPE_AT","Type":"Int64"}},"comment":"期望时间"}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"EntSQL":{"table":"Orders"}}},{"name":"User","config":{"Table":""},"edges":[{"name":"requested","type":"Order","annotations":{"EntGQL":{"OrderField":"REQUESTED_COUNT","RelayConnection":true}},"comment":"需求"},{"name":"received","type":"Order","ref_name":"receiver","inverse":true,"annotations":{"EntGQL":{"OrderField":"RECEIVED_COUNT","RelayConnection":true}},"comment":"接单"}],"fields":[{"name":"created_at","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"CREAT_AT","Type":"Int64"}},"comment":"创建时间"},{"name":"updated_at","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"UPDATED_AT","Type":"Int64"}},"comment":"更新时间"},{"name":"uname","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":30,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"UNAME"}},"comment":"用户名"},{"name":"passwd","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"sensitive":true,"comment":"密码"},{"name":"phone","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":15,"unique":true,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"手机号码"},{"name":"is_admin","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"IS_ADMIN"}},"comment":"是否管理员"},{"name":"is_secretary","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"IS_SECRETARY"}},"comment":"是否部员"},{"name":"is_active","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"STATE"}},"comment":"用户状态"}],"hooks":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"EntSQL":{"table":"Users"}}}],"Features":["namedges","privacy","schema/snapshot","entql"]}`
