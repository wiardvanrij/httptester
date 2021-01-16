# Simple http tester 

Either build it yourself or use the docker image: https://hub.docker.com/repository/docker/wiardvanrij/httptester

# Usage

`main.go -seconds=1 http://your-endpoint.whatever https://a-second.one`

The flag `seconds` is optional and defaults to `1`. 

# metrics

This application outputs a metric: `http_stats` which consists of:
- url
- http status
- type of http test

Example:

```
http_stats{status="200",type="Connect",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="Connect",url="https://sysrant.com",quantile="0.9"} 50
http_stats{status="200",type="Connect",url="https://sysrant.com",quantile="0.99"} 50
http_stats_sum{status="200",type="Connect",url="https://sysrant.com"} 50
http_stats_count{status="200",type="Connect",url="https://sysrant.com"} 4
http_stats{status="200",type="DNSLookup",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="DNSLookup",url="https://sysrant.com",quantile="0.9"} 17
http_stats{status="200",type="DNSLookup",url="https://sysrant.com",quantile="0.99"} 17
http_stats_sum{status="200",type="DNSLookup",url="https://sysrant.com"} 17
http_stats_count{status="200",type="DNSLookup",url="https://sysrant.com"} 4
http_stats{status="200",type="NameLookup",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="NameLookup",url="https://sysrant.com",quantile="0.9"} 17
http_stats{status="200",type="NameLookup",url="https://sysrant.com",quantile="0.99"} 17
http_stats_sum{status="200",type="NameLookup",url="https://sysrant.com"} 17
http_stats_count{status="200",type="NameLookup",url="https://sysrant.com"} 4
http_stats{status="200",type="Pretransfer",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="Pretransfer",url="https://sysrant.com",quantile="0.9"} 222
http_stats{status="200",type="Pretransfer",url="https://sysrant.com",quantile="0.99"} 222
http_stats_sum{status="200",type="Pretransfer",url="https://sysrant.com"} 222
http_stats_count{status="200",type="Pretransfer",url="https://sysrant.com"} 4
http_stats{status="200",type="ServerProcessing",url="https://sysrant.com",quantile="0.5"} 91
http_stats{status="200",type="ServerProcessing",url="https://sysrant.com",quantile="0.9"} 346
http_stats{status="200",type="ServerProcessing",url="https://sysrant.com",quantile="0.99"} 346
http_stats_sum{status="200",type="ServerProcessing",url="https://sysrant.com"} 621
http_stats_count{status="200",type="ServerProcessing",url="https://sysrant.com"} 4
http_stats{status="200",type="StartTransfer",url="https://sysrant.com",quantile="0.5"} 91
http_stats{status="200",type="StartTransfer",url="https://sysrant.com",quantile="0.9"} 568
http_stats{status="200",type="StartTransfer",url="https://sysrant.com",quantile="0.99"} 568
http_stats_sum{status="200",type="StartTransfer",url="https://sysrant.com"} 843
http_stats_count{status="200",type="StartTransfer",url="https://sysrant.com"} 4
http_stats{status="200",type="TCPConnection",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="TCPConnection",url="https://sysrant.com",quantile="0.9"} 32
http_stats{status="200",type="TCPConnection",url="https://sysrant.com",quantile="0.99"} 32
http_stats_sum{status="200",type="TCPConnection",url="https://sysrant.com"} 32
http_stats_count{status="200",type="TCPConnection",url="https://sysrant.com"} 4
http_stats{status="200",type="TLSHandshake",url="https://sysrant.com",quantile="0.5"} 0
http_stats{status="200",type="TLSHandshake",url="https://sysrant.com",quantile="0.9"} 171
http_stats{status="200",type="TLSHandshake",url="https://sysrant.com",quantile="0.99"} 171
http_stats_sum{status="200",type="TLSHandshake",url="https://sysrant.com"} 171
http_stats_count{status="200",type="TLSHandshake",url="https://sysrant.com"} 4
http_stats{status="200",type="Total",url="https://sysrant.com",quantile="0.5"} 91
http_stats{status="200",type="Total",url="https://sysrant.com",quantile="0.9"} 570
http_stats{status="200",type="Total",url="https://sysrant.com",quantile="0.99"} 570
http_stats_sum{status="200",type="Total",url="https://sysrant.com"} 846
http_stats_count{status="200",type="Total",url="https://sysrant.com"} 4
```