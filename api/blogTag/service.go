// package blog_tag

// import "fmt"

// type BlogTagService interface {
// 	GetTags(blogID string) ([]string, error)
// 	AddTag(blogID string, tag string) error
// 	DeleteTag(blogID string, tag string) error
// }

// type blogTagService struct {
// 	repo BlogTagRepository
// }

// func NewBlogTagService(repo BlogTagRepository) BlogTagService {
// 	return &blogTagService{repo: repo}
// }

// func (s *blogTagService) GetTags(blogID string) ([]string, error) {
// 	blog, err := s.repo.GetBlogByID(blogID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return blog.Tags, nil
// }

// func (s *blogTagService) AddTag(blogID string, tag string) error {
// 	blog, err := s.repo.GetBlogByID(blogID)
// 	if err != nil {
// 		return err
// 	}

// 	for _, existingTag := range blog.Tags {
// 		if existingTag == tag {
// 			return fmt.Errorf("tag already exists")
// 		}
// 	}

// 	blog.Tags = append(blog.Tags, tag)
// 	return s.repo.SaveBlog(blog)
// }

// func (s *blogTagService) DeleteTag(blogID string, tag string) error {
// 	blog, err := s.repo.GetBlogByID(blogID)
// 	if err != nil {
// 		return err
// 	}

// 	var updatedTags []string
// 	tagFound := false
// 	for _, existingTag := range blog.Tags {
// 		if existingTag != tag {
// 			updatedTags = append(updatedTags, existingTag)
// 		} else {
// 			tagFound = true
// 		}
// 	}

// 	if !tagFound {
// 		return fmt.Errorf("tag not found")
// 	}

// 	blog.Tags = updatedTags
// 	return s.repo.SaveBlog(blog)
// }
