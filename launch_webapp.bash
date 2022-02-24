#!/usr/bin/bash

# launches the dvpn node with default log file
./dvpn-node webapp -config wconfig.json |& tee /tmp/logger_dvpn.log
