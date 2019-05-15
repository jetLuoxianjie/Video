package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video/api/defs"
	"video/api/utils"
)



func AddUserCredential(loginName string,pwd string) error{
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string,error){
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("Delete user error: %s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int ,name string)(*defs.VideoInfo,error){
	vid, err:=utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t:=time.Now()
	ctime:=t.Format("Jan 02 2006, 15:04:05")

	fmt.Println(t)
	fmt.Println(ctime)

	stmtIns, err:=dbConn.Prepare(`INSERT INTO video_info
		(id, author_id, name, display_ctime) VALUES(?,?,?,?)`)
	if err!=nil {
		return nil, err
	}
	_, err=stmtIns.Exec(vid, aid, name, ctime)
	if err!=nil {
		return nil, err
	}
	res:=&defs.VideoInfo{Id: vid, AuthorId: aid, Name:name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil

}
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err:= dbConn.Prepare(`SELECT author_id, name, display_ctime FROM video_info 
		WHERE id=?`)
	if err!=nil {
		return nil, err
	}
	var aid int
	var dct string
	var name string
	err=stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err!=nil && err!=sql.ErrNoRows {
		return nil, err
	}
	if err==sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res:=&defs.VideoInfo{Id: vid, AuthorId: aid, Name:name, DisplayCtime:dct}
	return res, nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err:=dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name=? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time<=FROM_UNIXTIME(?)
		OREDER BY video_info.create_time DESC`)
	var res []*defs.VideoInfo
	if err!=nil{
		return res, err
	}
	rows, err:=stmtOut.Query(uname, from, to)
	if err!=nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err:=rows.Scan(&id, &aid, &name, &ctime); err!=nil{
			return res, err
		}
		vi:=&defs.VideoInfo{Id:id, AuthorId:aid, Name:name, DisplayCtime: ctime}
		res = append(res, vi)
	}
	defer stmtOut.Close()
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err!=nil {
		return err
	}
	_, err=stmtDel.Exec(vid)
	if err!=nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}
