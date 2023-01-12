# base64lines

base64 encode (or decode) line by line

---

it's like the script below, but a little faster 

```sh
while read cookie; do
    echo "$cookie" | base64
done < cookies
```
