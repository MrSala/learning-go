package articles

type Getter interface {
	GetAll() []Article
}

type Adder interface {
	Add(a Article)
}

type Article struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

// need to learn more about syntax
type Repo struct {
	Articles []Article
}

// need to learn more about syntax
func New() *Repo {
	return &Repo{
		Articles: []Article{},
	}
}

// need to learn more about syntax
func (r *Repo) Add(a Article) {
	r.Articles = append(r.Articles, a)
}

func (r *Repo) GetAll() []Article {
	return r.Articles
}
