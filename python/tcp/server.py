import socket

host = ''
port = 2000
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((host, port))
s.listen(1)
while True:
    conn, addr = s.accept()
    print("Connection accepted from " + repr(addr[1]))

    conn.send("Server approved connection\n".encode())
    print(repr(addr[1]) + ": " + conn.recv(1026).decode())
    conn.close()
