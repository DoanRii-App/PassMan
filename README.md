# PassMan
```
Backend     - GO Lang
Frontend    - React
Encryption  - AES
Database    - MySQL
```

## Build
Change docker-build.sh, docker-rebuild.sh and docker-clean.sh permision
```
$ chmod a+x docker-build.sh
```
Run docker-build.sh, docker-rebuild.sh and docker-clean.sh with permision
```
$ sudo ./docker-build.sh
```

## Git Push
Change git-push.sh permision
```
$ chmod a+x git-push.sh
```
Run git-push.sh with args "Commit message" "Personal access token" "Git project file link without https://"
```
$ ./git-push.sh "first commit" "123asd" "github.com/DoanRii-App/PassMan.git"
```