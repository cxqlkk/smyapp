package dao

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"smy/model"
	"smy/util"
)

type StudentDao struct {
}

var DefaultStudentDao  =&StudentDao{}
//添加学生
func (sd *StudentDao) AddStudent(stu model.Student) (int64, error) {
	db := OpenDb()
	defer db.Close()
	stmt, serr := db.Prepare("insert into tb_student (id,studentName,phoneNum,address,birthDay,status,createTime,updateTime,updated,created)" +
		"values(?,?,?,?,?,?,?,?,?,?)")
	if serr != nil {
		return 0, serr
	}
	n, _ := snowflake.NewNode(1)
	result, err := stmt.Exec(n.Generate().String(), stu.StudentName, stu.PhoneNum, stu.Address, stu.BirthDay, 1, stu.CreateTime, stu.UpdateTime, stu.Updated, stu.Created)
	if err != nil {
		DbError(err)
		return 0, err
	}

	return result.RowsAffected()
}

func (sd *StudentDao) LoginStudent(stu model.Student) (model.Student, bool) {
	var id, studentName, moto, ideal, favoritePerson, favoriteColor, heroMan string
	var phone string
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select id, phoneNum, studentName,  motto,ideal,favoritePerson,favoriteColor,heroMan " +
		"  from tb_student where phoneNum=? and studentName=? ")
	checkerr(err)

	row := stmt.QueryRow(stu.PhoneNum, stu.StudentName)
	//和名称无关，和顺序有关
	row.Scan(&id,&phone,&studentName, &moto, &ideal, &favoritePerson, &favoriteColor, &heroMan)
	return model.Student{
		Motto:          moto,
		Ideal:          ideal,
		FavoritePerson: favoritePerson,
		FavoriteColor:  favoriteColor,
		HeroMan:        heroMan,
		PhoneNum:phone,
		StudentName:studentName,
		ID:id,
	}, studentName != ""
}
func checkerr(e error) {
	if e != nil {
		fmt.Println(e)

	}
}
//分页查询总数
func (sd *StudentDao)Count(studentName string)int{
	var count int
	db:=OpenDb()
	defer db.Close()
	stmt,error:=db.Prepare("select count(1) count from tb_student where status =1 and studentName like concat('%',?,'%')" )
	DbError(error)
	stmt.QueryRow(studentName).Scan(&count)
	return count

}
//分页查询列表
func (sd *StudentDao)ListStudent(page *util.Page,studentName string)(students []model.Student){
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println(" 捕捉到 panic")
		}
	}()
	db:=OpenDb()
	defer db.Close()
	stmt,error:=db.Prepare("select address, id, phoneNum, studentName,  motto,ideal,favoritePerson,favoriteColor,heroMan " +
		"  from tb_student where status=1 and studentName like concat('%',?,'%') limit ?,?")
	defer stmt.Close()
	DbError(error)
	rows,err:=stmt.Query(studentName,(page.PageNo-1)*page.PageSize,page.PageSize)
	DbError(err)
	for rows.Next(){
		student:=model.Student{}
		rows.Scan(&student.Address,&student.ID,&student.PhoneNum,&student.StudentName,&student.Motto,&student.Ideal,&student.FavoritePerson,&student.FavoriteColor,&student.HeroMan)
		students=append(students,student)
		fmt.Println(student)
	}
	return students

}

func (sd *StudentDao)GetStudentById(id string)model.Student{
	db:=OpenDb()
	defer db.Close()
	stmt,error:=db.Prepare("select birthDay, address, id, phoneNum, studentName,  motto,ideal,favoritePerson,favoriteColor,heroMan " +
		"  from tb_student where status=1 and id=?")
	defer stmt.Close()
	DbError(error)
	student:=model.Student{}
	stmt.QueryRow(id).Scan(&student.BirthDay,&student.Address,&student.ID,&student.PhoneNum,&student.StudentName,&student.Motto,&student.Ideal,&student.FavoritePerson,&student.FavoriteColor,&student.HeroMan)

	return student

}

func ( sd *StudentDao)UpdateStudent(stu model.StudentVo)(int64,error){
	db:=OpenDb()
	defer db.Close()
	stmt,error:=db.Prepare("update tb_student set birthDay=?, address=?, phoneNum=?, studentName=?,  motto=?,ideal=?,favoritePerson=?,favoriteColor=?,heroMan=? where id=? ")
	defer stmt.Close()
	DbError(error)
	result,err:=stmt.Exec(stu.BirthDay,stu.Address,stu.PhoneNum,stu.StudentName,stu.Motto,stu.Ideal,stu.FavoritePerson,stu.FavoriteColor,stu.HeroMan,stu.ID)

	DbError(err)
	return result.RowsAffected()
}
func (sd *StudentDao)DeleteStudentById(id string)(int64,error){
	db:=OpenDb()
	defer db.Close()
	stmt,error:=db.Prepare("update tb_student set status =0 where id =?")
	DbError(error)
	defer stmt.Close()
	result,err:=stmt.Exec(id)
	DbError(err)
	return result.LastInsertId()
}