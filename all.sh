
for i in $(seq 1 $1);
do
	go run ./main.go --day $i
done
