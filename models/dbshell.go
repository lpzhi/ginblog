package models

import (
	"time"
)

var (
	tableLists  = [...]string{"log_actors","log_install","log_dau","log_economy","log_loading","log_online"}
)
type dbshell struct {
	Model
}

func CreateRoleTotalDatabase() {
	//sql := "create database role_total charset utf8"
	////sql := "select * from blog_tag"
	//rs := db.Exec(sql)
	sq := "use role_total"
	db.Exec(sq)
	createRoleTables()
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

func createRoleTables()  {
	start := time.Now() //.Format("20060102 ")
	oneYearLater :=time.Now().AddDate(1,0,0) //.Format("20060102 ")
	dateArr := getPrDates(start,oneYearLater,"20060102")
	for _,v :=range tableLists {
		for _,v1 :=range dateArr{
			sql :=getRoleTotalSql(v,v1)
			db.Exec(sql)
		}
	}
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
		break
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
		break
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
		break
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
		break
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
		break
	case "log_online":
		sql =  "CREATE TABLE log_online_"+i+
			"(`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
			"`logdate` datetime NULL  COMMENT '记录时间'," +
			"`pfrom_id` smallint(5) NOT NULL DEFAULT '0' COMMENT '平台标识'," +
			"`server_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '区服ID'," +
			"`count` int(10) NOT NULL DEFAULT '0' COMMENT '在线人数'," +
			"PRIMARY KEY (`id`),KEY `logdate` (`logdate`,`pfrom_id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ";
	default:
		break
	}

	return sql
}