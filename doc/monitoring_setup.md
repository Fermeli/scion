### Prometheus
Version: 2.37 (LTS)
https://prometheus.io/docs/prometheus/2.37/getting_started/

Installation
```
# configuration directory:
# data directory:
sudo mkdir -p /var/lib/prometheus

wget https://github.com/prometheus/prometheus/releases/download/v2.37.1/prometheus-2.37.1.linux-amd64.tar.gz
tar xvfz prometheus-*.tar.gz
cd prometheus-2.37.1.linux-amd64

sudo mv prometheus promtool /usr/local/bin/
prometheus --version
```

Start Application:
```
prometheus --config.file=prometheus.yml
```

The prometheus metrics can be browsed on port `9090`.

### Grafana

Installation:
```
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
sudo add-apt-repository "deb https://packages.grafana.com/oss/deb stable main"
sudo apt update && sudo apt install grafana
```

Start Application:
```
sudo systemctl daemon-reload
sudo systemctl start grafana-server
sudo systemctl status grafana-server
```

The Grafana web server is reachable on port `3000`. Per default, use username `admin` with password `admin` to log in.

For Grafana, first the Prometheus server has to be added as data source.
