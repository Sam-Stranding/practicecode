package main

import (
	"context"
	"fmt"
	"gorm/model"
	"gorm/query"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/edu.mall?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}

	q := query.Use(db)
	ctx := context.Background()
	userDAD := q.User

	//1.创建用户
	newUser := &model.User{
		Name:  "张三",
		Age:   25,
		Email: "zhangsan@163.com",
	}
	err = userDAD.WithContext(ctx).Create(newUser)
	if err != nil {
		log.Fatal("创建用户失败：", err)
	}
	fmt.Println("创建用户成功：", newUser)

	////2.查询用户
	////2.1 根据ID查询
	//user, err := userDAD.WithContext(ctx).Where(q.User.ID.Eq(newUser.ID)).First()
	//if err != nil {
	//	log.Fatal("查询用户失败：", err)
	//}
	//fmt.Printf("查询结果：%+v\n", user)
	//
	////2.2 条件查询
	//users, err := userDAD.WithContext(ctx).
	//	Where(q.User.Age.Gte(20), q.User.Name.Like("%张%")).
	//	Find()
	//if err != nil {
	//	log.Fatal("条件查询失败：", err)
	//}
	//fmt.Printf("查询到{%d}个用户", len(users))
	//
	////3.更新用户
	//_, err = userDAD.WithContext(ctx).
	//	Where(q.User.ID.Eq(newUser.ID)).
	//	Update(q.User.Age, 26)
	//if err != nil {
	//	log.Fatal("更新用户年龄失败：", err)
	//}
	//fmt.Println("更新用户年龄成功")
	//
	//// 多字段更新
	//_, err = userDAD.WithContext(ctx).
	//	Where(q.User.ID.Eq(newUser.ID)).
	//	Updates(model.User{
	//		Name: "张三丰",
	//		Age:  27,
	//	})
	//if err != nil {
	//	log.Fatal("多字段更新失败:", err)
	//}
	//
	////4.删除用户
	//_, err = userDAD.WithContext(ctx).
	//	Where(q.User.ID.Eq(newUser.ID)).
	//	Delete()
	//if err != nil {
	//	log.Fatal("删除用户失败：", err)
	//}
	//fmt.Println("删除用户成功")

}
