#!/bin/sh

find . -exec rename 's/ethereum/dogeum/' {} ";"
find . -exec rename 's/ether/doge/' {} ";"
find . -exec rename 's/geth/gdoge/' {} ";"