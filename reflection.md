# Reflection

Коммит для меня как сдать работу на хранение: у каждого коммита свой код (номер), и к любому снимку можно вернуться. Ветки нужны, чтобы было удобно работать, как с черновиком, не трогая основную ветку main. Самым сложным для меня было найти и настроить токен для GitHub, так как я до этого с ним не работала. Конфликт оказался не страшным: удобно, что VS Code сразу показывает разницу между версиями, и легко выбрать то, что нужно, не запутаешься. Я возьму привычку всё время проверять, где я нахожусь, через pwd, так как часто создаю папку не там.

## Bonus: Prompted Commits

Prompt (одинаковый для всех): Write a git commit message for this diff. Format: "type: short description", where type is one of add, update, fix, delete, docs, refactor. Diff: <вывод git diff --staged>

1. AI: docs: add instructions for running the program and tests
   Мой: docs: add instructions for running the program and tests in README
2. AI: refactor: simplify health score label condition
   Мой: refactor: simplify code in health.go
3. AI: test: add empty commit message case
   Мой: update: add empty message case in health_test.go
4. AI: docs: document Repo Health Analyzer functionality and usage
   Мой: docs: document Repo Health Analyzer functionality in workflow.md
5. AI: docs: document AI-assisted commit messages and reflection
   Мой: docs: add Prompted Commits bonus to reflection.md

   ИИ правильно соблюдал структуру "тип: описание", но не знал контекста проекта и почти никогда не называл изменённый файл, хотя имя файла было в diff, поэтому по его сообщениям трудно понять, где искать изменения.То есть мои сообщения были точнее, потому что я знаю правила проекта.
