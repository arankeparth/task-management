FROM mysql:latest

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=parth123
ENV MYSQL_USER=parth
ENV MYSQL_PASSWORD=19122000

# Copy database initialization script
COPY init.sql /docker-entrypoint-initdb.d/


# Expose MySQL port
EXPOSE 3306
