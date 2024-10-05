1.Dockerfile olmadan
import      Import the contents from a tarball to create a filesystem image


```
GOOS=linux go build -o bin/app
./bin/app
less ./bin/app 

tar cvf example.tar bin 
````
macOS üzerinde linux binary si insa ettik.ELF linux un executable file formati


docker images
REPOSITORY                   TAG       IMAGE ID       CREATED             SIZE
example                      latest    60bf49f7cd55   3 seconds ago       3.23MB
example                      scratch   b86bacb8c24a   30 minutes ago      3.23MB
example                      alpine    ebc78bf36765   36 minutes ago      15.3MB

2. https://github.com/ko-build/ko
3. jib
4. use single-node GKE clusters just like Minikube