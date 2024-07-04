Note that UTF-8 (Unicode Transformation Format - 8) is an encoding standard for Unicode along with UTF-16 and UTF-32. This latter standard is a fixed-width encoding, in which each 32-bit value represents one Unicode code point. Some programmers have UTF-32 in mind when they think about "Unicode." 

# Pointers
- A Pointer is an address to data in memory
- &(ampersand) operator returns the address of a variable/function
- *(start) operator returns data at an address(dereferencing)
-   var x = 1 -> 1 is in memory and x is reference to it.
    var y int
    var ip *int -> ip is pointer to int
    ip = &x -> ip now points to x
    y = *ip -> y is now 1

## New
- Another way to create a variable
- new() function creates a variable and returns a pointer to the variable 
- variable is initialized to zero
    ptr := new(int)
    *ptr = 3 -> the value 3 is placed at the address specified by ptr

# Variable Scope
- The places in code where a variable can be accessed
- In Go, variable scoping is done using blocks.
## Blocks 
- A sequence of declarations and statements within matching brackets. {}. Including function definitions.
- Hierarchy of implicit blocks also. without the curly brackets
- Universe block - all Go source that is the biggest block.
- Package block - all source in a package.
- File block - all source in a file. all the source code in a single file is within the file block. Package can be composed of many files.
- if, for, switch - all code inside the statement
- Clause in switch or select - individual clauses each get a block.
# Lexical Scoping
- Go is lexically scoped using blocks. The relationship of one block being defined inside another block.
- bi >= bj if bj is defined inside bi. defined inside is transitive. (var x = 4 b1, func f b2, func g b3. b1 is file block) for x it looks inside b2, then what is the next bigger that I have defined inside?
- Variable accessible from block bj if 1.variable is declared in block bi 2.and bi>=bj

# Deallocating Memory
- Variables are referring to some data in memory.
- When a variable is no longer needed, it should be deallocated. Memory space is made available.
- Otherwise we will eventually run out of memory. func f() var x = 4. each call f() creates an integer. allocate 100 different spaces for this variable x.
- You can eat up all your spaces very quickly this is called memory leak.
## Stack vs Heap 
- about where the space is stored in memory.
- Stack is an area of memory that is dedicated to function calls. Local variables for a function are stored here. They are allocated in the stack area of the memory. Deallocated after function completes.
- Heap is a persistent region of memory. it does not go away. You have to explicitly deallocate it in another languages like C.
- Data on the heap must be deallocated when it is done being used. In most compiled languages like C, this is done manually. x = malloc(32); free(x); 32 bytes. Error-prone but fast. deallocate at the wrong time, forgetting to deallocate. In an interpreted language the interpreter does it and it can take time.

# Garbage Collection
- Hard to determine when a variable is no longer in use.
    func foo() *int { x :=1 return &x} 
    func main() { var y *int y = foo() }
- This is legal in Go. when func ends, it should be deallocated. But Returning a pointer to x. Pointer make it difficult. 
- Deallocation is complicated. One way of dealing with that is to have garbage collection.
## Garbage Collection
- Automatic tool that eals with deallocation. 
- In interpreted languages, this is done by the interpreter. JVM, Python Interpreter. when no more references, no more pointers to variables it deallocates them.
- Easy for the programmer
- Slow (need an interpreter)
## Garbage Collection in Go
- Go is a compiled language which enables garbage collection built-in to it. That is a unique feature. Implementation is fast. The Go compiler can figure out when to follow these points. It keep tracks of the pointers to a particular object.
- Compiler determines stack vs heap. It allocates stuff on the heap and the stack itself. You as a programmer do not have to determine I want to this on the heap or stack. The Go compiler it will put code in there at compile-time, it will figure out this needs to go heap or stack.
- Garbage collection is background.
- The act of GC does take some time. But it is an efficient implementation. It makes programming a lot easier.

