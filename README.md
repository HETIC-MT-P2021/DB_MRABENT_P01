# Venus 

Venus is a basic API server that will provide endpoints to read data of a shop.

## Installation


TODO: 



```bash

$ git clone <this repo>
$ cd DB_MRABENT_P01
$ docker-compose up --build

```
You can import the DB by running :
```bash 
docker exec -i 27738ad5d9dd_db_mrabent_p01_db_1 sh -c 'exec mysql -ugoproject -p"goproject" image_goproject' < ./docker/data/database.sql
```

The API will be accessible at localhost:8080,  database on : 3309 .



## Technical choices

**Go version** : 1.15.2

**MySQL Server version** : 5.7.26  x(


## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
