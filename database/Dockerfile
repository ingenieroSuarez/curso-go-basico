FROM postgres:10.3
# el nombre 1.sql es para que lo ejecute de primero ya que docker corre de manera alfabetica
COPY up.sql /docker-entrypoint-initdb.d/1.sql
# comando para correr postgres: 
CMD ["postgres"]