# url-shortener

> A service for shortening URLs, collecting and displaying statistics on visits  
> Related services: Bitly, Clck, Tinyurl, Goosu
## Demo
[Link to demo](https://url-shortener.ngrink.ru)

![url-shortener-preview](https://github.com/ngrink/url-shortener/assets/47951318/d489b839-ba0f-4ede-95ed-12d769ebc7e2)

## Features
- Registration, authorization
- Shortening links
- Custom links
- Visit statistics
- Access control

## Technology stack
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
REDIRECT_CACHE_CONTROL_MAX_AGE=600
```

5. Build and run application
```bash
go build ./cmd/url-shortener
./url-shortener
```

### Using prebuilt binaries
Download prebuilt binaries from the [latest release](https://github.com/ngrink/url-shortener/releases/latest), choose appropriate os/arch version. For example:
```bash
wget https://github.com/ngrink/url-shortener/releases/download/v0.2.0/url-shortener_Linux_x86_64.tar.gz
tar -zxf url-shortener_Linux_x86_64.tar.gz
./url-shortener
```
