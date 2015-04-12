package service

import (
  "github.com/benschw/go-todo/api"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "log"
  "strconv"
  "time"
)

type TodoResource struct {
  db gorm.DB
}

func (tr *TodoResource) CreateToDo(c *gin.Context) {
  var todo api.Todo

  if !c.Bind(&todo) {
    c.JSON(400, api.NewError("Problem decoding body"))
    return
  }
  todo.Status = api.TodoStatus
  todo.Created = int32(time.Now().Unix())

  tr.db.Save(&todo)

  c.JSON(201, todo)
}

func (tr *TodoResource) GetAllTodos(c *gin.Context) {
  var todos []api.Todo

  tr.db.Order("created desc").Find(&todos)

  c.JSON(200, todos)
}

func (tr *TodoResource) GetTodo(c *gin.Context) {
  id, err := tr.getId(c)
  if err != nil {
    c.JSON(400, api.NewError("Problem decoding id sent"))
    return
  }

  var todo api.Todo

  if tr.db.First(&todo, id).RecordNotFound() {
    c.JSON(404, gin.H{"error": "not found"})
  } else {
    c.JSON(200, todo)
  }
}

func (tr *TodoResource) UpdateTodo(c *gin.Context) {
  id, err := tr.getId(c)
  if err != nil {
    c.JSON(400, api.NewError("Problem decoding id sent"))
    return
  }

  var todo api.Todo

  var existing api.Todo

  if tr.db.First(&existing, id).RecordNotFound() {
    c.JSON(404, api.NewError("not found"))
  } else {
    tr.db.Sabe(&todo)
    c.JSON(200, todo)
  }
}
