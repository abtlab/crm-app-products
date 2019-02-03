$name = "crm-app-products"

docker build -t $name .
docker run -it --rm --name $name -p 9000:9000 $name