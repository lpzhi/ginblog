package models

import (
	"fmt"
	"strconv"
	"time"
)

var (
	tableLists  = [...]string{"log_actors","log_install","log_dau","log_economy","log_loading","log_online"}
	enterLogTableLists = [...]string{"log_payment","account_info","t_server_config","t_agent_config"}
)
type dbshell struct {
	Model
}

func CreateEnterLogDatabase(pf string) error {
	var database string
	database = "enter_log_"+pf
	sql := "create database "+ database +" charset utf8"
	fmt.Println(sql)
	db := db.Exec(sql)
	useSql := "use "+database
	db.Exec(useSql)
	return db.Error
}

func CreateEnterLogDatabaseTables(pf string) error  {
	//创建数据库
	if err := CreateEnterLogDatabase(pf);err!=nil{
		return err
	}
	var err error

	//创建数据表
	for  _,v := range enterLogTableLists{
		if v=="account_info"{
			for i:=0;i<=256;i++ {
				sql := getRoleTotalSql(v,strconv.Itoa(i))
				if db:=db.Exec(sql);db.Error!=nil{
					err = db.Error
				}
			}
		}else {

			sql := getRoleTotalSql(v,"0")
			if db:=db.Exec(sql);db.Error!=nil{
				err = db.Error
			}
		}
	}

	return err
}

func CreateRoleTotalDatabase() error {
	//sql := "create database role_total charset utf8"
	////sql := "select * from blog_tag"
	//rs := db.Exec(sql)
	sq := "use role_total"
	db.Exec(sq)
	err := createRoleTables()
	return  err
}

func getPrDates(start,end time.Time , format string)   []string{
	var dateArr [] string
	var dtStart,dtEnd  time.Time

	dtStart = start
	dtEnd = end

	for {
		if dtStart.Before(dtEnd) {
			dateArr = append(dateArr,dtStart.Format(format))
			dtStart = dtStart.AddDate(0,0,1)
		}else {
			break
		}
	}

	return dateArr
}

func createRoleTables()  error {
	var err error

	start := time.Now() //.Format("20060102 ")
	oneYearLater :=time.Now().AddDate(1,0,0) //.Format("20060102 ")
	dateArr := getPrDates(start,oneYearLater,"20060102")
	for _,v :=range tableLists {
		for _,v1 :=range dateArr{
			sql :=getRoleTotalSql(v,v1)
			if db :=db.Exec(sql);db.Error!=nil{
				err = db.Error
			}
		}
	}

	return err
}


