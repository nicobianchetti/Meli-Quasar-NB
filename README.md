# Meli-Quasar-NB
Challenge Mercado Libre 2021 - Operación Fuego de Quasar.

## Desafío 

Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que retorne
la fuente y contenido del mensaje de auxilio. Para esto, cuentas con tres satélites que te
permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada
satélite debido al campo de asteroides frente a la nave.

### Pre-requisitos 📋

_Para correr localmente la aplicación se debe tener instalado Docker y Docker-Compose, clonar repositorio y dentro del folder raíz ejecutar el siguiente comando de docker-compose (la aplicación está expuesta por defecto en el puerto :5000 de localhost y las peticiones deben pasar por el puerto :80) :_

```
git clone https://github.com/nicobianchetti/Meli-Quasar-NB.git
```

```
docker-compose --compatibility up --build
```

### Ejecución 🚀


_Host de la aplicación:_

```
http://144.126.217.32
```
_En Headers agregar:_

```
api-key : una-api-key-muy-segura
```

### Ejemplo de Request para nivel 2:

_Endporint POST /quasar/topsecret/:_

```
http://144.126.217.32/quasar/topsecret/
```

_Body:_

```
{
    "satellites":[
            {
                "name": "kenobi",
                "distance": 100.00,
                "message": ["este", "", "", "mensaje", ""]
            },
            {
                "name": "skywalker",
                "distance": 115.5 ,
                "message": ["", "es", "", "", "secreto"]
            },
            {
                "name": "sato",
                "distance": 142.7 ,
                "message": ["este", "", "un", "", ""]
            }

        ]
}
```
_Response:_

```
{
    "position": {
        "x": -487.2859125,
        "y": 1557.014225
    },
    "message": "este es un mensaje secreto"
}
```

Imagen Ejemplo Request Nivel 1

### Ejemplo de Request para nivel 3:

Considerar que el mensaje ahora debe poder recibirse en diferentes POST al nuevo servicio
/topsecret_split/, respetando la misma firma que antes.
Crear un nuevo servicio /topsecret_split/ que acepte POST y GET. En el GET la
respuesta deberá indicar la posición y el mensaje en caso que sea posible determinarlo y tener
la misma estructura del ejemplo del Nivel 2. Caso contrario, deberá responder un mensaje de
error indicando que no hay suficiente información.

_Solución:_

_Para resolver el nivel 3 se decidió referenciar un "user" por cada petición que llega, sea POST o GET. Es decir, se podrán recibir los diferentes POST de cada uno de los satélites asignados a un user cuyo dato se envía en los headers, y luego se hace un GET para el mismo user. Si ya se enviaron los datos de los tres satélites y los mismos son consisentes, se mostrará la información asociada de Coordenada y Mensaje para ese emisor y se procederá a limpiar los datos agregados para ese usuario de modo que se puedan volver a agregar 3 nuevos POST. En caso de que falte agregar datos de un satélite , o la información no sea consistente, la api devuelve un 404._

_Aclaración: por cada "user" se pueden enviar varios POST de cada satélite. Siempre se tomará en cuenta el último satélite de cada uno que se solicitó. Ejemplo: Se realiza primero un POST para el satélite "kenobi" con "distancia" = 100 y para el "user":Nicolás y en realidad la distancia era 200, si todavía no se procesó un GET y se vuelve a procesar un POST con el mismo satélite , para el mismo usuario, y con "distancia" = 200 , se borra el dato anterior y se reemplaza por el nuevo._

_Endporints:_
_ POST /quasar/topsecret_split/{sastelite_name}:_

```
http://144.126.217.32/quasar/topsecret_split/kenobi
http://144.126.217.32/quasar/topsecret_split/skywalker
http://144.126.217.32/quasar/topsecret_split/sato
```







