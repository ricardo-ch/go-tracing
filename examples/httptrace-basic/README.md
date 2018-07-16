If you want to test, first start a zipkin server :

```bash
docker run -d -p 9411:9411 openzipkin/zipkin
```

then :

```bash
go run main.go
```

You can now go to find your trace on :

```html
http://localhost:9411/zipkin/
```


