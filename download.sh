source .env

(cd data && curl https://adventofcode.com/2024/day/$1/input --cookie "session=$SESSION" -o "$1")
