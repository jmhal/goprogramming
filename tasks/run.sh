#!/bin/bash

# Baixa a imagem do mongo
docker pull mongo

# Constrói a imagem do apidb
docker build -t apidb .

# Se os contêineres já existirem, vamos pará-los e removê-los
if docker ps -a -f "name=mongodb" | grep mongodb > /dev/null 
then
    docker stop mongodb
    docker rm mongodb
fi

if docker ps -a -f "name=apidb" | grep apidb > /dev/null 
then
    docker stop apidb
fi


# Criar o contêiner do mongo
docker run -d --name mongodb mongo

# Criar o contêiner da apidb
docker run -d --name apidb --rm -p 10000:10000 --link mongodb:mongo apidb