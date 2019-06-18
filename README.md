# Dbash
`dbash` is a simple tool for gaining access to your Docker containers. Simply run `dbash` for a list of all running containers and just select the one you need shell access to.

### Usage
##### Container access list view
```sh
dbash
```
##### Container logs list view
```sh
dbash -l
```
##### Access a known container
```sh
dbash -c [CONTAINER IDENTIFIER]
```
##### Logs from a known container
```sh
dbash -l -c [CONTAINER IDENTIFIER]
```
The `CONTAINER IDENTIFIER` here can be a short or long Docker container ID hash, or a container name

### Build from Source
##### Check and install dependancies
```sh
./Configure
```
##### Build binary
```sh
Make
```
##### Install
```sh
Make install
```