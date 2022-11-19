网盘 web 端

npm install 下载依赖包

npm run serve 打包

生产环境可以使用nginx进行转发配置，打包后根目录下会生成文件夹 dist，将 dist 文件夹下的文件放置于 nginx/html 目录下，并配置 nginx/conf/nginx.conf，具体配置如下：

```conf
server {
        listen       80;
		server_name localhost;  

        location / {
            root   html;
            index  index.html index.htm;
			try_files	$uri $uri/ /index.html; 
        }

		location /api/{
			#proxy_set_hearder host                $host;
			#proxy_set_header X-forwarded-for $proxy_add_x_forwarded_for;
			#proxy_set_header X-real-ip           $remote_addr;

			# 配置此处用于获取客户端的真实IP
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
			proxy_pass	http://localhost:8080/;
		}

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
```

修改完后重启 Nginx 服务器即可