# Scratchpay Challenge

## Starting dev environment
### `Running dev container`
```sh
$   docker-compose up -d
$   docker-compose exec app bash
All the commands bellow are done inside app container
```

### `Running application`
```sh
$   go run main.go
The app will start at localhost:8080.
```

`API usage`

| Route | Http verb | Query Params                                            | Description                                                                                                            |
| ----- | --------- | ------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
|`/`    | `GET`     | name (string), state (string), from (HH:MM), to (HH:MM) | `returns an array of json with practices that match the searching params. There are 4 query params as filters in the search that works cummulatively`                                                                       |

Request example: http://localhost:8080?name=scr&state=ca&from=10:00&to=12:00

You can use a collection to test the route. It´s in the folder ./collections [Download Insomnia](https://insomnia.rest/download)

### `Security`
```sh
Making requests to the api needs an authorization token in the request´s header (Authorization=Bearer token). The token is a secret in the .env variable APP_SECRET. The secret is in place of a hipothetic JWT token.
```

### `Testing application`
```sh
$   go test ./... -v
It will run all the tests and show its result
```

### `Testing coverage`
```sh
$   ./coverage/coverage.sh
It will run the tests and coverage tool. You can see the % of test coverage in the terminal and also opening the file cover.html that is generated in the coverage folder to see the test coverage in each file.
Ensure that you have execution permissions (chmod +x ./coverage/coverage.sh)
```

### `CI`
```sh
In a pull request on master branch will start the CI process. This process will run all tests in the application to ensure it´s OK beforing merging in the master branch. The repository is configured to allow merge only if all tests passed.
```

### Main Assumptions
```sh
For security reasons the api don´t return all practices if a search is performed with any parameter. It also assumes that when searching by availability it will search using the variables from and to independently if only one is supplied. As an example, if from=10:00 it will return all practices that have opening hour before that and clouse hour after that. The same occurs to to variable. If to=12:00 it will return all practices that have closure hour after that and opening hour before that.
```