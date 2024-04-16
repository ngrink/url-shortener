# url-shortener

> A service for shortening URLs, collecting and displaying statistics on visits  
> Related services: Bitly, Clck, Tinyurl, Goosu
## Demo
[Link to demo](https://url-shortener.ngrink.ru)

## Features
- Registration, authorization
- Shortening links
- Custom links
- Visit statistics

## Technology Stack
Linux, Nginx, REST API, SSR, JWT, Brypt, PostgreSQL  
Golang, Gorilla/mux, Gorm, Godotenv, html/template

## Architecture
This application has been developed using the principles of Clean Architecture and Dependency injection. The purpose of this architecture is to create flexible, extensible and easily maintained software by dividing the code into independent layers and layers.

#### Layers
- Controllers: handling requests, parsing data, forming response
- Services: perform bussiness logic
- Repositories: interact with database, ORM mapping
- Entities: describe fields

## Installation
To install and run this project, follow these steps:

1. Clone the repository:
```bash
git clone https://github.com/ngrink/url-shortener.git
```

2. Go to the project directory
```bash
cd url-shortener
```

3. Install dependencies
```bash
go mod tidy
```
4. Update env configuration
```bash
vim config/.env.local
```

```env
APP_PORT=7000
APP_HOST=http://localhost:7000

POSTGRES_HOST=<POSTGRES_HOST>
POSTGRES_PORT=<POSTGRES_PORT>
POSTGRES_USER=<POSTGRES_USER>
POSTGRES_PASSWORD=<POSTGRES_PASSWORD>
POSTGRES_DB=<POSTGRES_DB>

JWT_SECRET="jwUjak517ayqnJaBZHu8i9qybzz"
REDIRECT_CACHE_CONTROL_MAX_AGE=3600
```

5. Run application
```bash
go run ./cmd/url-shortener
```