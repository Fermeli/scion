### SCION
First of all SCION must be set up by following the instructions written [here](https://scion.docs.anapaya.net/en/latest/build/setup.html). 

Once SCION is setup, you can clone the repository:
```
git clone git@github.com:Fermeli/scion.git
```
Switch to the development branch:
```
cd scion
git checkout development_branch
```

Start the bazel server:
```
sudo ./scion.sh bazel_remote
```

Several topologies are available for the network. They can be found in the topology folder. Each .topo file is a different network topology. To set a cetain topology run the following command:
```
sudo ./scion.sh topology -c ./topology/example.topo 
```
where example.topo is the topology that you choose to run.
Finally, run the network:
```
 sudo ./scion.sh run
```

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
In the gen folder, a folder per AS is defined, each of them contains a prometheus config file prometheus.yml. Open a terminal at the path of the folder of the AS you are interested in and run the following command:
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
A Grafana dashboard is defined as a json file in the folder grafana_dashboard. It is called monitoring.json. Once you access Grafana, this dashboard can  be imported.

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
* -s which is the IPV4 address of the AS that performs the monitoring
* -address which specifies the IPV6 address of the AS on which the rate limit is set
* -ingress the ID of the ingress on which the rate limit is performed
* -egress the ID of the egress on which the rate limit is performed
* -cbs the CBS at which the rate limit is set
* -rate the CIR at which the rate limit is set

The addresses of the ASes can be found in gen/sciond_addresses.json.

### Example
In this example we test the rate limiter on the topology topology/tiny4.topo
Start the bazel server:
```
./scion.sh bazel_remote
```
Select topology tiny4.topo:
```
./scion.sh topology -c ./topology/tiny4.topo
```
Run the network:
```
./scion.sh run
```
Start the grafana server
```
systemctl daemon-reload
systemctl start grafana-server
```
Start the prometheus for the AS ff00_0_111.
```
prometheus --config.file=gen/ASff00_0_111/prometheus.yml
```
Open the Grafana at http://localhost:3000. Connect using "admin" as both username and password. Imort the dashboard grafana_dashboard/monitoring.json. Now you should be able to access the dashboard and to see some of the plots. You should see that AS ff00_0_110 sends packet to AS ff00_0_111 using ingressID 0 and egressID 41.

Set a rate limit from AS ff00_0_111 to AS ff00_0_110 on ingressID 0 and egressID 41:
```
go run go/scriptRL/main.go -s 127.0.0.17 -address 1-ff00:0:110 -cbs 400000 -rate 3000 -ingress 41 -egress 0
```

Open a terminal at the root of the scion-apps directory that you cloned and start the bandwidth test server in ASff00_0_111:
```
export SCION_DAEMON_ADDRESS=127.0.0.20:30255
sudo -E go run bwtester/bwtestserver/bwtestserver.go --listen 127.0.0.20:30255
```
Open another terminal at the same path and run the bandwidth test client in ASff00_0_110:
```
export SCION_DAEMON_ADDRESS=127.0.0.13:30255
sudo -E ./scion-bwtestclient -s 1-ff00:0:111,[127.0.0.20]:30255 -cs 1Mbps
```
Here we sent 1Mbps from AS ff00_0_110 to AS ff00_0_111. Some of the packets might be dropped by AS ff00_0_111. These information should be visible in the Grafana plots.
