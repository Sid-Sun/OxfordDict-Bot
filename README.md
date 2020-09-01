# *Oxford* Dictionary Search

## A Telegram bot to query the Oxford English Dictionary 

###  Oxford Dictionary Search bot finds the:

* **Definition** 
* **Examples**
* **Lexical Category**
* **Language**

## Usage

### Environment Variables:

| Name             | Description                                         | Required                   | Values 
| ---------------- | --------------------------------------------------- | :------------------------: | ---------------
| `APP_ENV`        | The type of env app is running in                   | `false` Fallback to `prod` | `dev` / `prod`
| `API_TOKEN`      | Bot's TG API Token                                  | `true`                     | string
| `ADMIN_CHAT_ID`  | Bot Admin's ChatID                                  | `true`                     | integer
| `APP_IDS`        | Dictionaries API APP IDs (corresponding)            | `true`                     | `ID-1;ID-2;ID-3`
| `APP_KEYS`       | Dictionaries API APP Keys (corresponding)           | `true`                     | `KEY-1;KEY-2;KEY-3`
| `REDIS_HOST`     | Redis server's Host adress                          | `true`                     | string
| `REDIS_PORT`     | Redis server's port number                          | `true`                     | integer
| `REDIS_PASS`     | Redis password                                      | `true`                     | string
| `REDIS_DB`       | Redis DB ID                                         | `true`                     | integer
| `REDIS_SSL`      | Use SSL for redis connection (cliet supports HTTPS) | `true`                     | boolean

### For *all* the available senses of the searched word

