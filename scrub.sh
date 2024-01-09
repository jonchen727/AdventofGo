#!/bin/bash

x=$(find .  -type f -wholename '*/input.txt' |  sed 's/^.\{2\}//')

for i in $x; do
    echo $i 
    git filter-repo --invert-paths --path $i
done 