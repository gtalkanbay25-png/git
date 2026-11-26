# Git Basics

Учебный репозиторий по основам git в Tomorrow School. Здесь я учусь создавать репозиторий, делать коммиты, работать с ветками, отправлять код на удалённый сервер и решать конфликты слияния. Также здесь есть Repo Health Analyzer на Go: программа, которая оценивает историю моих коммитов.

## Structure

    git/
    ├── README.md                      описание проекта (этот файл)
    ├── reflection.md                  что я узнала про git + бонус Prompted Commits
    ├── go.mod                         Go-модуль git-basics
    ├── analyzer/
    │   └── main.go                    анализатор: читает git log и печатает отчёт
    ├── internal/
    │   └── health/
    │       ├── health.go              ScoreMessage (оценка 0–3) и Label (tidy / acceptable / messy)
    │       └── health_test.go         unit-тесты для ScoreMessage и Label
    ├── feature/
    │   └── feature.txt                файл, созданный на ветке feature и слитый в main
    └── docs/
        └── workflow.md                как я работала с ветками, merge и конфликтом

## Requirements

- git
- Go 1.26 или новее

## How to run

1. Склонировать репозиторий и зайти в папку:

       git clone https://github.com/gtalkanbay25-png/git.git
       cd git

2. Запустить анализатор (из корня репозитория):

       go run ./analyzer

3. Запустить тесты:

       go test ./...

## Example output

    === Repo Health ===
    Commits: 16
    Span: 1 days
    Cadence: 16.00 commits/day
    Average message score: 3.0 / 3
    History: tidy

## Scoring rules

Каждое сообщение коммита получает от 0 до 3 баллов:
- +1 если в сообщении 3 и больше слов
- +1 если оно начинается с типа: add, update, fix, delete, docs, refactor
- +1 если оно упоминает файл (слово с точкой, например health.go) или тему (readme, branch, merge, conflict, workflow, reflection, feature, analyzer)

Средний балл 2.5 и выше даёт tidy, от 1.5 до 2.4 даёт acceptable, ниже 1.5 даёт messy.