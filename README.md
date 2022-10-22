# base64lines

base64 encode (or decode) line by line

---

it's like the script below, but a little faster 

```sh
cat example.txt | xargs -L1 -- sh -c 'echo $@ | tr -d "\n" | base64' _
```
