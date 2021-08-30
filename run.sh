

#if [ $1 = "test" ]; then
#    docker-compose -f docker-compose.test.yml up 
#elif [ $1 = "live" ]; then
#    docker-compose up
#else
#    echo "Invalid argument. Must be test OR live  ($1)."
#fi


if [ $1 = "database" ]; then
    docker-compose up unico-mysql &
elif [ $1 = "app" ]; then
    go run main.go 
elif [ $1 = "test" ]; then
    go test ./controllers & 
else
    echo "Invalid argument. Must be database OR app OR test  ($1)."
fi
