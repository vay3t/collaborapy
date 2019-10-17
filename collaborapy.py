#!/usr/bin/env python3
"""
Very simple HTTP server in python for logging requests
Usage::
    python3 collaborapy.py [<port>]
"""
from http.server import BaseHTTPRequestHandler, HTTPServer

class S(BaseHTTPRequestHandler):
    server_version = "Microsoft-IIS/6.0"
    sys_version = ""

    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()
        print("-"*50)

    def do_GET(self):
        print("-"*50)
        print("{0} {1} {2}\n{3}".format(str(self.command), str(self.path),
                str(self.request_version), str(self.headers)))
        self._set_response()

    def do_POST(self):
        print("-"*50)
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        print("{0} {1} {2}\n{3}\n{4}\n".format(str(self.command), str(self.path),
                str(self.request_version), str(self.headers), post_data.decode('utf-8')))
        self._set_response()

def run(server_class=HTTPServer, handler_class=S, port=8080):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print('Starting httpd...\n')
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()
    print('Stopping httpd...\n')

if __name__ == '__main__':
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()
