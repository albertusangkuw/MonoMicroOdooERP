### Service 10 - Odoo ,Product and Point of Sale
### Albertus Septian Angkuw

### Untuk Go Build Linux ,  
 export GOOS=linux
 export GOARCH=arm64
 go build

### Build docker
docker build -t service-artefact .
<br>
docker run --net odoo-ms-net --ip 172.18.0.51 --name odoo-service-10-go -p 1323:1323 -d service-artefact
