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

// don't know the idea behind it
type Repo struct {
	Articles []Article
}

// don't know the idea behind it
func New() *Repo {
	return &Repo{
		Articles: []Article{},
	}
}

func (r *Repo) Add(a Article) {
	r.Articles = append(r.Articles, a)
}

func (r *Repo) GetAll() []Article {
	return r.Articles
}
