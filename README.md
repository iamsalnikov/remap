How to use

```
remap <map> <file>
```

Remap just replaces all keys by it's values in a passed file and send it to stdout

For example, we have a file `file.conf`:

```
API_URL=<api_url>
API_KEY=<api_key>
```

We need to replace all placeholders (`<api_url>` and `<api_key>`) by specific value

We create a file with specific values - the map file (`map.conf`):

```
<api_url> = https://example.com
<api_key> = kksdfo93204jkljJKHJKsdf
```

Then we call remap:

```
remap map.conf file.conf
```

Result in stdout:

```
API_URL=https://example.com
API_KEY=kksdfo93204jkljJKHJKsdf
```

Remap just replaces all keys by it's values in a passed file.