#Esta basado en postgres 10.3
FROM postgres:10.3
#Copy los archivos descritos en la linea 4 y 8
COPY up.sql /docker-entrypoint-initdb.d/1.sql
#Directorio de trabajo
WORKDIR /usr/src
COPY [".", "usr/src/"]
#Corre este comando en el contenedor
CMD ["postgres"]

