```
FROM alpine
RUN apk add --no-cache curl

ENTRYPOINT ["curl"]
```

 1084  docker build -t curl .
 1085  docker images
 1086  docker tag curl gcr.io/kbarlak-public/curl


docker history curl
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
c284c6d72c07   3 minutes ago   ENTRYPOINT ["curl"]                             0B        buildkit.dockerfile.v0
<missing>      3 minutes ago   RUN /bin/sh -c apk add --no-cache curl # bui…   6.3MB     buildkit.dockerfile.v0
<missing>      2 weeks ago     /bin/sh -c #(nop)  CMD ["/bin/sh"]              0B        
<missing>      2 weeks ago     /bin/sh -c #(nop) ADD file:5758b97d8301c84a2…   8.46MB   


docker inspect curl
Parent 
ROOTFS -> tar ball lari indirmesi gerekiyor.


dive curl

crane manifest gcr.io/kbarlak-public/curl
crane blob gcr.io/kbarlak-public/curl@sha...
la
tar tvf layer.tar.gz | grep -v cert
