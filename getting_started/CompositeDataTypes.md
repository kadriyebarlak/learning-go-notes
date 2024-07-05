# Composite Data Types
- Complex data types. They are data types that put together that aggregate other data types.
# Arrays
- Fixed-length series of elements of a chosen type. It is known to compile time how big they are.
- Elements are initialized to zero value. In other languages like C when you create it, it is not initialized unless you write code to initialize it. If it is string, it is empty string.
    var x [5]int
    x[0] = 2
    fmt.Printf(x[1]) -> that will print 0
## Array Literal
- An array pre-defined with values
    var x [5]int = [5] {1,2,3,4,5} -> it has five literals inside the array literal
- Length of literal must be length of array
- ... for size in array literal infers size from number of initializers
    x := [...]int {1,2,3,4,5}
## Iterating Through Arrays
- Use a for loop with the range keyword
    x := [3]int {1,2,3}
    for i, v range x {
        fmt.Printf("ind %d, val %d", i, v)
    }
- range returns two values, index and value
# Slices
- A window on an underlying array
- Variable size, up to the whole array. You can increase the size of the slice.
- **Pointer** indicates the start of the slice
- **Length** is the number of elements in the slice
- **Capacity** is maximum number of elements. From start of slice to end of array.
    arr := [...]string {"a","b","c","d","e","f","g"}
    s1 := arr[1:3] -> b,c
    s2 := arr[2:5] -> c,d,e
- len() function returns the length
- cap() function returns the capacity
    a1 := [3]string ("a","b","c")
    sli1 := a1[0:1]
    len(sli1), cap(sli1) -> 1, 3
- **Accessing Slices** Writing to a slice changes underlying array. Overlapping slices refer to the same array elements
## Slice Literals
- Can be used to initialize  slice. when you initialize a slice that means you are creating an array.
- Creates the underlying array and creates a slice to reference it.
- Slice points to the start of the array, lenght is capacitiy. They are the same.
    sli := []int {1,2,3} -> anything in the brackets. Compiler knows that is a slice.

# Variable Slices
- This is the third way to make slice. You do not care about there is an array in the back but you want to initialize a slice to a particular size.
- Create a slice(and array) using **make()**
- 2-argument version: specify type and length/capacity
- Init to zero, length=capacity
    sli = make([]int, 10)
- 3-argument version: specify length and capacity separately
    sli = make([]int, 10, 15) -> that means that underlying array is bigger than the slice.
## Append
- Size of a slice can be increased by append()
- Adds elements to the end of a slice
- Inserts into underlying array
- Increases size of array if necessary. If you reach the limits of the array size, it will make a new underlying array that is bigger.
    sli = make([]int, 0, 3)
- Length of slice is 0
    sli = append(sli, 100)

# Hash Tables
- Key/value pairs
- **Hash function** is used to compute the slot or a key. A hash function process the key and generate the number of the slot and insert the value into it. You never call this hash function explicitly.
## Tradeoffs of Hash Tables
- Advantage: Faster lookup than lists. Constant-time vs linear time
- Advantage: Arbitrary keys. Not ints like slices or arrays.
- Disadvantage: May have collisions. Two keys hash to same slot. Linkedlist can be used to handle it. But the speed gets a little slower. Collisions are very rare.

# Maps
- Implementation of a hash table
- Use make() to create a map
    var idMap map[string]int
    idMap = make(map[string]int) -> create an empty map
- May define a map literal
    idMap := map[string]int {"joe":23}
- idMap["joe"] -> returns zero if key is not present
- delete(idMap, "joe")
## Map Functions
- Two-value assignment tests for existince of the key
    id, p := idMap["joe"] 
- id is value, p is presence of key(boolean)
- len(idMap) returns number values
## Iterating Through a Map
- Use a for loop with the range keyword. Two-value assignment with range.
    for key, val := range idMap{
        fmt.Println(key, value)
    }

# Structs
- An aggregate type meaning groups together other objects of arbitrary types into one object
- It is or organizational purposes.
- Bring together different variables that are related
    type struct Person {
        name string
        addr string
    }
    var p1 Person
- Each property is a field
## Initializing Structs
- Can use new()
- Initializes fields zero
    p1 := new(Person)
- Can initialize using a struct literal
    p1 := Person("name":"Joe", "addr":"a st.")


