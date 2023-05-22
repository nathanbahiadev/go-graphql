package graph

import "github.com/nathanbahiadev/go-graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
