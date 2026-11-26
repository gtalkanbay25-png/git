# Workflow

## Branches
У меня было три ветки: main, feature и readme-update. main это основная ветка (чистовик), а остальные как черновики: в них я пробую изменения, а потом сливаю в main.

## Feature branch
Я создала ветку командой git switch -c feature, добавила в ней файл feature/feature.txt и сделала коммит. Потом перешла на main и сделала git merge feature. Конфликта не было, git просто передвинул main вперёд (fast-forward).

## Merge conflict
Я создала ветку readme-update и изменила в README вторую строку. Потом в main изменила ту же строку по-другому. При git merge readme-update случился конфликт. VS Code показал метки <<<<<<<, ======= и >>>>>>>, я выбрала нужный текст, вручную удалила метки и сделала коммит слияния.

## Remotes
У меня два удалённых репозитория: origin (Gitea школы) и github (GitHub). В Gitea отправляю командой git push, а на GitHub командой git push github main.

## Analyzer
Repo Health Analyzer читает историю через git log, ставит каждому сообщению коммита от 0 до 3 баллов и выводит итог: tidy, acceptable или messy. Логика лежит в internal/health, запуск через go run .