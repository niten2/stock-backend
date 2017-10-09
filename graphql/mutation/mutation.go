package mutation

import (
	// "math/rand"
	"github.com/graphql-go/graphql"
)

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randStringRunes(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }


var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

    // // curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
		// "createTodo": &graphql.Field{
			// Type:        todoType, // the return type for this field
			// Description: "Create new todo",
			// Args: graphql.FieldConfigArgument{
				// "text": &graphql.ArgumentConfig{
					// Type: graphql.NewNonNull(graphql.String),
				// },
			// },
			// Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// // marshall and cast the argument value
				// text, _ := params.Args["text"].(string)

				// // figure out new id
				// newID := randStringRunes(8)

				// // perform mutation operation here
				// // for e.g. create a Todo and save to DB.
				// newTodo := Todo{
					// ID:   newID,
					// Text: text,
					// Done: false,
				// }

				// TodoList = append(TodoList, newTodo)

				// // return the new Todo object that we supposedly save to DB
				// // Note here that
				// // - we are returning a `Todo` struct instance here
				// // - we previously specified the return Type to be `todoType`
				// // - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
				// return newTodo, nil
			// },
		// },
		// /*
			// curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
		// */
		// "updateTodo": &graphql.Field{
			// Type:        todoType, // the return type for this field
			// Description: "Update existing todo, mark it done or not done",
			// Args: graphql.FieldConfigArgument{
				// "done": &graphql.ArgumentConfig{
					// Type: graphql.Boolean,
				// },
				// "id": &graphql.ArgumentConfig{
					// Type: graphql.NewNonNull(graphql.String),
				// },
			// },
			// Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// // marshall and cast the argument value
				// done, _ := params.Args["done"].(bool)
				// id, _ := params.Args["id"].(string)
				// affectedTodo := Todo{}

				// // Search list for todo with id and change the done variable
				// for i := 0; i < len(TodoList); i++ {
					// if TodoList[i].ID == id {
						// TodoList[i].Done = done
						// // Assign updated todo so we can return it
						// affectedTodo = TodoList[i]
						// break
					// }
				// }
				// // Return affected todo
				// return affectedTodo, nil
			// },
		// },
	},
})
