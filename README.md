
# Использования Github для OAuth

### Требования

Чтобы создать такое же OAuth приложение, нужно создать новое OAuth App на Github. 

![APP!](images/app.png)

***

1. Склонируйте репозиторий на вашу локальную машину:

   ```bash
   git clone https://github.com/Jhnvlglmlbrt/oauth

2. Перейдите в директорию проекта:

   ```bash
   cd basic-auth

3. Установить зависимости:
    
    ```bash
    go get 

4. Нужно создать свой .env файл в /oauth/ :

    ```bash
    CLIENT_ID=xxxxxxxxxxxxxxxxxxxx
    CLIENT_SECRET=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx


5. Убедитесь, что в случае, если создавали свои url, что поменяли их в коде.

6. Запустите код: 

    ```bash
    cd cmd && go run main.go

***
### Структура приложения

![DIAG!](images/diag.png)

1. Клиент (фронтенд) переходит на маршрут, предназначенный для входа через GitHub - '/login/github'.

2. Когда этот маршрут запрашивается у сервера (GET-запрос к серверу), сервер перенаправляет пользователя на специальный маршрут на github.com, где он может авторизовать приложение.

Маршрут имеет следующий вид: 

https://github.com/login/oauth/authorize?client_id=GITHUBCLIENTID&redirect_uri=/login/github/callback. 

Здесь пользователь входит в GitHub (если ещё не вошел) и на следующей странице разрешает приложению доступ к своей информации.

3. После того как пользователь разрешил, API GitHub отправляет запрос на redirect_uri, указанный в предыдущем запросе (/login/github/redirect в данном случае). 

В этом запросе мы получаем authorization code в качестве параметра запроса. Мы извлекаем этот code. Он будет использоваться позже для получения данных пользователя.

4. Затем мы отправляем POST-запрос к конечной точке https://github.com/login/oauth/access_token. 

В качестве параметров запроса мы отправляем наш идентификатор клиента (client id) и клиентский секрет (client secret), которые мы создали в нашей учетной записи GitHub, а также код, который мы получили на предыдущем этапе. На этом этапе мы можем быть уверены, что возвращенные данные будут действительными, так как они поступают непосредственно от GitHub.

5. Вышеуказанный запрос возвращает ответ следующего вида: access_token=e72e16c7e42f292c6912e7710c838347ae178b4a&token_type=bearer. 

Сервер должен извлечь access_token из этого ответа и использовать его для финального запроса на получение данных.

6. Теперь сервер отправляет GET-запрос к конечной точке https://api.github.com/user. С заголовком авторизации, например: Authorization: token e72e16c7e42f292c6912e7710c838347ae178b4a.

7. Ответ на этот запрос - это данные пользователя. 

