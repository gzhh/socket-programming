### 官方文档
https://www.php.net/manual/zh/book.sockets.php

### 重点函数
- socket_create ( int $domain , int $type , int $protocol ) : resource
- socket_bind ( resource $socket , string $address , int $port = 0 ) : bool
- socket_listen ( Socket $socket , int $backlog = 0 ) : bool
- socket_connect ( Socket $socket , string $address , int|null $port = null ) : bool
- socket_accept ( Socket $socket ) : Socket|false
- socket_read — Reads a maximum of length bytes from a socket
- socket_write ( Socket $socket , string $data , int|null $length = null ) : int|false
- socket_close ( resource $socket ) : void

- socket_select ( array|null &$read , array|null &$write , array|null &$except , int|null $seconds , int $microseconds = 0 ) : int|false
- socket_recv ( resource $socket , string &$buf , int $len , int $flags ) : int
- socket_send ( Socket $socket , string $data , int $length , int $flags ) : int|false
- socket_strerror(socket_last_error($socket))

### 常用预定义常量
**domain**
AF_UNIX (integer)
AF_INET (integer)
AF_INET6 (integer)

**type**
SOCK_STREAM (integer)
SOCK_DGRAM (integer)

**protocol**
SOL_SOCKET (integer)
SOL_TCP (integer)
SOL_UDP (integer)
