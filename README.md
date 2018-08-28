# remap - замена значений в файле по мапке

remap решает одну задачу - замена всех вхождений каких-то ключей в файле.

Например, мы имеем файл с конфигурацией для подключения к какому-то сервису:

**config.php**
```
return [
    "service" => [
        "apiKey" => "<api_key>",
        "ignoreEmails" => "<admin_email>,<developer_email>",
    ],
],
```

Чтобы заменить все значения в нем, нам нужен файл с мапкой для плейсхолдеров:

**map.conf**
```
<api_key> = mu super api key
<admin_email> = admin@example.com
<developer_email> = developer@example.com
```

А дальше просто нужно вызвать команду для замены всех плейсхолдеров:

```
remap map.conf config.php
```

Результат работы будет выведен на stdout. Если нужно его сохранить в файл, то можно перенаправить.
