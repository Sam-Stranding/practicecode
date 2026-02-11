1.设计API
    GET    /students                  #获取所有学生列表
    POST   /student       body=json   #创建新学生
    Get    /students?id=1             #获取特定学生
    PUT    /students?id=1 body=json   #更新特定学生
    DELETE /students?id=1 body=json   #删除特定学生
2.定义实现方法
func getStudents(w http.ResponseWriter, r *http.Request){
//处理GET /students接口
}
func createStudents(w http.ResponseWriter, r *http.Request){
//处理POST /students接口
}

3.注册路由
mux := http.NewServeMux()
mux.HandleFunc("/students", studentsHandler) //处理 /students 接口
4.启动服务
//监听8080端口，使用自定义路由
http.ListenAndServe(":8080",mux)