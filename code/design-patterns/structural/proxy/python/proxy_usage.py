from proxy import Nginx

nginx = Nginx(max_requests=3)

msg1 = nginx.handle_request("google.com")
msg2 = nginx.handle_request("google.com")
msg3 = nginx.handle_request("google.com")
assert msg1 == "WebServer handling request for: google.com"
assert msg2 == msg3 == "Returning cached response for: google.com"
msg4 = nginx.handle_request("google.com")
assert msg4 == "Access denied"
