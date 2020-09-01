# *Oxford* Dictionary Search

## A Telegram bot to query the Oxford English Dictionary 

###  Oxford Dictionary Search bot finds the:

* **Definition** 
* **Examples**
* **Lexical Category**
* **Language**

## Usage

### Environment Variables:

| Name             | Description                               | Required                   | Values 
| ---------------- | ----------------------------------------  | :------------------------: | ---------------
| `APP_ENV`        | The type of env app is running in         | `false` Fallback to `prod` | `dev` / `prod`
| `API_TOKEN`      | Bot's TG API Token                        | `true`                     | string
| `ADMIN_CHAT_ID`  | Bot Admin's ChatID                        | `true`                     | integer
| `APP_IDS`        | Dictionaries API APP IDs (corresponding)  | `true`                     | `ID-1;ID-2;ID-3`
| `APP_KEYS`       | Dictionaries API APP Keys (corresponding) | `true`                     | `KEY-1;KEY-2;KEY-3`

### For *all* the available senses of the searched word

