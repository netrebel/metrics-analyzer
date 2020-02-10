# metrics analyzer

Verifies if there's a key in prometheus.txt that doesn't exist in exporter-merger.txt

Required files:
- [files/exporter-merger.txt](files/exporter-merger.txt) dump from [exporter-merger](https://github.com/rebuy-de/exporter-merger) (/metrics)
- [files/prometheus.txt](files/prometheus.txt) dump from App prometheus metrics (/prometheus)

# How to run

`./metrics-analyzer`

# How to rebuild

`go build`