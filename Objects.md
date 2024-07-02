# Objects

- Go language is weakly object-oriented.
- Object-oriented programming really for for code organization. You organize your code by encapsulating your code.
- You group together data and functions which are related.
- Create a user-defined type which is specific to an application.
    Ex: ints have data(number of value) and functions(addition). Object-oriented programming is the same idea. You are creating types. specific to applications.

## Object example
- Implementing an application performing geometry in 3D.
- Many functions will operate on points.
- Each point has some data x,y and z coords.
- Points also have **functions** :DistToOrigin(), Quadrant()
- Point **class** defines data and functions. a general template.
- Point **objects** are instances of class. actual values.

## Objects in Go
**structs**
- No inheritance
- No constructors
- No generics
- Easier to code

# Concurrency
- Performance Limits: Cocurrency comes from the need for speed. A lot of the motivation not all.
- Moore's Law used to help performance. Number of transistors on a chip doubles every 18 months. It is not the case anymore recently. 
- More transistors used to lead to higher clock frequencies. The transistors got a little smaller and they would be closer to each other, you could increase the clock rate. zero to one, zero to one. 
- Power consumption/temperature constraints limit clock frequencies now. Moore's Law had to slow down. When you pack thesse transistors onto a chip, they generate heat. Every time they switch they consume power which generates heat. More frequently higher rate and more heat. The chip would physically melt. They got ans blowing over the chip. It distributes the heat.
- How do you get performance improvement?
## Parallelism 
- Number of cores still increases over time. Quad core four copies of the core. GPU have 1.000 processor core
- You can perform multiple tasks at the same time on different cores. If you have got 4 cores you, you can do 4 things at once.
- Difficulties at parallelism. When do tasks start/stop? What if one task needs data from another task? Do tasks conflict in memory variblae A and ovveride variable B in other task.
## concurrent Programming
- **Concurrency** is the management of multiple tasks at the same time. They might not actually be executing at the same time, maybe they are executing on a single core processor. They are alive at the same time. They could be executing at the same time if you had the resource. They need to be going on at the same time. Maybe one is paused while the others running but they are all alive at the same time at least from the user's perspective.
- Key requirement for large systems.
- Concurrent programming enables parallelism. Multiple tasks can be alive and communicating the same time then if you have the resources, multiple cores, multiple memory, then you can map them onto those parallel resources and get parallelism. The program is making decisions that allow things to run in parallel. Management of task execution. Communication between tasks. Synchronization between tasks. There are times where one task has to do something for the next task can start.
## Concurrency in Go
- Go includes concurrency primitives built-in to the language
- **Goroutines** represent concurrent tasks, basically a thread.
- **Channels** are used to communicate between tasks.
- **select** enables task synchronization
- Concurrency primitives are efficient and easy to use.

# Workspaces and Packages
- Three subdirectories: src contains source code files. pkg contains packages(libraries). bin contains executables.
- Workspace directory defined by **GOPATH** environment variable.
- GOPATH is defined during installation. Default gopath c:\Users\yourname\go
- Go tools assume that code is in GOPATH
- There must be one package called main. When you build the main package when you compile it, it makes an executable programme. When you build the non main packages, it does not make it executable for those. It will be incorporated into some other package.
- Main package needs a main() function
- main() is where code execution starts.

# Go Tool
- Import: When yo do an import, the Go Tool when it does a build, it has to find the imported packages. I searches through the directories specified by GOROOT and GOPATH environment variable. If you keep everything inside GOROOT and GOPATH, so inside your workspace, it will find them. If you decide you want to import some packege from some other place, maybe in a different directory, you have to change your GOROOT and GOPATH paths.  
- Go Tool: When you download Go, you get this Go Tool. A tool to manage Go source code. Several commands.
- **go build** - compiles the program. Arguments can be a list of packages or .go files. Creates an executable for the main package, same name as the first .go file. .exe suffix for executable in Windows
- **go doc**  - prints documentation for a package.
- **go fmt** - formats source code files.
- **go get** - downloads packages and installs them.
- **go list** - lists all installed packages.
- **go run** - compiles .go files and runs the executable.
- **go test** - runs tests using files ending in "_test.go"

# Variables
- Variables are basically data stored in memory somewhere.
- Most basic declaration var x, y int (keyword, name,type)
## Variable Types. 
- Integer only integral values. 
- Floating point fractional (decimal) values. 
- Strings byte(character) sequences
- the compiler need to know how much space needed for allocate. it needs to know how to convert it into the machine code instructions.
- Defining an alias for a type. type IDnum int, type Celsius float64. var pid IDnum. var temp Celsius
- var x int // x=0 string=""
- Short variable declarations x:= 100. Can only do this inside a function.
