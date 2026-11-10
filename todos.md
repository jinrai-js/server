# TODO

## 🐛 fetch_group.Wait() — риск бесконечного ожидания

**Где:** `internal/lib/fetch_group/fetch_group.go`, `internal/lib/app_state/app_state_type.go` (`GetWithoutSource`), `internal/lib/handler/hander.go` (`Render`).

**Проблема:** `globalFetchGroup` — общий глобальный `sync.WaitGroup` на весь сервер, и `Wait()` блокируется, пока счётчик не станет 0. Из-за этого:

1. **`Wait()` ждёт фетчи ВСЕХ параллельных запросов**, а не только текущего. Под нагрузкой новые `Run`/`Add` постоянно пополняют счётчик → `Wait` может не вернуться практически бесконечно.
2. **Гонка `Add` ↔ `Wait` при нулевом счётчике** — запрещённое использование `sync.WaitGroup` (доки: "calls with a positive delta that start when the counter is zero must happen before a Wait"). Последствия: `Wait` возвращается раньше времени (потеря результата) или panic `"WaitGroup is reused before previous Wait has returned"`.
3. **Гарантия завершения держится только на таймауте `http.Client{Timeout: 3s}`** в `SendRequest`. Неразрывность `Add`→`Done` означает: если фетч не завершится (уберут таймаут / добавят источник без таймаута / стриминг) — `Wait` повиснет навсегда.

**Бонус:** паника внутри goroutine фетча (`log.Panic` в `SendRequest`, `app_error.Create` → `panic(*server_panic_container)`) никем не ловится → крах всего процесса.

**Также:** `InitRequired` вызывается дважды за запрос (`internal/lib/jinrai/handler.go:32` напрямую + внутри `GetWithoutSource` через `JoinStates`) — второй вызов лишняя работа / лишний `pass.Exit()` для дубликатов ключей.

**Идеи для исправления:**
- Ждать только свои фетчи: per-request группа ожидания (в контексте), а не глобальный синглтон.
- Не использовать глобальный `WaitGroup` — перенести ожидание в контекст запроса или использовать `errgroup`/каналы.
- Добавить `recover` в goroutine фетча.