package controller

import (
	"chk/models"
	"chk/services"
	"chk/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	_ "github.com/joho/godotenv/autoload"
	"go/types"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type HomeController struct {
}

var count int

func (receiver HomeController) Upload(c *gin.Context) {
	mediaType, params, err := mime.ParseMediaType(c.GetHeader("Content-Type"))
	if err != nil {
		utils.HttpErrorResponse(c, http.StatusBadRequest, err, "parse media type error")
		return
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		utils.HttpErrorResponse(c, http.StatusBadRequest, err, "content type not support")
		return
	}

	var result models.Result
	mr := multipart.NewReader(c.Request.Body, params["boundary"])

	_, err = processCreate(mr)
	if err != nil {
		utils.HttpErrorResponse(c, http.StatusServiceUnavailable, result.Error, "process create error")
		return
	}

	if result.Error != nil {
		utils.HttpErrorResponse(c, http.StatusServiceUnavailable, result.Error, "service error")
		return
	}

	//fmt.Println("count")
	//fmt.Println(count)

	utils.HttpSuccessResponse(c, http.StatusOK, types.Interface{}, "ok")
}

func processCreate(mr *multipart.Reader) (int, error) {
	var err error
	numWps := 100
	result := make([]models.Result, 0)
	jobs := make(chan []byte, numWps)
	res := make(chan models.Result)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []byte, results chan<- models.Result) {
		for {
			select {
			case job, ok := <-jobs: // check for readable state of the channel.
				if !ok {
					return
				}
				results <- parseAndCreate(job)
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			p, err := mr.NextPart()

			if err == io.EOF {
				break
			}
			if err != nil {
				err = errors.New("reading next part error")
			}
			part, err := ioutil.ReadAll(p)
			if err != nil {
				err = errors.New("reading part error")
			}

			fmt.Println("part")
			fmt.Println(len(part))
			jobs <- part
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		result = append(result, r)
	}

	return 0, err
}

func parseAndCreate(part []byte) models.Result {
	var priceHistories []models.PriceHistory
	err := gocsv.UnmarshalBytes(part, &priceHistories)
	if err != nil {
		return models.Result{}
	} else {
		count += len(priceHistories)
		fmt.Println(len(priceHistories))
		priceHistoryService := new(services.PriceHistoryService)
		return priceHistoryService.Create(priceHistories)
	}
}

func (receiver HomeController) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	getParam := models.GetParam{
		Page:  page,
		Limit: limit,
	}

	priceHistoryService := new(services.PriceHistoryService)
	result := priceHistoryService.Get(getParam)
	if result.Error != nil {
		utils.HttpErrorResponse(c, http.StatusServiceUnavailable, result.Error, "service error")
	}

	utils.HttpSuccessResponse(c, http.StatusOK, result.Body, "ok")
}
