# Imagem para rodar o server do go
# docker build -t rutsatz/hello-go .
# docker run --rm -p 80:80 rutsatz/hello-go
# docker push rutsatz/hello-go

# docker build -t rutsatz/hello-go:v2 .
# docker push rutsatz/hello-go:v2

FROM golang:1.15
COPY . .
RUN go build -o server .
CMD ["./server"]
