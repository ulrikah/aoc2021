#! /usr/bin/env sh

DAY=$1
DIR=day$(printf "%02d" ${DAY})
AOCCOOKIE=$(cat .session)

curl \
  -fSL -o ./day${DAY}/input.txt \
  -H "Cookie: session=${AOCCOOKIE}" \
  https://adventofcode.com/2021/day/${DAY}/input
