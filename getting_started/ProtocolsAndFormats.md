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
