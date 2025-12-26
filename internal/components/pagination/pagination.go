package pagination

import (
	"fmt"
	"strings"
)

type Props struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Pages int `json:"pages"`
}

func Component(props Props) string {
	html := getHTML(props)

	return `<ul class="ant-pagination css-4fj0kk">` + html + `</ul>`
}

func getHTML(props Props) string {
	var html strings.Builder

	// Рассчитываем количество элементов слева и справа от центра
	sideItems := props.Size / 2
	start := props.Page - sideItems
	end := props.Page + sideItems

	// Если size четный, уменьшаем end на 1, чтобы общее кол-во было равно size
	if props.Size%2 == 0 {
		end--
	}

	// Корректировка границ, чтобы не выйти за пределы [1, pages]
	if start < 1 {
		end = end + (1 - start)
		start = 1
	}
	if end > props.Pages {
		start = start - (end - props.Pages)
		end = props.Pages
	}
	// Вторая проверка на случай, если size > pages
	if start < 1 {
		start = 1
	}

	for i := start; i <= end; i++ {
		classes := []string{"ant-pagination-item", fmt.Sprintf("ant-pagination-item-%d", i)}
		if i == props.Page {
			classes = append(classes, "ant-pagination-item-active")
		}
		// Показываем индикатор "прыжка" назад, если скрыты страницы в начале
		if i == start && i > 1 {
			classes = append(classes, "ant-pagination-item-after-jump-prev")
		}
		// Показываем индикатор "прыжка" вперед, если скрыты страницы в конце
		if i == end && i < props.Pages {
			classes = append(classes, "ant-pagination-item-before-jump-next")
		}

		html.WriteString(fmt.Sprintf(`<li class="%s" tabindex="0"><a rel="nofollow">%d</a></li>`, strings.Join(classes, " "), i))
	}

	return html.String()
}
