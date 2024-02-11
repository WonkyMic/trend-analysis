```
tailwindcss -i ./dist/main.css -o ./dist/tailwind.css
```

```
docker build -t htmx-server:v1 .
docker run -p 8080:8080 -e APP_ENV=prod htmx-server:v1
```

TODO