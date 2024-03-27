package domain

import (
	"regexp"
	"strconv"
)

func isValidBlogPost(blogPost string) bool {
	// Regular expression to validate blog post
	min := 200
	max := 2000
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s.,?!;:'"()\[\]{}-]{` + strconv.Itoa(min) + `,` + strconv.Itoa(max) + `}$`)

	// Check if the blog post matches the regular expression
	return regex.MatchString(blogPost)
}

func isValidComment(comment string) bool {
	// Regular expression to validate comment
	min := 1
	max := 200
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s.,?!;:'"()\[\]{}-]{` + strconv.Itoa(min) + `,` + strconv.Itoa(max) + `}$`)

	// Check if the comment matches the regular expression
	return regex.MatchString(comment)
}
