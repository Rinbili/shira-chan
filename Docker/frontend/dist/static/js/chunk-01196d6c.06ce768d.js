(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-01196d6c"],{"0cbc":function(t,e,n){"use strict";n("1d26")},"1d26":function(t,e,n){},"333d":function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"block",staticStyle:{"margin-top":"10px",position:"absolute",bottom:"0",right:"0"}},[n("el-pagination",{attrs:{"current-page":t.currentPage,"page-sizes":[10,20,30,40],"page-size":t.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:t.total},on:{"update:currentPage":function(e){t.currentPage=e},"update:current-page":function(e){t.currentPage=e},"update:pageSize":function(e){t.pageSize=e},"update:page-size":function(e){t.pageSize=e},"size-change":t.SizeChange,"current-change":t.CurrentChange}})],1)},i=[],s=(n("a9e3"),{props:{page:{type:Number,default:1},size:{type:Number,default:10},total:{type:Number,default:10}},data:function(){return{currentPage:1,pageSize:10}},watch:{Page:function(){this.currentPage=this.page},Size:function(){this.pageSize=this.size}},mounted:function(){this.currentPage=this.page,this.pageSize=this.size},methods:{SizeChange:function(t){this.pageSize=t,this.currentPage=1,this.$emit("handlePageSizeChange",this.currentPage,this.pageSize)},CurrentChange:function(t){this.currentPage=t,this.$emit("handlePageCurrentChange",this.currentPage,this.pageSize)}}}),r=s,o=n("2877"),l=Object(o["a"])(r,a,i,!1,null,"582eac06",null);e["a"]=l.exports},caae:function(t,e,n){"use strict";n.r(e);var a,s,r,o,l,c=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("选择时间范围：")]),n("el-date-picker",{staticStyle:{"margin-right":"20px"},attrs:{type:"daterange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期",size:"small"},model:{value:t.query.timeRange,callback:function(e){t.$set(t.query,"timeRange",e)},expression:"query.timeRange"}}),n("span",[t._v("工单标题：")]),n("el-input",{staticStyle:{width:"200px","margin-right":"10px"},attrs:{placeholder:"请输入工单标题",size:"small"},model:{value:t.query.selectInput,callback:function(e){t.$set(t.query,"selectInput",e)},expression:"query.selectInput"}}),n("el-button",{attrs:{type:"primary",icon:"el-icon-search",size:"small"},on:{click:t.search}},[t._v("查询")]),n("el-button",{attrs:{icon:"el-icon-refresh-right",size:"small"},on:{click:t.reset}},[t._v("重置")]),n("el-button",{attrs:{type:"success",icon:"el-icon-refresh",size:"small"},on:{click:t.refresh}},[t._v("刷新")])],1),n("div",{staticClass:"center-table",staticStyle:{height:"550px",position:"relative"}},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.tableLoading,expression:"tableLoading"}],staticStyle:{width:"100%"},attrs:{data:t.formList,border:"","header-cell-style":{background:"#f3f6fd"},"max-height":"500","element-loading-background":"rgba(255, 255, 255, 0.5)","element-loading-text":"数据正在加载中","element-loading-spinner":"el-icon-loading"}},[n("el-table-column",{attrs:{align:"center",prop:"createdAt",label:"提交时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(new Date(1e3*e.row.createdAt).toLocaleString())+" ")]}}])}),n("el-table-column",{attrs:{align:"center",prop:"title",label:"标题"}}),n("el-table-column",{attrs:{align:"center",prop:"type",label:"故障类型"},scopedSlots:t._u([{key:"default",fn:function(e){return["0"===e.row.type?n("span",[t._v("启动故障")]):t._e(),"1"===e.row.type?n("span",[t._v("系统故障")]):t._e(),"2"===e.row.type?n("span",[t._v("蓝屏故障")]):t._e(),"3"===e.row.type?n("span",[t._v("黑屏故障")]):t._e(),"4"===e.row.type?n("span",[t._v("死机故障")]):t._e(),"5"===e.row.type?n("span",[t._v("重启故障")]):t._e(),"6"===e.row.type?n("span",[t._v("电脑服务")]):t._e()]}}])}),n("el-table-column",{attrs:{align:"center",prop:"hopeAt",label:"预约时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(new Date(1e3*e.row.hopeAt).toLocaleString())+" ")]}}])}),n("el-table-column",{attrs:{align:"center",prop:"contact",label:"联系电话"}}),n("el-table-column",{attrs:{align:"center",prop:"status",label:"工单当前状态"},scopedSlots:t._u([{key:"default",fn:function(e){return[1===e.row.status?n("span",{staticClass:"status-publish"},[t._v("已发布")]):2===e.row.status?n("span",{staticClass:"status-send"},[t._v("已派单")]):3===e.row.status?n("span",{staticClass:"status-success"},[t._v("已结单")]):n("span",{staticClass:"status-error"},[t._v("被拒单")])]}}])}),n("el-table-column",{attrs:{align:"center",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"success",icon:"el-icon-view",size:"small"},on:{click:function(n){return t.details(e.row.id)}}},[t._v("详情")])]}}])})],1),n("Pagination",{attrs:{total:Number(t.total),page:t.query.page,size:t.query.limit},on:{handlePageSizeChange:t.getCursor,handlePageCurrentChange:t.getCursor}})],1),n("el-dialog",{attrs:{title:"工单详情",visible:t.formDilog,top:"30px",center:""},on:{"update:visible":function(e){t.formDilog=e}}},[n("div",{staticStyle:{width:"100%","margin-bottom":"20px"}},[n("div",{staticStyle:{width:"100%","text-align":"center","margin-bottom":"20px","font-size":"20px","font-weight":"bold"}},[t._v(" 报修单当前状态 ")]),n("el-steps",{attrs:{active:this.form.status,"finish-status":"success","align-center":""}},[n("el-step",{attrs:{title:"已发布",description:"等待管理员派单"}}),n("el-step",{attrs:{title:"已派单",description:"维修员已经接单，请等待维修员与您联系"}}),n("el-step",{attrs:{title:"已完成",description:"报修结束"}})],1)],1),n("div",{staticStyle:{"margin-bottom":"20px"}},[n("div",{staticStyle:{width:"100%","text-align":"center","margin-bottom":"20px","font-size":"20px","font-weight":"bold"}},[t._v(" 报修单详情 ")]),n("el-form",{ref:"form",attrs:{model:t.form,"label-width":"80px"}},[n("el-row",{attrs:{gutter:20}},[n("el-col",{attrs:{span:12}},[n("el-form-item",{attrs:{label:"预约时间"}},[n("el-input",{attrs:{readonly:!0},model:{value:new Date(1e3*t.form.hopeAt).toLocaleString(),callback:function(e){t.$set(new Date(1e3*t.form.hopeAt),"toLocaleString()",e)},expression:"new Date(form.hopeAt * 1000).toLocaleString()"}})],1)],1),n("el-col",{attrs:{span:12}},[n("el-form-item",{attrs:{label:"工单标题"}},[n("el-input",{attrs:{disabled:!0},model:{value:t.form.title,callback:function(e){t.$set(t.form,"title",e)},expression:"form.title"}})],1)],1)],1),n("el-row",{attrs:{gutter:20}},[n("el-col",{attrs:{span:12}},[n("el-form-item",{attrs:{label:"故障类型"}},[n("el-input",{attrs:{disabled:!0},model:{value:t.form.type,callback:function(e){t.$set(t.form,"type",e)},expression:"form.type"}})],1)],1),n("el-col",{attrs:{span:12}},[n("el-form-item",{attrs:{label:"联系电话"}},[n("el-input",{attrs:{disabled:!0},model:{value:t.form.contact,callback:function(e){t.$set(t.form,"contact",e)},expression:"form.contact"}})],1)],1)],1),n("el-row",[n("el-form-item",{attrs:{label:"故障描述"}},[n("el-input",{attrs:{type:"textarea",rows:5,disabled:!0},model:{value:t.form.content,callback:function(e){t.$set(t.form,"content",e)},expression:"form.content"}})],1)],1)],1),n("div",[3===this.form.status?n("div",{staticStyle:{display:"inline-block",float:"left"}},[n("div",[t._v("请对此报修流程进行评价")]),n("el-rate",{attrs:{"show-text":""},model:{value:t.form.evaluation,callback:function(e){t.$set(t.form,"evaluation",e)},expression:"form.evaluation"}})],1):t._e(),3===this.form.status?n("el-button",{staticStyle:{float:"right"},attrs:{type:"success",size:"medium"},on:{click:t.submitEvaluation}},[t._v("确定")]):t._e()],1)],1),n("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(e){t.formDilog=!1}}},[t._v("取 消")]),n("el-button",{attrs:{type:"primary"},on:{click:function(e){t.formDilog=!1}}},[t._v("确 定")])],1)])],1)],1)},u=[],d=(n("d81d"),n("caad"),n("2532"),n("333d")),p=n("8785"),f=n("5184"),h=n("2866");function m(t){return h["a"].query({url:"query",method:"post",headers:{"Content-Type":"application/json"},query:Object(f["a"])(a||(a=Object(p["a"])(["\n      query (\n        $first: Int\n        $after: Cursor\n        $isFinished: Boolean\n        $isClosed: Boolean\n        $hopeAtGTE: Int64\n        $hopeAtLTE: Int64\n        $titleContains: String\n      ) {\n        orders(\n          first: $first\n          after: $after\n          where: {\n            isClosed: $isClosed\n            isFinished: $isFinished\n            hopeAtGTE: $hopeAtGTE\n            hopeAtLTE: $hopeAtLTE\n            titleContains: $titleContains\n          }\n        ) {\n          edges {\n            cursor\n          }\n          pageInfo {\n            startCursor\n            endCursor\n          }\n        }\n      }\n    "]))),variables:{first:t.limit*(t.page-1),after:t.after,hopeAtGTE:t.timeRange[0]/1e3,hopeAtLTE:t.timeRange[1]/1e3,titleContains:t.selectInput,isFinished:t.isFinished,isClosed:t.isClosed}})}function g(t){return h["a"].query({url:"query",method:"post",headers:{"Content-Type":"application/json"},fetchPolicy:"no-cache",query:Object(f["a"])(s||(s=Object(p["a"])(["\n      query (\n        $first: Int\n        $after: Cursor\n        $hopeAtGTE: Int64\n        $hopeAtLTE: Int64\n        $titleContains: String\n        $id: ID\n      ) {\n        orders(\n          first: $first\n          after: $after\n          where: {\n            hopeAtGTE: $hopeAtGTE\n            hopeAtLTE: $hopeAtLTE\n            titleContains: $titleContains\n            hasRequesterWith: { id: $id }\n          }\n        ) {\n          edges {\n            node {\n              id\n              title\n              contact\n              type\n              hopeAt\n              createdAt\n              isClosed\n              isFinished\n            }\n          }\n          totalCount\n        }\n      }\n    "]))),variables:{first:t.limit,after:t.after,hopeAtGTE:parseInt(t.timeRange[0]/1e3),hopeAtLTE:parseInt(t.timeRange[1]/1e3),titleContains:t.selectInput,id:t.id}})}function b(t){return h["a"].query({url:"query",method:"post",headers:{"Content-Type":"application/json"},fetchPolicy:"no-cache",query:Object(f["a"])(r||(r=Object(p["a"])(["\n      query ($id: ID) {\n        orders(where: { id: $id }) {\n          edges {\n            node {\n              id\n              title\n              content\n              contact\n              type\n              hopeAt\n              createdAt\n              evaluation\n              isClosed\n              isFinished\n            }\n          }\n          totalCount\n        }\n      }\n    "]))),variables:{id:t}})}function y(){return h["a"].query({url:"query",method:"post",headers:{"Content-Type":"application/json"},fetchPolicy:"no-cache",query:Object(f["a"])(o||(o=Object(p["a"])(["\n      query {\n        orders(where: { hasReceiver: true }) {\n          edges {\n            node {\n              id\n            }\n          }\n        }\n      }\n    "])))})}function v(t){return h["a"].mutate({url:"query",method:"post",headers:{"Content-Type":"application/json"},mutation:Object(f["a"])(l||(l=Object(p["a"])(["\n      mutation ($id: ID!, $evaluation: Float!) {\n        updateOrder(input: { evaluation: $evaluation }, id: $id) {\n          evaluation\n        }\n      }\n    "]))),variables:{id:t.id,evaluation:parseFloat(t.evaluation)}})}var C={components:{Pagination:d["a"]},data:function(){return{query:{page:1,limit:10,after:null,timeRange:"",selectInput:"",id:this.$store.state.user.id},formList:[],receiverFormList:[],formDilog:!1,tableLoading:!1,active:1,form:"",total:""}},created:function(){var t=this;y().then((function(e){for(var n=0;n<e.data.orders.edges.length;n++)t.receiverFormList.push(e.data.orders.edges[n].node.id)})),this.getRecordsList()},methods:{getRecordsList:function(){var t=this;this.tableLoading=!0,g(this.query).then((function(e){t.formList=e.data.orders.edges.map((function(t){return{id:t.node.id,createdAt:t.node.createdAt,hopeAt:t.node.hopeAt,title:t.node.title,type:t.node.type,contact:t.node.contact,evaluation:t.node.evaluation,isClosed:t.node.isClosed,isFinished:t.node.isFinished}})),t.total=e.data.orders.totalCount;for(var n=0;n<t.formList.length;n++)t.receiverFormList.includes(t.formList[n].id)?(t.formList[n].status=2,!0===t.formList[n].isFinished?t.formList[n].status=3:!0===t.formList[n].isClosed&&(t.formList[n].status=4)):t.formList[n].status=1;t.tableLoading=!1})),this.tableLoading=!1},getCursor:function(t,e){var n=this;this.query.page=t,this.query.limit=e,1!==this.query.page?(this.query.after=null,m(this.query).then((function(t){n.query.after=t.data.orders.pageInfo.endCursor,n.getRecordsList()}))):(this.query.after=null,this.getRecordsList())},search:function(){this.getRecordsList()},reset:function(){this.query.timeRange="",this.query.selectInput="",this.getRecordsList()},refresh:function(){this.getRecordsList()},details:function(t){var e=this;b(t).then((function(t){switch(e.form=t.data.orders.edges[0].node,e.form.type){case"0":e.form.type="启动故障";break;case"1":e.form.type="系统故障";break;case"2":e.form.type="蓝屏故障";break;case"3":e.form.type="黑屏故障";break;case"4":e.form.type="死机故障";break;case"5":e.form.type="重启故障";break;case"6":e.form.type="其他故障";break}e.receiverFormList.includes(e.form.id)?(e.form.status=2,!0===e.form.isFinished?e.form.status=3:!0===e.formList[i].isClosed&&(e.form.status=4)):e.form.status=1,e.formDilog=!0}))},submitEvaluation:function(){var t=this;v(this.form).then((function(e){e.data.updateOrder.evaluation?(t.$message.success("评价成功"),t.formDilog=!1):t.$message.error("评价失败")})).catch((function(){t.$message.error("评价失败")}))}}},$=C,w=(n("0cbc"),n("2877")),_=Object(w["a"])($,c,u,!1,null,"3c28d7bd",null);e["default"]=_.exports}}]);