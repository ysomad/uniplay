upstream ui_node {
    server 127.0.0.1:3000;
}

upstream public_api {
    server 127.0.0.1:4433;
}

upstream admin_api {
    server 127.0.0.1:4434;
}

server {
    server_name dev.uni-play.ru;

    location /kratos {
        rewrite ^/kratos/(.*)$ /$1 break;

        proxy_pass http://public_api;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_set_header host $host;
        proxy_set_header x-real-ip $remote_addr;
        proxy_set_header x-forwarded-for $proxy_add_x_forwarded_for;
    }

    location /admin {
        rewrite /admin/(.*) /$1  break;

        set $allow 0;

        if ($remote_addr ~* "172.24.0.*") {
                set $allow 1;
        }

        if ($arg_secret = "GuQ8alL2") {
                set $allow 1;
        }

        if ($allow = 0) {
                return 403;
        }

        proxy_pass http://admin_api;
        proxy_redirect          off;
        proxy_set_header        Host            $host;
        proxy_set_header        X-Real-IP       $remote_addr;
        proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location /identities {
        proxy_pass http://admin_api;
        proxy_redirect          off;
        proxy_set_header        Host            $host;
        proxy_set_header        X-Real-IP       $remote_addr;
        proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location /auth {
        rewrite /auth/(.*) /$1  break;

        proxy_pass http://ui_node;
        proxy_redirect          off;
        proxy_set_header        Host            $host;
        proxy_set_header        X-Real-IP       $remote_addr;
        proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    error_page 401 = @error401;

    location @error401 {
        return 302 https://dev.uni-play.ru/auth/login;
    }
}
