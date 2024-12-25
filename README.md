# Password Manager

[![Go](https://img.shields.io/badge/-Go-00ADD8?style=flat&logo=Go&logoColor=ffffff)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/-SQLite-003B57?style=flat&logo=SQLite&logoColor=ffffff)](https://www.sqlite.org/)
[![Gorilla Mux](https://img.shields.io/badge/-Gorilla%20Mux-74D269?style=flat&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTAwMCIgaGVpZ2h0PSI1MDAiIHZpZXdCb3g9IjAgMCAxMDAwIDUwMCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJNMCAyNSA0MDAgMjUgNDAwIDI1SDBWMjVaTTAgNDc1IDQwMCA0NzUgNDAwIDQ3NUgwVjQ3NVpNNTI4IDI1SDYwMFY1MDBINTQyVjI1WiIgZmlsbD0iIzQ4QkQ2QyIvPjxwYXRoIGQ9Ik02MDAgMEgzMDBDMTQwIDAgMCAxNDAgMCAzMDBWMjAwSDEwMEwxMDAgMzAwSDBWMzAwQzAgNDQwIDE0MCA1ODAgMzAwIDYwMEgyMDAgNDgwSDMwMFM0MjAwIDUyMDAgNTIwMCA0MzAwTDYwMCA0MjAwVjI4MDAgTDAgMjQwMFYzNzAwTDMwMCAzODAwVjI2MDBMMzYwMCAyNTgwbDAgMCIgZmlsbD0id2hpdGUiLz48L3N2Zz4=)](https://www.gorillatoolkit.org/)
[![Godotenv](https://img.shields.io/badge/-Godotenv-42b983?style=flat&logo=dotenv&logoColor=ffffff)](https://github.com/joho/godotenv)
[![AES-256](https://img.shields.io/badge/-AES--256-007ACC?style=flat&logo=lock&logoColor=white)](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

Password Manager — это веб-приложение для управления паролями. Оно позволяет:

- Создавать и хранить зашифрованные пароли для различных сервисов.
- Генерировать случайные пароли с кастомными настройками.
- Получать список сохранённых паролей.
- Искать пароли по имени сервиса.

---

## **Особенности проекта**

- Шифрование паролей с использованием AES-256.
- Генерация случайных паролей с поддержкой опций для цифр, символов и букв.
- REST API для управления паролями.
- Простая архитектура и понятный код.

---

## **Технологии**

- **Язык:** Go
- **Фреймворк:** Gorilla Mux
- **Шифрование:** AES-256
- **Работа с переменными окружения:** Godotenv
- **База данных:** SQLite

---

## **Как запустить проект**

### **1. Клонируйте репозиторий**

```bash
git clone https://github.com/just4fun-xd/password-manager.git
cd password-manager
```

### **2. Настройка переменных окружения**

Создайте файл `.env` в корневой директории и укажите ключ шифрования:
```plaintext
ENCRYPTION_KEY=my32characterlongencryptionkey1!
```

### **3. Установите зависимости**

Убедитесь, что у вас установлен Go версии 1.20 или выше. Установите необходимые зависимости:

```bash
go mod tidy
```

### **4. Запустите сервер**

Запустите сервер с помощью команды:

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`.

---

## **API Эндпоинты**

### **1. Создание нового пароля**

**POST /passwords**

- **Описание:** Сохраняет новый пароль.
- **Пример запроса:**
  ```json
  {
      "service_name": "Instagram",
      "username": "photo_lover",
      "password": "instapass999"
  }
  ```
- **Пример ответа:**
  ```json
  {
      "id": 1,
      "service_name": "Instagram",
      "username": "photo_lover",
      "password": "instapass999",
      "created_at": "2024-12-26T00:00:00Z"
  }
  ```

---

### **2. Получение всех паролей**

**GET /passwords**

- **Описание:** Возвращает все сохранённые пароли.
- **Пример ответа:**
  ```json
  [
    {
      "id": 1,
      "service_name": "Instagram",
      "username": "photo_lover",
      "password": "instapass999",
      "created_at": "2024-12-26T00:00:00Z"
    }
  ]
  ```

---

### **3. Поиск паролей**

**GET /search?service=Instagram**

- **Описание:** Ищет пароли по имени сервиса.
- **Пример ответа:**
  ```json
  [
    {
      "id": 1,
      "service_name": "Instagram",
      "username": "photo_lover",
      "password": "instapass999",
      "created_at": "2024-12-26T00:00:00Z"
    }
  ]
  ```

---

## **Архитектура проекта**

```
password-manager/
├── README.md
├── api
│   └── handlers.go
├── config
│   └── config.go
├── db
│   ├── db.go
│   ├── init.sql
│   └── repository.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── password.go
├── passwords.db
└── utils
    └── encryption.go
```

---

## **Как добавить новые функции**

1. **Добавление нового маршрута:**
   - Добавьте новый маршрут в `api/handlers.go`.
   - Реализуйте обработчик для маршрута.

2. **Добавление полей в таблицу:**
   - Измените структуру в `models/password.go`.
   - Добавьте соответствующую миграцию в SQL-скрипты.

3. **Подключение к другой базе данных:**
   - Измените строку подключения в `db/init.go` на нужный драйвер (например, PostgreSQL).

---

## **Идеи для улучшения**

- **Авторизация:** Добавить поддержку пользователей и аутентификацию.
- **Статистика:** Добавить эндпоинт для вывода статистики по паролям.
- **Истечение срока действия паролей:** Добавить возможность задавать срок действия.

---

## **Контакты**

Если у вас есть вопросы или предложения, свяжитесь со мной:

- **Email:** k.shalygin@yandex.ru
- **GitHub:** [github.com/just4fun-xd](https://github.com/just4fun-xd)

---

