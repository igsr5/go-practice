package main
//
//import (
//  "fmt"
//  "errors"
//  "database/sql"
//  _"github.com/lib/pq"
//)
//
//type Post struct {
//  Id int
//  Content string
//  Author string
//}
//
//type Comment struct {
//  Id int
//  Content string
//  Author string
//  Post *Post
//}
//
//var Db *sql.DB
//
//// DB初期化
//func init() {
//  var err error
//  Db, err = sql.Open("postgres", "user=postgres dbname=go-web-app sslmode=disable")
//  if err != nil {
//    panic(err)
//  }
//}
//
//// Post
//func Posts(limit int) (posts []Post, err error) {
//  rows, err := Db.Query("select id, content, author from posts limit $1", limit)
//  if err != nil {
//    return
//  }
//  for rows.Next() {
//    post := Post{}
//    err = rows.Scan(&post.Id, &post.Content, &post.Author)
//    if err != nil {
//      return
//    }
//    posts = append(posts, post)
//  }
//
//  rows.Close()
//  return
//}
//
//func GetPost(id int) (post Post, err error) {
//  post = Post{}
//  err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
//  return
//}
//
//func (post *Post) Create() (err error) {
//  statement := "insert into posts (content, author) values ($1, $2) returning id"
//  stmt, err := Db.Prepare(statement)
//  if err != nil {
//    return
//  }
//  defer stmt.Close()
//  err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
//  return
//}
//
//func (post *Post) Update() (err error) {
//  _, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
//  return
//}
//
//func (post *Post) Delete() (err error) {
//  _, err = Db.Exec("delete from posts where id = $1", post.Id)
//  return
//}
//
//func (post *Post) Comment() (comments []Comment, err error) {
//  rows, err := Db.Query("select id, content, author from comments where post_id = $1", post.Id)
//  if err != nil {
//    fmt.Println(err)
//    return
//  }
//  for rows.Next() {
//    comment := Comment{Post: post}
//    err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
//    if err != nil {
//      fmt.Println(err)
//      return
//    }
//    comments = append(comments, comment)
//  }
//
//  rows.Close()
//  return
//}
//
//// Comment
//
//func (comment *Comment) Create() (err error) {
//  if comment.Post == nil {
//    err = errors.New("投稿が見つかりません。")
//    return
//  }
//  err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
//  return
//}
//
//func main() {
//  post := Post{Content: "Hello World!", Author: "Sau Sheong"}
//  post.Create()
//
//  comment := Comment{Content: "Nice!!", Author: "bonjour Monde!", Post: &post}
//  comment.Create()
//  comment.Create()
//  comment.Create()
//
//  comments,_ := post.Comment()
//  fmt.Println(comments[0].Post.Id)
//
//  //readPost.Content = "bonjour Monde!"
//  //readPost.Author = "Pierre"
//  //readPost.Update()
//
//  //fmt.Println(GetPost(readPost.Id))
//
//  //readPost.Delete()
//
//  // posts,_ := Posts(5)
//  // fmt.Println(posts)
//}