func getRoleTotalSql(table,i string) string {
	var sql string

	switch table {

	case "log_actors":
		sql = "CREATE TABLE log_actors_"+i+ "(" +
			"id int(10) NOT NULL AUTO_INCREMENT,pfrom_id smallint(5) NOT NULL DEFAULT '0' COMMENT '平台标示'," +
			"server_id int(11) DEFAULT '0' COMMENT '玩家所在的服务器的编号'," +
			"accountid int(11) DEFAULT NULL COMMENT '账户ID'," +
			"actorid int(11) NOT NULL COMMENT '角色ID'," +
			"actorname varchar(32) DEFAULT NULL COMMENT '玩家名字'," +
			"sex int(11) DEFAULT '0' COMMENT '性别'," +
			"level int(11) DEFAULT '0' COMMENT '等级'," +
			"job int(11) DEFAULT '0' COMMENT '职业'," +
			"yuanbao int(11) DEFAULT '0' COMMENT '非绑定元宝'," +
			"createtime datetime DEFAULT NULL COMMENT '角色的创建时间'," +
			"totalpower int(11) DEFAULT '0' COMMENT '战斗力'," +
			"recharge int(11) DEFAULT '0' COMMENT '充值元宝数'," +
			"accountname varchar(80) DEFAULT NULL COMMENT '账户的名字'," +
			"totalonline int(10) unsigned DEFAULT '0' COMMENT '总的在线时间（秒）'," +
			"dayilyonline int(10) unsigned DEFAULT '0' COMMENT '单日在线时间（秒）'," +
			"updatetime datetime NOT NULL COMMENT '最后更新时间'," +
			"turnlevel tinyint(2) NOT NULL DEFAULT 0 COMMENT '转生等级'," +
			"PRIMARY KEY (`id`)," +
			"KEY pfrom_id (`pfrom_id`)," +
			"KEY server_id (`server_id`)," +
			"KEY createtime (`createtime`)," +
			"KEY updatetime (`updatetime`)) " +
			"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='玩家的角色表' ";
	case "log_install":
		sql = "CREATE TABLE log_install_"+i+
			"(`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`regdate` datetime  NULL  COMMENT '注册时间'," +
			"`logindate` datetime  NULL   COMMENT '登录时间'," +
			"`account` varchar(64) NOT NULL DEFAULT '' COMMENT '注册用户ID'," +
			"`pf` varchar(32) NOT NULL DEFAULT '' COMMENT '渠道'," +
			"`via` varchar(32) NOT NULL DEFAULT '' COMMENT 'via参数'," +
			"`server_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '服务区号'," +
			"`isnewaccount`  tinyint NOT NULL DEFAULT 0 COMMENT '是否新账号0表示新账号'," +
			"PRIMARY KEY (`id`),KEY `regdate` (`regdate`)," +
			"KEY `isnewaccount` (`isnewaccount`)," +
			"KEY `logdate` (`logindate`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	case "log_dau":
		sql = "CREATE TABLE  log_dau_"+i+
			"(`id` int(10) unsigned NOT NULL AUTO_INCREMENT, " +
			"`logindate` datetime  NULL  COMMENT '登录时间'," +
			"`account` varchar(64) NOT NULL DEFAULT '' COMMENT '注册用户ID'," +
			"`pf` varchar(32) NOT NULL DEFAULT '' COMMENT '渠道'," +
			"`via` varchar(32) NOT NULL DEFAULT '' COMMENT 'via参数'," +
			"`server_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '服务区号'," +
			"`is_new` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否新账号'," +
			"`regdate` datetime NULL  COMMENT '注册时间'," +
			"PRIMARY KEY (`id`), KEY `pf` (`pf`)," +
			" KEY `account` (`account`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	case "log_economy":
		sql = "CREATE TABLE  log_economy_"+i+ "(" +
			"`id` int(10) NOT NULL AUTO_INCREMENT," +
			" `logdate` datetime  NULL  COMMENT '记录时间', " +
			"`actorid` bigint(10) DEFAULT NULL," +
			" `currency` varchar(32) DEFAULT NULL," +
			" `amount` varchar(32) DEFAULT NULL," +
			" `value` varchar(32) DEFAULT NULL," +
			" `kingdom` varchar(32) DEFAULT NULL," +
			" `phylum` varchar(32) DEFAULT NULL," +
			" `classfield` varchar(32) DEFAULT NULL, " +
			"`family` varchar(32) DEFAULT NULL, " +
			"`genus` varchar(32) DEFAULT NULL," +
			" `extra` varchar(128) DEFAULT NULL, " +
			"`extra2` varchar(128) DEFAULT NULL, " +
			"`serverid` int(11) DEFAULT '0'," +
			" `pfrom_id` smallint(5) NOT NULL DEFAULT '0' COMMENT '平台标示'," +
			"PRIMARY KEY (`id`)," +
			" KEY `logdate` (`logdate`,`currency`,`kingdom`,`phylum`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	case "log_loading":
		sql = "CREATE TABLE log_loading_"+i+ "" +
			" (`id` int(10) NOT NULL AUTO_INCREMENT," +
			"`pfrom_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台标示'," +
			"`server_id` smallint(5) NOT NULL DEFAULT '0' COMMENT '区服id'," +
			"`account` varchar(64) NOT NULL DEFAULT '' COMMENT '平台账号'," +
			"`counter` varchar(20) NOT NULL DEFAULT '' COMMENT '日志类型'," +
			"`kingdom` varchar(32) NOT NULL DEFAULT '' COMMENT '打点位置'," +
			"`is_new` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否新用户'," +
			"`exts` varchar(32) NOT NULL DEFAULT '' COMMENT 'ios为1,Andir为2'," +
			"`ip` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆ip ip2long存入'," +
			"`logdate` datetime  NULL  COMMENT '记录时间',," +
			"`channel`  varchar(50) NOT NULL DEFAULT '' COMMENT '渠道'," +
			"`level`  varchar(50) NOT NULL DEFAULT '' COMMENT '渠道'," +
			"PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	case "log_online":
		sql =  "CREATE TABLE log_online_"+i+
			"(`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`logdate` datetime NULL  COMMENT '记录时间'," +
			"`pfrom_id` smallint(5) NOT NULL DEFAULT '0' COMMENT '平台标识'," +
			"`server_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '区服ID'," +
			"`count` int(10) NOT NULL DEFAULT '0' COMMENT '在线人数'," +
			"PRIMARY KEY (`id`),KEY `logdate` (`logdate`,`pfrom_id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	case "t_agent_config":
		sql = "CREATE TABLE `t_agent_config` (`" +
			"id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`pfrom_id` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '代理商id'," +
			"`pfrom_name` varchar(64) NOT NULL DEFAULT '' COMMENT '代理商名'," +
			"`login_key` varchar(32) NOT NULL DEFAULT '' COMMENT '登录key'," +
			"`pay_key` varchar(32) NOT NULL DEFAULT '' COMMENT '充值key'," +
			"`home_url` varchar(255) NOT NULL DEFAULT '' COMMENT '官网地址'," +
			"`pay_url` varchar(255) NOT NULL DEFAULT '' COMMENT '充值地址'," +
			"`bbs_url` varchar(255) NOT NULL DEFAULT '' COMMENT '论坛地址'," +
			"`gm_url` varchar(255) NOT NULL DEFAULT '' COMMENT '客服地址'," +
			"`fcm_url` varchar(255) NOT NULL DEFAULT '' COMMENT '防沉迷地址'," +
			"`card_url` varchar(255) NOT NULL DEFAULT '' COMMENT '卡码地址'," +
			"`bi_server` varchar(255) NOT NULL DEFAULT '' COMMENT '上报地址'," +
			"`srvaddr` varchar(30) NOT NULL DEFAULT '' COMMENT '服务端loginserver'," +
			"`login_port` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '服务端loginserver端口'," +
			"`server_list_url` varchar(255) NOT NULL DEFAULT '' COMMENT '区服列表地址',`ip_pay_list` text COMMENT 'IP白名单'," +
			"`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态0新建1可用2审批中'," +
			"`desc` varchar(64) DEFAULT '' COMMENT '备注',PRIMARY KEY (`id`),UNIQUE KEY `pfid_pfname` (`pfrom_id`,`pfrom_name`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8"
	case "t_server_config":
		sql = "CREATE TABLE `t_server_config` (`" +
			"id` int(10) unsigned NOT NULL auto_increment," +
			"`server_id` int(10) unsigned NOT NULL default '0' COMMENT '区服ID'," +
			"`pfrom_id` smallint(5) unsigned NOT NULL default '0' COMMENT '平台ID'," +
			"`ip` varchar(128) default '' COMMENT '游戏数据库ip'," +
			"`port` int(11) NOT NULL default '80',`dbuser` varchar(50) default 'jzwc1'," +
			"`dbpwd` varchar(100) default '0987abc123',`dbname` varchar(50) default '' COMMENT 'actor库名',`md5` varchar(50) default NULL," +
			"`onlinedate` date default NULL COMMENT '开服时间',`pre_onlinedate` datetime DEFAULT NULL COMMENT '预开服时间'," +
			"`available` tinyint(1) NOT NULL default '2' COMMENT '0合服1现服2未开3特殊服'," +
			"`entranceUrl` varchar(200) default NULL COMMENT '入口地址',`name` varchar(20) default NULL COMMENT '服务器名称'," +
			"`dblogname` varchar(50) default NULL COMMENT '日志库的名称',`mergedate` date default NULL COMMENT '合服时间'," +
			"`merge_id` int(10) default '0' COMMENT '合服ID',`route_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip地址'," +
			"`route_port` varchar(20) NOT NULL DEFAULT '' COMMENT '网关端口'," +
			"`server_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0:无状态 1:新服 2:火爆'," +
			"PRIMARY KEY  (`id`)," +
			"UNIQUE KEY `server_id_pfrom` (`server_id`,`pfrom_id`)) " +
			"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8";
	case "account_info":
		fmt.Println(i)
		sql = "CREATE TABLE account_info_" +i+ "(" +
			"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`pf` varchar(16) NOT NULL DEFAULT '' COMMENT '平台'," +
			"`via` varchar(32) NOT NULL DEFAULT '' COMMENT '子渠道'," +
			"`regdate` datetime  NULL  COMMENT '帐号创建时间'," +
			"`logindate` datetime  NULL  COMMENT '登录时间'," +
			"`account` varchar(64) NOT NULL DEFAULT '' COMMENT '平台账号'," +
			"`serverid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '服务区号'," +
			"`is_paid` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否有充值'," +
			"`closed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被封禁'," +
			"PRIMARY KEY (`id`)," +
			"KEY `account` (`account`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8";
	case "log_payment":
		sql = "CREATE TABLE `log_payment` (" +
			"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`pf` varchar(32) NOT NULL DEFAULT '' COMMENT '平台标识'," +
			"`logdate` datetime NOT NULL COMMENT '系统时间'," +
			"`p_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '平台传过来time'," +
			"`account` varchar(64) NOT NULL DEFAULT '' COMMENT '平台账号'," +
			"`actorid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色id'," +
			"`actorname` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名'," +
			"`actorlevel` smallint(3) NOT NULL DEFAULT '0' COMMENT '角色等级'," +
			"`glod_remain` int(11) NOT NULL DEFAULT '0' COMMENT '剩余元宝'," +
			"`amount` double(10,2) DEFAULT '0.00' COMMENT '充值金额'," +
			"`gold` int(6) unsigned NOT NULL DEFAULT '0' COMMENT '购买元宝数量'," +
			"`phylum` varchar(16) NOT NULL DEFAULT '' COMMENT '分类'," +
			"`transactionid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单号'," +
			"`provider` varchar(32) NOT NULL DEFAULT '' COMMENT '支付的提供商'," +
			"`via` varchar(32) NOT NULL DEFAULT '' COMMENT '子渠道'," +
			"`first_pay` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否首充'," +
			"`serverid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '服务区号'," +
			"`regdate` datetime  NULL COMMENT '创建时间',PRIMARY KEY (`id`)," +
			"UNIQUE KEY `transactionid` (`transactionid`,`first_pay`),KEY `logdate` (`logdate`)," +
			"KEY `pf` (`pf`),KEY `account` (`account`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8"
	default:
		break
	}

	return sql
}