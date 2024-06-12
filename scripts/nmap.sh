#!/bin/bash

nmap -sV $1
nmap -Pn --script vuln $1
