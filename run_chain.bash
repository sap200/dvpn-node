#!/usr/bin/bash

# remove previous data
vineyardd unsafe-reset-all
rm ~/.vineyard/config/genesis.json
rm ~/.vineyard/config/gentx/$(ls ~/.vineyard/config/gentx/)


# Initialize app
vineyardd init demo
cat ~/.vineyard/config/genesis.json
vineyardd keys list
vineyardd keys add alice --keyring-backend test
vineyardd keys show alice

# Make yourself a proper validator
vineyardd add-genesis-account alice 100000000stake
grep -A 2 -B 2 denom ~/.vineyard/config/genesis.json

# generate tx so, you own max
vineyardd gentx alice 70000000stake --chain-id vineyard
# collect all txs
vineyardd collect-gentxs

# start creating new blocks
vineyardd start --log_level error &

vineyardd keys add bob
vineyardd query bank balances $(./vineyardd keys show alice -a)
vineyardd tx bank send $(vineyardd keys show alice -a) $(./simd keys show bob -a) 10stake --chain-id vineyard
root@localhost:/home/workspace# nano run_chain 
root@localhost:/home/workspace# cat run_chain 
#!/usr/bin/bash

# remove previous data
vineyardd unsafe-reset-all
rm ~/.vineyard/config/genesis.json
rm ~/.vineyard/config/gentx/$(ls ~/.vineyard/config/gentx/)


# Initialize app
vineyardd init demo
cat ~/.vineyard/config/genesis.json
vineyardd keys list
vineyardd keys add alice --keyring-backend test
vineyardd keys show alice

# Make yourself a proper validator
vineyardd add-genesis-account alice 100000000stake
grep -A 2 -B 2 denom ~/.vineyard/config/genesis.json

# generate tx so, you own max
vineyardd gentx alice 70000000stake --chain-id vineyard
# collect all txs
vineyardd collect-gentxs

# start creating new blocks
vineyardd start --log_level error &

vineyardd keys add bob
vineyardd query bank balances $(./vineyardd keys show alice -a)
vineyardd tx bank send $(vineyardd keys show alice -a) $(./simd keys show bob -a) 10stake --chain-id vineyard
