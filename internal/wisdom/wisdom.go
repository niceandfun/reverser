package wisdom

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://www.goodreads.com/quotes/tag/wisdom"

type Wisdom string

func (w *Wisdom) getRandomWisdom() {
	// Инициализация генератора случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Отправка HTTP-запроса
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: статус ответа", resp.StatusCode)
		return
	}

	// Парсинг HTML-контента
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при парсинге HTML:", err)
		return
	}

	// Извлечение цитат
	var quotes []string
	doc.Find(".quoteText").Each(func(i int, s *goquery.Selection) {
		quote := s.Text()
		quotes = append(quotes, quote)
	})

	// Проверка наличия цитат
	if len(quotes) == 0 {
		fmt.Println("Цитаты не найдены")
		return
	}

	// Выбор случайной цитаты
	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	// Вывод случайной цитаты
	fmt.Println("Случайная цитата:", randomQuote)
}

func (w *Wisdom) updateWisdom() {
	w.getRandomWisdom()
}

func (w *Wisdom) cleanWisdom() {
	*w = Wisdom("")
}

func New() Wisdom {
	var w Wisdom
	w.getRandomWisdom()
	return w
}
