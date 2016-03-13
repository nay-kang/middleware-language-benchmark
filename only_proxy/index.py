import SimpleHTTPServer
import SocketServer
import httplib

PORT = 18083

class APIHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):
	def do_GET(self):
		self.send_response(200)
		self.end_headers()
		conn = httplib.HTTPConnection("localhost:18082")
		conn.request("GET","/product_list.json")
		response = conn.getresponse()
		self.wfile.write(response.read())
		self.wfile.flush()
		self.wfile.close()

api_handler = APIHandler

httpd = SocketServer.TCPServer(("", PORT), api_handler)

print "serving at port", PORT
httpd.serve_forever()