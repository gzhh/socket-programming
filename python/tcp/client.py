import socket

host = '127.0.0.1'
port = 2000
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))
print(s.recv(1024))
inpt = input('type anything and click enter... ')
s.send(inpt.encode())
print("the message has been sent")