# Basic Data Types
## Comments, Printing, Integers
## Integers
## Floating Point
- They are basic real numbers.
- float32 6 digits of precision. Approximately 6 decimal bit digits of precision
- float64 15 digits of precision.
- The more precision is probably better. precision errors are an issue.
- Decimal or as the exponent for 10. e2 means 10 to 2.
- Complex numbers represented as two floats : real and imaginary. var z complex128 = complex(2,3)
# ASCII and Unicode
- String are sequences of bytes.
- Each character has to be coded according to a standardized code. ASCII was the first accepted one.
- ASCII American Standart Code for Information Exchange. Just a character coding. Each character is associated with an 8-bit number. Capital A in ASCII is the number 41 in hexadecimal. 8-bit means it can maximum represent 256 possible characters.
- Unicode is a character code that is a 32-bit long code.
- UTF-8 is a subset of Unicode. It is a variable length. It can be 8-bit up to 32-bit. 8-bit UTF codes are same as ASCII.
- The default in Go is UTF-8
- Code point is a term for a Unicode character.
- Rune - just a term for a code point in Go. Capital A has a rune which is represented with 0x41
# Strings
- Strings are arbitrary sequences of bytes represented in UTF-8. Each byte is a rune represented as u UTF-8 code point.
- Read-only. You can not modify a string. You can make a new string that is a modified version of an existing string.
- Often meant to be printed.
- String literal - notated by double quotes.
- Each byte is a rune(UTF-8 code point).
- Strings are made of Unicode runes.
## Unicode Package
- Runes are divided into many different categories.
- Provides a set of functions to test categories of runes. 
    IsDigit(r rune) 
    IsSpace(r rune)
    IsLetter(r rune)
    IsLower(r rune)
    IsPunct(r rune) -is punctuation
- Some functions perform conversions.
    ToUpper(r rune)
    ToLower(r rune)
## Strings Package
- Functions to manipulate UTF-8 encoded strings
- Look at the whole string not individual 
- String search functions 
    Compare(a, b) -returns an integer comparing two strings lexicographically. 0 if a==b, -1 if a < b, and 1 if a>b 
    Contains(s, substr) -returns true if substring is inside s
    HasPrefix(s, prefix) -returns true is the string s begins with prefix
    Index(s, substr) -returns the index of the first instance of substr in s
## String Manipulation
- Strings are immutable, but modified strings are returned. It returns a new string that is modified.
- Replace(s, old, new, n) -replace returns a copy of the string s with the first n instances of old replaced by new.
- ToLower(s), ToUpper(s)
- TrimSpace(s) -returns a new string with all leading and trailing white space removed.
## Strconv Package 
- Conversions to and from string representations of basic data types.
- Atoi(s) -converts s to int. ASCII to integer
- Itoa(s) -converts int(base 10) to string
- FormatFloat(f, fmt, prec, bitSize) -converts floating point number to a string
- ParseFloat(si bitSize) -converts a string to a floating point number

# Constants
- It is an expression whose value is known at compile time. It never changes.
- Type is inferred from righthand side. const x = 1.3 const (y = 4, z="hi")
## iota
- Generate a set of related but distinct constants
- Often represents a property which has several distinct possible values. Days of the week. Months of the year.
- Constants must be different but actual value is not important.
- Like an enumerated type in other languages.
- Example 
    type Grades int
    const (
        A Grades iota
        B
        C
    )
- Each constant is assigned to a unique integer automatically. Starts at 1 and increments but you can not guarantee. You do not care what the actual values of the constants are. You just want the constant values to be different.

# Control Flow
- If else
- For loop Example: for i:=0; i < 10; i++ {fmt.PrintF("hi ")}
    Example: i = 0 for i < 10 {fmt.PrintF("hi ") i++}
- Switch/Case is a multi-way if statement
- Tagless Switch
    switch {
        case x > 1:
            ...
        case x < -1:
            ...
        default:
            ...
    }
- Break and Continue
    Break exits the containing loop
    Continue skips the rest of the current iteration

# Scan
- Scan reads user input
- Takes a pointer as an argument
- Typed data is written to pointer
- Returns number of scanned items
- Example
    var appleNum int
    fmt.Printf("Number of apples?)
    num, err := fmt.Scan(&appleNum) // scan function takes the number and puts it into the appleNum variable
    fmt.Printf(appleNum)











