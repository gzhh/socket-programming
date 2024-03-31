### Wiki
- https://en.wikipedia.org/wiki/Network_socket
- https://en.wikipedia.org/wiki/Berkeley_sockets
- https://man7.org/linux/man-pages/man2/socket.2.html

### 重点函数
- socket
- bind
- listen
- connect
- accept
- send
- close
- select
- poll

- getsockopt
- setsockopt
- gethostbyname
- gethostbyaddr

### 常量
**domain**
- AF_INET for network protocol IPv4 (IPv4-only)
- AF_INET6 for IPv6 (and in some cases, backward compatible with IPv4)
- AF_UNIX for local socket (using a special filesystem node)

**type**
- SOCK_STREAM (reliable stream-oriented service or Stream Sockets)
- SOCK_DGRAM (datagram service or Datagram Sockets)
- SOCK_SEQPACKET (reliable sequenced packet service)
- SOCK_RAW (raw protocols atop the network layer)

**protocol**
- IPPROTO_TCP
- IPPROTO_UDP
- IPPROTO_SCTP
- IPPROTO_DCCP

### Ref
- https://github.com/shineyr/Socket
- https://www.cnblogs.com/skynet/archive/2010/12/12/1903949.html
