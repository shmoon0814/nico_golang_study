package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func Scrape(term string) {
	baseURL = "https://kr.indeed.com/jobs?q=+" + term + "&limit=50"
	var jobs []extractedJob
	ch := make(chan []extractedJob)
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		//extractedJobs := getPage(i, ch)
		go getPage(i, ch)
		//jobs = append(jobs, extractedJobs...)
	}

	for i := 0; i < totalPages; i++ {
		job := <-ch
		jobs = append(jobs, job...)
	}

	writeJobs(jobs)
	fmt.Println("Scrapping Done!")
}
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		w.Write(job.writeBody())
	}
}

func (job *extractedJob) writeBody() []string {
	return []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
}

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

func getPage(page int, ch chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	ch <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := card.Find(".title>a").Text()
	location := card.Find(".sjcl").Text()
	salary := card.Find(".salaryText").Text()
	summary := card.Find(".summary").Text()
	c <- extractedJob{id: id, title: cleanString(title), location: cleanString(location), salary: cleanString(salary), summary: cleanString(summary)}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}
