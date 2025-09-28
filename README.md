<div>
    <h2><bold>Proje Kurulumu</bold></h2>
    <br>
    <h4>Docker Kurulumu : </h4>
    <br>
    <span>Bilgisayara önce docker compose indirilmesi gerekli. Daha sonra;</span>
    <br>
    <code>docker run -d --hostname rabbitmq --name rabbitmq -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management</p></code>
    <br>
    <h4>Swagger Kurulumu : </h4>
    <br>
    <span>Bu işlemin IDE terminalin üzerinden yapılması gerekli!!</span>
    <code>go install github.com/swaggo/swag/cmd/swag@latest</code>
    <br>
    <h2><bold>Swagger Init İçin</bold></h2>
    <br>
    <span>Hangi controller için init edilmek isteniyorsa öncelikle onun path'ine girilmesi gerekli. Örneğin (internal\user)</span>
    <br>
    <span>Daha sonra terminal üzerinden<code> swag init -g controller.go</code> yapılmalı.</span>
    <br>
    <span>API ve Consumer configurationları: https://imgur.com/a/jrMEwTI</span>
</div>