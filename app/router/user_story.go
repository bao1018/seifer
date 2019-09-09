package router

import (
	"net/http"
	"sort"
	"sync"

	"bitbucket.org/kyicy/seifer/app/api"
	"bitbucket.org/kyicy/seifer/app/model"
	"github.com/labstack/echo/v4"
	"gonum.org/v1/gonum/mat"
)

// DIM vector dimension
const DIM = 512

type createUserStoryBody struct {
	Title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
}

// CreateUserStory handler
func CreateUserStory(c echo.Context) error {
	body := new(createUserStoryBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	userStory := new(model.UserStory)
	userStory.Title = body.Title
	userStory.Body = body.Body

	titleVector, bodyVector := titleAndBodyToVectors(body.Title, body.Body)

	userStory.Title = body.Title
	userStory.Body = body.Body
	userStory.TitleVector = titleVector
	userStory.BodyVector = bodyVector

	model.DB.Save(userStory)

	return c.NoContent(http.StatusOK)
}

type similarUserStoriesBody struct {
	Title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
}

// SimilarUserStories handler
func SimilarUserStories(c echo.Context) error {
	body := new(similarUserStoriesBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	titleVector, bodyVector := titleAndBodyToVectors(body.Title, body.Body)

	titleDense := mat.NewDense(1, DIM, titleVector)
	bodyDense := mat.NewDense(1, DIM, bodyVector)

	var userStories []model.UserStory

	model.DB.Find(&userStories)

	var ids []uint
	var titleVecs, bodyVecs []float64
	for _, userStory := range userStories {
		ids = append(ids, userStory.ID)
		titleVecs = append(titleVecs, userStory.TitleVector...)
		bodyVecs = append(bodyVecs, userStory.BodyVector...)
	}

	titleMatrix := mat.NewDense(len(userStories), DIM, titleVecs)
	bodyMatrix := mat.NewDense(len(userStories), DIM, bodyVecs)
	titleMatrixT := titleMatrix.T()
	bodyMatrixT := bodyMatrix.T()

	var titleRank mat.Dense
	titleRank.Mul(titleDense, titleMatrixT)
	titleRank.Scale(1.3, &titleRank)

	var bodyRank mat.Dense
	bodyRank.Mul(bodyDense, bodyMatrixT)

	var ScoreMat mat.Dense
	ScoreMat.Add(&titleRank, &bodyRank)

	Score := mat.Row(nil, 0, &ScoreMat)

	slice := NewSlice(Score...)

	sort.Sort(sort.Reverse(slice))

	var userStoryResult []*model.UserStory

	for _, index := range slice.idx {
		userStoryResult = append(userStoryResult, &userStories[index])
	}

	return c.JSON(http.StatusOK, userStoryResult)
}

func titleAndBodyToVectors(title, body string) (titleVector []float64, bodyVector []float64) {
	var wg sync.WaitGroup
	wg.Add(2)
	// process title
	go func() {
		defer wg.Done()
		titleVector = api.SentenceToVector(title)
	}()

	// process body
	go func() {
		defer wg.Done()
		bodyVector = api.SentenceToVector(body)
	}()
	wg.Wait()

	return
}

// Slice struct
type Slice struct {
	sort.Float64Slice
	idx []int
}

// Swap for order
func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// NewSlice helper
func NewSlice(n ...float64) *Slice {
	s := &Slice{Float64Slice: sort.Float64Slice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}
