# Hexlet Path Size

[![hexlet-check](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/hexlet-check.yml)
[![CI](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/CI.yml/badge.svg)](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/CI.yml)

**Hexlet Path Size** — консольная утилита для подсчёта суммарного размера файлов и директорий.


## Возможности

* Подсчёт размера файла.
* Подсчёт размера директории.
* Рекурсивный обход вложенных директорий.
* Учёт или игнорирование скрытых файлов и директорий.
* Вывод размера в удобочитаемом формате (`KB`, `MB`, `GB` и т.д.).
* Поддержка символических ссылок.

---

## Установка

Клонируйте репозиторий:

```bash
git clone https://github.com/bond6699/go-from-scratch-project-242.git
cd <repository>
```

Соберите проект:

```bash
make build
```

или

```bash
go build -o hexlet-path-size ./cmd/hexlet-path-size
```

---

## Использование

Общий синтаксис:

```text
hexlet-path-size [ФЛАГИ] <ПУТЬ>
```

### Флаги

| Флаг                | Описание                                      |
| ------------------- | --------------------------------------------- |
| `-r`, `--recursive` | Рекурсивно обрабатывать вложенные директории. |
| `-a`, `--all`       | Учитывать скрытые файлы и директории.         |
| `-H`, `--human`     | Выводить размер в удобочитаемом формате.      |
| `--help`            | Показать справку.                             |

> **Примечание**
>
> Если указан путь к файлу, флаг `--all` не влияет на результат.

---

## Примеры

Подсчитать размер файла:

```bash
hexlet-path-size file.txt
```

Вывод:

```text
2048B
```

Подсчитать размер директории:

```bash
hexlet-path-size ./project
```

Рекурсивный подсчёт:

```bash
hexlet-path-size -r ./project
```

Учитывать скрытые файлы:

```bash
hexlet-path-size -a ./project
```

Вывести результат в удобочитаемом формате:

```bash
hexlet-path-size -r -a -H ./project
```

Пример вывода:

```text
12.4MB
```

---

## Структура проекта

```text
cmd/
    hexlet-path-size/      Точка входа в приложение

internal/
    cli/                   Настройка CLI
    formatter/             Форматирование результата
    pathsize/              Логика подсчёта размера
```

---

## Разработка

Запуск тестов:

```bash
make test
```

Проверка линтером:

```bash
make lint
```

---

## Демонстрация

Запись работы программы:

[![asciicast](https://asciinema.org/a/vSBm8O7vYUjdh9wh.svg)](https://asciinema.org/a/vSBm8O7vYUjdh9wh)

---

## Используемые технологии

* Go
* urfave/cli/v3
* stretchr/testify

---