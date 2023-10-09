#!/bin/bash

swiftc RabinKarp.swift

if [[ $? -eq 0 ]]
then
    ./RabinKarp
    rm -rf RabinKarp.exe RabinKarp.exp RabinKarp.lib
fi