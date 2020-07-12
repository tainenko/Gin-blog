package db
func createBlogTagTable(){
	DB.Query(`
			CREATE TABLE IF NOT EXISTS blog_tag(
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(100) DEFAULT ""
			created_on int(10) unsigned DEFAULT "0"
			modified_on int(10) unsigned DEFAULT "0"
			modified_by varchar(100) DEFAULT ""
			state tinyint(3) unsigned DEFAULT "1"
			PRIMARY KEY("id")
		)`)
}