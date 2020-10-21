# rustleyourjimmies
A grinds my gear expression generator

I recently thoutgh of the structure of expression that I call "Grind my gear expressions" It's simple enough that I'm pretty sure someone already thought about it
but the closest I could find was [this jacksfilms video](https://www.youtube.com/watch?v=C7fS3RsFuUE) 

Setting up the certs goes like this

``
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
``

![Why though](https://i.kym-cdn.com/photos/images/newsfeed/001/255/097/022.jpg "Why though")

I was bored and wanted to learn go
