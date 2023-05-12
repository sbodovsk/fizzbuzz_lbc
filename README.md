# FizzBuzz Test

## Test api
To test the api, please run :
```
go mod tidy
go run api.go
```
then, with your preferred http client tool (I use Postman), call this route with these params :
```
localhost:7001/fizzbuzz?string1=fizz&string2=buzz&int1=3&int2=5&limit=100
```


Or if you prefer good old `curl`, here it is (to save you some time) :
```
curl 'localhost:7001/fizzbuzz?string1=fizz&string2=buzz&int1=3&int2=5&limit=100'
```

## Docker
The simple dockerfile let's you build an image and use it.

```
docker build -t api .
docker run -p 7001:7001 api
```