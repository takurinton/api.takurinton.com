package service

import (
	"fmt"
	"portfolio/model"
	"portfolio/utils"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type Blog struct{}

// RSS 用
func (b Blog) GetAllPosts() (p []*model.Rss, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("blogapp_post").Select("id, title, contents, pub_date").Where("open = true").Find(&p).Error; err != nil {
		return
	}

	for _, item := range p {
		item.Contents = utils.ParseContents(item.Contents)
	}

	return
}

// 管理画面用
// open = true をなくして降順ソートしてる
func (b Blog) GetAllPostsReverse() (p []*model.BlogappPost, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("blogapp_post").Select("blogapp_post.id, title, contents, open, blogapp_category.name as name, blogapp_post.pub_date").Joins("inner join blogapp_category on blogapp_post.category_id = blogapp_category.id").Order("blogapp_post.pub_date desc").Find(&p).Error; err != nil {
		return
	}

	for _, item := range p {
		item.Contents = utils.ParseContents(item.Contents)
	}

	return
}

// ページネーションも実装するからここの処理は重くなりそう
func (b *Blog) GetPosts(page int, category string) (pn *pagination.Paginator, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	posts := []*model.BlogappPost{}
	var query *gorm.DB
	if category == "" {
		// query = db.Raw("SELECT blogapp_post.id, title, name, substring(contents, 1, 200), pub_date FROM blogapp_post INNER JOIN blogapp_category ON blogapp_post.category_id = blogapp_category.id WHERE open = true").Scan(&posts)
		query = db.Table("blogapp_post").Select("blogapp_post.id, title, name, left(contents, 200) as contents, pub_date").Where("open = true").Joins("inner join blogapp_category on blogapp_post.category_id = blogapp_category.id").Order("pub_date desc").Find(&posts)
	} else {
		query = db.Table("blogapp_post").Select("blogapp_post.id, title, name, left(contents, 200) as contents, pub_date").Where("open = true and name = ?", category).Joins("inner join blogapp_category on blogapp_post.category_id = blogapp_category.id").Order("pub_date desc").Find(&posts)
	}

	pn = pagination.Paging(&pagination.Param{
		DB:      query,
		Page:    page,
		Limit:   5,
		ShowSQL: false,
	}, &posts)

	for _, post := range posts {
		post.Contents = fmt.Sprintf("%s...", utils.ParseContents(post.Contents))
	}

	return
}

func (b Blog) GetPost(id string) (p model.BlogappPost, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()
	if err = db.Table("blogapp_post").Select("blogapp_post.id, title, contents_image_url, name, contents, pub_date").Where("open = true and blogapp_post.id = ?", id).Joins("inner join blogapp_category on blogapp_post.category_id = blogapp_category.id").Find(&p).Error; err != nil {
		return
	}

	var comment []model.BlogappComment
	if comment, err = GetBlogComment(id); err != nil {
		return
	}
	p.Comment = comment

	return
}

func (b Blog) GetPostAdmin(id string) (p model.BlogappPost, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()
	if err = db.Table("blogapp_post").Select("blogapp_post.id, title, open, name, contents, pub_date").Where("blogapp_post.id = ?", id).Joins("inner join blogapp_category on blogapp_post.category_id = blogapp_category.id").Find(&p).Error; err != nil {
		return
	}

	var comment []model.BlogappComment
	if comment, err = GetBlogComment(id); err != nil {
		return
	}
	p.Comment = comment

	return
}

func GetBlogComment(id string) (p []model.BlogappComment, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Select("post_id_id, name, contents, pub_date").Table("blogapp_comment").Find(&p, "post_id_id=?", id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	return
}

func (b Blog) PostComment(id int64, name, contents string) error {
	t := time.Now()
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	c := model.BlogappComment{PostIdId: id, Name: name, Contents: contents, PubDate: t}

	if err := db.Table("blogapp_comment").Create(&c).Error; err != nil {
		return err
	}
	return nil
}

// こんな感じ
// {
// 'title': 'hgufa',
// 'contents': '# hello world',
// 'contentsImageUrl': '/images/markdownx/f22f7003-01a0-4a06-a885-3b8d87676e43.jpeg',
// 'name': 'インターン',
// 'open': True
// }
func (b Blog) CreatePost(post model.Post) (err error) {
	t := time.Now()
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	var category model.BlogappCategory
	if err = db.Select("id").Table("blogapp_category").Find(&category, "name = ?", post.Name).Error; err != nil {
		return
	}

	// var create model.BlogappPost
	db.Exec("INSERT INTO blogapp_post(title, contents, contents_image_field, contents_image_url, pub_date, view, category_id, open) VALUES (?, ?, ?, ?, ?, ?, ?, ?);", post.Title, post.Contents, "hoge", "hoge", t, 0, category.Id, post.Open)
	return nil
}

func (b Blog) UpdatePost(id string, post model.Post) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	var category model.BlogappCategory
	if err = db.Select("id").Table("blogapp_category").Find(&category, "name = ?", post.Name).Error; err != nil {
		return
	}

	db.Exec("UPDATE blogapp_post SET title = ?, contents = ?, contents_image_field = ?, contents_image_url = ?, pub_date = ?, view = ?, category_id = ?, open = ? WHERE id = ?;", post.Title, post.Contents, "hoge", "hoge", post.PubDate, 0, category.Id, post.Open, id)
	return nil
}

// 君はidだけでいいよ
func (b Blog) DeletePost(id string) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("DELETE FROM blogapp_post WHERE id = ?", id)
	return nil
}

func (b Blog) CreateCategory(name string) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("INSERT INTO blogapp_category(name) VALUES (?);", name)
	return nil
}

func (b Blog) UpdateCategory(id, name string) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	// var category model.BlogappCategory
	db.Exec("UPDATE blogapp_category SET name = ? WHERE id = ?;", name, id)
	return nil
}

func (b Blog) GetCategory() (res []model.BlogappCategory, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("blogapp_category").Find(&res).Error; err != nil {
		return
	}
	return
}

func (b Blog) GetDetailCategory(id string) (res model.BlogappCategory, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("blogapp_category").Where("id = ?", id).Find(&res).Error; err != nil {
		return
	}
	return
}

func (b Blog) DeleteCategory(id string) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("DELETE FROM blogapp_category WHERE id = ?;", id)
	return nil
}
