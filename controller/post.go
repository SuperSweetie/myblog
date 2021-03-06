package controller

import (
	"strings"

	"github.com/AmyangXYZ/SG_Sweetie/model"
	"github.com/AmyangXYZ/sweetygo"
)

// Home Page Handler.
func Home(ctx *sweetygo.Context) error {
	ctx.Set("title", "Home")
	posts, err := model.GetPosts("1")
	if err != nil {
		return err
	}
	ctx.Set("posts", posts)
	return ctx.Render(200, "home")
}

// PaginationHome returns 5 posts per page.
// for load-more button of Home Page.
func PaginationHome(ctx *sweetygo.Context) error {
	if page := ctx.Param("n"); page != "" {
		posts, err := model.GetPosts(page)
		if err != nil {
			return err
		}
		return ctx.JSON(200, 1, "success", posts)
	}
	return nil
}

// Show Post Page Handler.
func Show(ctx *sweetygo.Context) error {
	if title := ctx.Param("title"); title != "" {
		title := strings.Replace(title, "-", " ", -1)
		post, err := model.GetPostByTitle(title)
		if err != nil {
			return err
		}
		if post.ID == 0 {
			return ctx.Text(404, "404 not found")
		}
		ctx.Set("post", post)
		ctx.Set("title", title)
		ctx.Set("show", true)
		return ctx.Render(200, "posts/show")
	}
	return nil
}

// Cat shows posts sorted by category.
func Cat(ctx *sweetygo.Context) error {
	if cat := ctx.Param("cat"); cat != "" {
		posts, err := model.GetPostsByCat(cat, "1")
		if err != nil {
			return err
		}
		b := []byte(cat)
		b[0] -= 32 // uppercase
		cat = string(b)
		ctx.Set("cat", true)
		ctx.Set("posts", posts)
		ctx.Set("title", cat)
		return ctx.Render(200, "posts/cat")
	}
	return nil
}

// PaginationCat returns 5 posts per page.
// for load more of Cat Page.
func PaginationCat(ctx *sweetygo.Context) error {
	page := ctx.Param("n")
	cat := ctx.Param("cat")
	if page != "" && cat != "" {
		posts, err := model.GetPostsByCat(cat, page)
		if err != nil {
			return err
		}
		return ctx.JSON(200, 1, "success", posts)
	}
	return nil
}

// NewPage is Create Post Page Handler
func NewPage(ctx *sweetygo.Context) error {
	ctx.Set("title", "New")
	ctx.Set("editor", true)
	return ctx.Render(200, "posts/new")
}

// EditPage is Edit Post Page Handler
func EditPage(ctx *sweetygo.Context) error {
	ctx.Set("title", "Edit")
	ctx.Set("editor", true)
	title := ctx.Param("title")
	title = strings.Replace(title, "-", " ", -1)
	post, err := model.GetPostByTitle(title)
	if err != nil {
		return err
	}
	ctx.Set("post", post)
	return ctx.Render(200, "posts/edit")
}

// New Post API Handler.
//
// Usage:
//  "/api/posts/new" -X POST -d "title=xx&cat=xx&text=xx"
func New(ctx *sweetygo.Context) error {
	title := ctx.Param("title")
	cat := ctx.Param("cat")
	html := ctx.Param("html")
	md := ctx.Param("md")
	if title != "" && cat != "" && html != "" && md != "" {
		err := model.NewPost(title, cat, html, md)
		if err != nil {
			return ctx.JSON(500, 0, "create post error", nil)
		}
		return ctx.JSON(201, 1, "success", nil)
	}
	return ctx.JSON(406, 0, "I can't understand what u want", nil)
}

// Update Post API Handler.
//
// Usage:
// 	"/api/post" -X PUT -d "title=xx&cat=xx&text=xx"
func Update(ctx *sweetygo.Context) error {
	oldTitle := ctx.Param("title")     // from url
	newTitle := ctx.Param("new-title") // from form
	cat := ctx.Param("cat")
	html := ctx.Param("html")
	md := ctx.Param("md")
	if newTitle != "" && cat != "" && html != "" && md != "" {
		err := model.UpdatePost(newTitle, cat, html, md, oldTitle)
		if err != nil {
			return ctx.JSON(500, 0, "update post error", nil)
		}
		return ctx.JSON(201, 1, "success", nil)
	}
	return ctx.JSON(406, 0, "I can't understand what u want", nil)
}
