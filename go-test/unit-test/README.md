## test coverage
* go test -cover
* go test -cover -coverprofile=cover.out
* go tool cover -html=cover.out -o coverage.html
* go test -bench . -cpu 2 -> 2 cekirdek uzerinde calistir
* go test -bench . -benchtime 10x -> birim zamanda degil 10 kez calissin
* go test -bench . -benchtime 1s -> 1 sn de calissin
* go test -bench . -benchtime 3x -count 5 -> 3 defa calissin 5 kere ardarda. 3 for calisiyor resetleniyor her sey(5 kere)
* go test -bench . -gcflags='-N -l' -> compiler optimizasyonlari kapatilir
* go test -c -o gzip_bench.test  -> compile et tekrarli calistirmak icin
* ./gzip_bench.test -test.bench . -test.benchmem -test.count 10 > result.txt

## PProf
* Testlerden ya da runtimeda uygulamadan alinan snapshot profilelari pprof ile incelenir.
* ./gzip_bench.test -test.bench . -test.benchmem -test.memprofile=mem.profile -test.cpuprofile=cpu.profile
* go tool pprof cpu.profile 
* go tool pprof mem.profile 
* list Zip
* go test -test.bench . -test.benchmem -test.memprofile=old.mem.profile -test.cpuprofile=old.cpu.profile
* go test -test.bench . -test.benchmem -test.memprofile=new.mem.profile -test.cpuprofile=new.cpu.profile
* go tool pprof -base old.mem.profile new.mem.profile
* pprof list Zip -> total kazanimi gosterir
* pprof web 
* top 10 en fazla cpu kullanan yerler
* peek Zip 

## Trace
* go test -bench . -trace=trace.out -> Trace ciktisi
