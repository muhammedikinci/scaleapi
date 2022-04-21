<br />
<p align="center">
  <h3 align="center">Scaleflix</h3>
  <p align="center">
    Movie and Series API Project With Golang
    <br />
    <br />
  </p>
</p>

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#installation-without-docker">Installation Without Docker</a></li>
      </ul>
    </li>
    <li><a href="#api-reference">API Reference</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## Getting Started

### Prerequisites

* Golang
* Docker

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/muhammedikinci/scaleapi
   cd ./scaleapi
   ```
2. Run docker compose file (with data)
   ```sh
   docker compose -p "scaleflix_with_data" -f .\docker-compose.crawler.yml up
   ```

- API and Crawler can be restarted many times until database is ready to accept the connection
  ```yaml
  restart: on-failure
  ```

- Crawler project will terminate automatically after writing all movies and series that get from other APIs to the database

- If you want to start the project without Crawler, you can use the base docker-compose file in the project
    ```sh
   docker compose -p "scaleflix_without_data" -f .\docker-compose.yml up
   ```

### Installation Without Docker

This action needs the installation of PostgreSQL manually.

1. Clone the repo
   ```sh
   git clone https://github.com/muhammedikinci/scaleapi
   cd ./scaleapi
   ```
2. Install Go dependencies
   ```sh
   go mod download
   ```

3. Start PostgreSQL Service and Create `scaleflix` database

4. Start API without build
   ```sh
   go run ./cmd/server/.
   ```

5. Start crawler without build
   ```sh
   go run ./cmd/crawler/.
   ```

## API Reference

### POST /register

```sh
curl --request POST \
  --url http://localhost:8080/register \
  --header 'Content-Type: application/json' \
  --data '{
	"username":"muhammed",
	"password":"1234"
    }'
```

### POST /login

```sh
curl --request POST \
  --url http://localhost:8080/login \
  --header 'Content-Type: application/json' \
  --data '{
	"username":"muhammed",
	"password":"1234"
    }'
```

### GET /movies

```sh
curl --request GET \
  --url http://localhost:8080/movies \
  --header 'Authorization: Bearer {{TOKEN}}' \
  --header 'Content-Type: application/json'
```

- GET /movies/:id
- GET /movies/filter?title={{title}}&genre={{genre}}

### POST /movies (ADMIN REQUIRED)

```sh
curl --request POST \
  --url http://localhost:8080/movies \
  --header 'Authorization: Bearer {{TOKEN}}' \
  --header 'Content-Type: application/json' \
  --data '{
		"title": "Edge of Tomorrow",
		"image": "https://m.media-amazon.com/images/M/MV5BMTc5OTk4MTM3M15BMl5BanBnXkFtZTgwODcxNjg3MDE@._V1_SX300.jpg",
		"description": "A soldier fighting aliens gets to relive the same day over and over again, the day restarting every time he dies.",
		"rating": 7.9,
		"release_date": "06 Jun 2014",
		"director": "Doug Liman",
		"writer": "Christopher McQuarrie, Jez Butterworth, John-Henry Butterworth",
		"stars": "Tom Cruise, Emily Blunt, Bill Paxton",
		"duration": "113 min",
		"imdb_id": "tt1631867",
		"year": 2014,
		"genre": "Action, Adventure, Sci-Fi"
	}'
```

### GET /series

```sh
curl --request GET \
  --url http://localhost:8080/series \
  --header 'Authorization: Bearer {{TOKEN}}' \
  --header 'Content-Type: application/json'
```

- GET /series/:id
- GET /series/:id/seasons
- GET /series/:id/seasons/:season_id
- GET /series/filter?title={{title}}&genre={{genre}}

### POST /series (ADMIN REQUIRED)

```sh
curl --request POST \
  --url http://localhost:8080/series \
  --header 'Authorization: Bearer {{TOKEN}}' \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Under the Dome",
	"image": "https://static.tvmaze.com/uploads/images/original_untouched/81/202627.jpg",
	"description": "<p><b>Under the Dome</b> is the story of a small town that is suddenly and inexplicably sealed off from the rest of the world by an enormous transparent dome. The town'\''s inhabitants must deal with surviving the post-apocalyptic conditions while searching for answers about the dome, where it came from and if and when it will go away.</p>",
	"rating": 6.5,
	"release_date": "2013-06-24",
	"director": "",
	"writer": "",
	"stars": "",
	"imdb_id": "tt1553656",
	"year": 0,
	"genre": "Drama,Science-Fiction,Thriller"
    }'
```

### POST /favorite_movie/:id

```sh
curl --request POST \
  --url http://localhost:8080/favorite_movie/{{movie_id}} \
  --header 'Authorization: Bearer {{TOKEN}}'
```

### POST /favorite_serie/:id

```sh
curl --request POST \
  --url http://localhost:8080/favorite_serie/{{serie_id}} \
  --header 'Authorization: Bearer {{TOKEN}}'
```

### GET /favorite

```sh
curl --request GET \
  --url http://localhost:8080/favorites \
  --header 'Authorization: Bearer {{TOKEN}}'
```

## Contact

Muhammed İKİNCİ - muhammedikinci@outlook.com