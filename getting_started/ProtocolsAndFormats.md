# Protocols and Formats

# RFCs (Requests for Comments)
- It is a definition on Internet protocols and formats, a standard definition.
- Example protocols: 
    HTML Hypertext Markup Language, 1866
    URI Uniform Resource Identifier. the addressing method. every server can interpret the address.
    HTTP Hypertext Transfer Protocol. that desribes how the messages should be transferred on the network. what information and what the headers. the content and so on.
    The standard protocols that are defined by RFCs
## Protocol Packages
- Golang provides packages for important RFCs.
- Functions which encode and decode protocol format.
- ***"net/http"*** : web communication protocol
    http.Get(www.uci.edu)
-***"net"*** TCP/IP and socket programming.
    TCP/IP basically define the internet. web-based communication, or FTP, Secure Shell all of them share in common is basic TCP/IP,UDP stack. Set of protocols that they all have to adhere to.
    net.Dial("tcp", "uci.edu:80")

# JSON JavaScript Object Notation
- Data format. RFC 7159. Format to represent structured information.
- Attribute-value pairs. This is a natural for a Struct or a map
- Basic value types. boolean, number, string, array, object
- Once you convert something into JSON, it is going to be represented all as Unicode characters. So it is human-readable.
- Fairly compact representation, not completely compact. If it were most compact then it would not be human-readable anymore. It is as small as you are going to get.
- Types can be combined recursively. Array of structs, struct in struct.
## JSON Marshalling
- Generating JSON representation from an object.
- Marshal() returns JSON representation as []byte
    p1 := Person(name="joe", addr:="a st.")
    barr, err := json.Marshal(p1)
## JSON Marshalling
- Unmarshal() converts a JSON []byte into a Go object 
    err := json.Unmarshal(barr, &p2)
- Pointer to Go object is passed to Unmarshal()
- Object must fit JSON []byte

# File Access, ioutil
- Linear access, not random access. They were actually stored on physical tapes(casette tapes or big old tapes in movies). The beginning of the file is at one point of he tape, the end of the file is way over here at the other point of the tape. So if you want to access it, you got to start, you read one point of the tape and you got to turn the tape and read the next point. Turn the tape read the next point and so on. This is mechanical operation of physically turning the tape to access the next piece of the file. If you want to just access something at the beginning then something at the end, that would waste a lot of time. The beginning is here then you have got to take a lot of time turning to get to the end. Mechanical delay. Instead you read it linearly. You will say, I am at the beginning and I read everything in that vicinity. That is how file access still is today. you do not necessarily have this linear access constraint. If you look at physical disks, they still have linear access. There is still tracks and you have to go through the whole track. There is still some linear access. A file might be in a random access device like a flash memory, but still the way that we access files in programs is as if it is a linear access device.
## Basic operations
1. Open - get handle for access
2. Read - read bytes into []byte
3. Write - write []byte into file
4. Close - release handle
5. Seek - move read/write head. Sometimes you really do need to skip to a certain point and seek does that. 
## ioutil File Read
- "io/ioutil" package has basic functions
    dat, err := ioutil.ReadFile("test.txt")
- dat is []byte filled with contents of ntire file
- Explicit open/close are not needed
- Large files cause  problem. Files can be gigantic. You can have files that really take up most of your memory. It takes it and reads it into RAM, into its main memory. Main memory is much more limited. RAM might only 8 gb. I can not run anything because I have used up all my memory.
## ioutil File Write
    dat = "Hello world"
    err := ioutil.WriteFile("outfile.txt", dat, 0777)
- Writes []byte to file
- Creates a file
- Unix-style permission bytes
- WriteFile is not flexible. You can not say Let me just dd onto the file. It just creates a file, dumps everything, the whole string or the whole byte array into that file and then closes the file.

# File Access, os
- You have more control like read a little bit, write a little bit.
- os.Open() opens a file. Returns a file descriptor. (A File struct basically). You can use to access the file.
- os.Close() closes a file.
- os.Read() reads from a file into a []byte. Fills the []byte. You can control how much you read.
- os.Write() writes a []byte into a file.
- Opening and reading
    f, err := os.Open("dt.txt")
    barr := make([]byte, 10)
    nb, err := f.Read(barr) -> Reads and fills barr. Returns number of bytes read.
    f.Close()
- Sometimes the number of byte should be equal to the size of the byte array, it does not have to be. Maybe the byte array is size 10 but there are only 5 byte sets left in the file then you read only 5. nb would equal 5.
- os File Create/Write
    f, err := os.Create("outfile.txt")
    barr := []byte {1,2,3}
    nb, err := f.Write(arr)
    nb, err := f.WriteString("Hi") -> Any Unicode sequence



