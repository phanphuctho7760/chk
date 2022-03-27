### How to run

1. Run command ```go mod vendor```
2. Run command ```cp .\.env.example .\.env``` (if running on window) or ```cp .env.example .env``` on macos
3. Run command ```docker compose up -d --build```
4. Run command ```docker compose up -d```
5. Run command ```docker compose restart```

### Sample curl command to call API

1. API create price history from csv upload:

    ```
    curl --request POST 'localhost:9999/data' \--form 'file=@"/price_histories.csv"'
    ```
2. API get price history

   ```
   curl --request GET 'localhost:9999/data?page=1&limit=10'
   ```