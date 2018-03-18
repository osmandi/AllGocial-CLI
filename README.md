# AllGocial-CLI

Es un prototipo para publicar Tweets directamente de la consola. Destinado a formar parte de [AllGocial](https://allgocial.now.sh)

Para publicar:
```
go run allgocial.go -m "Tweet que quieras publicar"
```

Si el tweet fue o no publiciado avisará por consola.


# Nota importante

Las claves de acceso no son públicas, si quieres usarlo, debes obtenerlas por ti mismo en [apps.twitter.com](https://apps.twitter.com). Setearlas dentro del código o bien crear un archivo llamado **keys.txt**:

```
consumerKey,consumerSecret,accessToken,accessTokenSecret,
```

Todo pegado y separado por comas "," incluso al final.
