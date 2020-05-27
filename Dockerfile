FROM ubuntu:latest

# Crete Opt Dir
CMD mkdir /opt/hotsPicker

# Copy files and directories from the application
COPY heroes.go /opt/hotsPicker/
COPY main.go /opt/hotsPicker/

# Tell Docker we are going to use this port
EXPOSE 80

CMD ./hotsPicker -t 