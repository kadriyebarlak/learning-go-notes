vi Dockerfile

```
FROM python:2
COPY . /app
WORKDIR /app -> do cd 

RUN pip install -r requirements.txt

ENTRYPOINT python image_resizer.py
```




docker build --tag deneme .
docker images
deneme 920MB


 