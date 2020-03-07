<h1>Blog To App</h1>
Blog To app is website about gather some awesome article around the world about career, developement, startup,... And of course its has a goal of myself
App now live in: https://go-blog-app.herokuapp.com (warning: It will load in 20-30s because heroku will turn off service if app no access in 5 minutes)

<h1>Getting Started</h1>
<code>git clone https://github.com/AnhNguyenQuoc/go-blog.git</code>

<h1>Prerequisites</h1>
Blog To app has development by docker. Because of that, you need to install docker step by step in the link below:

<code>https://www.docker.com/products/docker-desktop</code>


<h1>Installing</h1>
<code>cd go-blog</code>

<code>docker-compose up --build</code>

After docker-compose running. You need to check IP of postgresql to config IP server in .env:

<code>docker inspect full_db_postgres</code>

Config your IP, PORT, USERNAME, PASSWORD and DB_NAME to run database for app.

DB_HOST=172.26.0.3 #change it everytime start docker-compose to correct IP host

DB_DRIVER=postgres

DB_USER=admin

DB_PASSWORD=admin

DB_NAME=go_blog

DB_PORT=5432

Config your PORT in env to running Blog To App in this port. For example run app in localhost:3000:

PORT=3000

<code>Open your browser and go to localhost:${YOUR_PORT_ENV}</code>

<h1>Built With</h1>
Golang language
Gorm - Used to database
httprouter - Used to router

<h1>Versioning</h1>
App now in beta with only register login and todo app

<h1>Authors</h1>
Nguyen Quoc Anh - Initial work
