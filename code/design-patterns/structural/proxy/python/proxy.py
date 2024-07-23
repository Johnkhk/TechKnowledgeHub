from typing import Dict


class ServiceInterface:
    def handle_request(self, url: str) -> str:
        pass


class WebServer(ServiceInterface):
    def handle_request(self, url: str) -> str:
        return f"WebServer handling request for: {url}"


class Nginx(ServiceInterface):
    def __init__(self, max_requests: int):
        self.web_server = WebServer()
        self.rate_limiter: Dict[str, int] = {}
        self.cache: Dict[str, str] = {}
        self.max_requests = max_requests

    def handle_request(self, url: str) -> str:
        # Check if rate limit exceeded
        if not self.control_access(url):
            return "Access denied"
        # Check if in cache
        if url in self.cache:
            return f"Returning cached response for: {url}"

        response = self.web_server.handle_request(url)
        self.cache_request(url, response)
        return response

    def control_access(self, url: str) -> bool:
        if url in self.rate_limiter:
            if self.rate_limiter[url] >= self.max_requests:
                return False
            self.rate_limiter[url] += 1
        else:
            self.rate_limiter[url] = 1
        return True

    def cache_request(self, url: str, response: str):
        self.cache[url] = response
