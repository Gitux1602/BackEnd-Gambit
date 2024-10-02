# Gambit Backend 🛒
Gambit es el backend de un sistema de ecommerce especializado en tecnología. Este proyecto se encarga de gestionar todas las operaciones principales de la plataforma, tales como la autenticación de usuarios, gestión de productos, stock, direcciones de envío, categorías, y órdenes de compra. Está desarrollado con Go y utiliza AWS Lambda para la ejecución de funciones serverless.

🛠️ Tecnologías Utilizadas
- Lenguaje: Go
- Plataforma: AWS
- Base de datos: MySQL
- Autenticación: JWT

🌐 API Endpoints

Método	Ruta	Descripción
GET	/users	Obtiene todos los usuarios
POST	/users	Crea un nuevo usuario
GET	/products	Obtiene todos los productos
POST	/products	Crea un nuevo producto
PUT	/products/{id}	Actualiza un producto específico
DELETE	/products/{id}	Elimina un producto
GET	/orders	Consulta las órdenes de compra
POST	/orders	Crea una nueva orden de compra
