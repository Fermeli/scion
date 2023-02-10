### SCION
First of all SCION must be set up by following the instructions can be found [here](https://scion.docs.anapaya.net/en/latest/build/setup.html). 

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

### Bandwidth tester
Installation:
```
git clone git@github.com:netsec-ethz/scion-apps.git
```
The instructions to set up the bandwidth tester can be found [here](https://github.com/netsec-ethz/scion-apps#installation).

Start Application:
Once the network is running with a certain topology, you need to pick one AS to be the client and one AS to be the server. A mapping from an AS IPV6 address to its IPV4 address can be found in gen/sciond_addresses.json. You also need to choose a port for the server and the client. Given servAddrIPV4, servAddrIPV6, clientAddrIPV4, clientAddrIPV6 the server address and the client addresses, and servPort and clientPort the ports that you chosse, open two terminals at the same path where you cloned the repository and run the following commands:

In the first terminal run the bandwidth server as follow:
```
cd scion-apps
export SCION_DAEMON_ADDRESS=servAddrIPV4:servPort
sudo -E go run bwtester/bwtestserver/bwtestserver.go --listen servAddrIPV4:servPort
```
Then, in the second terminal you can run the bandwidth client. Suppose that you want to send 1Mbps, run the following commands:
```
cd scion-apps
export SCION_DAEMON_ADDRESS=clientAddr:clientPort
sudo -E go run bwtester/bwtestclient/bwtestclient.go -s servAddrIPV6,[servAddrIPV4]:servPort -cs 1Mbps
```
### Setting rate limit
Once the network is running, rate limits can be set by using the script located in go/scriptRL/main.go.
The script takes the following argument:
* -s which is the IPV4 address of the AS which performs the monitoring
* -address which specifies the IPV6 address of the address on which the rate limit is set
* -ingress the ID of the ingress on which the rate limit is performed
* -egress the ID of the egress on which the rate limit is performed
* -cbs the CBS at which the rate limit is set
* -rate the CIR at which the rate limit is set

The addresses of the ASes can be found in gen/sciond_addresses.json.

