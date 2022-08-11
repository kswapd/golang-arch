##proxy test:

```
curl -Lv  --proxy http://127.0.0.1:9999 https://www.jetbrains.com/intellij-repository/snapshots/com/jetbrains/gateway/BUILD/222.3739.24-CUSTOM-SNAPSHOT/BUILD-222.3739.24-CUSTOM-SNAPSHOT.txt

#Start as a server:
nohup ./easy_proxy -port 9999 2>&1 > proxy.log&
#Build for other OS:
GOOS=linux GOARCH=amd64 go build -o easy_proxy main.go
```