package wisdom

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://www.goodreads.com/quotes/tag/wisdom"

type Wisdom string

func getRandomWisdom() (string, error) {
	// Инициализация генератора случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Отправка HTTP-запроса
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

	// Парсинг HTML-контента
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Извлечение цитат
	var quotes []string
	doc.Find(".quoteText").Each(func(i int, s *goquery.Selection) {
		quote := s.Text()
		quotes = append(quotes, quote)
	})

	// Проверка наличия цитат
	if len(quotes) == 0 {
		fmt.Println()
		return "", errors.New("цитаты не найдены")
	}

	// Выбор случайной цитаты
	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	return randomQuote, nil
}

func (w *Wisdom) updateWisdom() error {
	quote, err := getRandomWisdom()
	if err != nil {
		return err
	}

	*w = Wisdom(quote)
	return nil
}

func (w *Wisdom) cleanWisdom() {
	*w = Wisdom("")
}

func New() *Wisdom {
	var w Wisdom
	quote, err := getRandomWisdom()
	if err != nil {
		return &w
	}
	w = Wisdom(quote)
	return &w
}
