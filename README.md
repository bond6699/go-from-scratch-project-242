# hexlet-path-size

[![hexlet-check](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/hexlet-check.yml)
[![CI](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/CI.yml/badge.svg)](https://github.com/bond6699/go-from-scratch-project-242/actions/workflows/CI.yml)

Консольная утилита для подсчёта размера файлов и содержимого директорий (аналог `du` с расширенными возможностями).  
Написана на Go в рамках учебного проекта.

## Особенности

- Рекурсивный обход вложенных папок
- Учёт или игнорирование скрытых файлов и папок
- Вывод размера в байтах или в удобочитаемом формате (B, KB, MB, GB)
- Обработка больших директорий без переполнения стека (итеративный обход)
- Код ошибки при некорректном пути или проблемах чтения
- Ссылки (sumlink) игнорируются

## Установка

### 1. Клонирование репозитория

```bash
git clone https://github.com/bond6699/go-from-scratch-project-242.git
cd go-from-scratch-project-242
```

### 2. Сборка бинарного файла

```bash
go build -o hexlet-path-size ./cmd/hexlet-path-size
```

После сборки исполняемый файл `hexlet-path-size` появится в корне проекта.  
Можно переместить его в `$PATH` для глобального доступа:

```bash
sudo mv hexlet-path-size /usr/local/bin/
```

### 3. Альтернативная установка через `go install`

```bash
go install ./cmd/hexlet-path-size
```

Бинарник будет установлен в `$GOPATH/bin`.

## Использование

```bash
./hexlet-path-size [глобальные флаги] <путь>
```

Где `<путь>` — путь к файлу или директории.

### Флаги

| Флаг | Короткая форма | Описание | Значение по умолчанию |
|------|----------------|----------|-----------------------|
| `--recursive` | `-r` | Рекурсивно обходить все поддиректории | `false` |
| `--all` | `-a` | Включать в подсчёт скрытые файлы/папки (начинающиеся с точки) | `false` |
| `--human` | `-H` | Выводить размер в человекочитаемом формате (KB, MB, GB) | `false` |
| `--help` | `-h` | Показать справку | — |

### Примеры

#### 1. Размер одного файла (в байтах) (при передаче пути на один файл флаг --all включен по умолчанию)

```bash
./hexlet-path-size photo.jpg
# Вывод: 153600B
```

#### 2. Размер текущей директории (без рекурсии, без скрытых)

```bash
./hexlet-path-size .
# Вывод: 4096B (размер самой директории)
```

#### 3. Рекурсивный подсчёт всего проекта

```bash
./hexlet-path-size -r ~/myproject
# Вывод: 1048576B (общий размер всех файлов)
```

#### 4. Рекурсивно, со скрытыми файлами, в человекочитаемом формате

```bash
./hexlet-path-size -r -a -H ./project
# Вывод: 1.23GB
```

#### 5. Комбинированный вывод с указанием пути

```bash
./hexlet-path-size -r -H /var/log
# Вывод: 456.78MB
```

#### 6. ASCIINEMA Test, Build and Run

[![asciicast](https://asciinema.org/a/vSBm8O7vYUjdh9wh.svg)](https://asciinema.org/a/vSBm8O7vYUjdh9wh)