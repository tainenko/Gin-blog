package db

func createBlogTagTable() {
	DB.Query(`
			CREATE TABLE IF NOT EXISTS blog_tag(
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(100) DEFAULT "",
			created_on int(10) unsigned DEFAULT "0",
			modified_on int(10) unsigned DEFAULT "0",
			modified_by varchar(100) DEFAULT "",
			state tinyint(3) unsigned DEFAULT "1",
			PRIMARY KEY("id")
		)`)
}

func createBlogArticleTable() {
	DB.Query(`
			CREATE TABLE IF NOT EXISTS blog_tag(
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			tag_id int(10) unsigned DEFAULT "0",
			title varchar(100) DEFAULT "",
			desc varc har(255) DEFAULT "",
			content text,
			created_on int(11) DEFAULT NULL,
			created_by varchar(100) DEFAULT""
			modified_on int(10) unsigned DEFAULT "0",
			modified_by varchar(255) DEFAULT"",
			state tinyint(3) unsigned DEFAULT "1",
			PRIMARY KEY("id"))`
	)
}

func createBlogAuthTable() {
	DB.Query(`
			CREATE TABLE IF NOT EXISTS blog_auth(
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			username varchar(50) DEFAULT "",
			password varchar(50) DEFAULT "",
			PRIMARY KEY("id")
			)`)
}
