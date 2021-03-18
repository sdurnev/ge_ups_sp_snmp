#!/usr/bin/env bash


#sudo socat UDP4-LISTEN:161,fork TCP4:localhost:5555 &
socat UDP4-LISTEN:5544,fork TCP4:localhost:5555 &