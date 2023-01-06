# sau

Sau is a tool that organizes URLs by subdomain. It creates a directory for each subdomain and stores all the URLs
belonging to that subdomain in a file.

### Installation

```
go install github.com/eriktok/sau@latest
```

### Usage

```
echo "example.com" | sau 
cat urls.txt | sau 
```

If you want to combine the Sau tool with  [Waybackurls](https://github.com/tomnomnom/waybackurls), you can use the
waybackurls tool to generate a list of URLs for a
particular domain and then use Sau to sort those URLs by subdomain. This would allow you to easily organize and analyze
the URLs for a domain by subdomain.

```
echo example.com | waybackurls | sau 
echo example.com | waybackurls | httpx | sau 
```