#!/bin/bash

# deve diventare "2022-12-11/17:43,5.0,88.0"

cat temp.csv|while
 read linea
do
 #echo DA: $linea
 echo ${linea:0:4}-${linea:4:2}-${linea:6:2}/${linea:8:2}:${linea:10:2}${linea:12}
done