FROM scratch
MAINTAINER Thilina Piyasundara 

EXPOSE 8080
ADD goecho /

CMD ["/goecho"]
