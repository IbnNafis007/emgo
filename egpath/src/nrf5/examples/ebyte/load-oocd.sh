#!/bin/sh

INTERFACE=stlink-v2
TARGET=nrf52
TRACECLKIN=4000000

. ../../../../../scripts/load-oocd.sh $@
