FROM ubuntu:20.04

#Configure Timezone
#RUN export TZ=Asia/Calcutta
#RUN echo "Asia/Calcutta" | tee /etc/timezone && \
#    dpkg-reconfigure --frontend noninteractive tzdata
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Calcutta

#Install dependencies
RUN apt update; apt dist-upgrade -y; apt install -y vim apache2; 

#Print hello world message
RUN echo 'Hello World!' > /var/www/html/index.html

#Configure apache2
RUN echo '. /etc/apache2/envvars' > /root/run_apache.sh && \
    echo 'mkdir -p /var/run/apache2' >> /root/run_apache.sh && \
    echo 'mkdir -p /var/lock/apache2' >> /root/run_apache.sh && \
    echo '/usr/sbin/apache2 -D FOREGROUND' >> /root/run_apache.sh && \
    chmod 755 /root/run_apache.sh

EXPOSE 80

CMD /root/run_apache.sh 
