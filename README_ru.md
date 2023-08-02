# Image Editor API
![go-version](https://img.shields.io/github/go-mod/go-version/NooFreeNames/ImageEditor?style=flat-square)

Сайт представляет собой удобный инструмент для редактирования изображений. С помощью данного сайта пользователи могут обрезать и фильтровать изображения в форматах JPG и PNG. Особенностью данного сайта является использование горутин для обработки изображений, что позволяет сократить время ожидания и ускорить процесс редактирования. Благодаря простому и интуитивно понятному интерфейсу, пользователи могут быстро и легко редактировать свои изображения, получая при этом высокое качество результата.

## Используемые технологии

- html
- css
- js
- bootstrap 5
- golang
- godotenv
- testify

## Установка

1. Клонировать репозиторий: `git clone https://github.com/your-username/image-editor-api.git`.
2. Перейдите в директорию проекта.
3. Установка необходимых модулей: `go mod download`.
4. Запустите сервер: `go run ./cmd/main.go`.

## Настройка сервера

Сервер настраивается с помощью файла ./config/.env.

Содержимое файла:
```properties
SERVER_HOST=127.0.0.1
SERVER_PORT=8080
SITE_DIR=web
```

* `SERVER_HOST` – адрес хоста, на котором будет запущен сервер.
* `SERVER_PORT` – порт, на котором будет запущен сервер.
* `SITE_DIR` – директория, в которой расположены файлы сайта.

## Использование

### Как открыть cайт?

После запуска сервера перейдите по ссылке которая отобразится в консоли.

![Ссылка на сайт](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/0.png)<br>

У вас должна открыться главная страница:

![Сайт](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/1.png)

### Изменение темы сайта

Чтобы изменить тему, нажмите на зеленую кнопку в правом верхнем углу.

![Смена темы](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/2.png)<br>
![Светлая тема](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/3.png)

### Загрузка изображения

Чтобы загрузить изображение, нажмите на поле выбора файла и загрузите желаемое изображение.

![Загрузка изображения](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/4.png)<br>
![Изображение загружено](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/5.png)

### Обрезка изобрадения

Чтобы обрезать изображение, укажите новый размер и положение обрезки. Затем нажмите Submit.

![Обрезка изображения](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/6.png)<br>
![Изображение отбрезано](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/7.png)

### Фильтрация изобрадения

Чтобы применить фильтр к изображению, выберите его из выпадающего списка и нажмите Submit.

![Фильтрация изображения](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/8.png)<br>
![Изображение отфильтровано](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/9.png)