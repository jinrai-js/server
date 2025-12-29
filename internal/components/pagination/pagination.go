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

	return `<ul class="ant-pagination css-4fj0kk">
  <li class="ant-pagination-prev ant-pagination-disabled" aria-disabled="true">
    <button
      class="ant-pagination-item-link"
      type="button"
      tabindex="-1"
      disabled=""
    >
      <span role="img" aria-label="left" class="anticon anticon-left"
        ><svg
          viewBox="64 64 896 896"
          focusable="false"
          data-icon="left"
          width="1em"
          height="1em"
          fill="currentColor"
          aria-hidden="true"
        >
          <path
            d="M724 218.3V141c0-6.7-7.7-10.4-12.9-6.3L260.3 486.8a31.86 31.86 0 000 50.3l450.8 352.1c5.3 4.1 12.9.4 12.9-6.3v-77.3c0-4.9-2.3-9.6-6.1-12.6l-360-281 360-281.1c3.8-3 6.1-7.7 6.1-12.6z"
          ></path></svg
      ></span>
    </button>
  </li>` + html + `
  <li tabindex="0" class="ant-pagination-next" aria-disabled="false">
    <button class="ant-pagination-item-link" type="button" tabindex="-1">
      <span role="img" aria-label="right" class="anticon anticon-right"
        ><svg
          viewBox="64 64 896 896"
          focusable="false"
          data-icon="right"
          width="1em"
          height="1em"
          fill="currentColor"
          aria-hidden="true"
        >
          <path
            d="M765.7 486.8L314.9 134.7A7.97 7.97 0 00302 141v77.3c0 4.9 2.3 9.6 6.1 12.6l360 281.1-360 281.1c-3.9 3-6.1 7.7-6.1 12.6V883c0 6.7 7.7 10.4 12.9 6.3l450.8-352.1a31.96 31.96 0 000-50.4z"
          ></path></svg
      ></span>
    </button>
  </li>
</ul>`
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

		html.WriteString(fmt.Sprintf(`<li class="%s" tabindex="0"><a href="?page=%d" rel="nofollow">%d</a></li>`, strings.Join(classes, " "), i, i))
	}

	return html.String()
}
