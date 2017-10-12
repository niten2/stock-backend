goconvey
append(TodoList, todo1, todo2, todo3)
rand.Seed(time.Now().UnixNano())



  // s3Bucket := os.Getenv("S3_BUCKET")
  // secretKey := os.Getenv("SECRET_KEY")

	// fmt.Println(s3Bucket)
	// fmt.Println(secretKey)


curl -g 'http://localhost:8000/graphql?query={todo(id:\"b\"){id,text,done}}'

curl -g 'http://localhost:8000/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'


gin -p 4000 --appPort 4001 run main.go



fresh

http://localhost:8000/graphql?query={todoList{id,text,done}}

http://localhost:8000/graphql?query={todoList{id,name}}

http://localhost:8000/graphql?query={products{id,name}}

curl -g 'http://localhost:8000/graphql?query={products{id,name}}'

curl -g 'http://localhost:8000/graphql?query={todoList{id,title}}'


curl -g 'http://localhost:8000/graphql?query={todo{id,title, name}}'


curl -g 'http://localhost:8000/graphql?query={todoList{id,text,done}}'

curl -g 'http://localhost:8000/graphql?query=mutation+_{createTodo(text:"11111My+new+todo"){id,text,done}}'


// curl -g 'http://localhost:8000/graphql?query={todo(id:\"b\"){id,text,done}}'
// fmt.Println("Get single todo: curl -g 'http://localhost:8080/graphql?query={todo(id:\"b\"){id,text,done}}'")
// fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
// fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")
// fmt.Println("Load todo list: curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'")
// fmt.Println("Access the web app via browser at 'http://localhost:8080'")
//
//

