### Service Product - Odoo
### Albertus Septian Angkuw

### Untuk Go Build Linux ,  
 export GOOS=linux
 export GOARCH=arm64
 go build

### Build docker
docker build -t odoo-product .
<br>
docker run --net odoo-ms-net --ip 172.18.0.52 --name odoo-product-go -p 1324:1324 -d odoo-product 
