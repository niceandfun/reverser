package wisdom

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://www.goodreads.com/quotes/tag/wisdom"

type Wisdom string

func getRandomWisdom() (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println()
		return "", fmt.Errorf("ошибка: статус ответа %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var quotes []string
	doc.Find(".quoteText").Each(func(i int, s *goquery.Selection) {
		quote := s.Text()
		quotes = append(quotes, quote)
	})

	if len(quotes) == 0 {
		fmt.Println()
		return "", errors.New("цитаты не найдены")
	}

	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	return randomQuote, nil
}

func (w *Wisdom) UpdateWisdom() error {
	quote, err := getRandomWisdom()
	if err != nil {
		return err
	}

	*w = Wisdom(quote)
	return nil
}

func (w *Wisdom) CleanWisdom() {
	*w = Wisdom("")
}

func (w Wisdom) String() string {
	return string(w)
}

func New() (*Wisdom, error) {
	var w Wisdom

	quote, err := getRandomWisdom()
	if err != nil {
		return nil, err
	}

	quote = formatQuote(quote)

	w = Wisdom(quote)
	return &w, nil
}

func formatQuote(q string) string {
	q = strings.TrimSpace(q)
	q = strings.ReplaceAll(q, "\n", " ")
	q = strings.ReplaceAll(q, "\r", "")
	return q
}
