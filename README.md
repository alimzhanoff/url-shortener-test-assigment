# Тестовое задание на позицию стажера-бекендера в юнит Авто

## Задача:

Нужно сделать HTTP сервис для сокращения URL наподобие [Bitly](https://bitly.com/) и других сервисов.

UI не нужен, достаточно сделать JSON API сервис.  
Должна быть возможность: 
- сохранить короткое представление заданного URL
- перейти по сохраненному ранее короткому представлению и получить redirect на соответствующий исходный URL

### Требования:

- Язык программирования: Go/Python/PHP/Java/JavaScript
- Предоставить инструкцию по запуску приложения. В идеале (но не обязательно) – использовать контейнеризацию с возможностью запустить проект командой [`docker-compose up`](https://docs.docker.com/compose/)
- Требований к используемым технологиям нет - можно использовать любую БД для персистентности
- Код нужно выложить на github (просьба не делать форк этого репозитория, чтобы не плодить плагиат)

### Усложнения:

- Написаны тесты (постарайтесь достичь покрытия в 70% и больше)
- Добавлена валидация URL с проверкой корректности ссылки
- Добавлена возможность задавать кастомные ссылки, чтобы пользователь мог сделать их человекочитаемыми - [http://bit.ly/avito-auto-be](http://bit.ly/avito-auto-be)
- Проведено нагрузочное тестирование с целью понять, какую нагрузку на чтение может выдержать наш сервис
- Если вдруг будет желание, можно слепить простой UI и выложить сервис на бесплатный хостинг - Google Cloud, AWS и подобные. 
