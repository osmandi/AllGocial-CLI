# AllGocial-CLI

Es un prototipo para publicar Tweets directamente de la consola. Destinado a formar parte de [AllGocial](https://allgocial.now.sh)

Version: 0.4

# Requisitos

- Tener Go instalado en tu sistema.
- Buscar tus keys en [apps.twitter.com](https://apps.twitter.com) para consumir el API de Twitter.

# Instalación

Descargar el código fuente
```
git clone https://github.com/osmandi/AllGocial-CLI
```


Setear los key de la aplicación twitter (ejemplo en Linux)
```
export CONSUMERKEY=LO-OBTIENES-DESDE-EL-API
export CONSUMERSECRET=LO-OBTIENES-DESDE-EL-API
export ACCESSTOKEN=LO-OBTIENES-DESDE-EL-API
export ACCESSTOKENSECRET=LO-OBTIENES-DESDE-EL-API
```

# Uso
```
go run main.go -h -> Ayuda de comandos e instrucciones.

go run main.go message "Tweet a publicar" -> Publica tu mensaje.

```

> Si creas un alias o bien lo instalas, podrás usarlo desde cualquier path. Ejemplo: **allgocial message "este es un tweet"**.

Si el tweet fue o no publiciado avisará por consola.

# Nota importante

 Las claves de acceso de **AllGocial** no serán públicas, si quieres usarlo, debes obtenerlas por ti mismo en [apps.twitter.com](https://apps.twitter.com). 

O bien, esperar a que tenga el primer release del Back de AllGocial ✌
